package cmd

import (
	"crypto/x509"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/tangx/check-https-expiry/backend"
)

var (
	URLFlag  string
	FileFlag string
)

func init() {
	flag.StringVar(&URLFlag, "url", "", "read domain list from a url")
	flag.StringVar(&FileFlag, "file", "", "read domain list from a file")
}

func Do(params []string) {

	flag.Parse()

	var domainList []string
	if URLFlag != "" {
		list := backend.DomainsFromURL(URLFlag)
		domainList = append(domainList, list...)
	}

	if FileFlag != "" {
		list := backend.DomainsFromFile(FileFlag)
		domainList = append(domainList, list...)
		fmt.Println(len(domainList))
	}

	// domainList = append(domainList, params...)
	fmt.Println(domainList)
	for _, url := range domainList {
		// fmt.Printf("domain url is %v////", url)
		check(parse(url))
	}

}

func parse(hostOrURL string) string {
	// https;//tangx.in:443/
	if strings.HasPrefix(hostOrURL, "http://") || strings.HasPrefix(hostOrURL, "https://") {
		return hostOrURL
	}

	return fmt.Sprintf("%s%s", "https://", hostOrURL)
}

func check(url string) (cert *x509.Certificate) {
	// resp, err := http.Get(url)
	resp, err := httpGet(url)

	if err != nil {
		// panic(err)
		log.Fatalf("%s", err)
	}
	defer resp.Body.Close()

	if resp.TLS == nil {
		fmt.Printf("%-20s: has no tls support\n", url)

		return
	}

	chains := resp.TLS.VerifiedChains
	for _, certs := range chains {
		for _, cert := range certs {
			// 这里会返回多个证书，其中部分是 CA 机构的证书，不在检查范围之列
			if cert.IsCA {
				continue
			}

			fmt.Printf("%-20s: ", url)
			valid(cert)
		}
	}
	return nil
}

func valid(cert *x509.Certificate) {

	validDate := cert.NotAfter

	// expireDate := validDate.Sub(time.Now())// 距离当前时间，使用 time.Until 更优雅
	expireDate := time.Until(validDate)

	// CommonName, ExpireDate
	format := "Cert CommonName: %s, Expire in: %.f days\n"
	fmt.Printf(format, cert.Subject.CommonName, (expireDate / 24).Hours())

}
