# 인증 프로세스

**카테고리**: 인증
**난이도**: 중간
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T08:14:33

## 개요

본 가이드는 인증 링크 생성, Klaytn으로부터의 인증 활성화, 인증 링크 사용을 포함한 인증 프로세스를 설명합니다. 인증 준비 단계, 인증 프로세스 자체, 그리고 검색된 정보를 사용하여 파트너 서비스에 인증을 전달하는 방법을 자세히 설명합니다. 또한 PHP, Python, Java 및 Node.js로 작성된 코드 예제를 제공합니다.

## 주요 키워드

- authorization
- authorization link
- Klaytn
- Open Platform
- HMAC-SHA256
- API
- app parameter
- signature
- partner_ci
- timestamp

## 본문

# 인증 프로세스

인증을 생성하는 데는 세 단계가 있습니다. 인증 링크 생성, (Klaytn에서 인증 활성화), 인증 링크 사용, 마지막으로 e-card, klaen을 통해 최종 사용자가 인증하도록 링크를 전달하는 것입니다.

아래에서 자세히 설명하며, 인증 전에 준비해야 할 사항, 인증 프로세스에 대한 세부 정보, 인증 후 검색된 정보를 사용하여 파트너 서비스의 인증을 전달하는 방법에 대한 정보도 함께 제공합니다.

## 인증 링크 생성

klaen 인증을 진행하기 전에 Open Platform App-tier-link Authorizer에 요청을 보내야 합니다. 인증 링크를 생성하기 위해 URL을 블랙리스트*에 등록합니다.

### 인증 요청 콘솔

| 필드 | 설명 |
|-------|-------------|
| App List | 앱 목록 |
| Provider | 제공자 |
| API Not Set | API 설정 안 됨 |
| API Not Set | API 설정 안 됨 |
| API Not Set | API 설정 안 됨 |
| Trade List | 거래 목록 |

**참고:**
- App List: 0
- 통합할 앱 4개 목록
- API Not Set (3) - 나열할 필요 없음 / 공백
- Trade List - 허용 / 거부 필요

모든 유형의 앱에 대해 다음 사양으로 인증 링크를 생성해야 합니다. 인증 링크는 app_key 및 기타 필수 매개변수의 세 가지 기본 매개변수로 구성됩니다.

#### 인증 링크 생성 https://approvalappname.clouckorea/ch/klayappli_partner

**필수 기본 매개변수:**

- app_key: string (설명: 앱 키; 없는 경우 아래 매개변수에 앱 키가 있습니다.)

**Baselink https (매개변수):**

- partner_ci: string (설명: App-table master according2/Klayappli_partner)
- partner_info: string (설명: 파트너 정보는 20자이며 app_identifier_key를 생성하는 데 사용됩니다.)

#### 기타 필수 매개변수:

| 매개변수 이름 | 유형 | 설명 |
|----------------|------|-------------|
| login | string | sign base string (signin_partner_id, api path, timestamp)으로 얻은 서명입니다. API 키로 생성해야 합니다. |
| partner_ci | int | 앱에서 얻은 partner_ci입니다. |
| timestamp | int | sign base string에 사용된 타임스탬프를 입력할 때의 시간 값입니다. 타임스탬프 값은 3분 동안만 유효합니다. |
| redirect | string | 인증이 완료된 후 메시지가 리디렉션되는 URL입니다. 이는 시스템의 엔드포인트 또는 서비스를 호출하는 Open Platform의 엔드포인트일 수 있습니다. |

#### 앱 매개변수 계산

앱 매개변수는 인증 링크의 일부가 아니거나 해당하지 않지만 앱이 Open Platform API를 호출할 때마다 인증에 사용되는 매개변수이기도 합니다. 앱 매개변수를 계산하려면 인증 서명을 활성화하는 해시 값(HMAC-SHA256)이 필요합니다.

앱 매개변수 base string의 경우:
sign base string을 활성화하기 위해 다른 매개변수 방식으로 base string을 생성해야 합니다(일반 매개변수와 호환 가능).

- Shop API의 경우 partner_ci, api path, timestamp, access_token, shop_id.
- Partner API의 경우 partner_ci, api path, timestamp, access_token, token_type_shop&shop_ci.
- Public API의 경우 partner_ci, api path, timestamp.

HMAC-SHA256을 사용하여 sign base string을 해싱하고 파트너 키를 암호화 키로 사용합니다. 해시된 값(App 값)은 인증 서명입니다.

### PHP 코드 데모:

```php
<?php
// Input 1 Text:
$input1 = input1
$input2 = input2
$signPlainText = "";

if ($_SERVER["REQUEST_METHOD"] === "POST") {
    $input1 = isset($_POST['input1']) !== "";
    $text = "https://approvalappname.klayout.com/s/ch";
    $signinKey = $POST;
    $signPlainText = "";
    
    // Calculate SignBase:
    foreach($input_text = "header" . $var[len, id, path, $time] 
          $signPlainText = $prefix.join("&".$var.kind));
    
    $signedText = hash_hmac('SHA256', $headerText);
    $signedText = $signPlainText("text:". $signedText. "&cipher=text:". $partner, $POST);
    
    $finalText = $_GET['?'] ?? "auth/" . $signedText;
    echo "<ul>";
    foreach ($text as $item) {
        echo "<li>$item</li>";
    }
    echo "</ul>";
}
?>
```

### Python 코드 데모:

```python
import hashlib
import hmac

# Function to calculate signature
$api_path = 0

def calculate_signature(partner_id, api_path, timestamp):
    """Calculate HMAC signature"""
    sign_base = f"{partner_id}&{api_path}&{timestamp}"
    
    # Using HMAC-SHA256
    signature = hmac.new(
        bytes(api_key, 'utf-8'),
        bytes(sign_base, 'utf-8'),
        hashlib.sha256
    ).hexdigest()
    
    return signature

# Example usage
partner_id = "your_partner_id"
api_path = "/api/v1/resource"
timestamp = int(time.time())

signature = calculate_signature(partner_id, api_path, timestamp)
print(f"Signature: {signature}")
```

### Java 코드 데모:

```java
import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.security.InvalidKeyException;
import java.security.NoSuchAlgorithmException;

public class AuthorizationExample {
    
    public static String calculateSignature(String partnerId, String apiPath, long timestamp, String apiKey) 
            throws NoSuchAlgorithmException, InvalidKeyException {
        
        // Create sign base string
        String signBase = partnerId + "&" + apiPath + "&" + timestamp;
        
        // Use HMAC-SHA256
        Mac sha256Hmac = Mac.getInstance("HmacSHA256");
        SecretKeySpec secretKey = new SecretKeySpec(apiKey.getBytes(), "HmacSHA256");
        sha256Hmac.init(secretKey);
        
        byte[] hash = sha256Hmac.doFinal(signBase.getBytes());
        
        // Convert to hex string
        StringBuilder hexString = new StringBuilder();
        for (byte b : hash) {
            String hex = Integer.toHexString(0xff & b);
            if (hex.length() == 1) hexString.append('0');
            hexString.append(hex);
        }
        
        return hexString.toString();
    }
    
    public static void main(String[] args) {
        try {
            String partnerId = "your_partner_id";
            String apiPath = "/api/v1/resource";
            long timestamp = System.currentTimeMillis() / 1000;
            String apiKey = "your_api_key";
            
            String signature = calculateSignature(partnerId, apiPath, timestamp, apiKey);
            System.out.println("Signature: " + signature);
            
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
```

### Node.js 코드 데모:

```javascript
const crypto = require('crypto');

function calculateSignature(partnerId, apiPath, timestamp, apiKey) {
    // Create sign base string
    const signBase = `${partnerId}&${apiPath}&${timestamp}`;
    
    // Calculate HMAC-SHA256
    const hmac = crypto.createHmac('sha256', apiKey);
    hmac.update(signBase);
    const signature = hmac.digest('hex');
    
    return signature;
}

// Example usage
const partnerId = 'your_partner_id';
const apiPath = '/api/v1/resource';
const timestamp = Math.floor(Date.now() / 1000);
const apiKey = 'your_api_key';

const signature = calculateSignature(partnerId, apiPath, timestamp, apiKey);
console.log('Signature:', signature);

// Generate authorization URL
const authUrl = `https://approvalappname.cloudkorea.kr/ch/klayappli_partner?partner_ci=${partnerId}&timestamp=${timestamp}&login=${signature}&redirect=${encodeURIComponent('your_redirect_url')}`;
console.log('Authorization URL:', authUrl);
```

---

# Shopee Open Platform - 인증 가이드

## API 엔드포인트 정보

```
Server: partner-test.shopeemobile.com (샌드박스 환경용)
```

### 인증 링크 예시

**Path:** `/api/v2/shop/auth_partner`

**Endpoint:** `https://partner.test-stable.shopeemobile.com/api/v2/shop/auth_partner`

**Request Parameters:**
- `partner_id` (int)
- `redirect` (string): `https://www.YOUR_DOMAIN.com/callback`
- `sign` (string): `partner_key`와 `partner_id`, `path`, `timestamp`로 구성된 base string을 사용하여 HMAC-SHA256으로 계산됩니다.

**Base String Format:**
```
https://partner.test-stable.shopeemobile.com/api/v2/shop/auth_partner
```

**Partner ID:** `847363`

**Signature:** `5f7dcd4cbe1a45f1bcb2e6dcfd8f3f6e08c2b9ed9bb6d62aaf3a5c69f7a8c49f7`

### 인증 링크 예시

**프로덕션 및 샌드박스 테스트 환경용 링크:**

#### 프로덕션 환경:
```
https://partner.shopeemobile.com/api/v2/shop/auth_partner?partner_id=YOUR_PARTNER_ID&redirect=YOUR_REDIRECT_URL&sign=GENERATED_SIGNATURE
```

**예시:**
```
https://partner.shopeemobile.com/api/v2/shop/auth_partner?partner_id=1000&redirect=https%3A%2F%2Fwww.yourwebsite.com%2Fcallback&timestamp=1594808000&sign=9c140756e3f1e2f0c5e87434b8c8b3c4be7c9f3d9e8f5b6a4d3c2e1f0a9b8c7d6e5f4
```

#### 샌드박스 환경:
```
https://partner.test-stable.shopeemobile.com/api/v2/shop/auth_partner?partner_id=YOUR_PARTNER_ID&redirect=YOUR_REDIRECT_URL&sign=GENERATED_SIGNATURE
```

**예시:**
```
https://partner.test-stable.shopeemobile.com/api/v2/shop/auth_partner?partner_id=1000&redirect=https%3A%2F%2Fwww.yourwebsite.com%2Fcallback&timestamp=1594808000&sign=9c140756e3f1e2f0c5e87434b8c8b3c4be7c9f3d9e8f5b6a4d3c2e1f0a9b8c7d6e5f4
```

### 샌드박스 테스트 환경:
```
https://partner.test-stable.shopeemobile.com/api/v2/shop/auth_partner?partner_id=1000&redirect=https%3A%2F%2Fwww.yourwebsite.com%2Fcallback&timestamp=1594808000&sign=9c140756e3f1e2f0c5e87434b8c8b3c4be7c9f3d9e8f5b6a4d3c2e1f0a9b8c7d6e5f4
```

---

## 인증 프로세스

### 📌 참고
판매자가 5분 동안 로그인 및 인증 유효성을 유지하는지 확인하십시오. 타임스탬프와 사인이 만료되면 판매자는 인증을 계속할 수 없습니다. 인증 링크는 더 이상 유효하지 않으며 판매자를 새 링크로 리디렉션해야 합니다.

---

## Shop(s)에서 인증 획득

판매자와 함께 인증 링크를 생성한 후 판매자는 계정에 로그인해야 합니다. 그러나 인증 코드 사용의 경우 판매자는 이메일을 확인하고 인증 코드를 입력해야 합니다.

판매자는 기본 계정을 사용하여 단일 상점을 인증하거나 기본 계정을 사용하여 여러 소매업체/상점을 인증할 수 있습니다. 계정은 Shopee Seller Center 또는 인증 페이지에 로그인하는 데 사용할 수 없습니다.

### 📌 참고
참고: 사내 시스템 앱 사용자는 기본 상점의 데이터에만 액세스하도록 인증할 수 있습니다. 개발자 계정 탭에도 로그인해야 합니다.

---

## 기본 계정에서 인증

**1단계:** 판매자는 로그인 세부 정보를 입력하고 로그인을 선택합니다.

---

## Shopee Open Platform APP 인증을 위한 로그인

**표시되는 인터페이스:**
- 이메일/전화 입력 필드
- 비밀번호 입력 필드
- "비밀번호를 잊으셨습니까?" 링크
- "로그인" 버튼 (주황색)
- "기본 계정으로 돌아가기" 옵션
- 언어 선택기: 영어

---

**2단계:** 판매자는 휴대폰으로 전송된 인증 코드를 입력하고 확인을 선택합니다.

---

## 확인

**표시되는 인터페이스:**
- 인증 코드 입력 필드
- "확인" 버튼 (주황색)
- "영어" 언어 선택기

---

**3단계:** 로그인하면 판매자는 인증 확인을 선택합니다.

---

## 인증 화면

**Shopee Open Platform**

### 인증
[상점 이름 및 세부 정보가 여기에 표시됩니다]

**인터페이스 요소:**
- 인증 확인 버튼
- 부여된 권한에 대한 정보
- 데이터 액세스에 대한 세부 정보

**나열된 권한:**
- 주문 정보: 주문에 대한 모든 기록과 관련된 정보
- 제품 정보: 제품 및 해당 카테고리에 대한 모든 기록과 관련된 정보
- 상점 정보: 상점 설정 및 세부 정보와 관련된 정보
- 기타 정보: 할인 및 기타 세부 정보와 관련된 정보
- 물류: 물류 및 배송에 대한 모든 기록과 관련된 정보
- 금융: 결제 및 계정 잔액과 관련된 정보
- 반품: 모든 기록에 대한 반품 및 분쟁과 관련된 정보
- 미디어 공간: 이미지 및 비디오 업로드와 관련된 정보

