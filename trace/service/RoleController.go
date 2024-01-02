// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package service

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
)

// ServiceABI is the input ABI used to generate the binding from.
const ServiceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_user_id\",\"type\":\"string\"}],\"name\":\"resetDRRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_user_id\",\"type\":\"string\"}],\"name\":\"resetRRRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_user_id\",\"type\":\"string\"}],\"name\":\"setDRRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_user_id\",\"type\":\"string\"}],\"name\":\"setRRRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_user_id\",\"type\":\"string\"}],\"name\":\"resetPRRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_user_id\",\"type\":\"string\"}],\"name\":\"setPRRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"bool\"}],\"name\":\"SetPRRoleResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"bool\"}],\"name\":\"ResetPRRoleResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"bool\"}],\"name\":\"SetDRRoleResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"bool\"}],\"name\":\"ResetDRRoleResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"bool\"}],\"name\":\"SetRRRoleResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"\",\"type\":\"bool\"}],\"name\":\"ResetRRRoleResult\",\"type\":\"event\"}]"

// ServiceBin is the compiled bytecode used for deploying new contracts.
var ServiceBin = "0x608060405234801561001057600080fd5b5061001961007a565b604051809103906000f080158015610035573d6000803e3d6000fd5b506000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061008a565b60405161169380610c9183390190565b610bf8806100996000396000f300608060405260043610610078576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806303ea78c81461007d578063300a4420146100ba5780634c1cf480146100f75780636d8edf7c146101345780638b85e2b214610171578063b25fcc2e146101ae575b600080fd5b34801561008957600080fd5b506100a4600480360361009f9190810190610a42565b6101eb565b6040516100b19190610af6565b60405180910390f35b3480156100c657600080fd5b506100e160048036036100dc9190810190610a42565b610338565b6040516100ee9190610af6565b60405180910390f35b34801561010357600080fd5b5061011e60048036036101199190810190610a42565b610485565b60405161012b9190610af6565b60405180910390f35b34801561014057600080fd5b5061015b60048036036101569190810190610a42565b6105d2565b6040516101689190610af6565b60405180910390f35b34801561017d57600080fd5b5061019860048036036101939190810190610a42565b61071f565b6040516101a59190610af6565b60405180910390f35b3480156101ba57600080fd5b506101d560048036036101d09190810190610a42565b61086c565b6040516101e29190610af6565b60405180910390f35b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a63430126040805190810160405280600281526020017f445200000000000000000000000000000000000000000000000000000000000081525086866040518463ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161029d93929190610b11565b602060405180830381600087803b1580156102b757600080fd5b505af11580156102cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506102ef9190810190610a19565b90507f6f36176fd484d76b5a044beb372e479916909ea1d42b94a39cbe98d4ec5a7d72600082136040516103239190610af6565b60405180910390a16000811391505092915050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a63430126040805190810160405280600281526020017f525200000000000000000000000000000000000000000000000000000000000081525086866040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016103ea93929190610b11565b602060405180830381600087803b15801561040457600080fd5b505af1158015610418573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061043c9190810190610a19565b90507f2f43c47684e93292df983b142ac84564c8171798543d1e4c058a00b8cad71c1c600082136040516104709190610af6565b60405180910390a16000811391505092915050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd517fae6040805190810160405280600281526020017f445200000000000000000000000000000000000000000000000000000000000081525086866040518463ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161053793929190610b11565b602060405180830381600087803b15801561055157600080fd5b505af1158015610565573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506105899190810190610a19565b90507f266b0a5b2ad2251976f3606e8b06066e6ce8379e6ee3eeda1f934241d86dd7fa600082136040516105bd9190610af6565b60405180910390a16000811391505092915050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd517fae6040805190810160405280600281526020017f525200000000000000000000000000000000000000000000000000000000000081525086866040518463ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161068493929190610b11565b602060405180830381600087803b15801561069e57600080fd5b505af11580156106b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506106d69190810190610a19565b90507f938928b1b88fe5ccc23a27f6112852db8f7d30dd5df882d4afcc570a4ee2fac56000821360405161070a9190610af6565b60405180910390a16000811391505092915050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a63430126040805190810160405280600281526020017f505200000000000000000000000000000000000000000000000000000000000081525086866040518463ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016107d193929190610b11565b602060405180830381600087803b1580156107eb57600080fd5b505af11580156107ff573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506108239190810190610a19565b90507fd754aa3299490ca521803b63cc8041d8ac65b00182c5b50007c26d11195c57df600082136040516108579190610af6565b60405180910390a16000811391505092915050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd517fae6040805190810160405280600281526020017f505200000000000000000000000000000000000000000000000000000000000081525086866040518463ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161091e93929190610b11565b602060405180830381600087803b15801561093857600080fd5b505af115801561094c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506109709190810190610a19565b90507ff6a07132b4cd39d748fc7677eb7b47b35f1287ba285e018e8cf288f29b5f142e600082136040516109a49190610af6565b60405180910390a16000811391505092915050565b60006109c58251610b61565b905092915050565b60008083601f84011215156109e157600080fd5b8235905067ffffffffffffffff8111156109fa57600080fd5b602083019150836001820283011115610a1257600080fd5b9250929050565b600060208284031215610a2b57600080fd5b6000610a39848285016109b9565b91505092915050565b60008060208385031215610a5557600080fd5b600083013567ffffffffffffffff811115610a6f57600080fd5b610a7b858286016109cd565b92509250509250929050565b610a9081610b55565b82525050565b6000828452602084019350610aac838584610b6b565b610ab583610bad565b840190509392505050565b6000610acb82610b4a565b808452610adf816020860160208601610b7a565b610ae881610bad565b602085010191505092915050565b6000602082019050610b0b6000830184610a87565b92915050565b60006040820190508181036000830152610b2b8186610ac0565b90508181036020830152610b40818486610a96565b9050949350505050565b600081519050919050565b60008115159050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610b98578082015181840152602081019050610b7d565b83811115610ba7576000848401525b50505050565b6000601f19601f83011690509190505600a265627a7a7230582014f6fb3b2f8ae923ee18de83738f50457027c59fafa81331b88ef7b999934d4c6c6578706572696d656e74616cf5003760806040523480156200001157600080fd5b506110016000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166356004b6a6040805190810160405280600781526020017f74785f726f6c65000000000000000000000000000000000000000000000000008152506040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016200010191906200024a565b602060405180830381600087803b1580156200011c57600080fd5b505af115801562000131573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525062000157919081019062000174565b50620002f4565b60006200016c8251620002a3565b905092915050565b6000602082840312156200018757600080fd5b600062000197848285016200015e565b91505092915050565b6000620001ad8262000298565b808452620001c3816020860160208601620002ad565b620001ce81620002e3565b602085010191505092915050565b6000600482527f726f6c65000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000600982527f757365725f6164647200000000000000000000000000000000000000000000006020830152604082019050919050565b60006060820190508181036000830152620002668184620001a0565b905081810360208301526200027b81620001dc565b90508181036040830152620002908162000213565b905092915050565b600081519050919050565b6000819050919050565b60005b83811015620002cd578082015181840152602081019050620002b0565b83811115620002dd576000848401525b50505050565b6000601f19601f8301169050919050565b61138f80620003046000396000f300608060405260043610610062576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806310158b891461006757806324e0791314610090578063a6343012146100cd578063dd517fae1461010a575b600080fd5b34801561007357600080fd5b5061008e60048036036100899190810190610f58565b610147565b005b34801561009c57600080fd5b506100b760048036036100b29190810190610f58565b610196565b6040516100c491906110da565b60405180910390f35b3480156100d957600080fd5b506100f460048036036100ef9190810190610f58565b6102ac565b60405161010191906110f5565b60405180910390f35b34801561011657600080fd5b50610131600480360361012c9190810190610f58565b6105c8565b60405161013e91906110f5565b60405180910390f35b6101518282610196565b1515610192576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610189906111c7565b60405180910390fd5b5050565b6000806000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f23f63c96040805190810160405280600781526020017f74785f726f6c65000000000000000000000000000000000000000000000000008152506040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016102449190611110565b602060405180830381600087803b15801561025e57600080fd5b505af1158015610272573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506102969190810190610f06565b90506102a3818585610908565b91505092915050565b600080600080600092506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f23f63c96040805190810160405280600781526020017f74785f726f6c65000000000000000000000000000000000000000000000000008152506040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016103619190611110565b602060405180830381600087803b15801561037b57600080fd5b505af115801561038f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506103b39190810190610f06565b91508173ffffffffffffffffffffffffffffffffffffffff16637857d7c96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561041957600080fd5b505af115801561042d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506104519190810190610e8b565b90508073ffffffffffffffffffffffffffffffffffffffff1663cd30a1d1866040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016104a89190611192565b600060405180830381600087803b1580156104c257600080fd5b505af11580156104d6573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff166328bb211787836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610531929190611132565b602060405180830381600087803b15801561054b57600080fd5b505af115801561055f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506105839190810190610f2f565b92507fee205d8383de0d8cc6823192e93e3086dd317c369858ed095ed2d719d25bfd23836040516105b491906110f5565b60405180910390a182935050505092915050565b600080600080600092506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f23f63c96040805190810160405280600781526020017f74785f726f6c65000000000000000000000000000000000000000000000000008152506040518263ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040161067d9190611110565b602060405180830381600087803b15801561069757600080fd5b505af11580156106ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506106cf9190810190610f06565b91506106da86610b81565b1580156106ef57506106ed828787610908565b155b156108c5578173ffffffffffffffffffffffffffffffffffffffff166313db93466040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561075857600080fd5b505af115801561076c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506107909190810190610edd565b90508073ffffffffffffffffffffffffffffffffffffffff1663e942b516866040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016107e79190611192565b600060405180830381600087803b15801561080157600080fd5b505af1158015610815573d6000803e3d6000fd5b505050508173ffffffffffffffffffffffffffffffffffffffff166331afac3687836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610870929190611162565b602060405180830381600087803b15801561088a57600080fd5b505af115801561089e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506108c29190810190610f2f565b92505b7f14a121ebed99a8737803a36f9a293fa68078956202e0a990b3d04694d93b71e4836040516108f491906110f5565b60405180910390a182935050505092915050565b6000808473ffffffffffffffffffffffffffffffffffffffff16637857d7c96040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15801561096f57600080fd5b505af1158015610983573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506109a79190810190610e8b565b90508073ffffffffffffffffffffffffffffffffffffffff1663cd30a1d1846040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004016109fe9190611192565b600060405180830381600087803b158015610a1857600080fd5b505af1158015610a2c573d6000803e3d6000fd5b5050505060008573ffffffffffffffffffffffffffffffffffffffff1663e8434e3986846040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401610a89929190611132565b602060405180830381600087803b158015610aa357600080fd5b505af1158015610ab7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610adb9190810190610eb4565b73ffffffffffffffffffffffffffffffffffffffff1663949d225d6040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b158015610b3e57600080fd5b505af1158015610b52573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610b769190810190610f2f565b139150509392505050565b60006060600080849250600083511415610b9e5760019350610dc9565b600091505b8251821015610dc4578282815181101515610bba57fe5b9060200101517f010000000000000000000000000000000000000000000000000000000000000090047f010000000000000000000000000000000000000000000000000000000000000002905060207f010000000000000000000000000000000000000000000000000000000000000002817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614158015610cc3575060097f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614155b8015610d365750600a7f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614155b8015610da95750600d7f0100000000000000000000000000000000000000000000000000000000000000027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614155b15610db75760009350610dc9565b8180600101925050610ba3565b600193505b505050919050565b6000610ddd825161128c565b905092915050565b6000610df1825161129e565b905092915050565b6000610e0582516112b0565b905092915050565b6000610e1982516112c2565b905092915050565b6000610e2d82516112d4565b905092915050565b600082601f8301121515610e4857600080fd5b8135610e5b610e5682611214565b6111e7565b91508082526020830160208301858383011115610e7757600080fd5b610e82838284611302565b50505092915050565b600060208284031215610e9d57600080fd5b6000610eab84828501610dd1565b91505092915050565b600060208284031215610ec657600080fd5b6000610ed484828501610de5565b91505092915050565b600060208284031215610eef57600080fd5b6000610efd84828501610df9565b91505092915050565b600060208284031215610f1857600080fd5b6000610f2684828501610e0d565b91505092915050565b600060208284031215610f4157600080fd5b6000610f4f84828501610e21565b91505092915050565b60008060408385031215610f6b57600080fd5b600083013567ffffffffffffffff811115610f8557600080fd5b610f9185828601610e35565b925050602083013567ffffffffffffffff811115610fae57600080fd5b610fba85828601610e35565b9150509250929050565b610fcd81611276565b82525050565b610fdc816112de565b82525050565b610feb816112f0565b82525050565b610ffa81611282565b82525050565b600061100b8261124b565b80845261101f816020860160208601611311565b61102881611344565b602085010191505092915050565b600061104182611240565b808452611055816020860160208601611311565b61105e81611344565b602085010191505092915050565b6000600982527f757365725f6164647200000000000000000000000000000000000000000000006020830152604082019050919050565b6000601382527f6e6f742068617665207065726d697373696f6e000000000000000000000000006020830152604082019050919050565b60006020820190506110ef6000830184610fc4565b92915050565b600060208201905061110a6000830184610ff1565b92915050565b6000602082019050818103600083015261112a8184611036565b905092915050565b6000604082019050818103600083015261114c8185611000565b905061115b6020830184610fd3565b9392505050565b6000604082019050818103600083015261117c8185611000565b905061118b6020830184610fe2565b9392505050565b600060408201905081810360008301526111ab8161106c565b905081810360208301526111bf8184611000565b905092915050565b600060208201905081810360008301526111e0816110a3565b9050919050565b6000604051905081810181811067ffffffffffffffff8211171561120a57600080fd5b8060405250919050565b600067ffffffffffffffff82111561122b57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60008115159050919050565b6000819050919050565b600061129782611256565b9050919050565b60006112a982611256565b9050919050565b60006112bb82611256565b9050919050565b60006112cd82611256565b9050919050565b6000819050919050565b60006112e982611256565b9050919050565b60006112fb82611256565b9050919050565b82818337600083830152505050565b60005b8381101561132f578082015181840152602081019050611314565b8381111561133e576000848401525b50505050565b6000601f19601f83011690509190505600a265627a7a7230582006dd5c2ee889d3ac696c75645dfd02ea4e3914ce0272f1c8d88a86e8abd075656c6578706572696d656e74616cf50037"

