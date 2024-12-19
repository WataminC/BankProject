<template>
  <div class="register-container">
    <div class="register-box">
      <h1 class="title">Register</h1>
      <form @submit.prevent="handleRegister">
        <div class="form-group">
          <label for="name">Username</label>
          <input
            type="text"
            id="name"
            v-model="name"
            placeholder="Enter your username"
            required
          />
        </div>
        <div class="form-group">
          <label for="email">Email</label>
          <input
            type="email"
            id="email"
            v-model="email"
            placeholder="Enter your email"
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
        <button type="submit" class="register-button">Register</button>
      </form>
      <div class="login-link">
        Already have an account? <a @click="goToLogin" class="login-button">Log In</a>
      </div>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

export default {
  name: 'Register',
  setup() {
    const name = ref('');
    const email = ref('');
    const password = ref('');
    const router = useRouter();

    const handleRegister = async () => {
      try {
        const response = await axios.post('http://localhost:8080/api/auth/register', {
          name: name.value,
          email: email.value,
          password: password.value,
        });
        console.log(response.data.Message);
        // Handle successful registration (e.g., navigate to login page)
        router.push('/login');
      } catch (error) {
        console.error(error.response.data.Error);
        // Handle registration error (e.g., show error message)
      }
    };

    const goToLogin = () => {
      router.push('/login');
    };

    return { name, email, password, handleRegister, goToLogin };
  }
};
</script>

<style scoped>
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: linear-gradient(to bottom right, #e8f5e9, #f3e5f5);
  font-family: 'Arial', sans-serif;
  animation: fade-in 1.5s ease-in;
}

.register-box {
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

.register-button {
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

.register-button:hover {
  background-color: #4db6ac;
  transform: scale(1.05);
}

.login-button {
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

.login-button:hover {
  background-color: #80cbc4;
  color: white;
}

.login-link {
  margin-top: 20px;
  text-align: center;
}

.login-link a {
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
</style>