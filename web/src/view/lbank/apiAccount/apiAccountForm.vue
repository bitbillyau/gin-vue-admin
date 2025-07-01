
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="序号:" prop="id">
    <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入序号" />
</el-form-item>
        <el-form-item label="邮箱:" prop="name">
    <el-input v-model="formData.name" :clearable="true" placeholder="请输入邮箱" />
</el-form-item>
        <el-form-item label="apiKey:" prop="apiKey">
    <el-input v-model="formData.apiKey" :clearable="true" placeholder="请输入apiKey" />
</el-form-item>
        <el-form-item label="apiSecret:" prop="apiSecret">
    <el-input v-model="formData.apiSecret" :clearable="true" placeholder="请输入apiSecret" />
</el-form-item>
        <el-form-item label="uid:" prop="uid">
    <el-input v-model="formData.uid" :clearable="true" placeholder="请输入uid" />
</el-form-item>
        <el-form-item label="passphrase:" prop="passphrase">
    <el-input v-model="formData.passphrase" :clearable="true" placeholder="请输入passphrase" />
</el-form-item>
         <el-form-item label="交易所:" prop="exchId">
              <el-select  v-model="formData.exchId" clearable placeholder="请选择交易所">
                <el-option
                  v-for="item in exchangeOptions"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
          </el-form-item>

        <el-form-item label="是否生效:" prop="status">
                <el-select   v-model="formData.status" clearable placeholder="请选择是否有效">
                  <el-option label="有效" :value="1" />
                </el-select>
              </el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createApiAccount,
  updateApiAccount,
  findApiAccount
} from '@/api/lbank/apiAccount'

defineOptions({
    name: 'ApiAccountForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            id: undefined,
            name: '',
            apiKey: '',
            apiSecret: '',
            uid: '',
            passphrase: '',
            exchId: undefined,
            status: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findApiAccount({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createApiAccount(formData.value)
               break
             case 'update':
               res = await updateApiAccount(formData.value)
               break
             default:
               res = await createApiAccount(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