// DeployService deploys a new contract, binding an instance of Service to it.
func DeployService(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Service, error) {
	parsed, err := abi.JSON(strings.NewReader(ServiceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ServiceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Service{ServiceCaller: ServiceCaller{contract: contract}, ServiceTransactor: ServiceTransactor{contract: contract}, ServiceFilterer: ServiceFilterer{contract: contract}}, nil
}

func AsyncDeployService(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(ServiceABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(ServiceBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// Service is an auto generated Go binding around a Solidity contract.
type Service struct {
	ServiceCaller     // Read-only binding to the contract
	ServiceTransactor // Write-only binding to the contract
	ServiceFilterer   // Log filterer for contract events
}

// ServiceCaller is an auto generated read-only Go binding around a Solidity contract.
type ServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceTransactor is an auto generated write-only Go binding around a Solidity contract.
type ServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type ServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type ServiceSession struct {
	Contract     *Service          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ServiceCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type ServiceCallerSession struct {
	Contract *ServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ServiceTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type ServiceTransactorSession struct {
	Contract     *ServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ServiceRaw is an auto generated low-level Go binding around a Solidity contract.
type ServiceRaw struct {
	Contract *Service // Generic contract binding to access the raw methods on
}

// ServiceCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type ServiceCallerRaw struct {
	Contract *ServiceCaller // Generic read-only contract binding to access the raw methods on
}

// ServiceTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type ServiceTransactorRaw struct {
	Contract *ServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewService creates a new instance of Service, bound to a specific deployed contract.
func NewService(address common.Address, backend bind.ContractBackend) (*Service, error) {
	contract, err := bindService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Service{ServiceCaller: ServiceCaller{contract: contract}, ServiceTransactor: ServiceTransactor{contract: contract}, ServiceFilterer: ServiceFilterer{contract: contract}}, nil
}

// NewServiceCaller creates a new read-only instance of Service, bound to a specific deployed contract.
func NewServiceCaller(address common.Address, caller bind.ContractCaller) (*ServiceCaller, error) {
	contract, err := bindService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ServiceCaller{contract: contract}, nil
}

// NewServiceTransactor creates a new write-only instance of Service, bound to a specific deployed contract.
func NewServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*ServiceTransactor, error) {
	contract, err := bindService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ServiceTransactor{contract: contract}, nil
}

// NewServiceFilterer creates a new log filterer instance of Service, bound to a specific deployed contract.
func NewServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*ServiceFilterer, error) {
	contract, err := bindService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ServiceFilterer{contract: contract}, nil
}

// bindService binds a generic wrapper to an already deployed contract.
func bindService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Service *ServiceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Service.Contract.ServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Service *ServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Service.Contract.ServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Service *ServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Service.Contract.ServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Service *ServiceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Service.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Service *ServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _Service.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Service *ServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _Service.Contract.contract.Transact(opts, method, params...)
}

// ResetDRRole is a paid mutator transaction binding the contract method 0x03ea78c8.
//
// Solidity: function resetDRRole(string _user_id) returns(bool)
func (_Service *ServiceTransactor) ResetDRRole(opts *bind.TransactOpts, _user_id string) (*types.Transaction, *types.Receipt, error) {
	return _Service.contract.Transact(opts, "resetDRRole", _user_id)
}

func (_Service *ServiceTransactor) AsyncResetDRRole(handler func(*types.Receipt, error), opts *bind.TransactOpts, _user_id string) (*types.Transaction, error) {
	return _Service.contract.AsyncTransact(opts, handler, "resetDRRole", _user_id)
}

// ResetDRRole is a paid mutator transaction binding the contract method 0x03ea78c8.
//
// Solidity: function resetDRRole(string _user_id) returns(bool)
func (_Service *ServiceSession) ResetDRRole(_user_id string) (*types.Transaction, *types.Receipt, error) {
	return _Service.Contract.ResetDRRole(&_Service.TransactOpts, _user_id)
}

func (_Service *ServiceSession) AsyncResetDRRole(handler func(*types.Receipt, error), _user_id string) (*types.Transaction, error) {
	return _Service.Contract.AsyncResetDRRole(handler, &_Service.TransactOpts, _user_id)
}

// ResetDRRole is a paid mutator transaction binding the contract method 0x03ea78c8.
//
// Solidity: function resetDRRole(string _user_id) returns(bool)
func (_Service *ServiceTransactorSession) ResetDRRole(_user_id string) (*types.Transaction, *types.Receipt, error) {
	return _Service.Contract.ResetDRRole(&_Service.TransactOpts, _user_id)
}

func (_Service *ServiceTransactorSession) AsyncResetDRRole(handler func(*types.Receipt, error), _user_id string) (*types.Transaction, error) {
	return _Service.Contract.AsyncResetDRRole(handler, &_Service.TransactOpts, _user_id)
}

// ResetPRRole is a paid mutator transaction binding the contract method 0x8b85e2b2.
//
// Solidity: function resetPRRole(string _user_id) returns(bool)
func (_Service *ServiceTransactor) ResetPRRole(opts *bind.TransactOpts, _user_id string) (*types.Transaction, *types.Receipt, error) {
	return _Service.contract.Transact(opts, "resetPRRole", _user_id)
}

func (_Service *ServiceTransactor) AsyncResetPRRole(handler func(*types.Receipt, error), opts *bind.TransactOpts, _user_id string) (*types.Transaction, error) {
	return _Service.contract.AsyncTransact(opts, handler, "resetPRRole", _user_id)
}

// ResetPRRole is a paid mutator transaction binding the contract method 0x8b85e2b2.
//
// Solidity: function resetPRRole(string _user_id) returns(bool)
func (_Service *ServiceSession) ResetPRRole(_user_id string) (*types.Transaction, *types.Receipt, error) {
	return _Service.Contract.ResetPRRole(&_Service.TransactOpts, _user_id)
}

func (_Service *ServiceSession) AsyncResetPRRole(handler func(*types.Receipt, error), _user_id string) (*types.Transaction, error) {
	return _Service.Contract.AsyncResetPRRole(handler, &_Service.TransactOpts, _user_id)
}

// ResetPRRole is a paid mutator transaction binding the contract method 0x8b85e2b2.
//
// Solidity: function resetPRRole(string _user_id) returns(bool)
func (_Service *ServiceTransactorSession) ResetPRRole(_user_id string) (*types.Transaction, *types.Receipt, error) {
	return _Service.Contract.ResetPRRole(&_Service.TransactOpts, _user_id)
}

func (_Service *ServiceTransactorSession) AsyncResetPRRole(handler func(*types.Receipt, error), _user_id string) (*types.Transaction, error) {
	return _Service.Contract.AsyncResetPRRole(handler, &_Service.TransactOpts, _user_id)
}

// ResetRRRole is a paid mutator transaction binding the contract method 0x300a4420.
//
// Solidity: function resetRRRole(string _user_id) returns(bool)
func (_Service *ServiceTransactor) ResetRRRole(opts *bind.TransactOpts, _user_id string) (*types.Transaction, *types.Receipt, error) {
	return _Service.contract.Transact(opts, "resetRRRole", _user_id)
}

func (_Service *ServiceTransactor) AsyncResetRRRole(handler func(*types.Receipt, error), opts *bind.TransactOpts, _user_id string) (*types.Transaction, error) {
	return _Service.contract.AsyncTransact(opts, handler, "resetRRRole", _user_id)
}

// ResetRRRole is a paid mutator transaction binding the contract method 0x300a4420.
//
// Solidity: function resetRRRole(string _user_id) returns(bool)
func (_Service *ServiceSession) ResetRRRole(_user_id string) (*types.Transaction, *types.Receipt, error) {
	return _Service.Contract.ResetRRRole(&_Service.TransactOpts, _user_id)
}

func (_Service *ServiceSession) AsyncResetRRRole(handler func(*types.Receipt, error), _user_id string) (*types.Transaction, error) {
	return _Service.Contract.AsyncResetRRRole(handler, &_Service.TransactOpts, _user_id)
}

// ResetRRRole is a paid mutator transaction binding the contract method 0x300a4420.
//
// Solidity: function resetRRRole(string _user_id) returns(bool)
func (_Service *ServiceTransactorSession) ResetRRRole(_user_id string) (*types.Transaction, *types.Receipt, error) {
	return _Service.Contract.ResetRRRole(&_Service.TransactOpts, _user_id)
}

func (_Service *ServiceTransactorSession) AsyncResetRRRole(handler func(*types.Receipt, error), _user_id string) (*types.Transaction, error) {
	return _Service.Contract.AsyncResetRRRole(handler, &_Service.TransactOpts, _user_id)
}

// SetDRRole is a paid mutator transaction binding the contract method 0x4c1cf480.
//
// Solidity: function setDRRole(string _user_id) returns(bool)
func (_Service *ServiceTransactor) SetDRRole(opts *bind.TransactOpts, _user_id string) (*types.Transaction, *types.Receipt, error) {
	return _Service.contract.Transact(opts, "setDRRole", _user_id)
}

func (_Service *ServiceTransactor) AsyncSetDRRole(handler func(*types.Receipt, error), opts *bind.TransactOpts, _user_id string) (*types.Transaction, error) {
	return _Service.contract.AsyncTransact(opts, handler, "setDRRole", _user_id)
}

// SetDRRole is a paid mutator transaction binding the contract method 0x4c1cf480.
//
// Solidity: function setDRRole(string _user_id) returns(bool)
func (_Service *ServiceSession) SetDRRole(_user_id string) (*types.Transaction, *types.Receipt, error) {
	return _Service.Contract.SetDRRole(&_Service.TransactOpts, _user_id)
}

func (_Service *ServiceSession) AsyncSetDRRole(handler func(*types.Receipt, error), _user_id string) (*types.Transaction, error) {
	return _Service.Contract.AsyncSetDRRole(handler, &_Service.TransactOpts, _user_id)
}

// SetDRRole is a paid mutator transaction binding the contract method 0x4c1cf480.
//
// Solidity: function setDRRole(string _user_id) returns(bool)
func (_Service *ServiceTransactorSession) SetDRRole(_user_id string) (*types.Transaction, *types.Receipt, error) {
	return _Service.Contract.SetDRRole(&_Service.TransactOpts, _user_id)
}

func (_Service *ServiceTransactorSession) AsyncSetDRRole(handler func(*types.Receipt, error), _user_id string) (*types.Transaction, error) {
	return _Service.Contract.AsyncSetDRRole(handler, &_Service.TransactOpts, _user_id)
}

// SetPRRole is a paid mutator transaction binding the contract method 0xb25fcc2e.
//
// Solidity: function setPRRole(string _user_id) returns(bool)
func (_Service *ServiceTransactor) SetPRRole(opts *bind.TransactOpts, _user_id string) (*types.Transaction, *types.Receipt, error) {
	return _Service.contract.Transact(opts, "setPRRole", _user_id)
}

func (_Service *ServiceTransactor) AsyncSetPRRole(handler func(*types.Receipt, error), opts *bind.TransactOpts, _user_id string) (*types.Transaction, error) {
	return _Service.contract.AsyncTransact(opts, handler, "setPRRole", _user_id)
}

// SetPRRole is a paid mutator transaction binding the contract method 0xb25fcc2e.
//
// Solidity: function setPRRole(string _user_id) returns(bool)
func (_Service *ServiceSession) SetPRRole(_user_id string) (*types.Transaction, *types.Receipt, error) {
	return _Service.Contract.SetPRRole(&_Service.TransactOpts, _user_id)
}

func (_Service *ServiceSession) AsyncSetPRRole(handler func(*types.Receipt, error), _user_id string) (*types.Transaction, error) {
	return _Service.Contract.AsyncSetPRRole(handler, &_Service.TransactOpts, _user_id)
}

// SetPRRole is a paid mutator transaction binding the contract method 0xb25fcc2e.
//
// Solidity: function setPRRole(string _user_id) returns(bool)
func (_Service *ServiceTransactorSession) SetPRRole(_user_id string) (*types.Transaction, *types.Receipt, error) {
	return _Service.Contract.SetPRRole(&_Service.TransactOpts, _user_id)
}

func (_Service *ServiceTransactorSession) AsyncSetPRRole(handler func(*types.Receipt, error), _user_id string) (*types.Transaction, error) {
	return _Service.Contract.AsyncSetPRRole(handler, &_Service.TransactOpts, _user_id)
}

// SetRRRole is a paid mutator transaction binding the contract method 0x6d8edf7c.
//
// Solidity: function setRRRole(string _user_id) returns(bool)
func (_Service *ServiceTransactor) SetRRRole(opts *bind.TransactOpts, _user_id string) (*types.Transaction, *types.Receipt, error) {
	return _Service.contract.Transact(opts, "setRRRole", _user_id)
}

func (_Service *ServiceTransactor) AsyncSetRRRole(handler func(*types.Receipt, error), opts *bind.TransactOpts, _user_id string) (*types.Transaction, error) {
	return _Service.contract.AsyncTransact(opts, handler, "setRRRole", _user_id)
}

// SetRRRole is a paid mutator transaction binding the contract method 0x6d8edf7c.
//
// Solidity: function setRRRole(string _user_id) returns(bool)
func (_Service *ServiceSession) SetRRRole(_user_id string) (*types.Transaction, *types.Receipt, error) {
	return _Service.Contract.SetRRRole(&_Service.TransactOpts, _user_id)
}

func (_Service *ServiceSession) AsyncSetRRRole(handler func(*types.Receipt, error), _user_id string) (*types.Transaction, error) {
	return _Service.Contract.AsyncSetRRRole(handler, &_Service.TransactOpts, _user_id)
}

// SetRRRole is a paid mutator transaction binding the contract method 0x6d8edf7c.
//
// Solidity: function setRRRole(string _user_id) returns(bool)
func (_Service *ServiceTransactorSession) SetRRRole(_user_id string) (*types.Transaction, *types.Receipt, error) {
	return _Service.Contract.SetRRRole(&_Service.TransactOpts, _user_id)
}

func (_Service *ServiceTransactorSession) AsyncSetRRRole(handler func(*types.Receipt, error), _user_id string) (*types.Transaction, error) {
	return _Service.Contract.AsyncSetRRRole(handler, &_Service.TransactOpts, _user_id)
}

// ServiceResetDRRoleResult represents a ResetDRRoleResult event raised by the Service contract.
type ServiceResetDRRoleResult struct {
	Arg0 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// WatchResetDRRoleResult is a free log subscription operation binding the contract event 0x6f36176fd484d76b5a044beb372e479916909ea1d42b94a39cbe98d4ec5a7d72.
//
// Solidity: event ResetDRRoleResult(bool )
func (_Service *ServiceFilterer) WatchResetDRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.contract.WatchLogs(fromBlock, handler, "ResetDRRoleResult")
}

func (_Service *ServiceFilterer) WatchAllResetDRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.contract.WatchLogs(fromBlock, handler, "ResetDRRoleResult")
}

// ParseResetDRRoleResult is a log parse operation binding the contract event 0x6f36176fd484d76b5a044beb372e479916909ea1d42b94a39cbe98d4ec5a7d72.
//
// Solidity: event ResetDRRoleResult(bool )
func (_Service *ServiceFilterer) ParseResetDRRoleResult(log types.Log) (*ServiceResetDRRoleResult, error) {
	event := new(ServiceResetDRRoleResult)
	if err := _Service.contract.UnpackLog(event, "ResetDRRoleResult", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchResetDRRoleResult is a free log subscription operation binding the contract event 0x6f36176fd484d76b5a044beb372e479916909ea1d42b94a39cbe98d4ec5a7d72.
//
// Solidity: event ResetDRRoleResult(bool )
func (_Service *ServiceSession) WatchResetDRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.Contract.WatchResetDRRoleResult(fromBlock, handler)
}

func (_Service *ServiceSession) WatchAllResetDRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.Contract.WatchAllResetDRRoleResult(fromBlock, handler)
}

// ParseResetDRRoleResult is a log parse operation binding the contract event 0x6f36176fd484d76b5a044beb372e479916909ea1d42b94a39cbe98d4ec5a7d72.
//
// Solidity: event ResetDRRoleResult(bool )
func (_Service *ServiceSession) ParseResetDRRoleResult(log types.Log) (*ServiceResetDRRoleResult, error) {
	return _Service.Contract.ParseResetDRRoleResult(log)
}

// ServiceResetPRRoleResult represents a ResetPRRoleResult event raised by the Service contract.
type ServiceResetPRRoleResult struct {
	Arg0 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// WatchResetPRRoleResult is a free log subscription operation binding the contract event 0xd754aa3299490ca521803b63cc8041d8ac65b00182c5b50007c26d11195c57df.
//
// Solidity: event ResetPRRoleResult(bool )
func (_Service *ServiceFilterer) WatchResetPRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.contract.WatchLogs(fromBlock, handler, "ResetPRRoleResult")
}

func (_Service *ServiceFilterer) WatchAllResetPRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.contract.WatchLogs(fromBlock, handler, "ResetPRRoleResult")
}

// ParseResetPRRoleResult is a log parse operation binding the contract event 0xd754aa3299490ca521803b63cc8041d8ac65b00182c5b50007c26d11195c57df.
//
// Solidity: event ResetPRRoleResult(bool )
func (_Service *ServiceFilterer) ParseResetPRRoleResult(log types.Log) (*ServiceResetPRRoleResult, error) {
	event := new(ServiceResetPRRoleResult)
	if err := _Service.contract.UnpackLog(event, "ResetPRRoleResult", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchResetPRRoleResult is a free log subscription operation binding the contract event 0xd754aa3299490ca521803b63cc8041d8ac65b00182c5b50007c26d11195c57df.
//
// Solidity: event ResetPRRoleResult(bool )
func (_Service *ServiceSession) WatchResetPRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.Contract.WatchResetPRRoleResult(fromBlock, handler)
}

func (_Service *ServiceSession) WatchAllResetPRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.Contract.WatchAllResetPRRoleResult(fromBlock, handler)
}

// ParseResetPRRoleResult is a log parse operation binding the contract event 0xd754aa3299490ca521803b63cc8041d8ac65b00182c5b50007c26d11195c57df.
//
// Solidity: event ResetPRRoleResult(bool )
func (_Service *ServiceSession) ParseResetPRRoleResult(log types.Log) (*ServiceResetPRRoleResult, error) {
	return _Service.Contract.ParseResetPRRoleResult(log)
}

// ServiceResetRRRoleResult represents a ResetRRRoleResult event raised by the Service contract.
type ServiceResetRRRoleResult struct {
	Arg0 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// WatchResetRRRoleResult is a free log subscription operation binding the contract event 0x2f43c47684e93292df983b142ac84564c8171798543d1e4c058a00b8cad71c1c.
//
// Solidity: event ResetRRRoleResult(bool )
func (_Service *ServiceFilterer) WatchResetRRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.contract.WatchLogs(fromBlock, handler, "ResetRRRoleResult")
}

func (_Service *ServiceFilterer) WatchAllResetRRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.contract.WatchLogs(fromBlock, handler, "ResetRRRoleResult")
}

// ParseResetRRRoleResult is a log parse operation binding the contract event 0x2f43c47684e93292df983b142ac84564c8171798543d1e4c058a00b8cad71c1c.
//
// Solidity: event ResetRRRoleResult(bool )
func (_Service *ServiceFilterer) ParseResetRRRoleResult(log types.Log) (*ServiceResetRRRoleResult, error) {
	event := new(ServiceResetRRRoleResult)
	if err := _Service.contract.UnpackLog(event, "ResetRRRoleResult", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchResetRRRoleResult is a free log subscription operation binding the contract event 0x2f43c47684e93292df983b142ac84564c8171798543d1e4c058a00b8cad71c1c.
//
// Solidity: event ResetRRRoleResult(bool )
func (_Service *ServiceSession) WatchResetRRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.Contract.WatchResetRRRoleResult(fromBlock, handler)
}

func (_Service *ServiceSession) WatchAllResetRRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.Contract.WatchAllResetRRRoleResult(fromBlock, handler)
}

// ParseResetRRRoleResult is a log parse operation binding the contract event 0x2f43c47684e93292df983b142ac84564c8171798543d1e4c058a00b8cad71c1c.
//
// Solidity: event ResetRRRoleResult(bool )
func (_Service *ServiceSession) ParseResetRRRoleResult(log types.Log) (*ServiceResetRRRoleResult, error) {
	return _Service.Contract.ParseResetRRRoleResult(log)
}

// ServiceSetDRRoleResult represents a SetDRRoleResult event raised by the Service contract.
type ServiceSetDRRoleResult struct {
	Arg0 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// WatchSetDRRoleResult is a free log subscription operation binding the contract event 0x266b0a5b2ad2251976f3606e8b06066e6ce8379e6ee3eeda1f934241d86dd7fa.
//
// Solidity: event SetDRRoleResult(bool )
func (_Service *ServiceFilterer) WatchSetDRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.contract.WatchLogs(fromBlock, handler, "SetDRRoleResult")
}

func (_Service *ServiceFilterer) WatchAllSetDRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.contract.WatchLogs(fromBlock, handler, "SetDRRoleResult")
}

