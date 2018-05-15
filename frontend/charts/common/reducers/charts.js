import {
    GET_SERVERS_REQUEST,
    GET_SERVERS_SUCCESS,
    GET_SERVERS_ERROR,
    GET_SERVERS_RESET,
    DISMISS_SERVERS_ERROR,
  } from '../constants'
  
  const initialState = {
    data: [],
    time: '5m',
    groupby: '30s',
    errors: null,
    fetching: false
  }
  
export default function charts(state = initialState, action) {
  switch (action.type) {
    case GET_SERVERS_REQUEST:
      return { ...state, fetching: true }
    case GET_SERVERS_SUCCESS:
      return Object.assign({}, state, { data: action.response }, { time: '5m' }, {groupby: '30s'}, {fetching: false})
    case GET_SERVERS_ERROR:
      return { ...state, errors: action.error, fetching: false }
    case GET_SERVERS_RESET:
      return { ...state }
    case DISMISS_SERVERS_ERROR:
      return { ...state, errors: null }
    default:
      return state
  }
}
  