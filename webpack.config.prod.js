import path from 'path'

// const GLOBALS = {
//     'process.env.NODE_ENV': JSON.stringify('production')
// }

module.exports = {
    mode: 'development',
    entry: ['./frontend/main.js'],
    output: {
        path: path.resolve(__dirname, './backend/public'),
        filename: 'bundle/bundle.js'
    },
    target: 'web',
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