package ddb_test

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	. "github.com/runtakun/dynamodb-marshaler-go"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Unmarshal", func() {

	Context("input value", func() {

		var sut sample

		BeforeEach(func() {
			d := map[string]*dynamodb.AttributeValue{
				"str":    &dynamodb.AttributeValue{S: aws.String("foo")},
				"bool":   &dynamodb.AttributeValue{BOOL: aws.Bool(true)},
				"int":    &dynamodb.AttributeValue{N: aws.String(fmt.Sprintf("%d", -42775))},
				"int8":   &dynamodb.AttributeValue{N: aws.String(fmt.Sprintf("%d", -100))},
				"int16":  &dynamodb.AttributeValue{N: aws.String(fmt.Sprintf("%d", -12345))},
				"int32":  &dynamodb.AttributeValue{N: aws.String(fmt.Sprintf("%d", -1073741014))},
				"int64":  &dynamodb.AttributeValue{N: aws.String(fmt.Sprintf("%d", -1112222111222111))},
				"uint":   &dynamodb.AttributeValue{N: aws.String(fmt.Sprintf("%d", 42775))},
				"uint8":  &dynamodb.AttributeValue{N: aws.String(fmt.Sprintf("%d", 111))},
				"uint16": &dynamodb.AttributeValue{N: aws.String(fmt.Sprintf("%d", 12345))},
				"uint32": &dynamodb.AttributeValue{N: aws.String(fmt.Sprintf("%d", 1073741014))},
				"uint64": &dynamodb.AttributeValue{N: aws.String(fmt.Sprintf("%d", 1112222111222111))},
			}
			Unmarshal(d, &sut)
		})

		It("should not be nil", func() {
			Expect(sut).NotTo(BeNil())
		})

		It("should be struct which has `Str` column", func() {
			Expect(sut.Str).To(Equal("foo"))
			Expect(sut.Str).NotTo(BeEmpty())
		})

		It("should be struct which has `Bool` column", func() {
			Expect(sut.Bool).To(BeTrue())
			Expect(sut.Bool).NotTo(BeFalse())
		})

		It("should be struct which has `Int` column", func() {
			Expect(sut.Int).To(Equal(-42775))
		})

		It("should be struct which has `Int8` column", func() {
			Expect(sut.Int8).To(Equal(int8(-100)))
		})

		It("should be struct which has `Int16` column", func() {
			Expect(sut.Int16).To(Equal(int16(-12345)))
		})

		It("should be struct which has `Int32` column", func() {
			Expect(sut.Int32).To(Equal(int32(-1073741014)))
		})

		It("should be struct which has `Int64` column", func() {
			Expect(sut.Int64).To(Equal(int64(-1112222111222111)))
		})

		It("should be struct which has `Uint` column", func() {
			Expect(sut.Uint).To(Equal(uint(42775)))
		})

		It("should be struct which has `Uint8` column", func() {
			Expect(sut.Uint8).To(Equal(uint8(111)))
		})

		It("should be struct which has `Uint16` column", func() {
			Expect(sut.Uint16).To(Equal(uint16(12345)))
		})

		It("should be struct which has `Uint32` column", func() {
			Expect(sut.Uint32).To(Equal(uint32(1073741014)))
		})

		It("should be struct which has `Uint64` column", func() {
			Expect(sut.Uint64).To(Equal(uint64(1112222111222111)))
		})

	})

})
