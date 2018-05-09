import React, { PropTypes } from 'react'
import connect from 'react-redux/lib/connect/connect'
//Components
import Navbar from '../components/Navigation'
import Switcher from '../components/Switcher'

class Layout extends React.Component {
	constructor(...props) {
		super(...props)
	}
	
	render() {
		const { mode } = this.props
		console.log('RENDER <Layout>')
		console.log(mode)

		return <div className='container-fluid'>
						{<Navbar mode={mode} />}
						{<div className='main'>
							{<Switcher />}
							{this.props.children}
						</div>}
					
        </div>
	}
}

Layout.propTypes = {
	children: PropTypes.object.isRequired
}

function mapStateToProps (state) {
	return {
		user: state
	}
}

export default connect(mapStateToProps)(Layout)
