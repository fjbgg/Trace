/**
* 角色存储器
*/
pragma solidity ^0.4.25;
pragma experimental ABIEncoderV2;

import "./LibString.sol";
import "./Table.sol";

contract RoleStorage {

    using LibString for string;
    
    event AddRoleResult(int count);
    event RemoveRoleResult(int count);

    TableFactory tf;
    string constant TABLE_NAME = "tx_role";

    /**
    *  角色表
    * +----------------------+------------------------+-------------------------+
    * | Field                | Type                   | Desc                    |
    * +----------------------+------------------------+-------------------------+
    * | role                 | string                 | 角色                     |
    * | user_addr              |  string memory       | 用户标识，（ string memory） |
    * +----------------------+------------------------+-------------------------+
    */
    constructor() public {
        tf = TableFactory(0x1001);
        tf.createTable(TABLE_NAME, "role", "user_addr");
    }

    function addRole(string memory _role, string memory _user_addr) public returns(int) {
        int count = int(0);
        Table table = tf.openTable(TABLE_NAME);
        if(!_role.empty() && !_isExist(table, _role, _user_addr)){
            Entry entry = table.newEntry();
            entry.set("user_addr", _user_addr);
            count = table.insert(_role, entry);
        }
        emit AddRoleResult(count);
        return count;
    }

    function removeRole(string memory _role,  string memory _user_addr) public returns(int) {
        int count = int(0);
        Table table = tf.openTable(TABLE_NAME);
        Condition condition = table.newCondition();
        condition.EQ("user_addr", _user_addr);
        count = table.remove(_role, condition);
        emit RemoveRoleResult(count);
        return count;
    }

    function hasRole(string memory _role,   string memory _user_addr) public view returns(bool){
        Table table = tf.openTable(TABLE_NAME);
        
        return _isExist(table, _role, _user_addr);
    }

    function checkRole(string memory _role,  string memory _user_addr) public view{
        require(hasRole(_role, _user_addr), "not have permission");
    }

    function _isExist(Table _table, string memory _role, string memory _user_addr) internal view returns(bool) {
        Condition condition = _table.newCondition();
        condition.EQ("user_addr", _user_addr);
        return _table.select(_role, condition).size() > int(0);
    }    
   
}