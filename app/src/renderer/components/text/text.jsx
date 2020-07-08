import React, { Component } from 'react';
import classnames from 'classnames';

import colors from '7g-utils/colors.js';

import { propTypes, defaultProps } from './text-proptypes.js';

import './text.scss';

const getStyle = ({ color: propColor = 'textColor', error, success, italic, bold, wordBreak, textTransform }) => {
    const color = propColor === 'inherit' ? 'inherit' : colors[propColor];
    const style = {
        color
    };
    if (error) {
        style.color = colors.colorRed;
    } else if (success) {
        style.color = colors.greenText;
    }
    if (italic) {
        style.fontStyle = 'italic';
    }
    if (bold) {
        style.fontWeight = '600';
    }
    if (wordBreak) {
        style.wordBreak = wordBreak;
    }
    if (textTransform) {
        style.textTransform = textTransform;
    }
    return style;
};

const getClassNames = (className, { testId }) => {
    return classnames(className, { [`js-test-${testId}`]: testId });
};

export default class Text extends Component {
    static LargeText(props) {
        return (
            <span className={getClassNames('text--large', props)} style={getStyle(props)}>
                {props.children}
            </span>
        );
    }

    static SmallText(props) {
        return (
            <span className={getClassNames('text--small', props)} style={getStyle(props)}>
                {props.children}
            </span>
        );
    }

    static MicroText(props) {
        return (
            <span className={getClassNames('text--micro', props)} style={getStyle(props)}>
                {props.children}
            </span>
        );
    }

    static propTypes = propTypes;

    static defaultProps = defaultProps;

    render() {
        const { children, ...props } = this.props;
        const style = getStyle(props);
        return (
            <span className={getClassNames('text', props)} style={style}>
                {children}
            </span>
        );
    }
}

Text.LargeText.propTypes = propTypes;
Text.SmallText.propTypes = propTypes;
Text.MicroText.propTypes = propTypes;
Text.LargeText.defaultProps = defaultProps;
Text.SmallText.defaultProps = defaultProps;
Text.MicroText.defaultProps = defaultProps;

export const getTextComponentBySize = (size) => {
    if (size === 'micro') {
        return Text.MicroText;
    }
    if (size === 'small') {
        return Text.SmallText;
    }
    if (size === 'large') {
        return Text.LargeText;
    }
    return Text;
};
