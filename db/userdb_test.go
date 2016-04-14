package db_test

import (
	
	. "github.com/migdi/delphos-api/db"
	. "github.com/migdi/delphos-api/models"

	. "github.com/onsi/ginkgo"    
	. "github.com/onsi/gomega"
)

var _ = Describe("Userdb", func() {
	Context("User BD Functions", func() {
		var set Userdb
	
		It("starts off empty", func() {
			set = NewUserdb()
			Expect(set.Size()).Should(BeEquivalentTo(0))
		})
		
		It("Insert an User", func() {
			set = NewUserdb()
			
			Expect(set.Size()).Should(BeEquivalentTo(0))

			user1 := NewUser("Foo1", "bar2@bar.com", "foobar1")
			user2 := NewUser("Foo2", "bar2@bar.com", "foobar2")
			
			*user1 = set.Add(*user1)
			*user2 = set.Add(*user2)
			
			Expect(set.Size()).Should(BeEquivalentTo(2))
	
			Expect(user1.Id()).Should(BeNumerically(">", 0))
			Expect(user1.Id()).Should(BeNumerically("<", user2.Id()))

		})
		
		It("Constains  or not an user", func() {
			set = NewUserdb()
			user1 := NewUser("Foo1", "bar2@bar.com", "foobar1")
			user2 := NewUser("Foo2", "bar2@bar.com", "foobar2")
		
			*user1 =set.Add(*user1)
			
			Expect(set.Contains(*user1)).Should(BeTrue())
			Expect(set.Contains(*user2)).Should(BeFalse())
		})
		
		
		It("Searh user by email", func() {
			set = NewUserdb()
			user1 := NewUser("Foo1", "bar2@bar.com", "foobar1")
			user2 := NewUser("Foo2", "bar2@bar.com", "foobar2")
	
			*user1 = set.Add(*user1)
			*user2 = set.Add(*user2)

			userFound, err := set.RetrieveByEmail(user1.Email())
			Expect(err).Should(BeNil())
			Expect(userFound).Should(Equal(*user1))
			
			
			userFound, err = set.RetrieveByEmail("none@gmail.com")
			Expect(err).ShouldNot(BeNil())

		})
		
		
		
	})
})