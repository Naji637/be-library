package service

import "reflect"

func CloneStruct[T any](fromStruct any) T {
	toStruct := new(T)
	toVal := reflect.ValueOf(toStruct).Elem()
	fromVal := reflect.ValueOf(fromStruct)

	for i := 0; i < fromVal.NumField(); i++ {
		fromField := fromVal.Type().Field(i)
		toField := toVal.FieldByName(fromField.Name)

		if toField.IsValid() && toField.CanSet() && fromField.Type == toField.Type() {
			toField.Set(fromVal.Field(i))
		}
	}
	return *toStruct
}