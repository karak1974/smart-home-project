# Smart Home Project

# Developer documentation only
User documentation will be in the user_documentation.md after the project is done.

## API documentation  
POST `/addRecord`  
Request body:  
`{"lamp":"<lamp_name>", "status":<bool>}`  
Response body:  
`{"id":<id>, "lamp":"<lamp>", "date":"<date>", "status":<bool>}`  

GET `/getRecordById/<ID>`  
Response body:  
`{"id":<id>, "lamp":"<lamp>", "date":"<date>", "status":<bool>}`  

GET `/getRecordByLamp/<lamp>`  
Response body:  
`{"id":<id>, "lamp":"<lamp>", "date":"<date>", "status":<bool>}`  

GET `/getLast`  
Response body:  
`{"id":<id>, "lamp":"<lamp>", "date":"<date>", "status":<bool>}`  

GET `/getLast/<amount>`  
Response body:  
`[{"id":<id>, "lamp":"<lamp>", "date":"<date>", "status":<bool>}]`  
> This is an array, this means it will return an object array instead of a single object

