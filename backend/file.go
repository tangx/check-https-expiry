package backend

import (
	"io/ioutil"
	"log"
	"strings"
)

func DomainsFromFile(path string) []string {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("%s", err)
	}

	l := strings.Split(string(body), "\n")
	return l
}
