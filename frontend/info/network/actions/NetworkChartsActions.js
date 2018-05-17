import {
    GET_SERVERS_REQUEST_NET,
    GET_SERVERS_SUCCESS_NET,
    GET_SERVERS_ERROR_NET,
    GET_SERVERS_RESET_NET,
    DISMISS_SERVERS_ERROR_NET,
  } from '../constants'
  import fetch from 'isomorphic-fetch'
  import config from '../../../config'
  
  const receiveSuccess = response => ({ type: GET_SERVERS_SUCCESS_NET, response })
  const receiveFail = error => ({ type: GET_SERVERS_ERROR_NET, error })
  
  export function getNetCharts() {
    return (dispatch) => {
      dispatch({
        type: GET_SERVERS_REQUEST_NET
      })
      return fetch(config.baseDomain + '/api/get-network-charts', {
        method: 'get',
      })
        .then(response => response.json())
        .then(response => dispatch(receiveSuccess(response)))
        .catch(error => dispatch(receiveFail(error)))
    }
  }

  export function reset() {
    return (dispatch) => {
      return dispatch({
        type: GET_SERVERS_RESET_NET
      })
    }
  }
  export function dismissError() {
    return (dispatch) => {
      return dispatch({
        type: DISMISS_SERVERS_ERROR_NET
      })
    }
  }
  