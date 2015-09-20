package ddb

import (
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func Marshal(iv interface{}) map[string]*dynamodb.AttributeValue {

	vv := iv.(map[string]interface{})
	ret := make(map[string]*dynamodb.AttributeValue)

	for key, value := range vv {

		if value == nil {
			ret[key] = &dynamodb.AttributeValue{NULL: aws.Bool(true)}
			continue
		}

		var attrValue *dynamodb.AttributeValue

		switch v := value.(type) {
		case string:
			attrValue = &dynamodb.AttributeValue{S: aws.String(v)}
		case int:
			attrValue = makeInt64AttrValue(int64(v))
		case int64:
			attrValue = makeInt64AttrValue(v)
		case uint:
			attrValue = makeUInt64AttrValue(uint64(v))
		case uint64:
			attrValue = makeUInt64AttrValue(v)
		case bool:
			attrValue = &dynamodb.AttributeValue{BOOL: aws.Bool(v)}
		case float32:
			attrValue = makeFloat64AttrValue(float64(v))
		case float64:
			attrValue = makeFloat64AttrValue(float64(v))
		case []string:
			attrValue = makeStringSliceAttrValue(v)
		}

		reflectValue := reflect.ValueOf(value)
		t := reflectValue.Type()
		if t.Kind() == reflect.Map {
			attrValue = makeMapAttrValue(value)
		}

		ret[key] = attrValue
	}

	return ret
}

func makeInt64AttrValue(v int64) *dynamodb.AttributeValue {
	return makeNumberAttrValue(fmt.Sprintf("%d", v))
}

func makeUInt64AttrValue(v uint64) *dynamodb.AttributeValue {
	return makeNumberAttrValue(fmt.Sprintf("%d", v))
}

func makeFloat64AttrValue(v float64) *dynamodb.AttributeValue {
	return makeNumberAttrValue(fmt.Sprintf("%f", v))
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

func makeMapAttrValue(m interface{}) *dynamodb.AttributeValue {
	return &dynamodb.AttributeValue{M: Marshal(m)}
}
