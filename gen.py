import requests
from faker import Faker
import random
import time

BASE_URL = "http://localhost:8080"
fake = Faker()
accounts = []

# Functions for API requests
def register_user():
    name = fake.user_name()
    email = fake.email()
    password = fake.password()
    response = requests.post(f"{BASE_URL}/api/auth/register", json={"name": name, "email": email, "password": password})
    if response.status_code == 200 and "Message" in response.json():
        # print(f"Registered user: {name}, {email}")
        return name, password
    return None, None

def login_user(name, password):
    response = requests.post(f"{BASE_URL}/api/auth/login", json={"name": name, "password": password})
    if response.status_code == 200 and "token" in response.json():
        return response.json()["token"]
    return None

def get_info(token):
    headers = {"Authorization": f"{token}"}
    response = requests.get(f"{BASE_URL}/api/account", headers=headers)
    if response.status_code == 200:
        return response.json()["account_number"], response.json()["userID"]
    return None, None

def deposit(token, amount):
    headers = {"Authorization": f"{token}"}
    response = requests.post(f"{BASE_URL}/api/transaction/deposit", json={"amount": str(amount)}, headers=headers)
    # if response.status_code == 200:
    #     print(f"Deposited {amount}")

def withdraw(token, amount):
    headers = {"Authorization": f"{token}"}
    response = requests.post(f"{BASE_URL}/api/transaction/withdraw", json={"amount": str(amount)}, headers=headers)
    # if response.status_code == 200:
    #     print(f"Withdraw {amount}")

def transfer(id, token, target_account, amount):
    headers = {"Authorization": f"{token}"}
    response = requests.post(f"{BASE_URL}/api/transaction/transfer", 
                             json={"account_id": id, "target_account_number": target_account, "amount": amount}, 
                             headers=headers)
    # if response.status_code == 200:
    #     print(f"Transferred {amount} to {target_account}")

def generate_data():
    for round_times in range(100000):
        if round_times % 1000 == 0:
            print(round_times)
        name, password = register_user()
        if not name:
            continue

        token = login_user(name, password)
        if not token:
            continue
        
        account, id = get_info(token)
        accounts.append(account)

        # Random deposits and withdrawals
        for _ in range(random.randint(1, 5)):
            deposit(token, round(random.uniform(10, 1000), 2))
            withdraw(token, round(random.uniform(1, 100), 2))

        # Optional: Add transfer logic
        # transfer(token, "3281543062464578", round(random.uniform(1, 100), 2))
        for _ in range(random.randint(1, 5)):
            i = random.randint(0, len(accounts) - 1)
            if account == accounts[i]:
                continue
            transfer(id, token, accounts[i], round(random.uniform(1, 100), 2))

        time.sleep(0.1)  # To avoid overwhelming the server

if __name__ == "__main__":
    generate_data()