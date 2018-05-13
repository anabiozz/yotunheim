import React, { PropTypes } from 'react'
import connect from 'react-redux/lib/connect/connect'
import Select from 'react-select';

export default class Selector extends React.Component {
	constructor(...props) {
        super(...props)
        
        this.state = {
            selectedOption: '',
            time: '',
            groupby: ''
        }

        this.handleChange = this.handleChange.bind(this)
	}
    
    handleChange = (selectedOption) => {
        this.setState({ selectedOption });
        console.log(`Selected: ${selectedOption.label}`);
        this.props.onTimeChange(selectedOption)
    }

	render() {
        console.log('RENDER <Selector>')
        
        const { name, options } = this.props

		return (
			<Select
                name={name}
                value={this.state.selectedOption}
                onChange={this.handleChange}
                options={options}
            />
		) 
	}
}

Selector.propTypes = {
    name: PropTypes.string.isRequired,
    options: PropTypes.array.isRequired,
    onTimeChange: PropTypes.func.isRequired,
}