pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./EthereumBank.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-solidity/contracts/token/ERC20/ERC20Detailed.sol";
import "../ILendingPool.sol";
import "../Oracle.sol";

interface BandOracleInterface {
    struct ReferenceData {
        uint256 rate;
        uint256 lastUpdatedBase;
        uint256 lastUpdatedQuote;
    }

    function getReferenceData(string calldata _base, string calldata _quote)
        external
        view
        returns (ReferenceData memory);

    function getReferenceDataBulk(
        string[] calldata _bases,
        string[] calldata _quotes
    ) external view returns (ReferenceData[] memory);
}

contract BridgeBank is EthereumBank {
    using SafeMath for uint256;

    address public operator;
    uint256 public feeNumerator;
    uint256 public feeDenominator;
    uint256 public SAFE_NUMBER = 1e12;

    Oracle public oracle;
    HarmonyBridge public harmonyBridge;
    BandOracleInterface public bandOracleInterface;
    ILendingPool public lendingPool;

    event UpdateOracle(address _newOracle);
    event UpdateHarmonyBridge(address _newHarmonyBridge);
    event UpdateFee(uint256 _feeNumerator, uint256 _feeDenominator);
    event WithdrawETH(address _receiver, uint256 _amount);
    event WithdrawERC20(address _token, address _receiver, uint256 _amount);

    constructor(
        address _operatorAddress,
        address _oracleAddress,
        address _harmonyBridgeAddress,
        address _bandOracleAddress,
        uint256 _feeNumerator,
        uint256 _feeDenominator,
        address _lendingPool
    ) public {
        operator = _operatorAddress;
        oracle = Oracle(_oracleAddress);
        harmonyBridge = HarmonyBridge(_harmonyBridgeAddress);
        bandOracleInterface = BandOracleInterface(_bandOracleAddress);
        feeNumerator = _feeNumerator;
        feeDenominator = _feeDenominator;
        lendingPool = ILendingPool(_lendingPool);
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

    function removeToken(address _token) public onlyOperator {
        require(_token != address(0), "Invalid token");
        require(tokenPairMaps[_token] != address(0), "Token already removed!");
        removeTokenInternal(_token);
    }

    /*
     * @dev: Locks received Ethereum/ERC20 funds.
     *
     * @param _recipient: bytes representation of destination address.
     * @param _token: token address in origin chain (0x0 if ethereum)
     * @param _amount: value of deposit
     */
    function lockERC20(
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

        if (feeNumerator != 0) {
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
        lockERC20Asset(
            msg.sender,
            _harmonyReceiver,
            _ethereumToken,
            _amountEthereumToken
        );
    }

    function swapETHForONE(address _harmonyReceiver, uint256 _amountETH)
        public
        payable
        availableNonce()
    {
        uint256 fee;

        if (feeNumerator != 0) {
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

        BandOracleInterface.ReferenceData memory data = bandOracleInterface
            .getReferenceData("ETH", "ONE");
        uint256 amountONE = _amountETH.mul(data.rate);

        lockAndSwapETHForONE(
            msg.sender,
            _harmonyReceiver,
            _amountETH,
            amountONE
        );
    }

    function swapETHForToken(
        address _harmonyReceiver,
        uint256 _amountETH,
        string memory _destTokenSymbol
    ) public payable availableNonce() {
        require(symbolToToken[_destTokenSymbol] != address(0), "Invalid Token");

        uint256 fee;

        if (feeNumerator != 0) {
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

        BandOracleInterface.ReferenceData memory data = bandOracleInterface
            .getReferenceData("ETH", _destTokenSymbol);
        uint256 amountHarmonyToken = _amountETH.mul(data.rate);

        lockAndSwapETHForToken(
            msg.sender,
            _harmonyReceiver,
            _amountETH,
            _destTokenSymbol,
            amountHarmonyToken
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

        if (feeNumerator != 0) {
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

        string memory symbol = ERC20Detailed(_amountEthereumToken).symbol();

        BandOracleInterface.ReferenceData memory data = bandOracleInterface
            .getReferenceData(symbol, _destTokenSymbol);
        uint256 amountHarmonyToken = _amountEthereumToken.mul(data.rate);

        lockAndSwapTokenForToken(
            msg.sender,
            _harmonyReceiver,
            _ethereumToken,
            _amountEthereumToken,
            _destTokenSymbol,
            amountHarmonyToken
        );
    }

    function swapTokenForONE(
        address _harmonyReceiver,
        address _ethereumToken,
        uint256 _amountEthereumToken
    ) public availableNonce() {
        require(tokenPairMaps[_ethereumToken] != address(0), "Invalid Token");

        uint256 fee;

        if (feeNumerator != 0) {
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

        string memory symbol = ERC20Detailed(_ethereumToken).symbol();

        BandOracleInterface.ReferenceData memory data = bandOracleInterface
            .getReferenceData(symbol, "ONE");
        uint256 amountONE = _amountEthereumToken.mul(data.rate);

        lockAndSwapTokenForONE(
            msg.sender,
            _harmonyReceiver,
            _ethereumToken,
            _amountEthereumToken,
            amountONE
        );
    }

    function unlock(
        address payable _receiver,
        address _token,
        uint256 _amount
    ) public onlyHarmonyBridge {
        require(tokenPairMaps[_token] != address(0), "Invalid Token");

        require(
            getLockedFunds(_token) >= _amount,
            "The Bank does not hold enough locked tokens to fulfill this request."
        );

        if (_token == ETHAddress) {
            require(
                address(this).balance >= _amount,
                "Insufficient ethereum balance for delivery."
            );
        } else {
            require(
                IERC20(_token).balanceOf(address(this)) >= _amount,
                "Insufficient ERC20 token balance for delivery."
            );
        }
        unlockFunds(_receiver, _token, _amount, lendingPool);
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

    function withdrawETH(address payable _receiver, uint256 _amountETH)
        public
        onlyOperator
    {
        bool check = checkWithdrawable(ETHAddress, _amountETH);
        require(check, "Insufficient balance");
        _receiver.transfer(_amountETH);
        emit WithdrawETH(_receiver, _amountETH);
    }

    function withdrawERC20(
        address _token,
        address _receiver,
        uint256 _amount
    ) public onlyOperator {
        bool check = checkWithdrawable(_token, _amount);
        require(check, "Insufficient balance");
        IERC20(_token).transfer(_receiver, _amount);
        emit WithdrawERC20(_token, _receiver, _amount);
    }

    function checkWithdrawable(address _token, uint256 _amount)
        internal
        view
        returns (bool)
    {
        if (_token == ETHAddress) {
            return address(this).balance - getLockedFunds(_token) >= _amount;
        } else {
            return
                IERC20(_token).balanceOf(address(this)) -
                    getLockedFunds(_token) >=
                _amount;
        }
    }
}
