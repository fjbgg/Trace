package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
	"trace/service"
	"trace/utils"
)

type GetFoodInfo struct {
	beego.Controller
}

// 获取某个食品的溯源信息
func (g GetFoodInfo) Trace() {
	tn, fs := g.getTraceNumberAndService()
	g.Ctx.Output.JSON(fs.Trace(tn), false, false)
}

// 获取某个食品的当前信息
func (g GetFoodInfo) GetFood() {
	tn, fs := g.getTraceNumberAndService()
	g.Ctx.Output.JSON(fs.Get_food(tn), false, false)
}

// 获取位于生产商的的食物信息
func (g GetFoodInfo) Producing() {
	fs := g.getFoodInfoService()
	g.Ctx.Output.JSON(fs.Producing(), false, false)
}

// 获取位于中间商的食物信息
func (g GetFoodInfo) Distributing() {
	fs := g.getFoodInfoService()
	g.Ctx.Output.JSON(fs.Distributing(), false, false)
}

// 获取位于超市的食物信息
func (g GetFoodInfo) Retailing() {
	fs := g.getFoodInfoService()
	g.Ctx.Output.JSON(fs.Retailing(), false, false)
}

func (g GetFoodInfo) getFoodInfoService() *service.FoodInfoService {
	fs := &service.FoodInfoService{}
	return fs
}

func (g GetFoodInfo) getTraceNumberAndService() (string, *service.FoodInfoService) {
	tn := g.GetString("traceNumber")
	atoi, _ := strconv.Atoi(tn)
	if atoi < 0 || tn == "" {
		g.Ctx.ResponseWriter.Write(utils.Send_custom_json(0, "error", "invalid parameter"))
	}
	fs := &service.FoodInfoService{}
	return tn, fs
}
