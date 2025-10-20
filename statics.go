/**********************************************************************
Dev:                    Federico Bertolini
Email:                  golife@paranoici.org
Project:                Aurora-go
Version:                0.2
Date:                   20/10/2025
Note:                   Go 1.22; constant tables and commands unchanged
**********************************************************************/
package main

type Inverter struct {
	State      string
	Date       string
	Model      string
	Serial     string
	Global     string
	Inverter   string
	Alarm string
	Alarm_code string
	Pw_out     string
	Pw_in      string
	Total      string
	Today      string
	Code       string
}

var CMD_SN = []byte{
	byte(int(DEFAULT_ADDR)),
	byte(int(0x3f)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
}

var CMD_ST = []byte{
	byte(int(DEFAULT_ADDR)),
	byte(int(0x32)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
}

var CMD_CET = []byte{
	byte(int(DEFAULT_ADDR)),
	byte(int(0x4e)),
	byte(int(0x05)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
}

var CMD_CED = []byte{
	byte(int(DEFAULT_ADDR)),
	byte(int(0x4e)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
}

var CMD_DSP3 = []byte{
	byte(int(DEFAULT_ADDR)),
	byte(int(0x3b)),
	byte(int(0x03)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
}

var CMD_DSP23 = []byte{
	byte(int(DEFAULT_ADDR)),
	byte(int(0x3b)),
	byte(int(0x17)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
}

var CMD_DSP25 = []byte{
	byte(int(DEFAULT_ADDR)),
	byte(int(0x3b)),
	byte(int(0x19)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
}

var CMD_DSP26 = []byte{
	byte(int(DEFAULT_ADDR)),
	byte(int(0x3b)),
	byte(int(0x1A)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
}

var CMD_DSP27 = []byte{
	byte(int(DEFAULT_ADDR)),
	byte(int(0x3b)),
	byte(int(0x1B)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
}

var CMD_VR = []byte{
	byte(int(DEFAULT_ADDR)),
	byte(int(0x3A)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
	byte(int(0x00)),
}

var cmdarray = map[string][]byte{

	"ST":    CMD_ST,
	"CET":   CMD_CET,
	"CED":   CMD_CED,
	"DSP3":  CMD_DSP3,
	"DSP23": CMD_DSP23,
	"DSP25": CMD_DSP25,
	"DSP26": CMD_DSP26,
	"DSP27": CMD_DSP27,
	"SN":    CMD_SN,
	"VR":    CMD_VR,
}

var VersionPart1_Label = map[string]string{
	"i": "Aurora 2 kW indoor",
	"o": "Aurora 2 kW outdoor",
	"I": "Aurora 3.6 kW indoor",
	"O": "Aurora 3.0-3.6 kW outdoor",
	"5": "Aurora 5.0 kW outdoor",
	"6": "Aurora 6 kW outdoor",
	"P": "3-phase interface (3G74)",
	"C": "Aurora 50kW module",
	"4": "Aurora 4.2kW new",
	"3": "Aurora 3.6kW new",
	"2": "Aurora 3.3kW new",
	"1": "Aurora 3.0kW new",
	"D": "Aurora 12.0kW",
	"X": "Aurora 10.0kW",
}

var VersionPart2_Label = map[string]string{
	"A": "UL1741",
	"E": "VDE0126",
	"S": "DR 1663/2000",
	"I": "ENEL DK 5950",
	"U": "UK G83",
	"K": "AS 4777",
}

var VersionPart3_Label = map[string]string{
	"N": "Transformerless",
	"T": "Transformer",
}

var VersionPart4_Label = map[string]string{
	"W": "Wind",
	"N": "PV",
}

var GlobalState_Label = [...]string{"Sending Parameters", "Wait Sun/Grid", "Checking Grid", "Measuring Riso",
	"DcDc Start", "Inverter", "Run ", "Recovery", "Pause", "Ground Fault",
	"OTH Fault", "Address Setting", "Self Test", "Self Test Fail", "Sensor Test + Meas.Riso",
	"Leak Fault", "Waiting for manual reset",
	"Internal Error E026", "Internal Error E027", "Internal Error E028", "Internal Error E029", "Internal Error E030",
	"Sending Wind Table", "Failed Sending table", "UTH Fault", "Remote OFF", "Interlock Fail",
	"Executing Autotest", "Waiting Sun", "Temperature Fault", "Fan Staucked",
	"Int. Com. Fault", "Slave Insertion", "DC Switch Open", "TRAS Switch Open", "MASTER Exclusion",
	"Auto Exclusion", "Erasing Internal EEprom", "Erasing External EEpro", "Counting EEprom", "Freeze"}

var DCDCState_Label = [...]string{"DcDc OFF", "Ramp Start", "MPPT", "Not Used",
	"Input OC", "Input UV", "Input OV", "Input Low", "No Parameters", "Bulk OV", "Communication Error", "Ramp Fail", "Internal Error",
	"Input mode Error", "Ground Fault", "Inverter Fail", "DcDc IGBT Sat", "DcDc ILEAK Fail", "DcDc Grid Fail", "DcDc Comm. Error"}

var InverterState_Label = [...]string{"Stand By", "Checking Grid", "Run", "Bulk OV", "Out OC",
	"IGBT Sat", "Bulk UV", "Degauss Error", "No Parameters", "Bulk Low", "Grid OV",
	"Communication Error", "Degaussing", "Starting", "Bulk Cap Fail", "Leak Fail",
	"DcDc Fail", "Ileak Sensor Fail", "SelfTest: relay inverter", "SelfTest: wait for sensor test",
	"SelfTest: test relay DcDc + sensor", "SelfTest: relay inverter fail", "SelfTest timeout fail",
	"SelfTest: relay DcDc fail", "Self Test 1", "Waiting self test start", "Dc Injection", "Self Test 2",
	"Self Test 3", "Self Test 4", "Internal Error", "Internal Error", "", "", "", "", "", "", "", "", "Forbidden State",
	"Input UC", "Zero Power", "Grid Not Present", "Waiting Start", "MPPT", "Grid Fail", "Input OC"}

var AlarmState_Label = [...][2]string{
	[...]string{"No Alarm", ""},
	[...]string{"Sun Low", "W001"},
	[...]string{"Input OC", "E001"},
	[...]string{"Input UV", "W002"},
	[...]string{"Input OV", "E002"},
	[...]string{"Sun Low", "W001"},
	[...]string{"No Parameters", "E003"},
	[...]string{"Bulk OV", "E004"},
	[...]string{"Comm.Error", "E005"},
	[...]string{"Output OC", "E006"},
	[...]string{"IGBT Sat", "E007"},
	[...]string{"Bulk UV", "W011"},
	[...]string{"Internal error", "E009"},
	[...]string{"Grid Fail", "W003"},
	[...]string{"Bulk Low", "E010"},
	[...]string{"Ramp Fail", "E011"},
	[...]string{"Dc/Dc Fail", "E012"},
	[...]string{"Wrong Mode", "E013"},
	[...]string{"Ground Fault", "---"},
	[...]string{"Over Temp.", "E014"},
	[...]string{"Bulk Cap Fail", "E015"},
	[...]string{"Inverter Fail", "E016"},
	[...]string{"Start Timeout", "E017"},
	[...]string{"Ground Fault", "E018"},
	[...]string{"Degauss error", "---"},
	[...]string{"Ileak sens.fail", "E019"},
	[...]string{"DcDc Fail", "E012"},
	[...]string{"Self Test Error 1", "E020"},
	[...]string{"Self Test Error 2", "E021"},
	[...]string{"Self Test Error 3", "E019"},
	[...]string{"Self Test Error 4", "E022"},
	[...]string{"DC inj error", "E023"},
	[...]string{"Grid OV", "W004"},
	[...]string{"Grid UV", "W005"},
	[...]string{"Grid OF", "W006"},
	[...]string{"Grid UF", "W007"},
	[...]string{"Z grid Hi", "W008"},
	[...]string{"Internal error", "E024"},
	[...]string{"Riso Low", "E025"},
	[...]string{"Vref Error", "E026"},
	[...]string{"Error Meas V", "E027"},
	[...]string{"Error Meas F", "E028"},
	[...]string{"Error Meas Z", "E029"},
	[...]string{"Error Meas Ileak", "E030"},
	[...]string{"Error Read V", "E031"},
	[...]string{"Error Read I", "E032"},
	[...]string{"Table fail", "W009"},
	[...]string{"Fan Fail", "W010"},
	[...]string{"UTH", "E033"},
	[...]string{"Interlock fail", "E034"},
	[...]string{"Remote Off", "E035"},
	[...]string{"Vout Avg errror", "E036"},
	[...]string{"Battery low", "W012"},
	[...]string{"Clk fail", "W013"},
	[...]string{"Input UC", "E037"},
	[...]string{"Zero Power", "W014"},
	[...]string{"Fan Stucked", "E038"},
	[...]string{"DC Switch Open", "E039"},
	[...]string{"Tras Switch Open", "E040"},
	[...]string{"AC Switch Open", "E041"},
	[...]string{"Bulk UV", "E042"},
	[...]string{"Autoexclusion", "E043"},
	[...]string{"Grid df/dt", "W015"},
	[...]string{"Den switch Open", "W016"},
	[...]string{"box fail", "W017"}}
