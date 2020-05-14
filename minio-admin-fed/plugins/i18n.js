import Vue from 'vue'
import VueI18n from 'vue-i18n'

Vue.use(VueI18n)

export default ({ app }) => {
    app.i18n = new VueI18n({
        locale: 'zh',
        fallbackLocale: 'zh',
        messages: {
            'zh': require('../locales/zh.json'),
            'en': require('../locales/en.json'),
        }
    })
}
