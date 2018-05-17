import React, {PropTypes} from 'react'
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts'

const formattTimestamp = (element) => {
    let time = Date.parse(element)
    let date = new Date(time)
    let hours = date.getHours()
    let minutes = '0' + date.getMinutes()
    let seconds = '0' + date.getSeconds()
    return hours + ':' + minutes.substr(-2) + ':' + seconds.substr(-2)
}

const Chart = ({chartName, value}) => {
    
    let data = []
    value.forEach((element) => {
        let obj = {}
        obj['time'] = formattTimestamp(element[0])
        obj['value'] = element[1]
        data.push(obj)
    })

    console.log('RENDER <Chart>')

    return (
        <div className='chart z-depth-5' >
            <div className='chart_title'>
                <span className='chart_head_text'>{chartName}</span>
            </div>
            
            <LineChart width={700} height={300} data={data}>

                <XAxis 
                    dataKey='time' 
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
                    dataKey='value' 
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
    // i: PropTypes.number.isRequired,
}

export default Chart