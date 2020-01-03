package backend

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func DomainsFromURL(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%s", err)
	}

	l := strings.Split(string(body), "\n")
	// fmt.Println(l)
	return l
}
