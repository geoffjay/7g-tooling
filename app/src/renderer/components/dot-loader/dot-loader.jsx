import React from 'react';
import PropTypes from 'prop-types';
import classnames from 'classnames';


const DotLoader = ({ appLoader, contrast, size, inline }) => {
    const classes = classnames('loader__container', 'loading-pulse', {
        'loading-light': contrast,
        'loading-smaller': size === 'small',
        'loading-smallest': size === 'smallest',
    });
    const style = !inline ? {} : {
        display: 'inline-block',
    };
    if (appLoader) {
        style.textAlign = 'center';
        style.paddingTop = '50px';
    }
    return (
        <div className={classes} style={style}>
            <div /><div /><div />
        </div>
    );
};

DotLoader.propTypes = {
    /**
     * set to true if this is the dotloader shown when loading an app section
     */
    appLoader: PropTypes.bool,
    /**
     * Dark and light versions of the loader
     */
    contrast: PropTypes.bool,
    /**
     * set display css to 'inline-block'
     */
    inline: PropTypes.bool,
    /**
     * leave blank, or use "small" and "smallest"
     */
    size: PropTypes.oneOf(['', 'small', 'smallest']),
};

DotLoader.defaultProps = {
    appLoader: false,
    contrast: false,
    size: '',
    inline: false,
};

export default DotLoader;
