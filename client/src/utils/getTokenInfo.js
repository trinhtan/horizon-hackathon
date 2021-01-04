import dai from 'assets/icons/dai.png';
import knc from 'assets/icons/knc.png';
import link from 'assets/icons/link.png';
import eth from 'assets/icons/eth.png';
import one from 'assets/icons/one.png';
import Web3 from 'web3';
import { getETHContractAddress } from 'utils/getETHContractAddress';
import { Harmony } from '@harmony-js/core';
import { ChainID, ChainType } from '@harmony-js/utils';
const ERC20 = require('contracts/IERC20.json');
const { Client } = require('@bandprotocol/bandchain.js');

const hmy = new Harmony('https://api.s0.b.hmny.io', {
  chainType: ChainType.Harmony,
  chainId: ChainID.HmyTestnet
});

const tokenInfo = {
  42: [
    {
      symbol: 'DAI',
      ethAddress: '0xff795577d9ac8bd7d90ee22b6c1703490b6512fd',
      hmyAddress: '0xff795577d9ac8bd7d90ee22b6c1703490b6512fd',
      icon: dai
    },
    {
      symbol: 'KNC',
      ethAddress: '0x3F80c39c0b96A0945f9F0E9f55d8A8891c5671A8',
      hmyAddress: '0xff795577d9ac8bd7d90ee22b6c1703490b6512fd',
      icon: knc
    },
    {
      symbol: 'LINK',
      ethAddress: '0xAD5ce863aE3E4E9394Ab43d4ba0D80f419F61789',
      hmyAddress: '0xff795577d9ac8bd7d90ee22b6c1703490b6512fd',
      icon: link
    },
    {
      symbol: 'ETH',
      ethAddress: '0x0000000000000000000000000000000000000001',
      hmyAddress: '0xff795577d9ac8bd7d90ee22b6c1703490b6512fd',
      icon: eth
    },
    {
      symbol: 'ONE',
      ethAddress: '0xff795577d9ac8bd7d90ee22b6c1703490b6512fd',
      hmyAddress: '0x0000000000000000000000000000000000000001',
      icon: one
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

export const getAddressToken = (_chainId, _symbol) => {
  let address = '';
  let listToken = tokenInfo[_chainId] ? tokenInfo[_chainId] : [];
  listToken.forEach(token => {
    if (token.symbol === _symbol) {
      address = token.ethAddress;
    }
  });
  return address;
};

export const getHmyAddressToken = (_chainId, _symbol) => {
  let address = '';
  let listToken = tokenInfo[_chainId] ? tokenInfo[_chainId] : [];
  listToken.forEach(token => {
    if (token.symbol === _symbol) {
      address = token.hmyAddress;
    }
  });
  return address;
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

// chain: 0: ETH 1: Harmony
export const balanceOf = async (tokenAddress, walletAddress, roadSwap) => {
  let web3 = new Web3(window.ethereum);
  if (roadSwap) {
    web3 = hmy;
  }
  let balance;
  if (tokenAddress === '0x0000000000000000000000000000000000000001') {
    balance = web3.eth.getBalance(walletAddress);
  } else {
    const erc20 = new web3.eth.Contract(ERC20.abi, tokenAddress);
    balance = await erc20.methods.balanceOf(walletAddress).call();
  }
  return balance;
};

// chain: 0: ETH 1: Harmony
export const allowance = async (tokenAddress, walletAddress, chainId, roadSwap) => {
  let web3 = new Web3(window.ethereum);
  if (roadSwap) {
    web3 = hmy;
  }
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
