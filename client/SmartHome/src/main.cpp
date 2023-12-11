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
char host[] = "192.168.43.86"; // dev server
 
WebSocketClient webSocketClient;
WiFiClient client;

void lampController(String input) {
  int length = 8;
  bool binaryArray[length];

  for (int i = 0; i < length; ++i) {
    binaryArray[i] = input[i] - '0';
  }

  for (int i = 0; i < length; ++i) {
    Serial.printf("%d->%d ", i, binaryArray[i]);
  }
  Serial.printf("\n");

  /*
  GPIO12 	Relay 2
  GPIO13 	Relay 1
  GPIO14 	Relay 3
  GPIO25 	Relay 6
  GPIO26 	Relay 5
  GPIO27 	Relay 4
  GPIO32 	Relay 8
  GPIO33 	Relay 7 
  */
}

void setup() {
  Serial.begin(115200);
 
  // Connect to WIFI
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
 
  // Connect to WebSocket
  delay(5000);
  if (client.connect(host, 8087)) {
    Serial.println("INFO :: WebSocket :: Connected");
  } else {
    Serial.println("ERROR :: WebSocket :: Connection failed");
  }
 
  // Create WebSocket handshake
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
      Serial.printf("INFO :: WebSocket :: Received data: %s\n", data);
      lampController(data);
    }
 
  } else {
    Serial.println("ERROR :: WebSocket :: Client disconnected");
  }
 
  // Decrease at release
  delay(5000);
 
}