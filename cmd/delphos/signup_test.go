package main_test

import (
	"fmt"
	"io/ioutil"
	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Signup", func() {
	Describe("GET /signup/facebook", func() {
		It("Register the user using facebook", func() {
			By("generate request")
			req, err := http.NewRequest("GET", fmt.Sprintf("http://127.0.0.1:%d/signup/facebook", delphosPort), nil)
			Expect(err).NotTo(HaveOccurred())

			By("send request to server")
			resp, err := httpclient.Do(req)
			Expect(err).NotTo(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(http.StatusOK))

			By("reading the body")
			body, err := ioutil.ReadAll(resp.Body)
			Expect(err).NotTo(HaveOccurred())

			Expect(body).To(MatchJSON(`{"id": "0.0.1"}`))

		})
	})
})
