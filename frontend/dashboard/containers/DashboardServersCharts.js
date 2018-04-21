import React from 'react'
import connect from 'react-redux/lib/connect/connect'
import { getJson, reset, dismissError } from '../actions/DashboardActions'
import bindActionCreators from 'redux/lib/bindActionCreators'

//Components
import ChartCell from '../components/ServersCallGrid'
import Layout from '../../common/containers/Layout'

class DashboardServersCharts extends React.Component {
  constructor(...props) {
    super(...props)
    
    this.updateData = this.updateData.bind(this)
  }

  updateData() {
    this.props.dismissError()
    this.props.getJson()
  }

  componentWillUnmount() {
    clearTimeout(this.timeout)
  }

  handleAlertDismiss() {
    this.props.dismissError()
  }

  componentDidMount() {
    this.props.getJson()
  }

  render() {
    console.log('DashboardServersCharts render')

    clearTimeout(this.timeout)
    this.timeout = setTimeout(this.updateData, 4000)

    const { json } = this.props
    return <div className='row main-row'>
          <Layout mode={this.props.route.mode}>
              <div>
                { <ChartCell data={json}/>}
              </div>
          </Layout>
        </div>
  }
}
function mapStateToProps(state) {
  return {
    json: state.default.charts.json,
  }
}
function mapDispatchToProps(dispatch) {
  return {
    getJson: bindActionCreators(getJson, dispatch),
    reset: bindActionCreators(reset, dispatch),
    dismissError: bindActionCreators(dismissError, dispatch),
  }
}
export default connect(mapStateToProps, mapDispatchToProps)(DashboardServersCharts)