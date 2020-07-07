'use strict';

module.exports = {
    resolve: {
        alias: {
            '7g-utils': __dirname +  '/src/renderer/utils',
            '7g-components': __dirname + '/src/renderer/components',
            '7g-pages': __dirname + '/src/renderer/pages'
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
