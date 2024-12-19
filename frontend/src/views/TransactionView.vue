<template>
  <div class="transaction-container">
    <div class="transaction-box">
      <h1 class="title">Deposit/Withdraw</h1>
      <div class="balance-info">
        <p><strong>Current Balance:</strong> ${{ userInfo.balance }}</p>
      </div>
      <div class="form-group">
        <label for="amount">Amount</label>
        <input
          type="number"
          id="amount"
          v-model="amount"
          placeholder="Enter amount"
          required
        />
      </div>
      <div class="actions">
        <button @click="deposit" class="action-button">Deposit</button>
        <button @click="withdraw" class="action-button">Withdraw</button>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import axios from 'axios';

export default {
  name: 'TransactionView',
  setup() {
    const userInfo = ref({});
    const amount = ref(0);

    const fetchUserInfo = async () => {
      try {
        const response = await axios.get('http://localhost:8080/api/account');
        userInfo.value = response.data;
      } catch (error) {
        console.error('Failed to fetch user info:', error.response.data.Error);
        // Handle error (e.g., redirect to login page)
      }
    };

    const deposit = async () => {
      try {
        const response = await axios.post('http://localhost:8080/api/transaction/deposit', {
          amount: amount.value.toString(),
        });
        userInfo.value.balance += amount.value;
        alert('Deposit successful');
      } catch (error) {
        alert('Deposit failed');
      }
    };

    const withdraw = async () => {
      try {
        const response = await axios.post('http://localhost:8080/api/transaction/withdraw', {
          amount: amount.value.toString(),
        });
        userInfo.value.balance -= amount.value;
        alert('Withdraw successful');
      } catch (error) {
        alert('Withdraw failed');
      }
    };

    onMounted(() => {
      fetchUserInfo();
    });

    return { userInfo, amount, deposit, withdraw };
  }
};
</script>

<style scoped>
.transaction-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: linear-gradient(to bottom right, #e8f5e9, #f3e5f5);
  font-family: 'Arial', sans-serif;
  animation: fade-in 1.5s ease-in;
}

.transaction-box {
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

.balance-info {
  margin-bottom: 20px;
  text-align: left;
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

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}
</style>