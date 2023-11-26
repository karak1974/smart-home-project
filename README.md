# Smart Home Project
Smarthome thesis project

## Documentations
- [Developer documentation(ENG)](docs/dev_documentation_en.md)
- User documentation(ENG)
- Developer documentation(HUN)
- User documentation(HUN)

## Description
This a remote lamp controller project using [ESP32_Relay_X8](https://templates.blakadder.com/ESP32_Relay_X8.html).  
Client connect to a publicly reachable server and that communicate with the controller via websocket and provide instructions.
Due to the server is public the user have to authenticate with a password, this is managed with haproxy.
