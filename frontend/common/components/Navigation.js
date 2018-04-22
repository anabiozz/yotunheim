import React from 'react'
import PropTypes from 'prop-types'
import Link from 'react-router/lib/Link'

const Navbar =  ({name}) => {
    
    console.log('RENDER <Navbar>')

    return <nav className='navbar navbar-inverse'>

        <div className='navbar-header'>
            <Link to={'/dashboard'} className='navbar-brand'><span>Heimdall</span></Link>
        </div>

        {name !== 'Unknown' && <div id='navbar' className='collapse navbar-collapse'>

            <ul className={'nav navbar-nav navbar-right'}>
                <li><Link to='/settings' className='font16px'><span className='glyphicon glyphicon-cog font16px'/></Link></li>
            </ul>

        </div>}
    </nav>
}

Navbar.propTypes = {
    name: PropTypes.string.isRequired
}

export default Navbar