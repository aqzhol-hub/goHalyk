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
- POST **/api/protected/logout**
```json
LOG OUT
```

- POST **/api/protected/addstate**
```json
{
    "Name" : "Online",
    "Description" : "Online"
}
```

- POST **/api/protected/updateState**
```json
Form -> stateID = 7
```

- GET **/api/protected/mystatus**
```json
{
  "status" : "ok"
}
```

- GET **/api/protected/allstates**
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

