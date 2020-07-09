import React from 'react';
import PropTypes from 'prop-types';

import Box from '7g-components/box/box.jsx';
import Text from '7g-components/text/text.jsx';


const TableEmptyMessage = ({ children }) => {
    return (
        <Box paddingY="24px" justifyContent="center" testId="table-empty-message">
            <Text bold="bold">{children}</Text>
        </Box>
    );
};

TableEmptyMessage.propTypes = {
    children: PropTypes.string.isRequired,
};

export default TableEmptyMessage;
