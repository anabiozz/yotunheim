import {
    GET_SERVERS_REQUEST_COMMON,
    GET_SERVERS_SUCCESS_COMMON,
    GET_SERVERS_ERROR_COMMON,
    GET_SERVERS_RESET_COMMON,
    DISMISS_SERVERS_ERROR_COMMON,
  } from '../constants'
  import fetch from 'isomorphic-fetch'
  import config from '../../../config'
  
  const receiveSuccess = response => ({ type: GET_SERVERS_SUCCESS_COMMON, response })
  const receiveFail = error => ({ type: GET_SERVERS_ERROR_COMMON, error })
  
  export function getCommonCharts() {
    return (dispatch) => {
      dispatch({
        type: GET_SERVERS_REQUEST_COMMON
      })
      return fetch(config.baseDomain + '/api/get-common-charts', {
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
        type: GET_SERVERS_RESET_COMMON
      })
    }
  }
  export function dismissError() {
    return (dispatch) => {
      return dispatch({
        type: DISMISS_SERVERS_ERROR_COMMON
      })
    }
  }
  