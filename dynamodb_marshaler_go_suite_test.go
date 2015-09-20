package ddb_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDynamodbMarshalerGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DynamodbMarshalerGo Suite")
}
