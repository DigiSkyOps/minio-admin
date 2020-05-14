<template>
    <div id="change-policy">
        <el-form :model="editUserData"
                 :rules="rules"
                 ref="changePolicy">
            <el-form-item :label="$t('username')"
                          prop="name">
                <el-input v-model="editUserData.name"
                          :disabled="true">
                </el-input>
            </el-form-item>
            <el-form-item :label="$t('binding_policy')"
                          style="height: 77vh;overflow-y: scroll">
                <a-table :columns="columns"
                         :dataSource="policys"
                         :rowSelection="{
                            selectedRowKeys: selectedRowKeys,
                            onChange: onSelectChange,
                            type: 'radio'
                        }"
                         bordered>
                    <span slot="name"
                          slot-scope="text">{{ text }}
                    </span>
                    <template slot="content"
                              slot-scope="text">
                        <span>
                            <a-button type="primary"
                                      style="font-family: Arial;"
                                      size="small"
                                      @click="() => detail(text)">{{$t("detail")}}
                            </a-button>
                        </span>
                    </template>
                </a-table>
            </el-form-item>
        </el-form>
        <div class="dialog-footer">
            <el-button :loading="isChangePolicy"
                       :disabled="!editUserValid"
                       type="primary"
                       @click="changePolicy">{{$t("confirm")}}
            </el-button>
            <el-button :disabled="isChangePolicy"
                       @click="resetData">{{$t("reset")}}
            </el-button>
            <el-button :disabled="isChangePolicy"
                       @click="$emit('cancel')">{{$t("cancel")}}
            </el-button>
        </div>
        <div>
            <a-modal :title="$t('policy_detail')"
                     v-model="detailVisible"
                     :footer="null">
                <json-viewer :value="decode(detailContent)"
                             :expand-depth="5"
                             copyable
                             sort>
                </json-viewer>
            </a-modal>
        </div>
    </div>
</template>
<script>
import { mapActions, mapGetters } from "vuex";
import { common } from "../mixin/common";
import api from "../api/index";
export default {
    name: "changePolicy",
    mixins: [common],
    props: { editData: { type: Object, default: {} } },
    data() {
        return {
            columns: [],
            selectedRowKeys: [],
            isChangePolicy: false,
            editUserValid: false,
            detailVisible: false,
            detailContent: 'eyJWZXJzaW9uIjoiMjAxMi0xMC0xNyIsIlN0YXRlbWVudCI6W3siRWZmZWN0IjoiQWxsb3ciLCJBY3Rpb24iOlsiczM6KiJdLCJSZXNvdXJjZSI6WyJhcm46YXdzOnMzOjo6bnh4LW4vKiJdfV19',
            editUserData: {
                name: "",
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
        this.columns = [
            {
                title: this.$t("policy_name"),
                dataIndex: "name",
                key: "name",
                slots: { title: "customTitle" },
                scopedSlots: { customRender: "name" }
            },
            {
                title: this.$t("operation"),
                dataIndex: "content",
                scopedSlots: { customRender: "content" },
                key: "content"
            }
        ]
        this.editUserData.name = this.editData.name;
        this.editUserData.status = this.editData.status;
        this.editUserData.policyName = this.editData.policyName;

        this.selectedRowKeys.push(this.editUserData.policyName);
        this.validation("changePolicy");
    },
    async created() { },
    computed: {
        ...mapGetters({
            policys: "getPolicys"
        })
    },
    methods: {
        ...mapActions({}),
        onSelectChange(selectedRowKeys) {
            // console.log("selectedRowKeys changed: ", selectedRowKeys);
            this.selectedRowKeys = selectedRowKeys;
        },
        resetData() {
            this.editUserData.selectedRowKeys = [];
        },
        detail(text) {
            this.detailVisible = true
            this.detailContent = text
        },
        async validation(ref) {
            try {
                await this.$refs[ref].validate();
                this.editUserValid = true;
            } catch (e) {
                this.editUserValid = false;
            }
        },
        async changePolicy() {
            this.isChangePolicy = true;
            this.$emit("editingU");
            try {
                let data = {
                    accessKey: this.editUserData.name,
                    policyname: this.selectedRowKeys[0],
                    isgroup: false
                };
                await api.CHANGE_POLICY(this, data);
                this.handleSuccessMessage("绑定", "策略");
                this.$emit("finish");
                this.resetData();
            } catch (e) {
                this.handleErrorMessage(e, "绑定", "策略");
                this.$emit("error");
            } finally {
                this.isChangePolicy = false;
            }
        }
    }
};
</script>
