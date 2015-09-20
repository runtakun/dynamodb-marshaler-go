package ddb

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Marshal converts map or struct to dynamodb attribute value
func Marshal(iv interface{}) map[string]*dynamodb.AttributeValue {

	vv := iv.(map[string]interface{})
	ret := make(map[string]*dynamodb.AttributeValue)

	for key, value := range vv {
		ret[key] = marshal(value)
	}

	return ret
}

func marshal(value interface{}) *dynamodb.AttributeValue {

	if value == nil {
		return &dynamodb.AttributeValue{NULL: aws.Bool(true)}
	}

	switch v := value.(type) {
	case string:
		return &dynamodb.AttributeValue{S: aws.String(v)}
	case int:
		return makeInt64AttrValue(int64(v))
	case int64:
		return makeInt64AttrValue(v)
	case uint:
		return makeUInt64AttrValue(uint64(v))
	case uint64:
		return makeUInt64AttrValue(v)
	case bool:
		return &dynamodb.AttributeValue{BOOL: aws.Bool(v)}
	case float32:
		return makeFloat64AttrValue(float64(v))
	case float64:
		return makeFloat64AttrValue(float64(v))
	case []string:
		return makeStringSliceAttrValue(v)
	case []int:
		return makeIntSliceAttrValue(v)
	case []int64:
		return makeInt64SliceAttrValue(v)
	case []byte:
		return &dynamodb.AttributeValue{B: v}
	case [][]byte:
		return &dynamodb.AttributeValue{BS: v}
	}

	reflectValue := reflect.ValueOf(value)
	t := reflectValue.Type()
	switch t.Kind() {
	case reflect.Map:
		return makeMapAttrValue(value)
	case reflect.Slice:
		return makeListAttrValue(value)
	}

	return nil
}

func makeInt64AttrValue(v int64) *dynamodb.AttributeValue {
	return makeNumberAttrValue(fmt.Sprintf("%d", v))
}

func makeUInt64AttrValue(v uint64) *dynamodb.AttributeValue {
	return makeNumberAttrValue(fmt.Sprintf("%d", v))
}

func makeFloat32AttrValue(v float32) *dynamodb.AttributeValue {
	return makeNumberAttrValue(strconv.FormatFloat(float64(v), 'f', -1, 32))
}

func makeFloat64AttrValue(v float64) *dynamodb.AttributeValue {
	return makeNumberAttrValue(strconv.FormatFloat(v, 'f', -1, 64))
}

func makeNumberAttrValue(str string) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{N: aws.String(str)}
}

func makeStringSliceAttrValue(strs []string) *dynamodb.AttributeValue {
	slices := make([]*string, len(strs))

	for i, v := range strs {
		slices[i] = aws.String(v)
	}

	return &dynamodb.AttributeValue{SS: slices}
}

func makeIntSliceAttrValue(ints []int) *dynamodb.AttributeValue {
	slices := make([]*string, len(ints))

	for i, v := range ints {
		slices[i] = aws.String(fmt.Sprintf("%d", v))
	}

	return &dynamodb.AttributeValue{NS: slices}
}

func makeInt64SliceAttrValue(ints []int64) *dynamodb.AttributeValue {
	slices := make([]*string, len(ints))

	for i, v := range ints {
		slices[i] = aws.String(fmt.Sprintf("%d", v))
	}

	return &dynamodb.AttributeValue{NS: slices}
}

func makeMapAttrValue(m interface{}) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{M: Marshal(m)}
}

func makeListAttrValue(l interface{}) *dynamodb.AttributeValue {

	values := l.([]interface{})

	list := make([]*dynamodb.AttributeValue, len(values))
	for i, value := range values {
		list[i] = marshal(value)
	}

	return &dynamodb.AttributeValue{L: list}
}
