<template>
  <div class="dashboard-container">
    <div class="dashboard-box">
      <h1 class="title">User Dashboard</h1>
      <div class="user-info">
        <p><strong>Username:</strong> {{ userInfo.username }}</p>
        <p><strong>Email:</strong> {{ userInfo.email }}</p>
        <p><strong>Account Number:</strong> {{ userInfo.account_number }}</p>
        <p><strong>Balance:</strong> ${{ userInfo.balance }}</p>
      </div>
      <div class="actions">
        <button @click="handleTransaction" class="action-button">Deposit/Withdraw</button>
        <button @click="transfer" class="action-button">Transfer</button>
        <button @click="viewFriends" class="action-button">Friends List</button>
      </div>
    </div>
    <button @click="logout" class="logout-button">Logout</button>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

export default {
  name: 'UserDashboard',
  setup() {
    const userInfo = ref({});
    const router = useRouter();

    const fetchUserInfo = async () => {
      try {
        const response = await axios.get('http://localhost:8080/api/account');
        userInfo.value = response.data;
        localStorage.setItem('userID', userInfo.value.userID);
      } catch (error) {
        console.error('Failed to fetch user info:', error.response.data.Error);
        // Handle error (e.g., redirect to login page)
      }
    };

    const handleTransaction = () => {
      router.push('/transaction');
    };

    const transfer = () => {
      // Handle transfer logic
      router.push('/transfer');
    };

    const viewFriends = () => {
      // Handle view friends list logic
      console.log('Friends List clicked');
    };

    const logout = () => {
      localStorage.removeItem('token');
      router.push('/login');
    };

    onMounted(() => {
      fetchUserInfo();
    });

    return { userInfo, handleTransaction, transfer, viewFriends, logout };
  }
};
</script>

<style scoped>
.dashboard-container {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: linear-gradient(to bottom right, #e8f5e9, #f3e5f5);
  font-family: 'Arial', sans-serif;
  animation: fade-in 1.5s ease-in;
}

.dashboard-box {
  background: #ffffff;
  border-radius: 8px;
  padding: 40px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  text-align: center;
  width: 100%;
  max-width: 600px;
  animation: fadeIn 1.5s ease-in-out;
}

.title {
  margin-bottom: 20px;
  color: #4a4a4a;
  font-size: 24px;
  font-weight: bold;
}

.user-info {
  margin-bottom: 20px;
  text-align: left;
}

.actions {
  display: flex;
  justify-content: space-between;
}

.action-button {
  flex: 1;
  margin: 0 10px;
  padding: 10px 20px;
  font-size: 16px;
  background-color: #80cbc4;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s, transform 0.3s;
}

.action-button:hover {
  background-color: #4db6ac;
  transform: scale(1.05);
}

.logout-button {
  margin-top: 20px;
  padding: 10px 20px;
  font-size: 16px;
  background-color: #ff7043;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s, transform 0.3s;
}

.logout-button:hover {
  background-color: #ff5722;
  transform: scale(1.05);
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