pragma solidity >=0.4.22 <0.6.0;

contract RGYx {
  address private creator;

  mapping (address => uint256) public balance;
  address[] public shareholders;
  address public developer;
  string public name;
  uint256 genesisShares;
  uint256 private shares;

  uint256 public cost;

  bool public availableToTransfer = false;

  event ShareSold(address _to, uint256 _how);
  event AvailableToTransferChange(bool _available);

  constructor (string memory _name, address _developer, uint256 _shares, uint256 _cost) public {
    creator = msg.sender;
    name = _name;
    genesisShares = _shares;
    shares = _shares;
    cost = _cost;
    developer = _developer;
  }

  function sellShares(address to, uint256 how ) public returns (uint256) {
    require(msg.sender == creator, "you are not Renvestgy");
    require(shares >= how, "insuficent founds");

    balance[to] += how;
    shares -= how;

    require(shares >= 0, "invalid sell");

    uint256 available = shares/genesisShares*100;
    if (available < 20) {
      availableToTransfer = true;
      emit AvailableToTransferChange(availableToTransfer);
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

  function getTotalShareholders() public view returns (uint256) {
    return shareholders.length;
  }
}