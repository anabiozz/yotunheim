import React from 'react'
import connect from 'react-redux/lib/connect/connect'
import { getCharts, reset, dismissError } from '../actions/DashboardActions'
import bindActionCreators from 'redux/lib/bindActionCreators'
import config from '../../config'
import Select from 'react-select';

//Components
import ChartCell from '../components/ServersCallGrid'

class DashboardServersCharts extends React.Component {
  constructor(...props) {
    super(...props)

    this.state = {
      time: '',
      groupby: ''
    }
    
    this.updateData = this.updateData.bind(this)
    this.updateCharts = this.updateCharts.bind(this)
    this.handleTime = this.handleTime.bind(this)
    this.handleGroupBy = this.handleGroupBy.bind(this)
  }

  componentWillUnmount() {
    clearTimeout(this.timeout)
  }

  componentDidMount() {
    const { time, groupby } = this.state
    this.setState({
      time: time,
      groupby: groupby
    })
    this.props.getCharts(time, groupby)
  }

  updateData(time, groupby) {
    this.setState({
      time,
      groupby
    })
    this.props.dismissError()
    this.props.getCharts(time, groupby)
  }

  updateCharts() {
    const { time, groupby } = this.state
    this.props.dismissError()
    this.props.getCharts(time, groupby)
  }

  handleTime = (time) => {
    this.updateData(`${time.label}`, this.state.groupby)
  }

  handleGroupBy = (groupby) => {
    this.updateData(this.state.time, `${groupby.label}`)
  }

  render() {
    console.log('RENDER <DashboardServersCharts>')

    clearTimeout(this.timeout)
    this.timeout = setTimeout(this.updateCharts, config.timeInterval)

    const { charts } = this.props
    return (
      <div className='main_monitoring'>
          <div className='selects'>
            <Select
                  name="time"
                  value={this.state.time}
                  onChange={this.handleTime}
                  options={[
                    { value: '5m', label: '5m' },
                    { value: '15m', label: '15m' },
                    { value: '30m', label: '30m' },
                    { value: '1h', label: '1h' },
                    { value: '3h', label: '3h' },
                    { value: '8h', label: '8h' },
                    { value: '24h', label: '24h' },
                ]}
              />
              <Select
                  name="limit"
                  value={this.state.groupby}
                  onChange={this.handleGroupBy}
                  options={[
                    { value: '30s', label: '30s' },
                    { value: '1m', label: '1m' },
                    { value: '5m', label: '5m' },
                    { value: '10m', label: '10m' },
                    { value: '30m', label: '30m' },
                    { value: '1h', label: '1h' },
                    { value: '5h', label: '5h' },
                ]}
              />
            </div>

          <ChartCell data={ charts }/>
      </div>
      
    )
  }
}
function mapStateToProps(state) {
  
  return {
    charts: state.default.charts,
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