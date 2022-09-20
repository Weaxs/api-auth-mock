# api-auth-mock
Support api to mock request authentication.

<table>
<thead> 
<tr> 
<th colspan=2>mode 验证模式</th> 
<th>api 接口</th> 
<th>claims 断言用户</th> 
<th>secret 密钥</th> 
</tr> 
</thead> 
<tbody> 
<tr> 
<td colspan=2 align="center">basic authentication</td> 
<td>/api/basic/mock</td> 
<td>

```json
// Accounts
{
  "account1": "password1",
  "account2": "password2",
  "account3": "password3"
}
```

</td>
<td>none</td> 
</tr> 
<tr> 
<td rowspan=3>jwt </td> 
<td>HS256/HS384/HS512</td> 
<td>/api/jwt/mock/hmac</td> 
<td>

```json
{
  "key": "api-mock-rsa"
}
```

</td> 
<td>./conf/public_key.pub</td> 
</tr> 
<tr> 
<td>RS256/RS384/RS512</td> 
<td>/api/jwt/mock/rsa</td> 
<td>

```json
{
  "key": "api-mock-hmac"
}
```

</td> 
<td>./conf/hmac_key</td> 
</tr> 
</tbody>
</table>