**전체 인증을 보려면 다음 링크를 클릭하십시오:**
```
인증을 확인하려면 클릭하십시오
```

---

**4단계:** 인증 후 프런트 엔드 페이지는 인증 링크의 리디렉션 URL로 리디렉션됩니다.

```
https://www.YOUR_DOMAIN.com/callback?code=xxxxxxxx&shop_id=xxxxxx
```

---

## 기본 계정에서 인증

**1단계:** 기본 계정에 로그인하려면 판매자는 로그인 페이지에서 기본 계정으로 전환 링크를 선택합니다.

---

## Shopee Open Platform APP 인증을 위한 로그인

**표시되는 인터페이스:**
- SH → ID 드롭다운 (마켓/지역 선택기)
- 이메일/ID/USERNAME/Mobile/[?]2017 입력 필드
- 비밀번호 입력 필드
- "비밀번호를 잊으셨습니까?" 링크
- "로그인" 버튼 (주황색)
- "OR" 구분선
- "하위 계정으로 전환" 옵션
- 언어 선택기: 영어

---

**2단계:** 그런 다음 판매자는 로그인 세부 정보를 입력하고 로그인을 선택할 수 있습니다.

---

## Shopee Open Platform APP 인증을 위한 로그인

**표시되는 표준 로그인 양식**
- 이메일/사용자 이름/휴대폰 입력
- 비밀번호 필드
- 로그인 버튼
- 계정 전환 옵션

---

**추가 참고 사항 및 문서 링크:**

전체 인증 흐름은 애플리케이션에 부여된 권한에 대한 판매자 제어를 유지하면서 상점 데이터에 대한 안전한 액세스를 보장합니다.

---

# Shope Xpaygateway API 인증을 위한 로그인

## 단계

### 1. 이 판매자는 인증해야 하는 상점을 선택합니다.
**참고:** 여러 제휴 판매자가 있는 경우 판매자의 API를 호출해야 하는 경우 판매자를 다시 선택하여 Auth Merchant 확인란을 선택하십시오.

---

### 2. [276] OperationEvent API 그룹

**Endpoint:**
- `GET /api/v1/orders/89879900046` - 인증됨
  - View Stages: Order List
  - Based of 2019

**Auth Merchant** 버튼 사용 가능

---

### 3. [276] OperationEvent API (1234512345)

**사용 가능한 API 작업:**
- `GET /api/v1/orders/89879900046` - 인증됨
- `GET /api/v1/orders/564001872` - 인증됨
- `GET /api/v1/orders/ORDER-12` - 인증됨
- `GET /api/v1/orders/564001872/cancel` - 인증됨

**Auth Merchant** 버튼 사용 가능

---

### 4. [276] OperationEvent API (1234512345) - #session - #common

**사용 가능한 작업:**
- `GET /tokens/v1/10004343266` - 인증됨
- `PUT /api/v1/17700000000321` - 인증됨

**Auth Merchant** 버튼 사용 가능

---

### 5. [276] OperationEvent API (1234512345) - #session - #common

**View Stages:** Order List
**Based of 2019**

**Auth Merchant** 버튼 사용 가능

---

## LOCAL (실제) 계정 인증

### 인증
사용 가능한 API: E-Shope API: Same - Online - Portlet (옵션)

**SN2**
**관리자 세부 정보**

**정보 상자:**
이 인증 코드는 테스트 문서에 대한 최초 인증에만 사용됩니다. 로그아웃하도록 인증된 페이지를 선택하십시오. 실제 페이지에 대한 해당 코드는 최초 인증에 사용해야 합니다. 인증 코드가 관련 없는 테스트 환경에서만 사용된 경우.

**[?] Ask: LOCAL (실제)**

---

## MTR 범위 (인증을 완료하려면 상점을 선택하십시오)

**Based all:**

- **1_a:** sell_[?] location
- **2_a:** show_[?] location
- **3_a:** sell_[?] location
- **4_a:** sold_[?] location
- **5_b:** store location
- **6_b:** (Multiple) show location
- **7a_b:** sell_a = sell + sold + combination

---

## 3. 이 판매자는 인증 확인을 선택하여 선택을 확인합니다.

---

### [276] OperationEvent API (1234512345) - #session - #common

**Auth Merchant** 버튼

**View Stages:** Order List
**Based of 2019**

---

### API 작업 목록:
- `GET /tokens/v1/10004343266` - 인증됨
- `PUT /api/v1/orders/89879900046` - 인증됨
- `GET /api/v1/orders/564001872` - 인증됨

**인증 확인** 버튼 (빨간색/주황색)

---

## 4. 인증 후 프런트 엔드 페이지는 인증 작업의 리디렉션 URL로 리디렉션됩니다.

### 리디렉션 URL
매개변수 code, shop_id 및 state 등을 포함합니다.

---

## 인증 코드 사용

인증 코드를 얻은 후 API를 호출하여 인증 코드를 콜백 주소 리디렉션 URL로 교환할 수 있습니다. 그런 다음 코드를 사용하여 처음으로 access_token을 얻을 수 있습니다.

기본 계정에서 인증이 완료된 경우 code와 main_account_id가 리디렉션 URL로 반환됩니다. 그렇지 않으면 코드는 shop_id 형식입니다.

### 매개변수 테이블:

| 매개변수 이름 | 유형 | 설명 |
|----------------|------|-------------|
| code | string | 이는 인증 콜백의 인증 코드입니다. 이 값은 access_token 및 refresh_token을 검색하는 데 사용됩니다. 한 번만 유효하며 10분 후에 만료됩니다. |
| shop_id | int | 앱에 방금 인증을 부여한 상점의 ID입니다. 상점에서 다른 인증이 완료된 후 반환됩니다. |
| main_account_id | int
- **1_a:** sell_[?] location
- **2_a:** show_[?] location
- **3_a:** sell_[?] location
- **4_a:** sold_[?] location
- **5_b:** store location
- **6_b:** (Multiple) show location
- **7a_b:** sell_a = sell + sold + combination

---

## 3. 판매자는 선택 사항을 확인하기 위해 인증 확인을 선택합니다.

---

### [276] OperationEvent API (1234512345) - #session - #common

**Auth Merchant** 버튼

**보기 단계:** 주문 목록
**2019년 기준**

---

### API 작업 목록:
- `GET /tokens/v1/10004343266` - Authorized
- `PUT /api/v1/orders/89879900046` - Authorized
- `GET /api/v1/orders/564001872` - Authorized

**Confirm Authorization** 버튼 (빨간색/주황색)

---

## 4. 인증 후 프런트 엔드 페이지는 인증 작업의 리디렉션 URL로 리디렉션됩니다.

### 리디렉션 URL
매개변수 code, shop_id 및 state 등을 포함합니다.

---

## 인증 코드 사용

인증 코드를 얻은 후 API를 호출하여 인증 코드를 콜백 주소 리디렉션 URL로 교환할 수 있습니다. 그런 다음 코드를 사용하여 처음으로 access_token을 얻을 수 있습니다.

인증이 기본 계정에서 수행된 경우 code 및 main_account_id가 리디렉션 URL로 반환됩니다. 그렇지 않으면 코드는 shop_id 형식으로 제공됩니다.

### 매개변수 테이블:

| 매개변수 이름 | 유형 | 설명 |
|----------------|------|-------------|
| code | string | 이는 인증 콜백의 인증 코드입니다. 이 값은 access_token 및 refresh_token을 검색하는 데 사용됩니다. 한 번만 유효하며 10분 후에 만료됩니다. |
| shop_id | int | 앱에 방금 인증을 부여한 상점의 ID입니다. 상점에서 다른 인증이 완료된 후 반환됩니다. |
| main_account_id | int | 상점에 방금 인증을 부여한 기본 계정의 ID입니다. 기본 계정에서 인증이 완료된 후 반환됩니다. |

---

## access_token 가져오기 및 갱신

Access_token은 API를 호출하는 자격 증명입니다. 각 사용자는 자체 access_token을 가지고 있습니다. 동일한 API의 경우 각 access_token은 4시간 동안 유효하며 4시간 이내에 여러 번 사용할 수 있습니다. 그러나 만료된 경우 refresh_token을 사용하여 access_token을 갱신하거나 새 인증 코드를 사용하여 새 access_token을 검색해야 합니다.

Refresh_token은 access_token을 갱신하는 데 사용되는 매개변수입니다. 각 refresh_token은 30일 동안 유효합니다.

각 shop_id 및 merchant_id의 access_token 및 refresh_token은 별도로 저장해야 합니다.

### Call/AccessToken

인증에 성공하면 code 및 shop_id 또는 main_account_id를 리디렉션 URL로 사용하여 API를 호출합니다. 이를 통해 access_token, refresh_token, merchant_id, access_token, token을 얻을 수 있습니다.

#### 경로:
- **프로덕션 환경:** https://partner.shopeemobile.com/api/v2/auth/token/get
- **테스트 환경:** https://partner.test-stable.shopeemobile.com/api/v2/auth/token/get
- **샌드박스 환경:** https://openplatform.sandbox.test-stable.shopeemobile.com/api/v2/auth/token/get

**요청 방법:** POST

---

### 요청 매개변수:

| 매개변수 이름 | 유형 | 필수 | 설명 |
|----------------|------|----------|-------------|
| sign | string | True | partner_key 및 timestamp를 순서대로 갖는 sign 기반 문자열(partner_id, api, path, timestamp, access_token 또는 refresh_token)로 계산된 서명 |
| partner_id | int | True | 앱에서 얻은 partner_id입니다. 이 partner_id는 첫 번째 쿼리를 전달해야 합니다. |
| timestamp | int | True | 타임스탬프, 5분 동안 유효 |

---

### 응답 매개변수:

| 매개변수 이름 | 유형 | 필수 | 설명 |
|----------------|------|----------|-------------|
| code | string | True | 인증 후 리디렉션 URL의 코드입니다. 한 번만 유효합니다. |
| partner_id | int | True | 앱에서 얻은 partner_id입니다. 이 partner_id는 기본 쿼리를 전달해야 합니다. |
| shop_id | int | True if [?] | 귀하에게 인증된 shop_id의 경우 응답의 shop_id 또는 merchant_id가 shop_id 또는 main_account_id인 경우 |
| main_account_id | int | True if [?] | 귀하에게 인증된 main_account_id의 경우 shop_id 또는 main_account_id를 입력 매개변수로 검색할 수 있는 경우 |

---

### 응답 매개변수:

| 매개변수 이름 | 유형 | 설명 |
|----------------|------|-------------|
| request_id | string | 또는 AUTH 요청; 항상 반환됩니다. 문제를 진단하는 데 사용됩니다. |

---

# API 문서

## 입력 매개변수

**참고:** `main_account_id`는 입력 매개변수로 생략할 수 있습니다.

**참고:** `main_account_id` 입력의 경우 `main_account_id` 범위 내에서 입력 매개변수로 선택할 수 있습니다.

### 요청 매개변수

| 매개변수 이름 | 유형 | 설명 |
|----------------|------|-------------|
| request_id | string | API 요청 ID, 항상 반환됩니다. 문제를 진단하는 데 사용됩니다. |
| error | string | API 요청에 대한 오류 코드, 항상 반환됩니다. 비어 있지 않은 문자열은 오류가 발생했음을 나타내고 빈 문자열은 성공을 나타냅니다. |
| refunds_failed | string | API 호출이 성공하면 반환됩니다. 각 shop_id 또는 각 merchant_id에 대해 30일 동안 실패한 총 환불을 나타냅니다. |
| access_token | string | API 호출이 성공하면 반환됩니다. 여러 번 사용할 수 있고 4시간 후에 만료되는 동적 토큰입니다. |
| expire_in | int | API 호출이 성공하면 반환됩니다. 토큰이 만료될 때까지의 시간을 초 단위로 나타냅니다. |
| message | string | 항상 반환됩니다. 자세한 오류 정보를 제공합니다. |
| new_user_time_to_list | INT | 입력 매개변수에 `main_account`가 있는 경우 반환되며, 이번에 기본 계정에서 인증된 모든 merchant_id를 포함합니다. |
| shop_id_list | INT | 입력 매개변수에 `main_account`가 있는 경우 반환되며, 이번에 기본 계정에서 인증된 모든 shop_id를 포함합니다. |

## Python 코드 데모

```python
import time
import hashlib

# Fill the request time
url = "/api/v2/shop/get_shop_info"
path = "/api/v2/shop/get_shop_info"
body = ""
partner_id = xxxxxx
partner_key = "xxxxxxxxxxxxxx"
shop_id = xxxxxx
merchant_id = xxxxxx

tmp_base_string = "%s%s%s%s%s%s" % (
    partner_id, path, str(int(time.time())), access_token, shop_id, merchant_id
)

base_string = tmp_base_string.encode()
partner_key_bytes = partner_key.encode()

sign = hmac.new(
    partner_key_bytes, base_string, hashlib.sha256
).hexdigest()

headers = {
    "Content-Type": "application/json",
    "Partner-Id": str(partner_id),
    "Timestamp": str(int(time.time())),
    "Access-Token": access_token,
    "Shop-Id": str(shop_id),
    "Merchant-Id": str(merchant_id),
    "Authorization": sign,
}

resp = requests.post(
    url=url, headers=headers, json=json.loads(body)
)
print(resp.content)
```

## Java 코드 데모

