# Controller

Run dev server
```
cd devserver
ip addr show wlan0 | grep 'inet ' | awk '{print $2}' | rev | cut -c 4- | rev
XOR_KEY=0100010001000100 go run main.go
```
Key explained  
01 00 01 00 01 00 01 00  
go XOR_KEY=0100010001000100    
 c XOR_KEY="10101010"

Build && Upload
```
cd SmartHome
WIFI_SSID="name" WIFI_PASS="pass" XOR_KEY="10101010" platformio run --target upload --upload-port /dev/ttyUSB0
```

Monitor
```
platformio device monitor -b 115200
```


