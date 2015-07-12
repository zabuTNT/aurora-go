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
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

func jsonFunc(w http.ResponseWriter, r *http.Request) {

	m := getInformation()

	b, _ := json.Marshal(m)
	out := string(b[:])

	fmt.Fprint(w, out)

}

func txtFunc(w http.ResponseWriter, r *http.Request) {

	m := getInformation()
	out := ""
	if m["state"] != "OK" {
		out = "State: " + m["state"] + "\n" +
			"Error: " + m["code"]

	} else {
		out = "State: " + m["state"] + "\n" +
			"Model: " + m["model"] + "\n" +
			"Serial: " + m["serial"] + "\n" +
			"Global State: " + m["global"] + "\n" +
			"Inverter State: " + m["inverter"] + "\n" +
			"Alarm: " + m["alarm"]
		if m["alarm_code"] != "" {
			out += " (Code: " + m["alarm_code"] + ")"
		}

		out += "\n" +
			"Power out: " + m["pw_out"] + " W\n" +
			"Power in: " + m["pw_in"] + " W\n" +
			"Total Production: " + m["total"] + " kWh\n" +
			"Today Production:" + m["today"] + " kWh"
	}

	fmt.Fprint(w, out)

}

func xmlFunc(w http.ResponseWriter, r *http.Request) {

	m := getInformation()

	i := Inverter{
		m["state"],
		m["model"],
		m["serial"],
		m["global"],
		m["inverter"],
		m["alarm"],
		m["alarm_code"],
		m["pw_out"],
		m["pw_in"],
		m["total"],
		m["today"],
		m["code"],
	}

	b, e := xml.Marshal(i)

	if e != nil {
		fmt.Println(e)
	}

	out := string(b[:])

	fmt.Fprint(w, out)

}
