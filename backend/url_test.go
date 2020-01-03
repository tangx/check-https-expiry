package backend

import "testing"

func Test_URL(t *testing.T) {
	DomainsFromURL("http://127.0.0.1:8080/1.txt")
}
