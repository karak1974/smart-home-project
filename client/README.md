# Controller

Ugly way to run a dev server
```
ip addr show wlan0 | grep 'inet ' | awk '{print $2}' | rev | cut -c 4- | rev && cd devserver && go run main.go
```




