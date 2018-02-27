import React from 'react'
import IndexRoute from 'react-router/lib/IndexRoute'
import Route from 'react-router/lib/Route'
//Layouts
import Layout from './common/containers/Layout';
import DashboardServersCharts from './dashboard/containers/DashboardServersCharts';

export default (user) => {
    console.log("user " + user);
    var core_url = process.env.CORE_URL ? process.env.CORE_URL : '/';
    return (
        <Route path='/dashboard' component={DashboardServersCharts} mode={'full'}></Route>
    )
}