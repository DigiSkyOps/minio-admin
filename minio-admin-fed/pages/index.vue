<template>
    <div id="minio-app-index">
        <div class="app-header">
            <list-user />
        </div>
    </div>
</template>

<script>
export default {
    components: {
        "list-user": () => import("../components/listUser")
    },
    data() {
        return {
        };
    },
};
</script>

<style lang="less">
#minio-app-index {
    height: 100vh;
    justify-content: center;
    align-items: center;
}

.title {
    font-family: "Quicksand", "Source Sans Pro", -apple-system,
        BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial,
        sans-serif; /* 1 */
    display: block;
    font-weight: 300;
    font-size: 100px;
    color: #35495e;
    letter-spacing: 1px;
}

.subtitle {
    font-weight: 300;
    font-size: 42px;
    color: #526488;
    word-spacing: 5px;
    padding-bottom: 15px;
}

.links {
    padding-top: 15px;
}
</style>
