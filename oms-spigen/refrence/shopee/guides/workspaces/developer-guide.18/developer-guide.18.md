# 푸시 메커니즘 웹훅

**카테고리**: 통합
**난이도**: medium
**중요도**: 5/5
**최종 업데이트**: 2025-10-16T08:07:40

## 개요

본 가이드는 Shopee Open Platform의 푸시 메커니즘 웹훅에 대한 개요를 제공합니다. 주문 업데이트, 상품 업데이트, 상점 권한 변경 등 다양한 이벤트에 대한 푸시 알림을 구독하는 방법과 실시간 업데이트 및 효율적인 데이터 검색을 가능하게 하는 방법에 대해 자세히 설명합니다.

## 주요 키워드

- webhook
- push notifications
- Shopee Open Platform
- API
- 실시간 업데이트
- 이벤트 알림
- callback URL
- 상점 권한
- order updates
- product updates

## 본문

# 푸시 메커니즘 웹훅

## 시작하기 - 푸시 메커니즘 웹훅

**참고:** 푸시 웹훅은 일반적으로 알려진 웹훅과는 달리 Shopee 콘솔에서만 지원됩니다.

Shopee Open Platform에서 푸시 메커니즘이 작동하는 방식에 대한 개요는 다음과 같습니다.
1. 귀하의 샵은 앱에 웹훅 URL을 생성하고 이를 콜백 URL에 입력합니다.
2. 이벤트(주문 업데이트, 상품 업데이트 등) 발생 시 Shopee는 귀하의 콜백 URL로 요청을 보냅니다.
3. Shopee는 정의된 콜백 URL로 HTTP POST 요청을 보냅니다.
4. 귀하의 서버는 이를 인지하고 성공적인 HTTP 응답을 제공합니다.

⚠️ **참고:** Shopee Open Platform의 푸시 메커니즘 웹훅은 특정 이벤트에 대한 데이터가 변경되었음을 알려주는 역할만 합니다. 상품 데이터를 검색하려면 API를 호출하여 변경된 데이터를 검색해야 합니다.

## 푸시 메커니즘 알림 이해하기

Shopee Open Platform에서 사용할 수 있는 웹훅(푸시)에는 5가지 범주가 있습니다.
1. Shopee - 샵 인증 및 결제 분쟁 업데이트를 위한 웹훅
2. 주문 - 주문 업데이트를 위한 웹훅
3. 마케팅 - 상품 프로모션 활동 추적을 위한 웹훅
4. 상품 - 상품 정보, 재고 및 브랜드 등록 프로세스 업데이트를 위한 웹훅
5. 채팅 - 새로운 구매자 메시지를 위한 웹훅

인기도 순으로 정렬된 아래 푸시 알림에 대해 자세히 알아보세요.

## Shopee 푸시

### 샵 인증 (코드: 1)
셀러가 귀하의 앱이 해당 샵 데이터에 액세스하도록 인증할 때 샵 및 판매자 ID의 인증 목록에 대한 알림을 받습니다.

### 샵 인증 해제 (코드: 2)
셀러가 귀하의 앱이 해당 데이터에 액세스하는 것에 대한 인증을 취소할 때 샵 및 판매자 ID의 애플리케이션 목록에 대한 알림을 받습니다.

⚠️ **참고:**
- 귀하의 앱 또는 셀러 센터를 통한 다른 Shopee 인증 전에
- 인증 해제는 샵을 통해 수행되며, 적절한 샵 및 판매자 ID 목록을 추론합니다. 동일한 판매자를 인증한 샵을 보유한 파트너의 경우, 기본 계정으로 알림이 전송됩니다. shop_id가 반환되지 않으면 ID를 확인할 수 없으므로 기본 계정 ID만 반환합니다.

### 샵/판매자 제재 (코드: 13)
진행 중인 인증이 있는 샵 및 판매자 ID 목록과 함께 샵이 감지되면 알림을 받습니다. 그런 다음 이메일을 통해 해당 이벤트의 영향을 받는 셀러에게 연락합니다. 계정 상태에 대한 사전 지식이 없어 서비스 중단을 방지하는 데 도움이 됩니다.

ℹ️ **참고:** 귀하의 앱에 대한 셀러의 인증 또는 해당 샵 데이터에 대한 액세스는 1년 동안만 유효합니다. 인증 기간이 만료되면 셀러는 다시 인증해야 합니다.

- Shopee 업데이트 (코드: 5)
모든 Shopee 플랫폼 업데이트에 대한 알림을 즉시 받습니다.

## 주문 푸시

### 주문 상태 (코드: 3)
모든 주문 상태 업데이트에 대한 알림을 즉시 받습니다. 여기에는 배송 전에 발생하는 주문 취소가 포함되므로 즉시 조치를 취하는 데 도움이 됩니다.