// ParseSetDRRoleResult is a log parse operation binding the contract event 0x266b0a5b2ad2251976f3606e8b06066e6ce8379e6ee3eeda1f934241d86dd7fa.
//
// Solidity: event SetDRRoleResult(bool )
func (_Service *ServiceFilterer) ParseSetDRRoleResult(log types.Log) (*ServiceSetDRRoleResult, error) {
	event := new(ServiceSetDRRoleResult)
	if err := _Service.contract.UnpackLog(event, "SetDRRoleResult", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchSetDRRoleResult is a free log subscription operation binding the contract event 0x266b0a5b2ad2251976f3606e8b06066e6ce8379e6ee3eeda1f934241d86dd7fa.
//
// Solidity: event SetDRRoleResult(bool )
func (_Service *ServiceSession) WatchSetDRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.Contract.WatchSetDRRoleResult(fromBlock, handler)
}

func (_Service *ServiceSession) WatchAllSetDRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.Contract.WatchAllSetDRRoleResult(fromBlock, handler)
}

// ParseSetDRRoleResult is a log parse operation binding the contract event 0x266b0a5b2ad2251976f3606e8b06066e6ce8379e6ee3eeda1f934241d86dd7fa.
//
// Solidity: event SetDRRoleResult(bool )
func (_Service *ServiceSession) ParseSetDRRoleResult(log types.Log) (*ServiceSetDRRoleResult, error) {
	return _Service.Contract.ParseSetDRRoleResult(log)
}

