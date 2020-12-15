import { Harmony } from '@harmony-js/core';
import { ChainID, ChainType, fromWei, hexToNumber, Units } from '@harmony-js/utils';
const hmy = new Harmony('https://api.s0.b.hmny.io', {
  chainType: ChainType.Harmony,
  chainId: ChainID.HmyTestnet
});

export const SET_ADDRESS = '@@hmy/SET_ADDDRESS';
export const setAddress = address => async dispatch => {
  dispatch({ type: SET_ADDRESS, address });
};
