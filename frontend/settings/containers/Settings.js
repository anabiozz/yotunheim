import React from 'react'
import connect from 'react-redux/lib/connect/connect'
import bindActionCreators from 'redux/lib/bindActionCreators'
import Select from 'react-select'

//Actions
import { getSettings, saveSettings, reset } from '../actions/SettingsAction';

class Settings extends React.Component {
    constructor(...props) {
        super(...props)

        this.state = {
            time: '',
            groupby: ''
        }
    
        this.handleTime = this.handleTime.bind(this)
        this.handleGroupBy = this.handleGroupBy.bind(this)
        this.saveSettings = this.saveSettings.bind(this)
    }

    componentDidMount() {
        //this.props.settings.slo_metrics.length == 0 ? this.props.getSettings() : null
        this.props.getSettings();
    }

    // componentDidMount() {
    //     const { settings } = this.props
    //     this.setState({
    //         time: settings.time,
    //         groupby: settings.groupby
    //     })
    // }

    handleTime = (time) => {
        this.setState({
            time: `${time.label}`,
            groupby: this.state.groupby
        })
    }

    handleGroupBy = (groupby) => {
        this.setState({
            time: this.state.time,
            groupby: `${groupby.label}`
        })
    }

    saveSettings = () =>  {
        console.log("saveSettings")
        this.props.reset();
        this.props.saveSettings();
    }

    render() {
        console.log('RENDER <Settings>')
        console.log(this.state.time)
        console.log(this.state.groupby)
        return (
            <div className='setting_main'>
                <div className='selects'>
                    <Select
                            name='time'
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
                            name='limit'
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
                <button type='submit' className='btn btn-primary' onClick={this.saveSettings}>Save</button>
            </div>
        )
    }
}

function mapStateToProps(state) {
  return {
    settings: state.default.settings
  }
}
function mapDispatchToProps(dispatch) {
  return {
    getSettings: bindActionCreators(getSettings, dispatch),
    saveSettings: bindActionCreators(saveSettings, dispatch),
    reset: bindActionCreators(reset, dispatch),
  }
}
export default connect(mapStateToProps, mapDispatchToProps)(Settings)