import React, {PropTypes} from 'react'
import { Table } from 'react-materialize'

const Board = ({chartName, value}) => {
    
    console.log('RENDER <Board>')

    if (chartName != undefined) {
        return (
            <div>
                <div className='table_name'>{chartName}</div>
                <Table className='table' responsive={true} centered={true} bordered={true}>
                    <thead>
                    <tr>
                        {
                            Object.keys(value[0].payload_arr[0]).map((key, j) =>
                                <th key={j} data-field={key}>{key}</th>
                            )
                        }
                    </tr>
                    </thead>
                    <tbody>
                    {
                        value.map(function(val, i) {
                            return (
                                <tr key={i}>
                                    {
                                        Object.keys(val.payload_arr[0]).map((key, j) =>
                                            <td key={j}>{val.payload_arr[0][key]}</td>
                                        )
                                    }
                                </tr>
                            )
                        })
                    }
                    </tbody>
                </Table>
            </div>
        )
    } else {
        return null
    }
}

Board.propTypes = {
    chartName: PropTypes.string.isRequired,
    value: PropTypes.array.isRequired,
}

export default Board