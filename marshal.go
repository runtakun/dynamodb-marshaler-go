package ddb

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var typeOfBytes = reflect.TypeOf([]byte(nil))

// Marshal converts map or struct to dynamodb attribute value
func Marshal(iv interface{}) map[string]*dynamodb.AttributeValue {
	kind := reflect.TypeOf(iv).Kind()
	if kind == reflect.Map {
		return marshalMap(reflect.ValueOf(iv))
	} else if kind == reflect.Ptr {
		return marshalStruct(reflect.ValueOf(iv).Elem())
	}

	return nil
}

func marshalMap(value reflect.Value) map[string]*dynamodb.AttributeValue {
	ret := make(map[string]*dynamodb.AttributeValue)

	for _, keyValue := range value.MapKeys() {
		if keyValue.Type().Kind() == reflect.String {
			ret[keyValue.String()] = marshalValue(value.MapIndex(keyValue))
		} else {
			panic("map key must be string")
		}
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

		name, option := parseTag(f.Tag.Get("json"))
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

func marshalValue(value reflect.Value) *dynamodb.AttributeValue {
	switch value.Type().Kind() {
	case reflect.String:
		return marshalStringValue(value)
	case reflect.Bool:
		return marshalBoolValue(value)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return marshalInt64Value(value)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return marshalUint64Value(value)
	case reflect.Float32, reflect.Float64:
		return marshalFloat64Value(value)
	case reflect.Array:
		return marshalArrayValue(value)
	case reflect.Interface:
		return marshalInterfaceValue(value)
	case reflect.Map:
		return marshalMapValue(value)
	case reflect.Ptr:
		return marshalPtrValue(value)
	case reflect.Slice:
		if value.Type() == typeOfBytes {
			return marshalBytesValue(value)
		}
		return marshalSliceValue(value)
	case reflect.Struct:
		return marshalStructValue(value)
	case reflect.UnsafePointer:
		return nil
	}

	return nil
}

func marshalStringValue(value reflect.Value) *dynamodb.AttributeValue {
	str := value.String()
	if str == "" {
		return makeNullAttrValue()
	}

	return makeStringAttrValue(str)
}

func makeNullAttrValue() *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{NULL: aws.Bool(true)}
}

func makeStringAttrValue(str string) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{S: aws.String(str)}
}

func marshalBoolValue(value reflect.Value) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{BOOL: aws.Bool(value.Bool())}
}

func marshalInt64Value(value reflect.Value) *dynamodb.AttributeValue {
	return makeNumberAttrValue(fmt.Sprintf("%d", value.Int()))
}

func marshalUint64Value(value reflect.Value) *dynamodb.AttributeValue {
	return makeNumberAttrValue(fmt.Sprintf("%d", value.Uint()))
}

func marshalFloat64Value(value reflect.Value) *dynamodb.AttributeValue {
	return makeNumberAttrValue(strconv.FormatFloat(value.Float(), 'f', -1, 64))
}

func makeNumberAttrValue(str string) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{N: aws.String(str)}
}

func marshalArrayValue(value reflect.Value) *dynamodb.AttributeValue {
	length := value.Len()

	list := make([]*dynamodb.AttributeValue, length)
	for i := 0; i < length; i++ {
		list[i] = marshalValue(value.Index(i))
	}

	return &dynamodb.AttributeValue{L: list}
}

func marshalMapValue(value reflect.Value) *dynamodb.AttributeValue {
	if value.IsNil() {
		return makeNullAttrValue()
	}

	return &dynamodb.AttributeValue{M: marshalMap(value)}
}

func marshalInterfaceValue(value reflect.Value) *dynamodb.AttributeValue {
	if value.IsNil() {
		return makeNullAttrValue()
	}

	return marshalValue(value.Elem())
}

func marshalPtrValue(value reflect.Value) *dynamodb.AttributeValue {
	if value.IsNil() {
		return makeNullAttrValue()
	}

	return marshalValue(value.Elem())
}

func marshalBytesValue(value reflect.Value) *dynamodb.AttributeValue {
	if value.IsNil() {
		return makeNullAttrValue()
	}

	return &dynamodb.AttributeValue{B: value.Bytes()}
}

func marshalSliceValue(value reflect.Value) *dynamodb.AttributeValue {
	if value.IsNil() {
		return makeNullAttrValue()
	}

	return marshalArrayValue(value)
}

func marshalStructValue(value reflect.Value) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{M: marshalStruct(value)}
}
