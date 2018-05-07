import React from 'react'
import PropTypes from 'prop-types'
import Link from 'react-router/lib/Link'

const Switcher =  ({name}) => {
    
    console.log('RENDER <Switcher>')

    return <div className='switcher'>
      <ul>
        <li><Link to='/settings'><span className='glyphicon glyphicon-cog font16px'/>SADSA</Link></li>
      </ul>
    </div>
}

Switcher.propTypes = {
    name: PropTypes.string.isRequired
}

export default Switcher