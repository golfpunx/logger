# logger
Golang Logger
An easy way to use my own logging.

## How to use
1. go get 
```sh
$ go get github.com/golfpunx/logger 
```
2. import
```go
import "github.com/golfpunx/logger"
```

## Example
```go
package main

import "github.com/golfpunx/logger"

func main(){
	//No need to name files. We use the date automatically.
	//We need only directory
	//logger.SaveDir("") is optional 
	logger.SaveDir("/tmp/project_name")
	
	//Write the info log
	logger.Info("now created")
	
	//Write the error log
	logger.Error("is exception! ...")
}
```
