package iogo_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestIogo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Iogo Suite")
}
