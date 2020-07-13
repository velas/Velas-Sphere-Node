//TODO general
// Make it upgradable

pragma solidity >=0.4.22 <0.6.0;

contract Tips {
    mapping(string => AuthorTip) internal keys;
    mapping(address => uint256) internal tips;

    struct AuthorTip {
        address payable author;
        uint256 tips;
    }

    function addKey(string calldata key) external {
        address payable _author = msg.sender;
        //check if it is new key
        require(keys[key].author == address(0));
        keys[key].author = _author;
    }

    function getTips(string calldata key) external view returns(uint256) {
        return keys[key].tips;
    }

    function tip(string calldata key) external payable {
        require(msg.value > 0);
         //check if key exists
        require(keys[key].author != address(0));
        keys[key].tips += msg.value;
        tips[keys[key].author] += msg.value;
    }

    function claimTips() external {
        address _author = msg.sender;
       require(tips[_author] > 0);
       msg.sender.transfer(tips[_author]);
       tips[_author] = 0;
    }
}