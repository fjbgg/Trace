// 存贮食物的信息

// 食品溯源id
// 食物名称
// 当前用户名
// 质量
// 状态 （0:生产 1:分销 2:出售）
pragma solidity ^0.4.25;
pragma experimental ABIEncoderV2;

import "./LibString.sol";
import "./Table.sol";
import "./Role.sol";
import "./FoodTraceDetail.sol";
contract FoodInfo  is Role {

    using LibString for string;

    event PutResult(int count);
    FoodTraceDetail foodtracedetail;
    TableFactory tf;
    string constant TABLE_NAME = "tx_foodinfo";

    /**
    * 食品信息存贮表 
    */
    constructor() public {
        tf = TableFactory(0x1001);
        tf.createTable(TABLE_NAME, "traceNumber", "name,current_username,quality,status");
        foodtracedetail = new FoodTraceDetail();
    }

    /**
    * 创建食品
    */
//    食品名称，溯源码，当前用户名，质量
    function newFood(string _name, uint256 _traceNumber, string _traceName, string memory _quality) public   returns(int) {
        onlyPRRole(msg.sender);
        int count = int(0);
        Table table = tf.openTable(TABLE_NAME);
        if(!_name.empty() && !_traceName.empty() && !_isExist(table, _traceNumber)){
            Entry entry = table.newEntry();
            entry.set("name", _name);
            entry.set("current_username", _traceName);
            entry.set("quality", _quality);
            entry.set("status", "0");
            count = table.insert(uint256ToString(_traceNumber), entry);
            foodtracedetail.newFoodDetail(_traceNumber, _traceName, _quality);
        }
        emit PutResult(count);
        return count;
    }

    /**
    * 食品分销过程中增加溯源信息的接口
    */
    function addTraceInfoByDistributor(uint256 _traceNumber, string _traceName, uint256 _quality) public  returns(bool){
        onlyDRRole(msg.sender);
        Table table = tf.openTable(TABLE_NAME);
        Condition condition = table.newCondition();
        Entries entries = table.select(uint256ToString(_traceNumber), condition);
        require(entries.size() >0 ,"error");
        Entry en =  entries.get(0);
        string memory status = uint256ToString(uint256(en.getInt("status")));
        require(equal("0",status),"status must be distributing");
        Entry entry  =  table.newEntry();
        entry.set("status", "1");
        table.update(uint256ToString(_traceNumber),entry,condition);
        int _a = foodtracedetail.addTraceInfoByDistributor(_traceNumber, _traceName, _quality);
        return _a>0 ?true:false ;
    }
     /**
    * 食品出售中增加溯源信息的接口
    */
    function addTraceInfoByRetailer(uint256 _traceNumber, string _traceName, uint256 _quality) public  returns(bool) {
        onlyRRRole(msg.sender);
        Table table = tf.openTable(TABLE_NAME);
        Condition condition = table.newCondition();
        Entries entries = table.select(uint256ToString(_traceNumber), condition);
        require(entries.size() >0 ,"error");
        Entry en =  entries.get(0);
        string memory status = uint256ToString(uint256(en.getInt("status")));
        require(equal("1",status),"status must be distributing");
        Entry entry  =  table.newEntry();
        entry.set("status", "2");
        table.update(uint256ToString(_traceNumber),entry,condition);
        int _a = foodtracedetail.addTraceInfoByRetailer(_traceNumber, _traceName, _quality);
        return _a>0 ?true:false ;
    }

    function _isExist(Table _table, uint256 _traceNumber) internal view returns(bool) {
        Condition condition = _table.newCondition();
        return _table.select(uint256ToString(_traceNumber), condition).size() > int(0);
    }

    function getTraceDetailByTraceNumber(string memory _traceNumber) public returns(string[],string[],string[],string[]){
        Table table = tf.openTable(TABLE_NAME);
        Condition condition = table.newCondition();
        string[] memory a;
        string[] memory b;
        string[] memory c;
        string[] memory d;
        (a, b, c, d) = foodtracedetail.getTraceInfoByTraceNumber(_traceNumber);
        return  (a,b,c,d);
    }

    function getTraceinfoByTraceNumber(string memory _traceNumber) public returns (string, string, string, string) {
        Table table = tf.openTable(TABLE_NAME);
        Condition condition = table.newCondition();
        Entries entries = table.select(_traceNumber, condition);
        Entry entry = entries.get(0);
        
        string memory name = entry.getString("name");
        string memory currentUsername = entry.getString("current_username");
        string memory quality = entry.getString("quality");
        string memory status = uint256ToString(uint256(entry.getInt("status")));
        
        return (name, currentUsername, quality, status);
    }

    function joinStrings(string[] memory strings) internal pure returns (string memory) {
        uint256 length = 0;
        for (uint256 ii = 0; ii < strings.length; ii++) {
            length += bytes(strings[ii]).length;
        }
        string memory result = new string(length);
        uint256 currentIndex = 0;
        for (uint256 i = 0; i < strings.length; i++) {
            // Copy each string to the result string
            bytes memory stringBytes = bytes(strings[i]);
            for (uint256 j = 0; j < stringBytes.length; j++) {
                bytes(result)[currentIndex] = stringBytes[j];
                currentIndex++;
            }
        }
        return result;
    }
    function uint256ToString(uint256 _origin) public pure returns (string memory result) {
        if (_origin == 0) {
            return "0";
        }
        //计算字符串长度
        uint256 length = 0;
        uint256 temp = _origin;
        while (temp != 0){
            length++;
            temp /= 10;
        }
        //赋值
        bytes memory resultBytes = new bytes(length);
        while (_origin != 0){
            resultBytes[length - 1] = bytes1(uint8(48 + _origin % 10));
            _origin /= 10;
            length--;
        }
        result = string(resultBytes);
    }
    function equal(string memory self, string memory other) internal pure returns(bool){
        bytes memory self_rep = bytes(self);
        bytes memory other_rep = bytes(other);
        
        if(self_rep.length != other_rep.length){
            return false;
        }
        uint selfLen = self_rep.length;
        for(uint i=0;i<selfLen;i++){
            if(self_rep[i] != other_rep[i]) return false;
        }
        return true;           
    }
}