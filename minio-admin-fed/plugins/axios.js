import qs from 'qs'
import { Message, MessageBox } from 'element-ui';

export default function ({ $axios, headers, redirect,store }) {
    //设置headers
    // headers['Content-Type'] = 'application/json;charset=utf-8'
    $axios.onRequest((config) => {
        config.headers['Content-Type'] = 'application/json;charset=utf-8'
        // config.headers['Content-Type'] = 'application/x-www-form-urlencoded;charset=utf-8'
        // config.data = qs.stringify(config.data)
    })

    $axios.onResponse(response => {
        const data = response.data
        if(data.hasOwnProperty('code') && data.code !== 200){
            if(data.code === 401){
                Message.error("登陆失效，前往登陆。。。")
                setTimeout(()=>{
                    redirect('/')
                },1500)
            }else{
                let storeErrors = JSON.parse(JSON.stringify(store.state.errorLogs))
                storeErrors.unshift({
                    msg: data.message,
                    time: new Date().toLocaleString()
                })
                store.dispatch("saveErrorLogs",storeErrors)
                let e = new Error()
                e.response = {}
                e.response.status = data.code
                e.response.statusText = data.message
                e.message = data.message
                throw e
            }
        }else{
            return data
        }
    })

    $axios.onError(error => {
        let storeErrors = JSON.parse(JSON.stringify(store.state.errorLogs))
        storeErrors.unshift({
            msg: `${error.response.status}:${error.response.statusText}`,
            time: new Date().toLocaleString()
        })
        store.dispatch("saveErrorLogs",storeErrors)
    })
}
