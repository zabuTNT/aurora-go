/**********************************************************************
Dev:                    Federico Bertolini
Email:                  golife@paranoici.org
Project:                Aurora-go
Version:                0.1
Date:                   12/07/2015
Note:          		   Go 1.4.2
**********************************************************************/
package aurora_go

import (
	_ "fmt"
	_ "github.com/howeyc/crc16"
	"math"
	"net"
	_ "strconv"
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

func QueryInverter(cmdarray, results map[string][]byte) bool {

	//fmt.Println("Connection with " + DEFAULT_IP + " ...")

	seconds := 1
	wait := time.Duration(seconds) * time.Second

	fp, err := net.DialTimeout("tcp", REMOTE_IP+":"+REMOTE_PORT, wait)

	if err != nil {
		results["ERROR"] = []byte(err.Error())
		return false
	}

	defer fp.Close()

	for key, value := range cmdarray {

		//Do you want third party functions?
		//cmd_crc := crc16.ChecksumCCITT(value)

		//or you can use my function...
		cmd_crc := CRC(value)

		//CRC values
		a := byte(int(cmd_crc & 0xFF))
		b := byte(int((cmd_crc >> 8) & 0xFF))

		//Append CRC
		txcmd := append(value[:], a)
		txcmd = append(txcmd[:], b)

		//Send on socket
		fp.Write(txcmd)

		reply := make([]byte, 8)

		//Read from socket
		fp.Read(reply)

		rxbuff := string(reply)
		if len(rxbuff) == 8 {

			//Check CRC
			if CheckCRC(reply) {
				reply = reply[:len(reply)-2]

				results[key] = reply

			} else {
				results["ERROR"] = []byte("CRC Error")
				return false
			}
		}
	}

	return true
}
