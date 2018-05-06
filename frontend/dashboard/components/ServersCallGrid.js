import React from 'react'
import { Chart, Board } from '../../common/components'
import { Row, Col } from 'react-materialize'

 const ChartCell = ({data}) => {

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
                    charts.push(<Col s={4} l={4} m={4} className='grid-example'>{renderLine(data.Metrics[index].Metric)}</Col>)
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
                <Row className='chart_row'>
                    {charts}
                </Row>
                <Row className='table_row'>
                    <Col s={6} l={6} m={6} className='grid-example'>{tabels}</Col>
                </Row>
            </div>
        )
    }
    return <div className='no_data'>no data</div>
}

export default ChartCell