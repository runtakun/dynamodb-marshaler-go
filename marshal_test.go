package ddb_test

import (
	"testing"

	"github.com/runtakun/dynamodb-marshaler-go"
)

func TestMarshalMap(t *testing.T) {

	m := map[string]interface{}{
		"str":   "string",
		"int":   1,
		"float": 1.0,
		"null":  nil,
	}

	u := ddb.Marshal(m)

	if u == nil {
		t.Error("t is nil")
		return
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
}
