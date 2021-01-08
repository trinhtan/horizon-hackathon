const contractAddress = {
  0: {
    bridgeBank: 'one17hr4uufyp37rg63fnghwj5edr3dnsna5at2zy4'
  }
};

export const getHmyContractAddress = _chainId => {
  return contractAddress[_chainId];
};
