import { 
    GET_SETTINGS_REQUEST, 
    GET_SETTINGS_SUCCESS, 
    GET_SETTINGS_ERROR, 
    GET_SETTINGS_RESET, 
    DISMISS_SETTINGS_ERROR, 
    SET_REGION_SETTING,
    SAVE_SETTINGS,
    SAVE_SUCCESS,
    RESET_FLAGS 
} from '../constants';
    
const initialState = {
    time: '15m',
    groupby: '30s',
    fetching: false,
    errors: null,
    success: false
}
    
export default function settings(state = initialState, action) {
    switch (action.type) {
        case GET_SETTINGS_REQUEST:
            return { ...state, fetching: true }
        case GET_SETTINGS_SUCCESS:
            return { ...state, time: action.response.time, groupby: action.response.groupby, fetching: false, errors: null }
        case GET_SETTINGS_ERROR:
            return { ...state, errors: action.error, fetching: false }
        case GET_SETTINGS_RESET:
            return { ...state }
        case RESET_FLAGS:
            return { ...state, errors: null, success: false, fetching: false }
        case DISMISS_SETTINGS_ERROR:
            return { ...state, errors: null }
        case SET_REGION_SETTING:
            var regions = state.regions.map(region => {
                if(region.inst_location == action.inst_location) {
                    let obj = {}
                    obj[action.setting.name] = action.setting.value
                    return Object.assign({}, region, obj)
                }
                return region
            })
            return { ...state, regions }
        case SAVE_SETTINGS:
            return { ...state, errors: null, success: false, fetching: true}
        case SAVE_SUCCESS:
            return { ...state, ...action.response, fetching: false, success: true}
        default:
            return state;
    }
}
    