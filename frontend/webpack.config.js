const path = require('path')

module.exports = {
    mode: 'development',
    entry: ['./main.js'],
    output: {
        path: path.resolve(__dirname, '../backend/public'),
        filename: 'bundle.js'
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
                test: /\.(js)$/,
                exclude: /node_modules/,
                use: ['babel-loader'],
            },
        ]
    },
    performance : {
        hints : false
    }
}