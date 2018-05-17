import { combineReducers } from 'redux'
import commonCharts from '../../info/common/reducers/commonCharts'
import netCharts from '../../info/network/reducers/netCharts'
import appsInfo from '../../info/apps/reducers/appsInfo'
import settings from '../../settings/reducers/settings'

export default combineReducers({
    commonCharts,
    netCharts,
    appsInfo,
    settings
})