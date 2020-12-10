pragma solidity ^0.5.0;

import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-solidity/contracts/token/ERC20/ERC20Detailed.sol";

/*
 *  @title: EthereumBank
 *  @dev: Ethereum bank which locks Ethereum/ERC20 token deposits, and unlocks
 *        Ethereum/ERC20 tokens once the prophecy has been successfully processed.
 */
contract EthereumBank {
    using SafeMath for uint256;

    uint256 public lockNonce;

    mapping(address => uint256) public lockedFunds;
    mapping(address => address) public tokenPairMaps;
    mapping(string => address) public symbolToToken;

    address public ETHAddress;
    address public ONEAddress;

    event LogLock(
        address _ethereumSender,
        address _harmonyReceiver,
        address _ethereumToken,
        address _harmonyToken,
        uint256 _amountEthereumToken,
        uint256 _amountHarmonyToken,
        uint256 _nonce
    );

    event LogUnlock(
        address _to,
        address _token,
        uint256 _value
    );

    /*
     * @dev: Modifier declarations
     */

    modifier availableNonce() {
        require(lockNonce + 1 > lockNonce, "No available nonces.");
        _;
    }

    /*
     * @dev: Constructor which sets the lock nonce
     */
    constructor() public {
        lockNonce = 0;
        ETHAddress = address(0x1111111111111111111111111111111111111111);
        ONEAddress = address(0x2222222222222222222222222222222222222222);
    }

    function addTokenInternal(address _ethereumToken, address _harmonyToken) internal {
        tokenPairMaps[_ethereumToken] = _harmonyToken;
        string memory symbol = ERC20Detailed(_ethereumToken).symbol();
        symbolToToken[symbol] = _ethereumToken;
    }

    function removeTokenInternal(address _ethereumToken) internal {
        tokenPairMaps[_ethereumToken] = address(0);
        string memory symbol = ERC20Detailed(_ethereumToken).symbol();
        symbolToToken[symbol] = address(0);
    }

    function isAcceptedToken(address _token)
        public
        view
        returns (bool)
    {
        return tokenPairMaps[_token] != address(0);
    }

    function getTokenMappedAddress(address _token)
        public
        view
        returns (address)
    {
        return tokenPairMaps[_token];
    }

    /*
     * @dev: Gets the amount of locked tokens by symbol.
     *
     * @param _symbol: The asset's symbol.
     */
    function getLockedFunds(address _token)
        public
        view
        returns (uint256)
    {
        return lockedFunds[_token];
    }

    /*
     * @dev: Creates a new Ethereum deposit with a unique id.
     *
     * @param _sender: The sender's ethereum address.
     * @param _recipient: The intended recipient's cosmos address.
     * @param _token: The currency type, either erc20 or ethereum.
     * @param _amount: The amount of erc20 tokens/ ethereum (in wei) to be itemized.
     */
    function lockERC20Asset(
        address _ethereumSender,
        address _harmonyReceiver,
        address _ethereumToken,
        uint256 _amountEthereumToken
    ) internal {
        lockNonce = lockNonce.add(1);
        address _harmonyToken = tokenPairMaps[_ethereumToken];
        lockedFunds[_ethereumToken] = lockedFunds[_ethereumToken].add(_amountEthereumToken);

        emit LogLock(_ethereumSender, _harmonyReceiver, _ethereumToken, _harmonyToken, _amountEthereumToken, _amountEthereumToken, lockNonce);
    }

    function lockAndSwapETHForONE(
        address _ethereumSender,
        address _harmonyReceiver,
        uint256 _amountETH,
        uint256 _amountONE
    ) internal {
        lockNonce = lockNonce.add(1);
        lockedFunds[ETHAddress] = lockedFunds[ETHAddress].add(_amountETH);

        emit LogLock(_ethereumSender, _harmonyReceiver, ETHAddress, ONEAddress, _amountETH, _amountONE, lockNonce);
    }

    function lockAndSwapETHForToken(
        address _ethereumSender,
        address _harmonyReceiver,
        uint256 _amountETH,
        string memory _destTokenSymbol,
        uint256 _amountDestToken
    ) internal {
        lockNonce = lockNonce.add(1);
        lockedFunds[ETHAddress] = lockedFunds[ETHAddress].add(_amountETH);

        address harmonyToken = tokenPairMaps[symbolToToken[_destTokenSymbol]];

        emit LogLock(_ethereumSender, _harmonyReceiver, ETHAddress, harmonyToken, _amountETH, _amountDestToken, lockNonce);
    }

    function lockAndSwapTokenForToken(
        address _ethereumSender,
        address _harmonyReceiver,
        address _ethereumToken,
        uint256 _amountEthereumToken,
        string memory _destTokenSymbol,
        uint256 _amountDestToken
    ) internal {
        lockNonce = lockNonce.add(1);
        lockedFunds[_ethereumToken] = lockedFunds[_ethereumToken].add(_amountEthereumToken);

        address harmonyToken = tokenPairMaps[symbolToToken[_destTokenSymbol]];

        emit LogLock(_ethereumSender, _harmonyReceiver, _ethereumToken, harmonyToken, _amountEthereumToken, _amountDestToken, lockNonce);
    }

    function lockAndSwapTokenForONE(
        address _ethereumSender,
        address _harmonyReceiver,
        address _ethereumToken,
        uint256 _amountEthereumToken,
        uint256 _amountONE
    ) internal {
        lockNonce = lockNonce.add(1);
        lockedFunds[_ethereumToken] = lockedFunds[_ethereumToken].add(_amountEthereumToken);

        emit LogLock(_ethereumSender, _harmonyReceiver, _ethereumToken, ONEAddress, _amountEthereumToken, _amountONE, lockNonce);
    }

    /*
     * @dev: Unlocks funds held on contract and sends them to the
     *       intended recipient
     *
     * @param _recipient: recipient's Ethereum address
     * @param _token: token contract address
     * @param _amount: wei amount or ERC20 token count
     */
    function unlockFunds(
        address payable _receiver,
        address _token,
        uint256 _amount
    ) internal {
        // Decrement locked funds mapping by the amount of tokens to be unlocked
        lockedFunds[_token] = lockedFunds[_token].sub(_amount);

        // Transfer funds to intended recipient
        if (_token == ETHAddress) {
            _receiver.transfer(_amount);
        } else {
            require(
                IERC20(_token).transfer(_receiver, _amount),
                "Token transfer failed"
            );
        }

        emit LogUnlock(_receiver, _token, _amount);
    }
}
