// 检查网站的域名证书有效期
package httptlsexpiry

import (
	"fmt"
	"net/http"
)

func httpGet(url string) {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	chains := resp.TLS.VerifiedChains
	for _, chain := range chains {
		for _, cert := range chain {
			// certs.

			// 判断是否为CA机构
			if cert.IsCA {
				fmt.Printf("cert.IsCA , cert.Subject.CommonName=%s\n", cert.Subject.CommonName)
				continue
			}

			fmt.Printf("%s\n", cert.Subject)
			fmt.Printf("not after: %s ; not before: %s\n", cert.NotAfter, cert.NotBefore)
			fmt.Printf("\n\n")
		}
	}

}
