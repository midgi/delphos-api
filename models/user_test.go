package models_test

import (
	. "github.com/kirederik/delphos/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	Context("when creating users", func() {

		It("must encrypt the password", func() {
			user := NewUser("Foo", "bar@bar.com", "foobar123")
			Expect(user.GetPassword()).NotTo(Equal("foobar123"))
		})

	})
})