package iogo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/silverwyrda/iogo"
)

var _ = Describe("Router", func() {
	It("should know the result of 2+2", func() {
		router := iogo.New()
		Expect(func() {
			router.Handle("GET", "ehlo", nil)
		}).To(Panic())
	})
})
