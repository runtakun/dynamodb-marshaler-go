package ddb_test

import (
	"fmt"
	"math"
	"strconv"

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
				"str":     &dynamodb.AttributeValue{S: aws.String("foo")},
				"bool":    &dynamodb.AttributeValue{BOOL: aws.Bool(true)},
				"int":     &dynamodb.AttributeValue{N: aws.String(fmt.Sprintf("%d", -42775))},
				"int8":    &dynamodb.AttributeValue{N: aws.String(fmt.Sprintf("%d", -100))},
				"int16":   &dynamodb.AttributeValue{N: aws.String(fmt.Sprintf("%d", -12345))},
				"int32":   &dynamodb.AttributeValue{N: aws.String(fmt.Sprintf("%d", -1073741014))},
				"int64":   &dynamodb.AttributeValue{N: aws.String(fmt.Sprintf("%d", -1112222111222111))},
				"uint":    &dynamodb.AttributeValue{N: aws.String(fmt.Sprintf("%d", 42775))},
				"uint8":   &dynamodb.AttributeValue{N: aws.String(fmt.Sprintf("%d", 111))},
				"uint16":  &dynamodb.AttributeValue{N: aws.String(fmt.Sprintf("%d", 12345))},
				"uint32":  &dynamodb.AttributeValue{N: aws.String(fmt.Sprintf("%d", 1073741014))},
				"uint64":  &dynamodb.AttributeValue{N: aws.String(fmt.Sprintf("%d", 1112222111222111))},
				"float32": &dynamodb.AttributeValue{N: aws.String(strconv.FormatFloat(math.Pi, 'f', -1, 32))},
				"float64": &dynamodb.AttributeValue{N: aws.String(strconv.FormatFloat(math.E, 'f', -1, 64))},
				"slice": &dynamodb.AttributeValue{SS: []*string{
					aws.String("a"),
					aws.String("b"),
					aws.String("c"),
				}},
				"empty_slice": &dynamodb.AttributeValue{NS: []*string{
					aws.String("1"),
					aws.String("2"),
					aws.String("3"),
				}},
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

		It("should be struct which has `float32` column", func() {
			Expect(sut.Float32).To(Equal(float32(math.Pi)))
		})

		It("should be struct which has `float64` column", func() {
			Expect(sut.Float64).To(Equal(float64(math.E)))
		})

		It("should be struct which has `slice` column", func() {
			Expect(sut.Slice).Should(HaveLen(3))
			Expect(sut.Slice[0]).To(Equal("a"))
			Expect(sut.Slice[1]).To(Equal("b"))
			Expect(sut.Slice[2]).To(Equal("c"))
		})

		It("should be struct which has `empty_slice` column", func() {
			Expect(sut.EmptySlice).Should(HaveLen(3))
			Expect(sut.EmptySlice[0]).To(Equal(1))
			Expect(sut.EmptySlice[1]).To(Equal(2))
			Expect(sut.EmptySlice[2]).To(Equal(3))
		})

	})

})
