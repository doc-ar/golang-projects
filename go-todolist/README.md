# Introduction

This is a todo list web application I created only using the go standard library. I learned how to create a basic webserver by following along a [tutorial by Akhil Sharma](https://www.youtube.com/watch?v=ASBUp7stqjo). A key note is that this application does not use any javascript so its very fast in terms of response.

> Note: The css of this project is entirely generated using chatgpt. I have adjusted the generated css slightly to fit my needs.

# How to run

Clone this repository anywhere in your machine. `cd` into the go-todolist directory. You can either build this code or run it directly using go. by default the app will run in localhost with port 8000 but you can change it for your needs.

```
go run .

# Or expose it to network using wildcard ip
go run . 0.0.0.0:8000

```

```
go build .
./go-todolist

# Or expose it to network using wildcard ip
./go-todolist 0.0.0.0:8000
```

> If you are using Windows machine building will create an exe file so run it using `./main.exe`

# API Endpoints

`/`: This renders the static index.html file from the static folder
`/hello`: This endpoint dynamically generates html with the word hello
`/form.html`: This endpoint leads to a basic form
`/form`: The form.html redirects to this endpoint and reveals if the http request was successful or not

```

```
