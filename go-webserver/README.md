# Introduction

This project is a very bare bones webserver that is created using the go standard library. I learned how to created this server by following along a [tutorial by Akhil Sharma](https://www.youtube.com/watch?v=ASBUp7stqjo).

# How to run

Clone this repository anywhere in your machine. `cd` into the go-webserver directory. You can either build this code or run it directly using go

```
go run .main.go
```

```
go build main.go
./main
```

> If you are using Windows machine building will create an exe file so run it using `./main.exe`

# API Endpoints

`/`: This renders the static index.html file from the static folder
`/hello`: This endpoint dynamically generates html with the word hello
`/form.html`: This endpoint leads to a basic form
`/form`: The form.html redirects to this endpoint and reveals if the http request was successful or not
