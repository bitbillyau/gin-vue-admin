<template>
  <div>
    <el-card header="数据库操作" style="width: 100%">
      <el-button 
        v-if="btnAuth.syncServerStrategySymbolBtn" 
        type="primary" 
        :loading="loading"
        @click="handleSyncServerStrategySymbol"
      >
        同步server_strategy_symbol
      </el-button>
    </el-card>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useBtnAuth } from '@/utils/btnAuth'
import { syncServerStrategySymbol } from '@/api/lbank/db'
import { ElMessage } from 'element-plus'

const btnAuth = useBtnAuth()
const loading = ref(false)

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
</script>