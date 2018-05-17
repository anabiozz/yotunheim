import React from 'react'
import connect from 'react-redux/lib/connect/connect'
import { getAppsCharts, reset, dismissError } from '../actions/AppsChartsActions'
import bindActionCreators from 'redux/lib/bindActionCreators'
import config from '../../../config'

//Components
import Charts from '../components/Charts'

class AppsCharts extends React.Component {
  constructor(...props) {
    super(...props)

    this.updateCharts = this.updateCharts.bind(this)
  }

  updateCharts() {
    this.props.dismissError()
    this.props.getAppsCharts()
  }

  componentWillUnmount() {
    clearTimeout(this.timeout)
  }

  componentDidMount() {
    this.props.getAppsCharts()
  }

  render() {
    console.log('RENDER <AppsCharts>')

    clearTimeout(this.timeout)
    this.timeout = setTimeout(this.updateCharts, config.timeInterval)

    const { appsCharts } = this.props

    console.log(appsCharts)

    return (
      <div className='main_monitoring'>
          <Charts data={ appsCharts.data }/>
      </div>
      
    )
  }
}
function mapStateToProps(state) {

  return {
    appsCharts: state.default.appsCharts
  }
}
function mapDispatchToProps(dispatch) {
  return {
    getAppsCharts: bindActionCreators(getAppsCharts, dispatch),
    reset: bindActionCreators(reset, dispatch),
    dismissError: bindActionCreators(dismissError, dispatch),
  }
}
export default connect(mapStateToProps, mapDispatchToProps)(AppsCharts)