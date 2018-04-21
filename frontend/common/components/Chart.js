import React from 'react'
import PropTypes from 'prop-types'
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts'

export default class Chart extends React.Component {
    render() {

        const { chartName, value, i } = this.props

        console.log(chartName)

        console.log('RENDER <Chart>')

        if (chartName != undefined) {
            return (
                <div className='chart' key={i}>
                    <LineChart width={600} height={300} data={value} margin={{top: 10, right: 30, left: 30, bottom: 10}}>
                        <XAxis dataKey='xline'/>
                        <YAxis domain={[0, 100]} />
                        <CartesianGrid strokeDasharray='3 3'/>
                        <Tooltip/>
                        <Legend />
                        <Line type='monotone' dataKey='payload' name={chartName + ' usage'} stroke='#8884d8' activeDot={{r: 8}}/>
                    </LineChart>
                </div>
            )
        } else {
            return null
        }
    }
}

Chart.propTypes = {
    chartName: PropTypes.string.isRequired,
    value: PropTypes.array.isRequired,
    i: PropTypes.number.isRequired,
}