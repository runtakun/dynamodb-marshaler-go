package ddb_test

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/kr/pretty"
	. "github.com/runtakun/dynamodb-marshaler-go"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Unmarshal", func() {

	Context("input normal value", func() {

		var sut sample

		BeforeEach(func() {
			d := map[string]*dynamodb.AttributeValue{
				"str": &dynamodb.AttributeValue{S: aws.String("foo")},
			}
			Unmarshal(d, &sut)
			log.Println(pretty.Sprintf("sut1: %#v", sut))
		})

		It("should not be nil", func() {
			Expect(sut).NotTo(BeNil())
		})

		It("should be `str` index to struct", func() {
			Expect(sut.Str).To(Equal("foo"))
			Expect(sut.Str).NotTo(BeEmpty())
		})

	})

})
