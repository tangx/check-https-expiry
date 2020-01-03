package main

import (
	"fmt"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {

	URLs := []string{"tangx.in",
		"http://tangx.in",
		"http://www.lcdcf.cn",
		"https://www.baidu.com",
	}

	// URLs := []string{"http://tangx.in"}

	// for _, url := range URLs {
	// 	// fmt.Printf("%s   :   ", url)
	// 	cmd.Do()
	// }
}

func Test_split(t *testing.T) {
	url := `https://tangx.in:80/`
	l := strings.Split(url, ":")
	fmt.Println(l)

}
