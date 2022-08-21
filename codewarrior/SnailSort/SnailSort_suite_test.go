package SnailSort_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestSnailSort(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SnailSort Suite")
}
