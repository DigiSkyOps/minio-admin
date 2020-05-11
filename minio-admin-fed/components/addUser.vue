<template>
    <div id="change-password">
        <el-form :model="addUserData"
                 :rules="rules"
                 ref="addUser">
            <el-form-item label="用户名称"
                          prop="name">
                <el-input @input="validation('addUser')"
                          v-model="addUserData.name">
                </el-input>
            </el-form-item>
            <el-form-item label="密码"
                          prop="password">
                <el-input @input="validation('addUser')"
                          v-model="addUserData.password"
                          autocomplete="off">
                </el-input>
            </el-form-item>
        </el-form>
        <div class="dialog-footer">
            <el-button :disabled="!addUserValid"
                       type="primary"
                       @click="addUser">确 定
            </el-button>
            <el-button :disabled="isAddUser"
                       @click="resetData('addUser')">重置
            </el-button>
            <el-button :disabled="isAddUser"
                       @click="$emit('cancel')">取 消
            </el-button>
        </div>
    </div>
</template>
<script>
import { mapActions, mapGetters } from "vuex";
import { common } from "../mixin/common";
import api from "../api/index";
export default {
    name: "addUser",
    mixins: [common],
    data() {
        return {
            isAddUser: false,
            addUserValid: false,
            addUserData: {
                name: "",
                password: ""
            },
            rules: {
                name: [
                    {
                        type: "string",
                        required: true,
                        message: "请输入用户名",
                        trigger: "change"
                    }
                ],
                password: [
                    {
                        type: "string",
                        required: true,
                        message: "请输入用户密码",
                        trigger: "change"
                    }
                ]
            }
        };
    },
    async mounted() {
        this.validation("addUser");
    },
    async created() { },
    computed: {
        ...mapGetters({
            policys: "getMinioPolicys"
        })
    },
    methods: {
        ...mapActions({}),
        resetData(ref) {
            this.addUserData.password = "";
            this.addUserData.name = "";
            this.validation(ref);
        },
        async validation(ref) {
            try {
                await this.$refs[ref].validate();
                this.addUserValid = true;
            } catch (e) {
                this.addUserValid = false;
            }
        },
        async addUser() {
            this.isAddUser = true;
            this.$emit("addingU");
            try {
                let data = {
                    accessKey: this.addUserData.name,
                    secretKey: this.addUserData.password
                };
                await api.ADD_USER(this, data);
                this.handleSuccessMessage("添加", "用户");
                this.$emit("finish");
                this.resetData("addUser");
            } catch (e) {
                this.handleErrorMessage(e, "添加", "添加");
                this.$emit("error");
            } finally {
                this.isAddUser = false;
            }
        }
    }
};
</script>
