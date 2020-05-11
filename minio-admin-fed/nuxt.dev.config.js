const merge = require('webpack-merge')
const baseConfig = require('./nuxt.base.config')

module.exports = merge(baseConfig,{
    axios: {
        //执行代理
        proxy: true,
        //请求带cookie
        credentials: true,
        //nuxt上下文中，将客户端请求标头设置为axios默认请求标头。这对于在服务器端进行需要基于Cookie的验证的请求很有用。也有助于在SSR和客户端代码中提出一致的请求。
        proxyHeaders: true,
        //在服务器端使用和预先创建请求的基本URL
        // baseURL: '',
        //在客户端使用和预先创建请求的基本URL
        // browserBaseURL: '',
        //自动拦截失败的请求，并重新请求次数,设置true默认3次,retries设置次数
        // retry:true,
        retry: { retries: 1 },
    },
    //设置代理
    proxy:  {
        '/api': {
            target: 'http://xxx.com',
            changeOrigin: true,
            pathRewrite: {
            '^/api' : '/api'
            }
        }
    },
})
