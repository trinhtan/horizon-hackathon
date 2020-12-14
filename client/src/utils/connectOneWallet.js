export const connectOneWallet = async isSender => {
  let isOneWallet = window.onewallet && window.onewallet.isOneWallet;
  if (isOneWallet) {
    let onewallet = window.onewallet;
    const getAccount = await onewallet.getAccount();
    let address = getAccount.address;
    console.log('harmony address', address);
  } else {
    alert('Please connect Onewallet extension!');
  }
};
