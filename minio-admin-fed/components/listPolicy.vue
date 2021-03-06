<template>
    <div>
        <div style="margin-bottom: 16px">
            <a-popconfirm title="确定批量删除?"
                          @confirm="() => patchDelete()">
                <a-button type="primary"
                          style="margin-left: 8px"
                          :disabled="!hasSelected"
                          :loading="loading">{{$t("batch_delete")}}
                </a-button>
            </a-popconfirm>
            <a-button type="primary"
                      style="margin-left: 8px"
                      @click="add"
                      :disabled="hasSelected">{{$t("add_policy")}}
            </a-button>
        </div>
        <a-table :rowSelection="{selectedRowKeys: selectedRowKeys, onChange: onSelectChange}"
                 :columns="columns"
                 style="height: 77vh;overflow-y: scroll"
                 :dataSource="policys"
                 bordered>
            <template slot="name"
                      slot-scope="text, record, index">
                {{text}}
            </template>
            <template slot="bindeduser"
                      slot-scope="text, record, index">
                <el-tag v-for="tag in record.bindeduser"
                        color="info"
                        size="medium"
                        :key="tag">
                    {{tag}}
                </el-tag>
            </template>
            <template slot="operation"
                      slot-scope="text, record, index">
                <div class="editable-row-operations">
                    <span>
                        <a-button type="primary"
                                style="font-family: Arial;"
                                size="small"
                                @click="() => detail(record)">{{$t("detail")}}
                        </a-button>
                    </span>
                    <span>
                        <a-button type="primary"
                                  size="small"
                                  @click="() => edit(record,text,'editPolicy')">{{$t("edit_policy")}}
                        </a-button>
                    </span>
                    <a-popconfirm v-if="users.length"
                                  title="Sure to delete?"
                                  @confirm="() => onDelete(record.key)">
                        <a-button type="danger"
                                  size="small">{{$t("delete_policy")}}
                        </a-button>
                    </a-popconfirm>
                </div>
            </template>
        </a-table>
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
        <a-drawer placement="right"
                  :destroyOnClose="true"
                  :closable="false"
                  :maskClosable="!doing"
                  @close="handleCloseDrawer"
                  :visible="drawerVisible"
                  :width="drawerWidth"
                  :zIndex="99">
            <edit-policy v-show="type === 'editPolicy'"
                         ref="editPolicy"
                         :editDataContent="editDataContent"
                         :editName="editName"
                         :bindedUser="bindedUser"
                         @editingP="doing = true"
                         @error="doing = false"
                         @finish="handleFinishEditPolicy"
                         @cancel="handleCloseDrawer" />
            <add-policy v-if="type === 'addPolicy'"
                        ref="addPolicy"
                        @addingP="doing = true"
                        @error="doing = false"
                        @finish="handleFinishaddPolicy"
                        @cancel="handleCloseDrawer" />
        </a-drawer>
    </div>
</template>
<script>
import { mapGetters } from "vuex";
import { Base64 } from "js-base64";
import { common } from "../mixin/common";
import api from "../api/index";
export default {
    components: {
        "add-policy": () => import("./addPolicy"),
        "edit-policy": () => import("./editPolicy")
    },
    mixins: [common],
    created() {
        this.columns = [
            {
                title: this.$t("policy_name"),
                dataIndex: "name",
                width: "20%",
                scopedSlots: { customRender: "name" },
                sorter: (a, b) => a.name.length - b.name.length,
                sortDirections: ["descend"]
            },
            {
                title: this.$t("username"),
                dataIndex: "bindeduser",
                scopedSlots: { customRender: "bindeduser" }
            },
            {
                title: this.$t("operation"),
                dataIndex: "operation",
                width: "35%",
                scopedSlots: { customRender: "operation" }
            }
        ]
    },
    data() {
        return {
            columns: [],
            detailContent: 'eyJWZXJzaW9uIjoiMjAxMi0xMC0xNyIsIlN0YXRlbWVudCI6W3siRWZmZWN0IjoiQWxsb3ciLCJBY3Rpb24iOlsiczM6KiJdLCJSZXNvdXJjZSI6WyJhcm46YXdzOnMzOjo6bnh4LW4vKiJdfV19',
            detailVisible: false,
            drawerVisible: false,
            drawerWidth: 600,
            doing: false,
            type: "",
            bindedUser: "",
            selectedRowKeys: [], // Check here to configure the default column
            loading: false,
            editName: "",
            editDataContent: {}
        };
    },
    computed: {
        ...mapGetters({
            users: "getUsers",
            policys: "getPolicys"
        }),
        hasSelected() {
            return this.selectedRowKeys.length > 0;
        },
    },
    methods: {
        edit(record, text, type) {
            this.drawerWidth = 600
            this.drawerVisible = true;
            this.type = type;
            this.editName = record.name
            this.bindedUser = record.bindeduser[0]
            this.editDataContent = JSON.parse(Base64.decode(record.content));
        },
        handleChange(value, key, column) {
            const newData = [...this.users];
            const target = newData.filter(item => key === item.key)[0];
            if (target) {
                target[column] = value;
                this.users = newData;
            }
        },
        isForbid(record) {
            if (record.status === "enabled") {
                return true;
            } else {
                return false;
            }
        },
        hideModal() {
            this.visible = false;
        },
        detail(obj) {
            this.detailVisible = true
            this.detailContent = obj.content
        },
        async handleFinishEditPolicy() {
            this.handleCloseDrawer();
            await this.getMinioUsers()
            await this.getMinioPolicys();
        },
        async handleFinishaddPolicy() {
            this.handleCloseDrawer();
            await this.getMinioPolicys();
        },
        add() {
            this.drawerWidth = 1000;
            this.drawerVisible = true;
            this.type = "addPolicy";
        },
        async patchDelete() {
            try {
                for (let v of this.selectedRowKeys.values()) {
                    let data = {
                        "name": v
                    };
                    await api.DEL_POLICY(this, data);
                }
                this.handleSuccessMessage("删除", "策略");
            } catch (e) {
                this.handleErrorMessage(e, "删除", "策略");
            } finally {
                this.selectedRowKeys = [];
                this.getMinioPolicys()
                this.getMinioUsers();
            }
        },
        async onDelete(key) {
            try {
                let data = {
                    name: key
                };
                await api.DEL_POLICY(this, data);
                this.handleSuccessMessage("删除", "策略");
            } catch (e) {
                this.handleErrorMessage(e, "删除", "策略");
            } finally {
                this.getMinioPolicys()
            }
        },
        onSelectChange(selectedRowKeys) {
            //   console.log("selectedRowKeys changed: ", selectedRowKeys);
            this.selectedRowKeys = selectedRowKeys;
        }
    }
};
</script>
<style scoped>
.editable-row-operations a {
    margin-right: 8px;
}
</style>
