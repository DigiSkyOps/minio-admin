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
            <a-popconfirm title="确定批量禁用?"
                          @confirm="() => patchDisable()">
                <a-button type="primary"
                          style="margin-left: 8px"
                          :disabled="!hasSelected"
                          :loading="loading">{{$t("batch_disable")}}
                </a-button>
            </a-popconfirm>
            <a-button type="primary"
                      style="margin-left: 8px"
                      @click="add"
                      :disabled="hasSelected">{{$t("add_user")}}
            </a-button>
        </div>
        <a-table :rowSelection="{selectedRowKeys: selectedRowKeys, onChange: onSelectChange}"
                 :columns="columns"
                 :dataSource="users"
                 bordered>
            <template v-for="col in ['name','policyName','status']"
                      :slot="col"
                      slot-scope="text, record">
                <div :key="col">
                    <a-input v-if="record.editable"
                             style="margin: -5px 0"
                             :value="text"
                             @change="e => handleChange(e.target.value, record.key, col)" />
                    <template v-else>{{text}}</template>
                </div>
            </template>
            <template slot="operation"
                      slot-scope="text, record">
                <div class="editable-row-operations">
                    <span>
                        <a-button type="primary"
                                  size="small"
                                  @click="() => edit(record,'changePolicy')">{{$t("edit")}}
                        </a-button>
                    </span>
                    <span>
                        <a-button type="primary"
                                  size="small"
                                  @click="() => changePassword(record,'changePassword')">{{$t("change_password")}}
                        </a-button>
                    </span>
                    <a-popconfirm v-if="isForbid(record)"
                                  title="Sure to disabled?"
                                  @confirm="() => onDisable(record)">
                        <a-button v-if="isForbid(record)"
                                  :disabled="!isForbid(record)"
                                  type="danger"
                                  size="small">{{$t("disable")}}
                        </a-button>
                    </a-popconfirm>
                    <a-popconfirm v-if="!isForbid(record)"
                                  title="Sure to active?"
                                  @confirm="() => onActive(record)">
                        <a-button v-if="!isForbid(record)"
                                  type="primary"
                                  size="small">{{$t("active")}}
                        </a-button>
                    </a-popconfirm>
                    <a-popconfirm v-if="users.length"
                                  title="Sure to delete?"
                                  @confirm="() => onDelete(record.key)">
                        <a-button type="danger"
                                  size="small">{{$t("delete")}}
                        </a-button>
                    </a-popconfirm>
                </div>
            </template>
        </a-table>
        <a-drawer placement="right"
                  :destroyOnClose="true"
                  :closable="false"
                  :maskClosable="!doing"
                  @close="handleCloseDrawer"
                  :visible="drawerVisible"
                  width="600"
                  :zIndex="99">
            <change-user-policy v-if="type === 'changePolicy'"
                                ref="changePolicy"
                                :editData="editData"
                                @editingU="doing = true"
                                @error="doing = false"
                                @finish="handleFinishEditUser"
                                @cancel="handleCloseDrawer" />
            <change-password v-if="type === 'changePassword'"
                             ref="changePassword"
                             :changePasswordData="changePasswordData"
                             @changingU="doing = true"
                             @error="doing = false"
                             @finish="handleFinishChangePassword"
                             @cancel="handleCloseDrawer" />
            <add-user v-if="type === 'addUser'"
                      ref="addUser"
                      @addingU="doing = true"
                      @error="doing = false"
                      @finish="handleFinishaddUser"
                      @cancel="handleCloseDrawer" />
        </a-drawer>
    </div>
