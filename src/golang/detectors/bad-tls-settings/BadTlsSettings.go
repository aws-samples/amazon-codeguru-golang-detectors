package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

// {fact rule=insecure-cryptography@v1.0 defects=1}
func badTlsSettingsNoncompliant() {
	tr := &http.Transport{
		// Noncompliant: insecure cipher with tls
		TLSClientConfig: &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
			},
			MinVersion: tls.VersionTLS13,
		},
	}
	client := &http.Client{Transport: tr}
	_, err := client.Get("https://example.com/")
	if err != nil {
		fmt.Println(err)
	}
}

// {/fact}

// {fact rule=insecure-cryptography@v1.0 defects=0}
func badTlsSettingsCompliant() {
	tr := &http.Transport{
		// Compliant: secure cipher with tls
		TLSClientConfig: &tls.Config{
			CipherSuites: []uint16{
				tls.TLS_AES_128_GCM_SHA256,
				tls.TLS_AES_256_GCM_SHA384,
			},
			MinVersion: tls.VersionTLS13,
		},
	}
	client := &http.Client{Transport: tr}
	_, err := client.Get("https://example.com/")
	if err != nil {
		fmt.Println(err)
	}
}

// {/fact}
