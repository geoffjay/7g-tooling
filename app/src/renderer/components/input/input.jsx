import React, { forwardRef } from 'react';
import PropTypes from 'prop-types';
import classnames from 'classnames';

import InputBase from './input-base.jsx';
import './input.scss';


export const Input = ({ type, placeholder, value, hasError, isDisabled, testId, onChange, forwardedRef, ...props }) => (
    <InputBase onChange={onChange}>
        {inputProps => (
            <input
                ref={forwardedRef}
                className={classnames('input-component', {
                    'input-component--disabled': isDisabled,
                    'input-component--error': hasError,
                    'input-component--number': type === 'number',
                    [`js-test-${testId}`]: testId,
                })}
                type={type}
                placeholder={placeholder}
                disabled={isDisabled}
                value={value}
                {...inputProps}
                {...props}
            />
        )}
    </InputBase>
);

Input.propTypes = {
    forwardedRef: PropTypes.object,
    /**
     *  If the input is disabled
     */
    isDisabled: PropTypes.bool,
    /**
     *  If the error class/styles should be applied
     */
    hasError: PropTypes.bool,
    /**
     *  Placeholder text for the input (optional)
     */
    placeholder: PropTypes.string,
    /**
     *  testId string, for testing
     */
    testId: PropTypes.string,
    /**
     *  What type of input is this?
     */
    type: PropTypes.oneOf([
        'text',
        'password',
        'number',
        'email',
    ]),
    /**
     *  Value for the input component, if you want to use as a controlled component (optional)
     */
    value: PropTypes.oneOfType([
        PropTypes.string,
        PropTypes.number,
    ]).isRequired,
    /**
     *  Change handler function, so that the consumer knows about changes to the value (optional)
     */
    onChange: PropTypes.func.isRequired,
};

Input.defaultProps = {
    forwardedRef: null,
    isDisabled: false,
    hasError: false,
    placeholder: '',
    testId: '',
    type: 'text',
};

// eslint-disable-next-line react/display-name
export default forwardRef((props, ref) => (
    <Input {...props} forwardedRef={ref} />
));
