import React from 'react';
import _omit from 'lodash/omit';

import Box from './box';
import { getBoxPropsBlacklist } from './utils';

/**
 * Configures a Box component for your special needs.
 *
 * Pass it the keys from the extensibleApisPropTypes you would like to be configurable
 * and it returns a Box component that only allows those props to be configurable.
 * @param {Array of keys} boxPropsToUse
 */
export const getBoxComponent = (boxPropsToUse) => {
    const boxPropsBlacklist = getBoxPropsBlacklist(boxPropsToUse);
    // eslint-disable-next-line react/display-name
    return (props) => {
        return (
            <Box {..._omit(props, boxPropsBlacklist)} />
        );
    };
};
