import React from 'react'
import connect from 'react-redux/lib/connect/connect'
import { getAppsInfo, reset, dismissError } from '../actions/AppsInfoActions'
import bindActionCreators from 'redux/lib/bindActionCreators'
import config from '../../../config'

//Components
import App from '../components/App'

class AppsInfo extends React.Component {
  constructor(...props) {
    super(...props)
  }

  componentWillUnmount() {
    clearTimeout(this.timeout)
    this.props.reset()
  }

  componentDidMount() {
    this.props.dismissError()
    this.props.getAppsInfo()
    this.props.reset()
  }

  app(appsInfo) {
    return appsInfo.data.map((app, index) => {
      return <App name={app.Name} data={ app.Metrics } key={index}/>
    })
  }

  render() {
    console.log('RENDER <AppsCharts>')

    clearTimeout(this.timeout)
    this.timeout = setTimeout(this.updateCharts, config.timeInterval)

    const { appsInfo } = this.props

    console.log(appsInfo)

    return (
      <div className='apps_wrapper'>
          {
            appsInfo.data.length > 0 ? this.app(appsInfo) : "null"
          }
      </div>
    )
  }
}
function mapStateToProps(state) {

  return {
    appsInfo: state.default.appsInfo
  }
}
function mapDispatchToProps(dispatch) {
  return {
    getAppsInfo: bindActionCreators(getAppsInfo, dispatch),
    reset: bindActionCreators(reset, dispatch),
    dismissError: bindActionCreators(dismissError, dispatch),
  }
}
export default connect(mapStateToProps, mapDispatchToProps)(AppsInfo)