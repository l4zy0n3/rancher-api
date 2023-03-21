package rancher_api_test

import (
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRancherApi(t *testing.T) {
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		t.Error("Error loading .env file")
	}

	RegisterFailHandler(Fail)
	RunSpecs(t, "RancherApi Suite")
}
