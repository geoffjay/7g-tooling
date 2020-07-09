import PropTypes from 'prop-types';


export const propTypes = {
    bold: PropTypes.oneOf(['normal', 'bold', 'bolder', 'lighter', 100, 200, 300, 400, 500, 600, 700, 800, 900]),
    children: PropTypes.node.isRequired,
    color: PropTypes.string,
    error: PropTypes.bool,
    italic: PropTypes.bool,
    success: PropTypes.bool,
    testId: PropTypes.string,
    wordBreak: PropTypes.oneOf(['normal', 'break-all', 'keep-all', 'break-word', 'initial', 'inherit']),
    textTransform: PropTypes.oneOf(['capitalize', 'uppercase', 'lowercase', 'none', 'full-width', 'full-size-kana', 'initial', 'inherit']),
};

export const defaultProps = {
    color: 'slateText',
    error: false,
    success: false,
    italic: false,
    bold: null,
    testId: null,
    wordBreak: undefined,
    textTransform: undefined,
};
