const mix = require('laravel-mix')
const webpack = require('webpack')

mix.postCss('_static/css/app.css', 'static/css/', [
  require('postcss-import')({
    from: '_static/css/app.css'
  }),
  require('cssnano')()
])
