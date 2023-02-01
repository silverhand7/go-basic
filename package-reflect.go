package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"` // required and max are the tag
}

// validation is required
// make validation using reflection
func isRequired(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}

	return true
}

func main() {
	sample := Sample{"Refo"}
	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampleType, sampleType.NumField(), sampleType.Field(0))

	fmt.Println(sampleType.Field(0).Tag.Get("required"), sampleType.Field(0).Tag.Get("max")) // get value on tag

	fmt.Println("validation", isRequired(sample))

	sample.Name = ""
	fmt.Println("Validation", isRequired(sample))

}
