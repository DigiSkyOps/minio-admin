const GET_MINIOUSERS = (vm, data = '') => vm.$axios.get(`${process.env.platformUrl}/api/user/listuser`, data)
const ADD_USER = (vm, data = '') => vm.$axios.post(`${process.env.platformUrl}/api/user/adduser`, data)
const DEL_USER = (vm, data = '') => vm.$axios.post(`${process.env.platformUrl}/api/user/deluser`, data)
const CHANGE_USERPASSWORD = (vm, data = '') => vm.$axios.post(`${process.env.platformUrl}/api/user/setuser`, data)
const CHANGE_STATUS = (vm, data = '') => vm.$axios.post(`${process.env.platformUrl}/api/user/setuserstatus`, data)

const CHANGE_POLICY = (vm, data = '') => vm.$axios.post(`${process.env.platformUrl}/api/policy/setpolicy`, data)
const GET_MINIOPOLICYS = (vm, data = '') => vm.$axios.get(`${process.env.platformUrl}/api/policy/listpolicy`, data)
const ADD_POLICY = (vm, data = '') => vm.$axios.post(`${process.env.platformUrl}/api/policy/addpolicy`, data)
const DEL_POLICY = (vm, data = '') => vm.$axios.post(`${process.env.platformUrl}/api/policy/delpolicy`, data)


const GET_MINIOBUCKETS = (vm, data = '') => vm.$axios.get(`${process.env.platformUrl}/api/bucket/listbucket`, data)

export default {
    GET_MINIOUSERS,
    CHANGE_USERPASSWORD,
    CHANGE_STATUS,
    ADD_USER,
    DEL_USER,

    GET_MINIOPOLICYS,
    CHANGE_POLICY,
    ADD_POLICY,
    DEL_POLICY,

    GET_MINIOBUCKETS
}
