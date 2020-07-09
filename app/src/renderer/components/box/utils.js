import _without from 'lodash/without';

import {
    marginPropTypes,
    paddingPropTypes,
    spacingPropTypes,
    flexPropTypes,
    dimensionPropTypes,
    backgroundPropTypes,
    borderRadiusPropTypes,
    extensibleApisPropTypes,
} from './prop-types';


export const marginPropKeys = Object.keys(marginPropTypes);
export const paddingPropKeys = Object.keys(paddingPropTypes);
export const spacingPropKeys = Object.keys(spacingPropTypes);
export const flexPropKeys = Object.keys(flexPropTypes);
export const dimensionPropKeys = Object.keys(dimensionPropTypes);
export const backgroundPropKeys = Object.keys(backgroundPropTypes);
export const borderRadiusPropKeys = Object.keys(borderRadiusPropTypes);
export const extensibleApisPropKeys = Object.keys(extensibleApisPropTypes);


export const splitMarginProps = (props) => splitProps(props, marginPropKeys);
export const splitPaddingProps = (props) => splitProps(props, paddingPropKeys);
export const splitSpacingProps = (props) => splitProps(props, spacingPropKeys);
export const splitFlexProps = (props) => splitProps(props, flexPropKeys);
export const splitDimensionProps = (props) => splitProps(props, dimensionPropKeys);
export const splitBackgroundProps = (props) => splitProps(props, backgroundPropKeys);
export const splitBorderRadiusProps = (props) => splitProps(props, borderRadiusPropKeys);

export const splitBoxProps = (props) => splitProps(props, extensibleApisPropKeys);

export const getBoxPropsBlacklist = (allowedPropKeys) => {
    return _without(extensibleApisPropKeys, ...allowedPropKeys);
};

/**
 * Helper to help separate a props object
 * @param {Object} props
 * @param {Array} keys keys of the props that will be in matchedProps
 *
 * Returns
 * { matchedProps, remainingProps }
 */
export const splitProps = (props, keys) => {
    const matchedProps = {};
    const remainingProps = {};
    const propKeys = Object.keys(props);

    for (let i = 0; i < propKeys.length; i++) {
        const propKey = propKeys[i];
        const propValue = props[propKey];

        if (keys.includes(propKey)) {
            matchedProps[propKey] = propValue;
        } else {
            remainingProps[propKey] = propValue;
        }
    }

    return { matchedProps, remainingProps };
};
