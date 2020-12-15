import * as actions from './action';
const initState = {
  address: null
};

const reducer = (state = initState, action) => {
  switch (action.type) {
    case actions.SET_ADDRESS:
      return {
        ...state,
        address: action.address
      };
    default:
      return state;
  }
};

export { reducer as ethReducer };