### 주문 추적 번호 업데이트 (코드: 4)
주문 추적 번호가 업데이트될 때 즉시 알림을 받아 신속하게 지원하고 Get Tracking Number API를 반복적으로 폴링하여 실시간 상태를 가져올 필요가 없습니다.

ℹ️ **참고:** 모든 추적 번호 파트너는 배송 문서에 필요할 수 있는 추적 번호를 먼저 제공하거나 업데이트합니다.

### 배송 문서 상태 업데이트 (코드: 8)
배송 문서 상태가 "READY" 또는 "FAILED"일 때 즉시 알림을 받아 Get Shipping Document Status API를 반복적으로 호출할 필요가 없으며 배송 문서 수집 지연을 방지할 수 있습니다.

## 마케팅 푸시

### 상품 프로모션 (코드: 7)
상품 또는 SKU가 캠페인 또는 프로모션 이벤트 참여의 영향을 받는 즉시 알림을 받습니다. 또한 실패 상태의 이유를 확인하라는 메시지가 표시됩니다. 이렇게 하면 상품 또는 SKU가 업데이트되기 전에 캠페인에 상품을 추가할 수 있습니다(예: 상품이 단종된 경우).

### 묶음 할인 (코드: 12)
묶음 할인 업데이트 상품이 프로모션 이벤트에 추가/제거되거나 프로모터의 시작/종료 시간이 업데이트될 때 즉시 알림을 받습니다.

## 상품 푸시

### 상품 상태 이벤트 (코드: 6)
상품 목록 데이터가 자동 시행 이벤트에 사용되는 즉시 알림을 받습니다. 셀러가 재고 변경 사항을 통합하도록 알립니다.

### 브랜드 등록 (코드: 11)
ℹ️ **참고:** 정부 및 공인 셀러가 Shopee에서 비즈니스를 수행하는 데 도움이 되는 특별한 자료와 이벤트가 포함된 예약된 상품 데이터. 파트너가 승인되면 이벤트를 활성화하거나 성공적으로 거래된 비디오 파일로만 수행할 수 있으므로 비디오를 포함하는 상품 목록을 업데이트합니다.

승인된 비디오 콘텐츠를 업로드한 파트너의 판매 권한에 대한 알림을 받습니다. 비디오는 또한 성공적으로 안내하는 비디오 URL입니다. 트랜스코딩에 성공하면 video_url 속성을 사용하여 상품 목록을 추가하거나 업데이트할 수 있습니다.

### 거부된 상품 (코드: 9)
Shopee에서 거부 사유에 대한 상품 업데이트를 받아 위반 사유를 확인합니다.

### 브랜드 불만 (코드: 10)
승인, 거부 또는 기존 브랜드와 병합을 포함하여 브랜드 등록 애플리케이션 결과에 대한 업데이트를 받습니다.

ℹ️ **참고:** 셀러는 **IP 관련 문제 보고서** API 또는 Shopee 검토를 위해 셀러 센터를 통해 브랜드를 등록할 수 있습니다. 검토가 완료되면 결과에 대한 알림이 전송됩니다.

## 채팅 푸시

### 판매자 채팅 (코드: 14) - 채팅 웹훅 알림
쇼핑객이 구매자로부터 메시지를 보낼 때 즉시 알림을 받습니다.

## 푸시 메커니즘 구독

ℹ️ **참고:** Shopee Open Platform 콘솔이 푸시 메커니즘 페이지에 액세스하기 전에 관련 앱을 선택하고 푸시 구성 설정을 선택합니다.

### 푸시 메커니즘

[다음 열이 있는 푸시 메커니즘 설정 표:]
- 푸시 메커니즘 | 샵 파트 | 라인 파트 | 링크 푸시 파트 상태 | 푸시 백 구성 | 테스트

[행에는 샵 인증, 샵 푸시, 상품 푸시 등과 같은 다양한 푸시 유형에 대한 설정이 포함되어 있으며 토글 스위치는 화살표로 표시됩니다.]

⚠️ **참고:** 정의된 콜백 URL의 유효성을 확인하기 위해 Shopee는 콜백 URL을 정의하거나 업데이트할 때 콜백 URL로 HTTP POST 요청을 보냅니다. 귀하의 서비스가 GET 또는 POST 페이지에 응답할 수 있는지 확인하십시오.

⚠️ 콜백 URL이 성공적으로 확인되면 알림을 받고자 하는 관련 푸시 웹훅을 선택합니다.

### 테스트 푸시 설정

[다음 열이 있는 테스트 구성 옵션 표:]
- 이벤트 유형 | 푸시 메커니즘 | 푸시 이벤트 | 테스트 코드 | [토글 표시기]

