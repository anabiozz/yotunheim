import React from 'react'
import PropTypes from 'prop-types'
import Link from 'react-router/lib/Link'

const Navbar =  () => {
    
    console.log('RENDER <Navbar>')

    return <nav className='navbar navbar-inverse'>

        <div className='navbar-header'>
            <Link to={'/main'} className='navbar-brand'><span>Heimdall</span></Link>
        </div>

    </nav>
}

Navbar.propTypes = {
    name: PropTypes.string.isRequired
}

export default Navbar