package ddb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func Marshal(iv interface{}) map[string]*dynamodb.AttributeValue {

	vv := iv.(map[string]interface{})
	ret := make(map[string]*dynamodb.AttributeValue)

	for key, value := range vv {

		var attrValue *dynamodb.AttributeValue

		switch v := value.(type) {
		case string:
			attrValue = &dynamodb.AttributeValue{S: aws.String(v)}
		}

		ret[key] = attrValue
	}

	return ret
}
