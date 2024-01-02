// 食物轮流的

// 食品溯源id - 溯源码
// 各阶段时间戳
// 各阶段用户名
// 各解阶段用户地址
// 各接段食物质量


pragma solidity ^0.4.25;
pragma experimental ABIEncoderV2;

import "./LibString.sol";
import "./Table.sol";


contract FoodTraceDetail  {

    using LibString for string;

    event PutResult(int count);

    TableFactory tf;
    string constant TABLE_NAME = "tx_FoodTraceDetail";


    constructor() public {
        tf = TableFactory(0x1001);
        tf.createTable(TABLE_NAME, "traceNumber", "timestamp,traceName,traceAddress,traceQuality");
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
    /**
    * 插入数据，已有数据不添加
    */
   // 各阶段时间戳
// 各阶段用户名
// 各解阶段用户地址
// 各接段食物质量
    function newFoodDetail(uint256 _traceNumber, string _traceName, string memory _quality) public   returns(int) {
        int count = int(0);
        Table table = tf.openTable(TABLE_NAME);
        Entry entry = table.newEntry();
        entry.set("timestamp", uint256ToString(now));
        entry.set("traceName", _traceName);
        entry.set("traceAddress", msg.sender);
        entry.set("traceQuality", _quality);
        count = table.insert(uint256ToString(_traceNumber), entry);
        emit PutResult(count);
        return count;
    }
    function addTraceInfoByDistributor(uint256 _traceNumber, string _traceName, uint256 _quality) public view returns(int){
        int count = int(0);
        Table table = tf.openTable(TABLE_NAME);
        Entry entry = table.newEntry();
        entry.set("timestamp",uint256ToString(now));
        entry.set("traceName", _traceName);
        entry.set("traceAddress", msg.sender);
        entry.set("traceQuality",uint256ToString(_quality));
        count = table.insert(uint256ToString(_traceNumber), entry);
        return count ;
    }
    function addTraceInfoByRetailer(uint256 _traceNumber, string _traceName, uint256 _quality) public view returns(int){
        int count = int(0);
        Table table = tf.openTable(TABLE_NAME);
        Entry entry = table.newEntry();
        entry.set("timestamp",uint256ToString(now));
        entry.set("traceName", _traceName);
        entry.set("traceAddress", msg.sender);
        entry.set("traceQuality",uint256ToString(_quality));
        count = table.insert(uint256ToString(_traceNumber), entry);
        return count ;
    }
//    "timestamp,traceName,traceAddress,traceQuality");
    // string _traceNumber, string _traceName, uint8 _quality)
    
    function getTraceInfoByTraceNumber(string memory _traceNumber) public returns (string[], string[], string[], string[]) {
        Table table = tf.openTable(TABLE_NAME);
        Condition condition = table.newCondition();
        Entries entries = table.select(_traceNumber, condition);
        
        string[] memory timestamp = new string[](uint256(entries.size()));
        string[] memory traceName = new string[](uint256(entries.size()));
        string[] memory traceAddress = new string[](uint256(entries.size()));
        string[] memory traceQuality = new string[](uint256(entries.size()));
        
        for (uint256 i = 0; i < uint256(entries.size()); ++i) {
            Entry entry = entries.get(int256(i));
            timestamp[i] = entry.getString("timestamp");
            traceName[i] = entry.getString("traceName");
            traceAddress[i] = entry.getString("traceAddress");
            traceQuality[i] = entry.getString("traceQuality");
        }
        
        return (timestamp, traceName, traceAddress, traceQuality);
    }

    function _isExistDetail(Table _table, uint256 _traceNumber) internal view returns(bool) {
        Condition condition = _table.newCondition();
        return _table.select(uint256ToString(_traceNumber), condition).size() > int(0);
    }
}