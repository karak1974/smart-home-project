# Smart Home Project

## API documentation
POST `/addRecord`
Request body:
`{"lamp":"<lamp_name>", "stateus":<bool>}`
Response body:
`{"id":<id>, "lamp":"<lamp>", "date":"<date>", "status":<bool>}`

GET `/getRecordById/<ID>"`
Response body:
`{"id":<id>, "lamp":"<lamp>", "date":"<date>", "status":<bool>}`
