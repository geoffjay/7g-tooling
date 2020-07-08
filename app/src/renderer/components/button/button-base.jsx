import React from 'react';
import classnames from 'classnames';

import { getBoxComponent, marginPropTypes, marginPropKeys } from '7g-components/box';

import { propTypes as buttonPropTypes, defaultProps } from './button-prop-types.js';


const Box = getBoxComponent(marginPropKeys);

const propTypes = {
    ...marginPropTypes,
    /**
     * It may look like testId is not being handled, but Box adds it as a className, the reason its removed from here,
     * is because in testing, if both places handles it, you end up with two nodes with the same `.js-test-name`, which is
     * a little unexpected and means you have to code a little differently. This change does not change how rendering works
     * in the browser, and rather simulates it a little better.
     */
    ...buttonPropTypes,
};

export const buttonBaseClassName = ({
    type,
    color,
    size,
    isActive,
    isLoading,
    isDisabled,
    className,
}) => (
    classnames(
        'button-component',
        `button-component--type-${type}`,
        `button-component--color-${color}`,
        `button-component--size-${size}`,
        {
            'button-component--disabled': isDisabled || isLoading,
            'button-component--active': isActive,
            'js-test-button-active': isActive,
            [className]: !!className,
        },
    )
);

/**
 *  Note: This component isn't really designed to be used directly unless you are
 *  sure that you know what you are doing. Use the Button component instead.
 */

const ButtonBase = ({
    isA11yExpanded,
    a11yLabel,
    a11yControls,
    className,
    children,
    color,
    isActive,
    isDisabled,
    isLoading,
    size,
    type,
    onClick,
    testId,
    ...props
}) => (
    <Box
        is="button"
        aria-expanded={isA11yExpanded}
        aria-label={a11yLabel}
        aria-controls={a11yControls}
        type="button"
        className={buttonBaseClassName({
            type,
            color,
            size,
            isDisabled,
            isLoading,
            isActive,
            className,
        })}
        disabled={isDisabled || isLoading}
        onClick={onClick}
        testId={testId}
        {...props}
    >
        {children}
    </Box>
);

ButtonBase.propTypes = propTypes;
ButtonBase.defaultProps = defaultProps;

export default ButtonBase;
