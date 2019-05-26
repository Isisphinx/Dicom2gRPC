package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	myFile := "test/data/SR/dlp83.63ctdi2.29.sr"
	out := dcm2json(myFile)
	fmt.Println("file :", out)
}

func dcm2json(file string) string {
	out, err := exec.Command("./third_party/x64/dcmtk/dcm2json/dcm2json", "-fc", "-q", file).Output()
	if err != nil {
		log.Fatal(err)
	}
	outStr := string(out)
	if outStr == "" {
		log.Fatal("dcm2json error : no output")
	}
	return outStr
}