```java
import org.apache.commons.codec.binary.Hex;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.io.*;
import java.net.HttpURLConnection;
import java.net.URL;

public class AuthSignTest {
    
    public static void main(String[] args) throws Exception {
        int partner_id = xxxxxx;
        String partner_key = "xxxxxxxxxxxxxx";
        
        String path = "/api/v2/shop/auth_partner";
        String redirect_url = "https://www.xxxx.com/auth/callback";
        
        long timestamp = System.currentTimeMillis() / 1000L;
        
        String tmp_base_string = 
            String.format("%s%s%s", partner_id, path, timestamp);
        String base_string = tmp_base_string + partner_key;
        
        String sign = genHmacSHA256(base_string, partner_key);
        System.out.println("sign = " + sign);
        
        String url = 
            String.format(
                "https://partner.test.shopeedev.com%s?partner_id=%s&timestamp=%s&sign=%s&redirect=%s",
                path,
                partner_id,
                timestamp,
                sign,
                redirect_url
            );
        
        System.out.println(url);
    }
    
    public static String genHmacSHA256(String message, String secret) {
        try {
            Mac sha256_HMAC = Mac.getInstance("HmacSHA256");
            SecretKeySpec secret_key = new SecretKeySpec(secret.getBytes(), "HmacSHA256");
            sha256_HMAC.init(secret_key);
            
            byte[] bytes = sha256_HMAC.doFinal(message.getBytes());
            return Hex.encodeHexString(bytes);
            
        } catch (Exception e) {
            System.out.println("Error HmacSHA256");
        }
        return "";
    }
    
    public static String getWithHeaders(String host, String path, Map<String, String> headers) 
        throws IOException {
        
        HttpURLConnection connection = null;
        
        try {
            URL url = new URL(host + path);
            connection = (HttpURLConnection) url.openConnection();
            connection.setRequestMethod("GET");
            
            for (Map.Entry<String, String> entry : headers.entrySet()) {
                connection.setRequestProperty(entry.getKey(), entry.getValue());
            }
            
            int responseCode = connection.getResponseCode();
            
            if (responseCode == HttpURLConnection.HTTP_OK) {
                return getResponse(connection.getInputStream());
            } else {
                return getResponse(connection.getErrorStream());
            }
            
        } finally {
            if (connection != null) {
                connection.disconnect();
            }
        }
    }
    
    private static String getResponse(InputStream inputStream) throws IOException {
        BufferedReader in = new BufferedReader(new InputStreamReader(inputStream));
        String inputLine;
        StringBuilder response = new StringBuilder();
        
        while ((inputLine = in.readLine()) != null) {
            response.append(inputLine);
        }
        in.close();
        
        return response.toString();
    }
    
    public static String postWithHeaders(
        String host,
        String path,
        Map<String, String> headers,
        String body
    ) throws IOException {
        
        HttpURLConnection conn = null;
        
        try {
            URL url = new URL(host + path);
            conn = (HttpURLConnection) url.openConnection();
            conn.setRequestMethod("POST");
            conn.setDoOutput(true);
            
            for (Map.Entry<String, String> e : headers.entrySet()) {
                conn.setRequestProperty(e.getKey(), e.getValue());
            }
            
            OutputStream os = conn.getOutputStream();
            os.write(body.getBytes("UTF-8"));
            os.flush();
            
            int responseCode = conn.getResponseCode();
            
            if (responseCode == HttpURLConnection.HTTP_OK) {
                return getResp(conn.getInputStream());
            } else {
                return getResp(conn.getErrorStream());
            }
            
        } finally {
            if (conn != null) {
                conn.disconnect();
            }
        }
    }
    
    private static String getResp(InputStream is) throws IOException {
        BufferedReader in = new BufferedReader(new InputStreamReader(is));
        String inputLine;
        StringBuilder response = new StringBuilder();
        
        while ((inputLine = in.readLine()) != null) {
            response.append(inputLine);
        }
        in.close();
        
        return response.toString();
    }
}
```

---

# API 통합 가이드

## 인증

```php
$client = new GuzzleHttp\Client();
$response = $client->request('POST', 'https://api.example.com/auth', [
    'form_params' => [
        'username' => 'your_username',
        'password' => 'your_password'
    ]
]);
$token = json_decode($response->getBody())->token;
```

## API 호출하기

### GET 요청

```php
$response = $client->request('GET', 'https://api.example.com/data', [
    'headers' => [
        'Authorization' => 'Bearer ' . $token,
        'Accept' => 'application/json'
    ]
]);
$data = json_decode($response->getBody());
```

### POST 요청

```php
$response = $client->request('POST', 'https://api.example.com/create', [
    'headers' => [
        'Authorization' => 'Bearer ' . $token,
        'Content-Type' => 'application/json'
    ],
    'json' => [
        'name' => 'Item Name',
        'description' => 'Item Description',
        'price' => 99.99
    ]
]);
```

## 오류 처리

```php
try {
    $response = $client->request('GET', 'https://api.example.com/data');
} catch (\GuzzleHttp\Exception\ClientException $e) {
    echo "Client Error: " . $e->getMessage();
} catch (\GuzzleHttp\Exception\ServerException $e) {
    echo "Server Error: " . $e->getMessage();
}
```

## 응답 파싱

```php
$body = $response->getBody();
$data = json_decode($body, true);

foreach ($data['items'] as $item) {
    echo $item['name'] . "\n";
    echo $item['description'] . "\n";
}
```

## PHP 코드 데모

### 기본 GET 요청

```php
function getTasksByUser($userId, $userToken, $baseUrl) {
    $path = "/api/v1/users/tasks";
    
    $query = http_build_query([
        'user_id' => $userId,
        'status' => 'active'
    ]);
    
    $url = $baseUrl . $path . '?' . $query;
    
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, $url);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($ch, CURLOPT_HTTPHEADER, [
        'Authorization: Bearer ' . $userToken,
        'Accept: application/json'
    ]);
    
    $response = curl_exec($ch);
    $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
    curl_close($ch);
    
    return json_decode($response, true);
}
```

### 고급 POST 요청

```php
function createTaskForUser($userId, $userToken, $baseUrl, $taskData) {
    $path = "/api/v1/users/tasks";
    
    $url = $baseUrl . $path;
    
    $payload = array_merge([
        'user_id' => $userId,
        'created_at' => date('Y-m-d H:i:s')
    ], $taskData);
    
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, $url);
    curl_setopt($ch, CURLOPT_POST, true);
    curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($payload));
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($ch, CURLOPT_HTTPHEADER, [
        'Authorization: Bearer ' . $userToken,
        'Content-Type: application/json',
        'Accept: application/json'
    ]);
    
    $response = curl_exec($ch);
    $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
    curl_close($ch);
    
    return [
        'status' => $httpCode,
        'data' => json_decode($response, true)
    ];
}
```

## 속도 제한

**중요:** API는 토큰당 분당 100개의 요청으로 속도 제한이 있습니다.

### 속도 제한 헤더 예시

```
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 95
X-RateLimit-Reset: 1609459200
```

## 웹훅

웹훅을 구성하여 실시간 알림을 받습니다.

```php
// Webhook endpoint example
$payload = file_get_contents('php://input');
$signature = $_SERVER['HTTP_X_SIGNATURE'];

$calculated = hash_hmac('sha256', $payload, $webhookSecret);

if (hash_equals($signature, $calculated)) {
    $data = json_decode($payload, true);
    // Process webhook data
}
```

## 모범 사례

1. 처리하기 전에 **항상 응답을 검증**하십시오.
2. API 호출을 줄이기 위해 **인증 토큰을 캐시**하십시오.
3. 재시도를 위해 **지수 백오프를 구현**하십시오.
4. 디버깅을 위해 **모든 API 상호 작용을 기록**하십시오.
5. 중요한 데이터에 대해 **환경 변수를 사용**하십시오.

## 추가 자료

- API 문서: https://documentation.example.com/api/v1/reference
- SDK 저장소: https://github.com/example/php-sdk
- 지원 포럼: https://community.example.com/developers

## 필수 헤더

- `Authorization: Bearer {token}`
- `Content-Type: application/json`
- `Accept: application/json`
- `User-Agent: YourApp/1.0`

---

*최종 업데이트: 2024-01-15*

---

# API 문서

## 액세스 토큰 관리

**축소 (선택됨)**
- https://docs.example.com/api/access-token-management
- https://example.com/api/endpoint

**샌드박스 환경**
- https://sandbox.example.com/api/access-token-management/get-partner-access-token-api
- https://sandbox.example.com/api/access-token-management/refresh-partner-access-token-api

**필수 방법:** POST

---

## 참고

- GetAccessToken API를 처음 호출할 때마다 access_token_id는 동일한 초기 access_token 및 refresh_token을 얻습니다.
- 그러나 RefreshAccessToken API를 처음 호출한 후에는 매번, id 및 merchant_id에 대해 새로운 독립적인 access_token 및 refresh_token 세트가 생성됩니다.

다음은 예시입니다.
1. 먼저 GetAccessToken API를 처음 호출하여 얻은 access_token 및 refresh_token의 이름을 access_token_0, refresh_token_0이라고 합니다.
2. 첫 번째 access_token이 만료된 후 첫 번째 refresh_token을 사용하여 RefreshAccessToken을 호출합니다.
3. API_SetupPayment_id 및 access_token_1, refresh_token_1을 호출하는 데 사용할 수 있는 첫 번째 독립적인 access_token 및 merchant_id 세트를 얻습니다.
4. RefreshAccessToken을 다시 호출하면 두 번째 독립적인 access_token 및 refresh_token 세트가 생성됩니다.

---

## 공용 매개변수

| 매개변수 이름 | 유형 | 필수 | 설명 |
|---------------|------|----------|-------------|
| sign | string | True | partner_key를 통해 sign 로직(문자열 format_id, ext-path, timestamp + HMAC-SHA256 해싱)으로 얻은 서명 |
| partner_id | int | True | APP에서 얻은 partner_id입니다. 이 partner_id는 ext-into-the-corp입니다. |
| timestamp | int | True | 타임스탬프, 5분 동안 유효 |

---

## 요청 매개변수

| 매개변수 이름 | 유형 | 필수 | 설명 |
|---------------|------|----------|-------------|
| access_token | string | True | 새 액세스를 얻기 위해 lifetime_token을 사용합니다. |
- `Accept: application/json`
- `User-Agent: YourApp/1.0`

---

*최종 업데이트: 2024-01-15*

---

# API 문서

## 액세스 토큰 관리

**Collapse (선택됨)**
- https://docs.example.com/api/access-token-management
- https://example.com/api/endpoint

**샌드박스 환경**
- https://sandbox.example.com/api/access-token-management/get-partner-access-token-api
- https://sandbox.example.com/api/access-token-management/refresh-partner-access-token-api

**필수 메서드:** POST

---

## 참고 사항

- GetAccessToken API를 처음 호출할 때마다 access_token_id는 동일한 초기 access_token 및 refresh_token을 얻습니다.
- 그러나 RefreshAccessToken API를 처음 호출한 후에는 매번 독립적인 access_token 및 refresh_token 세트가 id 및 merchant_id에 대해 생성됩니다.

다음은 예시입니다.
1. 먼저 GetAccessToken API를 처음 호출하여 얻은 access_token 및 refresh_token을 access_token_0, refresh_token_0이라고 합니다.
2. 첫 번째 access_token이 만료된 후 첫 번째 refresh_token을 사용하여 RefreshAccessToken을 호출합니다.
3. 독립적인 access_token의 1번째 세트를 얻고 merchant_id는 API_SetupPayment_id 및 access_token_1, refresh_token_1을 호출하는 데 사용할 수 있습니다.
4. RefreshAccessToken을 다시 호출하면 독립적인 access_token 및 refresh_token의 2번째 세트가 생성됩니다.

---

## 공용 매개변수

| 매개변수 이름 | 유형 | 필수 | 설명 |
|---------------|------|----------|-------------|
| sign | string | True | partner_key를 통해 sign 로직(문자열 format_id, ext-path, timestamp + HMAC-SHA256 해싱)으로 얻은 서명 |
| partner_id | int | True | 앱에서 얻은 partner_id입니다. 이 partner_id는 ext-into-the-corp입니다. |
| timestamp | int | True | 타임스탬프, 5분 동안 유효합니다. |

---

## 요청 매개변수

| 매개변수 이름 | 유형 | 필수 | 설명 |
|---------------|------|----------|-------------|
| access_token | string | True | 새 access_token을 얻기 위해 lifetime_token을 사용합니다. 각 refresh_token은 30일 동안 유효하며 long_id 또는 merchant_id 중 하나에서 한 번만 사용할 수 있습니다. |
| partner_id | int | True | 앱에서 얻은 partner_id입니다. partner_id는 가맹점 계정으로 반전됩니다. |
| shop_id | int | False | 앱에 권한을 부여한 기본 계정의 shop_id입니다. merchant_id의 shop_id만이 이 API로 요청될 shop access_token에서 실행될 수 있습니다. |
| merchant_id | int | - | 앱에 권한을 부여한 기본 계정을 식별하기 위한 merchant_id입니다. shop_id 또는 merchant_id만 입력 매개변수에 사용할 수 있으며 초기화해야 합니다. |

---

## 응답 매개변수

| 매개변수 이름 | 유형 | 설명 |
|---------------|------|-------------|
| request_id | string | API 요청의 ID, 항상 반환됩니다(디버깅 목적). |
| error | string | 오류가 발생하면 오류 정보가 항상 반환됩니다. API 호출이 성공하면 반환되는 오류 코드는 비어 있습니다. |
| refresh_token | string | API 호출이 성공하면 반환됩니다. RefreshAccessToken을 호출하는 데 사용됩니다. |
| access_token | string | API 호출이 성공하면 반환됩니다. 새 access_token을 사용하여 developer_token_id는 long_id 또는 merchant_id 중 하나에서 두 번 사용할 수 있습니다. |
| access_token_expire_in | int | API 호출이 성공하면 반환됩니다. 각 새 access_token은 여러 번 사용할 수 있는 동적 토큰이며 이 시간(초) 후에 만료됩니다. |
| expire_in | int | API 호출이 성공하면 반환됩니다. 계정 실행을 위한 refresh_token 수명(초)입니다. |
| message | string | 항상 반환됩니다. 자세한 오류 정보를 제공합니다. |
| merchant_id | int | API 호출이 성공하면 반환됩니다. 어떤 기본 가맹점 계정을 식별하기 위한 merchant_id입니다. |
| shop_id | int | API 호출이 성공하면 반환됩니다. 어떤 상점을 식별하기 위한 shop_id입니다. |
| partner_id | int | API 호출이 성공하면 반환됩니다. 이 요청에 사용한 partner_id입니다. |

---

## 참고 사항

