import React from 'react'
import Link from 'react-router/lib/Link'

const Switcher =  () => {
    
    console.log('RENDER <Switcher>')

    return <div className='switcher z-depth-5' role='navigation'>
      <ul>
        <li><Link to='/main'><span className='font16px'/>Main</Link></li>
        <li><Link to='/apps'><span className='font16px'/>Apps</Link></li>
      </ul>
    </div>
}

export default Switcher