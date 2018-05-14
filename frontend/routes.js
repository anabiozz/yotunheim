import React from 'react'
import IndexRoute from 'react-router/lib/IndexRoute'
import Route from 'react-router/lib/Route'
//Layouts
import Layout from './common/containers/Layout'
import DashboardServersCharts from './dashboard/containers/DashboardServersCharts'
import AppsPage from './dashboard/containers/AppsPage'
import Settings from './settings/containers/Settings'

export default () => {
    return (
        <Route path='/' component={Layout}>
            <IndexRoute component={DashboardServersCharts} />
            <Route path='main' component={DashboardServersCharts} />
            <Route path='apps' component={AppsPage} />
            <Route path='settings' component={Settings}>
                <IndexRoute component={Settings} />
            </Route>
        </Route>
    )
}

