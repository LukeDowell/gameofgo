package game_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGameofgo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gameofgo Suite")
}
