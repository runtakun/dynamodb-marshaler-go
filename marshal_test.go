package ddb_test

import (
	"math"
	"strconv"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	. "github.com/runtakun/dynamodb-marshaler-go"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type sample struct {
	Str             string                 `json:"str"`
	Bool            bool                   `json:"bool"`
	Blob            []byte                 `json:"blob"`
	Int             int                    `json:"int"`
	Int8            int8                   `json:"int8"`
	Int16           int16                  `json:"int16"`
	Int32           int32                  `json:"int32"`
	Int64           int64                  `json:"int64"`
	Uint            uint                   `json:"uint"`
	Uint8           uint8                  `json:"uint8"`
	Uint16          uint16                 `json:"uint16"`
	Uint32          uint32                 `json:"uint32"`
	Uint64          uint64                 `json:"uint64"`
	Float32         float32                `json:"float32"`
	Float64         float64                `json:"float64"`
	Arr             [3]int                 `json:"arr"`
	InterfaceInt    interface{}            `json:"interface_int"`
	InterfaceString interface{}            `json:"interface_str"`
	Map             map[string]interface{} `json:"map"`
	Ptr             *string                `json:"ptr"`
	Slice           []string               `json:"slice"`
	EmptySlice      []int                  `json:"empty_slice"`
	Child           *child                 `json:"child"`
}

type child struct {
	Content string `json:"content"`
}

var _ = Describe("Marshal", func() {
	Context("input struct", func() {

		var sut map[string]*dynamodb.AttributeValue

		BeforeEach(func() {

			ptr := "ptr"

			s := &sample{
				Str:             "foo",
				Bool:            false,
				Blob:            []byte{0x0, 0x1, 0x2, 0x3, 0x4},
				Int:             1,
				Int8:            2,
				Int16:           3,
				Int32:           4,
				Int64:           5,
				Uint:            1,
				Uint8:           2,
				Uint16:          3,
				Uint32:          4,
				Uint64:          5,
				Float32:         math.E,
				Float64:         math.Pi,
				Arr:             [3]int{1, 2, 3},
				EmptySlice:      []int{},
				InterfaceInt:    12345,
				InterfaceString: "bar",
				Map: map[string]interface{}{
					"map_foo": "map_bar",
					"map_int": 54321,
				},
				Ptr:   &ptr,
				Slice: []string{"f", "o", "o"},
				Child: &child{Content: "bar_child"},
			}
			sut = Marshal(s)
			GinkgoWriter.Write([]byte(awsutil.Prettify(sut)))
		})

		It("should not be nil", func() {
			Expect(sut).NotTo(BeNil())
		})

		It("should be `str` element to dynamodb attribute value", func() {
			Expect(*sut["str"].S).To(Equal("foo"))
			Expect(sut["str"].N).To(BeNil())
		})

		It("should be `bool` element to dynamodb attribute value", func() {
			Expect(*sut["bool"].BOOL).To(BeFalse())
			Expect(sut["bool"].S).To(BeNil())
		})

		It("should be `blob` element to dynamodb attribute value", func() {
			Expect(sut["blob"].B).Should(HaveLen(5))
		})

		It("should be `int` element to dynamodb attribute value", func() {
			Expect(*sut["int"].N).To(Equal("1"))
			Expect(sut["int"].S).To(BeNil())
		})

		It("should be `int8` element to dynamodb attribute value", func() {
			Expect(*sut["int8"].N).To(Equal("2"))
			Expect(sut["int8"].S).To(BeNil())
		})

		It("should be `int16` element to dynamodb attribute value", func() {
			Expect(*sut["int16"].N).To(Equal("3"))
			Expect(sut["int16"].S).To(BeNil())
		})

		It("should be `int32` element to dynamodb attribute value", func() {
			Expect(*sut["int32"].N).To(Equal("4"))
			Expect(sut["int32"].S).To(BeNil())
		})

		It("should be `int64` element to dynamodb attribute value", func() {
			Expect(*sut["int64"].N).To(Equal("5"))
			Expect(sut["int64"].S).To(BeNil())
		})

		It("should be `uint` element to dynamodb attribute value", func() {
			Expect(*sut["uint"].N).To(Equal("1"))
			Expect(sut["uint"].S).To(BeNil())
		})

		It("should be `uint8` element to dynamodb attribute value", func() {
			Expect(*sut["uint8"].N).To(Equal("2"))
			Expect(sut["uint8"].S).To(BeNil())
		})

		It("should be `uint16` element to dynamodb attribute value", func() {
			Expect(*sut["uint16"].N).To(Equal("3"))
			Expect(sut["uint16"].S).To(BeNil())
		})

		It("should be `uint32` element to dynamodb attribute value", func() {
			Expect(*sut["uint32"].N).To(Equal("4"))
			Expect(sut["uint32"].S).To(BeNil())
		})

		It("should be `uint64` element to dynamodb attribute value", func() {
			Expect(*sut["uint64"].N).To(Equal("5"))
			Expect(sut["uint64"].S).To(BeNil())
		})

		It("should be `float32` element to dynamodb attribute value", func() {
			f32, _ := strconv.ParseFloat(*sut["float32"].N, 32)
			Expect(f32).Should(BeNumerically("~", math.E, 1e-6))
		})

		It("should be `float64` element to dynamodb attribute value", func() {
			f64, _ := strconv.ParseFloat(*sut["float64"].N, 64)
			Expect(f64).Should(BeNumerically("~", math.Pi, 1e-6))
		})

		It("should be `arr` element to dynamodb attribute value", func() {
			Expect(sut["arr"].L).Should(HaveLen(3))
			Expect(sut["arr"].S).To(BeNil())
		})

		It("should be `empty_slice` element to dynamodb attribute value", func() {
			Expect(sut["empty_slice"].L).Should(HaveLen(0))
			Expect(sut["empty_slice"].N).To(BeNil())
		})

		It("should be interface type to dynamodb attribute value", func() {
			Expect(*sut["interface_int"].N).To(Equal("12345"))
			Expect(*sut["interface_str"].S).To(Equal("bar"))
		})

		It("should be `map` type to dynamodb attribute value", func() {
			m := sut["map"].M
			Expect(*m["map_foo"].S).To(Equal("map_bar"))
			Expect(*m["map_int"].N).To(Equal("54321"))
		})

		It("should be `ptr` type to dynamodb attribute value", func() {
			Expect(*sut["ptr"].S).To(Equal("ptr"))
		})

		It("should be `slice` type to dynamodb attribute value", func() {
			Expect(sut["slice"].L).Should(HaveLen(3))
		})

		It("should be `child` type to dynamodb attribute value", func() {
			child := sut["child"].M
			Expect(*child["content"].S).To(Equal("bar_child"))
		})
	})

	Context("empty value", func() {

		var sut map[string]*dynamodb.AttributeValue

		BeforeEach(func() {
			s := &sample{
				Str:   "",
				Map:   nil,
				Child: nil,
			}
			sut = Marshal(s)
			GinkgoWriter.Write([]byte(awsutil.Prettify(sut)))
		})

		It("should be `str` type to null value", func() {
			Expect(*sut["str"].NULL).To(BeTrue())
			Expect(sut["str"].S).To(BeNil())
		})

		It("should be `map` type to null value", func() {
			Expect(*sut["map"].NULL).To(BeTrue())
			Expect(sut["map"].M).To(BeNil())
		})

		It("should be `child` type to null value", func() {
			Expect(*sut["child"].NULL).To(BeTrue())
			Expect(sut["child"].M).To(BeNil())
		})

	})

})