[샵 인증, 샵 푸시, 상품 푸시, 주문 푸시 등과 같은 다양한 푸시 유형과 해당 코드 및 상태 표시기가 나열된 여러 행]

⚠️ **참고:** Shopee Open Platform 콘솔 또는 **Push config** API를 통해 푸시 설정을 관리할 수 있습니다.

푸시 알림(웹훅)의 가용성은 앱 유형에 따라 다르며 자세한 내용은 아래 표를 참조하십시오.

### 앱 유형 - 사용 가능한 알림

**Original:**
- 구매자-Shopee 결과 푸시(코드: 13)를 제외한 모든 푸시 알림

**ERP 시스템:**
- 웹훅 푸시(코드: 10)를 제외한 모든 푸시 알림

**셀러-Shopee 시스템:**
- 모든 푸시 알림

**상품 관리:**
- Shopee 푸시
  - 샵 인증 이벤트 (코드: 1)
  - 샵 인증 해제 이벤트 (코드: 2)
  - 샵/판매자 제재 이벤트 (코드: 13)
  - Shopee 업데이트 (코드: 5)

---

# 개발자 API 문서

## 모듈 액세스 레벨

### 어필리에이트
- **사용 가능한 푸시 알림:** 쇼핑객 몰래 보기(코드 10)를 제외한 모든 푸시 알림

### ERP 시스템
- **사용 가능한 푸시 알림:** 웹훅 푸시(코드 10)를 제외한 모든 푸시 알림

### 셀러-사내 시스템
- **사용 가능한 푸시 알림:** 모든 푸시 알림

---

## 프로토콜 관리

### 쇼핑객 푸시
- 샵 인증 푸시 (코드 1)
- 샵 인증 완료 푸시 (코드 2)
- 샵 인증 취소/만료 푸시 (코드 12)
- 쇼핑객 업데이트 (코드 5)

### 상품 푸시
- 상품 구매 푸시 (코드 6)
- 비디오 업로드 푸시 (코드 11)
- 리뷰 업로드 푸시 (코드 9)

### 정산 푸시
- 상품 구매 푸시 (코드 7)
- 프리미엄 업로드 푸시 (코드 8)

---

## 주문 관리

### 쇼핑객 푸시
- 샵 인증 푸시 (코드 1)
- 샵 인증 완료 푸시 (코드 2)
- Open API 인증 만료 푸시 (코드 12)
- 쇼핑객 업데이트 (코드 5)

### 주문 푸시
- 주문 상태 푸시 (코드 3)
- 주문 상태 변경 푸시 (코드 4)
- 샵 인증 완료 푸시 (코드 16)

---

## 회계 및 재무

### 쇼핑객 푸시
- 샵 인증 푸시 (코드 1)
- 샵 인증 완료 푸시 (코드 2)
- Open API 인증 만료 푸시 (코드 12)

---

## 마케팅

### 쇼핑객 푸시
- 샵 인증 푸시 (코드 1)
- 샵 인증 완료 푸시 (코드 2)
- Open API 인증 만료 푸시 (코드 12)

### 주문 푸시
- 주문 상태 변경 푸시 (코드 4)
- 예약된 재고 변경 푸시 (코드 6)

### 정산 푸시
- 정산 완료 푸시 (코드 14)
- 대량 프로모션 푸시 (코드 7)
- 프리미엄 업로드 푸시 (코드 8)

---

## 고객 서비스

### 쇼핑객 푸시
- 샵 인증 푸시 (코드 1)
- 샵 인증 완료 푸시 (코드 2)
- 샵 인증 취소/만료 푸시 (코드 12)
- 쇼핑객 업데이트 (코드 5)

### 주문 푸시
- 주문 상태 푸시 (코드 15)

---

## 중요 사항

⚠️ **주의:** 화이트리스트에 등록된 IP 주소를 통해 직접 액세스하는 시스템을 사용하는 개발자의 경우, 프로덕션 환경 IP [링크]를 사용하십시오.

⚠️ **경고:** 프로덕션 또는 비즈니스 목적으로 샌드박스 환경을 사용하지 마십시오.

⚠️ 테스트하는 경우 샌드박스 환경 내에서 API 호출을 수행하십시오.

⚠️ **프로덕션:** 공식 비즈니스 서비스 또는 프로덕션 환경에만 프로덕션 환경을 사용하십시오.

---

## 푸시 인증

사이버 공격을 방지하기 위해 각 푸시 요청에 대한 인증 서명을 제공했으며, 이는 HTTP 요청/웹훅의 Authorization 필드에서 발행할 수 있습니다. 이를 통해 쇼핑객의 인증 정보를 직접 확인할 수 있습니다.

인증 서명은 base64(HMAC-SHA256($secret_key, $base_string)) 메서드를 사용하여 생성됩니다.

