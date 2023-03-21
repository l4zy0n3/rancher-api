package rancher_api_test

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type RequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var _ = Describe("LoginApi", func() {
	It("should login to rancher", func() {
		// Create a client with insecure verification disabled
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}

		// Create a request
		req, err := http.NewRequest("POST", os.Getenv("API_BASE_URL")+"/v3-public/localProviders/local?action=login", bytes.NewBuffer([]byte(`{"username":"`+os.Getenv("RANCHER_USERNAME")+`","password":"`+os.Getenv("RANCHER_PASSWORD")+`"}`)))
		req.Header.Add("Content-Type", "application/json")
		Expect(err).To(BeNil())

		// Send the request and check the response
		resp, err := client.Do(req)
		Expect(err).To(BeNil())
		Expect(resp.StatusCode).To(Equal(201))
	})
})
