/**
* 角色控制器
*/
pragma solidity ^0.4.25;
pragma experimental ABIEncoderV2;

import "./LibString.sol";
import "./RoleStorage.sol";

contract RoleController  {

    using LibString for string;

    string constant PR_Role = "PR";	// 生产
    string constant DR_Role = "DR";	// 中间
    string constant RR_Role = "RR";	// 零售

    RoleStorage private roleStorage;

    event SetPRRoleResult(bool);
    event ResetPRRoleResult(bool);
    event SetDRRoleResult(bool);
    event ResetDRRoleResult(bool);
    event SetRRRoleResult(bool);
    event ResetRRRoleResult(bool);

    constructor() public {
        roleStorage = new RoleStorage();
    }

    /** 授权生产方角色
    * _user_id 用户address
    */
    function setPRRole(string _user_id) external   returns(bool){
        int256 count = roleStorage.addRole(PR_Role, _user_id);
        emit SetPRRoleResult(count > int256(0));
        return count > int256(0);
    }

    /** 重置生产方角色
    * _user_id 用户id
    */
    function resetPRRole(string _user_id) external   returns(bool){
        int256 count = roleStorage.removeRole(PR_Role, _user_id);
        emit ResetPRRoleResult(count > int256(0));
        return count > int256(0);
    }

    /** 授权中间方角色
    * _user_id 用户id
    */
    function setDRRole(string _user_id) external   returns(bool){
        int256 count = roleStorage.addRole(DR_Role, _user_id);
        emit SetDRRoleResult(count > int256(0));
        return count > int256(0);
    }

    /** 重置中间方角色
    * _user_id 用户id
    */
    function resetDRRole(string _user_id) external   returns(bool){
        int256 count = roleStorage.removeRole(DR_Role, _user_id);
        emit ResetDRRoleResult(count > int256(0));
        return count > int256(0);
    }

    /** 授权零售方角色
    * _user_id 用户id
    */
    function setRRRole(string _user_id) external   returns(bool){
        int256 count = roleStorage.addRole(RR_Role, _user_id);
        emit SetRRRoleResult(count > int256(0));
        return count > int256(0);
    }

    /** 重置零售方角色
    * _user_id 用户id
    */
    function resetRRRole(string _user_id) external   returns(bool){
        int256 count = roleStorage.removeRole(RR_Role, _user_id);
        emit ResetRRRoleResult(count > int256(0));
        return count > int256(0);
    }

}