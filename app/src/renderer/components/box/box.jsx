import React, { Component } from 'react';
import classnames from 'classnames';
import defaultTo from 'lodash/defaultTo';
import omit from 'lodash/omit';

import colors from '../../utils/colors';

import { propTypes, defaultProps } from './prop-types.js';
import { splitBoxProps } from './utils';


export const numberToPx = (size) => `${size}px`;

class Box extends Component {
    static displayName = 'Box'

    static propTypes = propTypes;

    static defaultProps = defaultProps;

    get className() {
        const { testId, className } = this.props;
        return classnames(
            className,
            'js-test-box-component',
            {
                [`js-test-${testId}`]: !!testId,
            },
        );
    }

    get hasFlexStyles() {
        const { alignItems, flexDirection, justifyContent, flexWrap, flex } = this.props;
        return !!(alignItems || flexDirection || justifyContent || flexWrap || flex);
    }

    getMarginStyles = () => {
        const marginTop = defaultTo(this.props.marginTop, defaultTo(this.props.marginY, this.props.margin));
        const marginBottom = defaultTo(this.props.marginBottom, defaultTo(this.props.marginY, this.props.margin));
        const marginLeft = defaultTo(this.props.marginLeft, defaultTo(this.props.marginX, this.props.margin));
        const marginRight = defaultTo(this.props.marginRight, defaultTo(this.props.marginX, this.props.margin));
        return { marginTop, marginBottom, marginLeft, marginRight };
    }

    getPaddingStyles = () => {
        const paddingTop = defaultTo(this.props.paddingTop, defaultTo(this.props.paddingY, this.props.padding));
        const paddingBottom = defaultTo(this.props.paddingBottom, defaultTo(this.props.paddingY, this.props.padding));
        const paddingLeft = defaultTo(this.props.paddingLeft, defaultTo(this.props.paddingX, this.props.padding));
        const paddingRight = defaultTo(this.props.paddingRight, defaultTo(this.props.paddingX, this.props.padding));
        return { paddingTop, paddingBottom, paddingLeft, paddingRight };
    }

    getBorderRadiusStyles = () => {
        const borderTopLeftRadius = defaultTo(this.props.borderTopLeftRadius, this.props.borderRadius);
        const borderBottomLeftRadius = defaultTo(this.props.borderBottomLeftRadius, this.props.borderRadius);
        const borderTopRightRadius = defaultTo(this.props.borderTopRightRadius, this.props.borderRadius);
        const borderBottomRightRadius = defaultTo(this.props.borderBottomRightRadius, this.props.borderRadius);
        return { borderTopLeftRadius, borderBottomLeftRadius, borderTopRightRadius, borderBottomRightRadius };
    }

    getBackgroundStyles = () => {
        if (this.props.backgroundColor) {
            return { background: colors[this.props.backgroundColor] };
        }
        return null;
    }

    getDimensionStyles = () => {
        const dimensions = {};
        if (this.props.height) {
            dimensions.height = this.props.height;
        }
        if (this.props.width) {
            dimensions.width = this.props.width;
        }
        return dimensions;
    }

    getFlexStyles = () => {
        if (!this.hasFlexStyles) {
            return {};
        }
        const { alignItems, flex, flexDirection, flexWrap, justifyContent } = this.props;
        return {
            alignItems,
            justifyContent,
            display: 'flex',
            flexDirection,
            flex: flex ? flex : '0 1 auto',
            flexWrap: flexWrap ? 'wrap' : null,
        };
    }

    getElement = () => {
        const { inline, is } = this.props;
        if (!is) {
            return inline ?
                'span' :
                'div';
        }
        return is;
    }

    render() {
        const { remainingProps } = splitBoxProps(this.props);
        const { style = {}, children, inline, is, ...rest } = remainingProps;

        const boxStyle = {
            ...this.getMarginStyles(),
            ...this.getPaddingStyles(),
            ...this.getBackgroundStyles(),
            ...this.getBorderRadiusStyles(),
            ...this.getDimensionStyles(),
            ...this.getFlexStyles(),
            ...style,
        };
        return React.createElement(
            this.getElement(inline, is), // handles is and inline
            {
                ...omit(rest, ['testId']),
                className: this.className,
                style: boxStyle,
            },
            children,
        );
    }
}

export default Box;
export * from './prop-types';
