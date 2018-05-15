import React, {PropTypes} from 'react'
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts'


const Chart = ({chartName, value, i}) => {

    console.log('RENDER <Chart>')

    return (
        <div className='chart z-depth-5' key={i}>
            <div className='chart_title'>
                <span className='chart_head_text'>{chartName}</span>
            </div>
            
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
                    type='natural' 
                    dataKey='payload' 
                    legendType='none' 
                    stroke='#ffcf32' 
                    dot={{ stroke: '#ffcf32', strokeWidth: 3 }} 
                    activeDot={{ stroke: '#ffcf32', strokeWidth: 3, r: 5 }} 
                />
            </LineChart>
        </div>
    )
}

Chart.propTypes = {
    chartName: PropTypes.string.isRequired,
    value: PropTypes.array.isRequired,
    i: PropTypes.number.isRequired,
}

export default Chart