import React from 'react'
import IndexRoute from 'react-router/lib/IndexRoute'
import Route from 'react-router/lib/Route'
//Layouts
import Layout from './shared/containers/Layout'
import CommonCharts from './charts/common/containers/CommonCharts'
import NetworkCharts from './charts/network/containers/NetworkCharts'
import AppsPage from './charts/common/containers/AppsPage'
import Settings from './settings/containers/Settings'

export default () => {
    return (
        <Route path='/' component={Layout}>
            <IndexRoute component={CommonCharts} />
            <Route path='main' component={CommonCharts} />
            <Route path='apps' component={AppsPage} />
            <Route path='network' component={NetworkCharts} />
            <Route path='settings' component={Settings}>
                <IndexRoute component={Settings} />
            </Route>
        </Route>
    )
}

