export const SET_ADDRESS = '@@eth/SET_ADDDRESS';
export const setAddress = address => async dispatch => {
  dispatch({ type: SET_ADDRESS, address });
};
