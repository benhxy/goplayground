package books_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLearnGinkgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LearnGinkgo Suite")
}
