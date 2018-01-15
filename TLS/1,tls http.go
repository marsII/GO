package main

import (
    "crypto/tls"
    "io/ioutil"
    "log"
    "net/http"
    "crypto/x509"
	"fmt"
)

func main() {
    caCert, err := ioutil.ReadFile("GeoTrust_Global_CA.pem")
    if err != nil {
        log.Fatal(err)
    }
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

    client := &http.Client{
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                RootCAs:      caCertPool,
            },
        },
    }

    resp , err1 := client.Get("https://www.google.com")
    if err1 != nil {
        panic(err1)
    }
	
	
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Print(string(body))
}

//main source
//https://stackoverflow.com/questions/38822764/how-to-send-a-https-request-with-a-certificate-golang