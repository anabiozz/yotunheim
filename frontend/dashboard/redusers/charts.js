import {
    GET_SERVERS_REQUEST,
    GET_SERVERS_SUCCESS,
    GET_SERVERS_ERROR,
    GET_SERVERS_RESET,
    DISMISS_SERVERS_ERROR,
  } from '../constants';
  
  const initialState = {
    stats:[],
    errors: null,
    fetching: false,
    json: []
  }
  
  export default function charts(state = initialState, action) {
    switch (action.type) {
      case GET_SERVERS_REQUEST:
        return { ...state, fetching: true }
      case GET_SERVERS_SUCCESS:
        return { ...state, json: action.response, fetching: false, errors: null }
      case GET_SERVERS_ERROR:
        return { ...state, errors: action.error, fetching: false }
      case GET_SERVERS_RESET:
        return { ...state, json: [] }
      case DISMISS_SERVERS_ERROR:
        return { ...state, errors: null }
      default:
        return state;
    }
  }
  