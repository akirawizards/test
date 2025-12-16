<script setup>
// 1. ロジックを書く場所
import { ref } from 'vue'

// リアクティブな変数（画面に反映される変数）の定義
const title = ref('Hello Vue!')
const count = ref(0)

// 関数の定義
const increment = () => {
  count.value++
}
</script>

<template>
  <div class="app-container">
    <h1>{{ title }}</h1>
    <p>ボタンを押した回数: {{ count }}</p>
    
    <button @click="increment">
      カウントアップ
    </button>
  </div>
</template>

<style scoped>
/* 3. CSS（スタイル）を書く場所 */
.app-container {
  text-align: center;
  margin-top: 60px;
  font-family: Arial, sans-serif;
}

h1 {
  color: #42b983; /* Vueの公式カラー */
}

button {
  background-color: #42b983;
  color: white;
  border: none;
  padding: 10px 20px;
  border-radius: 4px;
  font-size: 16px;
  cursor: pointer;
}

button:hover {
  background-color: #3aa876;
}
</style>
