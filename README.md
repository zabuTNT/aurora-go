# aurora-go
Aurora Inverter Communication

Simple server application written in GoLang that let you get information from PowerOne Aurora series Inverters.

You can retrieve data in XML, JSON and use them for building Android/IOS apps that monitoring a remote inverter (for example).

It's based on Daniele De Santis PHP project:

InverterPowerMeterLITE monitor

http://www.desantix.it/index.php?page=download

#Usage:

Run application with default values with

./aurora-go

Or you can configure the server with arguments:

  -p=1470: Inverter Port
  
  -r="192.168.0.190": Inverter IP
  
  -s=8100: Server Listening Port

Example: 
./aurora-go -r=192.168.1.133 -s=80

Inverter IP:PORT : 192.168.0.133:1470

Simple Data URL : http://localhost:80/

Json Data URL : http://localhost:80/json/

XML Data URL : http://localhost:80/xml/


