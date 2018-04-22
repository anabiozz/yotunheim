import React from 'react'
import IndexRoute from 'react-router/lib/IndexRoute'
import Route from 'react-router/lib/Route'
//Layouts
import Layout from './common/containers/Layout'
import DashboardServersCharts from './dashboard/containers/DashboardServersCharts'

export default (user) => {
    // let core_url = process.env.CORE_URL ? process.env.CORE_URL : '/';
    return (
        <Route path='/' component={Layout}>
            <IndexRoute component={DashboardServersCharts} user={user}/>
        </Route>
    )
}

