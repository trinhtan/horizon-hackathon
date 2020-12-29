import dai from 'assets/icons/dai.png';
import knc from 'assets/icons/knc.png';
import link from 'assets/icons/link.png';
import eth from 'assets/icons/eth.png';
import one from 'assets/icons/one.png';
import { getETHContractAddress } from 'utils/getETHContractAddress';
const ERC20 = require('contracts/IERC20.json');
const { Client } = require('@bandprotocol/bandchain.js');

const tokenInfo = {
  42: [
    {
      symbol: 'DAI',
      ethAddress: '0xff795577d9ac8bd7d90ee22b6c1703490b6512fd',
      icon: dai
    },
    {
      symbol: 'KNC',
      ethAddress: '0x3F80c39c0b96A0945f9F0E9f55d8A8891c5671A8',
      icon: knc
    },
    {
      symbol: 'LINK',
      ethAddress: '0xAD5ce863aE3E4E9394Ab43d4ba0D80f419F61789',
      icon: link
    },
    {
      symbol: 'WETH',
      ethAddress: '0xd0A1E359811322d97991E03f863a0C30C2cF029C',
      icon: eth
    }
  ]
};

export const getSymbolETH = (_chainId, _address) => {
  let symbol = '';
  let listToken = tokenInfo[_chainId] ? tokenInfo[_chainId] : [];
  listToken.forEach(token => {
    if (token.ethAddress === _address) {
      symbol = token.symbol;
    }
  });
  return symbol;
};

export const getIconETH = (_chainId, _address) => {
  let icon;
  let listToken = tokenInfo[_chainId] ? tokenInfo[_chainId] : [];
  listToken.forEach(token => {
    if (token.ethAddress === _address) {
      icon = token.icon;
    }
  });
  return icon;
};

export const convertToken = async (src, target, amount) => {
  const endpoint = 'https://api-gm-lb.bandchain.org';
  const bandchain = new Client(endpoint);
  const res = await bandchain.getReferenceData([src + '/' + target]);
  let rate;
  if (res.length > 0) {
    rate = res[0].rate;
  }
  return rate * amount;
};

export const balanceOf = async (tokenAddress, walletAddress) => {
  const web3 = new Web3(window.ethereum);
  let balance;
  if (tokenAddress === '0x0000000000000000000000000000000000000001') {
    balance = web3.eth.getBalance(walletAddress);
  } else {
    const erc20 = new web3.eth.Contract(ERC20.abi, tokenAddress);
    balance = await erc20.methods.balanceOf(walletAddress).call();
  }
  return balance;
};

export const allowance = async (tokenAddress, walletAddress, chainId) => {
  const web3 = new Web3(window.ethereum);
  let tokenAllowance = 0;
  if (tokenAddress !== '0x0000000000000000000000000000000000000001') {
    const erc20 = new web3.eth.Contract(ERC20.abi, tokenAddress);
    let contractAddress = getETHContractAddress(chainId);
    tokenAllowance = await erc20.methods
      .allowance(walletAddress, contractAddress.bridgeBank)
      .call();
  }
  return tokenAllowance;
};
