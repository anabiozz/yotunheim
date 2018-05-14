import { routerReducer } from 'react-router-redux'
import { applyMiddleware, createStore } from 'redux'
import { combineReducers } from 'redux'
import thunk from 'redux-thunk'
import * as rootReducer from '../reducers'

rootReducer.routing = routerReducer

export default function configureStore(initialState) {
  return createStore(
    combineReducers(rootReducer),
    initialState,
    applyMiddleware(thunk)
  )
}
