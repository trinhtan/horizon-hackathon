import Web3 from 'web3';
import BridgeBank from 'contracts/BridgeBank.json';
import { getETHContractAddress } from 'utils/getETHContractAddress';
const ERC20 = require('contracts/IERC20.json');
const web3 = new Web3(window.ethereum);

export const approve = async (walletAddress, tokenAddress, chainId) => {
  const contractEthAddress = getETHContractAddress(chainId);
  const erc20 = new web3.eth.Contract(ERC20.abi, tokenAddress);
  await erc20.methods.approve(contractEthAddress.bridgeBank, 10e20).send({ from: walletAddress });
};

export const swapToken_1_1 = async (walletAddress, receiver, tokenAddress, amount, chainId) => {
  const contractAddress = getETHContractAddress(chainId);
  const bridgeBank = new web3.eth.Contract(BridgeBank.abi, contractAddress.bridgeBank);
  await bridgeBank.methods
    .swapToken_1_1(receiver, tokenAddress, amount)
    .send({ from: walletAddress });
};

export const swapETHForONE = async (walletAddress, receiver, amount, chainId) => {
  const contractAddress = getETHContractAddress(chainId);
  const bridgeBank = new web3.eth.Contract(BridgeBank.abi, contractAddress.bridgeBank);
  await bridgeBank.methods
    .swapToken_1_1(receiver, amount)
    .send({ value: amount, from: walletAddress });
};

export const swapETHForWETH = async (walletAddress, receiver, amount, chainId) => {
  const contractAddress = getETHContractAddress(chainId);
  const bridgeBank = new web3.eth.Contract(BridgeBank.abi, contractAddress.bridgeBank);
  await bridgeBank.methods
    .swapETHForWETH(receiver, amount)
    .send({ value: amount, from: walletAddress });
};

export const swapETHForToken = async (walletAddress, receiver, amount, destToken, chainId) => {
  const contractAddress = getETHContractAddress(chainId);
  const bridgeBank = new web3.eth.Contract(BridgeBank.abi, contractAddress.bridgeBank);
  await bridgeBank.methods
    .swapETHForToken(receiver, amount, destToken)
    .send({ value: amount, from: walletAddress });
};
export const swapTokenForToken = async (
  walletAddress,
  receiver,
  ethToken,
  amount,
  destToken,
  chainId
) => {
  const contractAddress = getETHContractAddress(chainId);
  const bridgeBank = new web3.eth.Contract(BridgeBank.abi, contractAddress.bridgeBank);
  await bridgeBank.methods
    .swapTokenForToken(receiver, ethToken, amount, destToken)
    .send({ from: walletAddress });
};
export const swapTokenForWETH = async (walletAddress, receiver, ethToken, amount, chainId) => {
  const contractAddress = getETHContractAddress(chainId);
  const bridgeBank = new web3.eth.Contract(BridgeBank.abi, contractAddress.bridgeBank);
  await bridgeBank.methods
    .swapTokenForWETH(receiver, ethToken, amount)
    .send({ from: walletAddress });
};
export const swapTokenForONE = async (walletAddress, receiver, ethToken, amount, chainId) => {
  const contractAddress = getETHContractAddress(chainId);
  const bridgeBank = new web3.eth.Contract(BridgeBank.abi, contractAddress.bridgeBank);
  await bridgeBank.methods
    .swapTokenForONE(receiver, ethToken, amount)
    .send({ from: walletAddress });
};
