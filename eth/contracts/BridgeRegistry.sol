pragma solidity ^0.5.0;

contract BridgeRegistry {
    address public harmonyBridge;
    address public bridgeBank;
    address public oracle;
    address public valset;

    event LogContractsRegistered(
        address _harmonyBridge,
        address _bridgeBank,
        address _oracle,
        address _valset
    );

    constructor(
        address _harmonyBridge,
        address _bridgeBank,
        address _oracle,
        address _valset
    ) public {
        harmonyBridge = _harmonyBridge;
        bridgeBank = _bridgeBank;
        oracle = _oracle;
        valset = _valset;

        emit LogContractsRegistered(harmonyBridge, bridgeBank, oracle, valset);
    }
}
