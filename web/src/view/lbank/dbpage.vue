<template>
  <div>
    <el-card header="同步表" style="width: 100%">
      <el-button 
        v-if="btnAuth.syncServerStrategySymbolBtn" 
        type="primary" 
        :loading="loading"
        @click="handleSyncServerStrategySymbol"
      >
        同步server_strategy_symbol
      </el-button>
    </el-card>

    <el-card header="关联母账号-子账号" style="width: 100%; margin-top: 20px">
      <el-form label-width="80px">
        <el-form-item label="母账号">
          <el-select v-model="selectedAccount" placeholder="请选择母账号" style="width: 300px">
            <el-option
              v-for="item in accountList"
              :key="item.id"
              :label="item.name"
              :value="item.id"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="子账号">
          <el-checkbox-group v-model="checkedSymbols">
            <el-checkbox
              v-for="item in symbolList"
              :key="item.id"
              :label="item.id"
            >
              {{ item.name }}
            </el-checkbox>
          </el-checkbox-group>
        </el-form-item>
        <el-form-item style="text-align: right;">
          <el-button type="primary" @click="handleSubmit">提交</el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useBtnAuth } from '@/utils/btnAuth'
import { syncServerStrategySymbol, getSymbolsForLinkedLbankAccount, getLbankAccounts, submitLinkedLbankSubAccounts } from '@/api/lbank/db'
import { ElMessage } from 'element-plus'

const btnAuth = useBtnAuth()
const loading = ref(false)
const symbolList = ref([])
const checkedSymbols = ref([])
const accountList = ref([])
const selectedAccount = ref()

const fetchSymbolList = async () => {
  try {
    const res = await getSymbolsForLinkedLbankAccount()
    if (res.code === 0) {
      symbolList.value = res.data || []
    } else {
      ElMessage.error(res.msg || '获取symbol列表失败')
    }
  } catch (error) {
    console.error('获取symbol列表失败:', error)
    ElMessage.error('获取symbol列表失败，请稍后重试')
  }
}

const fetchAccountList = async () => {
  try {
    const res = await getLbankAccounts()
    if (res.code === 0) {
      accountList.value = res.data || []
    } else {
      ElMessage.error(res.msg || '获取账号列表失败')
    }
  } catch (error) {
    console.error('获取账号列表失败:', error)
    ElMessage.error('获取账号列表失败，请稍后重试')
  }
}

const handleSyncServerStrategySymbol = async () => {
  try {
    loading.value = true
    const res = await syncServerStrategySymbol()
    if (res.code === 0) {
      ElMessage.success(res.msg || '同步成功')
    } else {
      ElMessage.error(res.msg || '同步失败')
    }
  } catch (error) {
    console.error('同步失败:', error)
    ElMessage.error('同步失败，请稍后重试')
  } finally {
    loading.value = false
  }
}

const handleSubmit = async () => {
  if (!selectedAccount.value) {
    ElMessage.error('请选择母账号')
    return
  }
  if (!checkedSymbols.value.length) {
    ElMessage.error('请至少选择一个子账号')
    return
  }
  try {
    const res = await submitLinkedLbankSubAccounts({
      parentId: selectedAccount.value,
      subAccountIds: checkedSymbols.value
    })
    if (res.code === 0) {
      ElMessage.success(res.msg || '提交成功')
    } else {
      ElMessage.error(res.msg || '提交失败')
    }
  } catch (error) {
    console.error('提交失败:', error)
    ElMessage.error('提交失败，请稍后重试')
  }
}

onMounted(() => {
  fetchSymbolList()
  fetchAccountList()
})
</script>

<style scoped>
.custom-tree-node {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 14px;
  padding-right: 8px;
}

.node-count {
  color: #909399;
  font-size: 12px;
}
</style>