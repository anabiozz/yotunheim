import React from 'react'
import connect from 'react-redux/lib/connect/connect'
import { getCharts, reset, dismissError } from '../actions/DashboardActions'
import bindActionCreators from 'redux/lib/bindActionCreators'
import config from '../../config'

//Components
import ChartCell from '../components/ServersCallGrid'

class DashboardServersCharts extends React.Component {
  constructor(...props) {
    super(...props)

    this.updateCharts = this.updateCharts.bind(this)
  }

  updateCharts() {
    const { time, groupby } = this.state
    this.props.dismissError()
    this.props.getCharts(time, groupby)
  }

  componentWillUnmount() {
    clearTimeout(this.timeout)
  }

  componentDidMount() {
    const { settings } = this.props
    this.setState({
      time: settings.time,
      groupby: settings.groupby
    })
    
    this.props.getCharts(settings.time, settings.groupby)
  }

  render() {
    console.log('RENDER <DashboardServersCharts>')

    clearTimeout(this.timeout)
    this.timeout = setTimeout(this.updateCharts, config.timeInterval)

    const { charts } = this.props

    return (
      <div className='main_monitoring'>
          <ChartCell data={ charts.data }/>
      </div>
      
    )
  }
}
function mapStateToProps(state) {
  console.log(state)
  return {
    charts: state.default.charts,
    settings: state.default.settings
  }
}
function mapDispatchToProps(dispatch) {
  return {
    getCharts: bindActionCreators(getCharts, dispatch),
    reset: bindActionCreators(reset, dispatch),
    dismissError: bindActionCreators(dismissError, dispatch),
  }
}
export default connect(mapStateToProps, mapDispatchToProps)(DashboardServersCharts)