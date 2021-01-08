const contractAddress = {
  42: {
    bridgeBank: '0xb297f87654AC5d33Be4D27592E248D922147FE9B'
  }
};

export const getETHContractAddress = _chainId => {
  return contractAddress[_chainId];
};
