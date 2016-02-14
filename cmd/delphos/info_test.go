package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Info", func() {
	Describe("GET /info", func() {
		It("returns information about the delphos API", func() {
			By("generate request")
			req, err := http.NewRequest("GET", fmt.Sprintf("http://127.0.0.1:%d/info", delphosPort), nil)
			Expect(err).NotTo(HaveOccurred())

			By("send request to server")
			resp, err := httpclient.Do(req)
			Expect(err).NotTo(HaveOccurred())

			By("reading the body")
			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())

			By("checking contents")
			Expect(body).To(MatchJSON(`{"version": "0.0.1"}`))
		})
	})
})
