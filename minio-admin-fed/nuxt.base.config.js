module.exports = {
    cache: true,
    loading: {
        color: '#409EFF',
        failedColor: 'red'
    },
    //配置路由权鉴
    // router: {
    // middleware: 'auth'
    // },
    // 配置环境变量
    env: {
        platformUrl: process.env.PLATFORM_URL || 'http://localhost:9000'
    },
    head: {
        title: '',
        meta: [
            {
                charset: 'utf-8'
            }, {
                name: 'viewport',
                content: 'width=device-width, initial-scale=1'
            }, {
                hid: 'description',
                name: 'description',
                content: 'Nuxt.js project'
            }
        ],
        //配置icon
        link: [
            {
                rel: 'icon',
                type: 'image/x-icon',
                href: '/favicon.ico'
            }
        ],
        script: [
            { src: 'https://cdn.bootcss.com/jquery/3.3.1/jquery.min.js' },
            { innerHTML: 'if (typeof module === \'object\') {window.jQuery = window.$ = module.exports;};', type: 'text/javascript' }
        ]
    },
    //配置后台渲染或者spa模式
    mode: 'spa',
    //引入nuxt/axios插件 proxy设置代理
    modules: ['@nuxtjs/axios'],
    css: [
        '~/assets/css/main.css',
        'element-ui/lib/theme-chalk/index.css',
        '~/static/iconfont.css'],
    build: {
        //依赖的第三方模块
        vendor: ['axios', 'element-ui']
    },
    //vue使用的第三方库 vue use之类的
    plugins: [
        { src: '~/plugins/axios', ssr: false },
        { src: '~/plugins/element-ui', ssr: true },
        { src: '~/plugins/localStorage', ssr: false },
        { src: '~/plugins/ant-ui', ssr: true },
        { src: '~/plugins/json-view', ssr: true },
        { src: '~/plugins/v-jsoneditor', ssr: true },
        { src: '~/plugins/i18n', ssr: true },
    ],
}
