<template>
    <el-row :gutter="10">
        <el-col :span="span">
            <el-card class="form">
                <el-form :model="addPolicyData"
                         :rules="rules"
                         :label-position="labelPosition"
                         label-width="100px"
                         ref="addPolicy">
                    <el-form-item label="策略名称"
                                  prop="name">
                        <el-input style="width:300px"
                                  @input="validation('addPolicy')"
                                  v-model="addPolicyData.name">
                        </el-input>
                    </el-form-item>
                    <el-form-item label="自定义编写">
                        <el-switch v-model="customEdit"
                                   @change="changeEdit"></el-switch>
                    </el-form-item>
                    <el-form-item v-if="!customEdit">
                        <el-button @click="addRule">
                            <span class="darken"><i class="iconfont icon-tianjia"></i></span>
                            <span>Add A Rule</span>
                        </el-button>
                    </el-form-item>

                    <div v-if="!customEdit"
                         v-for="(item,index) in addPolicyData.rules"
                         :key="index"
                         class="rule-item">
                        <el-button @click="delRule(index)">
                            <span class="darken"><i class="iconfont icon-jian"></i></span>
                        </el-button>
                        <el-form :model="item"
                                 :rules="rules"
                                 v-if="!customEdit"
                                 :label-position="labelPosition"
                                 label-width="80px"
                                 :ref="'rule'+ index">
                            <el-form-item prop="bucket"
                                          label="存储桶">
                                <el-select v-model="item.bucket"
                                           @change="validation('addPolicy')"
                                           placeholder="请选择存储桶">
                                    <el-option v-for="item in buckets"
                                               :key="item"
                                               :label="item"
                                               :value="item">
                                    </el-option>
                                </el-select>
                            </el-form-item>
                            <el-form-item label="路径"
                                          prop="path">
                                <el-input style="width:300px"
                                          @input="validation('addPolicy')"
                                          v-model="item.path">
                                </el-input>
                            </el-form-item>
                            <el-form-item prop="action"
                                          label="Action">
                                <el-radio-group v-model="item.action">
                                    <el-radio label="readonly"></el-radio>
                                    <el-radio label="writeonly"></el-radio>
                                    <el-radio label="readwrite"></el-radio>
                                </el-radio-group>
                            </el-form-item>
                            <el-form-item prop="effect"
                                          label="Effect">
                                <el-radio-group v-model="item.effect">
                                    <el-radio label="Allow"></el-radio>
                                    <el-radio label="Deny"></el-radio>
                                </el-radio-group>
                            </el-form-item>
                        </el-form>
                    </div>

                    <el-form-item v-if="customEdit">
                        <template>
                            <v-jsoneditor :value="content"
                                          ref="jsoneditor"
                                          :options="jsonOptions"
                                          :plus="false"
                                          height="400px">
                            </v-jsoneditor>
                        </template>
                    </el-form-item>
                </el-form>
                <div class="dialog-footer">
                    <el-button :disabled="!addPolicyValid"
                               type="primary"
                               @click="addPolicy">确 定
                    </el-button>
                    <el-button :disabled="isAddPolicy"
                               @click="resetData('addPolicy')">重置
                    </el-button>
                    <el-button :disabled="isAddPolicy"
                               @click="$emit('cancel')">取 消
                    </el-button>
                </div>
            </el-card>
        </el-col>
        <el-col :span="span"
                v-if="!customEdit">
            <el-card class="box-card">
                <div slot="header"
                     class="clearfix">
                    <span>策略内容</span>
                </div>
                <pre class="json">{{ jsonString }}</pre>
            </el-card>
        </el-col>
    </el-row>
