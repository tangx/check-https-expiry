package backend

import (
	"bufio"
	"io"
	"log"
	"os"
)

func DomainsFromFile(path string) (result []string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	br := bufio.NewReader(f)
	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		result = append(result, string(line))
	}

	return result
}