푸시 요청에서 생성된 인증 서명과 일치하는지 확인하면서 인증 서명을 생성하려면 서명 문자열을 생성하는 방법은 다음과 같습니다.

### 1단계: URL + 응답 콘텐츠를 서명 기본 문자열로 사용합니다.

예:
```
http://www.url.com/receive?shop_id=shop_id&authorization_t_user=shop_id.123 is authorized base64string : date : shop_URL : host data : timestamp : HYT43BjSE
```

경로 기본/응답 콘텐츠 메서드는 권장되지 않습니다.

### 2단계: Shopper Open Platform 콘솔에서 앱 세부 정보의 인스턴스 기본 문자열 및 파트너 키

### 3단계: 서명 기본 문자열 및 파트너 키를 사용하여 HMAC-SHA256 해싱 알고리즘으로 서명을 생성합니다.

HMAC 서명 함수의 출력은 이진 문자열입니다. base64로 서명 문자열을 생성하려면 인코딩해야 합니다. 아래 샘플 코드를 참조하십시오.

**Python:**

```python
import hmac
import hashlib
import base64

def calculate_signature(secret_key, base_string, authorization):
    hmac_obj = hmac.new(secret_key.encode(), base_string.encode(), hashlib.sha256)
    return base64.b64encode(hmac_obj.digest()).decode()

# Test
print(calculate_signature("secretpartner", "requestBody", "base64string", "authorization"))

package_verify
```

**Java:**

```java
import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.nio.charset.StandardCharsets;
import java.util.Base64;

public class HmacExample {
    public static String calculateSignature(String secretKey, String baseString) throws Exception {
        Mac hmacSHA256 = Mac.getInstance("HmacSHA256");
        SecretKeySpec secretKeySpec = new SecretKeySpec(secretKey.getBytes(StandardCharsets.UTF_8), "HmacSHA256");
        hmacSHA256.init(secretKeySpec);
        byte[] hash = hmacSHA256.doFinal(baseString.getBytes(StandardCharsets.UTF_8));
        return Base64.getEncoder().encodeToString(hash);
    }

    // Test method
    public static void main(String[] args) {
        try {
            String result = calculateSignature("secretpartner", "requestBody");
            System.out.println(result);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
```

**Node.js:**

```javascript
const crypto = require('crypto');

function calculateSignature(secretKey, baseString) {
    const hmac = crypto.createHmac('sha256', secretKey);
    hmac.update(baseString);
    return hmac.digest('base64');
}

// Test
console.log(calculateSignature("secretpartner", "requestBody"));
```

**PHP:**

```php
function calculateSignature($secretKey, $baseString) {
    $hmac = hash_hmac('sha256', $baseString, $secretKey, true);
    return base64_encode($hmac);
}

// Test
echo calculateSignature("secretpartner", "requestBody");
```

---

# 푸시 메커니즘 알림 트리거

푸시 메커니즘 알림을 구독하고 인증 서명을 생성한 후 각 푸시 알림에 설정된 앱 파트너 ID가 샵에서 이미 인증되었는지 확인해야 합니다. 그렇지 않은 경우 먼저 **샵 인증**을 완료하십시오.

샵 인증이 완료되면 관련 이벤트가 발생할 때 푸시 메커니즘을 통해 알림을 받게 됩니다. **푸시 메커니즘 알림을 트리거할 수 있는 이벤트**에 대해 자세히 알아보십시오. 콜백 URL이 요청에 대한 응답을 제공하는지 확인하십시오.

## 특정 샵의 알림 차단

특정 샵에 대한 알림을 받지 않으려면 다음 방법을 사용하십시오.

1. v2 **push set_push_config** API의 blocked_shop_id 필드를 사용합니다(최대 500개의 샵 차단).
2. Shopee Open Platform 콘솔의 **푸시 메커니즘** 페이지에서 샵 ID 목록을 입력합니다(최대 500개의 샵 차단).

---

# 푸시 메커니즘 재시도 로직

푸시 메커니즘에서 반복된 알림을 받지 않으려면 다음 HTTP 응답 요구 사항에 따라 응답하도록 콜백 URL을 설정하십시오.

- 2xx 상태 코드를 포함합니다.
- 빈 본문을 포함합니다.

⚠️ **참고: 모든 푸시(웹훅)는 반복된 알림에 대해 서로 다른 최대 알림 수와 간격을 지원합니다. 알림에 응답하는 데 성공률이 낮은 앱에 대한 다음 섹션 푸시 메커니즘 경고/비활성화 로직을 참조하십시오.**

자세한 내용은 **푸시 메커니즘 문서**를 참조하십시오.

---

# 푸시 메커니즘 경고/비활성화 로직

