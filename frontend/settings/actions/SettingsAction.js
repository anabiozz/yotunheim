import { 
    GET_SETTINGS_REQUEST, 
    GET_SETTINGS_SUCCESS, 
    GET_SETTINGS_ERROR, 
    GET_SETTINGS_RESET, 
    SET_REGION_SETTING,
    DISMISS_SETTINGS_ERROR, 
    SAVE_SETTINGS,
    SAVE_SUCCESS,
} from '../constants';
import fetch from 'isomorphic-fetch';
import config from '../../config';

const receiveSuccess = response => ({ type: GET_SETTINGS_SUCCESS, response })
const receiveFail = error => ({ type: GET_SETTINGS_ERROR, error })


export function dispatchError(error) {
    return (dispatch) => {
        dispatch(receiveFail(error));
    }
}

export function getSettings(page) {
    return (dispatch) => {
        dispatch({
            type: GET_SETTINGS_REQUEST
        })

        let domain = config.baseDomain+process.env.CORE_URL
        if (typeof window !== 'undefined') {
            domain = window.location.origin+process.env.CORE_URL
        }
        let url = page ? domain+'api/settings?page='+page : domain+'api/settings'
        return fetch(url, {
            method: 'get',
            credentials: 'include'
        })
        .then(response => {
            if( 200 == response.status ) {
                return response
            } else {
                throw new Error('Cannot load data from server. Response status ' + response.status)
            }
        })
        .then(response => response.json())
        .then(response =>  dispatch(receiveSuccess(response)))
        .catch(error => dispatch(receiveFail(error)))
    }
}

export function setSettings(id, index, state_id, setting) {
    return (dispatch) => {
        return dispatch({
            type: SET_THRESHOLDS_SETTING, id, index, state_id, setting
        })
    }
}

export function setRegSettings(inst_location, setting) {
    return (dispatch) => {
        return dispatch({
            type: SET_REGION_SETTING, inst_location, setting
        })
    }
}

export function saveSettings() {
    console.log("saveSettings Action")
    return (dispatch, getState) => {
        dispatch({
            type: SAVE_SETTINGS
        })
        let settings = getState().default.settings
        return fetch(config.baseDomain + '/api/settings', {
            method: 'put',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(settings)
        })
        .then(response => response.json())
        .then((response) => {
            console.log(response)
            if (!response.error) {
                dispatch({
                    type: SAVE_SUCCESS, response
                })
                return response
            } else {
                var error = new Error(response.message)
                console.log(response.message)
                error.response = response
                throw error
            }
        })
        .catch(error => dispatch(receiveFail(error)))
    }
}

export function reset() {
    return (dispatch) => {
        return dispatch({
            type: GET_SETTINGS_RESET
        })
    }
}

export function dismissError() {
    return (dispatch) => {
        return dispatch({
            type: DISMISS_SETTINGS_ERROR
        })
    }
}

export function resetFlags() {
    return (dispatch) => {
        dispatch({
            type: RESET_FLAGS
        })
    }
}