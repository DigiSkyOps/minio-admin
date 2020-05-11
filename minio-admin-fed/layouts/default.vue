<template>
    <a-layout id="minio-layout">
        <a-layout-sider :trigger="null"
                        collapsible
                        v-model="collapsed">
            <div class="logo" />
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
            <a-layout-header style="background: #fff; padding: 0">
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
            routes: [
                {
                    path: "/user",
                    name: "用户管理"
                },
                {
                    path: "/policy",
                    name: "策略管理",
                }
            ],
            selectedPath: [this.$route.path]
        };
    },
    async created() {
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
    }
};
</script>
<style>
#minio-layout .trigger {
    height: 100vh;
    font-size: 18px;
    line-height: 64px;
    padding: 0 24px;
    cursor: pointer;
    transition: color 0.3s;
}

#minio-layout .trigger:hover {
    color: #1890ff;
}

#minio-layout .logo {
    height: 32px;
    background: rgba(255, 255, 255, 0.2);
    margin: 16px;
}
</style>
