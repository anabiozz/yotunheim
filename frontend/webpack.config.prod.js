import path from 'path'

// const GLOBALS = {
//     'process.env.NODE_ENV': JSON.stringify('production')
// }


module.exports = {
    debug: false,
    mode: 'development',
    entry: ['./main.js'],
    output: {
        path: path.resolve(__dirname, '../backend/public'),
        filename: 'bundle.js'
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