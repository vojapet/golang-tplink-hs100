package hs1x0_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHs110(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hs110 Suite")
}
