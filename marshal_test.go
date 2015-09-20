package ddb_test

import (
	"testing"

	"github.com/runtakun/dynamodb-marshaler-go"
)

func TestMarshalMap(t *testing.T) {

	m := map[string]interface{}{
		"str": "string",
		"num": 1,
	}

	u := ddb.Marshal(m)
	if u == nil {
		t.Error("t is nil")
		return
	}
}
