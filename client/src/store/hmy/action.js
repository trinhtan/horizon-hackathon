import Web3 from 'web3';
import { Harmony } from '@harmony-js/core';
import { ChainID, ChainType, fromWei, hexToNumber, Units } from '@harmony-js/utils';

import WalletConnectProvider from '@walletconnect/web3-provider';

const provider = new WalletConnectProvider({
  rpc: {
    0: 'https://api.s0.b.hmny.io',
    1: 'https://api.s1.b.hmny.io'
    // ...
  },
  qrcodeModalOptions: {
    mobileLinks: ['trust']
  }
});

const hmy = new Harmony('https://api.s0.b.hmny.io', {
  chainType: ChainType.Harmony,
  chainId: ChainID.HmyTestnet
});

export const SET_HMY = '@@hmy/SET_HMY';
export const setHmy = hmy => async dispatch => {
  dispatch({ type: SET_HMY, hmy });
};

export const SET_ADDRESS = '@@hmy/SET_ADDDRESS';
export const setAddress = address => async dispatch => {
  dispatch({ type: SET_ADDRESS, address });
};

export const SET_IS_AUTHORIZED = '@@hmy/SET_IS_AUTHORIZED';
export const setIsAuthorized = isAuthorized => async dispatch => {
  dispatch({ type: SET_IS_AUTHORIZED, isAuthorized });
};
