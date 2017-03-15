package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

/**
Quick throwaway script to escape JSON from stdin to stdout
*/
func main() {
	infile := os.Stdin
	defer infile.Close()
	inbytes, err := ioutil.ReadAll(infile)
	handleErr(err)
	outbytes, err := json.Marshal(string(inbytes))
	handleErr(err)
	fmt.Println(string(outbytes))
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
