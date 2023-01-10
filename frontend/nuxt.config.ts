// https://v3.nuxtjs.org/api/configuration/nuxt.config
export default defineNuxtConfig({
  modules: ["@nuxtjs/tailwindcss"], // モジュールの使用
  runtimeConfig: {
    public: {
      PUBLIC_BACKEND_URL: process.env.PUBLIC_BACKEND_URL,
    },
  },
});
