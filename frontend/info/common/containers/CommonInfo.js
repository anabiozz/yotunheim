import React from 'react'
import connect from 'react-redux/lib/connect/connect'
import { getCommonCharts, reset, dismissError } from '../actions/CommonChartsActions'
import bindActionCreators from 'redux/lib/bindActionCreators'
import config from '../../../config'

//Components
import Charts from '../components/Charts'

class CommonCharts extends React.Component {
  constructor(...props) {
    super(...props)

    this.updateCharts = this.updateCharts.bind(this)
  }

  updateCharts() {
    this.props.dismissError()
    this.props.getCommonCharts()
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
    
    this.props.getCommonCharts()
  }

  render() {
    console.log('RENDER <CommonCharts>')

    clearTimeout(this.timeout)
    this.timeout = setTimeout(this.updateCharts, config.timeInterval)

    const { commonCharts } = this.props

    console.log(commonCharts)

    return (
      <div className='main_monitoring'>
          <Charts data={ commonCharts.data }/>
      </div>
      
    )
  }
}
function mapStateToProps(state) {
  return {
    commonCharts: state.default.commonCharts,
    settings: state.default.settings
  }
}
function mapDispatchToProps(dispatch) {
  return {
    getCommonCharts: bindActionCreators(getCommonCharts, dispatch),
    reset: bindActionCreators(reset, dispatch),
    dismissError: bindActionCreators(dismissError, dispatch),
  }
}
export default connect(mapStateToProps, mapDispatchToProps)(CommonCharts)