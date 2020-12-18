pragma solidity ^0.5.0;

import "../node_modules/openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./Valset.sol";
import "../node_modules/openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";
import "./BridgeBank/BridgeBank.sol";

contract HarmonyBridge {
    using SafeMath for uint256;

    address public operator;
    address public oracle;
    bool public hasOracle;
    bool public hasBridgeBank;
    uint256 public unlockClaimCount;

    Valset public valset;
    BridgeBank public bridgeBank;

    mapping(uint256 => UnlockClaim) public unlockClaims;

    enum Status {Null, Pending, Success, Failed}

    struct UnlockClaim {
        address harmonySender;
        address payable ethereumReceiver;
        address originalValidator;
        address token;
        uint256 amount;
        Status status;
    }

    event LogOracleSet(address _oracle);

    event LogBridgeBankSet(address _bridgeBank);

    event LogNewUnlockClaim(
        uint256 _unlockID,
        address _harmonySender,
        address payable _ethereumReceiver,
        address _validatorAddress,
        address _tokenAddress,
        uint256 _amount
    );

    event LogUnlockCompleted(uint256 _unlockID);

    modifier isPending(uint256 _unlockID) {
        require(isUnlockClaimActive(_unlockID), "Unlock claim is not active");
        _;
    }

    modifier onlyOperator() {
        require(msg.sender == operator, "Must be the operator.");
        _;
    }

    modifier isActive() {
        require(
            hasOracle == true && hasBridgeBank == true,
            "The Operator must set the oracle and bridge bank for bridge activation"
        );
        _;
    }

    constructor(address _operator, address _valset) public {
        unlockClaimCount = 0;
        operator = _operator;
        valset = Valset(_valset);
        hasOracle = false;
        hasBridgeBank = false;
    }

    function setOracle(address _oracle) public onlyOperator {
        require(
            !hasOracle,
            "The Oracle cannot be updated once it has been set"
        );

        hasOracle = true;
        oracle = _oracle;

        emit LogOracleSet(oracle);
    }

    function setBridgeBank(address payable _bridgeBank) public onlyOperator {
        require(
            !hasBridgeBank,
            "The Bridge Bank cannot be updated once it has been set"
        );

        hasBridgeBank = true;
        bridgeBank = BridgeBank(_bridgeBank);

        emit LogBridgeBankSet(address(bridgeBank));
    }

    function newUnlockClaim(
        address _harmonySender,
        address payable _ethereumReceiver,
        address _token,
        uint256 _amount
    ) public isActive {
        require(
            valset.isActiveValidator(msg.sender),
            "Must be an active validator"
        );
        require(bridgeBank.isAcceptedToken(_token) == true, "Invalid Token");
        require(
            bridgeBank.getLockedFunds(_token) >= _amount,
            "Not enough locked assets to complete the proposed prophecy"
        );

        // Create the new ProphecyClaim
        UnlockClaim memory unlockClaim = UnlockClaim(
            _harmonySender,
            _ethereumReceiver,
            msg.sender,
            _token,
            _amount,
            Status.Pending
        );

        // Increment count and add the new ProphecyClaim to the mapping
        unlockClaimCount = unlockClaimCount.add(1);
        unlockClaims[unlockClaimCount] = unlockClaim;

        emit LogNewUnlockClaim(
            unlockClaimCount,
            _harmonySender,
            _ethereumReceiver,
            msg.sender,
            _token,
            _amount
        );
    }

    function completeUnlockClaim(uint256 _unlockID)
        public
        isPending(_unlockID)
    {
        require(
            msg.sender == oracle,
            "Only the Oracle may complete prophecies"
        );

        unlockClaims[_unlockID].status = Status.Success;

        unlockTokens(_unlockID);

        emit LogUnlockCompleted(_unlockID);
    }

    function unlockTokens(uint256 _unlockID) internal {
        UnlockClaim memory unlockClaim = unlockClaims[_unlockID];

        if (unlockClaim.token == bridgeBank.ETHAddress()){
            bridgeBank.unlockETH(
                unlockClaim.ethereumReceiver,
                unlockClaim.amount
            );
        } else {
            bridgeBank.unlockERC20(
                unlockClaim.ethereumReceiver,
                unlockClaim.token,
                unlockClaim.amount
            );
        }
    }

    function isUnlockClaimActive(uint256 _unlockID) public view returns (bool) {
        return unlockClaims[_unlockID].status == Status.Pending;
    }

    function isUnlockClaimValidatorActive(uint256 _unlockID)
        public
        view
        returns (bool)
    {
        return
            valset.isActiveValidator(unlockClaims[_unlockID].originalValidator);
    }
}
