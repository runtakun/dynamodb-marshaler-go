package ddb_test

import (
	"math"
	"strconv"
	"testing"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/runtakun/dynamodb-marshaler-go"
)

func TestMarshalMap(t *testing.T) {

	m := map[string]interface{}{
		"str":     "string",
		"int":     1,
		"int64":   int64(1),
		"float32": float32(math.Pi),
		"float64": math.Pi,
		"byte":    []byte{0x0, 0x1, 0x2, 0x3, 0x4},
		"null":    nil,
		"ss":      []string{"a", "b", "c", "d", "e"},
		"ns":      []int{1, 2, 3},
		"bs": [][]byte{
			[]byte{0x0, 0x1, 0x2, 0x3},
			[]byte{0x4, 0x5, 0x6, 0x7},
			[]byte{0x8, 0x9, 0xa, 0xb},
			[]byte{0xc, 0xd, 0xe, 0xf},
		},
		"map": map[string]interface{}{
			"id":    "hogehoge",
			"child": true,
		},
		"list": []interface{}{
			map[string]interface{}{
				"key":  1,
				"name": "name1",
			},
			map[string]interface{}{
				"key":  2,
				"name": "name2",
			},
		},
	}

	u := ddb.Marshal(m)

	if u == nil {
		t.Error("t is nil")
	}

	if _, ok := u["str"]; !ok {
		t.Error("not includes element")
		return
	}

	str := u["str"].S
	if *str != "string" {
		t.Error("str does not match")
		return
	}

	n32, _ := strconv.ParseInt(*u["int"].N, 10, 32)
	if n32 != 1 {
		t.Error("int does not match")
		return
	}

	n64, _ := strconv.ParseInt(*u["int64"].N, 10, 64)
	if n64 != 1 {
		t.Error("int64 does not match")
		return
	}

	if len(u["byte"].B) != 5 {
		t.Error("byte does not match")
		return
	}

	if _, ok := u["float32"]; !ok {
		t.Error("not includes element")
		return
	}

	if _, ok := u["float64"]; !ok {
		t.Error("not includes element")
		return
	}

	if !*u["null"].NULL {
		t.Error("null does not match")
		return
	}

	if len(u["ss"].SS) != 5 {
		t.Error("ss set does not match")
		return
	}

	if len(u["ns"].NS) != 3 {
		t.Error("ns does not match")
		return
	}

	if len(u["bs"].BS) != 4 {
		t.Error("bs does not match")
		return
	}

	l := u["list"].L
	if len(l) != 2 {
		t.Error("list does not match")
		return
	}

	t.Logf("return value: %s", awsutil.Prettify(u))
}
