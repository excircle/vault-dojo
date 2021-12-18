package main

import Conf "vault-dojo/conf"
import "os"
import "log"

var InfoLogger *log.Logger

func init() {
	file, err := os.OpenFile("./dojo.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "[INFO] ", log.LstdFlags)
}

func main() {
	status := Conf.Check(configFile)
	InfoLogger.Println(status)
}