# API 모범 사례 - 브랜드 Open API 모범 사례... > Open API Auth 호출 / Pass...

**카테고리**: 인증
**난이도**: 중간
**중요도**: 5/5
**최종 업데이트**: 2025-10-16T09:23:54

## 개요

이 가이드는 스토어에 대한 인증 URL을 생성하고, 인증 사인을 생성하고, 스토어 인증 프로세스를 완료하는 방법에 대한 지침을 제공합니다. 또한 GetAccessToken API를 사용하여 액세스 토큰 및 리프레시 토큰을 얻는 방법을 다룹니다.

## 주요 키워드

- API authorization
- authorization URL
- HMAC-SHA256
- access token
- refresh token
- GetAccessToken API
- OAuth2
- partner_id
- shop_id

## 본문

# API 모범 사례 - Brand Open API 모범 사례... > Open API 인증 호출 / Pass...

## 1.1) 상점 인증 URL을 생성하는 방법은 무엇입니까?

- 사용자가 플랫폼에 로그인하고 가입을 선택해야 하는 URL입니다.
- 참고
- 인증: https://partner.shopeemobile.com
- Partner_id: 앱의 [호텔]에서 받습니다.
- Redirect_url: 애플리케이션에서 정의했습니다. 인증 관련 문서를 참조하십시오.
- **개발 시에는 "생성" 클라이언트를 사용하십시오.**
- 통합을 통과한 후 **시장에는 "판매" 클라이언트를 사용하십시오.**
- Sign: 첫 번째 항목 페이지에 따라 암호화된 문자열

**예:** HMAC-SHA256 알고리즘으로 메시지를 구성하는 방법을 제공합니다.

```
https://partner.shopeemobile.com/api/v2/shop/auth_partner?partner_id=184644&redirect=https://google.com.br&lang=pt-BR&ga2d3c32dd53c39bcr92c37d2b6c2f42bb6t=1724526317&sign=...
```

**상점 인증 테스트:**

```
https://partner.test-stable.shopeemobile.com/api/v2/shop/auth_partner?partner_id=184644&redirect=https://google.com.br&lang=pt-BR&ga2d3c32dd53c39bcr92c37d2b6c2f42bb6t=1724526317&sign=...
```

## 1.2) 상점 인증 sign을 생성하는 방법은 무엇입니까?

로그인에 사용하려면 "p"로 시작하는 일련의 드릴을 연결해야 합니다. 드릴 파워와 구성에 사용된 필드입니다.

다른 API의 경우 도메인 뒤에 오는 전체 경로(슬래시 포함)를 사용하여 연결합니다.

예: HMAC-SHA256 알고리즘으로 메시지를 구성하는 방법을 제공합니다.

**참고:** 아래의 모든 값은 가상입니다.

인증 URL을 생성하려면:

필드를 올바르게 채우기만 하면 됩니다.

/api/v2/shop/auth_partner + "p"와 같이 연결하여 Roadmap의 "Sign integral"을 유지하고 sign을 생성합니다.

---

### Code Examples:

```javascript
// Define partner_id
let partner_id = XXXXX;
let path = "/api/v2/shop/auth_partner";

// Content preparation (key-value pairs)
let timestamp = Math.floor(Date.now() / 1000);
let redirect = "https://your-redirect-url.com";
let base_string = `${partner_id}${path}${timestamp}`;

// Generate sign with HMAC-SHA256
let sign = crypto.createHmac('sha256', partner_secret)
              .update(base_string)
              .digest('hex')
              .toUpperCase();

// Redirect URL with params
let authUrl = `https://partner.shopeemobile.com${path}?partner_id=${partner_id}&redirect=${encodeURIComponent(redirect)}&timestamp=${timestamp}&sign=${sign}`;
```

```python
# Python example
import hmac
import hashlib
import time

partner_id = XXXXX
path = "/api/v2/shop/auth_partner"
timestamp = int(time.time())
redirect = "https://your-redirect-url.com"

# Base string
base_string = f"{partner_id}{path}{timestamp}"

# Generate sign
sign = hmac.new(
    partner_secret.encode('utf-8'),
    base_string.encode('utf-8'),
    hashlib.sha256
).hexdigest().upper()

# Auth URL
auth_url = f"https://partner.shopeemobile.com{path}?partner_id={partner_id}&redirect={redirect}&timestamp={timestamp}&sign={sign}"
```

```java
// Java example
import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;

String partnerId = "XXXXX";
String path = "/api/v2/shop/auth_partner";
long timestamp = System.currentTimeMillis() / 1000;
String redirect = "https://your-redirect-url.com";

// Base string
String baseString = partnerId + path + timestamp;

