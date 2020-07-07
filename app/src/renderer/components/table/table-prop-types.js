import PropTypes from 'prop-types';
import noop from 'lodash/noop';

export const tableProps = {
    columns: PropTypes.arrayOf(PropTypes.exact({
        header: PropTypes.oneOfType([
            PropTypes.func,
            PropTypes.string,
        ]).isRequired,
        property: PropTypes.string.isRequired,
        headerAlignment: PropTypes.string,
        cellAlignment: PropTypes.string,
        render: PropTypes.func,
        primary: PropTypes.bool,
        sortable: PropTypes.bool,
        ellipsis: PropTypes.bool,
        centered: PropTypes.bool,
        width: PropTypes.string,
    })).isRequired,
    cellHeight: PropTypes.string,
    data: PropTypes.arrayOf(PropTypes.object).isRequired,
    emptyMessage: PropTypes.oneOfType([PropTypes.string, PropTypes.element]),
    keyProperty: PropTypes.string.isRequired,
    hasSummary: PropTypes.bool,
    hasOuterPadding: PropTypes.bool,
    sortProperty: PropTypes.string,
    sortDirection: PropTypes.string,
    testId: PropTypes.string,
    onSort: PropTypes.func,
    customRowRenderer: PropTypes.func,
    isLoading: PropTypes.bool,
    loadingRows: PropTypes.number,
};

export const defaultTableProps = {
    cellHeight: null,
    emptyMessage: 'No matches found',
    testId: null,
    hasSummary: false,
    hasOuterPadding: true,
    sortProperty: null,
    sortDirection: null,
    onSort: noop,
    customRowRenderer: null,
    isLoading: false,
    loadingRows: 5,
};

export const rowProps = {
    className: PropTypes.string,
    hasSummary: PropTypes.bool,
    cellHeight: PropTypes.string,
    row: PropTypes.object.isRequired,
    index: PropTypes.number.isRequired,
    columns: PropTypes.arrayOf(PropTypes.exact({
        header: PropTypes.oneOfType([
            PropTypes.func,
            PropTypes.string,
        ]).isRequired,
        headerAlignment: PropTypes.string,
        cellAlignment: PropTypes.string,
        property: PropTypes.string.isRequired,
        render: PropTypes.func,
        primary: PropTypes.bool,
        sortable: PropTypes.bool,
        ellipsis: PropTypes.bool,
        centered: PropTypes.bool,
        width: PropTypes.string,
    })).isRequired,
    keyProperty: PropTypes.string.isRequired,
    isLoading: PropTypes.bool,
};

export const defaultRowProps = {
    hasSummary: false,
    cellHeight: null,
    className: '',
    isLoading: false,
};
