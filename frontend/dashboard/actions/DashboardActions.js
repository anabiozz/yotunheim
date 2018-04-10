import {
    GET_SERVERS_REQUEST,
    GET_SERVERS_SUCCESS,
    GET_SERVERS_ERROR,
    GET_SERVERS_RESET,
    DISMISS_SERVERS_ERROR,
  } from '../constants';
  import fetch from 'isomorphic-fetch';
  import config from '../../config';
  
  const receiveSuccess = response => ({ type: GET_SERVERS_SUCCESS, response })
  const receiveFail = error => ({ type: GET_SERVERS_ERROR, error })
  
  export function getJson() {
    return (dispatch) => {
      dispatch({
        type: GET_SERVERS_REQUEST
      })
      return fetch(config.baseDomain + '/api/get-json', {
        method: 'get',
              // credentials: 'include'
      })
        .then(response => response.json())
        .then(response => dispatch(receiveSuccess(response)))
        .catch(error => dispatch(receiveFail(error)))
    }
  }
  
  export function reset() {
    return (dispatch) => {
      return dispatch({
        type: GET_SERVERS_RESET
      })
    }
  }
  export function dismissError() {
    return (dispatch) => {
      return dispatch({
        type: DISMISS_SERVERS_ERROR
      })
    }
  }
  