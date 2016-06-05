# PSN-Go
A PSN API wrapper written in Go.

This is an implementation of my [PSN API wrapper written in PHP](https://github.com/Tustin/psn-php) but written in Golang. This is by no means 
done and is actually my first project written in Go. Please bear with me as I learn the ins and outs!

## Installation
Navigate to your working directory for Go and run this Go command:
```
go get github.com/Tustin/psn-go
```
This will create a new directory tree in your `src` folder with the source.

## Usage
After installing the code, create a new Go file and import the package like so:
``````go
package main

import(
    "github.com/Tustin/psn-go"
    "fmt"
)

func main(){
    oauth, err := psn.Login("myemail@psn.com", "hunter2")
    if err != nil {
        panic(err)
    }
    fmt.Println(oauth)
}
``````

If you use the correct PSN credentials, your OAuth token should get output to the command line. Otherwise the program will panic and output the error that occurred.

**Please note that this wrapper is a WIP and will include a lot more functionality in the future!**
