package ddb

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Marshal converts map or struct to dynamodb attribute value
func Marshal(iv interface{}) map[string]*dynamodb.AttributeValue {
	kind := reflect.TypeOf(iv).Kind()
	if kind == reflect.Map {
		return marshalMap(iv)
	} else if kind == reflect.Ptr {
		return marshalStruct(reflect.ValueOf(iv).Elem())
	}

	return nil
}

func marshalMap(iv interface{}) map[string]*dynamodb.AttributeValue {
	vv := iv.(map[string]interface{})
	ret := make(map[string]*dynamodb.AttributeValue)

	for key, value := range vv {
		ret[key] = marshal(value)
	}

	return ret
}

func marshalStruct(value reflect.Value) map[string]*dynamodb.AttributeValue {

	t := value.Type()

	numField := t.NumField()

	ret := make(map[string]*dynamodb.AttributeValue)
	for i := 0; i < numField; i++ {
		f := t.Field(i)

		if f.PkgPath != "" {
			continue
		}

		name, option := parseTag(f.Tag.Get("dynamodb"))
		if name == "-" {
			continue
		}
		if option == "omitifempty" {
			continue
		}
		if name == "" {
			name = f.Name
		}
		ret[name] = marshalValue(value.FieldByIndex(f.Index))
	}

	return ret
}

func parseTag(tag string) (string, string) {
	if idx := strings.Index(tag, ","); idx != -1 {
		return tag[:idx], tag[idx+1:]
	}
	return tag, ""
}

func marshal(value interface{}) *dynamodb.AttributeValue {

	if value == nil {
		return makeNullAttrValue()
	}

	switch v := value.(type) {
	case string:
		return makeStringAttrValue(v)
	case bool:
		return makeBoolAttrValue(v)
	case int:
		return makeInt64AttrValue(int64(v))
	case int64:
		return makeInt64AttrValue(v)
	case uint:
		return makeUInt64AttrValue(uint64(v))
	case uint64:
		return makeUInt64AttrValue(v)
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

	return marshalValue(reflectValue)
}

func makeNullAttrValue() *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{NULL: aws.Bool(true)}
}

func makeStringAttrValue(v string) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{S: aws.String(v)}
}

func makeBoolAttrValue(v bool) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{BOOL: aws.Bool(v)}
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

func marshalValue(value reflect.Value) *dynamodb.AttributeValue {
	fmt.Println(value)

	switch value.Type().Kind() {
	case reflect.String:
		return makeStringAttrValue(value.String())
	case reflect.Bool:
		return makeBoolAttrValue(value.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return makeInt64AttrValue(value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return makeUInt64AttrValue(value.Uint())
	case reflect.Float32, reflect.Float64:
		return makeFloat64AttrValue(value.Float())
	}

	return nil
}
