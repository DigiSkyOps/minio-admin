<template>
    <a-layout id="minio-layout">
        <a-layout-sider :trigger="null"
                        collapsible
                        v-model="collapsed">
            <div class="logo">
                <img src="../assets/images/minio.svg" alt="">
            </div>
            <a-menu theme="dark"
                    mode="inline"
                    :defaultSelectedKeys="selectedPath"
                    @click="handleClick">
                <template v-for="item in routes">
                    <a-menu-item :key="item.path">
                        <span>{{ item.name }}</span>
                    </a-menu-item>
                </template>
            </a-menu>
        </a-layout-sider>
        <a-layout>
            <a-layout-header style="background: #fff; padding: 0;overflow: hidden">
                <a-select :default-value="locale" style="width: 120px;float: right;margin: 15px;" @change="handleChange">
                    <a-select-option value="zh">
                        中文
                    </a-select-option>
                    <a-select-option value="en">
                        English
                    </a-select-option>
                </a-select>
            </a-layout-header>
            <a-layout-content id="minio-view-container"
                              :style="{ margin: '24px 16px', padding: '24px', background: '#fff', minHeight: '280px' }">
                <nuxt />
            </a-layout-content>
        </a-layout>
    </a-layout>
</template>
<script>
import { mapGetters, mapActions } from "vuex";
import { common } from "../mixin/common";
import api from "../api/index";
export default {
    mixins: [common],
    data() {
        return {
            collapsed: false,
            routes: [],
            selectedPath: [this.$route.path],
            locale: "zh"
        };
    },
    async created() {
        if(window.localStorage.getItem("minioAdminLocale")){
            this.$i18n.locale = window.localStorage.getItem("minioAdminLocale")
            this.locale = window.localStorage.getItem("minioAdminLocale")
        }
        this.routes = [
            {
                path: "/user",
                name: this.$t("menu_user")
            },
            {
                path: "/policy",
                name: this.$t("menu_policy"),
            }
        ]
        try {
            await this.getMinioUsers();
            await this.getMinioPolicys();
            await this.getMinioBuckets();
        } catch (e) {
            this.$notification.open({
                message: "初始化失败",
                description: `${e.message},${e.data.message}`,
                icon: <a-icon type="close-circle" style="color: red" />
            });
        }
    },
    methods: {
        ...mapActions({
            saveMinioUsers: "saveMinioUsers",
            saveMinioPolicys: "saveMinioPolicys",
            saveMinioBuckets: "saveMinioBuckets",
        }),
        handleClick(e) {
            this.selectedPath = [e.key];
            this.$router.push(e.key);
        },
        handleChange(data){
            window.localStorage.setItem("minioAdminLocale",data)
            this.locale = data
            this.$i18n.locale = data
            this.$router.go(0)
        }
    }
};
</script>
<style lang="less">
#minio-layout {
    .trigger:hover {
        color: #1890ff;
    }
    .logo{
        height: 66px;
        background: rgba(255, 255, 255, 0.2);
        margin: 16px;
        position: relative;
        text-align: center;
        padding: 5px 0;
        img{
            width: 30px;
        }
    }
    .trigger {
        height: 100vh;
        font-size: 18px;
        line-height: 64px;
        padding: 0 24px;
        cursor: pointer;
        transition: color 0.3s;
    }
}
</style>
