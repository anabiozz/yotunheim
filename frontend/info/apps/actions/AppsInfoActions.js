import {
    GET_SERVERS_REQUEST_APP,
    GET_APPS_COUNT,
    GET_SERVERS_SUCCESS_APP,
    GET_SERVERS_ERROR_APP,
    GET_SERVERS_RESET_APP,
    DISMISS_SERVERS_ERROR_APP,
  } from '../constants'
  import fetch from 'isomorphic-fetch'
  import config from '../../../config'
  
  const receiveSuccess = response => ({ type: GET_SERVERS_SUCCESS_APP, response })
  const receiveFail = error => ({ type: GET_SERVERS_ERROR_APP, error })
  
  export function getAppsInfo() {
    return (dispatch) => {
      dispatch({
        type: GET_SERVERS_REQUEST_APP
      })
      return fetch(config.baseDomain + '/api/get-apps-info', {
        method: 'get',
      })
        .then(response => response.json())
        .then(response => dispatch(receiveSuccess(response)))
        .catch(error => dispatch(receiveFail(error)))
    }
  }

  export function getAppsCount() {
    return (dispatch) => {
      dispatch({
        type: GET_APPS_COUNT
      })
      return fetch(config.baseDomain + '/api/get-apps-count', {
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
        type: GET_SERVERS_RESET_APP
      })
    }
  }
  export function dismissError() {
    return (dispatch) => {
      return dispatch({
        type: DISMISS_SERVERS_ERROR_APP
      })
    }
  }
  