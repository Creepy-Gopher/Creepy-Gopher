package parser

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

func printStructFields(s interface{}) {
	val := reflect.ValueOf(s)
	typ := val.Type()

	if typ.Kind() != reflect.Struct {
		fmt.Println("Provided value is not a struct")
		return
	}

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i)
		fmt.Printf("%s: %v\n", field.Name, value)
	}
}

func handlerConvertablePrice(content string) (uint64, error) {
	if content == "اجاره رایگان" {
		return 0, nil
	}
	parts := strings.Split(content, " ")
	price, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, err
	}
	if parts[1] == "میلیارد" {
		price = price * 1000
	}
	return uint64(price), nil
}

func handlerFloor(content string) (uint, error) {
	parts := strings.Split(string(content), " از ")
	num, err := strconv.ParseUint(parts[0], 10, 0)
	if err == nil {
		return uint(num), nil
	}
	return 0, err
}

func handlerArea(content string) (uint64, error) {
	num, err := strconv.ParseUint(content, 10, 64)
	if err == nil {
		return uint64(uint(num)), nil
	}
	return 0, err
}

func hanldeNonConvertablePrices(content string) (uint64, error) {
	parts := strings.Split(string(content), " ")
	priceSTR := parts[0]
	priceSTR = strings.ReplaceAll(priceSTR, "٬", "")
	price, err := strconv.Atoi(priceSTR)
	price /= 1000000
	if err != nil {
		return 0, err
	}
	return uint64(price), nil
}
