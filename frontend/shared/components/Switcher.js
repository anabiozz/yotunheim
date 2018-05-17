import React from 'react'
import Link from 'react-router/lib/Link'

const Switcher =  () => {
    
    console.log('RENDER <Switcher>')

    return <div className='switcher' role='navigation'>
      <ul>
        <li><Link to='/server/common'><span className='font16px'/>Common</Link></li>
        <li><Link to='/server/apps'><span className='font16px'/>Apps</Link></li>
        <li><Link to='/server/network'><span className='font16px'/>Network</Link></li>
      </ul>
    </div>
}

export default Switcher