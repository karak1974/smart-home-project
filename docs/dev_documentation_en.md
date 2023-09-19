# Developer documentation

## Setup

### Backend  
The backend have only one dependency which is [Docker](https://www.docker.com/).  
To start it `docker compose up`  
To start it as a daemon `docker compose up -d`  

## API documentation
### Add record
POST `/addRecord`  
Request body:  
`{"lamp":"<lamp_name>", "status":<bool>}`  
Response body:  
`{"id":<id>, "lamp":"<lamp>", "date":"<date>", "status":<bool>}`

### Get record by ID
GET `/getRecordById/<ID>`  
Response body:  
`{"id":<id>, "lamp":"<lamp>", "date":"<date>", "status":<bool>}`

### Get last record by lamp
GET `/getLastByLamp/<lamp>`  
Response body:  
`{"id":<id>, "lamp":"<lamp>", "date":"<date>", "status":<bool>}`

### Get last record
GET `/getLast`  
Response body:  
`{"id":<id>, "lamp":"<lamp>", "date":"<date>", "status":<bool>}`

### Get last X amount record
GET `/getLast/<amount>`  
Response body:  
`[{"id":<id>, "lamp":"<lamp>", "date":"<date>", "status":<bool>}]`

### Geg last X amount record by lamp
GET `/getLast/<lamp>/<amount>`  
Response body:  
`[{"id":<id>, "lamp":"<lamp>", "date":"<date>", "status":<bool>}]`
