import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      "/query": {
        target: "http://localhost:8080",
        changeOrigin: true
      },
      "/login": {
        target: "http://localhost:8080",
        changeOrigin: true
      }
    }
  }
})
