import React, { PropTypes } from 'react'
import connect from 'react-redux/lib/connect/connect'
//Components
import Navbar from '../components/Navigation'
import Switcher from '../components/Switcher'
import Footer from '../components/Footer'

class Layout extends React.Component {
	constructor(...props) {
		super(...props)
	}
	
	render() {
		console.log('RENDER <Layout>')

		return (
			<div className='react_root'>
				<Navbar/>
			<div className='wrapper'>
				<Switcher/>
			{this.props.children}
			</div>
				<Footer/>
			</div>
		) 
	}
}

Layout.propTypes = {
	children: PropTypes.object.isRequired
}

function mapStateToProps (state) {
	return {state}
}

export default connect(mapStateToProps)(Layout)
