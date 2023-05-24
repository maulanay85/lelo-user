package common

import (
	"errors"
	"reflect"
)

func ConvertModel(source any, destination any) error {
	sourceType := reflect.TypeOf(source)
	destinationType := reflect.TypeOf(destination)
	if sourceType == nil {
		return errors.New("source is nil")
	}

	if destinationType == nil {
		return errors.New("destination is nil")
	}

	destValue := reflect.ValueOf(destination)
	sourceValue := reflect.ValueOf(source).Elem()
	dest := destValue.Elem()

	for i := 0; i < sourceValue.NumField(); i++ {
		name := sourceValue.Type().Field(i).Name
		value := sourceValue.Field(i).Interface()
		// t := sourceValue.Type().Field(i).Type
		// tt := fmt.Sprintf("%v", t)

		dest.FieldByName(name).Set(reflect.ValueOf(value))

		// switch tt {
		// case "string":
		// 	dest.FieldByName(name).SetString(fmt.Sprintf("%s", value))
		// case "int8":
		// 	dest.FieldByName(name).Set(value)
		// case "int16":
		// 	dest.FieldByName(name).SetInt(value.(int64))
		// case "int32":
		// 	dest.FieldByName(name).SetInt(value.(int64))
		// case "int64":
		// 	dest.FieldByName(name).SetInt(value.(int64))
		// case "bool":
		// 	dest.FieldByName(name).SetBool(value.(bool))
		// case "float8":
		// 	dest.FieldByName(name).SetFloat(value.(float64))
		// case "float16":
		// 	dest.FieldByName(name).SetFloat(value.(float64))
		// case "float32":
		// 	dest.FieldByName(name).SetFloat(value.(float64))
		// case "float64":
		// 	dest.FieldByName(name).SetFloat(value.(float64))
		// }
	}
	return nil
}
