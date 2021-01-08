const contractAddress = {
  0: {
    bridgeBank: '0x0b7fB1bfEebAB5F14A6f28cfbDCBB79Af8cf0937'
  }
};

export const getHmyContractAddress = _chainId => {
  return contractAddress[_chainId];
};
