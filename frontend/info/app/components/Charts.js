import React from 'react'
import { Chart, Board } from '../../../shared/components'

 const Charts = ({data}) => {

    const renderLine = (data, i) => {
        return <Chart chartName={data.ChartName} value={data.Metric[0].Value} key={i}/>
    } 

    const renderTable = (data, i) => {
        return <Board chartName={data.ChartName} value={data.Metric[0]} key={i}/>
    }

    let charts = []
    let tabels = []

    if(data.length > 0) {

        for (let index = 0; index < data.length; index++) {

            switch (data[index].ChartType) {
                case 'counter':
                    charts.push(renderLine(data[index], index))
                    break
                case  'histogram':
                    // charts.push(this._renderBar(chartData.Metrics[index].Metric))
                    break
                case 'table':
                    tabels.push(renderTable(data[index], index))
                    break
            }
        }
        return (
            <div className='table'>
                {tabels}
            </div>
        )
    }
    return <div className='no_data'>no data</div>
}

export default Charts