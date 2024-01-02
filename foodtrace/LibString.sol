pragma solidity ^0.4.25;
library LibString{
    function empty(string memory src) internal pure returns(bool){
        bytes memory src_rep = bytes(src);
        if(src_rep.length == 0) return true;

        for(uint i=0;i<src_rep.length;i++){
            byte b = src_rep[i];
            if(b != 0x20 && b != byte(0x09) && b!=byte(0x0A) && b!=byte(0x0D)) return false;
        }
        return true;
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
     function concat(string memory self, string memory str) internal returns (string memory  _ret)  {
        _ret = new string(bytes(self).length + bytes(str).length);

        uint selfptr;
        uint strptr;
        uint retptr;
        assembly {
            selfptr := add(self, 0x20)
            strptr := add(str, 0x20)
            retptr := add(_ret, 0x20)
        }
        
        memcpy(retptr, selfptr, bytes(self).length);
        memcpy(retptr+bytes(self).length, strptr, bytes(str).length);
    }
    function memcpy(uint dest, uint src, uint len) private {
        // Copy word-length chunks while possible
        for(; len >= 32; len -= 32) {
            assembly {
                mstore(dest, mload(src))
            }
            dest += 32;
            src += 32;
        }

        // Copy remaining bytes
        uint mask = 256 ** (32 - len) - 1;
        assembly {
            let srcpart := and(mload(src), not(mask))
            let destpart := and(mload(dest), mask)
            mstore(dest, or(destpart, srcpart))
        }
    }

}