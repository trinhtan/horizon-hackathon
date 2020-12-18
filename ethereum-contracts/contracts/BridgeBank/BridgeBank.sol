pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-solidity/contracts/token/ERC20/ERC20Detailed.sol";
import "../interfaces/aave-protocol-v2/ILendingPool.sol";
import "../interfaces/aave-protocol-v2/IWETHGateway.sol";
import "../interfaces/band-oracle/BandOracleInterface.sol";
import "../Oracle.sol";

contract BridgeBank {
    using SafeMath for uint256;

    address public operator;
    uint256 public feeNumerator;
    uint256 public feeDenominator;
    uint256 public SAFE_NUMBER = 1e12;
    address public WETH;
    uint256 public lockNonce;
    address public ETHAddress = address(0x1111111111111111111111111111111111111111);
    address public ONEAddress = address(0x2222222222222222222222222222222222222222);

    mapping(address => uint256) public lockedFunds;
    mapping(address => address) public tokenPairMaps;
    mapping(string => address) public symbolToToken;

    Oracle public oracle;
    HarmonyBridge public harmonyBridge;
    BandOracleInterface public bandOracleInterface;
    ILendingPool public lendingPool;
    IWETHGateway public wethGateway;

    event UpdateOracle(address _newOracle);
    event UpdateHarmonyBridge(address _newHarmonyBridge);
    event UpdateFee(uint256 _feeNumerator, uint256 _feeDenominator);
    event WithdrawETH(address _receiver, uint256 _amount);
    event WithdrawERC20(address _token, address _receiver, uint256 _amount);
    event LogLock(
        address _ethereumSender,
        address _harmonyReceiver,
        address _ethereumToken,
        address _harmonyToken,
        uint256 _amountEthereumToken,
        uint256 _amountHarmonyToken,
        uint256 _nonce
    );
    event LogUnlock(address _to, address _token, uint256 _value);

    constructor(
        address _operatorAddress,
        address _oracleAddress,
        address _harmonyBridgeAddress,
        address _bandOracleAddress,
        address _lendingPool,
        address _wethGateway,
        address _weth
    ) public {
        operator = _operatorAddress;
        oracle = Oracle(_oracleAddress);
        harmonyBridge = HarmonyBridge(_harmonyBridgeAddress);
        bandOracleInterface = BandOracleInterface(_bandOracleAddress);
        lendingPool = ILendingPool(_lendingPool);
        wethGateway = IWETHGateway(_wethGateway);
        WETH = _weth;
        lockNonce = 0;
    }

    modifier availableNonce() {
        require(lockNonce + 1 > lockNonce, "No available nonces.");
        _;
    }

    modifier onlyOperator() {
        require(msg.sender == operator, "Must be BridgeBank operator.");
        _;
    }

    modifier onlyOracle() {
        require(
            msg.sender == address(oracle),
            "Access restricted to the oracle"
        );
        _;
    }

    modifier onlyHarmonyBridge() {
        require(
            msg.sender == address(harmonyBridge),
            "Access restricted to the harmony bridge"
        );
        _;
    }

    function() external payable onlyOperator {}

    function addToken(address _token, address _hmyToken) public onlyOperator {
        require(
            _token != address(0) && _hmyToken != address(0),
            "Invalid token"
        );
        require(tokenPairMaps[_token] == address(0), "Token already added!");
        addTokenInternal(_token, _hmyToken);
    }

    function addTokenInternal(address _ethereumToken, address _harmonyToken) internal {
        tokenPairMaps[_ethereumToken] = _harmonyToken;
        string memory symbol = ERC20Detailed(_ethereumToken).symbol();
        symbolToToken[symbol] = _ethereumToken;
    }

    function removeToken(address _token) public onlyOperator {
        require(_token != address(0), "Invalid token");
        require(tokenPairMaps[_token] != address(0), "Token already removed!");
        removeTokenInternal(_token);
    }

    function removeTokenInternal(address _ethereumToken) internal {
        tokenPairMaps[_ethereumToken] = address(0);
        string memory symbol = ERC20Detailed(_ethereumToken).symbol();
        symbolToToken[symbol] = address(0);
    }

    function isAcceptedToken(address _token) public view returns (bool) {
        return tokenPairMaps[_token] != address(0);
    }

    function getTokenMappedAddress(address _token)
        public
        view
        returns (address)
    {
        return tokenPairMaps[_token];
    }

    function getLockedFunds(address _token) public view returns (uint256) {
        return lockedFunds[_token];
    }

    /*
     * @dev: Locks received Ethereum/ERC20 funds.
     *
     * @param _recipient: bytes representation of destination address.
     * @param _token: token address in origin chain (0x0 if ethereum)
     * @param _amount: value of deposit
     */
    function swapToken_1_1(
        address _harmonyReceiver,
        address _ethereumToken,
        uint256 _amountEthereumToken
    ) public availableNonce() {
        require(
            _ethereumToken != ETHAddress &&
                tokenPairMaps[_ethereumToken] != address(0),
            "Invalid token"
        );

        uint256 fee;

        if (feeNumerator != 0 && feeDenominator != 0) {
            fee = _amountEthereumToken
                .mul(feeNumerator)
                .mul(SAFE_NUMBER)
                .div(feeDenominator)
                .div(SAFE_NUMBER);
        }

        require(
            IERC20(_ethereumToken).transferFrom(
                msg.sender,
                address(this),
                _amountEthereumToken + fee
            ),
            "Contract token allowances insufficient to complete this lock request"
        );
        IERC20(_ethereumToken).approve(
            address(lendingPool),
            _amountEthereumToken + fee
        );
        lendingPool.deposit(
            _ethereumToken,
            _amountEthereumToken + fee,
            address(this),
            0
        );

        lockNonce = lockNonce.add(1);
        lockedFunds[_ethereumToken] = lockedFunds[_ethereumToken].add(
            _amountEthereumToken
        );

        address _harmonyToken = tokenPairMaps[_ethereumToken];

        emit LogLock(
            msg.sender,
            _harmonyReceiver,
            _ethereumToken,
            _harmonyToken,
            _amountEthereumToken,
            _amountEthereumToken,
            lockNonce
        );
    }

    function swapETHForONE(address _harmonyReceiver, uint256 _amountETH)
        public
        payable
        availableNonce()
    {
        uint256 fee;

        if (feeNumerator != 0 && feeDenominator != 0) {
            fee = _amountETH
                .mul(feeNumerator)
                .mul(SAFE_NUMBER)
                .div(feeDenominator)
                .div(SAFE_NUMBER);
        }

        require(
            msg.value == _amountETH + fee,
            "The transactions value must be equal the specified amount (in wei)"
        );

        (bool success,) = address(wethGateway).call.value(msg.value)(
            abi.encodeWithSignature(
                "depositETH(address,uint16)",
                address(this),
                0
            )
        );

        require(success, "Aave LendingPool: deposit ETH failed!");

        BandOracleInterface.ReferenceData memory data = bandOracleInterface
            .getReferenceData("ETH", "ONE");

        uint256 amountONE = _amountETH.mul(data.rate);

        lockNonce = lockNonce.add(1);
        lockedFunds[ETHAddress] = lockedFunds[ETHAddress].add(_amountETH);

        emit LogLock(
            msg.sender,
            _harmonyReceiver,
            ETHAddress,
            ONEAddress,
            _amountETH,
            amountONE,
            lockNonce
        );
    }

    function swapETHForToken(
        address _harmonyReceiver,
        uint256 _amountETH,
        string memory _destTokenSymbol
    ) public payable availableNonce() {
        require(symbolToToken[_destTokenSymbol] != address(0), "Invalid Token");

        uint256 fee;

        if (feeNumerator != 0 && feeDenominator != 0) {
            fee = _amountETH
                .mul(feeNumerator)
                .mul(SAFE_NUMBER)
                .div(feeDenominator)
                .div(SAFE_NUMBER);
        }

        require(
            msg.value == _amountETH + fee,
            "The transactions value must be equal the specified amount (in wei)"
        );

        (bool success,) = address(wethGateway).call.value(msg.value)(
            abi.encodeWithSignature(
                "depositETH(address,uint16)",
                address(this),
                0
            )
        );

        require(success, "Aave LendingPool: deposit ETH failed!");

        BandOracleInterface.ReferenceData memory data = bandOracleInterface
            .getReferenceData("ETH", _destTokenSymbol);
        uint256 amountHarmonyToken = _amountETH.mul(data.rate);

        lockNonce = lockNonce.add(1);
        lockedFunds[ETHAddress] = lockedFunds[ETHAddress].add(_amountETH);

        address harmonyToken = tokenPairMaps[symbolToToken[_destTokenSymbol]];

        emit LogLock(
            msg.sender,
            _harmonyReceiver,
            ETHAddress,
            harmonyToken,
            _amountETH,
            amountHarmonyToken,
            lockNonce
        );
    }

    function swapTokenForToken(
        address _harmonyReceiver,
        address _ethereumToken,
        uint256 _amountEthereumToken,
        string memory _destTokenSymbol
    ) public availableNonce() {
        require(
            symbolToToken[_destTokenSymbol] != address(0) &&
                tokenPairMaps[_ethereumToken] != address(0),
            "Invalid Token"
        );

        uint256 fee;

        if (feeNumerator != 0 && feeDenominator != 0) {
            fee = _amountEthereumToken
                .mul(feeNumerator)
                .mul(SAFE_NUMBER)
                .div(feeDenominator)
                .div(SAFE_NUMBER);
        }

        require(
            IERC20(_ethereumToken).transferFrom(
                msg.sender,
                address(this),
                _amountEthereumToken + fee
            ),
            "Contract token allowances insufficient to complete this lock request"
        );
        IERC20(_ethereumToken).approve(
            address(lendingPool),
            _amountEthereumToken + fee
        );
        lendingPool.deposit(
            _ethereumToken,
            _amountEthereumToken + fee,
            address(this),
            0
        );

        string memory symbol = ERC20Detailed(_amountEthereumToken).symbol();

        BandOracleInterface.ReferenceData memory data = bandOracleInterface
            .getReferenceData(symbol, _destTokenSymbol);
        uint256 amountHarmonyToken = _amountEthereumToken.mul(data.rate);

        lockNonce = lockNonce.add(1);
        lockedFunds[_ethereumToken] = lockedFunds[_ethereumToken].add(
            _amountEthereumToken
        );

        address harmonyToken = tokenPairMaps[symbolToToken[_destTokenSymbol]];

        emit LogLock(
            msg.sender,
            _harmonyReceiver,
            _ethereumToken,
            harmonyToken,
            _amountEthereumToken,
            amountHarmonyToken,
            lockNonce
        );
    }

    function swapTokenForONE(
        address _harmonyReceiver,
        address _ethereumToken,
        uint256 _amountEthereumToken
    ) public availableNonce() {
        require(tokenPairMaps[_ethereumToken] != address(0), "Invalid Token");

        uint256 fee;

        if (feeNumerator != 0 && feeDenominator != 0) {
            fee = _amountEthereumToken
                .mul(feeNumerator)
                .mul(SAFE_NUMBER)
                .div(feeDenominator)
                .div(SAFE_NUMBER);
        }

        require(
            IERC20(_ethereumToken).transferFrom(
                msg.sender,
                address(this),
                _amountEthereumToken + fee
            ),
            "Contract token allowances insufficient to complete this lock request"
        );
        IERC20(_ethereumToken).approve(
            address(lendingPool),
            _amountEthereumToken + fee
        );
        lendingPool.deposit(
            _ethereumToken,
            _amountEthereumToken + fee,
            address(this),
            0
        );

        string memory symbol = ERC20Detailed(_ethereumToken).symbol();

        BandOracleInterface.ReferenceData memory data = bandOracleInterface
            .getReferenceData(symbol, "ONE");
        uint256 amountONE = _amountEthereumToken.mul(data.rate);

        lockNonce = lockNonce.add(1);
        lockedFunds[_ethereumToken] = lockedFunds[_ethereumToken].add(
            _amountEthereumToken
        );

        emit LogLock(
            msg.sender,
            _harmonyReceiver,
            _ethereumToken,
            ONEAddress,
            _amountEthereumToken,
            amountONE,
            lockNonce
        );
    }

    function unlockERC20(
        address payable _receiver,
        address _token,
        uint256 _amount
    ) public onlyHarmonyBridge {
        require(tokenPairMaps[_token] != address(0), "Invalid Token");

        DataTypes.ReserveData memory reserve = lendingPool.getReserveData(
            _token
        );
        address aToken = reserve.aTokenAddress;
        uint256 totalAmount = IERC20(aToken).balanceOf(address(this));
        require(_amount <= totalAmount, "Not enough aToken fund");

        lendingPool.withdraw(_token, _amount, _receiver);
        lockedFunds[_token] = lockedFunds[_token].sub(
            _amount
        );

        emit LogUnlock(_receiver, _token, _amount);
    }

    function unlockETH(
        address payable _receiver,
        uint256 _amount
    ) public onlyHarmonyBridge {
        DataTypes.ReserveData memory reserve = lendingPool.getReserveData(
            WETH
        );
        address aToken = reserve.aTokenAddress;
        uint256 totalAmount = IERC20(aToken).balanceOf(address(this));
        require(_amount <= totalAmount, "Not enough aWETH fund");
        IERC20(aToken).approve(address(wethGateway), _amount);
        wethGateway.withdrawETH(_amount, _receiver);
        lockedFunds[ETHAddress] = lockedFunds[ETHAddress].sub(
            _amount
        );

        emit LogUnlock(_receiver, ETHAddress, _amount);
    }

    function updateOracle(address _oracleAddress) public onlyOperator {
        oracle = Oracle(_oracleAddress);
        emit UpdateOracle(_oracleAddress);
    }

    function updateHmyBridge(address _harmonyBridge) public onlyOperator {
        harmonyBridge = HarmonyBridge(_harmonyBridge);
        emit UpdateHarmonyBridge(_harmonyBridge);
    }

    function updateFee(uint256 _feeNumerator, uint256 _feeDenominator)
        public
        onlyOperator
    {
        feeNumerator = _feeNumerator;
        feeDenominator = _feeDenominator;
        emit UpdateFee(_feeNumerator, _feeDenominator);
    }
}
