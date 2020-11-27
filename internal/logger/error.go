package logger

import (
	"errors"
	"fmt"
	"reflect"
)

func Error(err interface{}) error {
	t := reflect.ValueOf(err).Type().String()
	if Level == 2 || Level == 3 {
		if t == "string" {
			fmt.Println(err.(string))
		} else {
			fmt.Println(err.(error).Error())
		}
	}
	if t == "string" {
		return errors.New(err.(string))
	} else {
		return err.(error)
	}
}
