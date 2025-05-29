package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace		*log.Logger		// Just about anything
	Info		*log.Logger		// Important Information
	Warning		*log.Logger		// Warning
	Error		*log.Logger		// Critical problem
)

func init() { 
	file, err := os.OpenFile("errors.txt",
		os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	
	Trace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate | log.Lmicroseconds | log.Llongfile)

	Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate | log.Lmicroseconds | log.Llongfile)

	Warning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate | log.Lmicroseconds | log.Llongfile)

	Error = log.New(io.MultiWriter(file, os.Stderr),
		"ERROR: ",
		log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {

	// Standard Logger
	Trace.Println("A standard log message")
	Info.Println("Special Information")
	Warning.Println("Be Aware")
	Error.Println("An Error Message from somewhere")

}