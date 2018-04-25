package oasis_dex_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOasisDex(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OasisDex Suite")
}
