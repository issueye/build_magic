import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import Icons from 'unplugin-icons/vite'
import Components from 'unplugin-vue-components/vite'
import { lazyImport, VxeResolver } from 'vite-plugin-lazy-import'
import monacoEditorPluginModule from 'vite-plugin-monaco-editor'
import { resolve } from 'path';


const isObjectWithDefaultFunction = (module: unknown): module is { default: typeof monacoEditorPluginModule } => (
  module != null &&
  typeof module === 'object' &&
  'default' in module &&
  typeof module.default === 'function'
)

const monacoEditorPlugin = isObjectWithDefaultFunction(monacoEditorPluginModule)
  ? monacoEditorPluginModule.default
  : monacoEditorPluginModule

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    Icons({}),
    // monaco
    monacoEditorPlugin({}),

    Components({}),
    lazyImport({
      resolvers: [
        VxeResolver({
          libraryName: 'vxe-table'
        }),
        // VxeResolver({
        //   libraryName: 'vxe-pc-ui',
        // })
      ]
    })
  ],
  resolve: {
    alias: [
      { find: '@', replacement: resolve(__dirname, 'src') },
    ],
  },

  build: {
    //false 的话会将项目中的所以 css 提取到一个 css 文件中
    cssCodeSplit: false,
  }
})