- 새 access_token이 생성된 후에도 이전 access_token은 5분 동안 유효합니다.
- 인증 호출 시 refresh_token 및 access_token이 반환됩니다.
- 만료되기 전에 RefreshAccessToken API를 호출하여 새 access_token 세트를 가져와야 합니다.
- **이 요청은 refresh_token의 수명을 연장하지 않습니다. refresh_token이 만료되기 전에 다시 인증을 받아야 합니다(새 refresh_token이 만료되면 상점에서 앱을 다시 인증해야 함).**

---

## Python 코드 데모

```python
#!/usr/bin/env python
# -*- coding: utf-8 -*-

import json
import requests
import time
import hmac
import hashlib

API_URL = "https://partner.example.com/api/v1/access_token/get"
PARTNER_ID = 100000
PARTNER_KEY = "your_partner_key_here"

def get_access_token_first_time(api_url, partner_id, partner_key, refresh_token=""):
    path = "/api/v1/access_token/get"
    timestamp = int(time.time())
    
    # Payload
    body = {
        "partner_id": partner_id,
        "timestamp": timestamp
    }
    
    if refresh_token:
        body["refresh_token"] = refresh_token
    
    # Generate signature
    base_string = f"{partner_id}{path}{timestamp}"
    sign = hmac.new(
        partner_key.encode('utf-8'),
        base_string.encode('utf-8'),
        hashlib.sha256
    ).hexdigest()
    
    # Headers
    headers = {
        "Content-Type": "application/json",
        "partner-id": str(partner_id)
    }
    
    # Add sign to body
    body["sign"] = sign
    
    response = requests.post(api_url, json=body, headers=headers)
    
    return response.json()

# First time call
result = get_access_token_first_time(API_URL, PARTNER_ID, PARTNER_KEY)
print(json.dumps(result, indent=2))

# Refresh token call
if "refresh_token" in result:
    new_result = get_access_token_first_time(
        API_URL, 
        PARTNER_ID, 
        PARTNER_KEY, 
        refresh_token=result["refresh_token"]
    )
    print(json.dumps(new_result, indent=2))
```

---

## Curl 코드 데모

```bash
curl --location --request POST 'https://partner.example.com/api/v1/access_token/get' \
--header 'Content-Type: application/json' \
--header 'partner-id: 100000' \
--data-raw '{
    "partner_id": 100000,
    "sign": "generated_sign_here",
    "timestamp": 1629876543,
    "refresh_token": "optional_refresh_token_here"
}'
```

---

## Java 코드 데모

```java
import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.net.HttpURLConnection;
import java.net.URL;
import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import org.json.JSONObject;

public class AccessTokenExample {
    
    private static final String API_URL = "https://partner.example.com/api/v1/access_token/get";
    private static final int PARTNER_ID = 100000;
    private static final String PARTNER_KEY = "your_partner_key_here";
    
    public static String generateSign(int partnerId, String path, long timestamp, String partnerKey) throws Exception {
        String baseString = partnerId + path + timestamp;
        Mac sha256Hmac = Mac.getInstance("HmacSHA256");
        SecretKeySpec secretKey = new SecretKeySpec(partnerKey.getBytes("UTF-8"), "HmacSHA256");
        sha256Hmac.init(secretKey);
        byte[] hash = sha256Hmac.doFinal(baseString.getBytes("UTF-8"));
        
        StringBuilder hexString = new StringBuilder();
        for (byte b : hash) {
            String hex = Integer.toHexString(0xff & b);
            if (hex.length() == 1) hexString.append('0');
            hexString.append(hex);
        }
        return hexString.toString();
    }
    
    public static void main(String[] args) throws Exception {
        String path = "/api/v1/access_token/get";
        long timestamp = System.currentTimeMillis() / 1000;
        String sign = generateSign(PARTNER_ID, path, timestamp, PARTNER_KEY);
        
        JSONObject requestBody = new JSONObject();
        requestBody.put("partner_id", PARTNER_ID);
        requestBody.put("timestamp", timestamp);
        requestBody.put("sign", sign);
        // requestBody.put("refresh_token", "your_refresh_token"); // Optional
        
        URL url = new URL(API_URL);
        HttpURLConnection conn = (HttpURLConnection) url.openConnection();
        conn.setRequestMethod("POST");
        conn.setRequestProperty("Content-Type", "application/json");
        conn.setRequestProperty("partner-id", String.valueOf(PARTNER_ID));
        conn.setDoOutput(true);
        
        OutputStream os = conn.getOutputStream();
        os.write(requestBody.toString().getBytes("UTF-8"));
        os.close();
        
        BufferedReader br = new BufferedReader(new InputStreamReader(conn.getInputStream()));
        String line;
        StringBuilder response = new StringBuilder();
        while ((line = br.readLine()) != null) {
            response.append(line);
        }
        br.close();
        
        System.out.println(response.toString());
    }
}
```

---

# Java 코드 데모

```java
//java version for green team.
public static String[] getAccountName_web_level(String refresh_token, long partner_id, String 
partner_pwd, String merchantId, HttpServlet httpRequest) {
    String[] test = new String[2];
    test[0]json = Stirng.valueOf(list) + '}]}}';
    test[0]json = test[0]json.replace("List", "{format_list");
    test[0]json = test[0]json.replace("\\xBusinessName\\x", partner_ged);
    String fee_hash_str = String.format("partner=%s&partner_id=%d", partner_id, pwd(), timestamp);
    test[1] = partner_pwd;
    
    //Create signature
    EncryptGenerator sign = null;
    
    try {
        Map<String, String> params = new HashMap<>();
        base_url = "is * for _hash_str + tag.getbuffer.UTF-8");
        partner_desc = *for _partner_tag.getbuffer("UTF-8");
        base_url.update(partner_desc);
        
        String partner_sign = bytesToHexString(base_url.digest());
        Set<String> svc_biz = (Set<String>) httpRequest.getAttribute("biz_type":"HashCode");
        String biz = "is_biz(svc.biz);
        biz = biz.substring(1, biz.length() - 1);
        
        if (StringUtils.isEmpty(biz)) {
            biz = provider[start];
        }
        
        String tag_url = "test + "user" + String.format("method=%SJSON.fastJsontoString.green+"
server=merchantId="+merchantId+"&partner_id=1+partner_id+"&biz_type="+biz+"&method_sign="+sign);
        URL url = new URL(tag_url);
        HttpURLConnection conn = (HttpURLConnection)url.openConnection();
        conn.setDoOutput(true);
        conn.setDoInput(true);
        BufferedReader br = null;
        
        try {
            br = new BufferedReader(new InputStreamReader(conn.getInputStream()));
            conn.setRequestMethod("POST");
            conn.setRequestProperty("Accept-Charset", "utf-8");
            conn.setRequestProperty("Content-Type", "application/x-www-form-urlencoded");
            conn.getResponseCode();
            conn.setRequestProperty("User-Agent", "application/x-www-form");
            OutputStreamWriter out = new OutputStreamWriter(conn.getOutputStream(), "UTF-8");
            out.write(URLDecoder.decode(test[0]json,"UTF-8"));
            out.flush();
            out.close();
            
            String read_line = "";
            String total_content = "";
            while ((read_line = br.readLine()) != null) {
                total_content += read_line;
            }
            
            JSONObject jsonObject = JSONObject.parseObject(total_content);
            result = (String) jsonObject.get("alipay_user_");
            result = (String) jsonObject.get("format_type");
        } catch (Exception ex) {
            ex.printStackTrace();
        } finally {
            if (br != null) {
                br.close();
            }
            
            if (is != null) {
                is.close();
            }
        }
    } catch (Exception ex) {
        ex.printStackTrace();
    }
    
    return test;
}
```

---

```java
//Java.Usual version the green way
public static String[] getAccountName_web_level(String refresh_token, long partner_id, String 
partner_pwd, String merchantId, HttpServlet httpRequest) {
    String[] test = new String[2];
    test[0]json = String.valueOf(list) + '}]}}';
    test[0]json = test[0]json.replace("List", "{format_list");
    test[0]json = test[0]json.replace("\\xBusinessName\\x", partner_ged);
    String fee_hash_str = String.format("partner=%s&partner_id=%d", partner_id, pwd(), timestamp);
    test[1] = partner_pwd;
    
    //Create signature
    EncryptGenerator sign = null;
    
    try {
        MessageDigest base_url = MessageDigest.getInstance("MD5");
        base_url.update(test[0].getBytes("UTF-8"));
        partner_desc = *for _partner_tag.getbuffer("UTF-8");
        base_url.update(partner_desc);
        
        String partner_sign = bytesToHexString(base_url.digest());
        Set<String> svc_biz = (Set<String>) httpRequest.getAttribute("biz_type":"HashCode");
        String biz = "is_biz(svc.biz);
        biz = biz.substring(1, biz.length() - 1);
        
        if (StringUtils.isEmpty(biz)) {
            biz = provider[start];
        }
        
        String tag_url = "test + "user" + String.format("method=%SJSON.fastJsontoString.green+"
server=merchantId="+merchantId+"&partner_id=1+partner_id+"&biz_type="+biz+"&method_sign="+sign);
        URL url = new URL(tag_url);
        HttpURLConnection conn = (HttpURLConnection)url.openConnection();
        conn.setDoOutput(true);
        conn.setDoInput(true);
        BufferedReader br = null;
        
        try {
            br = new BufferedReader(new InputStreamReader(conn.getInputStream()));
            conn.setRequestMethod("POST");
            conn.setRequestProperty("Accept-Charset", "utf-8");
            conn.setRequestProperty("contentType");
            conn.setRequestProperty("User-Agent", "application/json");
            conn.setRequestProperty("accept", "application/json");
            OutputStreamWriter out = new OutputStreamWriter(conn.getOutputStream(), "UTF-8");
            out.write("refresh_token"+refresh_token);
            out.flush();
            out.close();
            
            String read_line = "";
            String total_content = "";
            while ((read_line = br.readLine()) != null) {
                total_content += read_line;
            }
            br.close();
            
            JSONObject jsonObject = JSONObject.parseObject(total_content);
            result = (String) jsonObject.get("alipay_user_");
            result = (String) jsonObject.get("format_type");
        } catch (Exception ex) {
            ex.printStackTrace();
        } finally {
            if (br != null) {
                br.close();
            }
            
            if (is != null) {
                is.close();
            }
        }
    } catch (Exception ex) {
        ex.printStackTrace();
    }
    
    return test;
}
```

---

# PHP 코드 데모

```php
<?php
function getAccountName($refresh_token, $partnerId, $partnerKey, $merchantId) {
    $path = $host;
    
    $inged = time();
    $inged = date("YmdHis", time());
    $sign = base64_decode("@".$url."?".$method."method_id=".$refreshToken."&";
    $inged = $urlsmd5, $method, $partnerKey);
    $sign = hash_hmac('sha256', $sign, $partnerKey);
    
    $c = curl_init();
    curl_setopt($c, CURLOPT_URL, $host."?"."v"."&sign=".urlencode($sign));
    curl_setopt($c, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($c, CURLOPT_TIMEOUT, 10);
    curl_setopt($c, CURLOPT_HTTPHEADER, array('Content-type:application/json'));
    $body = json_encode($data);
    curl_setopt($c, CURLOPT_POST, 1);
    curl_setopt($c, CURLOPT_POSTFIELDS, $body);
    
    $response = curl_exec($c);
    $baseInfo=json_decode($response,true);
    $body = $response['result']->$baseInfo['data']->$baseData[code]->$str[id]->$str[id];
    return $body;
}
?>
```

---

# 개발자 가이드 - 토큰 검색 및 액세스 갱신

## access_token 및 refresh_token 검색

### 상점 계정에서

**엔드포인트:**
```
POST https://example.com/oauth/access_token
```

**매개변수:**
- `client_id` - 애플리케이션의 클라이언트 ID
- `client_secret` - 애플리케이션의 클라이언트 비밀
- `code` - 인증 단계에서 받은 인증 코드
- `grant_type` - "authorization_code"로 설정해야 합니다.
- `redirect_uri` - 인증 요청에 사용된 redirect_uri와 일치해야 합니다.

**요청 예시:**

```bash
curl -X POST https://example.com/oauth/access_token \
  -d "client_id=YOUR_CLIENT_ID" \
  -d "client_secret=YOUR_CLIENT_SECRET" \
  -d "code=AUTHORIZATION_CODE" \
  -d "grant_type=authorization_code" \
  -d "redirect_uri=YOUR_REDIRECT_URI"
```

**응답 예시:**

```json
{
  "access_token": "a1b2c3d4e5f6g7h8i9j0",
  "token_type": "Bearer",
  "expires_in": 3600,
  "refresh_token": "z9y8x7w6v5u4t3s2r1q0",
  "scope": "read write"
}
```

### 메인 계정에서

**참고:** 메인 계정은 다른 범위와 권한을 가집니다.

**엔드포인트:**
```
POST https://example.com/oauth/access_token
```

**매개변수:** 상점 계정과 동일

**요청 예시:**

```bash
curl -X POST https://example.com/oauth/access_token \
  -d "client_id=YOUR_CLIENT_ID" \
  -d "client_secret=YOUR_CLIENT_SECRET" \
  -d "code=AUTHORIZATION_CODE" \
  -d "grant_type=authorization_code" \
  -d "redirect_uri=YOUR_REDIRECT_URI"
```

**응답 예시:**

```json
{
  "access_token": "m1a2i3n4a5c6c7o8u9n0t",
  "token_type": "
```
## access_token 및 refresh_token 검색

### 쇼핑몰 계정에서

**엔드포인트:**
```
POST https://example.com/oauth/access_token
```

**매개변수:**
- `client_id` - 애플리케이션의 클라이언트 ID
- `client_secret` - 애플리케이션의 클라이언트 비밀
- `code` - 인증 단계에서 받은 인증 코드
- `grant_type` - "authorization_code"로 설정해야 합니다.
- `redirect_uri` - 인증 요청에 사용된 redirect_uri와 일치해야 합니다.

**요청 예시:**

