import React from 'react'
import { LineChart, BarChart, Bar, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts'
import ReactTable from 'react-table'
import {
  Table,
  TableBody,
  TableHeader,
  TableHeaderColumn,
  TableRow,
  TableRowColumn,
} from 'material-ui/Table';
import injectTapEventPlugin from 'react-tap-event-plugin';
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider'
import {green100, green500, green700} from 'material-ui/styles/colors';
import getMuiTheme from 'material-ui/styles/getMuiTheme';

export default class ChartCell extends React.Component {
    constructor(...props) {
        super(...props)
    }

    _renderLine(chartData){
        return Object.entries(chartData).map(([key, value], i) => {
            return (
                <div className='chart' key={i}>
                    <LineChart width={600} height={300} data={value} margin={{top: 10, right: 30, left: 30, bottom: 10}}>
                        <XAxis dataKey='xline'/>
                        <YAxis domain={[0, 100]} />
                        <CartesianGrid strokeDasharray='3 3'/>
                        <Tooltip/>
                        <Legend />
                        <Line type='monotone' dataKey='payload' name={key + ' usage'} stroke='#8884d8' activeDot={{r: 8}}/>
                    </LineChart>
                </div>
            )
        })
    }

    _renderBar(chartData){
        return Object.entries(chartData).map(([key, value], i) => {
            console.log(key)
            return (
                <div className='chart' key={i}>
                    <BarChart width={600} height={300} data={value} margin={{top: 10, right: 30, left: 30, bottom: 10}}>
                        <XAxis dataKey='xline'/>
                        <YAxis domain={[0, 100]} />
                        <CartesianGrid strokeDasharray='3 3'/>
                        <Tooltip/>
                        <Legend />
                        <Bar dataKey='payload' name={key + ' usage'} fill='#8884d8' />
                    </BarChart>
                </div>
            )
        })
    }

    _renderTable(chartData){

        const muiTheme = getMuiTheme({
            palette: {
              primary1Color: green500,
              primary2Color: green700,
              primary3Color: green100,
            },
            },
            {
            avatar: {
              borderColor: null,
            },
          });

        return Object.entries(chartData).map(([key, value], i) => {
           
            value = value[0].payload_arr

            return (
                <div key={i}>
                    <MuiThemeProvider muiTheme={muiTheme}>
                        <Table>
                        <TableHeader>
                            <TableRow key={i}>
                                {
                                    Object.keys(value[0]).map(key => 
                                        <TableHeaderColumn key={i}>{key}</TableHeaderColumn>
                                    )
                                }
                            </TableRow>
                        </TableHeader>
                        <TableBody>
                            {
                                value.map(function(val, i) {
                                    return (
                                        <TableRow key={i}>
                                        {
                                            Object.keys(val).map(key => 
                                                <TableRowColumn key={i}> {val[key]}</TableRowColumn>
                                            )
                                        }
                                        </TableRow>
                                    )
                                })
                            }
                        </TableBody>
                        </Table>
                    </MuiThemeProvider>
                </div>
            )
        })
    }

    render() {
        let { chartData } = this.props
        console.log(chartData)
        let result = []
        if(chartData.Metrics !== undefined) {
            for (let index = 0; index < chartData.Metrics.length; index++) {
                switch (chartData.Metrics[index].ChartType) {
                    case 'counter':
                        result.push(this._renderLine(chartData.Metrics[index].Metric))
                        break
                    case  'histogram':
                        result.push(this._renderBar(chartData.Metrics[index].Metric))
                        break
                    case  'table':
                        result.push(this._renderTable(chartData.Metrics[index].Metric))
                        break
                }
            }
            return <div className='row'>{ result }</div>
        }
        return <div>no data</div>
    }
}

ChartCell.propTypes = {}