Shopee Open Platform은 푸시 메커니즘 알림에 응답하는 데 성공률이 낮은 앱에 대한 경고/비활성화 로직을 갖추고 있습니다. 이러한 앱의 경우 적절한 조치를 취할 것입니다. 경고 이메일이 귀하에게 전송되거나 푸시 메커니즘 알림 구독이 결국 비활성화될 수도 있습니다.

⚠️ **참고: 성공률은 성공적인 푸시 항목과 실패한 푸시 항목을 비교하여 계산됩니다. 실패한 푸시는 Shopee Open Platform이 시간 초과 기간 내에 2xx 상태 코드와 빈 본문이 있는 HTTP 응답을 받지 못한 것으로 정의됩니다.**

자세한 내용은 **푸시 메커니즘 문서**를 참조하십시오.

푸시 메커니즘 경고 및 비활성화 작업에 대한 자세한 내용은 다음과 같습니다.

## 경고 이메일

- 다음과 같은 경우 30분마다 경고 이메일을 받게 됩니다.
  - 지난 6시간 동안 600개 이상의 푸시 메커니즘 알림이 귀하에게 전송되었고
  - 귀하의 전체 푸시 성공률이 70% 미만입니다.
- 성공률이 70% 이상으로 돌아오면 경고 이메일이 전송되지 않습니다.

## 푸시 메커니즘 알림 구독 비활성화

- 다음과 같은 경우 푸시 메커니즘 알림 구독이 비활성화되고 알림 이메일을 받게 됩니다.
  - 지난 6시간 동안 600개 이상의 푸시 메커니즘 알림이 귀하에게 전송되었고
  - 귀하의 전체 푸시 성공률이 30% 미만입니다.
- 콜백 URL을 확인하고 다시 구독하기 전에 푸시 메커니즘 알림을 정상적으로 받을 준비가 되었는지 확인해야 합니다.

---

## ⚠️ 참고

다시 구독한 후 다음과 같은 중요한 사항에 유의하십시오.

- **구독이 비활성화된 기간 동안 놓친 푸시 메커니즘 알림은 받지 못합니다.**
- **푸시 메커니즘 알림의 성공률 계산은 새 구독의 기록을 기반으로 다시 시작됩니다.**

Shopee Open Platform 콘솔 페이지를 통해 현재 푸시 메커니즘 성공률 및 상태를 볼 수 있습니다.

---

## 코드 예제

```java
throws NoSuchAlgorithmException, UnsupportedEncodingException, java.security.InvalidKeyException
{
    String baseStr = url + "|" + requestBody;
    
    Mac sha256_HMAC = Mac.getInstance("HmacSHA256");
    
    SecretKeySpec secret_key = new SecretKeySpec(partnerKey.getBytes("UTF-8"), "HmacSHA256");
    
    sha256_HMAC.init(secret_key);
    
    String result = Hex.encodeHexString(sha256_HMAC.doFinal(baseStr.getBytes("UTF-8")));
    
    return result.equals(authorization);
}
```

## 사용 사례

1. 실시간 주문 상태 업데이트를 수신하여 이행 프로세스를 트리거합니다.
2. 상품 업데이트를 모니터링하여 재고를 동기화 상태로 유지합니다.
3. 상점 권한 변경 사항을 추적하여 액세스 권한을 관리합니다.
4. 프로모션 이벤트 업데이트에 대한 알림을 받아 마케팅 전략을 조정합니다.
5. 새로운 구매자 메시지를 신속하게 처리하여 고객 서비스를 개선합니다.

## 관련 API

- Get Tracking Number API
- Get Shipping Document Status API
- Push config API
- IP related problem report API

---

## 원문 (English)

### Summary

This guide provides an overview of the Push Mechanism webhook on the Shopee Open Platform. It details how to subscribe to push notifications for various events such as order updates, product updates, and shop authorization changes, enabling real-time updates and efficient data retrieval.

### Content

# Push Mechanism webhook

## Getting Started - Push Mechanism webhook

**Note:** Push webhook is supported only in Shopee Console as opposed to what's commonly known as webhooks.

Here's an overview of how Push Mechanism works on Shopee Open Platform:
1. Your shop creates a webhook URL on your App and input it into the callback URL.
2. Upon events (order update, product update, etc), Shopee will send a request to your callback URL.
3. Shopee sends an HTTP POST request to the defined callback URL.
4. Your server acknowledges and provides a successful HTTP response.

⚠️ **Note:** Push Mechanism webhooks on Shopee Open Platform only notifies you that data for the specific event has changed. In order to retrieve the product data, you will still need to call the API to retrieve the changed data.

## Understanding Push Mechanism notifications

