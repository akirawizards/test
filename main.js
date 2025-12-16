import Vue from 'vue'
import App from './App.vue'

// Vue 2.3.3 の設定（プロダクションのヒントを非表示）
Vue.config.productionTip = false

new Vue({
  // h は createElement のエイリアスとして慣習的に使われます
  render: h => h(App),
}).$mount('#app')
