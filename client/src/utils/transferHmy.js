import { Harmony } from '@harmony-js/core';
import { ChainID, ChainType } from '@harmony-js/utils';
import BridgeBankHmy from 'contracts/BridgeBankHmy.json';
import { getHmyContractAddress } from 'utils/getHmyContractAddress';

const ERC20 = require('contracts/IERC20.json');

const options = {
  gasLimit: 6721900,
  gasPrice: 1000000000
};

const hmy = new Harmony('https://api.s0.b.hmny.io', {
  chainType: ChainType.Harmony,
  chainId: ChainID.HmyTestnet
});

export const approve = async (walletAddress, tokenAddress, chainId) => {
  const contractEthAddress = getHmyContractAddress(chainId);
  const erc20 = hmy.contracts.createContract(ERC20.abi, tokenAddress);
  await erc20.methods.approve(contractEthAddress.bridgeBank, 10e20).send({ from: walletAddress });
};

export const swapToken_1_1_hmy = async (walletAddress, receiver, tokenAddress, amount, chainId) => {
  const contractAddress = getHmyContractAddress(chainId);
  const bridgeBank = hmy.contracts.createContract(BridgeBankHmy.abi, contractAddress.bridgeBank);
  await bridgeBank.methods
    .swapToken_1_1(receiver, tokenAddress, amount)
    .send({ from: walletAddress });
};

export const swapONEForETH = async (walletAddress, receiver, amount, chainId) => {
  const contractAddress = getHmyContractAddress(chainId);
  const bridgeBank = hmy.contracts.createContract(BridgeBankHmy.abi, contractAddress.bridgeBank);
  await bridgeBank.methods
    .swapONEForETH(receiver, amount)
    .send({ value: amount, from: walletAddress });
};

export const swapONEForWONE = async (walletAddress, receiver, amount, chainId) => {
  const contractAddress = getHmyContractAddress(chainId);
  const bridgeBank = hmy.contracts.createContract(BridgeBankHmy.abi, contractAddress.bridgeBank);
  await bridgeBank.methods
    .swapONEForWONE(receiver, amount)
    .send({ value: amount, from: walletAddress });
};

export const swapONEForToken = async (walletAddress, receiver, amount, destToken, chainId) => {
  const contractAddress = getHmyContractAddress(chainId);
  const bridgeBank = hmy.contracts.createContract(BridgeBankHmy.abi, contractAddress.bridgeBank);
  await bridgeBank.methods
    .swapETHForToken(receiver, amount, destToken)
    .send({ value: amount, from: walletAddress });
};

export const swapTokenForToken_hmy = async (
  walletAddress,
  receiver,
  hmyToken,
  amount,
  destToken,
  chainId
) => {
  const contractAddress = getHmyContractAddress(chainId);
  const bridgeBank = hmy.contracts.createContract(BridgeBankHmy.abi, contractAddress.bridgeBank);
  await bridgeBank.methods
    .swapTokenForToken(receiver, hmyToken, amount, destToken)
    .send({ from: walletAddress });
};

export const swapTokenForWONE = async (walletAddress, receiver, harmonyToken, amount, chainId) => {
  const contractAddress = getHmyContractAddress(chainId);
  const bridgeBank = hmy.contracts.createContract(BridgeBankHmy.abi, contractAddress.bridgeBank);
  await bridgeBank.methods
    .swapTokenForWONE(receiver, harmonyToken, amount)
    .send({ from: walletAddress });
};

export const swapTokenForETH = async (walletAddress, receiver, hmyToken, amount, chainId) => {
  const contractAddress = getHmyContractAddress(chainId);
  const bridgeBank = hmy.contracts.createContract(BridgeBankHmy.abi, contractAddress.bridgeBank);
  await bridgeBank.methods
    .swapTokenForETH(receiver, hmyToken, amount)
    .send({ from: walletAddress });
};
