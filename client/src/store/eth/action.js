import BridgeBank from 'contracts/BridgeBank.json';
import { getETHContractAddress } from 'utils/getETHContractAddress';

var contractAddress;
export const SET_ADDRESS = '@@eth/SET_ADDDRESS';
export const setAddress = address => async dispatch => {
  dispatch({ type: SET_ADDRESS, address });
};

export const SET_CHAINID = '@@eth/SET_CHAINID';
export const setChainId = chainId => dispatch => {
  dispatch({ type: SET_CHAINID, chainId });
};

export const SET_WEB3 = '@@eth/SET_WEB';
export const setWeb3 = web3 => async (dispatch, getState) => {
  dispatch({ type: SET_WEB3, web3 });

  let state = getState();
  let { chainId } = state.eth;
  contractAddress = getETHContractAddress(chainId);

  const bridgeBankInstance = new web3.eth.Contract(BridgeBank.abi, contractAddress.bridgeBank);
  dispatch(setBridgeBank(bridgeBankInstance));
};

export const SET_BRIDGE_BANK = '@@eth/SET_BRIDGE_BANK';
export const setBridgeBank = bridgeBankInstance => dispatch => {
  dispatch({ type: SET_BRIDGE_BANK, bridgeBankInstance });
};
