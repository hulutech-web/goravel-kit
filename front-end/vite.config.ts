import { defineConfig, loadEnv } from 'vite';
import vue from '@vitejs/plugin-vue';
import path from 'path';
import Components from 'unplugin-vue-components/vite';
import { AntDesignVueResolver } from 'unplugin-vue-components/resolvers';
import { AntdvLessPlugin, AntdvModifyVars } from 'stepin/lib/style/plugins';
import AutoImport from "unplugin-auto-import/vite";
const timestamp = new Date().getTime();
const prodRollupOptions = {
  onwarn(warning, warn) {
    if (warning.code === 'TS_ERROR') return // 忽略 TS 错误
    warn(warning)
  },
  output: {
    chunkFileNames: (chunk) => {
      return 'assets/' + chunk.name + '.[hash]' + '.' + timestamp + '.js';
    },
    manualChunks(id: string) {
      if (id.includes('node_modules')) {
        return id.split('/node_modules/').pop()?.split('/')[0]
      }
    },
    assetFileNames: (asset) => {
      const name = asset.name;
      if (name && (name.endsWith('.css') || name.endsWith('.js'))) {
        const names = name.split('.');
        const extname = names.splice(names.length - 1, 1)[0];
        return `assets/${names.join('.')}.[hash].${timestamp}.${extname}`;
      }
      return 'assets/' + asset.name;
    },
  },
};
// vite 配置
export default ({ command, mode }) => {
  // 获取环境变量
  const env = loadEnv(mode, process.cwd());

  return defineConfig({
    server: {
      proxy: {
        '/api': {
          target: `${env.VITE_API_URL}/api`,
          ws: true,
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, ''),
        },
      },
      hmr: true,
    },

    resolve: {
      alias: {
        '@': path.resolve(__dirname, 'src'),
      },
    },
    esbuild: {
      legalComments: 'none',
      jsxFactory: 'h',
      jsxFragment: 'Fragment',
    },
    build: {
      minify: 'esbuild', // 替换 'terser'
      sourcemap: true,
      cssTarget: "chrome89",
      chunkSizeWarningLimit: 2048,
      rollupOptions: mode === 'production' ? prodRollupOptions : {},
    },
    plugins: [
      vue({
        template: {
          compilerOptions:{
            whitespace: 'preserve', // 避免 Vue 模板解析严格模式
          },
          transformAssetUrls: {
            img: ['src'],
            'a-avatar': ['src'],
            'stepin-view': ['logo-src', 'presetThemeList'],
            'a-card': ['cover'],
          },
        },
      }),
      Components({
        resolvers: [AntDesignVueResolver({ importStyle: mode === 'development' ? false : 'less' })],
      }),
      AutoImport({
        imports: ["vue", "vue-router"],
        dirs: [
          "src/composables/**/*",
          "src/store/**/*",
          "src/components/business/**/*",
        ],
        dts: "core/types/auto-imports.d.ts",
        vueTemplate: true,
        defaultExportByFilename: true,
      }),
    ],
    css: {
      preprocessorOptions: {
        less: {
          plugins: [AntdvLessPlugin],
          modifyVars: AntdvModifyVars,
          javascriptEnabled: true,
        },
      },
    },
    base: env.VITE_BASE_URL,
  });
};