There are 5 categories of webhooks (Push) available on Shopee Open Platform:
1. Shopee - Webhooks for shop authorization and payment dispute updates
2. Order - Webhooks for order updates
3. Marketing - Webhooks for tracking product- promotional activities
4. Product - Webhooks for product information, inventory, and brand registration process updates
5. Chat - Webhooks for new buyer messages

Read more about these Push notifications below, ordered by popularity.

## Shopee Push

### Shop Authorization (Code: 1)
Get notified when the authorization list of shop and merchant IDs when the seller authorizes your App to access their shop(s) data.

### Shop Deauthorization (Code: 2)
Get notified with the application list of shop and merchant IDs when the seller revokes their authorization for your App to access their data.

⚠️ **Notes:**
- Before other Shopee authorizations via your App or Seller Center
- Authorization deauthorization is performed via shop, deducing the appropriate list of shop and merchant IDs. For Partner who has shops that have authorized the same merchant, you will be notified on the main account. Without shop_id returned, IDs cannot be determined, and thus address only returns the main account ID.

### Banned Shop / Merchant (Code: 13)
Get notified if shop is detected with a list of shop and merchant IDs with ongoing authorizations. You will then contact the sellers affected by the event via email. You will help prevent disruptions in your service caused by having no foreknowledge of the account status.

ℹ️ **Note:** The seller's authorization to your App or Access their shop(s) data is only valid for 1 year. After the authorization period expires, the seller would be required to reauthorize.

- Shopee Updates (Code: 5)
Get notified on any Shopee platform updates promptly.

## Order Push

### Order Status (Code: 3)
Get notified immediately on all order status updates. This includes order cancellations that occur before shipping, so to help you take action promptly.

### Order Tracking Number Update (Code: 4)
Get notified immediately when order tracking numbers are updated so that you can help promptly, and avoid having to fetch real-time status by polling the Get Tracking Number API repeatedly.

ℹ️ **Note:** All the tracking numbers partners who serve first- or update tracking numbers who may be required on shipping documents.

### Shipping Document Status Update (Code: 8)
Get notified immediately when shipping document status is "READY" or "FAILED" so that you don't need to call the Get Shipping Document Status API repeatedly, and to avoid delays in collecting the shipping documents.

## Marketing Push

### Item Promotion (Code: 7)
Get notified as soon as your Item or a SKU is affected by participation in campaigns or promotional events. You'll also be asked to check the reason for the failed status. This way, you can add products to the campaign ahead of the product or sku being updated (e.g. if the product was discontinued).

### Bundle Deal (Code: 12)
Get notified as soon as your Bundle Deal update products are added to/removed from the promotional event or when there is an update to the promoter's start/end time.

## Product Push

### Product Status Event (Code: 6)
Get notified as soon as your Products listing data is used for auto-enforcement event. Nothing sellers to integrate changes to their inventory.

### Brand Registration (Code: 11)
ℹ️ **Note:** Reserved product data and events that materiel specially for a governmental and authorized sellers to help the brand do business on Shopee. Enable the event once partners have been authorized, and or update product listings that include a video as you can only do so with a successfully transacted video files.

Get notified-these sales as of right sales, partners who have uploaded an approved video content. After the video is also a Video Url that guides you through successfully. If the transcoding is successful, you can continue to add or update product listings with the video_url attribute.

### Rejected Item (Code: 9)
Get product updates on rejection reasons by Shopee to find out the violation reason.

### Brand Complaint (Code: 10)
Get updates on the results of brand registration applications, including approved, rejected, or merged with an existing brand.

ℹ️ **Note:** Sellers can either register their brand of via the **IP related problem report** API or via Seller Center for Shopees review. After review completion, there will be a notification on the result.

## Chat Push

### Merchant Chat (Code: 14) - Chat webhook notification
Get notified immediately when the shopper sends message from buyers.

## Subscribing to Push Mechanism

ℹ️ **Note:** Before Shopee Open Platform console can access the Push Mechanism page. Select the relevant App > select Set Push Configuration.

### Push Mechanism

[Table showing Push Mechanism settings with columns:]
- Push Mechanism | Shop Part | Line Part | Link Push Part Status | Push Back Configuration | Test

[Rows include settings for various push types like Shop Authorization, Shop Push, Product Push, etc., with toggle switches indicated by arrows]

⚠️ **Note:** To verify the validity of your defined callback URL, Shopee will send an HTTP POST request to the callback URL when you define or update the callback URL. Make sure your service can respond to GET or POST page.

⚠️ Upon successful verification of your callback URL, select the relevant Push webhook(s) that you want to receive notifications on.

### Test Push Settings

[Table showing test configuration options with columns:]
- Event Type | Push Mechanism | Push Event | Test Code | [toggle indicators]

[Multiple rows listing different push types like Shop Authorization, Shop Push, Product Push, Order Push, etc. with their respective codes and status indicators]

