import PropTypes from 'prop-types';
import range from 'lodash/range';


const validPxSizes = [0, 4, 8, 12, 16, 20, 24, 32, 40, 48, 56, 64, 72, 80].map(size => `${size}px`);
const validPercentageSizes = range(100).map(size => `${size}%`);

export const spacingSizePropType = (props, propName, componentName) => {
    if (props[propName] && (!validPxSizes.includes(props[propName]) && props[propName] !== 'auto')) {
        return new Error(`Invalid prop '${propName}' supplied to '${componentName}'. It should be a valid pixel size: (${validPxSizes.join(', ')}).`);
    }
};
export const radiusSizePropType = PropTypes.oneOfType([
    spacingSizePropType,
    PropTypes.oneOf(validPercentageSizes),
]);

export const isPercentOrPixelPropType = (props, propName, componentName) => {
    if (props[propName] && !/^(auto|0)$|^[+-]?[0-9]+.?([0-9]+)?(px|%)$/.test(props[propName])) {
        return new Error(`Invalid prop '${propName}' supplied to '${componentName}'. It should be a valid pixel or percentage size.`);
    }
};

export const marginPropTypes = {
    /**
     * Margin on a 8x grid
     */
    margin: spacingSizePropType,
    /**
     * Margin bottom on a 8x grid
     */
    marginBottom: spacingSizePropType,
    /**
     * Margin left on a 8x grid
     */
    marginLeft: spacingSizePropType,
    /**
     * Margin right on a 8x grid
     */
    marginRight: spacingSizePropType,
    /**
     * Margin top on a 8x grid
     */
    marginTop: spacingSizePropType,
    /**
     * Margin left & right on a 8x grid
     */
    marginX: spacingSizePropType,
    /**
     * Margin top & bottom on a 8x grid
     */
    marginY: spacingSizePropType,
};

export const marginDefaultProps = {
    margin: undefined,
    marginBottom: undefined,
    marginLeft: undefined,
    marginRight: undefined,
    marginTop: undefined,
    marginX: undefined,
    marginY: undefined,
};

export const paddingPropTypes = {
    /**
     * Padding on a 8x grid
     */
    padding: spacingSizePropType,
    /**
     * Padding bottom on a 8x grid
     */
    paddingBottom: spacingSizePropType,
    /**
     * Padding left on a 8x grid
     */
    paddingLeft: spacingSizePropType,
    /**
     * Padding right on a 8x grid
     */
    paddingRight: spacingSizePropType,
    /**
     * Padding top on a 8x grid
     */
    paddingTop: spacingSizePropType,
    /**
     * Padding left & right on a 8x grid
     */
    paddingX: spacingSizePropType,
    /**
     * Padding top & bottom on a 8x grid
     */
    paddingY: spacingSizePropType,
};

export const paddingDefaultProps = {
    padding: undefined,
    paddingBottom: undefined,
    paddingLeft: undefined,
    paddingRight: undefined,
    paddingTop: undefined,
    paddingX: undefined,
    paddingY: undefined,
};

export const spacingPropTypes = {
    ...marginPropTypes,
    ...paddingPropTypes,
};

export const spacingDefaultProps = {
    ...marginDefaultProps,
    ...paddingDefaultProps,
};

export const flexPropTypes = {
    /**
     * How to align the contents along the cross axis.
     */
    alignItems: PropTypes.oneOf([
        'flex-start',
        'center',
        'flex-end',
        'baseline',
        'stretch',
        'normal',
    ]),
    /**
     * flex for the Box's content:
     */
    flex: PropTypes.string,
    /**
     * flex-direction for the Box's content:
     */
    flexDirection: PropTypes.oneOf(['row', 'column', 'row-reverse', 'column-reverse']),
    /**
     *  Whether flex children will wrap if they can't all fit.
     */
    flexWrap: PropTypes.bool,
    /**
     * How to align the contents along the main axis.
     */
    justifyContent: PropTypes.oneOf([
        'flex-start',
        'center',
        'space-between',
        'space-around',
        'space-evenly',
        'flex-end',
        'normal',
    ]),
};

export const flexDefaultProps = {
    alignItems: undefined,
    flex: undefined,
    flexDirection: undefined,
    flexWrap: false,
    justifyContent: undefined,
};

export const dimensionPropTypes = {
    /**
     * Box height
     */
    height: isPercentOrPixelPropType,
    /**
     * Box width
     */
    width: isPercentOrPixelPropType,
};

export const dimensionDefaultProps = {
    height: undefined,
    width: undefined,
};

export const backgroundPropTypes = {
    /**
     * Background color
     */
    backgroundColor: PropTypes.string,
};

export const backgroundDefaultProps = {
    backgroundColor: undefined,
};

export const borderRadiusPropTypes = {
    /**
     * borderRadius on a 8x grid
     */
    borderRadius: radiusSizePropType,
    /**
     * borderRadius top left on a 8x grid
     */
    borderTopLeftRadius: radiusSizePropType,
    /**
     * borderRadius top right on a 8x grid
     */
    borderTopRightRadius: radiusSizePropType,
    /**
     * borderRadius bottom left on a 8x grid
     */
    borderBottomLeftRadius: radiusSizePropType,
    /**
     * borderRadius bottom right on a 8x grid
     */
    borderBottomRightRadius: radiusSizePropType,
};

export const borderRadiusDefaultProps = {
    borderBottomLeftRadius: undefined,
    borderBottomRightRadius: undefined,
    borderRadius: undefined,
    borderTopLeftRadius: undefined,
    borderTopRightRadius: undefined,
};

export const extensibleApisPropTypes = {
    ...spacingPropTypes,
    ...flexPropTypes,
    ...dimensionPropTypes,
    ...borderRadiusPropTypes,
    ...backgroundPropTypes,
};

export const propTypes = {
    ...extensibleApisPropTypes,
    /**
     * id for box
     */
    id: PropTypes.string,
    /**
     * dom node to render this Box as
     */
    is: PropTypes.node,
    children: PropTypes.node,
    /**
     * Is the box inline or block
     */
    inline: PropTypes.bool,
    /**
     * Test ID
     */
    testId: PropTypes.string,
};

export const defaultProps = {
    ...spacingDefaultProps,
    ...flexDefaultProps,
    ...dimensionDefaultProps,
    ...borderRadiusDefaultProps,
    ...backgroundDefaultProps,
    id: undefined,
    children: '',
    is: undefined,
    inline: false,
    testId: undefined,
};
