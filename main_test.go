package gocva_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestValidator(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "goCVA")
}

var _ = Describe("goCVA", func() {
	classTests()
	cvaTests()
})
