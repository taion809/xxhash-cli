package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/cespare/xxhash"
)

func main() {
	var (
		file *os.File
		err  error
	)

	if len(os.Args) < 2 || os.Args[1] == "-" {
		file = os.Stdin
	} else {
		filename := filepath.Clean(os.Args[1])
		file, err = os.Open(filename)
		if err != nil {
			log.Fatalf("error reading file %q, err: %q", filename, err)
		}
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("error reading file %q, err: %q", file.Name(), err)
	}

	sum := xxhash.Sum64(buf)
	fmt.Printf("%s: %d\n", file.Name(), sum)
}
