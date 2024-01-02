package conf

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

var (
	ContractAddress    string
	DistributorAddress string
	DistributorPK      string
	ProducerAddress    string
	ProducerPK         string
	RetailerAddress    string
	RetailerPK         string
)

func init() {
	err := beego.LoadAppConfig("ini", "conf/app.conf")
	if err != nil {
		fmt.Println("Failed to load config file:", err)
		return
	}
	ContractAddress, _ = beego.AppConfig.String("contractAddress")
	DistributorAddress, _ = beego.AppConfig.String("distributorAddress")
	DistributorPK, _ = beego.AppConfig.String("distributorPK")
	ProducerAddress, _ = beego.AppConfig.String("producerAddress")
	ProducerPK, _ = beego.AppConfig.String("producerPK")
	RetailerAddress, _ = beego.AppConfig.String("retailerAddress")
	RetailerPK, _ = beego.AppConfig.String("retailerPK")

}
