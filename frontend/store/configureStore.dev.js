import { routerReducer } from 'react-router-redux'
import { applyMiddleware, createStore } from 'redux'
import reduxImmutableStateInvariant from 'redux-immutable-state-invariant'
import thunk from 'redux-thunk'
import rootReducer from '../reducers'

rootReducer.routing = routerReducer

export default function configureStore(initialState) {
  return createStore(
    rootReducer,
    initialState,
    applyMiddleware(thunk, reduxImmutableStateInvariant())
  )
}
