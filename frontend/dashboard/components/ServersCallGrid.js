import React from 'react';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts';

export default class ChartCell extends React.Component {
    constructor(...props) {
        super(...props)
    }

    _renderObject(chartData){
        return Object.entries(chartData).map(([key, value], i) => {
			return (
				<div key={i}>
					<LineChart width={600} height={300} data={value} margin={{top: 5, right: 30, left: 20, bottom: 5}}>
                    <XAxis dataKey={key + ' usage'}/>
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

    render() {
        let { chartData } = this.props
        if(chartData != undefined) {
            return <div>
                {this._renderObject(chartData)}
            </div>
        }
        return <div>no data</div>
    }
}
ChartCell.propTypes = {
  
}


