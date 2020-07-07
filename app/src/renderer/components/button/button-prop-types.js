import _isString from 'lodash/isString';
import PropTypes from 'prop-types';

import buttonConstants from './button-constants.js';


export const propTypes = {
    /**
     *  aria-expanded for the Button
     */
    isA11yExpanded: PropTypes.bool,
    /**
     *  aria-label for the Button
     */
    a11yLabel: PropTypes.string,
    /**
     *  aria-controls for the Button
     *  Refers to the id of element that contains all the content that is shown or hidden on button press
     */
    a11yControls: PropTypes.string,
    /**
     *  Color for the button (defaults to blue)
     */
    color: PropTypes.oneOf([
        buttonConstants.COLOR.BLUE,
        buttonConstants.COLOR.RED,
    ]),
    children: (props, propName) => {
        if (_isString(props[propName])) {
            return new Error('For strings, use text prop instead of children prop. Alternatively wrap your children prop in <span />');
        }
    },
    /**
     *  use <a /> instead of <button />
     */
    href: PropTypes.string,
    /**
     *  show the icon on the left side of the text
     */
    icon: PropTypes.string,
    /**
     *  If the button is currently active or not
     */
    isActive: PropTypes.bool,
    /**
     *  If the button is currently disabled or not
     */
    isDisabled: PropTypes.bool,
    /**
     *  If the button is currently loading or not
     */
    isLoading: PropTypes.bool,
    /**
     *  Size of the button
     */
    size: PropTypes.oneOf([
        buttonConstants.SIZE.NORMAL,
        buttonConstants.SIZE.SMALL,
    ]),
    /**
     *  Button text
     */
    text: PropTypes.string,
    /**
     *  Test ID for selecting button during tests
     */
    testId: PropTypes.string,
    /**
     *  Type of button
     */
    type: PropTypes.oneOf(
        Object.keys(buttonConstants.TYPE).map(k => buttonConstants.TYPE[k]),
    ),
    /**
     *  Handle when the button is clicked
     */
    onClick: PropTypes.func,
};

export const defaultProps = {
    isA11yExpanded: null,
    a11yLabel: null,
    a11yControls: null,
    children: null,
    color: buttonConstants.COLOR.BLUE,
    href: null,
    icon: null,
    isDisabled: false,
    isLoading: false,
    isActive: false,
    size: buttonConstants.SIZE.NORMAL,
    text: '',
    testId: '',
    type: buttonConstants.TYPE.OUTLINE,
    onClick: () => null,
};
