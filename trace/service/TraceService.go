package service

import (
	"fmt"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/client"
	conf2 "github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	//"github.com/FISCO-BCOS/go-sdk/client"
	"trace/conf"
)

var Tracecli TraceSession

func init() {
	configs, err := conf.ParseConfigFile("config.toml")
	if err != nil {
		log.Fatalf("ParseConfigFile failed, err: %v", err)
	}
	client, err := client.Dial((*conf2.Config)(&configs[0]))
	if err != nil {
		log.Fatal(err)
	}
	address := conf.ContractAddress

	contractAddress := common.HexToAddress(address)
	instance, err := NewTrace(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	Tracecli = TraceSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()}
}
func LoadUser(user string) {
	fromAddress := ""
	privateKey := ""
	switch user {
	case "producer":
		fromAddress = conf.ProducerAddress
		privateKey = conf.ProducerPK
	case "distributor":
		fromAddress = conf.DistributorAddress
		privateKey = conf.DistributorPK
	case "retailer":
		fromAddress = conf.RetailerAddress
		privateKey = conf.RetailerPK
	default:
		fmt.Println("not foun role")
		return
	}
	key, _ := crypto.HexToECDSA(privateKey)
	auth := bind.NewKeyedTransactor(key)
	auth.From = common.HexToAddress(fromAddress)
	files, _ := conf.ParseConfigFile("config.toml")
	client, _ := client.Dial((*conf2.Config)(&files[0]))
	instance, _ := NewTrace(common.HexToAddress(conf.ContractAddress), client)
	Tracecli = TraceSession{Contract: instance, CallOpts: *client.GetCallOpts(), TransactOpts: *auth}
}
