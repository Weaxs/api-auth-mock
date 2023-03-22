# api-auth-mock

[![GoDoc](https://godoc.org/github.com/Weaxs/api-auth-mock?status.svg)](https://godoc.org/github.com/Weaxs/api-auth-mock)
[![Go Report Card](https://goreportcard.com/badge/github.com/Weaxs/api-auth-mock)](https://goreportcard.com/report/github.com/Weaxs/api-auth-mock)
[![License](https://img.shields.io/github/license/Weaxs/api-auth-mock)](https://github.com/Weaxs/api-auth-mock/blob/main/LICENSE)

Support api to mock request authentication.

<table>
<thead> 
<tr> 
<th colspan=2 style="background:#F0F0F0F0;">mode</th> 
<th style="background:#F0F0F0F0;">api</th> 
<th style="background:#F0F0F0F0;">claims</th> 
<th style="background:#F0F0F0F0;">secret</th> 
</tr> 
</thead> 
<tbody> 
<tr> 
<td colspan="2" style="text-align: center;">basic authentication</td> 
<td>/api/basic/mock</td> 
<td>

```json
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
<td rowspan=3 style="text-align: center;">jwt </td> 
<td style="text-align: center;">HS256<br/>HS384<br/>HS512</td> 
<td>/api/jwt/mock/hmac</td> 
<td>

```json
{
  "key": "api-mock-hmac"
}
```

</td> 
<td>./conf/hmac_key</td> 
</tr> 
<tr> 
<td style="text-align: center;">RS256<br/>RS384<br/>RS512</td> 
<td>/api/jwt/mock/rsa</td> 
<td>

```json
{
  "key": "api-mock-rsa"
}
```

</td> 
<td>public : ./conf/public_key.pub <br/> private : ./conf/private_key </td> 
</tr> 
<tr> 
<td style="text-align: center;">ES256<br/>ES384<br/>ES512</td> 
<td>/api/jwt/mock/ecdsa</td> 
<td>

```json
{
  "key": "api-mock-ecdsa"
}
```

</td> 
<td>
es256 public : ./conf/ec256-public.pem<br/>
es256 private : ./conf/ec256-private.pem<br/>
es384 public : ./conf/ec384-public.pem<br/>
es384 private : ./conf/ec384-private.pem<br/>
es512 public : ./conf/ec512-public.pem<br/>
es512 private : ./conf/ec512-private.pem<br/>
</td> 
</tr> 
<tr>
<td colspan="2" style="text-align: center;">oauth</td> 
<td>token: /api/oauth/mock/token<br/>authorize: /api/oauth/mock/authorize</td> 
<td>

```json
{
  "client_id": "id0001",
  "client_secret": "secret0001"
}
```

</td>
<td>none</td> 
</tr>
</tbody>
</table>