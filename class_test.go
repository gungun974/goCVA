package gocva_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/gungun974/gocva"
)

func classTests() {
	Describe("Class", func() {
		It("should render empty class", func() {
			// arrange
			cb := gocva.ClassBuilder{}

			// act
			result := cb.String()

			// assert
			Expect(result).To(Equal(""))
		})

		It("should render one class", func() {
			// arrange
			cb := gocva.ClassBuilder{}

			// act
			cb.Add("first")

			result := cb.String()

			// assert
			Expect(result).To(Equal("first"))
		})

		It("should render multiple classes in right order", func() {
			// arrange
			cb := gocva.ClassBuilder{}

			// act
			cb.Add("first")
			cb.Add("second")
			cb.Add("third")

			result := cb.String()

			// assert
			Expect(result).To(Equal("first second third"))
		})
	})
}
