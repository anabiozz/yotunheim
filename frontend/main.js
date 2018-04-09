import React from 'react';
import ReactDOM from 'react-dom';
import { createStore, combineReducers, applyMiddleware } from 'redux'
import Provider from 'react-redux/lib/components/Provider';
import browserHistory from 'react-router/lib/browserHistory';
import Router from 'react-router/lib/Router';
import thunkMiddleware from 'redux-thunk';
import routes from './routes';
import * as reducers from './flux/reduser';
import { syncHistoryWithStore, routerReducer } from 'react-router-redux';
import DashboardServersCharts from './dashboard/containers/DashboardServersCharts'

reducers.routing = routerReducer;

const store = createStore(combineReducers(reducers), applyMiddleware(thunkMiddleware));
const history = syncHistoryWithStore(browserHistory, store);

function run () {

    if (module.hot) {
        ReactDOM.render(
            <Provider store={store}>
                <Router routes={routes('alex')} history={history}/>
            </Provider>
            , document.getElementById('root'));
    }

    ReactDOM.render(
            <Provider store={store}>
                <Router routes={routes('alex')} history={history}/>
            </Provider>
    , document.getElementById('root'));
}

function init () {
    run();
    store.subscribe(run);
}
  
init();