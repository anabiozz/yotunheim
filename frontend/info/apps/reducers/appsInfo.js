import {
    GET_SERVERS_REQUEST_APP,
    GET_SERVERS_SUCCESS_APP,
    GET_SERVERS_ERROR_APP,
    GET_SERVERS_RESET_APP,
    DISMISS_SERVERS_ERROR_APP,
  } from '../constants'
  
  const initialState = {
    data: [],
    errors: null,
    fetching: false
  }
  
export default function appsInfo(state = initialState, action) {
  switch (action.type) {
    case GET_SERVERS_REQUEST_APP:
      return { ...state, fetching: true }
    case GET_SERVERS_SUCCESS_APP:
      return Object.assign({}, state, { data: action.response }, {fetching: false})
    case GET_SERVERS_ERROR_APP:
      return { ...state, errors: action.error, fetching: false }
    case GET_SERVERS_RESET_APP:
      return { ...state, data: initialState }
    case DISMISS_SERVERS_ERROR_APP:
      return { ...state, errors: null }
    default:
      return state
  }
}
  