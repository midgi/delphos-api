package models_test

import (
	"github.com/migdi/delphos-api/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	Context("when creating users", func() {
		It("must encrypt the password", func() {
			user := models.NewUser("Foo", "bar@bar.com", "foobar123")
			Expect(user.Password()).NotTo(Equal("foobar123"))
		})
	})
})
