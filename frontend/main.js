
import React from 'react'
import { render } from 'react-dom'
import Provider from 'react-redux/lib/components/Provider'
import browserHistory from 'react-router/lib/browserHistory'
import Router from 'react-router/lib/Router'
import routes from './routes'
import * as reducers from './reducers'
import { syncHistoryWithStore, routerReducer } from 'react-router-redux'
import { loadState, SaveState } from './store/localStorage'
import configureStore from './store'

reducers.routing = routerReducer

const prersistedState = loadState()

const store = configureStore(prersistedState)
store.subscribe(() => {
    SaveState(store.getState())
})



const history = syncHistoryWithStore(browserHistory, store)

render(
    <Provider store={store}>
        <Router routes={routes()} history={history}/>
    </Provider>, document.getElementById('root')
)