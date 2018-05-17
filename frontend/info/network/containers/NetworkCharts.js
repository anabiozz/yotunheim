import React from 'react'
import connect from 'react-redux/lib/connect/connect'
import { getNetCharts, reset, dismissError } from '../actions/NetworkChartsActions'
import bindActionCreators from 'redux/lib/bindActionCreators'
import config from '../../../config'

//Components
import Charts from '../components/Charts'

class NetworkCharts extends React.Component {
  constructor(...props) {
    super(...props)

    this.updateCharts = this.updateCharts.bind(this)
  }

  updateCharts() {
    this.props.dismissError()
    this.props.getNetCharts()
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
    
    this.props.getNetCharts()
  }

  render() {
    console.log('RENDER <NetworkCharts>')

    clearTimeout(this.timeout)
    this.timeout = setTimeout(this.updateCharts, config.timeInterval)

    const { netCharts } = this.props

    console.log(netCharts)

    return (
      <div className='main_monitoring'>
          <Charts data={ netCharts.data }/>
      </div>
      
    )
  }
}
function mapStateToProps(state) {
  return {
    netCharts: state.default.netCharts,
    settings: state.default.settings
  }
}
function mapDispatchToProps(dispatch) {
  return {
    getNetCharts: bindActionCreators(getNetCharts, dispatch),
    reset: bindActionCreators(reset, dispatch),
    dismissError: bindActionCreators(dismissError, dispatch),
  }
}
export default connect(mapStateToProps, mapDispatchToProps)(NetworkCharts)