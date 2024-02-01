<script setup lang="ts">
import PocketBase from 'pocketbase'

const router = useRouter()
const pb = new PocketBase('http://127.0.0.1:8090')
const email = ref('')
const password = ref('')

async function login() {
  const authData = await pb.collection('users').authWithPassword(
    email.value,
    password.value,
    )
  if (pb.authStore.isValid) {
    router.push('/game')
  } else {
    console.log('Invalid credentials')
  }
  console.log(authData)
}

function navigateToRegister() {
  router.push('/register')
}
</script>

<template>
  <div class="login-container">
    <h1>Login</h1>
    <form @submit.prevent="login()">
      <div class="input-container">
        <label for="email">Email</label>
        <input v-model="email" id="email" type="email" required>
      </div>
      <div class="input-container">
        <label for="password">Password</label>
        <input v-model="password" id="password" type="password" required>
      </div>
      <button type="submit">Login</button>
    </form>
    <button @click="navigateToRegister()">Register</button>
  </div>
</template>

<style scoped>
.login-container {
  width: 300px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  margin-top: 50px;
}

h1 {
  text-align: center;
  margin-bottom: 20px;
}

.input-container {
  margin-bottom: 20px;
}

label {
  display: block;
  margin-bottom: 5px;
}

input {
  width: 100%;
  padding: 5px;
  border: 1px solid #ccc;
  border-radius: 3px;
}

button {
  width: 100%;
  padding: 5px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 3px;
  cursor: pointer;
  margin-top: 20px;
}

button:hover {
  background-color: #0056b3;
}
</style>