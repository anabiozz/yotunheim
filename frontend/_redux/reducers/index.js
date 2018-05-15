import { combineReducers } from 'redux'
import charts from '../../charts/common/reducers/charts'
import settings from '../../settings/reducers/settings'

export default combineReducers({
    charts,
    settings
})