</template>
<script>
import { mapActions, mapGetters } from "vuex";
import { common } from "../mixin/common";
import api from "../api/index";
export default {
    name: "addPolicy",
    mixins: [common],
    data() {
        return {
            labelPosition: "right",
            isAddPolicy: false,
            customEdit: false,
            addPolicyValid: false,
            content: {},
            addPolicyData: {
                name: `polilcy-${this.formatDate(new Date())}`,
                content: {},
                // bucket: "",
                // path: "*",
                // action: "readonly",
                // effect: "Allow",
                rules: []
            },
            rules: {
                name: [
                    {
                        type: "string",
                        required: true,
                        message: "请输入策略名称",
                        trigger: "change"
                    }
                ],
                content: [
                    {
                        required: true,
                        message: "请输入策略内容",
                        trigger: "change"
                    }
                ],
                bucket: [
                    {
                        required: true,
                        message: "请选择存储桶",
                        trigger: "change"
                    }
                ],
                path: [
                    {
                        type: "string",
                        required: true,
                        message: "请输入路径",
                        trigger: "change"
                    }
                ],
            },
            jsonOptions: {
                mode: 'code',
                onValidate: async json => {
                    var errors = [];
                    if (!json) {
                        errors.push({
                            path: ['custom'],
                            message: '格式错误'
                        })
                        this.addPolicyValid = falses
                        throw Error
                    } else {
                        try {
                            this.addPolicyData.content = json
                            await this.validation("addPolicy")
                        } catch (e) {
                            this.addPolicyValid = false
                        }
                    }
                    return errors;
                },
                onChange: async function () {
                    this.onValidate()
                },
                // onValidationError: async errors => {
                //     console.log('json err')
                //     this.addPolicyValid = false
                // }
            }
        }
    },
    async created() { },
    computed: {
        ...mapGetters({
            policys: "getPolicys",
            buckets: "getBuckets"
        }),
        jsonString() {
            return JSON.stringify(this.model, null, 2).trim();
        },
        span() {
            if (!this.customEdit) {
                return 12
            } else {
                return 24
            }
        },
        Statement() {
            let sm = []
            for (let v of this.addPolicyData.rules.values()) {
                let action
                switch (v.action) {
                    case 'readonly':
                        action = ["s3:GetBucketLocation", "s3:GetObject", "s3:ListBucket"]
                        break;
                    case 'writeonly':
                        action = ["s3:PutObject"]
                        break;
                    case 'readwrite':
                        action = ["s3:*"]
                        break;
                }
                let path = v.path
                if (v.path[0] === "/"){
                    path = path.replace(/\//i,'')
                }
                let data = {
                    "Action": action,
                    "Effect": v.effect,
                    "Resource": `arn:aws:s3:::${v.bucket}/${path}`
                };
                sm.push(data)
            }
            return sm
        },
        model() {
            let data = {
                "Version": "2012-10-17",
                "Statement": this.Statement
            }
            return data
        }
    },
    methods: {
        ...mapActions({}),
        async resetData(ref) {
            if (this.customEdit) {
                this.$refs['addPolicy'].resetFields()
                this.$refs['jsoneditor'].editor.setText('')
                this.addPolicyValid = false
            } else {
                this.addPolicyData.name = ''
                this.addPolicyData.rules = []
                await this.validation(ref)
            }
        },
        formatDate(date) {
            var y = date.getFullYear()
            var m = date.getMonth() + 1
            m = m < 10 ? ('0' + m) : m
            var d = date.getDate()
            d = d < 10 ? ('0' + d) : d
            var h = date.getHours()
            var minute = date.getMinutes()
            minute = minute < 10 ? ('0' + minute) : minute
            return y + '-' + m + '-' + d + '-' + h + '-' + minute
        },
        async addRule() {
            let rule = {
                "bucket": "",
                "path": "*",
                "action": "readonly",
                "effect": "Allow"
            }
            this.addPolicyData.rules.push(rule)
            this.validateRule()
        },
        async delRule(index) {
            eval(`this.addPolicyData.rules.splice(index,1)`)
            await this.validation("addPolicy")
            await this.validateRule()
        },
        async changeEdit() {
            await this.validation("addPolicy")
            if (this.customEdit) {
                this.$refs['jsoneditor'].editor.set(this.addPolicyData.content)
            }
        },
        async validation(ref) {
            if (this.customEdit) {
                try {
                    await this.$refs[ref].validate()
                    this.addPolicyValid = true
                } catch (e) {
                    this.addPolicyValid = false
                }
            } else {
                try {
                    await this.$refs[ref].validate()
                    await this.validateRule()

                } catch (e) {
                    this.addPolicyValid = false
                }
            }
        },
        async validateRule() {
            this.addPolicyData.rules.length == 0 ? this.addPolicyValid = false : this.addPolicyValid = true
            try {
                for (let index of this.addPolicyData.rules.keys()) {
                    let ref = "rule" + index
                    await this.$refs[ref][0].validate()
                    this.addPolicyValid = true
                }
            } catch (e) {
                this.addPolicyValid = false
            }
        },
        async addPolicyByForm() {
            try {
                let data = {
                    "policyname": this.addPolicyData.name,
                    "policycontent": JSON.stringify(this.model)
                }
                await api.ADD_POLICY(this, data)
                this.handleSuccessMessage("添加", "策略")
                this.$emit("finish")
                this.resetData("addPolicy")
            } catch (e) {
                this.handleErrorMessage(e, "添加", "策略")
                this.$emit("error")
            } finally {
                this.isAddPolicy = false
            }
        },
        async addPolicyByCustom() {
            try {
                let data = {
                    "policyname": this.addPolicyData.name,
                    "policycontent": JSON.stringify(this.addPolicyData.content)
                };
                let res = await api.ADD_POLICY(this, data)
                this.handleSuccessMessage("添加", "策略")
                this.$emit("finish")
                this.resetData("addPolicy")
            } catch (e) {
                this.handleErrorMessage(e, "添加", "策略")
                this.$emit("error")
            } finally {
                this.isAddPolicy = false
            }
        },
        async addPolicy() {
            this.isAddPolicy = true
            this.$emit("addingP")
            if (this.customEdit) {
                this.addPolicyByCustom()
            } else {
                this.addPolicyByForm()
            }
        }
    }
};
</script>
<style>
.json {
    text-align: left;
}
.rule-item {
    background: #f2f5f5;
    padding: 20px;
    margin: 10px;
    border: 1px solid #dde4e6;
}
</style>
