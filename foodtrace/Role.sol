pragma solidity ^0.4.25;

import "./RoleStorage.sol";

// 此合约无需部署，在合约内用
contract Role {

    string constant private PR_Role = "PR";	// 生产商 Producer
    string constant private DR_Role = "DR";	// 中间商 Distributor
    string constant private RR_Role = "RR";	// 零售商 Retailer 

    RoleStorage private roleStorage;

    constructor() public {
        roleStorage = new RoleStorage();
    }
// 用地址对应着权限 - 传入地址查询权限
    function onlyPRRole(address _user_addr) public view{
        require(roleStorage.hasRole(PR_Role, addressToString(_user_addr)), "not has pr permission");
    }

    function onlyDRRole(address _user_addr) public view{
        require(roleStorage.hasRole(DR_Role, addressToString(_user_addr)), "not has dr permission");
    }

    function onlyRRRole(address _user_addr) public view{

        require(roleStorage.hasRole(RR_Role, addressToString(_user_addr)), "not has rr permission");
    }
    function addressToString(address _address) internal pure returns (string memory) {
        bytes32 value = bytes32(uint256(_address));
        bytes memory alphabet = "0123456789abcdef";

        bytes memory str = new bytes(42);
        str[0] = "0";
        str[1] = "x";

        for (uint256 i = 0; i < 20; i++) {
            str[2+i*2] = alphabet[uint8(value[i + 12] >> 4)];
            str[3+i*2] = alphabet[uint8(value[i + 12] & 0x0f)];
        }
        return string(str);
    }
    

}
