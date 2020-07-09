import React, { PureComponent } from 'react';
import _omit from 'lodash/omit';

import Link from '7g-components/link/link.jsx';
import DotLoader from '7g-components/dot-loader/dot-loader.jsx';
import { marginPropTypes } from '7g-components/box';

import ButtonIcon from './button-icon.jsx';
import ButtonBase, { buttonBaseClassName } from './button-base.jsx';
import buttonConstants from './button-constants.js';
import { propTypes as buttonPropTypes, defaultProps } from './button-prop-types.js';

import './button.scss';


const propTypes = {
    ...marginPropTypes,
    ...buttonPropTypes,
};

class Button extends PureComponent {
    static Base = (props) => {
        const textOrChildren = props.text
            ? <span>{props.text}</span>
            : props.children;
        const content = props.isLoading
            ? <DotLoader size="small" />
            : textOrChildren;
        const icon = !!props.icon && !props.isLoading && (
            <ButtonIcon
                icon={props.icon}
                hasText={!!textOrChildren}
            />
        );
        return props.href ? (
            <Link
                {..._omit(props, ['isA11yExpanded', 'a11yLabel', 'a11yControls', 'isDisabled', 'isActive', 'isLoading', 'text'])}
                aria-expanded={props.isA11yExpanded}
                aria-label={props.a11yLabel}
                aria-controls={props.a11yControls}
                disabled={props.isDisabled || props.isLoading}
                onClick={props.onClick}
                className={buttonBaseClassName(props)}
            >
                {icon}{content}
            </Link>
        ) : (
            <ButtonBase {..._omit(props, ['text'])}>
                {icon}{content}
            </ButtonBase>
        );
    };

    static Outline = (props) => (
        <Button.Base
            {...props}
            onClick={Button.getOnClick(props)}
            type={buttonConstants.TYPE.OUTLINE}
        />
    );

    static Filled = (props) => (
        <Button.Base
            {...props}
            onClick={Button.getOnClick(props)}
            type={buttonConstants.TYPE.FILLED}
        />
    );

    static Minimal = (props) => (
        <Button.Base
            {...props}
            onClick={Button.getOnClick(props)}
            type={buttonConstants.TYPE.MINIMAL}
        />
    );

    static Inline = props => (
        <Button.Base
            {...props}
            onClick={Button.getOnClick(props)}
            type={buttonConstants.TYPE.INLINE}
        />
    );

    static handleDisabledClick = () => null;

    static getOnClick = ({ isDisabled, onClick }) => isDisabled
        ? Button.handleDisabledClick
        : onClick;

    static displayName = 'Button';

    static propTypes = propTypes;

    static defaultProps = defaultProps;

    render() {
        return (
            <Button.Base
                {...this.props}
                onClick={Button.getOnClick(this.props)}
            />
        );
    }
}

/* eslint-disable react/static-property-placement */
Button.Outline.propTypes = propTypes;
Button.Outline.defaultProps = defaultProps;
Button.Outline.displayName = 'Button.Outline';

Button.Filled.propTypes = propTypes;
Button.Filled.defaultProps = defaultProps;
Button.Filled.displayName = 'Button.Filled';

Button.Minimal.propTypes = propTypes;
Button.Minimal.defaultProps = defaultProps;
Button.Minimal.displayName = 'Button.Minimal';

Button.Inline.propTypes = propTypes;
Button.Inline.defaultProps = defaultProps;
Button.Inline.displayName = 'Button.Inline';

Button.Base.propTypes = propTypes;
Button.Base.defaultProps = defaultProps;
Button.Base.displayName = 'Button.Base';
/* eslint-enable */

export default Button;
