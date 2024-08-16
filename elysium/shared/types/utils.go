package types

import (
	"encoding/json"
	"fmt"
	"golang.org/x/xerrors"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func Analyze(bytes []byte) (Prototype, error) {

	primo := Primordial{}
	if err := json.Unmarshal(bytes, &primo); err != nil {
		return nil, err
	}

	switch primo.Subject {
	case SyntheticStudent{}.GetSubject():
		student := SyntheticStudent{}
		if err := json.Unmarshal(bytes, &student); err != nil {
			return nil, err
		}
		return student, nil
	default:
		return nil, xerrors.Errorf("error unknown prototype: %s", primo.Subject)
	}
}

func ToInsertValue(prototype Prototype) string {

	fields := reflect.TypeOf(prototype)
	values := reflect.ValueOf(prototype)
	insertValues := make([]string, 0)
	for idx := 0; idx < fields.NumField(); idx++ {
		field := fields.Field(idx)
		value := values.Field(idx)
		switch field.Tag.Get("ctype") {
		case "String":
			insertValues = append(insertValues, fmt.Sprintf("'%s'", value.String()))
		case "Array(String)":
			elements := make([]string, 0)
			v := value.Interface().([]string)
			for _, element := range v {
				elements = append(elements, fmt.Sprintf("'%s'", element))
			}
			insertValues = append(insertValues, fmt.Sprintf("[%s]", strings.Join(elements, ",")))
		case "Array(UInt8)", "Array(UInt32)", "Array(UInt64)":
			elements := make([]string, 0)
			v := value.Interface().([]int)
			for _, element := range v {
				elements = append(elements, strconv.Itoa(element))
			}
			insertValues = append(insertValues, strings.Join(elements, ","))
		case "DateTime", "DateTime32", "DateTime64":
			v := value.Interface().(time.Time)
			insertValues = append(insertValues, fmt.Sprintf("parseDateTimeBestEffort('%s')", v.UTC().String()))
		case "Date":
			v := value.Interface().(time.Time)
			insertValues = append(
				insertValues, fmt.Sprintf("parseDateTimeBestEffort('%d-%d-%d')", v.Year(), v.Month(), v.Day()),
			)
		case "UInt8", "UInt32", "UInt64":
			v := int(value.Uint())
			insertValues = append(insertValues, strconv.Itoa(v))
		case "Int8", "Int32", "Int64":
			v := int(value.Int())
			insertValues = append(insertValues, strconv.Itoa(v))
		case "Float8", "Float32", "Float64":
			v := value.Float()
			insertValues = append(insertValues, fmt.Sprintf("%f", v))
		}
	}

	return fmt.Sprintf("(%s)", strings.Join(insertValues, ","))
}