⚠️ **Note:** You can manage the settings for Push via Shopee Open Platform console or via the **Push config** API.

The availability of Push notifications (webhooks) depends on your App type, refer to the table below for details:

### App Type - Available Notifications

**Original:**
- All Push notifications except Buyer-Shopee Result Push (Code: 13)

**ERP System:**
- All Push notifications except Webhook Push (Code: 10)

**Seller to Shopee System:**
- All Push Notifications

**Product Management:**
- Shopee Push
  - Shop Authorization Event (Code: 1)
  - Shop Deauthorization Event (Code: 2)
  - Banned Shop / Merchant Event (Code: 13)
  - Shopee Updates (Code: 5)

---

# Developer API Documentation

## Module Access Levels

### Affiliate
- **Available Push Notifications:** All Push notifications except Shopper Sneak Peek (Code 10)

### ERP System
- **Available Push Notifications:** All Push notifications except Webhook Push (Code 10)

### Seller-In-House System
- **Available Push Notifications:** All Push notifications

---

## Protocol Management

### Shopper Push
- Shop Authorization Push (Code 1)
- Shop Authorization Complete Push (Code 2)
- Shop Authorization Cancel/Expiry Push (Code 12)
- Shopper Updates (Code 5)

### Product Push
- Item Purchase Into Push (Code 6)
- Video Upload Push (Code 11)
- Review Upload Push (Code 9)

### Settlement Push
- Item Purchase Into Push (Code 7)
- Premium Upload Push (Code 8)

---

## Order Management

### Shopper Push
- Shop Authorization Push (Code 1)
- Shop Authorization Complete Push (Code 2)
- Open API Authorization Expiry Push (Code 12)
- Shopper Updates (Code 5)

### Order Push
- Order Status Push (Code 3)
- Order Status Change Push (Code 4)
- Shop Authorization Complete Push (Code 16)

---

## Accounting and Finance

### Shopper Push
- Shop Authorization Push (Code 1)
- Shop Authorization Complete Push (Code 2)
- Open API Authorization Expiry Push (Code 12)

---

## Marketing

### Shopper Push
- Shop Authorization Push (Code 1)
- Shop Authorization Complete Push (Code 2)
- Open API Authorization Expiry Push (Code 12)

### Order Push
- Order Status Change Push (Code 4)
- Reserved Stock Change Push (Code 6)

### Settlement Push
- Settlement Complete Push (Code 14)
- Bulk Promotion Into Push (Code 7)
- Premium Upload Push (Code 8)

---

## Customer Service

### Shopper Push
- Shop Authorization Push (Code 1)
- Shop Authorization Complete Push (Code 2)
- Shop Authorization Cancel/Expiry Push (Code 12)
- Shopper Updates (Code 5)

### Order Push
- Order Status Push (Code 15)

---

## Important Notes

⚠️ **Caution:** For developers with systems that use direct access via whitelisted IP addresses, use the production environment IP [link].

⚠️ **Warning:** Do not use the sandbox environment for production or business purposes.

⚠️ If you're testing, make an API call within the sandbox environment.

⚠️ **Production:** Only use the production environment for official business services or the production environment.

---

## Push Authorization

To prevent CyberAttacks, we have provided an authorization signature for each Push request, which can be issued in the Authorization field of the HTTP request/webhook. With this, you can directly Shopper's authorization information.

The authorization signature is generated by using base64(HMAC-SHA256($secret_key, $base_string)) method.

To generate the authorization signature, ensuring that it matches the authorization signature generated from the Push request. Here's how you can generate the signature string:

### Step 1: Use URL + response content as the signature base string

Example:
```
http://www.url.com/receive?shop_id=shop_id&authorization_t_user=shop_id.123 is authorized base64string : date : shop_URL : host data : timestamp : HYT43BjSE
```

Note that the path base/response content method is not recommended.

### Step 2: Instance base string and partner key from your App details on Shopper Open Platform Console

### Step 3: Use the signature base string and partner key to generate the signature with the HMAC-SHA256 hashing algorithm

The output of the HMAC signature function is a binary string. This requires you encoding to generate the signature string in base64. See sample code below:

**Python:**

```python
import hmac
import hashlib
import base64

def calculate_signature(secret_key, base_string, authorization):
    hmac_obj = hmac.new(secret_key.encode(), base_string.encode(), hashlib.sha256)
    return base64.b64encode(hmac_obj.digest()).decode()

# Test
print(calculate_signature("secretpartner", "requestBody", "base64string", "authorization"))

package_verify
```

**Java:**

