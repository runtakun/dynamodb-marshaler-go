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
				"blob":    &dynamodb.AttributeValue{B: []byte{0x1, 0x2, 0x3}},
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
				"map": &dynamodb.AttributeValue{M: map[string]*dynamodb.AttributeValue{
					"map_foo":   &dynamodb.AttributeValue{S: aws.String("map_bar")},
					"map_int":   &dynamodb.AttributeValue{N: aws.String("54321")},
					"map_long":  &dynamodb.AttributeValue{N: aws.String("1223362036844775800")},
					"map_float": &dynamodb.AttributeValue{N: aws.String("3.14")},
					"map_ss":    &dynamodb.AttributeValue{SS: []*string{aws.String("b"), aws.String("a"), aws.String("r")}},
					"map_ns":    &dynamodb.AttributeValue{NS: []*string{aws.String("1"), aws.String("2"), aws.String("3")}},
					"map_bs":    &dynamodb.AttributeValue{BS: [][]byte{[]byte{0x1, 0x2, 0x3}, []byte{0x3, 0x2, 0x1}, []byte{0x3, 0x3, 0x3}}},
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

		It("should be struct which has `Blob` column", func() {
			Expect(sut.Blob).Should(HaveLen(3))
			Expect(sut.Blob[0]).To(Equal(byte(1)))
			Expect(sut.Blob[1]).To(Equal(byte(2)))
			Expect(sut.Blob[2]).To(Equal(byte(3)))
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

		It("should be map which has `map_foo` column", func() {
			Expect(sut.Map).ShouldNot(BeNil())

			v, ok := sut.Map["map_foo"]
			Expect(ok).Should(BeTrue())
			Expect(v).To(Equal("map_bar"))
		})
		It("should be map which has `map_int` column", func() {
			Expect(sut.Map).ShouldNot(BeNil())

			v, ok := sut.Map["map_int"]
			Expect(ok).Should(BeTrue())
			Expect(v).To(Equal(54321))
		})

		It("should be map which has `map_long` column", func() {
			Expect(sut.Map).ShouldNot(BeNil())

			v, ok := sut.Map["map_long"]
			Expect(ok).Should(BeTrue())
			Expect(v).To(Equal(1223362036844775800))
		})

		It("should be map which has `map_float` column", func() {
			Expect(sut.Map).ShouldNot(BeNil())

			v, ok := sut.Map["map_float"]
			Expect(ok).Should(BeTrue())
			Expect(v).To(Equal(3.14))
		})

		It("should be map which has `map_ss` column", func() {
			Expect(sut.Map).ShouldNot(BeNil())

			v, ok := sut.Map["map_ss"]
			Expect(ok).Should(BeTrue())

			vv := v.([]string)

			Expect(vv).Should(HaveLen(3))
			Expect(vv[0]).To(Equal("b"))
			Expect(vv[1]).To(Equal("a"))
			Expect(vv[2]).To(Equal("r"))
		})

		It("should be map which has `map_ns` column", func() {
			Expect(sut.Map).ShouldNot(BeNil())

			v, ok := sut.Map["map_ns"]
			Expect(ok).Should(BeTrue())

			vv := v.([]interface{})

			Expect(vv).Should(HaveLen(3))
			Expect(vv[0]).To(Equal(1))
			Expect(vv[1]).To(Equal(2))
			Expect(vv[2]).To(Equal(3))
		})

		It("should be map which has `map_bs` column", func() {
			Expect(sut.Map).ShouldNot(BeNil())

			v, ok := sut.Map["map_bs"]
			Expect(ok).Should(BeTrue())

			vv := v.([][]byte)

			Expect(vv).Should(HaveLen(3))
			Expect(vv[0]).To(Equal([]byte{0x1, 0x2, 0x3}))
			Expect(vv[1]).To(Equal([]byte{0x3, 0x2, 0x1}))
			Expect(vv[2]).To(Equal([]byte{0x3, 0x3, 0x3}))
		})
	})
})
