## Intro
This project is to demonstrate how to build a web server using Golang.

## Compile
1. Download dependency:
``` shell
$ go get -u github.com/julienschmidt/httprouter
```
2. Download code:
``` shell
$ git clone git@github.com:cjun714/go-webserver-example
```
3. Build:
``` shell
$ cd go-webserver-example
$ go build -o ./main ./src
```

## Run
``` shell
$ cd go-webserver-example
$ ./main 8080 # run server on 8080
```

## Test
Open web browser, access below URLs:
- http://localhost:8080
- http://localhost:8080/hello
- http://localhost:8080/user/get/1
- http://localhost:8080/user/del/1
- http://localhost:8080/sample.html
- http://localhost:8080/json/sample.json



