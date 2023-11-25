#include <stdio.h>
#include <stdarg.h>

#include <Arduino.h>
#include <WiFi.h>
#include <WebSocketClient.h>
#include "lwip/ip4_addr.h"

#define xstr(s) str(s)
#define str(s) #s

const char* ssid = xstr(WIFI_SSID);
const char* pass = xstr(WIFI_PASS);

char path[] = "/smart-home";
char host[] = "192.168.1.102"; // dev server
 
WebSocketClient webSocketClient;
WiFiClient client;
 
void setup() {
  Serial.begin(115200);
 
  // WIFI
  Serial.printf("INFO :: WIFI :: Connecting to %s\n", ssid);
  WiFi.begin(ssid, pass);
  int i = 1;
  while (WiFi.status() != WL_CONNECTED) {
    delay(1000);
    Serial.printf("INFO :: WIFI :: Attempt %d\n", i);
    i++;
  }
  Serial.printf("INFO :: WIFI :: IP address: ");
  Serial.println(WiFi.localIP());
 
  // WEBSOCKET
  delay(5000);
  if (client.connect(host, 8087)) {
    Serial.println("INFO :: WebSocket :: Connected");
  } else {
    Serial.println("ERROR :: WebSocket :: Connection failed");
  }
 
  webSocketClient.path = path;
  webSocketClient.host = host;
  if (webSocketClient.handshake(client)) {
    Serial.println("INFO :: WebSocket :: Handshake successful");
  } else {
    Serial.println("INFO :: WebSocket :: Handshake failed");
  }
 
}
 
void loop() {
  String data;
 
  if (client.connected()) {
 
    // Sending alive signal
    webSocketClient.sendData("OK");
 
    webSocketClient.getData(data);
    if (data.length() > 0) {
      Serial.print("Received data: ");
      Serial.println(data);
    }
 
  } else {
    Serial.println("ERROR :: WebSocket :: Client disconnected");
  }
 
  delay(3000);
 
}