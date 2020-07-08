import React, { PureComponent } from 'react';
import _isString from 'lodash/isString';
import _range from 'lodash/range';
import classNames from 'classnames';

import Icon from '7g-components/icon/icon.jsx';
// TODO: skeleton stuff
// import Skeleton from '7g-components/skeleton/skeleton.jsx';

import {
    tableProps,
    defaultTableProps,
    defaultRowProps,
    rowProps,
} from './table-prop-types.js';
import TableEmptyMessage from './table-empty-message.jsx';
import { SORTING, SORT_ICONS } from './table-constants';


import './table.scss';


class Table extends PureComponent {

    static Row = (props) => {
        const {
            row,
            index,
            cellHeight,
            columns,
            hasSummary,
            isLoading,
            keyProperty,
            className,
        } = props;

        const defaultCellRenderFunc = (key) => {
            return (dataRow) => dataRow[key];
        };

        // const initializingCellRenderFunc = () => <Skeleton.Bar />;

        const getCellRenderFunc = (key) => {
            // if (isLoading) {
            //     return initializingCellRenderFunc;
            // }
            const col = columns.find((c) => c.property === key);
            if (col.render) {
                return col.render;
            }
            return defaultCellRenderFunc(key);
        };

        return (
            <tr
                key={row[keyProperty]}
                className={classNames(
                    'js-test-table-row',
                    'sg-table__row',
                    {
                        'sg-table__row--summary': index === 0 && hasSummary,
                    },
                    className,
                )}
            >
                {columns.map(({ property, width, ellipsis, centered, cellAlignment }) => {
                    const style = {};
                    if (width) {
                        style.width = width;
                    }
                    if (ellipsis) {
                        style.whiteSpace = 'nowrap';
                        style.overflow = 'hidden';
                        style.textOverflow = 'ellipsis';
                    }
                    if (centered) {
                        style.display = 'flex';
                        style.alignItems = 'center';
                    }
                    if (cellHeight) {
                        style.height = cellHeight;
                    }
                    if (cellAlignment) {
                        style.textAlign = cellAlignment;
                    }

                    return (
                        <td key={property} className="sg-table__cell">
                            <div style={style}>
                                {getCellRenderFunc(property)(row)}
                            </div>
                        </td>
                    );
                })}
            </tr>
        );
    };

    static propTypes = tableProps;

    static defaultProps = defaultTableProps;

    static SORTING = SORTING;

    static SORT_ICONS = SORT_ICONS;

    defaultHeaderRenderFunc(header, sortable, sortIcon, onSort) {
        if (!sortable) {
            return <div className="sg-table__header--default">{header}</div>;
        }
        return (
            <button
                className="sg_table__sort-button js-test-sg-table-sortable-header"
                type="button"
                onClick={onSort}
            >
                {header}
                {sortIcon}
            </button>
        );
    }

    render() {
        const {
            columns,
            data,
            emptyMessage,
            loadingRows,
            isLoading,
            sortProperty,
            sortDirection,
            onSort,
            cellHeight,
            hasSummary,
            hasOuterPadding,
            keyProperty,
            testId,
        } = this.props;

        const headClass = classNames(
            'sg-table__thead',
            { 'sg-table__thead--with-outer-padding': hasOuterPadding },
        );
        const bodyClass = classNames(
            'sg-table__tbody',
            { 'sg-table__tbody--with-outer-padding': hasOuterPadding },
        );

        return (
            <React.Fragment>
                <table
                    className={classNames('sg-table', {
                        [`js-test-${testId}`]: !!testId,
                    })}
                >
                    <thead className={headClass}>
                        <tr>
                            {columns.map(({ header, property, sortable, width, headerAlignment }) => {
                                let sortIconName = Table.SORT_ICONS.unsorted;
                                let sortIconSize = '12';
                                let sortIconColor = 'slate30';
                                let newSortDirection = Table.SORTING.asc;
                                if (property === sortProperty) {
                                    sortIconName = Table.SORT_ICONS[sortDirection];
                                    sortIconColor = 'slate';
                                    if (sortDirection === Table.SORTING.asc) {
                                        newSortDirection = Table.SORTING.desc;
                                    }
                                    if (sortDirection !== Table.SORTING.unsorted) {
                                        sortIconSize = '8';
                                    }
                                }
                                const sortIcon = <Icon icon={sortIconName} size={sortIconSize} color={sortIconColor} />;
                                const handleSort = () => onSort(property, newSortDirection);

                                const style = {};
                                if (width) {
                                    style.width = width;
                                }
                                if (headerAlignment) {
                                    style.textAlign = headerAlignment;
                                }

                                return (
                                    <th key={property} className="sg-table__header">
                                        <div style={style}>
                                            {_isString(header) ? this.defaultHeaderRenderFunc(header, sortable, sortIcon, handleSort) : header(sortIcon, handleSort)}
                                        </div>
                                    </th>
                                );
                            })}
                        </tr>
                    </thead>

                    <tbody className={bodyClass}>
                        {isLoading ? _range(loadingRows).map((id) => (
                            <Table.Row
                                key={id}
                                hasSummary={hasSummary}
                                cellHeight={cellHeight}
                                index={id}
                                columns={columns}
                                keyProperty="id"
                                row={{ id }}
                                isLoading
                            />
                        )) : data.map((row, index) => {
                            const Row = this.props.customRowRenderer ? this.props.customRowRenderer : Table.Row;
                            return (
                                <Row
                                    key={`row-${row[keyProperty]}`}
                                    hasSummary={hasSummary}
                                    cellHeight={cellHeight}
                                    row={row}
                                    index={index}
                                    columns={columns}
                                    keyProperty={keyProperty}
                                />
                            );
                        })}
                    </tbody>
                </table>
                {data.length === 0 && emptyMessage && emptyMessage !== '' && !isLoading && (
                    React.isValidElement(emptyMessage) ? emptyMessage : <TableEmptyMessage testId="table-empty-message">{emptyMessage}</TableEmptyMessage>
                )}
            </React.Fragment>
        );
    }
}

/* eslint-disable react/static-property-placement */
Table.Row.propTypes = rowProps;
Table.Row.defaultProps = defaultRowProps;
/* eslint-enable */

export default Table;
