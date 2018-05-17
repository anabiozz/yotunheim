import React from 'react'
import Link from 'react-router/lib/Link'

const Navbar =  () => {
    
    console.log('RENDER <Navbar>')

    return <nav className='navbar'>
        <div className='navbar-header'>
            <Link to={'/server/common'} className='navbar-brand'><span>Heimdall</span></Link>   
        </div>
        <div className='settings'>
            <Link to={'/settings'} className='settigns_link'><span>Settings</span></Link> 
        </div> 

    </nav>
}

export default Navbar