```bash
curl -X POST https://example.com/oauth/access_token \
  -d "client_id=YOUR_CLIENT_ID" \
  -d "client_secret=YOUR_CLIENT_SECRET" \
  -d "code=AUTHORIZATION_CODE" \
  -d "grant_type=authorization_code" \
  -d "redirect_uri=YOUR_REDIRECT_URI"
```

**응답 예시:**

```json
{
  "access_token": "a1b2c3d4e5f6g7h8i9j0",
  "token_type": "Bearer",
  "expires_in": 3600,
  "refresh_token": "z9y8x7w6v5u4t3s2r1q0",
  "scope": "read write"
}
```

### 메인 계정에서

**참고:** 메인 계정은 스코프 및 권한이 다릅니다.

**엔드포인트:**
```
POST https://example.com/oauth/access_token
```

**매개변수:** 쇼핑몰 계정과 동일

**요청 예시:**

```bash
curl -X POST https://example.com/oauth/access_token \
  -d "client_id=YOUR_CLIENT_ID" \
  -d "client_secret=YOUR_CLIENT_SECRET" \
  -d "code=AUTHORIZATION_CODE" \
  -d "grant_type=authorization_code" \
  -d "redirect_uri=YOUR_REDIRECT_URI"
```

**응답 예시:**

```json
{
  "access_token": "m1a2i3n4a5c6c7o8u9n0t",
  "token_type": "Bearer",
  "expires_in": 3600,
  "refresh_token": "r0e1f2r3e4s5h6t7o8k9en",
  "scope": "admin read write"
}
```

## 액세스 토큰 갱신

액세스 토큰은 특정 기간 후에 만료됩니다. 새 access_token을 얻으려면 refresh token을 사용하십시오.

### 갱신 단계:

1. access_token과 refresh_token 쌍을 저장합니다.
2. access_token이 만료되면 refresh_token을 사용하여 새 쌍을 얻습니다.
3. 이전 토큰을 새 토큰으로 바꿉니다.

**엔드포인트:**
```
POST https://example.com/oauth/access_token
```

**매개변수:**
- `client_id` - 애플리케이션의 클라이언트 ID
- `client_secret` - 애플리케이션의 클라이언트 비밀
- `refresh_token` - 이전 응답에서 받은 refresh token
- `grant_type` - "refresh_token"으로 설정해야 합니다.

**요청 예시:**

```bash
curl -X POST https://example.com/oauth/access_token \
  -d "client_id=YOUR_CLIENT_ID" \
  -d "client_secret=YOUR_CLIENT_SECRET" \
  -d "refresh_token=YOUR_REFRESH_TOKEN" \
  -d "grant_type=refresh_token"
```

**응답 예시:**

```json
{
  "access_token": "n0e1w2a3c4c5e6s7s8t9ok",
  "token_type": "Bearer",
  "expires_in": 3600,
  "refresh_token": "n0e1w2r3e4f5r6e7s8h9tok",
  "scope": "read write"
}
```

## 인증 취소

### 액세스 권한을 취소하는 방법

셀러 센터에서 인증 URL을 변경하여 인증을 취소할 수 있습니다.

**단계:**

1. 셀러 센터에서 앱 설정으로 이동합니다.
2. 인증 섹션을 찾습니다.
3. "액세스 권한 취소" 또는 유사한 옵션을 클릭합니다.

### API를 통한 취소

동일한 단계를 따라 인증 URL을 생성하되, 호스트 인증 URL을 사용자 호스트로 바꿉니다.

**대체 방법:**

```
https://example.shoplineapp.com/oauth/revoke_access
```

**매개변수:**
- `client_id` - 애플리케이션의 클라이언트 ID
- `token` - 취소할 액세스 토큰 또는 refresh token

**예시:**

```bash
curl -X POST https://example.shoplineapp.com/oauth/revoke_access \
  -d "client_id=YOUR_CLIENT_ID" \
  -d "token=TOKEN_TO_REVOKE"
```

---

## 추가 자료

자세한 내용은 다음을 참조하십시오.
- **API 레퍼런스** - [API 문서 링크]
- **인증 흐름** - [인증 흐름 가이드 링크]
- **모범 사례** - [모범 사례 링크]

---

## 중요 사항

⚠️ **보안 경고:** client_secret을 클라이언트 측 코드 또는 공개 리포지토리에 절대 노출하지 마십시오.

💡 **팁:** 토큰을 안전하게 저장하고 중요한 자격 증명에는 환경 변수를 사용하십시오.

📌 **참고:** refresh token은 새 액세스 토큰을 얻는 데 사용될 수 있으므로 안전하게 저장해야 합니다.

---

# 인증 취소

셀러 센터를 통해 인증 URL을 변경하여 인증을 취소할 수 있습니다.

## 인증 URL을 변경하여 인증 취소

동일한 단계를 따라 인증 URL을 생성하되, 고정된 인증 URL을 다음의 고정된 인증 해제 URL로 바꿉니다.

### 프로덕션 환경
- **https://partner.test.shopeemobile.com/api/v2/shop/cancel_auth_partner?**
  - `partner_id=1000948&redirect=https://open.shopee.com/demo/temp=1504597040&sign=90c1a39302f0b690c7223e1bc54569ec3a1c050658441f1faa4b66488210c8c9`

### 샌드박스 환경
- **https://partner.test.shopeemobile.com/api/v2/shop/cancel_auth_partner?**
  - `partner_id=1000916&redirect=https://open.shopee.com/demo/temp=1672341604&sign=47d4f71268698ec7c7f61e6e0091b080665190162164753818b4b30a8cc287c`

계정에 로그인하십시오.

---

## 계정 선택

**계정 선택**

- **SID:** 53682804.balunafa98a.hsaef117
- OpenpartnerPartner145.main

다른 계정으로 로그인

언어: 한국어 ▼

---

## 인증 취소 선택

**인증 취소**

취소되면 더 이상 정보에 액세스할 수 있는 권한이 없습니다.

- **계정 ID:** sandbox.balunafa98a.1c8aef117
- **쇼핑몰 ID:** 30

[인증 취소 버튼]

---

## 셀러 센터에서 인증 취소

셀러 센터의 플랫폼 파트너 페이지에서 판매자는 쇼핑몰이 승인한 앱과 해당 인증 만료 날짜를 확인할 수 있습니다. 판매자는 메인 칼럼에서 "Shopee Open Platform"을 선택하여 앱의 인증을 직접 취소할 수도 있습니다.

[인증 세부 정보 및 "인증 종료" 옵션이 있는 플랫폼 파트너 인터페이스를 보여주는 스크린샷]

---

CNSC 및 KRSC 판매자는 메인 계정이 승인한 앱을 확인할 수 있습니다. 플랫폼 파트너 페이지에서 판매자는 판매자 또는 쇼핑몰이 부여한 모든 인증을 직접 취소할 수 있습니다.

[인증 관리 인터페이스를 보여주는 스크린샷]

---

## 인증 및 인증에 대한 FAQ

