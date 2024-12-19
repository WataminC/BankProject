Development:

1. Read yaml configuration file
2. Define models
3. Connect database and auto migrate

4. Back end API
   1. Register and Login
   2. Login middleware
   3. Get Balance 

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