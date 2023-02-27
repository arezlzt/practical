package practical

import (
	"reflect"
)

// StructToMap
// struct to map
func StructToMap[T any](obj T) map[string]interface{} {
	if reflect.TypeOf(reflect.ValueOf(obj).Interface()).Kind() != reflect.Struct {
		panic("param not is struct")
	}
	oValue := reflect.ValueOf(obj)
	tValue := reflect.TypeOf(obj)
	result := make(map[string]interface{})
	for i := 0; i < oValue.NumField(); i++ {
		if reflect.TypeOf(oValue.Field(i).Interface()).Kind() == reflect.Struct {
			return structToMap(oValue.Field(i).Interface(), result)
		} else {
			result[CamelToCase(tValue.Field(i).Name)] = oValue.Field(i).Interface()
		}

	}
	return result
}
func structToMap[T any](obj T, result map[string]interface{}) map[string]interface{} {
	oValue := reflect.ValueOf(obj)
	tValue := reflect.TypeOf(obj)
	for i := 0; i < oValue.NumField(); i++ {
		if reflect.TypeOf(oValue.Field(i).Interface()).Kind() == reflect.Struct {
			return structToMap(oValue.Field(i).Interface(), result)
		} else {
			result[CamelToCase(tValue.Field(i).Name)] = oValue.Field(i).Interface()
		}
	}
	return result
}

func StructColumn[T comparable, E comparable](list []T, fieldName string) []E {
	arr := make([]E, 0)
	if len(list) == 0 {
		return arr
	}
	if reflect.TypeOf(list[0]).Kind() != reflect.Struct {
		panic("T not is struct")
	}
	fieldName = CaseToCamel(fieldName)
	for _, value := range list {
		oValue := reflect.ValueOf(value)
		tValue := reflect.TypeOf(value)
		for i := 0; i < tValue.NumField(); i++ {
			if tValue.Field(i).Name == fieldName {
				arr = append(arr, oValue.Field(i).Interface().(E))
			}
		}
	}
	return arr
}
