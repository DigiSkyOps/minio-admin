<template>
    <div id="edit-policy">
        <el-form :model="editPolicyData"
                 :rules="rules"
                 label-width="100px"
                 ref="editPolicy">
            <el-form-item :label="$t('policy_name')"
                          prop="name">
                <el-input style="width:300px"
                          @input="validation('editPolicy')"
                          v-model="editPolicyData.name">
                </el-input>
            </el-form-item>
            <el-form-item :label="$t('policy_detail')">
                <v-jsoneditor
                    v-model="editPolicyData.content"
                    :options="jsonOptions"
                    :plus="false"
                    height="400px">
                </v-jsoneditor>
            </el-form-item>
        </el-form>

        <div slot="footer"
             class="dialog-footer">
            <el-button :loading="isEditPolicy"
                       :disabled="!editPolicyValid"
                       type="primary"
                       @click="save">{{$t("confirm")}}
            </el-button>
            <el-button :disabled="isEditPolicy"
                       @click="resetData('editPolicy')">{{$t("reset")}}
            </el-button>
            <el-button :disabled="isEditPolicy"
                       @click="$emit('cancel')">{{$t("cancel")}}
            </el-button>
        </div>
    </div>
</template>

<script>
import { common } from '../mixin/common'
import { mapActions, mapGetters } from 'vuex'
import api from "../api/index"
export default {
    name: "editPolicy",
    mixins: [common],
    props: { editName: { type: String, default: "" }, bindedUser: { type: String, default: "" }, editDataContent: { type: Object, default: {} } },
    data() {
        return {
            rules: {},
            editPolicyValid: false,
            isEditPolicy: false,
            editPolicyData: {
                "name": "",
                "content": "",
            },
            jsonOptions: {
                mode: 'code',
                onValidate: json => {
                    var errors = [];
                    if (!json) {
                        this.editPolicyValid = false
                        errors.push({
                            path: ['custom'],
                            message: '格式错误'
                        });
                    } else {
                        if (this.validation('editPolicy')) {
                            this.editPolicyValid = true
                        } else {
                            this.editPolicyValid = false
                        }
                    }
                    return errors;
                },
                onChange: function () {
                    this.onValidate()
                },
            }
        }
    },
    mounted() {
        this.editPolicyData.name = this.editName
        this.editPolicyData.content = this.editDataContent
    },
    created() {
        this.rules = {
            name: [
                {
                    type: "string",
                    required: true,
                    message: this.$t("rule_policy_name"),
                    trigger: "change"
                }
            ],
        }
    },
    computed: {
        ...mapGetters({
            users: "getUsers"
        }),
    },
    methods: {
        async validation(ref) {
            try {
                await this.$refs[ref].validate()
                this.editPolicyValid = true
            } catch (e) {
                this.editPolicyValid = false
            }
        },
        // onError() {
        //     this.editPolicyValid = false;
        // },
        async save() {
            this.isEditPolicy = true;
            this.$emit("editingP");
            // 添加新的策略
            // 重新绑定用户
            // 找出原有用户
            // 重新绑定
            // 删除原有策略
            try {
                let data = {
                    policyname: this.editPolicyData.name,
                    policycontent: JSON.stringify(this.editPolicyData.content)
                };
                await api.ADD_POLICY(this, data);
                let changeData = {
                    "policyname": this.editPolicyData.name,
                    "accesskey":  this.bindedUser,
                    "isgroup": false
                }
                await api.CHANGE_POLICY(this,changeData)
                let delData = {
                    "name": this.editName
                }
                await api.DEL_POLICY(this,delData)
                this.handleSuccessMessage("修改", "策略");
                this.$emit("finish");
                this.resetData("editPolicy");
            } catch (e) {
                this.handleErrorMessage(e, "修改", "策略");
                this.$emit("error");
            } finally {
                this.isEditPolicy = false;
            }
        },
        resetData(ref) {
            // this.$refs[ref].resetFields()
            this.editPolicyData.content = {}
            this.editPolicyData.name = ''
            this.validation(ref);
        },
    }
}
</script>
