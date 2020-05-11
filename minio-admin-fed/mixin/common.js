import { mapActions,mapGetters } from "vuex"
import { Base64 } from "js-base64";
import api from "../api/index"
export const common = {
    data() {
        return {
            loading: null,
            props: {
                label: 'title',
                children: 'children'
            },
        }
    },
    computed: {
        ...mapGetters({
            users: "getUsers",
            policys: "getPolicys"
        })
    },
    methods: {
        ...mapActions({
            saveMinioUsers: "saveMinioUsers",
            saveMinioPolicys: "saveMinioPolicys",
            saveMinioBuckets: "saveMinioBuckets",
        }),
        handleShowLoading(text) {
            this.loading = this.$loading({
                // target: document.querySelector('.qin-view-container'),
                lock: true,
                text: text,
                spinner: 'el-icon-loading',
                background: 'rgba(0, 0, 0, 0.7)'
            })
        },
        handleCloseLoading() {
            setTimeout(() => {
                if (this.loading) {
                    this.loading.close()
                }
            }, 300)
        },
        handleShowSkeleton() {
            this.showSkeleton = true
        },
        handleCloseSkeleton() {
            setTimeout(() => {
                this.showSkeleton = false
            }, 300)
        },
        handleSuccessMessage(type, description) {
            this.$notification.open({
                // message: `${type}成功`,
                description: `${type}${description}成功`,
                icon: < a-icon type="check-circle"
                    style="color: green" />,
                duration: 1
            })
        },
        handleErrorMessage(e, type, name) {
            if (typeof e === 'boolean') {
                this.$notification.open({
                    message: '添加失败',
                    description: '请检查所填项',
                    icon: < a-icon type="close-circle"
                        style="color: red" />,
                    duration: 1
                })
            } else {
                this.$notification.open({
                    message: `${type}失败`,
                    description: `${type}${name}失败:${e.message}`,
                    icon: < a-icon type="close-circle"
                        style="color: red" />,
                    duration: 1
                })
            }
        },
        handleThrowError(e, message) {
            let error = new Error()
            error.data = e
            error.message = message
            throw error
        },
        handleCloseDrawer() {
            this.drawerVisible = false
            this.doing = false
            this.type = ''
        },
        findUser(policyName) {
            let arr = []
            for (var [k, v] of this.users.entries()) {
                if (v.policyName === policyName) {
                    arr.push(v.name)
                }
            }
            return arr
        },
        async getMinioPolicys() {
            try {
                let res = await api.GET_MINIOPOLICYS(this);
                // 转为数组
                let arr = Array.from(res.data);
                for (var [k, v] of Object.entries(res.data)) {
                    var o = {};
                    o["key"] = k;
                    o["name"] = k;
                    o["content"] = v;
                    o["bindeduser"] = this.findUser(k)
                    arr.push(o);
                }
                this.saveMinioPolicys(arr);
            } catch (e) {
                this.handleThrowError(e, "获取 minio 策略失败");
            }
        },
        async getMinioUsers() {
            try {
                let res = await api.GET_MINIOUSERS(this);
                // 转为数组
                var arr = [];
                for (var [k, v] of Object.entries(res.data)) {
                    var o = {};
                    o["key"] = k;
                    o["name"] = k;
                    o["policyName"] = v.policyName;
                    o["status"] = v.status;
                    arr.push(o);
                }
                this.saveMinioUsers(arr);
            } catch (e) {
                this.handleThrowError(e, "获取 minio 用户失败");
            }
        },
        async getMinioPolicys() {
            try {
                let res = await api.GET_MINIOPOLICYS(this);
                // 转为数组
                var arr = Array.from(res.data);
                for (var [k, v] of Object.entries(res.data)) {
                    var o = {};
                    o["key"] = k;
                    o["name"] = k;
                    o["content"] = v;
                    o["bindeduser"] = this.findUser(k)
                    arr.push(o);
                }
                this.saveMinioPolicys(arr);
            } catch (e) {
                this.handleThrowError(e, "获取 minio 策略失败");
            }
        },
        decode(content) {
            return JSON.parse(Base64.decode(content));
        },
        async getMinioBuckets() {
            try {
                let arr = []
                let res = await api.GET_MINIOBUCKETS(this);
                for (let v of res.data.values()) {
                    arr.push(v.name)
                }
                this.saveMinioBuckets(arr);
            } catch (e) {
                this.handleThrowError(e, "获取 minio 存储桶失败");
            }
        }
    },
}
