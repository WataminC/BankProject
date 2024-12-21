<template>  <div class="admin-container">    <h1 class="admin-title">Admin Dashboard</h1>    <table class="loan-table" v-if="loanRequests.length">      <thead>        <tr>          <th>ID</th>          <th>User</th>          <th>Amount</th>          <th>Interest</th>          <th>Status</th>          <th>Actions</th>        </tr>      </thead>
      <tbody>
        <tr 
          v-for="item in loanRequests" 
          :key="item.id" 
          class="loan-row" 
          @animationend="rowAnimationEnded"
        >
          <td>{{ item.id }}</td>
          <td>{{ item.user_id }}</td>
          <td>{{ item.amount }}</td>
          <td>{{ item.interest }}</td>
          <td>{{ item.status }}</td>
          <td>
            <button class="approve-button" @click="processRequest(item.id, true)">Approve</button>
            <button class="reject-button" @click="processRequest(item.id, false)">Reject</button>
          </td>
        </tr>
      </tbody>
    </table>
    <div v-else class="no-request">No pending requests</div>
    <button class="logout-button" @click="logout">Logout</button>
  </div>
</template>

<script>
import axios from 'axios';
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';

export default {
  name: 'AdminView',
  setup() {
    const loanRequests = ref([]);
    const router = useRouter();

    const fetchLoanRequests = async () => {
      try {
        const res = await axios.get('http://localhost:8080/api/admin/loan');
        loanRequests.value = res.data;
      } catch (error) {
        console.error(error);
      }
    };

    const processRequest = async (id, approved) => {
      try {
        await axios.post('http://localhost:8080/api/admin/loan', {
          id,
          is_approved: approved,
        });
        loanRequests.value = loanRequests.value.filter(item => item.id !== id);
      } catch (error) {
        console.error(error);
      }
    };

    const logout = () => {
      localStorage.removeItem('token');
      router.push('/login');
    };

    const rowAnimationEnded = () => {
      // ...
    };

    onMounted(() => {
      fetchLoanRequests();
    });

    return { loanRequests, processRequest, logout, rowAnimationEnded };
  },
};
</script>

<style scoped>
.admin-container {
  margin: 40px auto;
  max-width: 800px;
  text-align: center;
  animation: fadeIn 1s ease-in-out;
}

.admin-title {
  font-size: 24px;
  margin-bottom: 20px;
  color: #4a4a4a;
}

.loan-table {
  width: 100%;
  border-collapse: collapse;
  margin-bottom: 30px;
  animation: slide-up 1s ease-in-out;
}

.loan-table th, .loan-table td {
  border: 1px solid #ccc;
  padding: 12px;
}

.no-request {
  margin: 20px 0;
  font-size: 18px;
  color: #666;
}

.logout-button {
  padding: 10px 20px;
  background-color: #80cbc4;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  color: #fff;
  transition: background-color 0.3s, transform 0.3s;
}

.logout-button:hover {
  background-color: #4db6ac;
  transform: scale(1.05);
}

.approve-button, .reject-button {
  width: 100%;
  padding: 10px;
  font-size: 14px;
  background-color: #80cbc4;
  color: #fff;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s, transform 0.3s;
  margin-bottom: 6px;
}

.approve-button:hover, .reject-button:hover {
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