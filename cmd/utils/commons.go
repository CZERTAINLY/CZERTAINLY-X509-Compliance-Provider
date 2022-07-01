package utils

import (
	"log"
	"reflect"
	"strconv"
)

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// InterfaceAsStringArray converts an interface to a string array
func InterfaceAsStringArray(value interface{}) []string {
	var resultingInfo []string
	if reflect.TypeOf(value).String() == "string" {
		resultingInfo = []string{value.(string)}
	} else {
		aInterface := value.([]interface{})
		resultingInfo = make([]string, len(aInterface))
		for i, v := range aInterface {
			resultingInfo[i] = v.(string)
		}
	}
	return resultingInfo
}

// InterfaceAsIntArray converts an interface to a int array
func InterfaceAsInteger(value interface{}) int {
	incomingDataType := reflect.TypeOf(value).String()
	switch incomingDataType {
	case "string":
		i, _ := strconv.Atoi(value.(string))
		return i
	case "float32":
		return int(value.(float32))
	case "float64":
		return int(value.(float64))
	case "int":
		return value.(int)
	default:
		log.Fatal("Invalid input for conversion")
		return 0
	}
}
