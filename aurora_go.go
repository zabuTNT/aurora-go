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
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

var REMOTE_IP = "192.168.0.190"
var REMOTE_PORT = "1470"
var DEFAULT_ADDR = 0x02

func getInformation() map[string]string {
	results := make(map[string][]byte)
	m := make(map[string]string)

	m["date"] = time.Now().Format("02/01/2006 15:04")

	if QueryInverter(cmdarray, results) {
		//VERSION
		version := VersionPart1_Label[string(results["VR"][2:3])]
		//GLOBAL STATE
		global_state := GlobalState_Label[results["ST"][2]]
		//INVERTER STATE
		inverter_state := InverterState_Label[results["ST"][3]]
		//ERROR STATE
		alarm_state := AlarmState_Label[results["ST"][5]][0]
		alarm_code := AlarmState_Label[results["ST"][5]][1]
		//TOTAL PRODUCTION
		cet := (int(results["CET"][2]) << 24) + (int(results["CET"][3]) << 16) + (int(results["CET"][4]) << 8) + int(results["CET"][5])
		//DAILY PRODUCTION
		ced := (int(results["CED"][2]) << 24) + (int(results["CED"][3]) << 16) + (int(results["CED"][4]) << 8) + int(results["CED"][5])
		//POWER OUT
		power_out := DSPValue(string(results["DSP3"]))
		//POWER IN
		dsp23 := DSPValue(string(results["DSP23"]))
		dsp25 := DSPValue(string(results["DSP25"]))
		dsp26 := DSPValue(string(results["DSP26"]))
		dsp27 := DSPValue(string(results["DSP27"]))
		power_in := (dsp23 * dsp25) + (dsp26 * dsp27)
		//SERIAL NUMBER
		serial := string(results["SN"])

		m["state"] = "OK"
		m["model"] = version
		m["serial"] = serial
		m["global"] = global_state
		m["inverter"] = inverter_state
		m["alarm"] = alarm_state
		m["alarm_code"] = alarm_code
		m["pw_out"] = strconv.FormatFloat(power_out, 'f', 2, 64)
		m["pw_in"] = strconv.FormatFloat(power_in, 'f', 2, 64)
		m["total"] = strconv.FormatFloat((float64(cet) / 1000), 'f', 2, 64)
		m["today"] = strconv.FormatFloat((float64(ced) / 1000), 'f', 2, 64)

	} else {
		m["state"] = "ERROR"
		m["code"] = string(results["ERROR"])
	}

	return m

}

func main() {

	remote := flag.String("r", "192.168.0.190", "Inverter IP")

	r_port := flag.Int("p", 1470, "Inverter Port")

	s := flag.Int("s", 8100, "Server Listening Port")

	flag.Parse()

	REMOTE_IP = *remote
	REMOTE_PORT = strconv.Itoa(*r_port)

	//sanity checks
	testIP := net.ParseIP(REMOTE_IP)

	if testIP.To4() == nil {
		fmt.Println(REMOTE_IP + " isn't a valid IPv4")
		os.Exit(0)
	}

	if *r_port < 1 || *r_port > 65535 {
		fmt.Println("Remote Port not in valid range (1-65535)")
		os.Exit(0)
	}

	if *s < 1 || *s > 65535 {
		fmt.Println("Local Port not in valid range (1-65535)")
		os.Exit(0)
	}

	//define Handlers
	http.HandleFunc("/", txtFunc)
	http.HandleFunc("/json/", jsonFunc)
	http.HandleFunc("/xml/", xmlFunc)

	fmt.Println("********** AURORA-GO **********")
	fmt.Println("Inverter IP:PORT : " + REMOTE_IP + ":" + REMOTE_PORT)
	fmt.Println("Simple Data URL : http://localhost:" + strconv.Itoa(*s) + "/")
	fmt.Println("Json Data URL : http://localhost:" + strconv.Itoa(*s) + "/json/")
	fmt.Println("XML Data URL : http://localhost:" + strconv.Itoa(*s) + "/xml/")
	fmt.Println("*******************************")

	if err := http.ListenAndServe(":"+strconv.Itoa(*s), nil); err != nil {

		fmt.Println("CRITICAL ERROR: ", err.Error())

	}

}
