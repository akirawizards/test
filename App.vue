<template>
  <div id="app" class="app-container">
    <h1>{{ title }}</h1>
    <p>ボタンを押した回数: {{ count }}</p>
    
    <button @click="increment">
      カウントアップ
    </button>
  </div>
</template>

<script>
export default {
  name: 'App',
  // Vue 2では data は関数である必要があります
  data() {
    return {
      title: 'Hello Vue 2.3.3!',
      count: 0
    }
  },
  // 関数は methods の中に記述します
  methods: {
    increment() {
      // this を使ってデータにアクセスします
      this.count++
    }
  }
}
</script>

<style scoped>
.app-container {
  text-align: center;
  margin-top: 60px;
  font-family: Arial, sans-serif;
}
h1 {
  color: #42b983;
}
</style>
