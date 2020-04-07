# Warpin

- API for sending a message Just send one parameter string for message After sending should be get response
- API for collect message that has been sent out API can get all previously sent messages
- API for display message in real time API should be long live connection to retrieve messages after send at realtime

## Requirements
- go 1.13 above (with go modules)
- sqlite DB (default in os)

## Pattern
- golang standard project layout
https://github.com/golang-standards/project-layout
- clean architecture by unclebob
https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

## Packages used
- github.com/go-playground/locales v0.13.0
- github.com/go-playground/universal-translator v0.17.0
- github.com/go-playground/validator/v10 v10.2.0
- github.com/gorilla/websocket v1.4.2
- github.com/jinzhu/gorm v1.9.12
- github.com/joho/godotenv v1.3.0
- github.com/labstack/echo/v4 v4.1.16
- github.com/sirupsen/logrus v1.5.0
- github.com/spf13/cast v1.3.1

## How to run
copy file `.env.example` to `.env`
```bash
$ cp .env.example .env
```

run the main file
```bash
$ go run main.go
```

or create binary file, then run it
```bash
$ go build main.go
$ ./main
```

import postman files in `api` folder:
```bash
$ ls -la api/
Warpin.postman_collection.json   Warpin.postman_environment.json
```

open your postman, and endpoint request in `insert` using `POST` and `findall` using `GET` (default api starts in [http://localhost:1234](http://localhost:1234 "http://localhost:1234"))

websocket, point your browser to [http://localhost:1234/chats/room](http://localhost:1234/chats/room "http://localhost:1234/chats/room")