</template>
<script>
import { mapGetters } from "vuex";
import { common } from "../mixin/common";
import api from "../api/index";
export default {
    components: {
        "change-password": () => import("./changePassword"),
        "change-user-policy": () => import("./changeUserPolicy"),
        "add-user": () => import("./addUser")
    },
    mixins: [common],
    data() {
        return {
            jsonData: {
                total: 25,
                limit: 10,
                skip: 0,
                links: {
                    previous: undefined,
                    next: "next"
                }
            },
            columns: [],
            drawerVisible: false,
            doing: false,
            editingKey: "",
            type: "",
            selectedRowKeys: [], // Check here to configure the default column
            loading: false,
            editData: {},
            changePasswordData: {}
        };
    },
    created() {
        this.columns = [{
            title: this.$t("username"),
            dataIndex: "name",
            width: "25%",
            scopedSlots: { customRender: "name" },
            sorter: (a, b) => a.name.length - b.name.length,
            sortDirections: ["descend"]
        },
        {
            title: this.$t("binding_policy"),
            dataIndex: "policyName",
            width: "25%",
            scopedSlots: { customRender: "policyName" }
        },
        {
            title: this.$t("state"),
            dataIndex: "status",
            width: "15%",
            scopedSlots: { customRender: "status" }
        },
        {
            title: this.$t("operation"),
            width: "35%",
            dataIndex: "operation",
            scopedSlots: { customRender: "operation" }
        }]
    },
    computed: {
        ...mapGetters({
            users: "getUsers"
        }),
        hasSelected() {
            return this.selectedRowKeys.length > 0;
        }
    },
    methods: {
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
        async handleFinishEditUser() {
            this.handleCloseDrawer();
            this.getMinioUsers();
        },
        async handleFinishChangePassword() {
            this.handleCloseDrawer();
            this.getMinioUsers();
        },
        async handleFinishaddUser() {
            this.handleCloseDrawer();
            this.getMinioUsers();
        },
        edit(record, type) {
            this.drawerVisible = true;
            this.type = type;
            this.editData = record;
        },
        add() {
            this.drawerVisible = true;
            this.type = "addUser";
        },
        async patchDelete() {
            try {
                for (let v of this.selectedRowKeys.values()) {
                    let data = {
                        accessKey: v
                    };
                    await api.DEL_USER(this, data);
                }
                this.handleSuccessMessage("删除", "用户");
            } catch (e) {
                this.handleErrorMessage(e, "删除", "用户");
            } finally {
                this.selectedRowKeys = [];
                this.getMinioUsers();
            }
        },
        async patchDisable() {
            try {
                for (let v of this.selectedRowKeys.values()) {
                    let data = {
                        accessKey: v,
                        status: "disabled"
                    };
                    await api.CHANGE_STATUS(this, data);
                }
                this.handleSuccessMessage("批量禁用", "用户");
            } catch (e) {
                this.handleErrorMessage(e, "批量禁用", "用户");
            } finally {
                this.selectedRowKeys = [];
                this.getMinioUsers();
            }
        },
        changePassword(record, type) {
            this.drawerVisible = true;
            this.type = type;
            this.changePasswordData = record;
        },
        async onDisable(record) {
            try {
                let data = {
                    accessKey: record.name,
                    status: "disabled"
                };
                await api.CHANGE_STATUS(this, data);
                this.handleSuccessMessage("禁用", "用户");
            } catch (e) {
                this.handleErrorMessage(e, "禁用", "用户");
            } finally {
                this.getMinioUsers();
            }
        },
        async onActive(record) {
            try {
                let data = {
                    accessKey: record.name,
                    status: "enabled"
                };
                await api.CHANGE_STATUS(this, data);
                this.handleSuccessMessage("激活", "用户");
            } catch (e) {
                this.handleErrorMessage(e, "激活", "用户");
            } finally {
                this.getMinioUsers();
            }
        },
        async onDelete(key) {
            try {
                let data = {
                    accessKey: key
                };
                await api.DEL_USER(this, data);
                this.handleSuccessMessage("删除", "用户");
            } catch (e) {
                this.handleErrorMessage(e, "删除", "用户");
            } finally {
                this.getMinioUsers();
            }
        },
        save(key) { },
        cancel(key) { },
        async start() {
            this.loading = true;
            // ajax request after empty completing
            setTimeout(() => {
                this.loading = false;
                this.selectedRowKeys = [];
            }, 1000);
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
