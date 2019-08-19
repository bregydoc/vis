pragma solidity >=0.4.22 <0.6.0;

contract RGYx {
  address private creator;

  mapping (address => uint256) public balance;
  address[] private shareholders;

  string public name;
  uint256 genesisShares;
  uint256 private shares;

  bool public availableToTransfer = false;

  event ShareSold(address _to, uint256 _how);
  event AvailableToTransfer();

  constructor (string memory _name, uint256 _shares) public {
    creator = msg.sender;
    name = _name;
    genesisShares = _shares;
    shares = _shares;
  }

  function sellShares(address to, uint256 how ) public returns (uint256) {
    require(msg.sender == creator, "you are not Renvestgy");
    require(shares >= how, "insuficent founds");

    balance[to] += how;
    shares -= how;

    require(shares >= 0, "invalid sell");

    uint256 availableShares = shares/genesisShares*100;
    if (availableShares < 20) {
      availableToTransfer = true;
    }

    // ! existance check
    bool exist = false;
    // * consume a lot of gas
    for (uint i = 0; i<shareholders.length; i++) {
      if (shareholders[i] == to) {
        exist = true;
        break;
      }
    }
    if (!exist) {
      shareholders.push(to);
    }

    emit ShareSold(to, how);
    return shares;
  }

  function getAvailableShares() public view returns (uint256) {
    return shares;
  }

  function getSoldShares() public view returns (uint256) {
    uint256 total;
    
    for (uint i = 0; i<shareholders.length; i++) {
      total += balance[shareholders[i]];
    }

    return total;
  }

  function getGenesis() public view returns (uint256) {
    return genesisShares;
  }
}