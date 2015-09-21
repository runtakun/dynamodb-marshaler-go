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
	Str             string                 `dynamodb:"str"`
	Bool            bool                   `dynamodb:"bool"`
	Int             int                    `dynamodb:"int"`
	Int8            int8                   `dynamodb:"int8"`
	Int16           int16                  `dynamodb:"int16"`
	Int32           int32                  `dynamodb:"int32"`
	Int64           int64                  `dynamodb:"int64"`
	Uint            uint                   `dynamodb:"uint"`
	Uint8           uint8                  `dynamodb:"uint8"`
	Uint16          uint16                 `dynamodb:"uint16"`
	Uint32          uint32                 `dynamodb:"uint32"`
	Uint64          uint64                 `dynamodb:"uint64"`
	Float32         float32                `dynamodb:"float32"`
	Float64         float64                `dynamodb:"float64"`
	Arr             [3]int                 `dynamodb:"arr"`
	InterfaceInt    interface{}            `dynamodb:"interface_int"`
	InterfaceString interface{}            `dynamodb:"interface_str"`
	Map             map[string]interface{} `dynamodb:"map"`
	Ptr             *string                `dynamodb:"ptr"`
	Slice           []string               `dynamodb:"slice"`
	Child           child                  `dynamodb:"child"`
}

type child struct {
	Content string `dynamodb:"content"`
}

var _ = Describe("Marshal", func() {
	Context("input struct", func() {

		var sut map[string]*dynamodb.AttributeValue

		BeforeEach(func() {

			ptr := "ptr"

			s := &sample{
				Str:             "foo",
				Bool:            false,
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
				InterfaceInt:    12345,
				InterfaceString: "bar",
				Map: map[string]interface{}{
					"map_foo": "map_foo",
					"map_int": 54321,
				},
				Ptr:   &ptr,
				Slice: []string{"f", "o", "o"},
				Child: child{Content: "bar_child"},
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

		It("should be `bool` element to dynamodb attribute value", func() {
			Expect(*sut["bool"].BOOL).To(BeFalse())
			Expect(sut["bool"].S).To(BeNil())
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

		It("should be interface type to dynamodb attribute value", func() {
			Expect(*sut["interface_int"].N).To(Equal("12345"))
			Expect(*sut["interface_str"].S).To(Equal("bar"))
		})

		It("should be `map` type to dynamodb attribute value", func() {
			// Expect(sut["map"].M).Should(HaveKeyWithValue("map_foo", "map_bar"))
			// Expect(sut["map"].M).Should(HaveKeyWithValue("map_int", 54321))
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
})
