pragma solidity ^0.4.24;

contract TableFactory {
    function openTable(string) public constant returns (Table);  
    function createTable(string,string,string) public returns(int); 
}


contract Condition {

    function EQ(string, int) public;
    function EQ(string, string) public;
}


contract Entry {
    function getInt(string) public constant returns(int);
    function getAddress(string) public constant returns(address);
    function getBytes64(string) public constant returns(byte[64]);
    function getBytes32(string) public constant returns(bytes32);
    function getString(string) public constant returns(string);

    function set(string, int) public;
    function set(string, string) public;
    function set(string, address) public;
}


contract Entries {
    function get(int) public constant returns(Entry);
    function size() public constant returns(int);
}


contract Table {
    function select(string, Condition) public constant returns(Entries);
    function insert(string, Entry) public returns(int);
    function update(string, Entry, Condition) public returns(int);
    function remove(string, Condition) public returns(int);

    function newEntry() public constant returns(Entry);
    function newCondition() public constant returns(Condition);
}