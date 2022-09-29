package main

import (
	"log"
	"os/exec"
)

func main() {
	cmd:= exec.Command("calc.exe")
	err:= cmd.Run()

	if(err!=nil){
		log.Fatal(err)
	}
}