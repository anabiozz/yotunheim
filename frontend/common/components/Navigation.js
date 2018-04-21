import React from 'react'
import PropTypes from 'prop-types'
import Link from 'react-router/lib/Link'
import NavDropdown from 'react-bootstrap/lib/NavDropdown'
import MenuItem from 'react-bootstrap/lib/MenuItem'

export default class Navbar extends React.Component {
    render() {
        const { name, makeRefresh } = this.props
        console.log('RENDER <Navbar>')

        return <nav className='navbar navbar-inverse navbar-fixed-top'>
            <div className='navbar-header'>
                <Link to={'/dashboard'} className='navbar-brand'><span>Heimdall</span></Link>
            </div>
            {name !== 'Unknown' && <div id='navbar' className='collapse navbar-collapse'>
                <ul className={'nav navbar-nav navbar-right'}>
                    <li><a onClick={makeRefresh}><span className='glyphicon glyphicon-refresh font16px'/></a></li>
                    <li><Link to='/settings' className='font16px'><span className='glyphicon glyphicon-cog font16px'/></Link></li>
                    <NavDropdown eventKey='4' title={name} id='nav-dropdown' className='font16px'>
                        <MenuItem eventKey='4.1' href='/oauth2/logout'>Log out</MenuItem>
                    </NavDropdown>
                </ul>
            </div>}
        </nav>
    }
}

Navbar.propTypes = {
    name: PropTypes.string.isRequired,
    makeRefresh: PropTypes.func.isRequired
}
