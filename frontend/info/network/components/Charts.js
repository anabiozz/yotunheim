import React from 'react'
import { Chart, Board } from '../../../shared/components'

 const Charts = ({data}) => {

    const renderLine = (data, i) => {
        return <Chart chartName={data.InfoName} value={data.Metric[0].Value} key={i}/>
    } 

    const renderTable = (data) => {
        return <Board chartName={'key'} value={data}/>
    }

    let charts = []
    let tabels = []

    if(data.length > 0) {

        for (let index = 0; index < data.length; index++) {

            switch (data[index].InfoType) {
                case 'counter':
                    charts.push(renderLine(data[index], index))
                    break
                case  'histogram':
                    // charts.push(this._renderBar(chartData.Metrics[index].Metric))
                    break
                case  'table':
                    tabels.push(renderTable(data.Metric))
                    break
            }
        }

        return (
            <div className='charts'>
                {charts}
            </div>
        )

    }
    return <div className='no_data'>no data</div>
}

export default Charts