import { Harmony } from '@harmony-js/core';
import { ChainID, ChainType, fromWei, hexToNumber, Units } from '@harmony-js/utils';
const hmy = new Harmony('https://api.s0.b.hmny.io', {
  chainType: ChainType.Harmony,
  chainId: ChainID.HmyTestnet
});
