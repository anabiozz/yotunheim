import React from 'react';
import { LineChart, BarChart, Bar, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts';

export default class ChartCell extends React.Component {
    constructor(...props) {
        super(...props)
    }

    // _renderLine(chartData){
    //     return Object.entries(chartData.Metrics).map(([key, value], i) => {
            
    //     })
    // }

    _renderBar(chartData){
        return Object.entries(chartData.Metrics).map(([key, value], i) => {
            if (key == 'disk') {
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
            } else {
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
            }
        })
    }

    render() {
        let { chartData } = this.props
        console.log(chartData.Data)
        if(chartData.Metrics != undefined) {
            console.log(chartData.ChartType)
            return <div>
                {
                   this._renderBar(chartData)
                }
            </div>
        }
        return <div>no data</div>
    }
}
ChartCell.propTypes = {
  
}


