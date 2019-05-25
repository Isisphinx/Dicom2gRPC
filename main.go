package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	myFile := "testdata/SR/dlp83.63ctdi2.29.sr"
	out := dcm2json(myFile)
	fmt.Println("file :", out)
}

func dcm2json(file string) string {
	out, err := exec.Command("./bin/x64/dcm2json/dcm2json", "-fc", "+fo", "-q", file).Output()
	if err != nil {
		log.Fatal(err)
	}
	outStr := string(out)
	if outStr == "" {
		log.Fatal("dcm2json error : no output")
	}
	return outStr
}

// func dcm2jsonstream(file string) {
// 	cmd := exec.Command("./bin/x64/dcm2json/dcm2json", "-fc", "+fo", "-q", file)

// 	stdout, _ := cmd.StdoutPipe()
// 	cmd.Start()

// 	scanner := bufio.NewScanner(stdout)
// 	scanner.Split(bufio.ScanWords)
// 	for scanner.Scan() {
// 		m := scanner.Text()
// 		fmt.Println(m)
// 	}
// 	cmd.Wait()
// }
