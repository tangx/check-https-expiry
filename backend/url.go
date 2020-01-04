package backend

import (
	"bufio"
	"io"
	"log"
	"net/http"
)

// func DomainsFromURL(url string) []string {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatalf("%s", err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatalf("%s", err)
// 	}

// 	l := strings.Split(string(body), "\n")
// 	// fmt.Println(l)
// 	return l
// }

func DomainsFromURL(url string) (result []string) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer resp.Body.Close()

	br := bufio.NewReader(resp.Body)

	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		result = append(result, string(line))
	}

	return result
}
