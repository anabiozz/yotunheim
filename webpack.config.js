var path = require('path');
var webpack = require('webpack');
var ManifestPlugin = require('webpack-manifest-plugin');

var ExtractTextPlugin = require("extract-text-webpack-plugin");

var autoprefixer = require('autoprefixer');
var CompressionPlugin = require("compression-webpack-plugin");

var host = process.env.HOST || 'localhost'
var devServerPort = 3808;

process.env.NODE_ENV = 'production'
var production = process.env.NODE_ENV === 'production';

const extractSass = new ExtractTextPlugin({
  filename: production ? "[name].[contenthash].css" : "[name].css",
});

var sassExtractor = () => {
  return ['css-hot-loader'].concat(extractSass.extract({
    use: [{
      loader: "css-loader",
      options: {
        sourceMap: true
      }
    }, {
      loader: 'postcss-loader',
      options: {
        sourceMap: true,
        plugins: [
          //require('postcss-import')({ root: loader.resourcePath }),
          //require('postcss-cssnext')(),
          autoprefixer({
            browsers:['ie >= 9', 'last 4 version', "> 1%"]
          })
          //require('cssnano')()
        ]
      }
    }, {
      loader: "sass-loader",
      options: {
        sourceMap: true
      }
    }],
    fallback: "style-loader"
  }))
}


var config = {
  entry: {
    vendor: [
      'babel-polyfill',
    ],
    application: 'application.es6'
  },

  module: {
      rules: [
          {
              exclude: [/node_modules/],
              include: [
                  path.resolve(__dirname, "frontend"),
              ],
              loader: 'react-hot-loader'
          },
          {
              test: /\.js$/,
              include: [
                  path.resolve(__dirname, "frontend"),
              ],
              loader: 'babel-loader',
              query: {
                  plugins: ['transform-runtime']
              }
          },
          { test: /\.(jpe?g|png|gif)$/i, use: "file-loader" },
          {
            test: /\.woff($|\?)|\.woff2($|\?)|\.ttf($|\?)|\.eot($|\?)|\.svg($|\?)/,
            use: production ? 'file-loader' : 'url-loader'
          },
          { test: /\.sass$/, use: sassExtractor() },
          { test: /\.scss$/, use: sassExtractor() },
          { test: /\.css$/, use: sassExtractor() }
      ]
  },

  output: {
    path: path.join(__dirname, 'backend/public', 'webpack'),
    publicPath: '/webpack/',

    filename: production ? '[name]-[chunkhash].js' : '[name].js'
  },

  resolve: {
    modules: [path.resolve(__dirname, "webpack"), path.resolve(__dirname, "node_modules")],
    extensions: [".js", ".es6", ".css", ".sass", ".scss"],
  },

  plugins: [
    extractSass,
    new ManifestPlugin({
      writeToFileEmit: true,
      basePath: "",
      publicPath: production ? "/webpack/" : 'http://' + host + ':' + devServerPort + '/webpack/',
    }),
  ]
};

if (production) {
  config.plugins.push(
    new webpack.optimize.CommonsChunkPlugin({name: 'vendor', filename: 'vendor-[chunkhash].js'}),
    new webpack.optimize.UglifyJsPlugin({
      compressor: { warnings: false },
      sourceMap: false
    }),
    new webpack.DefinePlugin({ // <--key to reduce React's size
      'process.env': { NODE_ENV: JSON.stringify('production') }
    }),
    new CompressionPlugin({
        asset: "[path].gz",
        algorithm: "gzip",
        test: /\.js$|\.css$/,
        threshold: 4096,
        minRatio: 0.8
    })
  );
} else {
  config.plugins.push(
    new webpack.optimize.CommonsChunkPlugin({name: 'vendor', filename: 'vendor.js'}),
    new webpack.NamedModulesPlugin()
  )

  config.devServer = {
    port: devServerPort,
    headers: { 'Access-Control-Allow-Origin': '*' },
  };
  config.output.publicPath = 'http://' + host + ':' + devServerPort + '/webpack/';
  config.devtool = 'source-map';
}

module.exports = config;