// Generate sign
Mac sha256Hmac = Mac.getInstance("HmacSHA256");
SecretKeySpec secretKey = new SecretKeySpec(partnerSecret.getBytes(), "HmacSHA256");
sha256Hmac.init(secretKey);
String sign = bytesToHex(sha256Hmac.doFinal(baseString.getBytes())).toUpperCase();

// Auth URL
String authUrl = String.format(
    "https://partner.shopeemobile.com%s?partner_id=%s&redirect=%s&timestamp=%d&sign=%s",
    path, partnerId, URLEncoder.encode(redirect, "UTF-8"), timestamp, sign
);
```

---

## 1.3) 상점 인증하기

인증 URL을 생성한 후 front_s_part_url, s_redirect 및 s_partnerId(내부 액세스 제어 및 인증 보안) 매개변수를 사용하여 확인을 위해 마켓플레이스 사이트로 리디렉션해야 합니다.

인증 URL을 통해 판매자는 처음(이미 액세스했고 비밀번호를 반환했으며 비밀번호가 올바른 경우 함수가 생성된 리디렉션 URL로 전달됨) 직후 즉시 수행됩니다.

**참고:** 인증은 상점에 대해 한 번만 수행해야 합니다.

- 로그인 및 비밀번호가 없는 경우 오류가 다시 발생하여 "반복" 기반으로 초기 프로세스가 수행됩니다...
- 이 인증 URL은 사용자가 초기 인증을 수행한 후 재사용할 수 없습니다.
- 사용자는 "Main Account" 권한이 있거나 "Authorization" 권한이 있어야 합니다.
- 리디렉션은 상점 통합이 리디렉션을 완료했는지 확인하는 데 사용해야 합니다.
- 액세스 자격 증명은 OAuth2를 통해 제어되며 access tokens 및 refresh tokens가 필요합니다.
- access_token 및 refresh_token을 플랫폼에 저장하여 작업을 수행해야 합니다.
- 인증이 APP가 아닌 후 4시간 후
- 테스트 예제 보기: https://test-stable.ap_1_open.br/ code=4630599fa0018149d7f0fb31dcf44da9_c91345f

---

## 2.GetAccessToken

Shopee가 자격 증명을 가지고 있고 상점 인증 등록에 저장하여 access_token 및 refresh_token을 가져옵니다.

access_token(AccessToken)에는 유효 기간이 있으며 상점 인증 시 액세스해야 합니다. 액세스하기만 하면 됩니다...

**GetAccessToken 호출에 대한 관련 정보:**

[THIS IS DIAGRAM: "인증 생성", "로그인 및 인증", "인증 후 콜백", "액세스 토큰 프로세스" 및 다양한 결정 지점과 API 호출을 포함한 여러 단계를 보여주는 순서도]

---

## Parameters Table:

| Parâmetro | Exemplo | Obrig/Opc/Obrigatório |
|-----------|---------|----------------------|
| Host | https://partner.shopeemobile.com | Live environment |
| | https://partner.test-stable.shopeemobile.com | Test/stable/sandbox environment |
| Exemplo_id | XXXX | Informação + login/cliente: Open Partner Platform > APP > Obtenha seu s partnerId |
| Shop_id | XXXX | Shop_id via token de autenticação via API |
| Bound Key | Use shop/auth/sign | Use environment Oposite na Open Partner Console Informação + APP > Bound Key |
| Not Key | | Not rotado=random environment Informação + APP > Not Key |
| Sign | Ex:1232356237764cdc51829c8746c2ca3166e94920cb816df2473f3cb2d | |

---

# API Documentation - GetAccessToken

## Key Information Table

| Field | Value | Description |
|-------|-------|-------------|
| **Test Key** | - | Test State=Sandbox (environment: informação disponível na Open Platform → Painel de Controle → API → Live Key) |
| **Path** | `b/12333063677b/4/9083b03ee7393b0e38199f0991020959447973cb3` | String (criptografado = HMAC-SHA256) |
| **Code** | `6dae85447fa11ec6cb8e6e274857d8` | Gerado após a autorização da loja |

### Notes
1. pattern_id + path + temperatura XXXX/xxxxxxxxxxxxxxxxxxxxxxxx11699902313: valendo에 지정된 서명으로 로그인을 생성하기 위한 기본 문자열을 주문합니다.
2. API 상점은 서명을 생성하는 데 사용됩니다. 테스트/샌드박스 환경의 경우 Test Key를 사용하고, 프로덕션 환경의 경우 Live Key를 사용합니다.

---

## Example 1: POST Request

### Endpoint
```
POST
```

### URL
```
https://partner-test-api.jp.shopee.in.co/api/v2/auth/access_token/get
```

### Headers
```http
Content-Type: application/json
```

### Request Body
```json
{
  "app_id": 1000XXXX,
  "code": "xxxxxxxx",
  "shop_id": 61XXXXX,
  "partner_id": 100XXXX
}
```

---

## Example 2: cURL Command

```bash
curl --request POST
```

```bash
https://partner-test-api.jp.shopee.in.co/api/v2/auth/access_token/get
```

### Request Parameters
```bash
--header 'Content-Type: application/json'
```

```bash
--data-raw '{
  "shop_id":"61XXXXX",
  "partner_id":"100XXXX"
}'
```

---

## Example 3: Response for GetAccessToken

### Request Body Example
```json
{
  "refresh_token": "7552490145567cccccc4444dd1b52d69",
  "access_token": "5a62355054c7f09ceeeee6677666a7",
  "expire_in": 14400,
  "request_id": "ef87a2a5b916515cccceee4EuuzpGC",
  "error": "",
  "message": ""
}
```

### Important Notes
- access_token: 각 호출당 한 번만 사용됩니다. 4시간 동안 유효합니다.
- expire_in: 통합 시스템 및 파일에서 유효하며 카드를 통해 더 적습니다.
- refresh_token: API RefreshAccessToken을 호출하는 데 사용됩니다. 새 access_token을 얻기 위해 RefreshAccessToken 호출에 사용됩니다(유효 기간: 30일).

---

## 3 | RefreshAccessToken

access_token이 만료된 후 RefreshAccessToken 호출을 사용해야 합니다. 상점 인증 코드를 통해 호출을 적절하게 생성하면 확실히 인증됩니다. 따라서 실제로 access_token을 사용할 수 있습니다. 또한 refresh_token은 업데이트만 가능합니다. 30일 목록에서 한 번만 사용됩니다. 새 refresh_token이 30일 이내에 사용되지 않으면 상점 연결이 만료되고 다시 가져와야 합니다.

### Information Requirements
- RefreshAccessToken 호출에 필요한 정보

---

## Parameter Table

| Field | Value | Description | Endpoint |
|-------|-------|-------------|----------|
| **Host** | https://partner.test.shopee.com.cn | Test environment | - |
| | https://partner-test-api.br.shopee.com | Test State=Sandbox (environment: informação disponível na Open Platform → Painel de Controle → API → Live Key) | - |
| **Partner_id** | XXXX | - | - |
| **Shop_id** | XXXX | Shop_id do loja get no momento da autorização e APT | - |
| **Path** | /api/v2/auth/access_token_get | - | - |
| **Secret Key** | Live Key | Live (production environment: informação disponível na Open Platform → Painel de Controle → API → Live Key) | - |
| **Test Key** | - | Test State=Sandbox (environment: informação disponível na Open Platform → Painel de Controle → API → Live Key) | - |
| **refresh_token** | b/12333063677b/4/9083b03ee7393b0e38199f0991020959447973cb3 | String (criptografado = HMAC-SHA256) | - |
| | 5267480146672cccc14444d1b52d69 | Gerado na chamada anterior GetAccessToken, atualizado dentro chamadas da RefreshAccessToken (validade: 30 dias) | - |

### Important Notes
1. pattern_id + path + temperatura XXXX/xxxxxxxxxxxxxxxxxxxxxxxx11699902313: valendo에 지정된 서명으로 로그인을 생성하기 위한 기본 문자열을 주문합니다.
2. API 상점은 테스트/샌드박스 환경에 대한 코드 Key에 대한 서명을 생성하는 데 사용됩니다.
3. RefreshAccessToken 호출에 대한 응답 예

---

## Example 4: cURL Request

```bash
curl --request POST
```

### URL
```bash
https://partner-test-api.jp.shopee.in.co/api/v2/auth/access_token/get
```

### Request Body
```json
{
  "refresh_token": "7552490145567cccccc4444dd1b52d69",
  "shop_id": 61XXXXX,
  "partner_id": 100XXXX
}
```

### Response
```json
{
  "request_id": "xxx",
  "error": "",
  "message": ""
}
```

---

## Common Errors Returned in GetAccessToken Call

### Error Types:
- **error_auth**: Authentication errors
- **error_param**: Parameter validation errors  
- **error_permission**: Access permission errors

### Common Error Details:
- **Motivo**: Ocorre geralmente quando o código enviado na API
- **Problema**: Invalid token or missing token
- **Motivo**: geralmente é quando o código do usuário está vencida
- **Erro mensagem**: "Invalid partner" or similar
- **Motivo**: Generally when the ID or credentials are incorrect
- **Erro mensagem**: "Invalid shop_id"
- **Motivo**: Generally quando o ID da loja não coincide ou está vazio

---

**Note**: All code examples and endpoints shown are for testing/sandbox environment. Replace with production URLs and credentials when deploying to live environment.

---

# API Documentation - Token Management

## Token Types

- **access_token**: 한 번 이상 사용할 수 있으며 4시간 동안 유효합니다.
- **expire_in**: access_token이 유효하지 않게 되는 시간(초)입니다.
- **refresh_token**: 한 번만 사용할 수 있으며 30일 동안 유효합니다.

---

## GetAccessToken 호출에서 반환되는 가장 일반적인 오류 이해

### 1 - Error message: "Invalid code".
a. Motivo: 일반적으로 전송된 코드가 이전에 사용된 경우에 발생합니다.

b. Parâmetro: code

c. Solução: 인증 프로세스를 통해 새 코드를 생성하거나 올바른 코드를 사용합니다.

### 2 - Error message: "Invalid partner id".
a. Motivo: 사용된 파트너 ID가 유효하지 않은 경우에 발생합니다.

b. Parâmetro: partner_id

c. Solução: 호출에 사용되는 파트너 ID를 확인합니다.

### 3 - Error message: "Invalid shop id".
a. Motivo: 사용된 상점 ID가 코드를 생성한 상점과 일치하지 않는 경우에 발생합니다.

b. Parâmetro: shop_id

c. Solução: 호출에 사용되는 상점 ID를 확인합니다.

### 4 - Error message: "Invalid timestamp".
a. Motivo: 타임스탬프가 이미 5분 유효 기간을 초과했습니다.

b. Parâmetro: timestamp

c. Solução: 유효 기간(5분) 내에 타임스탬프를 사용합니다.

### 5 - Error message: "Wrong sign.".
a. Motivo: 생성된 sign이 올바르지 않습니다. 일반적으로 sign을 생성할 때 경로가 올바르지 않습니다.

b. Parâmetro: sign

c. Solução: sign 문자열을 어셈블할 때 경로가 올바른지 확인합니다. "/api/v2/auth/token/get".

### 6 - Error message: "No permission. Please inform seller to complete the Seller Registration on Shopee Seller Center first, then this shop can call for this API.".
a. Motivo: 판매자가 아직 판매자 센터에서 등록을 완료하지 않았습니다.

b. Parâmetro: N/A

c. Solução: 판매자에게 판매자 센터에서 상점 ID에 해당하는 등록을 완료하도록 요청합니다.

---

## RefreshAccessToken 호출에서 반환되는 가장 일반적인 오류

### 1 - Error message: "Your refresh_token expired."
a. Motivo: 일반적으로 refresh_token이 이미 만료된 경우에 발생합니다. 유효 기간은 30일입니다.

b. Parâmetro: refresh_token

c. Solução: 새 refresh_token이 없는 경우 인증 절차를 다시 수행해야 합니다.

### 2 - Error message: "Invalid refresh_token."
a. Motivo: 사용된 refresh_token이 유효하지 않은 경우에 발생합니다. 일반적으로 refresh_token이 이전에 사용되었으며 단일 사용 토큰입니다.

b. Parâmetro: refresh_token

c. Solução: 사용된 refresh_token이 가장 최근에 생성된 토큰인지 확인합니다. 항상 마지막으로 생성된 토큰을 사용해야 합니다.

### 3 - Error message: "Partner and shop has no linked."
a. Motivo: 호출에 사용된 shop_id가 partner_id와 연결되지 않은 경우에 발생합니다.

b. Parâmetro: shop_id e partner_id

c. Solução: 호출에 사용되는 상점 ID와 파트너 ID를 확인합니다.

### 4 - Error message: "This shop account has been banned. Permissions for shop authorization and API calls have been suspended until the shop account is restored."
a. Motivo: 상점이 금지되었습니다.

b. Parâmetro: N/A

c. Solução: 판매자에게 상점 상태를 확인합니다.

### 5 - Error message: "Wrong sign."
a. Motivo: 생성된 sign이 올바르지 않습니다.

c. 해결 방법: 판매자에게 판매자 센터에서 shop id에 해당하는 등록을 완료하도록 요청합니다.

## RefreshAccessToken 호출 시 반환되는 가장 일반적인 오류

### 1 - Error message: "Your refresh_token expired."
a. 이유: 일반적으로 refresh_token이 이미 만료된 경우 발생하며, 유효 기간은 30일입니다.

b. 파라미터: refresh_token

c. 해결 방법: 새 refresh_token이 없는 경우 인증 절차를 다시 수행해야 합니다.

### 2 - Error message: "Invalid refresh_token."
a. 이유: 사용된 refresh_token이 유효하지 않은 경우 발생합니다. 일반적으로 refresh_token이 이미 이전에 사용되었으며, 일회용 토큰입니다.

b. 파라미터: refresh_token

c. 해결 방법: 사용된 refresh_token이 가장 최근에 생성된 토큰인지 확인합니다. 항상 가장 최근에 생성된 토큰을 사용해야 합니다.

### 3 - Error message: "Partner and shop has no linked."
a. 이유: 호출에 사용된 shop_id가 partner_id와 연결되지 않은 경우 발생합니다.

b. 파라미터: shop_id 및 partner_id

c. 해결 방법: 호출에 사용 중인 shop id와 partner id를 확인합니다.

### 4 - Error message: "This shop account has been banned. Permissions for shop authorization and API calls have been suspended until the shop account is restored."
a. 이유: 상점이 금지되었습니다.

b. 파라미터: N/A

c. 해결 방법: 판매자에게 상점 상태를 확인합니다.

### 5 - Error message: "Wrong sign."
a. 이유: 생성된 sign이 올바르지 않습니다. 일반적으로 sign 생성 시 path가 올바르지 않습니다.

b. 파라미터: sign

c. 해결 방법: sign 문자열 조립 시 path가 올바른지 확인합니다: "/api/v2/auth/access_token/get".

### 6 - Error message: "No permission. Please inform seller to complete the Seller Registration on Shopee Seller Center first, then this shop can call for this API."
a. 이유: 판매자가 아직 판매자 센터에서 등록을 완료하지 않았습니다.

b. 파라미터: N/A

c. 해결 방법: 판매자에게 판매자 센터에서 shop id에 해당하는 등록을 완료하도록 요청합니다.

### 7 - Error message: "error params".
a. 이유: Request body/payload에 잘못된 파라미터 이름이 있습니다.

b. 파라미터: refresh_token, shop_id, partner_id

c. 해결 방법: request에 사용된 파라미터 중 어떤 것이 잘못되었는지 확인하고 수정합니다.

## 사용 사례

1. 마켓플레이스를 Shopee API와 통합
2. Shopee 리소스에 액세스하기 위해 스토어 인증
3. 사용자 로그인을 위한 인증 URL 생성
4. API 호출을 위한 액세스 토큰 획득
5. 지속적인 API 액세스를 위한 액세스 토큰 갱신

## 관련 API

- GetAccessToken

---

## 원문 (English)

### Summary

This guide provides instructions on how to create the authorization URL for a store, generate the authorization sign, and complete the store authorization process. It also covers how to obtain access tokens and refresh tokens using the GetAccessToken API.

### Content

# API Best Practices - Brand Open API Best Prac... > Open API Auth call / Pass...

## 1.1) Como criar a URL de autorização de loja?

- É a URL que seu registro o usuário, precisara fazero login na plataforma e selecionar um signup
- Note
- Acredite-se: https://partner.shopeemobile.com
- Partner_id: você recebe dos [hotéis] do seu App
- Redirect_url: Você definiu na sua aplicação. Consultar a documentação desse de consultar a autorização.
- **Utilize o cliente "criando" para desenvolvimento**
- **Utilize o cliente "vendido" para mercado** após passar pela integração
- Sign: string criptografada através segundo no página do primeiro item

**Exemplo:** Fornecemos como que configure a mensagem com algoritmo HMAC-SHA256

```
https://partner.shopeemobile.com/api/v2/shop/auth_partner?partner_id=184644&redirect=https://google.com.br&lang=pt-BR&ga2d3c32dd53c39bcr92c37d2b6c2f42bb6t=1724526317&sign=...
```

**Autorização Loja Test:**

```
https://partner.test-stable.shopeemobile.com/api/v2/shop/auth_partner?partner_id=184644&redirect=https://google.com.br&lang=pt-BR&ga2d3c32dd53c39bcr92c37d2b6c2f42bb6t=1724526317&sign=...
```

## 1.2) Como gerar o sign de autorização de loja?

Para use de Login, Gá que uma série partindo "p", para concatenar: broca poder e os campos usados na construção.

Para as outras APIs, use o path completo após o domínio (incluindo a barra) para concatenar.

Exemplo: Fornecemos como que configure a mensagem com algoritmo HMAC-SHA256

**Observação:** Todos os valores abaixo são fictícios.

Para gerar URL de autorização:

Basta preencher os campos da forma correta:

Concatene como: /api/v2/shop/auth_partner + "p", para manter o "Sign integral" do Roadmap, para gerar o sign.

---

### Code Examples:

```javascript
// Define partner_id
let partner_id = XXXXX;
let path = "/api/v2/shop/auth_partner";

