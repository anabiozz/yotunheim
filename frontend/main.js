
import React from 'react'
import { render } from 'react-dom'
import { createStore, combineReducers, applyMiddleware } from 'redux'
import Provider from 'react-redux/lib/components/Provider'
import browserHistory from 'react-router/lib/browserHistory'
import Router from 'react-router/lib/Router'
import thunk from 'redux-thunk'
import reduxImmutableStateInvariant from 'redux-immutable-state-invariant'
import routes from './routes'
import * as reducers from './reducers'
import { syncHistoryWithStore, routerReducer } from 'react-router-redux'

reducers.routing = routerReducer

const store = createStore(
    combineReducers(reducers),
    applyMiddleware(thunk, reduxImmutableStateInvariant())
)

const history = syncHistoryWithStore(browserHistory, store)

render(
    <Provider store={store}>
        <Router routes={routes()} history={history}/>
    </Provider>, document.getElementById('root')
)