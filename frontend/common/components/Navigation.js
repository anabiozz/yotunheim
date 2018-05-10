import React from 'react'
import Link from 'react-router/lib/Link'

const Navbar =  () => {
    
    console.log('RENDER <Navbar>')

    return <nav className='navbar navbar-inverse z-depth-5'>

        <div className='navbar-header'>
            <Link to={'/main'} className='navbar-brand'><span>Heimdall</span></Link>
        </div>

    </nav>
}

export default Navbar