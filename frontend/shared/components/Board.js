import React, {PropTypes} from 'react'
import { Table } from 'react-materialize'

const Board = ({chartName, value}) => {
    
    console.log(value)
    
    console.log('RENDER <Board>')

    if (chartName != undefined) {
        return (
            <div className='table_main'>
                <div className='table_name'>{chartName}</div>
                <Table className='table' responsive={true} bordered={true}>
                    <thead>
                    <tr>
                        {
                            value.Titles.map((element, i) => {
                                return<th key={i} data-field={element}>{element}</th>
                            })
                        }
                    </tr>
                    </thead>
                    <tbody>
                    {
                        value.Value.map((array, j) => {
                            return <tr key={j}>
                                {
                                    array.map((val, k) => {
                                        return <th key={k} data-field={val}>{val}</th>
                                    })
                                }
                            </tr>
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
    value: PropTypes.object.isRequired,
}

export default Board