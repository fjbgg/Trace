package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"trace/conf"
	"trace/utils"
)

type User struct {
	beego.Controller
}

func (u User) Userinof() {
	userName := u.GetString("userName")
	var jsons []byte
	switch userName {
	case "producer":
		jsons = utils.Send_custom_json(200, "address", conf.ProducerAddress)
	case "distributor":
		jsons = utils.Send_custom_json(200, "address", conf.DistributorAddress)
	case "retailer":
		jsons = utils.Send_custom_json(200, "address", conf.RetailerAddress)
	default:
		jsons = utils.Send_custom_json(200, "error", "user not found")
	}
	u.Ctx.ResponseWriter.Write(jsons)
}
