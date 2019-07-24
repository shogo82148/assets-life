// +build ignore

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	for i := 0; i < 256; i++ {
		if err := os.MkdirAll(fmt.Sprintf("bench/%02x", i), 0755); err != nil {
			log.Fatal(err)
		}
		for j := 0; j < 256; j++ {
			err := ioutil.WriteFile(fmt.Sprintf("bench/%02x/%02x.txt", i, j), []byte(fmt.Sprintf("%02x%02x", i, j)), 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
