import React, {PropTypes} from 'react'
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts'

const Chart = ({chartName, value, i}) => {

    console.log('RENDER <Chart>')

    if (chartName != undefined) {
        return (
            <div className='chart' key={i}>

                <span className='chart_head_text'>{chartName}</span>
                
                <LineChart width={550} height={300} data={value}>

                    <XAxis 
                        dataKey='xline' 
                        tick={{stroke: '#ddd'}} 
                    />

                    <YAxis 
                        domain={[0, 100]} 
                        allowDecimals={false} 
                        tick={{stroke: '#ddd'}}
                        padding={{ bottom: 50 }}
                        width={110}
                    />

                    <CartesianGrid strokeDasharray='3 3'/>
                    <Tooltip/>
                    <Legend />
                    <Line 
                        type='monotoneY' 
                        dataKey='payload' 
                        legendType='none' 
                        stroke='#FFC300' 
                        dot={{ stroke: '#FFC300', strokeWidth: 2 }} 
                        activeDot={{ stroke: '#FFC300', strokeWidth: 2, r: 5 }} 
                    />
                </LineChart>
            </div>
        )
    } else {
        return null
    }
}

Chart.propTypes = {
    chartName: PropTypes.string.isRequired,
    value: PropTypes.array.isRequired,
    i: PropTypes.number.isRequired,
}

export default Chart