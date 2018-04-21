import React from 'react'
import { Chart, Board } from '../../common/components'
import { Row, Col } from 'react-materialize'

export default class ChartCell extends React.Component {
    constructor(...props) {
        super(...props)

        this.renderLine = this.renderLine.bind(this)
        this.renderTable = this.renderTable.bind(this)
    }

    renderLine(data) {
        return Object.entries(data).map(([key, value], i) => {
            return <Chart chartName={key} value={value} i={i}/>
        })
    } 

    renderTable(data) {
        return Object.entries(data).map(([key, value], i) => {
            return <Board chartName={key} value={value} i={i}/>
        })
    }

    render() {
        let { data } = this.props
        let charts = []
        let tabels = []
        if(data.Metrics !== undefined) {
            for (let index = 0; index < data.Metrics.length; index++) {
                switch (data.Metrics[index].ChartType) {
                    case 'counter':
                        charts.push(this.renderLine(data.Metrics[index].Metric))
                        break
                    case  'histogram':
                        // charts.push(this._renderBar(chartData.Metrics[index].Metric))
                        break
                    case  'table':
                        tabels.push(this.renderTable(data.Metrics[index].Metric))
                        break
                }
            }
            return (
                <div>
                    <Row className='chart_row'>
                        <Col l={12} m={12} className='grid-example'>{charts}</Col>
                    </Row>
                    <Row className='table_row'>
                        <Col l={12} m={12} className='grid-example'>{tabels}</Col>
                    </Row>
                </div>
            )
        }
        return <div>no data</div>
    }
}

ChartCell.propTypes = {}