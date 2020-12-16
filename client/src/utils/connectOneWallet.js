import store from 'store';
import { setAddress, setHmy, setIsAuthorized } from 'store/hmy/action.js';

export const connectOneWallet = async () => {
  let isOneWallet = window.onewallet && window.onewallet.isOneWallet;
  if (isOneWallet) {
    let onewallet = window.onewallet;
    const getAccount = await onewallet.getAccount();
    let address = getAccount.address;
    store.dispatch(setAddress(address));
    store.dispatch(setHmy(onewallet));
    store.dispatch(setIsAuthorized(true));
    localStorage.setItem(
      'harmony_hmy_session',
      JSON.stringify({
        address: address
      })
    );
    console.log('Wallet harmony connected');
  } else {
    alert('Please connect Onewallet extension!');
  }
};

export const signOut = async () => {
  const { isAuthorized, hmy } = store.getState().hmy;
  if (isAuthorized) {
    return hmy
      .forgetIdentity()
      .then(() => {
        store.dispatch(setIsAuthorized(false));
        store.dispatch(setAddress(''));
        store.dispatch(setHmy(null));
        localStorage.setItem(
          'harmony_hmy_session',
          JSON.stringify({
            address: ''
          })
        );
        console.log('Wallet harmony is sign out');
        return Promise.resolve();
      })
      .catch(err => {
        console.error(err.message);
      });
  }
};
