import React from 'react';
import connect from 'react-redux/lib/connect/connect';
import { getJson, reset, dismissError } from '../actions/DashboardActions';
import config from '../../config';
import bindActionCreators from 'redux/lib/bindActionCreators';

//Components
import ChartCell from '../components/ServersCallGrid';
import Layout from '../../common/containers/Layout';

class DashboardServersCharts extends React.Component {
  constructor(...props) {
    super(...props);
    this.state = {
      stats:[],
      errors: null,
      fetching: false,
      json: []
    }
  }

  componentWillUnmount() {
  }
  handleAlertDismiss() {
  }

  componentDidMount() {
    this.props.getJson()
  }
  render() {
    console.log(' DashboardServersCharts render');
    const { json } = this.props
    return <div className='row main-row'>
        <Layout mode={this.props.route.mode}> 
          <div className='col-sm-9 col-sm-offset-3 col-md-offset-2 col-md-10 col-lg-8 main'>
            { <ChartCell chartData={json}/>}
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