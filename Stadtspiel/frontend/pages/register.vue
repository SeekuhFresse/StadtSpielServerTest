<script setup lang="ts">
import PocketBase from 'pocketbase'

const pb = new PocketBase('http://127.0.0.1:8090')
const username = ref('')
const email = ref('')
const password = ref('')
const confirmPassword = ref('')

async function register() {
    if (password.value !== confirmPassword.value) {
        alert('Passwords do not match. Please try again.')
        return
    }

    try {
        const userRecord = await pb.collection('users').create({
            username: username.value,
            email: email.value,
            password: password.value
        })
        await navigateTo('/game')
    } catch (error) {
        console.error('Error creating user:', error)
    }
}

function navigateToLogin() {
    navigateTo('/login')
}
</script>

<template>
    <div class="register-container">
        <h1>Register</h1>
        <form @submit.prevent="register()">
            <div class="input-container">
                <label for="username">Username</label>
                <input v-model="username" id="username" type="text" required>
            </div>
            <div class="input-container">
                <label for="email">Email</label>
                <input v-model="email" id="email" type="email" required>
            </div>
            <div class="input-container">
                <label for="password">Password</label>
                <input v-model="password" id="password" type="password" required>
            </div>
            <div class="input-container">
                <label for="confirm-password">Confirm Password</label>
                <input v-model="confirmPassword" id="confirm-password" type="password" required>
            </div>
            <button type="submit">Register</button>
        </form>
        <button @click="navigateToLogin()">Back</button>
    </div>
</template>


<style scoped>
.register-container {
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