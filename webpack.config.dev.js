const path = require('path')

module.exports = {
    mode: 'development',
    entry: ['./frontend/main.js'],
    output: {
        path: path.resolve(__dirname, './backend/public/bundle'),
        filename: 'bundle.js'
    },
    module: {
        rules: [
            {
                test: /\.html$/,
                use: {
                    loader: 'html-loader'
                }
            },
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: {
                    loader: 'babel-loader'
                }
            },
            {
                test: /\.s?css$/,
                use: [
                  'style-loader',
                  'css-loader',
                  'sass-loader'
                ]
            }
        ]
    },
    performance : {
        hints : false
    }
}