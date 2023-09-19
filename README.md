# Smart Home Project

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