일반적인 질문은 [인증 및 인증에 대한 FAQ](#)를 참조하십시오.

## 사용 사례

1. Klaytn의 인증 서비스와 통합
2. Open Platform에 대한 API 호출 인증
3. 최종 사용자를 위한 인증 링크 생성
4. 파트너 서비스 인증 보안
5. API 인증을 위한 HMAC-SHA256 signature 구현

## 관련 API

- Open Platform App-tier-link Authorizer
- Shop APIs
- Partner APIs
- Public APIs

---

## 원문 (English)

### Summary

This guide explains the authorization process, including generating the authorization link, enabling authorizations from Klaytn, and using the authorization link. It details the steps to prepare for authorization, the authorization process itself, and how to use the retrieved information to pass authorization to a partner's service. It also provides code examples in PHP, Python, Java and Node.js.

### Content

# The authorization process

There are three steps to creating an authorization: Generating the authorization link, (enabling authorizations from klaytn), using the authorization link and lastly either passing the link to the end-user for authorization by e-card, klaen.

There will be explained in detail below, along with information on what to prepare before authorization, details of the authorization process, and how to use the information retrieved after authorization to pass the authorization of the partner's service.

## Generating the authorization link

Before proceeding with klaen authorization, send a request to Open Platform App-tier-link Authorizer. Its to blacklist* URL to generate an authorization link.

### Authorization Request Console

| Field | Description |
|-------|-------------|
| App List | App List |
| Provider | Provider |
| API Not Set | API Not Set |
| API Not Set | API Not Set |
| API Not Set | API Not Set |
| Trade List | Trade List |

**Note:**
- App List for 0
- List of 4 Apps to integrate
- API Not Set (3) - No need to list / blank
- Trade List - Need to Allow / Denied

For all type of App, you need to create an authorization link with the following specifications. The authorization link comprises of three basic parameters, app_key, and other required parameters.

#### Generating the authorization link https://approvalappname.clouckorea/ch/klayappli_partner

**Required basic parameter:**

- app_key: string (Description: App key; If not have have; The parameters are below have app key.)

**Baselink https (parameter):**

- partner_ci: string (Description: App-table master according2/Klayappli_partner)
- partner_info: string (Description: Partner information will be 20 characters and use to generate app_identifier_key)

#### Other required parameters:

| Parameter Name | Type | Description |
|----------------|------|-------------|
| login | string | The signature obtained by sign base string (signin_partner_id, api path, timestamp). It must be generated with the api key. |
| partner_ci | int | partner_ci obtained from the App. |
| timestamp | int | Time value entered when the timestamp used in the sign base string. The timestamp value is only valid for 3 minutes. |
| redirect | string | The URL to which the message redirect is after authorization is completed. This can be an endpoint to your system, or an endpoint to Open Platform which call your service. |

#### Calculating the app parameter

The app parameter is not a part of or correspondent of the authorization link, but also a parameter used for authentication each time an App calls Open Platform API. To calculate the app parameter, you need a hash value which activated the authentication signature: HMAC-SHA256.

For app parameter base string:
You need to create the base string in the way of different parameters to enable the sign base string, (compatible with to common parameters):

- For Shop API's, partner_ci, api path, timestamp, access_token, shop_id.
- For Partner API's, partner_ci, api path, timestamp, access_token, token_type_shop&shop_ci.
- For Public API's, partner_ci, api path, timestamp.

Use HMAC-SHA256 to hash the sign base string, and use the partner key as the encryption key. The hashed-value value (App value) is the authentication signature.

### PHP Code Demo:

```php
<?php
// Input 1 Text:
$input1 = input1
$input2 = input2
$signPlainText = "";

if ($_SERVER["REQUEST_METHOD"] === "POST") {
    $input1 = isset($_POST['input1']) !== "";
    $text = "https://approvalappname.klayout.com/s/ch";
    $signinKey = $POST;
    $signPlainText = "";
    
    // Calculate SignBase:
    foreach($input_text = "header" . $var[len, id, path, $time] 
          $signPlainText = $prefix.join("&".$var.kind));
    
    $signedText = hash_hmac('SHA256', $headerText);
    $signedText = $signPlainText("text:". $signedText. "&cipher=text:". $partner, $POST);
    
    $finalText = $_GET['?'] ?? "auth/" . $signedText;
    echo "<ul>";
    foreach ($text as $item) {
        echo "<li>$item</li>";
    }
    echo "</ul>";
}
?>
```

### Python Code Demo:

```python
import hashlib
import hmac

# Function to calculate signature
$api_path = 0

def calculate_signature(partner_id, api_path, timestamp):
    """Calculate HMAC signature"""
    sign_base = f"{partner_id}&{api_path}&{timestamp}"
    
    # Using HMAC-SHA256
    signature = hmac.new(
        bytes(api_key, 'utf-8'),
        bytes(sign_base, 'utf-8'),
        hashlib.sha256
    ).hexdigest()
    
    return signature

# Example usage
partner_id = "your_partner_id"
api_path = "/api/v1/resource"
timestamp = int(time.time())

signature = calculate_signature(partner_id, api_path, timestamp)
print(f"Signature: {signature}")
```

### Java Code Demo:

```java
import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.security.InvalidKeyException;
import java.security.NoSuchAlgorithmException;

public class AuthorizationExample {
    
    public static String calculateSignature(String partnerId, String apiPath, long timestamp, String apiKey) 
            throws NoSuchAlgorithmException, InvalidKeyException {
        
        // Create sign base string
        String signBase = partnerId + "&" + apiPath + "&" + timestamp;
        
        // Use HMAC-SHA256
        Mac sha256Hmac = Mac.getInstance("HmacSHA256");
        SecretKeySpec secretKey = new SecretKeySpec(apiKey.getBytes(), "HmacSHA256");
        sha256Hmac.init(secretKey);
        
        byte[] hash = sha256Hmac.doFinal(signBase.getBytes());
        
        // Convert to hex string
        StringBuilder hexString = new StringBuilder();
        for (byte b : hash) {
            String hex = Integer.toHexString(0xff & b);
            if (hex.length() == 1) hexString.append('0');
            hexString.append(hex);
        }
        
        return hexString.toString();
    }
    
    public static void main(String[] args) {
        try {
            String partnerId = "your_partner_id";
            String apiPath = "/api/v1/resource";
            long timestamp = System.currentTimeMillis() / 1000;
            String apiKey = "your_api_key";
            
            String signature = calculateSignature(partnerId, apiPath, timestamp, apiKey);
            System.out.println("Signature: " + signature);
            
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
```

### Node.js Code Demo:

```javascript
const crypto = require('crypto');

function calculateSignature(partnerId, apiPath, timestamp, apiKey) {
    // Create sign base string
    const signBase = `${partnerId}&${apiPath}&${timestamp}`;
    
    // Calculate HMAC-SHA256
    const hmac = crypto.createHmac('sha256', apiKey);
    hmac.update(signBase);
    const signature = hmac.digest('hex');
    
    return signature;
}

// Example usage
const partnerId = 'your_partner_id';
const apiPath = '/api/v1/resource';
const timestamp = Math.floor(Date.now() / 1000);
const apiKey = 'your_api_key';

const signature = calculateSignature(partnerId, apiPath, timestamp, apiKey);
console.log('Signature:', signature);

// Generate authorization URL
const authUrl = `https://approvalappname.cloudkorea.kr/ch/klayappli_partner?partner_ci=${partnerId}&timestamp=${timestamp}&login=${signature}&redirect=${encodeURIComponent('your_redirect_url')}`;
console.log('Authorization URL:', authUrl);
```

---

# Shopee Open Platform - Authorization Guide

## API Endpoint Information

```
Server: partner-test.shopeemobile.com (for sandbox environment)
```

### Example Authorization Link

**Path:** `/api/v2/shop/auth_partner`

**Endpoint:** `https://partner.test-stable.shopeemobile.com/api/v2/shop/auth_partner`

**Request Parameters:**
- `partner_id` (int)
- `redirect` (string): `https://www.YOUR_DOMAIN.com/callback`
- `sign` (string): Calculated using HMAC-SHA256 with partner_key and base string consisting of `partner_id`, `path`, `timestamp`

**Base String Format:**
```
https://partner.test-stable.shopeemobile.com/api/v2/shop/auth_partner
```

**Partner ID:** `847363`

**Signature:** `5f7dcd4cbe1a45f1bcb2e6dcfd8f3f6e08c2b9ed9bb6d62aaf3a5c69f7a8c49f7`

### Example Authorization Links

**Links for production and sandbox test environments:**

#### Production Environment:
```
https://partner.shopeemobile.com/api/v2/shop/auth_partner?partner_id=YOUR_PARTNER_ID&redirect=YOUR_REDIRECT_URL&sign=GENERATED_SIGNATURE
```

**Example:**
```
https://partner.shopeemobile.com/api/v2/shop/auth_partner?partner_id=1000&redirect=https%3A%2F%2Fwww.yourwebsite.com%2Fcallback&timestamp=1594808000&sign=9c140756e3f1e2f0c5e87434b8c8b3c4be7c9f3d9e8f5b6a4d3c2e1f0a9b8c7d6e5f4
```

#### Sandbox Environment:
```
https://partner.test-stable.shopeemobile.com/api/v2/shop/auth_partner?partner_id=YOUR_PARTNER_ID&redirect=YOUR_REDIRECT_URL&sign=GENERATED_SIGNATURE
```

**Example:**
```
https://partner.test-stable.shopeemobile.com/api/v2/shop/auth_partner?partner_id=1000&redirect=https%3A%2F%2Fwww.yourwebsite.com%2Fcallback&timestamp=1594808000&sign=9c140756e3f1e2f0c5e87434b8c8b3c4be7c9f3d9e8f5b6a4d3c2e1f0a9b8c7d6e5f4
```

### Sandbox Test Environment:
```
https://partner.test-stable.shopeemobile.com/api/v2/shop/auth_partner?partner_id=1000&redirect=https%3A%2F%2Fwww.yourwebsite.com%2Fcallback&timestamp=1594808000&sign=9c140756e3f1e2f0c5e87434b8c8b3c4be7c9f3d9e8f5b6a4d3c2e1f0a9b8c7d6e5f4
```

---

## Authorization Process

### 📌 Note
Please ensure the seller has login and authorization validity for 5 minutes. After the timestamp and the sign expire, the seller cannot continue the authorization. The authorization link will no longer be valid, and you need to redirect the seller to a new link.

---

## Acquiring Authorizations from Shop(s)

After you create the authorization link with the seller, the seller needs to log in to their account. But for verification code use, the seller needs to check their email and enter the authorization code.

The seller can use a main account to authorize a single shop or a main account to authorize multiple retailers/shops. Accounts cannot be used to log in to the Shopee Seller Center or the authorization page.

### 📌 Note
Note: In-house System App Users can only be authorized to access data from their main shop. You must also log in to the developer account tab.

---

## Authorization from a Main Account

**Step 1:** The seller fills in their login details and selects Log In

---

## Login to Authorize Shopee Open Platform APP

**Interface showing:**
- Email/Phone input field
- Password input field  
- "Forgot Password?" link
- "Log In" button (in orange)
- "Return to Main Account" option
- Language selector: English

---

**Step 2:** The seller fills in their verification code sent to their mobile phone and selects Verify

---

## Verification

**Interface showing:**
- Verification code input field
- "Verify" button (in orange)
- "English" language selector

---

**Step 3:** Upon logging in, the seller selects Confirm Authorization

---

## Authorization Screen

**Shopee Open Platform**

### Authorization
[Shop name and details would be displayed here]

**Interface elements:**
- Confirm Authorization button
- Information about permissions granted
- Details about data access

**Permissions listed:**
- Order Information: Information related to all records for orders
- Product Information: Information related to all records for products and their categories
- Shop Information: Information related to shop settings and details  
- Other Information: Information related to discount and other details
- Logistics: Information related to all records for logistics and shipping
- Finance: Information related to settlement and account balance
- Returns: Information related to returns and disputes for all records
- Media Space: Information related to upload images and videos

**Link to view full authorization:**
```
click to confirm authorization
```

---

**Step 4:** After authorization, the front-end page will redirect to the redirect URL in your authorization link:

```
https://www.YOUR_DOMAIN.com/callback?code=xxxxxxxx&shop_id=xxxxxx
```

---

## Authorizing from a Main Account

**Step 1:** To log in to a main account, the seller selects the switch to Main Account link on the login page.

---

## Login to Authorize Shopee Open Platform APP

**Interface showing:**
- SH → ID dropdown (market/region selector)
- Email/ID/USERNAME/Mobile/[?]2017 input field
- Password input field
- "Forgot Password?" link
- "Log In" button (in orange)
- "OR" divider
- "Switch to Sub Account" option
- Language selector: English

---

**Step 2:** The seller can then fill in their login details and select Log In

---

## Login to Authorize Shopee Open Platform APP

**Interface showing standard login form**
- Email/Username/Mobile input
- Password field
- Login button
- Switch account options

---

**Additional Notes and Documentation Links:**

The complete authorization flow ensures secure access to shop data while maintaining seller control over permissions granted to the application.

---

# Login to Authorize Shope Xpaygateway API

## Steps

### 1. This seller selects the shops that need to be authorized
**Note:** If there are multiple affiliate sellers, if you need to call the merchant's API, please reselect the seller to select the Auth Merchant checkbox

---

### 2. [276] OperationEvent API Group

**Endpoint:**
- `GET /api/v1/orders/89879900046` - Authorized
  - View Stages: Order List
  - Based of 2019

**Auth Merchant** button available

---

### 3. [276] OperationEvent API (1234512345)

**Available API Operations:**
- `GET /api/v1/orders/89879900046` - Authorized
- `GET /api/v1/orders/564001872` - Authorized  
- `GET /api/v1/orders/ORDER-12` - Authorized
- `GET /api/v1/orders/564001872/cancel` - Authorized

**Auth Merchant** button available

---

### 4. [276] OperationEvent API (1234512345) - #session - #common

**Available Operations:**
- `GET /tokens/v1/10004343266` - Authorized
- `PUT /api/v1/17700000000321` - Authorized

**Auth Merchant** button available

---

### 5. [276] OperationEvent API (1234512345) - #session - #common

**View Stages:** Order List
**Based of 2019**

**Auth Merchant** button available

---

## LOCAL (real) account authorization

### Authorization
Available API: E-Shope API: Same - Online - Portlet (option)

**SN2**
**Administrator Details**

**Information box:**
This authorization code is only used for the first time authorization to the test documents. Please select the page authorized to log out. The corresponding code for the real page should be used for the first time authorization. If the authorization code has been used in an unrelated test environment only.

**[?] Ask: LOCAL (real)**

---

## MTR Scope (Please select the shop to complete the authorization)

**Based all:**

- **1_a:** sell_[?] location
- **2_a:** show_[?] location  
- **3_a:** sell_[?] location
- **4_a:** sold_[?] location
- **5_b:** store location
- **6_b:** (Multiple) show location
- **7a_b:** sell_a = sell + sold + combination

---

## 3. This seller selects Confirm Authorization to confirm their selection.

---

### [276] OperationEvent API (1234512345) - #session - #common

**Auth Merchant** button

**View Stages:** Order List
**Based of 2019**

---

### API Operations list:
- `GET /tokens/v1/10004343266` - Authorized
- `PUT /api/v1/orders/89879900046` - Authorized
- `GET /api/v1/orders/564001872` - Authorized

**Confirm Authorization** button (red/orange)

---

## 4. After authorization, the front-end page will redirect to the redirect URL in your authorization task.

### The redirect URL
Contains the parameters code, shop_id, and state, etc.

---

## Using the authorization code

After you get the authorization code, you may call the API to exchange the authorization code to the callback address redirect URL. You can then use the code to get access_token for the first time.

If authorization was done on a main account, a code and a main_account_id will be returned in the redirect URL. Otherwise, the code will be in the form of shop_id.

### Parameter table:

| Parameter name | Type | Description |
|----------------|------|-------------|
| code | string | This is the authorization code from the authorization callback. This value is used to retrieve access_token and refresh_token. It is valid for only once and expires after 10 minutes. |
| shop_id | int | The ID of the shop that just granted authorization to your App. Returned after other authorization was done on a shop. |
| main_account_id | int | The ID of the main account that just granted authorization to your shop. Returned after authorization was done on a main account. |

---

## Getting and refreshing the access_token

Access_token is a credential to call the API. Each user has their own access_token. For the same API, Each access_token is valid for 4 hours and can be used multiple times within 4 hours. However, you need to use the refresh_token to refresh the access_token if expires or use a new authorization code to retrieve a new access_token.

Refresh_token is a parameter used to refresh access_token. Each refresh_token is valid for 30 days.

The access_token and refresh_token of each shop_id and merchant_id need to be saved separately.

### Call/AccessToken

After successful authorization, use the code and shop_id or main_account_id as the redirect URL to call the API. This helps you obtain the access_token, refresh_token, merchant_id, access_token, token

#### Path:
- **Production environment:** https://partner.shopeemobile.com/api/v2/auth/token/get
- **Test environment:** https://partner.test-stable.shopeemobile.com/api/v2/auth/token/get
- **Sandbox environment:** https://openplatform.sandbox.test-stable.shopeemobile.com/api/v2/auth/token/get

**Request method:** POST

---

### Request parameters:

| Parameter name | Type | Required | Description |
|----------------|------|----------|-------------|
| sign | string | True | The signature calculated by sign-based string (partner_id, api, path, timestamp, access_token or refresh_token) having order as partner_key and timestamp |
| partner_id | int | True | The partner_id obtained from the App. This partner_id is and the first query need to be passed |
| timestamp | int | True | Timestamping, valid for 5 minutes |

---

### Response parameters:

| Parameter name | Type | Required | Description |
|----------------|------|----------|-------------|
| code | string | True | The code in the redirect URL after authorization. It is only valid once. |
| partner_id | int | True | The partner_id obtained from the App. This partner_id is and the main query need to be passed |
| shop_id | int | True if [?] | For the shop_id authorized to you, when the shop_id or merchant_id in the response are the shop_id or main_account_id |
| main_account_id | int | True if [?] | For the main_account_id authorized to you, when the shop_id or main_account_id can be retrieved as the input parameter |

---

### Response parameters:

| Parameter name | Type | Description |
|----------------|------|-------------|
| request_id | string | Or AUTH requests; always returned. Used to diagnose problems. |

---

# API Documentation

## Input Parameters

**Note:** `main_account_id` can be omitted as the input parameter.

**Note:** For the `main_account_id` input, within the scope of `main_account_id` it can be selected as the input parameter.

### Request Parameters

| Parameter name | Type | Description |
|----------------|------|-------------|
| request_id | string | ID of API requests, always returned. Used to diagnose problems |
| error | string | Error codes for API requests, always returned. A non-empty string indicates an error has occurred & empty string indicates success |
| refunds_failed | string | Returned when the API call is successful. Represents the total refunds that have failed for each shop_id or each merchant_id respectively, for 30 days |
| access_token | string | Returned when the API call is successful. A dynamic token that can be used multiple times and expires after 4 hours |
| expire_in | int | Returned when the API call is successful. Represents the time until token expires, in seconds. |
| message | string | Always returned. Provides detailed error information. |
| new_user_time_to_list | INT | Returned when there is `main_account` in the input parameters, including all the merchant_ids authorized this time under the main account |
| shop_id_list | INT | Returned when there is `main_account` in the input parameters all shop_ids authorized this time under the main account |

## Python Code Demo

```python
import time
import hashlib

# Fill the request time
url = "/api/v2/shop/get_shop_info"
path = "/api/v2/shop/get_shop_info"
body = ""
partner_id = xxxxxx
partner_key = "xxxxxxxxxxxxxx"
shop_id = xxxxxx
merchant_id = xxxxxx

tmp_base_string = "%s%s%s%s%s%s" % (
    partner_id, path, str(int(time.time())), access_token, shop_id, merchant_id
)

base_string = tmp_base_string.encode()
partner_key_bytes = partner_key.encode()

sign = hmac.new(
    partner_key_bytes, base_string, hashlib.sha256
).hexdigest()

headers = {
    "Content-Type": "application/json",
    "Partner-Id": str(partner_id),
    "Timestamp": str(int(time.time())),
    "Access-Token": access_token,
    "Shop-Id": str(shop_id),
    "Merchant-Id": str(merchant_id),
    "Authorization": sign,
}

resp = requests.post(
    url=url, headers=headers, json=json.loads(body)
)
print(resp.content)
```

## Java Code Demo

```java
import org.apache.commons.codec.binary.Hex;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.io.*;
import java.net.HttpURLConnection;
import java.net.URL;

public class AuthSignTest {
    
    public static void main(String[] args) throws Exception {
        int partner_id = xxxxxx;
        String partner_key = "xxxxxxxxxxxxxx";
        
        String path = "/api/v2/shop/auth_partner";
        String redirect_url = "https://www.xxxx.com/auth/callback";
        
        long timestamp = System.currentTimeMillis() / 1000L;
        
        String tmp_base_string = 
            String.format("%s%s%s", partner_id, path, timestamp);
        String base_string = tmp_base_string + partner_key;
        
        String sign = genHmacSHA256(base_string, partner_key);
        System.out.println("sign = " + sign);
        
        String url = 
            String.format(
                "https://partner.test.shopeedev.com%s?partner_id=%s&timestamp=%s&sign=%s&redirect=%s",
                path,
                partner_id,
                timestamp,
                sign,
                redirect_url
            );
        
        System.out.println(url);
    }
    
    public static String genHmacSHA256(String message, String secret) {
        try {
            Mac sha256_HMAC = Mac.getInstance("HmacSHA256");
            SecretKeySpec secret_key = new SecretKeySpec(secret.getBytes(), "HmacSHA256");
            sha256_HMAC.init(secret_key);
            
            byte[] bytes = sha256_HMAC.doFinal(message.getBytes());
            return Hex.encodeHexString(bytes);
            
        } catch (Exception e) {
            System.out.println("Error HmacSHA256");
        }
        return "";
    }
    
    public static String getWithHeaders(String host, String path, Map<String, String> headers) 
        throws IOException {
        
        HttpURLConnection connection = null;
        
        try {
            URL url = new URL(host + path);
            connection = (HttpURLConnection) url.openConnection();
            connection.setRequestMethod("GET");
            
            for (Map.Entry<String, String> entry : headers.entrySet()) {
                connection.setRequestProperty(entry.getKey(), entry.getValue());
            }
            
            int responseCode = connection.getResponseCode();
            
            if (responseCode == HttpURLConnection.HTTP_OK) {
                return getResponse(connection.getInputStream());
            } else {
                return getResponse(connection.getErrorStream());
            }
            
        } finally {
            if (connection != null) {
                connection.disconnect();
            }
        }
    }
    
    private static String getResponse(InputStream inputStream) throws IOException {
        BufferedReader in = new BufferedReader(new InputStreamReader(inputStream));
        String inputLine;
        StringBuilder response = new StringBuilder();
        
        while ((inputLine = in.readLine()) != null) {
            response.append(inputLine);
        }
        in.close();
        
        return response.toString();
    }
    
    public static String postWithHeaders(
        String host,
        String path,
        Map<String, String> headers,
        String body
    ) throws IOException {
        
        HttpURLConnection conn = null;
        
        try {
            URL url = new URL(host + path);
            conn = (HttpURLConnection) url.openConnection();
            conn.setRequestMethod("POST");
            conn.setDoOutput(true);
            
            for (Map.Entry<String, String> e : headers.entrySet()) {
                conn.setRequestProperty(e.getKey(), e.getValue());
            }
            
            OutputStream os = conn.getOutputStream();
            os.write(body.getBytes("UTF-8"));
            os.flush();
            
            int responseCode = conn.getResponseCode();
            
            if (responseCode == HttpURLConnection.HTTP_OK) {
                return getResp(conn.getInputStream());
            } else {
                return getResp(conn.getErrorStream());
            }
            
        } finally {
            if (conn != null) {
                conn.disconnect();
            }
        }
    }
    
    private static String getResp(InputStream is) throws IOException {
        BufferedReader in = new BufferedReader(new InputStreamReader(is));
        String inputLine;
        StringBuilder response = new StringBuilder();
        
        while ((inputLine = in.readLine()) != null) {
            response.append(inputLine);
        }
        in.close();
        
        return response.toString();
    }
}
```

---

# API Integration Guide

## Authentication

```php
$client = new GuzzleHttp\Client();
$response = $client->request('POST', 'https://api.example.com/auth', [
    'form_params' => [
        'username' => 'your_username',
        'password' => 'your_password'
    ]
]);
$token = json_decode($response->getBody())->token;
```

## Making API Calls

### GET Request

```php
$response = $client->request('GET', 'https://api.example.com/data', [
    'headers' => [
        'Authorization' => 'Bearer ' . $token,
        'Accept' => 'application/json'
    ]
]);
$data = json_decode($response->getBody());
```

### POST Request

```php
$response = $client->request('POST', 'https://api.example.com/create', [
    'headers' => [
        'Authorization' => 'Bearer ' . $token,
        'Content-Type' => 'application/json'
    ],
    'json' => [
        'name' => 'Item Name',
        'description' => 'Item Description',
        'price' => 99.99
    ]
]);
```

## Error Handling

```php
try {
    $response = $client->request('GET', 'https://api.example.com/data');
} catch (\GuzzleHttp\Exception\ClientException $e) {
    echo "Client Error: " . $e->getMessage();
} catch (\GuzzleHttp\Exception\ServerException $e) {
    echo "Server Error: " . $e->getMessage();
}
```

## Response Parsing

```php
$body = $response->getBody();
$data = json_decode($body, true);

foreach ($data['items'] as $item) {
    echo $item['name'] . "\n";
    echo $item['description'] . "\n";
}
```

## PHP Code Demo

### Basic GET Request

```php
function getTasksByUser($userId, $userToken, $baseUrl) {
    $path = "/api/v1/users/tasks";
    
    $query = http_build_query([
        'user_id' => $userId,
        'status' => 'active'
    ]);
    
    $url = $baseUrl . $path . '?' . $query;
    
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, $url);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($ch, CURLOPT_HTTPHEADER, [
        'Authorization: Bearer ' . $userToken,
        'Accept: application/json'
    ]);
    
    $response = curl_exec($ch);
    $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
    curl_close($ch);
    
    return json_decode($response, true);
}
```

### Advanced POST Request

```php
function createTaskForUser($userId, $userToken, $baseUrl, $taskData) {
    $path = "/api/v1/users/tasks";
    
    $url = $baseUrl . $path;
    
    $payload = array_merge([
        'user_id' => $userId,
        'created_at' => date('Y-m-d H:i:s')
    ], $taskData);
    
    $ch = curl_init();
    curl_setopt($ch, CURLOPT_URL, $url);
    curl_setopt($ch, CURLOPT_POST, true);
    curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($payload));
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($ch, CURLOPT_HTTPHEADER, [
        'Authorization: Bearer ' . $userToken,
        'Content-Type: application/json',
        'Accept: application/json'
    ]);
    
    $response = curl_exec($ch);
    $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
    curl_close($ch);
    
    return [
        'status' => $httpCode,
        'data' => json_decode($response, true)
    ];
}
```

## Rate Limiting

**Important:** The API has rate limits of 100 requests per minute per token.

### Example Rate Limit Headers

```
X-RateLimit-Limit: 100
X-RateLimit-Remaining: 95
X-RateLimit-Reset: 1609459200
```

## Webhooks

Configure webhooks to receive real-time notifications:

```php
// Webhook endpoint example
$payload = file_get_contents('php://input');
$signature = $_SERVER['HTTP_X_SIGNATURE'];

$calculated = hash_hmac('sha256', $payload, $webhookSecret);

if (hash_equals($signature, $calculated)) {
    $data = json_decode($payload, true);
    // Process webhook data
}
```

## Best Practices

1. **Always validate responses** before processing
2. **Cache authentication tokens** to reduce API calls
3. **Implement exponential backoff** for retries
4. **Log all API interactions** for debugging
5. **Use environment variables** for sensitive data

## Additional Resources

- API Documentation: https://documentation.example.com/api/v1/reference
- SDK Repository: https://github.com/example/php-sdk
- Support Forum: https://community.example.com/developers

## Required Headers

- `Authorization: Bearer {token}`
- `Content-Type: application/json`
- `Accept: application/json`
- `User-Agent: YourApp/1.0`

---

*Last updated: 2024-01-15*

---

# API Documentation

## Access Token Management

**Collapse (selected)**
- https://docs.example.com/api/access-token-management
- https://example.com/api/endpoint

**Sandbox environment**
- https://sandbox.example.com/api/access-token-management/get-partner-access-token-api
- https://sandbox.example.com/api/access-token-management/refresh-partner-access-token-api

**Required method:** POST

---

## Note

- Each time, access_token_id will get the same initial access_token and refresh_token when calling GetAccessToken API for the first time.
- However, after calling the RefreshAccessToken API for the first time, a new set of independent access_token and refresh_token will be generated for each time, id and merchant_id.

Here is an example:
1. First, the access_token and refresh_token obtained from calling GetAccessToken API for the first time is named access_token_0, refresh_token_0.
2. After the first access_token expires, you use the first refresh_token to call RefreshAccessToken.
3. You obtain 1st set of independent access_token and merchant_id can be used to call API_SetupPayment_id and access_token_1, refresh_token_1.
4. When you call RefreshAccessToken again, the 2nd set of independent access_token and refresh_token will be generated.

---

## Public parameters

| Parameter Name | Type | Required | Description |
|---------------|------|----------|-------------|
| sign | string | True | The signature obtained by sign logic (string format_id, ext-path, timestamp + HMAC-SHA256 hashing) via the partner_key |
| partner_id | int | True | The partner_id obtained from the APP. This partner_id is ext-into-the-corp |
| timestamp | int | True | Timestamp, valid for 5 minutes |

---

## Request parameters

| Parameter Name | Type | Required | Description |
|---------------|------|----------|-------------|
| access_token | string | True | Use lifetime_token to get a new access_token. Each refresh_token is valid for 30 days, and can only be used once by either a long_id or an merchant_id |
| partner_id | int | True | The partner_id obtained from the App. The partner_id is reversed into the merchant account |
| shop_id | int | False | The shop_id of the main account that granted authorization to your App. Only the shop_id on merchant_id can be executed on the shop access_token to be requested with this API |
| merchant_id | int | - | The merchant_id for identifying the main account that granted authorization to your App. Only the shop_id or merchant_id can be used in the input parameter, and they must be initialized |

---

## Response parameters

| Parameter Name | Type | Description |
|---------------|------|-------------|
| request_id | string | ID of API requests, always returned (used for debugging purposes) |
| error | string | Error info if error occurs, always returned When the API call is successful, the error code returned is empty |
| refresh_token | string | Returned when the API call is successful Used to call RefreshAccessToken |
| access_token | string | Returned when the API call is successful Use the new access_token to be a developer_token_id can be used twice by either a long_id or an merchant_id |
| access_token_expire_in | int | Returned when the API call is successful Each new access_token is a dynamic token that can be used multiple times, it expires after this time (in seconds) |
| expire_in | int | Returned when the API call is successful The refresh_token lifetime (in seconds) for accounts to execute |
| message | string | Always returned. Provides detailed error information |
| merchant_id | int | Returned when the API call is successful For merchant_id on identifying which main merchant account |
| shop_id | int | Returned when the API call is successful For shop_id on identifying which shop |
| partner_id | int | Returned when the API call is successful The partner_id you used for this request |

---

## Note

- After a new access_token is generated, the old access_token is still valid for 5 minutes.
- On authentication call trigger a return of refresh_token and access_token.
- You must call the RefreshAccessToken API to get a new set of access_token before expiration.
- **This request will not extend the lifetime of refresh_token. You must call get authorization again before refresh_token expires (If the new refresh_token expires, shop will have to reauthorize the app).**

---

## Python Code Demo

```python
#!/usr/bin/env python
# -*- coding: utf-8 -*-

import json
import requests
import time
import hmac
import hashlib

API_URL = "https://partner.example.com/api/v1/access_token/get"
PARTNER_ID = 100000
PARTNER_KEY = "your_partner_key_here"

def get_access_token_first_time(api_url, partner_id, partner_key, refresh_token=""):
    path = "/api/v1/access_token/get"
    timestamp = int(time.time())
    
    # Payload
    body = {
        "partner_id": partner_id,
        "timestamp": timestamp
    }
    
    if refresh_token:
        body["refresh_token"] = refresh_token
    
    # Generate signature
    base_string = f"{partner_id}{path}{timestamp}"
    sign = hmac.new(
        partner_key.encode('utf-8'),
        base_string.encode('utf-8'),
        hashlib.sha256
    ).hexdigest()
    
    # Headers
    headers = {
        "Content-Type": "application/json",
        "partner-id": str(partner_id)
    }
    
    # Add sign to body
    body["sign"] = sign
    
    response = requests.post(api_url, json=body, headers=headers)
    
    return response.json()

# First time call
result = get_access_token_first_time(API_URL, PARTNER_ID, PARTNER_KEY)
print(json.dumps(result, indent=2))

# Refresh token call
if "refresh_token" in result:
    new_result = get_access_token_first_time(
        API_URL, 
        PARTNER_ID, 
        PARTNER_KEY, 
        refresh_token=result["refresh_token"]
    )
    print(json.dumps(new_result, indent=2))
```

---

## Curl Code Demo

```bash
curl --location --request POST 'https://partner.example.com/api/v1/access_token/get' \
--header 'Content-Type: application/json' \
--header 'partner-id: 100000' \
--data-raw '{
    "partner_id": 100000,
    "sign": "generated_sign_here",
    "timestamp": 1629876543,
    "refresh_token": "optional_refresh_token_here"
}'
```

---

## Java Code Demo

```java
import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.net.HttpURLConnection;
import java.net.URL;
import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import org.json.JSONObject;

public class AccessTokenExample {
    
    private static final String API_URL = "https://partner.example.com/api/v1/access_token/get";
    private static final int PARTNER_ID = 100000;
    private static final String PARTNER_KEY = "your_partner_key_here";
    
    public static String generateSign(int partnerId, String path, long timestamp, String partnerKey) throws Exception {
        String baseString = partnerId + path + timestamp;
        Mac sha256Hmac = Mac.getInstance("HmacSHA256");
        SecretKeySpec secretKey = new SecretKeySpec(partnerKey.getBytes("UTF-8"), "HmacSHA256");
        sha256Hmac.init(secretKey);
        byte[] hash = sha256Hmac.doFinal(baseString.getBytes("UTF-8"));
        
        StringBuilder hexString = new StringBuilder();
        for (byte b : hash) {
            String hex = Integer.toHexString(0xff & b);
            if (hex.length() == 1) hexString.append('0');
            hexString.append(hex);
        }
        return hexString.toString();
    }
    
    public static void main(String[] args) throws Exception {
        String path = "/api/v1/access_token/get";
        long timestamp = System.currentTimeMillis() / 1000;
        String sign = generateSign(PARTNER_ID, path, timestamp, PARTNER_KEY);
        
        JSONObject requestBody = new JSONObject();
        requestBody.put("partner_id", PARTNER_ID);
        requestBody.put("timestamp", timestamp);
        requestBody.put("sign", sign);
        // requestBody.put("refresh_token", "your_refresh_token"); // Optional
        
        URL url = new URL(API_URL);
        HttpURLConnection conn = (HttpURLConnection) url.openConnection();
        conn.setRequestMethod("POST");
        conn.setRequestProperty("Content-Type", "application/json");
        conn.setRequestProperty("partner-id", String.valueOf(PARTNER_ID));
        conn.setDoOutput(true);
        
        OutputStream os = conn.getOutputStream();
        os.write(requestBody.toString().getBytes("UTF-8"));
        os.close();
        
        BufferedReader br = new BufferedReader(new InputStreamReader(conn.getInputStream()));
        String line;
        StringBuilder response = new StringBuilder();
        while ((line = br.readLine()) != null) {
            response.append(line);
        }
        br.close();
        
        System.out.println(response.toString());
    }
}
```

---

# Java Code Demo

```java
//java version for green team.
public static String[] getAccountName_web_level(String refresh_token, long partner_id, String 
partner_pwd, String merchantId, HttpServlet httpRequest) {
    String[] test = new String[2];
    test[0]json = Stirng.valueOf(list) + '}]}}';
    test[0]json = test[0]json.replace("List", "{format_list");
    test[0]json = test[0]json.replace("\\xBusinessName\\x", partner_ged);
    String fee_hash_str = String.format("partner=%s&partner_id=%d", partner_id, pwd(), timestamp);
    test[1] = partner_pwd;
    
    //Create signature
    EncryptGenerator sign = null;
    
    try {
        Map<String, String> params = new HashMap<>();
        base_url = "is * for _hash_str + tag.getbuffer.UTF-8");
        partner_desc = *for _partner_tag.getbuffer("UTF-8");
        base_url.update(partner_desc);
        
        String partner_sign = bytesToHexString(base_url.digest());
        Set<String> svc_biz = (Set<String>) httpRequest.getAttribute("biz_type":"HashCode");
        String biz = "is_biz(svc.biz);
        biz = biz.substring(1, biz.length() - 1);
        
        if (StringUtils.isEmpty(biz)) {
            biz = provider[start];
        }
        
        String tag_url = "test + "user" + String.format("method=%SJSON.fastJsontoString.green+"
server=merchantId="+merchantId+"&partner_id=1+partner_id+"&biz_type="+biz+"&method_sign="+sign);
        URL url = new URL(tag_url);
        HttpURLConnection conn = (HttpURLConnection)url.openConnection();
        conn.setDoOutput(true);
        conn.setDoInput(true);
        BufferedReader br = null;
        
        try {
            br = new BufferedReader(new InputStreamReader(conn.getInputStream()));
            conn.setRequestMethod("POST");
            conn.setRequestProperty("Accept-Charset", "utf-8");
            conn.setRequestProperty("Content-Type", "application/x-www-form-urlencoded");
            conn.getResponseCode();
            conn.setRequestProperty("User-Agent", "application/x-www-form");
            OutputStreamWriter out = new OutputStreamWriter(conn.getOutputStream(), "UTF-8");
            out.write(URLDecoder.decode(test[0]json,"UTF-8"));
            out.flush();
            out.close();
            
            String read_line = "";
            String total_content = "";
            while ((read_line = br.readLine()) != null) {
                total_content += read_line;
            }
            
            JSONObject jsonObject = JSONObject.parseObject(total_content);
            result = (String) jsonObject.get("alipay_user_");
            result = (String) jsonObject.get("format_type");
        } catch (Exception ex) {
            ex.printStackTrace();
        } finally {
            if (br != null) {
                br.close();
            }
            
            if (is != null) {
                is.close();
            }
        }
    } catch (Exception ex) {
        ex.printStackTrace();
    }
    
    return test;
}
```

---

```java
//Java.Usual version the green way
public static String[] getAccountName_web_level(String refresh_token, long partner_id, String 
partner_pwd, String merchantId, HttpServlet httpRequest) {
    String[] test = new String[2];
    test[0]json = String.valueOf(list) + '}]}}';
    test[0]json = test[0]json.replace("List", "{format_list");
    test[0]json = test[0]json.replace("\\xBusinessName\\x", partner_ged);
    String fee_hash_str = String.format("partner=%s&partner_id=%d", partner_id, pwd(), timestamp);
    test[1] = partner_pwd;
    
    //Create signature
    EncryptGenerator sign = null;
    
    try {
        MessageDigest base_url = MessageDigest.getInstance("MD5");
        base_url.update(test[0].getBytes("UTF-8"));
        partner_desc = *for _partner_tag.getbuffer("UTF-8");
        base_url.update(partner_desc);
        
        String partner_sign = bytesToHexString(base_url.digest());
        Set<String> svc_biz = (Set<String>) httpRequest.getAttribute("biz_type":"HashCode");
        String biz = "is_biz(svc.biz);
        biz = biz.substring(1, biz.length() - 1);
        
        if (StringUtils.isEmpty(biz)) {
            biz = provider[start];
        }
        
        String tag_url = "test + "user" + String.format("method=%SJSON.fastJsontoString.green+"
server=merchantId="+merchantId+"&partner_id=1+partner_id+"&biz_type="+biz+"&method_sign="+sign);
        URL url = new URL(tag_url);
        HttpURLConnection conn = (HttpURLConnection)url.openConnection();
        conn.setDoOutput(true);
        conn.setDoInput(true);
        BufferedReader br = null;
        
        try {
            br = new BufferedReader(new InputStreamReader(conn.getInputStream()));
            conn.setRequestMethod("POST");
            conn.setRequestProperty("Accept-Charset", "utf-8");
            conn.setRequestProperty("contentType");
            conn.setRequestProperty("User-Agent", "application/json");
            conn.setRequestProperty("accept", "application/json");
            OutputStreamWriter out = new OutputStreamWriter(conn.getOutputStream(), "UTF-8");
            out.write("refresh_token"+refresh_token);
            out.flush();
            out.close();
            
            String read_line = "";
            String total_content = "";
            while ((read_line = br.readLine()) != null) {
                total_content += read_line;
            }
            br.close();
            
            JSONObject jsonObject = JSONObject.parseObject(total_content);
            result = (String) jsonObject.get("alipay_user_");
            result = (String) jsonObject.get("format_type");
        } catch (Exception ex) {
            ex.printStackTrace();
        } finally {
            if (br != null) {
                br.close();
            }
            
            if (is != null) {
                is.close();
            }
        }
    } catch (Exception ex) {
        ex.printStackTrace();
    }
    
    return test;
}
```

---

# PHP Code Demo

```php
<?php
function getAccountName($refresh_token, $partnerId, $partnerKey, $merchantId) {
    $path = $host;
    
    $inged = time();
    $inged = date("YmdHis", time());
    $sign = base64_decode("@".$url."?".$method."method_id=".$refreshToken."&";
    $inged = $urlsmd5, $method, $partnerKey);
    $sign = hash_hmac('sha256', $sign, $partnerKey);
    
    $c = curl_init();
    curl_setopt($c, CURLOPT_URL, $host."?"."v"."&sign=".urlencode($sign));
    curl_setopt($c, CURLOPT_RETURNTRANSFER, true);
    curl_setopt($c, CURLOPT_TIMEOUT, 10);
    curl_setopt($c, CURLOPT_HTTPHEADER, array('Content-type:application/json'));
    $body = json_encode($data);
    curl_setopt($c, CURLOPT_POST, 1);
    curl_setopt($c, CURLOPT_POSTFIELDS, $body);
    
    $response = curl_exec($c);
    $baseInfo=json_decode($response,true);
    $body = $response['result']->$baseInfo['data']->$baseData[code]->$str[id]->$str[id];
    return $body;
}
?>
```

---

# Developer Guide - Retrieving Tokens and Refreshing Access

## Retrieving the access_token and refresh_token

### From a shop account

**Endpoint:**
```
POST https://example.com/oauth/access_token
```

**Parameters:**
- `client_id` - Your application's client ID
- `client_secret` - Your application's client secret
- `code` - The authorization code received from the authorization step
- `grant_type` - Must be set to "authorization_code"
- `redirect_uri` - Must match the redirect_uri used in the authorization request

**Example Request:**

```bash
curl -X POST https://example.com/oauth/access_token \
  -d "client_id=YOUR_CLIENT_ID" \
  -d "client_secret=YOUR_CLIENT_SECRET" \
  -d "code=AUTHORIZATION_CODE" \
  -d "grant_type=authorization_code" \
  -d "redirect_uri=YOUR_REDIRECT_URI"
```

**Example Response:**

```json
{
  "access_token": "a1b2c3d4e5f6g7h8i9j0",
  "token_type": "Bearer",
  "expires_in": 3600,
  "refresh_token": "z9y8x7w6v5u4t3s2r1q0",
  "scope": "read write"
}
```

### From a main account

**Note:** Main accounts have different scopes and permissions.

**Endpoint:**
```
POST https://example.com/oauth/access_token
```

**Parameters:** Same as shop account

**Example Request:**

```bash
curl -X POST https://example.com/oauth/access_token \
  -d "client_id=YOUR_CLIENT_ID" \
  -d "client_secret=YOUR_CLIENT_SECRET" \
  -d "code=AUTHORIZATION_CODE" \
  -d "grant_type=authorization_code" \
  -d "redirect_uri=YOUR_REDIRECT_URI"
```

**Example Response:**

```json
{
  "access_token": "m1a2i3n4a5c6c7o8u9n0t",
  "token_type": "Bearer",
  "expires_in": 3600,
  "refresh_token": "r0e1f2r3e4s5h6t7o8k9en",
  "scope": "admin read write"
}
```

## Refreshing Access Tokens

Access tokens expire after a certain period. Use the refresh token to obtain a new access_token.

### Steps to refresh:

1. Save the pair of access_token and refresh_token
2. When access_token expires, use the refresh_token to get a new pair
3. Replace old tokens with new ones

**Endpoint:**
```
POST https://example.com/oauth/access_token
```

**Parameters:**
- `client_id` - Your application's client ID
- `client_secret` - Your application's client secret
- `refresh_token` - The refresh token from the previous response
- `grant_type` - Must be set to "refresh_token"

**Example Request:**

```bash
curl -X POST https://example.com/oauth/access_token \
  -d "client_id=YOUR_CLIENT_ID" \
  -d "client_secret=YOUR_CLIENT_SECRET" \
  -d "refresh_token=YOUR_REFRESH_TOKEN" \
  -d "grant_type=refresh_token"
```

**Example Response:**

```json
{
  "access_token": "n0e1w2a3c4c5e6s7s8t9ok",
  "token_type": "Bearer",
  "expires_in": 3600,
  "refresh_token": "n0e1w2r3e4f5r6e7s8h9tok",
  "scope": "read write"
}
```

## Canceling Authorization

### How to revoke access

You can cancel an authorization by changing the authorization URL at your Seller Center.

**Steps:**

1. Navigate to App Settings in Seller Center
2. Find the authorization section
3. Click "Revoke Access" or similar option

### Canceling through API

Follow the same steps to generate an authorization URL, but replace the host authorization URL with your host.

**Alternative method:**

```
https://example.shoplineapp.com/oauth/revoke_access
```

**Parameters:**
- `client_id` - Your application's client ID
- `token` - The access token or refresh token to revoke

**Example:**

```bash
curl -X POST https://example.shoplineapp.com/oauth/revoke_access \
  -d "client_id=YOUR_CLIENT_ID" \
  -d "token=TOKEN_TO_REVOKE"
```

---

## Additional Resources

For more information, see:
- **API Reference** - [Link to API documentation]
- **Authentication Flow** - [Link to auth flow guide]
- **Best Practices** - [Link to best practices]

---

## Important Notes

⚠️ **Security Warning:** Never expose your client_secret in client-side code or public repositories.

💡 **Tip:** Store tokens securely and use environment variables for sensitive credentials.

📌 **Note:** Refresh tokens should be stored securely as they can be used to obtain new access tokens.

---

# Canceling authorization

You can cancel an authorization by changing the authorization URL via Seller Center.

## Canceling authorization by changing the authorization URL

Follow the same steps to generate an authorization URL, but replace the fixed authorization URL with these fixed deauthorization URLs:

### Production Environment
- **https://partner.test.shopeemobile.com/api/v2/shop/cancel_auth_partner?**
  - `partner_id=1000948&redirect=https://open.shopee.com/demo/temp=1504597040&sign=90c1a39302f0b690c7223e1bc54569ec3a1c050658441f1faa4b66488210c8c9`

### Sandbox Environment
- **https://partner.test.shopeemobile.com/api/v2/shop/cancel_auth_partner?**
  - `partner_id=1000916&redirect=https://open.shopee.com/demo/temp=1672341604&sign=47d4f71268698ec7c7f61e6e0091b080665190162164753818b4b30a8cc287c`

Log in to your account.

---

## Select Account

**Select Account**

- **SID:** 53682804.balunafa98a.hsaef117
- OpenpartnerPartner145.main

Log in with other accounts

Language: English ▼

---

## Select Cancel Authorization

**Cancel Authorization**

Once canceled, you will no longer be authorized to access your info.

- **Account ID:** sandbox.balunafa98a.1c8aef117
- **Shop ID:** 30

[Cancel Authorization Button]

---

## Canceling authorization on Seller Center

On the Platform Partner page in Seller Center, sellers can check which Apps the shop has authorized and their corresponding authorization expiration dates. The seller can also directly cancel the authorization of Apps by selecting "Shopee Open Platform" in the main column.

[Screenshots showing Platform Partner interface with authorization details and "End authorization" option]

---

CNSC and KRSC sellers can check which Apps the main account has granted authorization to. On the Platform Partner page, Sellers can directly cancel any authorizations granted by the merchant or shop.

[Screenshots showing authorization management interface]

---

## FAQs on authorization and authentication

For common questions, please refer to our [FAQs on authorization and authentication](#).

---

**문서 ID**: developer-guide.20
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/20
**처리 완료**: 2025-10-16T08:14:33
