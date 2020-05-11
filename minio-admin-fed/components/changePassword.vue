<template>
    <div id="change-password">
        <el-form :model="editUserData"
                 :rules="rules"
                 ref="changePassword">
            <el-form-item label="用户名称"
                          prop="name">
                <el-input v-model="editUserData.name"
                          :disabled="true">
                </el-input>
            </el-form-item>
            <el-form-item label="密码"
                          prop="password">
                <el-input @input="validation('changePassword')"
                          v-model="editUserData.password"
                          autocomplete="off">
                </el-input>
            </el-form-item>
        </el-form>
        <div class="dialog-footer">
            <el-button :loading="isChangePassword"
                       :disabled="!changePasswordValid"
                       type="primary"
                       @click="saveUser">确 定
            </el-button>
            <el-button :disabled="isChangePassword"
                       @click="resetData('changePassword')">重置
            </el-button>
            <el-button :disabled="isChangePassword"
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
    name: "changePassword",
    mixins: [common],
    props: { changePasswordData: { type: Object, default: {} } },
    data() {
        return {
            isChangePassword: false,
            changePasswordValid: false,
            editUserData: {
                name: "",
                status: "",
                policyName: "",
                password: ""
            },
            rules: {
                password: [
                    {
                        type: "string",
                        required: true,
                        message: "请输入用户新密码",
                        trigger: "change"
                    }
                ]
            }
        };
    },
    async mounted() {
        this.editUserData.name = this.changePasswordData.name;
        this.editUserData.status = this.changePasswordData.status;
        this.editUserData.policyName = this.changePasswordData.policyName;
        this.validation("changePassword");
    },
    async created() { },
    computed: {
        ...mapGetters({
            policys: "getPolicys"
        })
    },
    methods: {
        ...mapActions({}),
        resetData(ref) {
            this.editUserData.password = "";
            this.validation(ref);
        },
        async validation(ref) {
            try {
                await this.$refs[ref].validate();
                this.changePasswordValid = true;
            } catch (e) {
                this.changePasswordValid = false;
            }
        },
        async saveUser() {
            this.isChangePassword = true;
            this.$emit("changingU");
            try {
                let data = {
                    accessKey: this.editUserData.name,
                    secretKey: this.editUserData.password,
                    status: "enabled"
                };
                await api.CHANGE_USERPASSWORD(this, data);
                this.handleSuccessMessage("更新", "密码");
                this.$emit("finish");
                this.resetData("changePassword");
            } catch (e) {
                this.handleErrorMessage(e, "更新", "密码");
                this.$emit("error");
            } finally {
                this.isChangePassword = false;
            }
        }
    }
};
</script>
