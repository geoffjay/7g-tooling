import React from 'react';
import PropTypes from 'prop-types';
import classnames from 'classnames';

import Icon from '7g-components/icon/icon.jsx';


const ButtonIcon = ({ icon, hasText }) => {
    const buttonClassName = classnames('button-component__icon', {
        'button-component__icon--left': hasText,
    });
    return (
        <span className={buttonClassName}>
            <Icon icon={icon} />
        </span>
    );
};

ButtonIcon.propTypes = {
    icon: PropTypes.string.isRequired,
    hasText: PropTypes.bool,
};

ButtonIcon.defaultProps = {
    hasText: false,
};

export default ButtonIcon;
