# aurora-go
Aurora Inverter Communication

Server application written in GoLang that let you get information from PowerOne Aurora series Inverters.
PowerOne Aurora PVI series (and other compatible models) inverters utilise a proprietary communication protocol over the standard RS-485 bus (3-wire D+/D-/GND). Implementing a communication library for these devices with such a protocol isn't easy without a proper document (and google can't help you in this case)

To reduce your work you can use this application or try to understand the protocoll reading the code.

You can retrieve data in XML, JSON and use them for building Android/IOS apps that monitor a remote inverter (for example).

It's based on Daniele De Santis PHP InverterPowerMeterLITE monitor project

http://www.desantix.it/index.php?page=download

###Usage:

Run application with default values with

./aurora-go

Or you can configure the server with arguments:

  -p=1470: Inverter Port
  
  -r="192.168.0.190": Inverter IP
  
  -s=8100: Server Listening Port

####Example: 

#####Run:
./aurora-go -r=192.168.1.133 -s=80

#####Output:

Inverter IP:PORT : 192.168.0.133:1470

Simple Data URL : http://localhost:80/

Json Data URL : http://localhost:80/json/

XML Data URL : http://localhost:80/xml/


