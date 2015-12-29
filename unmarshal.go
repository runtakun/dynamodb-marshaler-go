package ddb

import (
	"errors"
	"log"
	"reflect"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/kr/pretty"
)

// Unmarshal converts dynamodb attribute value map to map or struct
func Unmarshal(item map[string]*dynamodb.AttributeValue, v interface{}) error {
	t := reflect.TypeOf(v)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	} else {
		return errors.New("value must be a pointer")
	}

	isptr := false
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		isptr = true
	}

	if t.Kind() == reflect.Struct {
		dest := reflect.ValueOf(v)
		if isptr {
			dest = dest.Elem()
		}

		if dest.IsNil() {
			dest.Set(reflect.New(t))
		}

		log.Println(pretty.Sprintf("dest: %#v", dest))

		destType := dest.Elem().Type()
		numOfField := destType.NumField()
		for i := 0; i < numOfField; i++ {
			f := destType.Field(i)

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

			if value, ok := item[name]; ok {
				switch f.Type.Kind() {
				case reflect.String:
					if value.S != nil {
						dest.Elem().FieldByIndex(f.Index).SetString(*value.S)
					}
				}
			}
		}

	}

	return nil
}
