import { applyMiddleware, createStore } from 'redux'
import thunk from 'redux-thunk'
import rootReducer from '../reducers'
import { routerReducer } from 'react-router-redux'

rootReducer.routing = routerReducer

export default function configureStore(initialState) {
  return createStore(
    rootReducer,
    initialState,
    applyMiddleware(thunk)
  )
}
