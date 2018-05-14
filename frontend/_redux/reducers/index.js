import { combineReducers } from 'redux'
import charts from '../../dashboard/reducers/charts'
import settings from '../../settings/reducers/settings'

export default combineReducers({
    charts,
    settings
})