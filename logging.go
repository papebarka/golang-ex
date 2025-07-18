package main

import "log"

func init() { 
	log.SetPrefix("TRACE:" )
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {

	// Standard Logger
	log.Println("message")
	

	//Fatalln is Println() followed by a call to os.Exit(1)
	log.Fatalln("fatal message")

	//Panicln is Println() followed by a call to panic()
	log.Panicln("fatal message")
}