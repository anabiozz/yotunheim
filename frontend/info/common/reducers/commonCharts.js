import {
    GET_SERVERS_REQUEST_COMMON,
    GET_SERVERS_SUCCESS_COMMON,
    GET_SERVERS_ERROR_COMMON,
    GET_SERVERS_RESET_COMMON,
    DISMISS_SERVERS_ERROR_COMMON,
  } from '../constants'
  
  const initialState = {
    data: [],
    time: '5m',
    groupby: '30s',
    errors: null,
    fetching: false
  }
  
export default function commonCharts(state = initialState, action) {
  switch (action.type) {
    case GET_SERVERS_REQUEST_COMMON:
      return { ...state, fetching: true }
    case GET_SERVERS_SUCCESS_COMMON:
      return Object.assign({}, state, { data: action.response }, { time: '5m' }, {groupby: '30s'}, {fetching: false})
    case GET_SERVERS_ERROR_COMMON:
      return { ...state, errors: action.error, fetching: false }
    case GET_SERVERS_RESET_COMMON:
      return { ...state }
    case DISMISS_SERVERS_ERROR_COMMON:
      return { ...state, errors: null }
    default:
      return state
  }
}
  