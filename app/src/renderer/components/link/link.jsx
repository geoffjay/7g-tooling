import React from 'react';
import PropTypes from 'prop-types';
import classnames from 'classnames';
import omit from 'lodash/omit';

import { getBoxComponent, marginPropTypes, marginPropKeys, flexPropKeys, flexPropTypes } from '../box';


const Box = getBoxComponent([ ...marginPropKeys, ...flexPropKeys ]);

const Link = ({ children, openInNewTab, href, text, className, ...props }) => {
    const linkClassNames = classnames(
        className,
        { 'js-prevent-router': openInNewTab },
    );

    return (
        <Box
            is="a"
            target={openInNewTab ? '_blank' : '_self'}
            rel={openInNewTab ? 'noreferrer noopener' : null}
            href={href}
            className={linkClassNames}
            {...omit(props, ['isDisabled'])}
        >
            {children || (text ? text : href)}
        </Box>
    );
};


Link.propTypes = {
    ...flexPropTypes,
    ...marginPropTypes,
    /**
    * elements (such as an `<DeprecatedIcon />`) pulled in from component
    */
    children: PropTypes.node,
    /**
    * Optional className
    */
    className: PropTypes.string,
    /**
    * Link to navigate to
    */
    href: PropTypes.string,
    /**
    * Link open in new tab
    */
    openInNewTab: PropTypes.bool,
    /**
    * Text to display as link
    */
    text: PropTypes.string,
    /**
    * testId to get this component by (rendered by Box)
    */
    testId: PropTypes.string,
    /**
    * Type of link
    */
    type: PropTypes.string,
    /**
    * If link leads to a downloadable file, the file will be named here
    */
    download: PropTypes.string,
};

Link.defaultProps = {
    children: null,
    className: '',
    download: null,
    href: null,
    openInNewTab: false,
    type: 'text/html',
    text: '',
    testId: undefined,
};

export default Link;
