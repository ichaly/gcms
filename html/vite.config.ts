import {resolve} from 'path'
import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'

import Components from "unplugin-vue-components/vite";
import {ElementPlusResolver} from "unplugin-vue-components/resolvers";

// https://vitejs.dev/config/
export default defineConfig({
    resolve: {
        // 导入文件夹别名
        alias: {
            '~': resolve(__dirname, './src'),
        },
    },
    server: {
        // 是否自动在浏览器打开
        // open: true,
        // 是否开启 https
        // https: false,
        // 服务端渲染
        // ssr: false,
        proxy: {
            '/api': {
                target: 'http://localhost:3333/',
                changeOrigin: true,
                ws: true,
                rewrite: (pathStr) => pathStr.replace('/api', '')
            },
        },
    },
    plugins: [
        vue(),
        Components({
            resolvers: [
                ElementPlusResolver({
                    importStyle: "sass",
                }),
            ],
        }),
    ],
    css: {
        preprocessorOptions: {
            scss: {
                additionalData: `@use "~/styles/element/index.scss" as *;`,
            },
        },
    },
    base: '/',
    root: './src/pages/',
    build: {
        emptyOutDir: true,
        outDir: '../../dist',
        rollupOptions: {
            input: {
                index: resolve(__dirname, 'src/pages/index.html'),
                admin: resolve(__dirname, 'src/pages/admin/index.html'),
            },
            output: {
                chunkFileNames: 'static/js/[name]-[hash].js',
                entryFileNames: 'static/js/[name]-[hash].js',
                assetFileNames: 'static/[ext]/[name]-[hash].[ext]',
            }
        },
    }
})
