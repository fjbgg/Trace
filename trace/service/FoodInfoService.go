package service

import (
	"fmt"
	"strconv"
)

type FoodInfoService struct{}

func (fs *FoodInfoService) Trace(traceNumber string) interface{} {
	trace := get_trace(traceNumber)
	return trace
}
func (fs *FoodInfoService) Get_food(traceNumber string) map[string]interface{} {
	response := make(map[string]interface{})
	name, current, quality, _, err := Tracecli.GetTraceinfoByTraceNumber(traceNumber)
	byTraceNumber, i, i2, _, err := Tracecli.GetTraceDetailByTraceNumber(traceNumber)
	if err != nil {
		fmt.Println(err)
	}
	response["timestamp"] = byTraceNumber[len(byTraceNumber)-1]
	response["produce"] = i[0]
	response["name"] = name
	response["current"] = current
	response["address"] = i2[0]
	response["quality"] = quality
	return response
}
func (fs *FoodInfoService) GetFoodList() []int {
	food, _ := Tracecli.GetAllFood()
	intSlice := make([]int, len(food))
	for i, v := range food {
		intSlice[i] = int(v.Int64())
	}
	return intSlice
}

func get_trace(traceNumber string) []interface{} {
	_, s, _, _, err := Tracecli.GetTraceinfoByTraceNumber(traceNumber)
	var responses []interface{}
	byTraceNumber, i, i2, i3, err := Tracecli.GetTraceDetailByTraceNumber(traceNumber)
	if err != nil {
		fmt.Println(err)
	}
	for index := 0; index < len(byTraceNumber); index++ {
		if index == 0 {
			response := make(map[string]interface{})
			response["traceNumber"] = traceNumber
			response["name"] = s
			response["produce_time"] = byTraceNumber[0]
			response["timestamp"] = byTraceNumber[index]
			response["from"] = i[index]
			response["quality"] = i3[index]
			response["from_address"] = i2[index]
			i4 := append(responses, response)
			responses = i4
		} else {
			response := make(map[string]interface{})
			response["traceNumber"] = traceNumber
			response["name"] = s
			response["produce_time"] = byTraceNumber[0]
			response["timestamp"] = byTraceNumber[index]
			response["from"] = i[index-1]
			response["to"] = i[index]
			response["quality"] = i3[index]
			response["from_address"] = i2[index-1]
			i4 := append(responses, response)
			responses = i4
		}
	}
	return responses
}

func (fs *FoodInfoService) GetTrace(traceNumber string) []interface{} {
	_, s, _, _, err := Tracecli.GetTraceinfoByTraceNumber(traceNumber)
	var responses []interface{}
	byTraceNumber, i, i2, i3, err := Tracecli.GetTraceDetailByTraceNumber(traceNumber)
	if err != nil {
		fmt.Println(err)
	}
	for index := 0; index < len(byTraceNumber); index++ {
		if index == 0 {
			response := make(map[string]interface{})
			response["traceNumber"] = traceNumber
			response["name"] = s
			response["produce_time"] = byTraceNumber[0]
			response["timestamp"] = byTraceNumber[index]
			response["from"] = i[index]
			response["quality"] = i3[index]
			response["from_address"] = i2[index]
			i4 := append(responses, response)
			responses = i4
		} else {
			response := make(map[string]interface{})
			response["traceNumber"] = traceNumber
			response["name"] = s
			response["produce_time"] = byTraceNumber[0]
			response["timestamp"] = byTraceNumber[index]
			response["from"] = i[index-1]
			response["quality"] = i3[index]
			response["from_address"] = i2[index-1]
			i4 := append(responses, response)
			responses = i4
		}
	}
	return responses
}
func (f FoodInfoService) Producing() []interface{} {
	var resList []interface{}
	numList := f.GetFoodList()
	for index := 0; index < len(numList); index++ {
		trace := f.GetTrace(strconv.Itoa(numList[index]))
		if len(trace) == 1 {
			i := append(resList, trace[0])
			resList = i
		}
	}
	fmt.Println("Producing: ", resList)
	return resList
}
func (f FoodInfoService) Distributing() []interface{} {
	numList := f.GetFoodList()
	var resList []interface{}
	for index := 0; index < len(numList); index++ {
		trace := f.GetTrace(strconv.Itoa(numList[index]))
		if len(trace) == 2 {
			i := append(resList, trace[1])
			resList = i
		}
	}
	return resList
}
func (f FoodInfoService) Retailing() []interface{} {
	numList := f.GetFoodList()
	var resList []interface{}
	for index := 0; index < len(numList); index++ {
		trace := f.GetTrace(strconv.Itoa(numList[index]))
		if len(trace) == 3 {
			i := append(resList, trace[2])
			resList = i
		}
	}
	return resList
}
