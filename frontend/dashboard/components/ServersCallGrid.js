import React from 'react';
import { LineChart, BarChart, Bar, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts';

export default class ChartCell extends React.Component {
    constructor(...props) {
        super(...props)
    }

    _renderLine(chartData){
        return Object.entries(chartData).map(([key, value], i) => {
            return (
                <div key={i}>
                    <LineChart width={600} height={300} data={value} margin={{top: 10, right: 30, left: 30, bottom: 10}}>
                        <XAxis dataKey='timestamp'/>
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
            return (
                <div key={i}>
                    <BarChart width={600} height={300} data={value} margin={{top: 10, right: 30, left: 30, bottom: 10}}>
                        <XAxis dataKey='timestamp'/>
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

    render() {
        let { chartData } = this.props
        
        let result = []

        console.log(chartData)

        if(chartData.Metrics != undefined) {
            for (let index = 0; index < chartData.Metrics.length; index++) {
                if (chartData.Metrics[index].ChartType == 'line') {
                    result.push(this._renderLine(chartData.Metrics[index].Metric))
                } else if(chartData.Metrics[index].ChartType == 'bar') {
                    result.push(this._renderBar(chartData.Metrics[index].Metric))
                }
                
            }
            return <div>{ result }</div>
        }

        return <div>no data</div>
    }
}
ChartCell.propTypes = {
  
}


