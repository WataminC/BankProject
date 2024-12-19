<template>
  <div class="login-container">
    <div class="login-box" v-if="showForm" @animationend="animationEnded">
      <h1 class="title">Welcome to Bank</h1>
      <form @submit.prevent="handleLogin">
        <div class="form-group">
          <label for="username">Username</label>
          <input
            type="text"
            id="username"
            v-model="username"
            placeholder="Enter your username"
            required
          />
        </div>
        <div class="form-group">
          <label for="password">Password</label>
          <input
            type="password"
            id="password"
            v-model="password"
            placeholder="Enter your password"
            required
          />
        </div>
        <button type="submit" class="login-button">Log In</button>
      </form>
      <div class="register-link">
        First Use? <a @click="goToRegister" class="register-button">Click me to register</a>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

export default {
  name: 'Login',
  setup() {
    const username = ref('');
    const password = ref('');
    const showForm = ref(false);
    const router = useRouter();

    const handleLogin = async () => {
      try {
        const response = await axios.post('http://localhost:8080/api/auth/login', {
          name: username.value,
          password: password.value,
        });
        console.log('Token:', response.data.token);
        // Handle successful login (e.g., store token, navigate to user dashboard)
        console.log("Test", response.data.token);
        localStorage.setItem('token', response.data.token);
        router.push('/dashboard');
      } catch (error) {
        console.error(error.response.data.Error);
        // Handle login error (e.g., show error message)
        alert('Login failed: ' + error.response.data.Error);
      }
    };

    const goToRegister = () => {
      router.push('/register');
    };

    const animationEnded = () => {
      console.log('Animation ended');
    };

    onMounted(() => {
      setTimeout(() => {
        showForm.value = true;
      }, 100);
    });

    return { username, password, handleLogin, showForm, animationEnded, goToRegister };
  }
};
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: linear-gradient(to bottom right, #e8f5e9, #f3e5f5);
  font-family: 'Arial', sans-serif;
  animation: fade-in 1.5s ease-in;
}

.login-box {
  background: #ffffff;
  border-radius: 8px;
  padding: 40px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  text-align: center;
  width: 100%;
  max-width: 500px;
  animation: fadeIn 1.5s ease-in-out;
}

.title {
  margin-bottom: 20px;
  color: #4a4a4a;
  font-size: 24px;
  font-weight: bold;
}

.form-group {
  margin-bottom: 15px;
  text-align: left;
}

label {
  display: block;
  margin-bottom: 5px;
  color: #666;
  font-weight: 600;
}

input {
  width: 100%;
  padding: 10px;
  font-size: 14px;
  border: 1px solid #ccc;
  border-radius: 4px;
  transition: border-color 0.3s;
}

input:focus {
  border-color: #80cbc4;
  outline: none;
}

.login-button {
  width: 100%;
  padding: 10px;
  font-size: 16px;
  background-color: #80cbc4;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s, transform 0.3s;
}

.login-button:hover {
  background-color: #4db6ac;
  transform: scale(1.05);
}

.register-button {
  margin-top: 10px;
  padding: 10px;
  font-size: 14px;
  background-color: transparent;
  color: #80cbc4;
  border: 1px solid #80cbc4;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s, color 0.3s;
}

.register-button:hover {
  background-color: #80cbc4;
  color: white;
}

.register-link {
  margin-top: 20px;
  text-align: center;
}

.register-link a {
  color: #80cbc4;
  cursor: pointer;
  text-decoration: none;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

@keyframes slide-up {
  from {
    transform: translateY(50px);
    opacity: 0;
  }
  to {
    transform: translateY(0);
    opacity: 1;
  }
}
</style>