// ServiceSetPRRoleResult represents a SetPRRoleResult event raised by the Service contract.
type ServiceSetPRRoleResult struct {
	Arg0 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// WatchSetPRRoleResult is a free log subscription operation binding the contract event 0xf6a07132b4cd39d748fc7677eb7b47b35f1287ba285e018e8cf288f29b5f142e.
//
// Solidity: event SetPRRoleResult(bool )
func (_Service *ServiceFilterer) WatchSetPRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.contract.WatchLogs(fromBlock, handler, "SetPRRoleResult")
}

func (_Service *ServiceFilterer) WatchAllSetPRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.contract.WatchLogs(fromBlock, handler, "SetPRRoleResult")
}

// ParseSetPRRoleResult is a log parse operation binding the contract event 0xf6a07132b4cd39d748fc7677eb7b47b35f1287ba285e018e8cf288f29b5f142e.
//
// Solidity: event SetPRRoleResult(bool )
func (_Service *ServiceFilterer) ParseSetPRRoleResult(log types.Log) (*ServiceSetPRRoleResult, error) {
	event := new(ServiceSetPRRoleResult)
	if err := _Service.contract.UnpackLog(event, "SetPRRoleResult", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchSetPRRoleResult is a free log subscription operation binding the contract event 0xf6a07132b4cd39d748fc7677eb7b47b35f1287ba285e018e8cf288f29b5f142e.
//
// Solidity: event SetPRRoleResult(bool )
func (_Service *ServiceSession) WatchSetPRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.Contract.WatchSetPRRoleResult(fromBlock, handler)
}

func (_Service *ServiceSession) WatchAllSetPRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.Contract.WatchAllSetPRRoleResult(fromBlock, handler)
}

