import store from 'store';
import { setAddress } from 'store/hmy/action.js';
export const connectOneWallet = async () => {
  let isOneWallet = window.onewallet && window.onewallet.isOneWallet;
  if (isOneWallet) {
    let onewallet = window.onewallet;
    const getAccount = await onewallet.getAccount();
    let address = getAccount.address;
    store.dispatch(setAddress(address));
  } else {
    alert('Please connect Onewallet extension!');
  }
};
