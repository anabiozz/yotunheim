import { combineReducers } from 'redux'
import commonCharts from '../../charts/common/reducers/commonCharts'
import netCharts from '../../charts/network/reducers/netCharts'
import appsCharts from '../../charts/apps/reducers/appsCharts'
import settings from '../../settings/reducers/settings'

export default combineReducers({
    commonCharts,
    netCharts,
    appsCharts,
    settings
})