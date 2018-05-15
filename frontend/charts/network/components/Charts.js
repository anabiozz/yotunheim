import React from 'react'
import { Chart, Board } from '../../../shared/components'

 const Charts = ({data}) => {

    const renderLine = (data) => {
        return Object.entries(data).map(([key, value], i) => {
            return <Chart chartName={key} value={value} i={i}/>
        })
    } 

    const renderTable = (data) => {
        return Object.entries(data).map(([key, value]) => {
            return <Board chartName={key} value={value}/>
        })
    }

    let charts = []
    let tabels = []

    if(data.Metrics !== undefined) {

        for (let index = 0; index < data.Metrics.length; index++) {

            switch (data.Metrics[index].ChartType) {
                case 'counter':
                    charts.push(renderLine(data.Metrics[index].Metric))
                    break
                case  'histogram':
                    // charts.push(this._renderBar(chartData.Metrics[index].Metric))
                    break
                case  'table':
                    tabels.push(renderTable(data.Metrics[index].Metric))
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