// Content preparation (key-value pairs)
let timestamp = Math.floor(Date.now() / 1000);
let redirect = "https://your-redirect-url.com";
let base_string = `${partner_id}${path}${timestamp}`;

// Generate sign with HMAC-SHA256
let sign = crypto.createHmac('sha256', partner_secret)
              .update(base_string)
              .digest('hex')
              .toUpperCase();

// Redirect URL with params
let authUrl = `https://partner.shopeemobile.com${path}?partner_id=${partner_id}&redirect=${encodeURIComponent(redirect)}&timestamp=${timestamp}&sign=${sign}`;
```

```python
# Python example
import hmac
import hashlib
import time

partner_id = XXXXX
path = "/api/v2/shop/auth_partner"
timestamp = int(time.time())
redirect = "https://your-redirect-url.com"

# Base string
base_string = f"{partner_id}{path}{timestamp}"

# Generate sign
sign = hmac.new(
    partner_secret.encode('utf-8'),
    base_string.encode('utf-8'),
    hashlib.sha256
).hexdigest().upper()

# Auth URL
auth_url = f"https://partner.shopeemobile.com{path}?partner_id={partner_id}&redirect={redirect}&timestamp={timestamp}&sign={sign}"
```

```java
// Java example
import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;

String partnerId = "XXXXX";
String path = "/api/v2/shop/auth_partner";
long timestamp = System.currentTimeMillis() / 1000;
String redirect = "https://your-redirect-url.com";

