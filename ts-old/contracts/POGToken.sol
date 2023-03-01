pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

struct TwitchViewer {
    address addr;
    string username;
    bool is_subscribed;
}

contract POGToken is Ownable, ERC20 {
    mapping(address => TwitchViewer) public viewers;

    event LogViewerAdded(address addr, string username);
    event LogViewerRemoved(address addr, string username);

    constructor(uint256 initialSupply) ERC20("PogChamp Token", "POG") {
        _mint(msg.sender, initialSupply);
    }

    function mint(address account, uint256 amount) public onlyOwner {
        _mint(account, amount);
    }

    function burn(address account, uint256 amount) public onlyOwner {
        _burn(account, amount);
    }

    function addViewer(
        address _addr,
        string memory _username,
        bool _is_subscribed
    ) public {
        bytes memory tmpUsername = bytes(_username);
        require(tmpUsername.length != 0, "Username must be non-empty");
        require(_addr != address(0x00));

        require(
            viewers[_addr].addr == address(0x00),
            "A viewer with the given address has already been added"
        );

        TwitchViewer memory viewer = TwitchViewer(
            _addr,
            _username,
            _is_subscribed
        );
        viewers[viewer.addr] = viewer;

        emit LogViewerAdded(viewer.addr, viewer.username);
    }

    function removeViewer(address _addr) public {
        require(viewers[_addr].addr == _addr, "Viewer not found");

        TwitchViewer memory viewer = viewers[_addr];
        delete viewers[_addr];

        emit LogViewerRemoved(viewer.addr, viewer.username);
    }
}
