import React, { PureComponent } from 'react';
import PropTypes from 'prop-types';
import classnames from 'classnames';

import colors from '../../utils/colors.js';
import { getBoxComponent, marginPropTypes, marginPropKeys } from '../box';


const Box = getBoxComponent(marginPropKeys);
export default class Icon extends PureComponent {

    static propTypes = {
        // Add spacing api from Box
        ...marginPropTypes,
        /**
         * Color: Can be any colour variable specified in colors.less
         * defaults to currentColor
         */
        color: PropTypes.string,
        /**
         * Icon name
         */
        icon: PropTypes.string.isRequired,
        /**
         * Icon size: Recommended sizes are 12, 16, 20, 24, 40. Defaults to inherit
         */
        size: PropTypes.oneOfType([
            PropTypes.string,
            PropTypes.number,
        ]),
        /**
         * Icon test id
         */
        testId: PropTypes.string,
        /**
         * Icon css vertical align
         */
        verticalAlign: PropTypes.string,
    };

    static defaultProps = {
        color: 'inherit',
        size: 'inherit',
        verticalAlign: '-15%',
        testId: undefined, // Handled by Box
    };

    render() {
        const {
            color: propColor,
            size: propSize,
            verticalAlign,
            icon,
            ...restProps
        } = this.props;
        const className = classnames(
            'js-test-icon-component',
        );
        const color = propColor === 'inherit' ? 'inherit' : colors[propColor];
        if (!color) {
            throw new Error(`This colour does not exist: ${propColor}`);
        }
        let fontSize = 'inherit';
        if (propSize !== 'inherit') {
            const size = Number(propSize);
            fontSize = isNaN(size) ? propSize : `${size}px`;
        }
        const styleSvg = {
            color,
            fill: 'currentColor',
            display: 'inline-block',
            fontSize,
            height: `1em`,
            width: `1em`,
            lineHeight: `1em`,
            userSelect: 'none',
            flexShrink: '0',
            verticalAlign,
        };
        return (
            <Box is="svg" style={styleSvg} {...restProps} className={className}>
                <use xlinkHref={`#icon-${icon}`} />
            </Box>
        );
    }
}
