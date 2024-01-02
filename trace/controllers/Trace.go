package controllers

import (
	"encoding/json"
	"math/big"
	"strconv"
	"trace/models"
	"trace/service"
	"trace/utils"

	beego "github.com/beego/beego/v2/server/web"
)

type Trace struct {
	beego.Controller
}

func (t Trace) NewFood() {
	t.newfood()
}
func (t Trace) Adddistribution() {
	t.adddistribution()
}
func (t Trace) Addretail() {
	t.addretail()
}
func(t Trace) Produce(){
		t.produce()
}
func (t Trace) produce() {
	var food models.Food
	data := t.Ctx.Input.RequestBody
	b, i := paramjson(data, &food)
	if !b {
		t.Ctx.ResponseWriter.Write(utils.Send_json(0, "传入参数不正确", i))
		return  // 添加此 return 语句以在参数无效时退出函数
	}
	service.LoadUser("producer")
	_, receipt, err := service.Tracecli.NewFood(food.Name, food.TraceNumber, food.TraceName, strconv.Itoa(int(food.Quality)))
	if err != nil {
		t.Ctx.ResponseWriter.Write(utils.Send_json(0, "traceNumber already exist", err))
		return  // 添加此 return 语句以在出现错误时退出函数
	}
	t.Ctx.ResponseWriter.Write(utils.Send_json(1, "ok", receipt))
}

func (t Trace) newfood() {
	var food models.Food
	data := t.Ctx.Input.RequestBody
	paramjson(data, &food)
	b, i := paramjson(data, &food)
	if !b {
		t.Ctx.ResponseWriter.Write(utils.Send_json(0, "传入参数不正确", i))
	}
	service.LoadUser("producer")
	_, receipt, err := service.Tracecli.NewFood(food.Name, food.TraceNumber, food.TraceName, strconv.Itoa(int(food.Quality)))
	if err != nil {
		t.Ctx.ResponseWriter.Write(utils.Send_json(0, "traceNumber already exist", err))
	}
	t.Ctx.ResponseWriter.Write(utils.Send_json(1, "ok", receipt))
}

func (t Trace) adddistribution() {
	var food models.Food
	data := t.Ctx.Input.RequestBody
	paramjson(data, &food)
	b, i := paramjson(data, &food)
	if !b {
		t.Ctx.ResponseWriter.Write(utils.Send_json(0, "传入参数不正确", i))
	}
	service.LoadUser("distributor")
	_, receipt, err := service.Tracecli.AddTraceInfoByDistributor(food.TraceNumber, food.TraceName, new(big.Int).SetUint64(uint64(food.Quality)))
	if err != nil {
		t.Ctx.ResponseWriter.Write(utils.Send_json(0, "error", err))
	}
	t.Ctx.ResponseWriter.Write(utils.Send_json(1, "ok", receipt))
}
func (t Trace) addretail() {
	var food models.Food
	data := t.Ctx.Input.RequestBody
	paramjson(data, &food)
	b, i := paramjson(data, &food)
	if !b {
		t.Ctx.ResponseWriter.Write(utils.Send_json(0, "传入参数不正确", i))
	}
	service.LoadUser("retailer")
	_, receipt, err := service.Tracecli.AddTraceInfoByRetailer(food.TraceNumber, food.TraceName, new(big.Int).SetUint64(uint64(food.Quality)))
	if err != nil {
		t.Ctx.ResponseWriter.Write(utils.Send_json(0, "error", err))
	}
	t.Ctx.ResponseWriter.Write(utils.Send_json(1, "ok", receipt))
}
func paramjson(data []byte, food *models.Food) (bool, models.Food) {
	if data == nil {
			return false, *food
	}
	if err := json.Unmarshal(data, food); err != nil {
			return false, *food
	} else {
			return true, *food
	}
}