'use strict';

const path = require('path');

module.exports = {
    resolve: {
        alias: {
            '7g-utils': path.resolve(__dirname, '/src/renderer/utils'),
            '7g-components': path.resolve(__dirname, '/src/renderer/components'),
            '7g-pages': path.resolve(__dirname, '/src/renderer/pages')
        },
        extensions: ['.ts', '.tsx', '.js', '.jsx']
    },
    module: {
        rules: [
            {
                test: /\.jsx$/,
                loader: 'babel-loader',
                exclude: /node_modules/,
            }
        ]
    }
};

// export default module;