// ParseSetPRRoleResult is a log parse operation binding the contract event 0xf6a07132b4cd39d748fc7677eb7b47b35f1287ba285e018e8cf288f29b5f142e.
//
// Solidity: event SetPRRoleResult(bool )
func (_Service *ServiceSession) ParseSetPRRoleResult(log types.Log) (*ServiceSetPRRoleResult, error) {
	return _Service.Contract.ParseSetPRRoleResult(log)
}

// ServiceSetRRRoleResult represents a SetRRRoleResult event raised by the Service contract.
type ServiceSetRRRoleResult struct {
	Arg0 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// WatchSetRRRoleResult is a free log subscription operation binding the contract event 0x938928b1b88fe5ccc23a27f6112852db8f7d30dd5df882d4afcc570a4ee2fac5.
//
// Solidity: event SetRRRoleResult(bool )
func (_Service *ServiceFilterer) WatchSetRRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.contract.WatchLogs(fromBlock, handler, "SetRRRoleResult")
}

func (_Service *ServiceFilterer) WatchAllSetRRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.contract.WatchLogs(fromBlock, handler, "SetRRRoleResult")
}

// ParseSetRRRoleResult is a log parse operation binding the contract event 0x938928b1b88fe5ccc23a27f6112852db8f7d30dd5df882d4afcc570a4ee2fac5.
//
// Solidity: event SetRRRoleResult(bool )
func (_Service *ServiceFilterer) ParseSetRRRoleResult(log types.Log) (*ServiceSetRRRoleResult, error) {
	event := new(ServiceSetRRRoleResult)
	if err := _Service.contract.UnpackLog(event, "SetRRRoleResult", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WatchSetRRRoleResult is a free log subscription operation binding the contract event 0x938928b1b88fe5ccc23a27f6112852db8f7d30dd5df882d4afcc570a4ee2fac5.
//
// Solidity: event SetRRRoleResult(bool )
func (_Service *ServiceSession) WatchSetRRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.Contract.WatchSetRRRoleResult(fromBlock, handler)
}

func (_Service *ServiceSession) WatchAllSetRRRoleResult(fromBlock *uint64, handler func(int, []types.Log)) error {
	return _Service.Contract.WatchAllSetRRRoleResult(fromBlock, handler)
}

// ParseSetRRRoleResult is a log parse operation binding the contract event 0x938928b1b88fe5ccc23a27f6112852db8f7d30dd5df882d4afcc570a4ee2fac5.
//
// Solidity: event SetRRRoleResult(bool )
func (_Service *ServiceSession) ParseSetRRRoleResult(log types.Log) (*ServiceSetRRRoleResult, error) {
	return _Service.Contract.ParseSetRRRoleResult(log)
}
