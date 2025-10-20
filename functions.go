/**********************************************************************
Dev:                    Federico Bertolini
Email:                  golife@paranoici.org
Project:                Aurora-go
Version:                0.2
Date:                   20/10/2025
Note:                   Go 1.22; TCP with deadlines and retry/backoff
**********************************************************************/
package main

import (
    "context"
    "fmt"
    "io"
    "math"
    "net"
    "time"
)

func DSPValue(hex string) float64 {
	v := (int(hex[2]) << 24) + (int(hex[3]) << 16) + (int(hex[4]) << 8) + int(hex[5])
	x := (v & ((1 << 23) - 1)) + (1<<23)*(v>>31|1)
	exp := (v >> 23 & 0xFF) - 127
	return float64(float64(x) * (math.Pow(2, float64(exp-23))))
}

func CRC(cmd []byte) int {
	BccLo := 0xFF
	BccHi := 0xFF
	for x := 0; x < len(cmd); x++ {
		c := int(cmd[x]) ^ BccLo
		tmp := (c << 4) & 0xFF
		c = tmp ^ c
		tmp = (c >> 5) & 0xFF
		BccLo = BccHi
		BccHi = c ^ tmp
		tmp = (c << 3) & 0xFF
		BccLo = BccLo ^ tmp
		tmp = (c >> 4) & 0xFF
		BccLo = BccLo ^ tmp
	}
	CRC_L := (-(BccLo) - 1) & 0xFF
	CRC_H := (-(BccHi) - 1) & 0xFF
	return ((CRC_H << 8) & 0xFF00) + ((CRC_L) & 0xFF)
}

func CheckCRC(rxdata []byte) bool {
	rx_crc := []byte{rxdata[6], rxdata[7]}
	my_crc := CRC(rxdata[:len(rxdata)-2])
	my_crc_tmp := []byte{byte(my_crc & 0xFF), byte((my_crc >> 8) & 0xFF)}

	if (int(rx_crc[0]) == int(my_crc_tmp[0])) && (int(rx_crc[1]) == int(my_crc_tmp[1])) {
		return true
	} else {
		return false
	}

}

// QueryInverter queries the inverter by sending a batch of commands.
// - Dials with timeout and TCP keepalive
// - Sets write/read deadlines and reads exactly 8 bytes
// - Applies exponential backoff retries to handle standby/power cycles
func QueryInverter(cmdarray, results map[string][]byte) bool {
    // exponential backoff to handle inverter standby/power cycle
    const maxAttempts = 5

    runBatch := func() error {
        // context timeout for the whole dial operation
        ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
        defer cancel()

        dialer := net.Dialer{Timeout: 1 * time.Second, KeepAlive: 10 * time.Second}
        conn, err := dialer.DialContext(ctx, "tcp", REMOTE_IP+":"+REMOTE_PORT)
        if err != nil {
            return fmt.Errorf("dial error: %w", err)
        }
        defer conn.Close()

        for key, value := range cmdarray {
            cmd_crc := CRC(value)
            // CRC values
            a := byte(int(cmd_crc & 0xFF))
            b := byte(int((cmd_crc >> 8) & 0xFF))
            // append CRC
            txcmd := append(value[:], a)
            txcmd = append(txcmd[:], b)

            // write with deadline
            _ = conn.SetWriteDeadline(time.Now().Add(1 * time.Second))
            if _, err := conn.Write(txcmd); err != nil {
                return fmt.Errorf("write error on %s: %w", key, err)
            }

            reply := make([]byte, 8)
            // read exactly 8 bytes with deadline
            _ = conn.SetReadDeadline(time.Now().Add(1 * time.Second))
            if _, err := io.ReadFull(conn, reply); err != nil {
                return fmt.Errorf("read error on %s: %w", key, err)
            }

            // validate response
            if len(reply) != 8 || !CheckCRC(reply) {
                return fmt.Errorf("CRC error or invalid length on %s", key)
            }

            // strip CRC and store
            r := reply[:len(reply)-2]
            results[key] = r
        }
        return nil
    }

    for attempt := 0; attempt < maxAttempts; attempt++ {
        if err := runBatch(); err != nil {
            results["ERROR"] = []byte(err.Error())
            // minimal logging via fmt to avoid extra deps
            fmt.Println("QueryInverter attempt", attempt+1, "failed:", err)
            time.Sleep(time.Duration(250*(1<<uint(attempt))) * time.Millisecond)
            continue
        }
        return true
    }

    if _, ok := results["ERROR"]; !ok {
        results["ERROR"] = []byte("all attempts failed")
    }
    return false
}
