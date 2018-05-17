import React from 'react'
import IndexRoute from 'react-router/lib/IndexRoute'
import Route from 'react-router/lib/Route'
//Layouts
import Layout from './shared/containers/Layout'
import CommonInfo from './info/common/containers/CommonInfo'
import NetworkCharts from './info/network/containers/NetworkCharts'
import AppsInfo from './info/apps/containers/AppsInfo'
import Settings from './settings/containers/Settings'

export default () => {
    return (
        <Route path='server' component={Layout}>
            <IndexRoute component={CommonInfo} />
            <Route path='common' component={CommonInfo} />
            <Route path='apps' component={AppsInfo}/>
            <Route path='apps/:app' component={NetworkCharts}/>
            <Route path='network' component={NetworkCharts} />
            <Route path='settings' component={Settings}>
                <IndexRoute component={Settings} />
            </Route>
        </Route>
    )
}

