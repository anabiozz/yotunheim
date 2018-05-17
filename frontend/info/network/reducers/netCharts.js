import {
    GET_SERVERS_REQUEST_NET,
    GET_SERVERS_SUCCESS_NET,
    GET_SERVERS_ERROR_NET,
    GET_SERVERS_RESET_NET,
    DISMISS_SERVERS_ERROR_NET,
  } from '../constants'
  
  const initialState = {
    data: [],
    errors: null,
    fetching: false
  }
  
export default function netCharts(state = initialState, action) {
  switch (action.type) {
    case GET_SERVERS_REQUEST_NET:
      return { ...state, fetching: true }
    case GET_SERVERS_SUCCESS_NET:
      return Object.assign({}, state, { data: action.response }, {fetching: false})
    case GET_SERVERS_ERROR_NET:
      return { ...state, errors: action.error, fetching: false }
    case GET_SERVERS_RESET_NET:
      return { ...state }
    case DISMISS_SERVERS_ERROR_NET:
      return { ...state, errors: null }
    default:
      return state
  }
}
  