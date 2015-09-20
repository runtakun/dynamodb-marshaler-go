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
		"null":    nil,
		"ss":      []string{"a", "b", "c", "d", "e"},
		"ns":      []int{1, 2, 3},
		"map": map[string]interface{}{
			"id":    "hogehoge",
			"child": true,
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

	if *u["float32"].N != "3.141593" {
		t.Error("float32 does not match")
		return
	}

	if *u["float64"].N != "3.141593" {
		t.Error("float64 does not match")
		return
	}

	if !*u["null"].NULL {
		t.Error("null does not match")
		return
	}

	ss := u["ss"].SS
	if len(ss) != 5 {
		t.Error("ss set does not match")
		return
	}

	ns := u["ns"].NS
	if len(ns) != 3 {
		t.Error("ns does not match")
		return
	}

	t.Logf("return value: %s", awsutil.Prettify(u))
}
