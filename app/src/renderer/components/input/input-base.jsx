import { PureComponent } from 'react';
import PropTypes from 'prop-types';

import utils from '../../utils/utils';


class InputBase extends PureComponent {

    static propTypes = {
        /**
         *  Function-as-child, recives a getInputProps function which allows you to spread those props on the input component
         */
        children: PropTypes.func.isRequired,
        /**
         *  Change handler function, so that the consumer knows about changes to the value (optional)
         */
        onChange: PropTypes.func.isRequired,
    }

    get inputProps() {
        return {
            'data-lpignore': 'true',
            ...utils.getOnInputOrOnChangeProp(this.handleChange),
        };
    }

    handleChange = (event) => {
        const { value } = event.target;
        return this.props.onChange(value, event);
    };

    render() {
        return this.props.children(this.inputProps);
    }
}

export default InputBase;
