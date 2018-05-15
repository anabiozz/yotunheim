import React from 'react'
import connect from 'react-redux/lib/connect/connect'
import { getCharts, reset, dismissError } from '../actions/CommonChartsActions'
import bindActionCreators from 'redux/lib/bindActionCreators'
import config from '../../../config'
import { Row } from 'react-materialize'

//Components
import Charts from '../components/Charts'

class DashboardServersCharts extends React.Component {
  constructor(...props) {
    super(...props)
    
    this.updateData = this.updateData.bind(this)
  }

  updateData() {
    this.props.dismissError()
    this.props.getCharts()
  }

  componentWillUnmount() {
    clearTimeout(this.timeout)
  }

  componentDidMount() {
    this.props.getCharts()
  }

  render() {
    console.log('RENDER <DashboardServersCharts>')

    clearTimeout(this.timeout)
    this.timeout = setTimeout(this.updateData, config.timeInterval)

    const { stats } = this.props

    return (
      <Row className='main_monitoring'>
         
      </Row>
      
    )
  }
}
function mapStateToProps(state) {
  return {
    stats: state.default.charts.stats,
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