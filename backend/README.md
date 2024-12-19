Development:

1. Read yaml configuration file
2. Define models
3. Connect database and auto migrate

4. Back End API
   1. Register and Login
   2. Login middleware
   3. Get Balance 
   4. Deposit and Withdraw
   5. Get and Add Transaction

# API 

1. 用户注册接口
- 接口路径：/api/auth/register
- 方法：POST

输入示例：
```json
{
  "name": "user123",
  "email": "user@example.com",
  "password": "securepassword"
}
```
输出：

- 成功：
```json
{
  "Message": "Add user successfully"
}
```
- 失败（如邮箱格式错误、邮箱/用户名已注册等）：
```json
{
  "Error": "Email or Name already registered"
}
```
功能：

- 注册新用户，验证输入是否合法。
- 若邮箱或用户名重复，返回错误。
- 为新用户创建初始账户，余额为0。

2. 用户登录接口
- 接口路径： /api/auth/login
- 方法： POST

输入示例：
```json
{
  "name": "user123",
  "password": "securepassword"
}
```

输出：
- 成功：
```json
{
  "token": "jwt-token-string"
}
```

- 失败（用户名或密码错误）：
```json
{
  "Error": "wrong credentials"
}
```

功能：
- 验证用户名和密码。
- 返回用户的JWT令牌，用于后续验证。

3. 查看用户信息接口
- 路径: /api/user
- 方法: GET

- 请求头:
```json
{
  "Authorization": "Bearer generated-jwt-token"
}
```
```json
{
    "account_number": "3281543062464578",
    "balance": 15.4,
    "email": "admin@example.com",
    "userID": 1,
    "username": "admin"
}
```
- 功能描述: 根据用户的 JWT token 返回用户的基本信息。

4. 存款接口
- 路径: /api/deposit
- 方法: POST

请求体:
```json
{
  "amount": 100.0
}
```

响应体:
```json
{
  "message": "Deposit successful"
}
```
- 功能描述: 用户向指定的账户存款。

5. 取款接口
- 路径: /api/withdraw
- 方法: POST

请求体:
```json
{
    "amount" : "0.01"
}

```
响应体:
```json
{
    "message": "Withdraw successfully",
    "new_balance": 3.39
}
```
- 功能描述: 用户从指定的账户中取款。

6. 获得交易记录接口
- 路径: /api/transaction
- 方法: GET

请求头:
```json
{
  "Authorization": "Bearer generated-jwt-token"
}
```

响应体:
- 有交易记录时
```json
[
  {
    "ID": 23,
    "AccountID": 2,
    "Type": "Transfer",
    "Amount": 1,
    "TargetAccountNumber": "3281543062464578",
    "CreatedAt": "2024-12-19T22:17:15.906692+08:00",
    "UpdatedAt": "2024-12-19T22:17:15.906692+08:00"
  },
  {
    "ID": 24,
    "AccountID": 2,
    "Type": "Transfer",
    "Amount": 1,
    "TargetAccountNumber": "3281543062464578",
    "CreatedAt": "2024-12-19T22:18:19.905653+08:00",
    "UpdatedAt": "2024-12-19T22:18:19.905653+08:00"
  }
]
```
- 无交易记录时
```json
"message": "No transactions found"
```

功能描述:
- 根据用户的 JWT token 返回用户的交易记录。

7. 转账接口

- 路径: /api/transaction/transfer
- 方法: POST

请求体:
```json
{
  "account_id": 1,
  "target_account_number": "3281543062464578",
  "amount": 100.0
}
```

响应体:
```json
{
  "message": "Transaction successful"
}
```

功能描述:
- 用户从一个账户向另一个账户转账。