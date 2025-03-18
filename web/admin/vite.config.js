import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite' // ✅ 导入 AutoImport
import Components from 'unplugin-vue-components/vite' // ✅ 导入 Components
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers' // ✅ 导入 ElementPlusResolver

export default defineConfig({
  plugins: [
    vue(),
    AutoImport({
      resolvers: [ElementPlusResolver()], // 自动导入 Element Plus 组件
    }),
    Components({
      resolvers: [ElementPlusResolver()], // 自动解析 Element Plus 组件
    }),
  ],
})
