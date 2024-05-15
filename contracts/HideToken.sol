// SPDX-License-Identifier: MIT
pragma solidity >=0.7.0 <0.9.0;

import "./openzeppelin/contracts/access/Ownable.sol";

interface IBalanceVerifier {
    function verifyProof(
        uint256[2] calldata _pA,
        uint256[2][2] calldata _pB,
        uint256[2] calldata _pC,
        uint256[3] calldata _pubSignals
    ) external view returns (bool);
}

interface ITransferVerifier {
    function verifyProof(
        uint256[2] calldata _pA,
        uint256[2][2] calldata _pB,
        uint256[2] calldata _pC,
        uint256[4] calldata _pubSignals
    ) external view returns (bool);
}

contract HideToken is Ownable {
    struct ZKProof {
        uint256[2] _pA;
        uint256[2][2] _pB;
        uint256[2] _pC;
    }

    enum TransferState {
        Pending,
        Completed,
        Revoked
    }

    struct TransferInfo {
        bytes valueBytes;
        uint256 value;
        uint256 expiredBlockNumber;
        TransferState state;
    }

    IBalanceVerifier balanceVerifier;
    ITransferVerifier transferVerifier;

    mapping(address => uint256) private _balances;
    mapping(address => bytes) private _pubKeys;
    mapping(address => mapping(address => TransferInfo)) private _transfers;

    string private _name;
    string private _symbol;
    uint8 private _decimals;
    uint256 private _effectiveBlock;

    constructor(address balanceVerifierAddr_, address transferVerifierAddr_,uint256 effectiveBlock)
    {
        balanceVerifier = IBalanceVerifier(balanceVerifierAddr_);
        transferVerifier = ITransferVerifier(transferVerifierAddr_);
        _effectiveBlock = effectiveBlock;
    }

    function initHideToken(
        string memory name_,
        string memory symbol_,
        uint8 decimals_
    ) external onlyOwner {
        _name = name_;
        _symbol = symbol_;
        _decimals = decimals_;
    }

    event RegisterWallet(address indexed sender);

    function registerWallet(
        ZKProof calldata proof,
        uint256[3] calldata _pubSignals,
        bytes memory publicKey
    ) public {
        address addr = _msgSender();
        require(_pubSignals[2] == uint256(uint160(addr)), "addr incoherence");
        require(_balances[addr] == 0, "address is registered");
        require(_pubSignals[1] == 0, "proof balance input must be zero");
        require(checkPublicKey(publicKey) == addr,"public key must be the sender");
        require(
            balanceVerifier.verifyProof(
                proof._pA,
                proof._pB,
                proof._pC,
                _pubSignals
            ) == true,
            "proof verify failed"
        );
        _balances[addr] = _pubSignals[0];
        _pubKeys[addr] = publicKey;

        emit RegisterWallet(addr);
    }

    function checkPublicKey(bytes memory publicKey) internal pure returns(address) {
        bytes32 pubHash = keccak256(publicKey);
        return address(bytes20(pubHash << 96));
    }
    event Transfer(address indexed from,address indexed to, uint256 value);
    // mint
    function mint(
        address addr,
        ZKProof calldata proof,
        uint256[4] calldata _pubSignals
    ) public returns (bool) {
        require(_balances[addr] != 0, "address must registed");
        require(
            _balances[addr] == _pubSignals[0],
            "proof balance input must be address balance"
        );
        require(
            transferVerifier.verifyProof(
                proof._pA,
                proof._pB,
                proof._pC,
                _pubSignals
            ) == true,
            "proof verify failed"
        );
        require(_pubSignals[3] == 1, "add must be add");
        _balances[addr] = _pubSignals[2];

        emit Transfer(address(0),addr,_pubSignals[1]);
        return true;
    }

    event RequestTransfer(address indexed from,address indexed to, uint256 value, bytes valueBytes);
    event ConfirmTransfer(address indexed from,address indexed to, uint256 value);
    event CancelTransfer(address indexed from,address indexed to, uint256 value);

    function transfer(
        address to,
        bytes memory valueBytes,
        ZKProof calldata fromProof,
        uint256[4] calldata fromPubSignals
    ) public returns(bool) {
        address from = _msgSender();

        require(_checkExpireTransfer(from,to),"last transfer has not expired");
        require(_balances[from] != 0, "address `from` must registed");
        require(_balances[to] != 0, "address `to` must registed");
        require(fromPubSignals[3] == 0, "must be sub");
        require(
            fromPubSignals[0] == _balances[from],
            "proof intput balance must be from balance"
        );

        _transfers[from][to].valueBytes = valueBytes;
        _transfers[from][to].value = fromPubSignals[1];
        _transfers[from][to].expiredBlockNumber = block.number + _effectiveBlock;
        _transfers[from][to].state = TransferState.Pending;

         require(
            transferVerifier.verifyProof(
                fromProof._pA,
                fromProof._pB,
                fromProof._pC,
                fromPubSignals
            ) == true,
            "proof verify failed"
        );

        _balances[from] = fromPubSignals[2];

        emit RequestTransfer(from, to, fromPubSignals[1], valueBytes);

        return true;
    }

    function _checkExpireTransfer(address from,address to) internal view returns(bool) {
        TransferInfo storage info = _transfers[from][to];
        return info.state != TransferState.Pending || info.expiredBlockNumber < block.number;
    }

    function confirmTransfer(
        address from, 
        ZKProof calldata toProof,
        uint256[4] calldata toPubSignals
    ) public returns(bool) {
        address to = _msgSender();
        require(_checkExpireTransfer(from,to) == false ,"the transfer has expired");
        require(_balances[from] != 0, "address `from` must registed");
        require(_balances[to] != 0, "address `to` must registed");
        require(toPubSignals[3] == 1, "add must be sub");
        require(
            toPubSignals[0] == _balances[to],
            "proof intput balance must be to balance"
        );
        require(_transfers[from][to].value == toPubSignals[1], "value must be equal");
        require(
            transferVerifier.verifyProof(
                toProof._pA,
                toProof._pB,
                toProof._pC,
                toPubSignals
            ) == true,
            "proof verify failed"
        );

        _balances[to] = toPubSignals[2];
        _transfers[from][to].state = TransferState.Completed;

        emit ConfirmTransfer(from,to,toPubSignals[1]);

        return true;
    }

    function cancelTransfer(
        address to, 
        ZKProof calldata fromProof,
        uint256[4] calldata fromPubSignals
    ) public returns(bool) {
        address from = _msgSender();
        require(_transfers[from][to].state == TransferState.Pending ,"the transfer has expired");
        require(_balances[from] != 0, "address `from` must registed");
        require(_balances[to] != 0, "address `to` must registed");
        require(fromPubSignals[3] == 1, "add must be sub");
        require(
            fromPubSignals[0] == _balances[from],
            "proof intput balance must be from balance"
        );
        require(_transfers[from][to].value == fromPubSignals[1], "value must be equal");
        require(
            transferVerifier.verifyProof(
                fromProof._pA,
                fromProof._pB,
                fromProof._pC,
                fromPubSignals
            ) == true,
            "proof verify failed"
        );

        _balances[from] = fromPubSignals[2];
        _transfers[from][to].state = TransferState.Revoked;

        emit CancelTransfer(from,to,fromPubSignals[1]);

        return true;
    }

    function _transfer(
        address from,
        address to,
        ZKProof memory fromProof,
        uint256[4] memory fromPubSignals,
        ZKProof memory toProof,
        uint256[4] memory toPubSignals
    ) internal returns (bool) {
        require(_balances[from] != 0, "address must registed");
        require(_balances[to] != 0, "address must registed");
        require(
            fromPubSignals[0] == _balances[from],
            "from proof intput balance must be address balance"
        );
        require(
            toPubSignals[0] == _balances[to],
            "to proof intput balance must be address balance"
        );
        require(fromPubSignals[3] == 0, "add must be sub");
        require(toPubSignals[3] == 1, "add must be add");
        require(fromPubSignals[1] == toPubSignals[1], "value must be equal");
        require(
            transferVerifier.verifyProof(
                fromProof._pA,
                fromProof._pB,
                fromProof._pC,
                fromPubSignals
            ) == true,
            "proof verify failed"
        );
        require(
            transferVerifier.verifyProof(
                toProof._pA,
                toProof._pB,
                toProof._pC,
                toPubSignals
            ) == true,
            "proof verify failed"
        );
        _balances[from] = fromPubSignals[2];
        _balances[to] = toPubSignals[2];
        return true;
    }

    function queryTransfer(address from,address to) public view returns(bytes memory,uint256,uint256 ) {
        
        require(_checkExpireTransfer(from,to) == false ,"the transfer has expired");

        return (_transfers[from][to].valueBytes,_transfers[from][to].value,_transfers[from][to].expiredBlockNumber);
    }

    function queryPublicKey(address to) public view returns(bytes memory) {
        return _pubKeys[to];
    }

    function balanceOf(address account) public view returns (uint256) {
        return _balances[account];
    }

    function name() public view returns (string memory) {
        return _name;
    }

    function symbol() public view returns (string memory) {
        return _symbol;
    }

    function decimals() public view returns (uint8) {
        return _decimals;
    }
}