package models_test

import (
	"github.com/migdi/delphos-api/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Advice", func() {
	Context("when creating advices", func() {
		It("must return adivce model", func() {
			user := models.NewUser("Foo", "bar@bar.com", "foobar123")
			advice := models.NewAdvice("id", "Content of Advice", user)
			Expect(advice.GetId()).To(Equal("id"))
			Expect(advice.GetId()).To(Equal("Content"))
		})
	})
})
