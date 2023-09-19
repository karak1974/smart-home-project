# Developer documentation
User documentation will be in the user_documentation.md after the project is done.

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

### Get record by lamp
GET `/getRecordByLamp/<lamp>`  
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
