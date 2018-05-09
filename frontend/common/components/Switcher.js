import React from 'react'
import PropTypes from 'prop-types'
import Link from 'react-router/lib/Link'

const Switcher =  () => {
    
    console.log('RENDER <Switcher>')

    return <div className='switcher'>
      <ul>
        <li><Link to='/main'><span className='font16px'/>Main</Link></li>
        <li><Link to='/apps'><span className='font16px'/>Apps</Link></li>
      </ul>
    </div>
}

Switcher.propTypes = {
    name: PropTypes.string.isRequired
}

export default Switcher