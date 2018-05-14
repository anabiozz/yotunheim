const path = require('path')

module.exports = {
    mode: 'development',
    entry: ['./frontend/main.js'],
    output: {
        path: path.resolve(__dirname, './backend/public'),
        filename: 'bundle/bundle.js'
    },
    module: {
        rules: [
            {
                test: /\.html$/,
                use: [
                    {
                        loader: 'html-loader'
                    }
                ]
            },
            {
                test: /\.js$/,
                exclude: /node_modules/,
                loader: 'babel-loader'
            },
        ]
    },
    performance : {
        hints : false
    }
}