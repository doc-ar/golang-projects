# Introduction

This project is a very simple one. Convert a string to morse code. Whitespaces do not have a designated morse code. In this instance the spaces `" "` are converted to `<SP>` to easily understand the output. Each character is seperated by a normal space.
The codes were taken from a pdf on the [US Navy Website](https://www.history.navy.mil/content/dam/museums/undersea/education/Morse%20Code.pdf) which contains more information on morse code.

#### Example

The text `the quick brown fox` will be translated as

```
- .... . <SP> --.- ..- .. -.-. -.- <SP> -... .-. --- .-- -. <SP> ..-. --- -..-
```

# How to run

Clone this repository anywhere in your machine. `cd` into the go-morsecode directory. You can either build this code or run it directly using go. by default the app will run in localhost with port 8000 but you can change it for your needs.

```
go run main.go
```

```
go build main.go
./main
```

> If you are using Windows machine building will create an exe file so run it using `./main.exe`

# Learnings

- The standard library package `os` is used to work with os commands and files. I used the function `os.ReadFile()` for reading the data in the json file
- The standard library package `encoding/json` is used to handle json data. For parsing the json data I used `json.Unmarshal()` function.
- The `fmt.scanf()` function uses whitespace as a delimeter. So spaces in a string are converted to `"\n"`. To counter this issue we have to used our own reader using the `bufio` package. By creating a new reader using `bufio.NewReader(os.Stdin)` we can read strings with spaces.

# Time Complexity

Right now the morse code is stored as an array of structs. So searching for an element takes `O(n)` time. The array size is 37 so it takes `O(37)`.
An improvement can be made to get morse code in constant time `O(1)` by converting the array of structs into a hashmap of key value pairs.

# Todo

- [ ] Use hashmap instead of array for constant time
- [ ] Create a webapp using HTMX