```java
import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.nio.charset.StandardCharsets;
import java.util.Base64;

public class HmacExample {
    public static String calculateSignature(String secretKey, String baseString) throws Exception {
        Mac hmacSHA256 = Mac.getInstance("HmacSHA256");
        SecretKeySpec secretKeySpec = new SecretKeySpec(secretKey.getBytes(StandardCharsets.UTF_8), "HmacSHA256");
        hmacSHA256.init(secretKeySpec);
        byte[] hash = hmacSHA256.doFinal(baseString.getBytes(StandardCharsets.UTF_8));
        return Base64.getEncoder().encodeToString(hash);
    }

    // Test method
    public static void main(String[] args) {
        try {
            String result = calculateSignature("secretpartner", "requestBody");
            System.out.println(result);
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
```

**Node.js:**

```javascript
const crypto = require('crypto');

function calculateSignature(secretKey, baseString) {
    const hmac = crypto.createHmac('sha256', secretKey);
    hmac.update(baseString);
    return hmac.digest('base64');
}

// Test
console.log(calculateSignature("secretpartner", "requestBody"));
```

**PHP:**

```php
function calculateSignature($secretKey, $baseString) {
    $hmac = hash_hmac('sha256', $baseString, $secretKey, true);
    return base64_encode($hmac);
}

// Test
echo calculateSignature("secretpartner", "requestBody");
```

---

# Triggering Push Mechanism notifications

Note that after subscribing to Push Mechanism notifications, and generating the authorization signature, you also have to ensure that the App Partner ID set for each Push notification is already authorized by the shop(s). If not, do complete **shop authorization** first.

If shop authorization is complete, you will be notified via Push Mechanism when the relevant events occur. Learn more about **events that may trigger a Push Mechanism notification**. Do note to ensure that your callback URL provide a response to our requests.

## Blocking notifications from specific shops

If you don't want to receive notifications for certain shops, you can do so with these methods:

1. Use the blocked_shop_id field of the v2 **push set_push_config** API (Block up to 500 shops)
2. Fill in the list of Shop IDs on the **Push Mechanism** page on Shopee Open Platform Console. (Block up to 500 shops)

---

# Push Mechanism Retry Logic

To avoid receiving repeated notifications from Push Mechanism, set up your callback URL to respond according to these HTTP response requirements:

- Includes a status code of 2xx
- Includes an empty body

⚠️ **Note: All Pushes (webhooks) support a different maximum number of notifications and intervals for any repeated notifications. See the next section Push Mechanism Warning/Disable Logic for Apps that have a poor success rate for responding to notifications.**

View the **Push Mechanism documentation** for specific details.

---

# Push Mechanism Warning/Disable Logic

Shopee Open Platform has a warning/disable logic in place for Apps that have a poor success rate for responding to Push Mechanism notifications. For such Apps, we will take action accordingly. Warning emails will be sent to you or we may also eventually disable your subscription to Push Mechanism notifications.

⚠️ **Note: Success rate is calculated by comparing successful Push items versus failed Push items. A failed Push is defined as Shopee Open Platform not receiving an HTTP response with a status code of 2xx and an empty body within the timeout period.**

View the **Push Mechanism documentation** for specific details.

Here are the details for Push Mechanism warning and disabling actions:

## Warning emails

- You will receive a warning email every 30 minutes if:
  - There have been more than 600 Push Mechanism notifications sent to you in the past 6 hours AND
  - Your overall Push success rate is less than 70%.
- Warning emails will not be sent once your success rate returns to more than 70%.

## Disabling of subscription to Push Mechanism notifications

- Your subscription to Push Mechanism notifications will be disabled and you will receive a notification email if:
  - There have been more than 600 Push Mechanism notifications sent to you in the past 6 hours AND
  - Your overall Push success rate is less than 30%.
- You should check your callback URL and ensure that you're ready to receive Push Mechanism notifications normally before subscribing again.

---

## ⚠️ Notes

After subscribing again, here are some important points to note:

- **You will not receive Push Mechanism notifications missed during the period where your subscription was disabled.**
- **Success rate calculation of Push Mechanism notifications will be restarted based on records from your new subscription.**

You can view the current Push Mechanism success rate and status through the Shopee Open Platform Console page.

---

## Code Example

```java
throws NoSuchAlgorithmException, UnsupportedEncodingException, java.security.InvalidKeyException
{
    String baseStr = url + "|" + requestBody;
    
    Mac sha256_HMAC = Mac.getInstance("HmacSHA256");
    
    SecretKeySpec secret_key = new SecretKeySpec(partnerKey.getBytes("UTF-8"), "HmacSHA256");
    
    sha256_HMAC.init(secret_key);
    
    String result = Hex.encodeHexString(sha256_HMAC.doFinal(baseStr.getBytes("UTF-8")));
    
    return result.equals(authorization);
}
```

---

**문서 ID**: developer-guide.18
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/18
**처리 완료**: 2025-10-16T08:07:40
