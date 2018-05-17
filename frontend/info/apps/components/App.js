import React from 'react'
import { Chart, Board } from '../../../shared/components'
import Link from 'react-router/lib/Link'

 const App = ({name, data}) => {

    return (
        <Link to={'/server/apps/'+name} className='app_block z-depth-5'>
            <p>{name}</p>
            <div className='common_info'>
                {
                    data.map((object, index) => {
                        return <div  key={index}> {object.Title + ": "} <span>{object.Value}</span> </div>
                    })
                }
            </div>
        </Link>
    )
}

export default App