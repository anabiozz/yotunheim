import React from 'react'
import { LineChart, BarChart, Bar, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts'
import {Table, Row, Col} from 'react-materialize';

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

        return Object.entries(chartData).map(([key, value], i) => {
            return (
                <div>
                    <div className='table_name'>{key}</div>
                    <Table className='table' responsive={true} centered={true} bordered={true}>
                        <thead>
                        <tr>
                            {
                                Object.keys(value[0].payload_arr[0]).map((key, j) =>
                                    <th key={j} data-field={key}>{key}</th>
                                )
                            }
                        </tr>
                        </thead>
                        <tbody>
                        {
                            value.map(function(val, i) {
                                return (
                                    <tr key={i}>
                                        {
                                            Object.keys(val.payload_arr[0]).map((key, j) =>
                                                <td key={j}>{val.payload_arr[0][key]}</td>
                                            )
                                        }
                                    </tr>
                                )
                            })
                        }
                        </tbody>
                    </Table>
                </div>
            )
        })
    }

    render() {
        let { chartData } = this.props
        // console.log(chartData)
        let charts = []
        let tabels = []
        if(chartData.Metrics !== undefined) {
            for (let index = 0; index < chartData.Metrics.length; index++) {
                switch (chartData.Metrics[index].ChartType) {
                    case 'counter':
                        charts.push(this._renderLine(chartData.Metrics[index].Metric))
                        break
                    case  'histogram':
                        charts.push(this._renderBar(chartData.Metrics[index].Metric))
                        break
                    case  'table':
                        tabels.push(this._renderTable(chartData.Metrics[index].Metric))
                        break
                }
            }
            return (
                <div>
                    <Row className='chart_row'>
                        <Col l={12} m={12} className='grid-example'>{charts}</Col>
                    </Row>
                    <Row className='table_row'>
                        <Col l={12} m={12} className='grid-example'>{tabels}</Col>
                    </Row>
                </div>
            )
        }
        return <div>no data</div>
    }
}

ChartCell.propTypes = {}


