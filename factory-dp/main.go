package main

import "fmt"

type Logger interface {
	log(s string)
}

type XMLLogger struct{}

type JSONLogger struct{}

func newLogger(kind string) Logger {
	if kind == "xml" {
		return XMLLogger{}
	} else if kind == "json" {
		return JSONLogger{}
	} else {
		return nil
	}
}

func (xmlLogger XMLLogger) log(s string) {
	fmt.Printf("<log>%v</log>\n", s)
}

func (jsonLogger JSONLogger) log(s string) {
	fmt.Printf("{ \"log\" : \"%v\" }", s)
}

func main() {
	xmlLogger := newLogger("xml")
	xmlLogger.log("xml")

	jsonLogger := newLogger("json")
	jsonLogger.log("json")

}
