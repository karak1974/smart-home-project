# Controller

Ugly way to run a dev server
```
cd devserver
ip addr show wlan0 | grep 'inet ' | awk '{print $2}' | rev | cut -c 4- | rev && go run main.go
```

Build && Upload
```
cd SmartHome
WIFI_SSID="name" WIFI_PASS="pass" platformio run --target upload --upload-port /dev/ttyUSB0
```

Monitor
```
platformio device monitor -b 115200
```


