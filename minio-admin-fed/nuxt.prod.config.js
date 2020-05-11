const merge = require('webpack-merge')
const baseConfig = require('./nuxt.base.config')

module.exports = merge(baseConfig,{
    //打包的文件路径
    generate: {
        // dir: './dist/home',
        interval: 2,
        retry: { retries: 1 },
    },
    //打包构建静态文件输出路径
    build: {
        extend (config) {
            config.output.publicPath = `/static/`;
        },
        publicPath: 'static'
    },
    //访问路径
    // router: {
    //     base: '/m/'
    // },
    axios: {
        credentials: true,
        baseURL: '/'
    },
    //引入nuxt/component-cache 缓存组件提高性能
    // modules: [['@nuxtjs/component-cache', {
    //     max: 10000,
    //     maxAge: 1000 * 60 * 60
    // }]],
})
