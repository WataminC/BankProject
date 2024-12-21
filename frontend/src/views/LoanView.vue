<template>
  <div class="loan-container">
    <div class="loan-box">
      <h1 class="title">Loan Management</h1>
      <div v-if="loanExists" class="loan-info">
        <h2>Current Loan</h2>
        <table>
          <tr>
            <th>Amount</th>
            <th>Interest</th>
            <th>Total Amount</th>
            <th>Status</th>
          </tr>
          <tr>
            <td>{{ loan.amount }}</td>
            <td>{{ loan.interest }}</td>
            <td>{{ loan.total_amount }}</td>
            <td>{{ loan.status }}</td>
          </tr>
        </table>
      </div>
      <div class="loan-form">
        <h2>Request a Loan</h2>
        <div class="form-group">
          <label for="loanAmount">Amount</label>
          <input type="number" id="loanAmount" v-model="loanAmount" placeholder="Enter loan amount" required />
        </div>
        <div class="form-group">
          <label for="loanInterest">Interest</label>
          <input type="number" id="loanInterest" v-model="loanInterest" placeholder="Enter loan interest" required />
        </div>
        <button @click="requestLoan" class="loan-button">Request Loan</button>
      </div>
      <div class="repay-form">
        <h2>Repay Loan</h2>
        <div class="form-group">
          <label for="repayAmount">Amount</label>
          <input type="number" id="repayAmount" v-model="repayAmount" placeholder="Enter repayment amount" required />
        </div>
        <button @click="repayLoan" class="loan-button">Repay Loan</button>
      </div>
      <button @click="goBack" class="back-button">Back to Dashboard</button>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';

export default {
  name: 'LoanView',
  setup() {
    const loanAmount = ref('');
    const loanInterest = ref('');
    const repayAmount = ref('');
    const loanExists = ref(false);
    const loan = ref({});
    const router = useRouter();

    const checkLoanStatus = async () => {
      try {
        const response = await axios.post('http://localhost:8080/api/loan/request/query', {
          user_id: 1, // Replace with actual user ID
        });
        if (response.data.is_succeed) {
          alert('Loan request succeeded!');
        }
      } catch (error) {
        console.error(error);
      }
    };

    const requestLoan = async () => {
      try {
        await axios.post('http://localhost:8080/api/loan/request', {
          user_id: 1, // Replace with actual user ID
          amount: loanAmount.value,
          interest: loanInterest.value,
        });
        alert('Loan request submitted successfully');
      } catch (error) {
        console.error(error);
        alert('Loan request failed');
      }
    };

    const getLoan = async () => {
      try {
        const response = await axios.post('http://localhost:8080/api/loan', {
          user_id: 1, // Replace with actual user ID
        });
        loan.value = response.data;
        loanExists.value = true;
      } catch (error) {
        console.error(error);
        loanExists.value = false;
      }
    };

    const repayLoan = () => {
      alert('Repay loan functionality not implemented yet');
    };

    const goBack = () => {
      router.push('/dashboard');
    };

    onMounted(() => {
      checkLoanStatus();
      getLoan();
    });

    return { loanAmount, loanInterest, repayAmount, loanExists, loan, requestLoan, repayLoan, goBack };
  }
};
</script>

<style scoped>
.loan-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: linear-gradient(to bottom right, #e8f5e9, #f3e5f5);
  font-family: 'Arial', sans-serif;
}

.loan-box {
  background: #ffffff;
  border-radius: 8px;
  padding: 40px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  text-align: center;
  width: 100%;
  max-width: 600px;
}

.title {
  margin-bottom: 20px;
  color: #4a4a4a;
  font-size: 24px;
  font-weight: bold;
}

.loan-info, .loan-form, .repay-form {
  margin-bottom: 20px;
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

.loan-button, .back-button {
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

.loan-button:hover, .back-button:hover {
  background-color: #4db6ac;
  transform: scale(1.05);
}

table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 20px;
}

th, td {
  padding: 10px;
  border: 1px solid #ccc;
  text-align: left;
}

th {
  background-color: #f5f5f5;
}
</style>
