export const state = () => {
    return {
        miniousers: [],
        miniopolicys: [],
        errorLogs: [],
        buckets: [],
    }
}

export const actions = {
    async saveMinioUsers({ commit }, params) {
        commit('SAVE_MINIO_USERS', params)
    },
    async saveMinioPolicys({ commit }, params) {
        commit('SAVE_MINIO_POLICYS', params)
    },
    async saveErrorLogs({ commit }, params) {
        commit('SAVE_ERROR_LOGS', params)
    },
    async saveMinioBuckets({commit},params){
        commit('SAVE_MINIO_BUCKETS',params)
    }
}

export const mutations = {
    SAVE_MINIO_USERS(state, params) {
        state.miniousers = params;
    },
    SAVE_MINIO_POLICYS(state, params) {
        state.miniopolicys = params;
    },
    SAVE_ERROR_LOGS(state, params) {
        state.errorLogs = params;
    },
    SAVE_MINIO_BUCKETS(state,params){
        state.buckets = params
    }
}

export const getters = {
    "getPolicys": state => state.miniopolicys,
    "getUsers": state => state.miniousers,
    "getErrorLogs": state => state.errorLogs,
    "getBuckets": state => state.buckets,
}