// Base string
String baseString = partnerId + path + timestamp;

// Generate sign
Mac sha256Hmac = Mac.getInstance("HmacSHA256");
SecretKeySpec secretKey = new SecretKeySpec(partnerSecret.getBytes(), "HmacSHA256");
sha256Hmac.init(secretKey);
String sign = bytesToHex(sha256Hmac.doFinal(baseString.getBytes())).toUpperCase();

// Auth URL
String authUrl = String.format(
    "https://partner.shopeemobile.com%s?partner_id=%s&redirect=%s&timestamp=%d&sign=%s",
    path, partnerId, URLEncoder.encode(redirect, "UTF-8"), timestamp, sign
);
```

---

## 1.3) Fazendo a autorização da loja

Após gerar a URL de autorização, será necessário redirecioná-lo ao seu site de marketplace de verificação com seguintes parâmetros: front_s_part_url, s_redirect e s_partnerId (para controle de acesso interno e segurança na autorização).

Depois de uma via URL de autorização, o vendedor é imediato após inicialmente (se já acessou e retornou a senha e a senha é certa, a função é passada para a URL do redirect criada).

**Observação:** A autorização precisa ser feita uma única vez para a loja.

- Se ele não tiver login e senha, um novamente erro será "repetition" baseando e processos inicial como uma...
- Este não URL de autorização pode ser reutilizada após o usuário fazer a autorização inicial
- O usuário deverá ter permissão de "Main Account" ou ter a permissão de "Authorization"
- O redirecionamento deve ser usado para confirmar a integração de sua loja completação dá redirect
- Credenciais de acesso são controladas via OAuth2 e requer access tokens e refresh tokens
- E obrigatório armazenar a access_token e refresh_token na sua plataforma para fazer suas
- Depois de 4 horas que um authorization é não APP
- Veja um exemplo de teste: https://test-stable.ap_1_open.br/ code=4630599fa0018149d7f0fb31dcf44da9_c91345f

---

## 2.GetAccessToken

Se a Shopee teve seus credenciais e armazá-o no registro da autorização de loja para obter o access_token e o refresh_token.

O access_token (AccessToken) tem um prazo de validade e exige, no momento da autorização da loja é preciso fazer acesso, basta-e obter o acesso...

**Informações relevantes sobre a chamada do GetAccessToken:**

[THIS IS DIAGRAM: A flowchart showing the authorization process with multiple steps including "Criação de autorização", "Login e autorização", "Callback após a autorização", "Processo de access token", and various decision points and API calls]

---

## Parameters Table:

| Parâmetro | Exemplo | Obrig/Opc/Obrigatório |
|-----------|---------|----------------------|
| Host | https://partner.shopeemobile.com | Live environment |
| | https://partner.test-stable.shopeemobile.com | Test/stable/sandbox environment |
| Exemplo_id | XXXX | Informação + login/cliente: Open Partner Platform > APP > Obtenha seu s partnerId |
| Shop_id | XXXX | Shop_id via token de autenticação via API |
| Bound Key | Use shop/auth/sign | Use environment Oposite na Open Partner Console Informação + APP > Bound Key |
| Not Key | | Not rotado=random environment Informação + APP > Not Key |
| Sign | Ex:1232356237764cdc51829c8746c2ca3166e94920cb816df2473f3cb2d | |

---

# API Documentation - GetAccessToken

## Key Information Table

| Field | Value | Description |
|-------|-------|-------------|
| **Test Key** | - | Test State=Sandbox (environment: informação disponível na Open Platform → Painel de Controle → API → Live Key) |
| **Path** | `b/12333063677b/4/9083b03ee7393b0e38199f0991020959447973cb3` | String (criptografado = HMAC-SHA256) |
| **Code** | `6dae85447fa11ec6cb8e6e274857d8` | Gerado após a autorização da loja |

### Notes
1. Order a string base para gerar o login com a assinatura especificada no valendo: pattern_id + path + temperatura XXXX/xxxxxxxxxxxxxxxxxxxxxxxx11699902313
2. As lojas de API são usadas para gerar assinaturas. Para ambientes de teste/sandbox, use o Test Key; para o ambiente de produção, use o Live Key.

---

## Example 1: POST Request

### Endpoint
```
POST
```

### URL
```
https://partner-test-api.jp.shopee.in.co/api/v2/auth/access_token/get
```

### Headers
```http
Content-Type: application/json
```

### Request Body
```json
{
  "app_id": 1000XXXX,
  "code": "xxxxxxxx",
  "shop_id": 61XXXXX,
  "partner_id": 100XXXX
}
```

---

## Example 2: cURL Command

```bash
curl --request POST
```

```bash
https://partner-test-api.jp.shopee.in.co/api/v2/auth/access_token/get
```

### Request Parameters
```bash
--header 'Content-Type: application/json'
```

```bash
--data-raw '{
  "shop_id":"61XXXXX",
  "partner_id":"100XXXX"
}'
```

---

## Example 3: Response for GetAccessToken

### Request Body Example
```json
{
  "refresh_token": "7552490145567cccccc4444dd1b52d69",
  "access_token": "5a62355054c7f09ceeeee6677666a7",
  "expire_in": 14400,
  "request_id": "ef87a2a5b916515cccceee4EuuzpGC",
  "error": "",
  "message": ""
}
```

### Important Notes
- access_token: Usado apenas uma vez por cada call; válido por 4 horas
- expire_in: Válido do sistema integrado e arquivo, lesser através de um cartão
- refresh_token: Usado para chamar a API RefreshAccessToken, É usado na chamada RefreshAccessToken para obter novo access_token (validade: 30 dias)

---

## 3 | RefreshAccessToken

Depois que o access_token expirar, deverá utilizando a chamada de RefreshAccessToken, a chamada de gerar uso adequada, pelo código do autorização da loja será definitivamente autorizado. Assim, na verdade se o access_token é disponível. Além disso, refresh_token pode ser atualizado apenas; uma única vez a lista utilizado de 30 dias. Caso o novo refresh_token não seja utilizado dentro de 30 dias, a conexão com a loja expiratá e será necessário reimportação.

### Information Requirements
- Informação necessária para a chamada de RefreshAccessToken

---

## Parameter Table

| Field | Value | Description | Endpoint |
|-------|-------|-------------|----------|
| **Host** | https://partner.test.shopee.com.cn | Test environment | - |
| | https://partner-test-api.br.shopee.com | Test State=Sandbox (environment: informação disponível na Open Platform → Painel de Controle → API → Live Key) | - |
| **Partner_id** | XXXX | - | - |
| **Shop_id** | XXXX | Shop_id do loja get no momento da autorização e APT | - |
| **Path** | /api/v2/auth/access_token_get | - | - |
| **Secret Key** | Live Key | Live (production environment: informação disponível na Open Platform → Painel de Controle → API → Live Key) | - |
| **Test Key** | - | Test State=Sandbox (environment: informação disponível na Open Platform → Painel de Controle → API → Live Key) | - |
| **refresh_token** | b/12333063677b/4/9083b03ee7393b0e38199f0991020959447973cb3 | String (criptografado = HMAC-SHA256) | - |
| | 5267480146672cccc14444d1b52d69 | Gerado na chamada anterior GetAccessToken, atualizado dentro chamadas da RefreshAccessToken (validade: 30 dias) | - |

### Important Notes
1. Order a string base para gerar o login com a assinatura especificada no valendo: pattern_id + path + temperatura XXXX/xxxxxxxxxxxxxxxxxxxxxxxx11699902313
2. As lojas de API são usadas para gerar assinaturas para o seu código Key para ambiente de teste/sandbox
3. Exemplo de resposta para a chamada de RefreshAccessToken

---

## Example 4: cURL Request

```bash
curl --request POST
```

### URL
```bash
https://partner-test-api.jp.shopee.in.co/api/v2/auth/access_token/get
```

### Request Body
```json
{
  "refresh_token": "7552490145567cccccc4444dd1b52d69",
  "shop_id": 61XXXXX,
  "partner_id": 100XXXX
}
```

### Response
```json
{
  "request_id": "xxx",
  "error": "",
  "message": ""
}
```

---

## Common Errors Returned in GetAccessToken Call

### Error Types:
- **error_auth**: Authentication errors
- **error_param**: Parameter validation errors  
- **error_permission**: Access permission errors

### Common Error Details:
- **Motivo**: Ocorre geralmente quando o código enviado na API
- **Problema**: Invalid token or missing token
- **Motivo**: geralmente é quando o código do usuário está vencida
- **Erro mensagem**: "Invalid partner" or similar
- **Motivo**: Generally when the ID or credentials are incorrect
- **Erro mensagem**: "Invalid shop_id"
- **Motivo**: Generally quando o ID da loja não coincide ou está vazio

---

**Note**: All code examples and endpoints shown are for testing/sandbox environment. Replace with production URLs and credentials when deploying to live environment.

---

# API Documentation - Token Management

## Token Types

- **access_token**: pode ser usado mais de uma vez, válido por 4 horas.
- **expire_in**: depois de quantos segundos o access_token deixará de ser válido;
- **refresh_token**: só pode ser usado uma vez e é válido por 30 dias.

---

## Entenda os erros mais comuns retornados na chamada ao GetAccessToken

### 1 - Error message: "Invalid code".
a. Motivo: Ocorre, geralmente, quando o código enviado já foi utilizado antes.

b. Parâmetro: code

c. Solução: Gerar um novo código através do processo de autorização, ou utilizar o código correto.

### 2 - Error message: "Invalid partner id".
a. Motivo: Ocorre quando o partner id utilizado não é válido.

b. Parâmetro: partner_id

c. Solução: Verificar o partner id que está sendo utilizado na chamada.

### 3 - Error message: "Invalid shop id".
a. Motivo: Ocorre quando o shop id utilizado não corresponde com a loja que gerou o código.

b. Parâmetro: shop_id

c. Solução: Verificar o shop id que está sendo utilizado na chamada.

### 4 - Error message: "Invalid timestamp".
a. Motivo: Timestamp já passou da validade de 5 minutos.

b. Parâmetro: timestamp

c. Solução: Utilizar um timestamp dentro da sua validade (5 minutos).

### 5 - Error message: "Wrong sign.".
a. Motivo: O sign gerado está incorreto, geralmente o path na hora de gerar o sign está incorreto.

b. Parâmetro: sign

c. Solução: Verificar se o path na montagem da string do sign está correto: "/api/v2/auth/token/get".

### 6 - Error message: "No permission. Please inform seller to complete the Seller Registration on Shopee Seller Center first, then this shop can call for this API.".
a. Motivo: O vendedor ainda não finalizou o cadastro na Central do Vendedor.

b. Parâmetro: N/A

c. Solução: Solicitar ao vendedor que finalize o cadastro correspondente ao shop id na Central do Vendedor.

---

## Erros mais comuns retornados na chamada ao RefreshAccessToken

### 1 - Error message: "Your refresh_token expired."
a. Motivo: Ocorre, geralmente, quando o refresh_token já expirou, ele tem validade de 30 dias.

b. Parâmetro: refresh_token

c. Solução: Caso não tenha um refresh_token novo, deverá realizar o procedimento de autorização novamente.

### 2 - Error message: "Invalid refresh_token."
a. Motivo: Ocorre quando o refresh_token utilizado não é válido, geralmente o refresh_token já foi utilizado anteriormente, ele é um token de uso único.

b. Parâmetro: refresh_token

c. Solução: Verificar se o refresh_token utilizado é o token gerado mais recente. Deve sempre utilizar o último gerado.

### 3 - Error message: "Partner and shop has no linked."
a. Motivo: Ocorre quando o shop_id utilizado na chamada não está conectado com o partner_id.

b. Parâmetro: shop_id e partner_id

c. Solução: Verificar o shop id e o partner id que estão sendo utilizados na chamada.

### 4 - Error message: "This shop account has been banned. Permissions for shop authorization and API calls have been suspended until the shop account is restored."
a. Motivo: Loja está banida.

b. Parâmetro: N/A

c. Solução: Verificar status da loja com o seller.

### 5 - Error message: "Wrong sign."
a. Motivo: O sign gerado está incorreto, geralmente o path na hora de gerar o sign está incorreto.

b. Parâmetro: sign

c. Solução: Verificar se o path na montagem da string do sign está correto: "/api/v2/auth/access_token/get".

### 6 - Error message: "No permission. Please inform seller to complete the Seller Registration on Shopee Seller Center first, then this shop can call for this API."
a. Motivo: O vendedor ainda não finalizou o cadastro na Central do Vendedor.

b. Parâmetro: N/A

c. Solução: Solicitar ao vendedor que finalize o cadastro correspondente ao shop id na Central do Vendedor.

### 7 - Error message: "error params".
a. Motivo: Request body/payload com nome de parâmetro incorreto.

b. Parâmetro: refresh_token, shop_id, partner_id

c. Solução: Verificar qual dos parâmetros utilizado no request está incorreto e corrigir.

---

**문서 ID**: developer-guide.385
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/385
**처리 완료**: 2025-10-16T09:23:54
