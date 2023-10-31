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

### Get last X amount record by lamp
GET `/getLast/<lamp>/<amount>`  
Response body:  
`[{"id":<id>, "lamp":"<lamp>", "date":"<date>", "status":<bool>}]`

### Get lamps
GET `/getLamps`
`null` or `["lamp0", "lamp1", ...]` 

### HealthCheck
GET `/hc`
Response body:
`OK` or `NOT_OK`

## Stupid notes
At the start there's no lamp in the frontend only an add new lamp button
'add new' will be always
if you add a new lamp it's name will be stored in an array(also store array in the db)


when we start the app the js ask the backend what lamps do we have
backend return the array
frontend ask for these lamps state

a lamp card be like
```
/----------\
|   name   |
|----------|
|   state  |
|----------|
|  up/down |
\----------/
```
up / down depends on the state (!state)

