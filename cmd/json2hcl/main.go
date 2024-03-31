package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/winebarrel/json2hcl"
)

func init() {
	log.SetFlags(0)
}

func main() {
	var r io.Reader
	flags := parseFlags()

	if flags.file == "" {
		r = os.Stdin
	} else {
		f, err := os.Open(flags.file)

		if err != nil {
			log.Fatalf("failed to open '%s': %s", flags.file, err)
		}

		defer f.Close()
		r = bufio.NewReader(f)
	}

	b, err := json2hcl.UnmarshalFrom(r)

	if err != nil {
		log.Fatalf("failed to unmarshal JSON: %s", err)
	}

	fmt.Println(string(b))
}
