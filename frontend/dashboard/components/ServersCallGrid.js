// first of course react!
var React = require('react');
// require `react-d3-core` for Chart component, which help us build a blank svg and chart title.
var Chart = require('react-d3-core').Chart;
// require `react-d3-basic` for Line chart component.
var LineChart = require('react-d3-basic').LineChart;

export default class ChartCell extends React.Component {
    constructor(...props) {
        super(...props)
    }

    render() {
        const { chartData } = this.props
        return <LineChart
            width= {700}
            height= {300}
            data= {chartData}
            chartSeries= {[
                {
                    field: 'age',
                    name: 'Age',
                    color: '#ff7f0e',
                    style: {
                        "strokeWidth": 2,
                        "strokeOpacity": .9,
                        "fillOpacity": .2
                    }
                }
            ]}
            x= {function(d) {
                return d.index;
            }}
            />
    }
}
ChartCell.propTypes = {
  
}
