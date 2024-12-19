<template>
  <div class="transfer-container">
    <div class="transfer-box">
      <h1 class="title">Transfer</h1>
      <div class="form-group">
        <label for="targetAccount">Target Account Number</label>
        <input
          type="text"
          id="targetAccount"
          v-model="targetAccount"
          placeholder="Enter target account number"
          required
        />
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
      <button @click="transfer" class="action-button">Transfer</button>
      <button @click="goBack" class="back-button">Back to Dashboard</button>
    </div>
    <div class="transaction-history">
      <h2 class="subtitle">Transaction History</h2>
      <ul>
        <li v-for="transaction in transactions" :key="transaction.ID">
          <p><strong>Type:</strong> {{ transaction.Type }}</p>
          <p><strong>Amount:</strong> ${{ transaction.Amount }}</p>
          <p><strong>Target Account:</strong> {{ transaction.TargetAccountNumber }}</p>
          <p><strong>Date:</strong> {{ new Date(transaction.CreatedAt).toLocaleString() }}</p>
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

export default {
  name: 'TransferView',
  setup() {
    const targetAccount = ref('');
    const amount = ref(0);
    const transactions = ref([]);
    const router = useRouter();

    const fetchTransactions = async () => {
      try {
        const response = await axios.get('http://localhost:8080/api/transaction');
        transactions.value = response.data;
      } catch (error) {
        console.error('Failed to fetch transactions:', error.response.data.Error);
        // Handle error (e.g., redirect to login page)
      }
    };

    const transfer = async () => {
      try {
        console.log(localStorage.getItem('userID'))
        const response = await axios.post('http://localhost:8080/api/transaction/transfer', {
          account_id : Number(localStorage.getItem('userID')),
          target_account_number: targetAccount.value,
          amount: amount.value
        });
        alert('Transfer successful');
        await fetchTransactions(); // Refresh transaction history
      } catch (error) {
        console.error('Failed to transfer:', error.response.data.Error);
        alert('Transfer failed: ' + error.response.data.Error);
      }
    };

    const goBack = () => {
      router.push('/dashboard');
    };

    onMounted(() => {
      fetchTransactions();
    });

    return { targetAccount, amount, transactions, transfer, goBack };
  }
};
</script>

<style scoped>
.transfer-container {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  height: 100vh;
  padding: 20px;
  background: linear-gradient(to bottom right, #e8f5e9, #f3e5f5);
  font-family: 'Arial', sans-serif;
  animation: fade-in 1.5s ease-in;
}

.transfer-box {
  background: #ffffff;
  border-radius: 8px;
  padding: 40px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  text-align: center;
  width: 45%;
  animation: slide-in 0.8s ease-out;
}

.title {
  margin-bottom: 20px;
  color: #4a4a4a;
  font-size: 24px;
  font-weight: bold;
}

.subtitle {
  margin-bottom: 20px;
  color: #4a4a4a;
  font-size: 20px;
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

.action-button {
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

.action-button:hover {
  background-color: #4db6ac;
  transform: scale(1.05);
}

.transaction-history {
  width: 45%;
  background: #ffffff;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  animation: slide-in 0.8s ease-out;
  overflow-y: auto;
  max-height: 80vh;
}

.transaction-history ul {
  list-style: none;
  padding: 0;
}

.transaction-history li {
  margin-bottom: 15px;
  padding: 10px;
  border-bottom: 1px solid #ccc;
}

.transaction-history li:last-child {
  border-bottom: none;
}

.back-button {
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

.back-button:hover {
  background-color: #ff5722;
  transform: scale(1.05);
}

@keyframes fade-in {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

@keyframes slide-in {
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