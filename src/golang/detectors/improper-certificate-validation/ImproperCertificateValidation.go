package main

import (
	"crypto/tls"
	"net/http"
)

func improperCertificateValidationNoncompliant(authReq *http.Request) *http.Response {
	// {fact rule=improper-certificate-validation@v1.0 defects=1}
	// Noncompliant : `InsecureSkipVerify` parameter set to true
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			MinVersion:         tls.VersionTLS13,
		},
	}
	// {/fact}
	client := &http.Client{Transport: tr}
	res, _ := client.Do(authReq)
	return res
}

func improperCertificateValidationCompliant(authReq *http.Request) *http.Response {
	// {fact rule=improper-certificate-validation@v1.0 defects=0}
	// Compliant : `InsecureSkipVerify` parameter set to false
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: false,
			MinVersion:         tls.VersionTLS13,
		},
	}
	// {/fact}
	client := &http.Client{Transport: tr}
	res, _ := client.Do(authReq)
	return res
}
