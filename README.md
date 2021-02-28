# Golang Halyk Project


### PUBLIC REST:
- POST **/api/public/signup**
```json
{
    "FirstName": "Aqzhol",
    "LastName" : "Baqatay",
    "Username" : "aqzhol@gmail.com",
    "Password" : "123"
}
```

- POST **/api/public/login**
```json
{
    "Username" : "aqzhol@gmail.com",
    "Password" : "123"
}
```

### PROTECTED REST:
- POST **/api/public/logout**
```json
LOG OUT
```

- POST **/api/public/addstate**
```json
{
    "Name" : "Online",
    "Description" : "Online"
}
```

- POST **/api/public/updateState**
```json
Form -> stateID = 7
```

- GET **/api/public/mystatus**
```json
{
  "status" : "ok"
}
```

- GET **/api/public/allstates**
```json
{
    "result": [
        {
            "id": 1,
            "name": "Online",
            "description": "Online"
        },
        {
            "id": 2,
            "name": "Offline",
            "description": "Offline"
        }
    ]
}

```

