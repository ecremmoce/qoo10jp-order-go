# API 통합 & 물류 가이드

**카테고리**: 통합
**난이도**: medium
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T09:05:20

## 개요

본 가이드는 Seller Logistics API와의 통합에 대한 포괄적인 정보를 제공합니다. API 구성, 인증, 요청/응답 매개변수, 물류 및 전자 상거래 애플리케이션을 위한 구현 단계를 다룹니다.

## 주요 키워드

- API integration
- REST API
- Logistics
- Seller API
- Authentication
- Request parameters
- Response parameters
- Shipping
- Quotation
- Ecommerce

## 본문

# API 통합 및 물류

## API Rest 콘솔 물류

API Rest 콘솔은 주로 다음을 담당합니다.
- 상점 콘솔에서 등록(판매 생성)할 수 있는 제품(SKU)을 노출합니다.
- 시뮬레이션 및 배송료 경로를 통합하여 각 견적에 정보를 보냅니다(예: "고객 유니폼 SKU").

### URL 구성: REST API

Shpppr에 대해 실행하고 각 배송료 및 구매 완료 시 발생하는 단계를 실행할 수 있도록 하는 것이 좋습니다. 하지만 예시는 각 시설에 대한 문서에 있습니다.

실행을 허용하여 Shpppr에 보고하려면 API에서 코드 op-se 또는 desenter URL을 사용해야 합니다.

## 승인/통합 요구 사항

통합 tajeta로 지정되며, API 통합은 sender 포털의 데이터 업데이트로 구성됩니다.
Shpppr "Seller 물류 API"는 환경 및 상점 운영자를 위한 REST API입니다. 가능한 예시로
Shpppr API 컨트롤러의 deseenho는 테스트로 구성됩니다.

테스트로 API 구성:
- 전송: 판매자 환경에 등록됨(API POST)
- 구성의 유연성을 제공하여 보고된 제품 중 하나를 전달할 수 있도록 허용합니다.
- URL은 특정 구조를 가지며 구성의 유연성을 제공합니다.
- 패널: 제품 제한 없이 Shpppr에서 판매 응답 생성
- 구성 테스트: 판매자 센터에서 환경 등록
- 전송 기능: Shpppr에서 설정한 tófica를 통해 개발(샌드박스 환경)을 위해 프로덕션 환경을 덮어씁니다.
- 등록된 URL을 통해 JSON을 통해 Shpppr에 코드 및 배송료를 전송하면 현재 제품 코드가 표시됩니다.

[API 구성, Criar API Seller Logística, 인증 및 구현 단계에 대한 추가 섹션이 계속됩니다.]

**참고:** 이미지에는 API 통합에 대한 여러 섹션이 있는 포르투갈어 기술 문서 페이지가 포함되어 있습니다. 이미지 품질과 여러 스크린샷 및 주석이 있는 복잡한 레이아웃으로 인해 일부 텍스트가 부분적으로 가려지거나 불분명할 수 있습니다. 콘텐츠에는 물류/판매자 API 시스템과 통합하기 위한 설정 지침, 요구 사항 및 단계별 절차가 포함되어 있습니다.

---

# API 통합 가이드

## 6. 404 오류/페이지 구성: 전자 상거래 호출, 추가됨, Cabo info msg podđeni francesco uni "catálogo" e "password", vv post dados로 검색할 수 있습니다.

### Shop App 구성

**페이지 통합 단계:**

1. "App ID"의 경우 해당 API에만 액세스할 필요가 없는 경우 애플리케이션의 ID proudly를 가져옵니다. IT 액세스 비활성화를 클릭하고 활성화를 위한 정보를 입력합니다.

---

### 페이로드 패턴 구성

**구성 단계:**

- 다음에서 Platform em Touch API에 액세스합니다. com o escalar "Seller Logístico" diretamente o blog, API req posto. Deve isso premerer, precaregar ou tag conforme a lista de 5Pages finally a cookie de URL. 각 API 또는 Shopper에 연결할 수 있는 URL이 있습니다.

**중요 참고 사항:**

- 통합이 "Seller" 및 "partner_id" Lise nova técn-le이 아닌 경우 이메일을 통해 품질을 확인할 수 없습니다. Assim fim o integração.

- "Post ID" Shopper responsáber alto bales를 반환하여 구성이 완료되었음을 확인하고 인증에 대한 수정 사항을 수행하면 LIVE에서 낮은 프로세스를 realiar할 수 있습니다.

---

## 견적 요청 및 응답의 표준 페이로드

### 요청 매개변수(쿼리)

#### API 세분화

| 필드 | 예시 | HTTP 주소 |
|-------|---------|--------------|
| URL | https://api.fixxel | 판매자의 ERP에서 형성된 URL |

#### 매개변수 테이블

| 이름 | 유형 | 예시 | 설명 |
|------|------|---------|-----------|
| platform_id | int | 1 | 파트너가 통합된 후 배포되는 애플리케이션 ID입니다. 시스템에서 애플리케이션을 식별합니다. |
| timestamp | timestamp | 1670000000 | 요청의 타임스탬프를 나타냅니다. 인증 공격을 방지하기 위해 필수입니다. 5분 후에 만료됩니다. |
| token | string | a519&5e2017f8f4a9f55622701f1017be77a6f2a5a56e5e6550b95516fde3 | 연결("platform_id" + api_key hash_hmac) 타임스탬프에 의해 생성된 서명입니다. 생성된 토큰의 SHA256을 수행해야 합니다. 참조: https://www.google.com/amp/s/acervolima.com/funcao-hash-hmac-no-php/ |

---

### Body Request 매개변수

| 이름 | 유형 | 필수 | 예시 | 설명 |
|------|------|-------------|---------|-----------|
| drop_id | int | True | 12234525 | 각 요청의 고유 식별자 |
| origin_zip_code | string | True | 12345000 | 8자리 숫자로만 구성된 발신자 우편 번호(대시 및 하이픈 없음) |
| destination_zip_code | string | True | 12346000 | 8자리 숫자로만 구성된 구매자 우편 번호(대시 및 하이픈 없음) |
| items | array | True | - | 제품 목록 |
| items_id | int | True | 12334678 | Shoppe의 항목 식별 |
| sku | string | False | 8695-abc | Shoppe의 상점 운영자가 등록한 SKU |
| price_id | int | False | 12334678 | Shoppe의 가격 식별. 참고: POST에서는 필수가 아니며, model에서 정보를 제공하는 경우 가격을 보내야 하지만 응답에서 반환해서는 안 됩니다. |
| category_id | int | True | 12344678 | 항목 카테고리 식별. 참고: POST에서는 필수가 아니며, 정보를 제공하는 경우 가격을 보내야 하지만 반드시 응답에서 반환해야 합니다. |
| quantity | int | True | 12344678 | 항목 수량 |
| price | float | False | - | - |
| dimensions | object | True | - | 제품 크기 |
| length | int | True | 10 | 센티미터 단위의 길이 |
| width | int | True | 10 | 센티미터 단위의 너비 |
| height | int | True | 10 | 센티미터 단위의 높이 |
| weight | int | True | 100 | 그램 단위의 무게 |

---

### 응답 매개변수 / 견적 응답(항목)

| 이름 | 유형 | 필수 | 예시 | 설명 |
|------|------|-------------|---------|-----------|
| id | int | True | - | 판매자가 수행한 견적 식별자 |
| message | string | True | - | 구매자 우편 번호 |
| request_id | string | True | - | API 호출의 bd 식별자 |

---

### 응답 매개변수 / 견적 응답(성공)

| 이름 | 유형 | 필수 | 예시 | 설명 |
|------|------|-------------|---------|-----------|
| quotation_id | int | True | 0912234000 | 판매자가 완료한 최종 견적 식별자 |

---

*문서 끝*

---

# API 문서 추출

## 요청 매개변수

| 이름 | 유형 | 필수 | 예시 | 설명 |
|------|------|------------|---------|-------------|
| message | string | True | | 제품 정보가 포함된 알림 메시지 |
| request_id | string | True | | API 사용자 식별 |

## 응답 매개변수 / Consulta 응답(응답)

| 이름 | 유형 | 필수 | 예시 | 설명 |
|------|------|------------|---------|-------------|
| request_id | string | True | 0912345678 | API 사용자 식별 코드 |
| session_key_code | string | True | | 제품 알림 코드 폴더 |
| packages | array | True | | 패키지 목록 |
| dimensions | object | True | | 제품 크기 |
| length | int | True | 10 | 센티미터 단위의 길이 |
| width | int | True | 10 | 센티미터 단위의 너비 |
| height | int | True | 10 | 센티미터 단위의 높이 |
| weight | int | True | 100 | 그램 단위의 무게 |
| items | array | True | | 제품 항목 목록 |
| item_id | int | True | 12345678 | Shoppe 항목 식별 |
| model_id | int | False | 12345678 | Shoppe 모델 식별 |
| model_id | int | False | 12345678 | Shoppe의 분류 모델 식별 |
| category_id | int | False | 12345678 | Shoppe의 제품 항목 카테고리 |
| quantity_id | int | True | 2 | 항목 수량 |
| price | int | True | 100.4 | 제품 가격 |
| dimensions | object | True | | 제품 크기 |
| length | int | True | 10 | 센티미터 단위의 길이 |
| width | int | True | 10 | 센티미터 단위의 너비 |
| height | int | True | 10 | 센티미터 단위의 높이 |
| weight | int | True | 100 | 그램 단위의 무게 |
| quotations | array | True | | 견적 목록 |
| price | float | True | 100.4 | 무료 배송 비교 값 |
| handling_time | int | True | 20 | 주문 준비/처리 시간(주문 준비 또는 배송/운송 시간 + 분리 시간) |
| shipping_time | int | True | 10 | 배송/운송 시간 데이터 |
| promise_time | int | True | 30 | 약속된 준비 시간 + 운송 시간 |
| service_code | string | True | M1020 | 공급자 서비스를 식별하는 코드(M1020의 "shipping_carrier" vs "shipping_method" 매개변수 코드) |

## 요청 예시(Ejemplo de requisição)

```bash
curl -i -H "GET" "https://api.freights/" \
-H "Content-Type: application/json" \
-H "x-session-user-type: API" \
-H "x-session-key: YOUR_SESSION_KEY" \
-d '{
    "req_id": "REQ123456",
    "category_id": "CATEGORY",
    "package_id": 0,
    "dimensions": {
        "width": 10,
        "length": 30,
        "height": 5,
        "weight": 100
    }
}'
```

## 성공 응답 예시(Exemplo de retorno (sucesso))

**HTTP 상태 코드:** 200

```json
{
    "return_id": "123456789",
    "session_key_code": "FREESPACE",
    "packages": [
        {
            "category_id": "CATEGORY",
            "items": [
                {
                    "width": 10,
                    "length": 30,
                    "height": 5,
                    "weight": 100,
                    "price": [],
                    "item_id": [],
                    "model_id": [],
                    "quantity": []
                }
            ],
            "quotations": [
                {
                    "price": 35.50,
                    "handling_time": 2,
                    "shipping_time": 10,
                    "promise_time": 12,
                    "service_code": "M1020"
                }
            ]
        }
    ]
}
```

## 오류 응답 예시(Exemplo de retorno (erro))

**HTTP 상태 코드:** 400

```json
{
    "request_id": "12345678910.012345",
    "message": "Invalid parameters",
    "code": 400
}
```

## 지연 오류에 대한 오류 응답(Resposta padrão para erros após requisição / Modelo de erros)

**HTTP 상태 코드:** 500

```json
{
    "request_id": "NOVALIDREQUEST00REQUEST",
    "message": "INTERNAL ERROR",
    "code": 500
}
```

---

# API 문서

## 오류 테이블에 매핑되지 않은 오류에 대한 표준 보고서

- HTTP 상태 코드 500 사용

```json
{
  "request_id": "KeyGen42xDN5xAs3cEu6AzY",
  "error": "internal server error",
  "message": "internal server error"
}
```

---

## 예상되는 반환 코드

| HTTP 상태 코드 | Error | Message | Area |
|-----------------|-------|---------|------|
| 200 | success API call, none Empty | success API call, none Empty | False |
| 403 | error_partner_id | partner_id, id 쿼리를 참조합니다. | False |
| 403 | error_partner_id | partner_id가 유효하지 않습니다. | False |
| 403 | error_login | 로그인이 유효하지 않습니다. | False |
| 403 | error_login | 오류 로그인이 유효하지 않습니다. | False |
| 403 | error_timestamp | 쿼리에 타임스탬프가 없습니다. | False |
| 403 | error_timestamp | 오류 타임스탬프가 유효하지 않습니다. | False |
| 403 | error_timestamp | 게시 타임스탬프가 유효하지 않습니다. | False |
| 403 | error_debug_id | debug_id가 유효하지 않습니다. | False |
| 403 | invalid origin_zip_code | origin_zip_code가 유효하지 않습니다. | False |
| 403 | invalid destination_zip_code | destination_zip_code가 유효하지 않습니다. | False |
| 403 | invalid destination_city_code | destination_city가 유효하지 않습니다. | False |
| 403 | invalid item_id | item_id가 유효하지 않습니다. | False |
| 403 | invalid model_id | model_id가 유효하지 않습니다. | False |
| 403 | invalid title | 제목이 유효하지 않습니다. | False |
| 403 | error_category_id | category_id가 유효하지 않습니다. | False |
| 403 | invalid quantity | 수량이 유효하지 않습니다. | False |
| 403 | invalid price | 가격이 유효하지 않습니다. | False |
| 403 | error_dimensions | 크기가 유효하지 않습니다. | False |
| 403 | invalid_weight | 무게가 유효하지 않습니다. | False |
| 403 | error_width | 너비가 유효하지 않습니다. | False |
| 403 | error_height | 높이가 유효하지 않습니다. | False |
| 500 | Internal system error | 내부 시스템 오류 | - |

---

## 응답 유효성 검사(Shopee URL 포함)

모든 빈 필드를 반환하는 "get_shipping_parameter" 호출의 응답에서 수행된 유효성 검사이며, 하나의 정보 변수만 사용합니다.

- 이 API의 통합 테스트에서는 필드를 완전하게 유효성을 검사/나열할 수 있는 캐리어의 경로/URL 중 일부를 사용하는 것이 좋습니다.
- "x-quotation" 매개변수 내에서 모든 매개변수가 포맷된 URL을 보내는 데 사용해야 합니다.
- 응답 형식 유효성 검사에서 URL 예시의 패턴을 따릅니다.

```
curl --request GET \
  --url 'https://partner-api-ur.test.api_mall.id/api/v1/shipping/CARRIER-TEST-1-534?' \
  --header 'Content-Type: application/json' \
  --header 'X-Partner-ID: PARTNER_ID' \
  --header 'debug-id: DEBUG_ID' \
  --header 'origin-zip-code: 01310-000' \
  --header 'partner-id: 2' \
  --header 'timestamp: 2024-05-07T14:30:00.000Z' \
  --header 'x-quotation: {"partner_id": 2, "timestamp": "2024-05-07T14:30:00.000Z", "origin": {"zip_code": "01310-000"}, "destination": {"zip_code": "04563-040", "city": "São Paulo"}, "items": [{"id": "MLB123456", "quantity": 2, "unit_price": 150.00, "dimensions": {"width": 10, "height": 5, "length": 15, "weight": 2}}]}' \
  --data ''
```

---

## 견적 배열이 있는 Jasper 응답("validation_errors" 매개변수 사용)

각 캐리어의 가격을 평가하기 전에 Jasper에서 반환된 배열의 표준 구조와 유사한 응답 배열을 만듭니다.

```json
{
  "api call id customer": [],
  "debug_id": "",
  "version": "Jasper 견적 배열",
  "to": {
    "zip_code": "",
    "address": {
      "sublocality": {"name": ""},
      "locality": {"name": ""},
      "administrative_area_level_2": {"name": ""},
      "administrative_area_level_1": {
        "short_name": "",
        "name": ""
      },
      "country": {
        "short_name": "",
        "name": ""
      }
    }
  },
  "from": {
    "zip_code": ""
  },
  "items": [{
    "id": "INTERNAL",
    "title": "MIRROR",
    "variation_id": 0,
    "category_id": "",
    "model_id": "",
    "seller": {
      "id": 0
    },
    "dimensions": {
      "category_id": "",
      "model_id": "",
      "length": 0,
      "width": 0,
      "height": 0,
      "weight": 0
    }
  }],
  "options": [{
    "id": 0,
    "name": "INTERNAL",
    "display": "MIRROR",
    "currency_id": "BRL",
    "list_cost": 0,
    "cost": 0,
    "type": "",
    "carrier_id": "",
    "shipping_method_id": 0,
    "estimated_delivery": {
      "date": "0",
      "type": "",
      "offset": {
        "date": "0"
      },
      "time_from": 0,
      "time_to": 0
    },
    "estimated_handling_limit": {
      "date": ""
    }
  }],
  "validation_errors": {
    "zip_code": "",
    "message": "Error code details"
  }
}
```

---

# API 통합 가이드

## 오류 매핑으로 응답 트리거

```json
{
  "response": {
    "trigger": {
      "height": 1,
      "weight": 1,
      "width":
## Jasper 응답 (견적 배열 포함, "validation_errors" 파라미터 사용)

각 운송업체의 가격을 평가하기 전에 Jasper가 반환하는 표준 배열 구조와 유사한 응답 배열을 생성합니다.

```json
{
  "api call id customer": [],
  "debug_id": "",
  "version": "Jasper 견적 배열",
  "to": {
    "zip_code": "",
    "address": {
      "sublocality": {"name": ""},
      "locality": {"name": ""},
      "administrative_area_level_2": {"name": ""},
      "administrative_area_level_1": {
        "short_name": "",
        "name": ""
      },
      "country": {
        "short_name": "",
        "name": ""
      }
    }
  },
  "from": {
    "zip_code": ""
  },
  "items": [{
    "id": "INTERNAL",
    "title": "MIRROR",
    "variation_id": 0,
    "category_id": "",
    "model_id": "",
    "seller": {
      "id": 0
    },
    "dimensions": {
      "category_id": "",
      "model_id": "",
      "length": 0,
      "width": 0,
      "height": 0,
      "weight": 0
    }
  }],
  "options": [{
    "id": 0,
    "name": "INTERNAL",
    "display": "MIRROR",
    "currency_id": "BRL",
    "list_cost": 0,
    "cost": 0,
    "type": "",
    "carrier_id": "",
    "shipping_method_id": 0,
    "estimated_delivery": {
      "date": "0",
      "type": "",
      "offset": {
        "date": "0"
      },
      "time_from": 0,
      "time_to": 0
    },
    "estimated_handling_limit": {
      "date": ""
    }
  }],
  "validation_errors": {
    "zip_code": "",
    "message": "Error code details"
  }
}
```

---

# API 통합 가이드

## 오류 매핑을 사용한 트리거 응답

```json
{
  "response": {
    "trigger": {
      "height": 1,
      "weight": 1,
      "width": 15
    }
  },
  "duration": 1,
  "grid": 56.5,
  "pending_time": 0,
  "sorting_time": 30,
  "service_time": 30,
  "service_order": "UNDEFINED"
}

"method_url": "POST"
```

> 전체 선택 오류 예제가 포함된 트리거 응답.

```json
{
  "response": {
    "error": "There should be a valid string 'Error'."
  },
  "url": "https://api.onfleet.com/v2/tasks/:taskId/clone",
  "method": "POST",
  "headers": {
    "Content-Type": "application/json",
    "Authorization": "Bearer dGVzdGluZ0Bjb3VsZGJlYmV0dGVyQWVpdGhlcg=="
  },
  "webhook_url": "AMEND",
  "trigger": "CREATED"
}

"with_body": {
  "url_id": "CHECKED",
  "trigger": "CREATED",
  "webhook": "UNDEFINED",
  "type": "UPDATED",
  "time_id": "UPDATED",
  "items": [
    {
      "time_id": "UPDATED",
      "watch": 5,
      "dev": "UPDATED",
      "obj": {
        "dimension": "CREATED"
      },
      "field": {
        "width": 1,
        "dev": 1,
        "time": 1,
        "items": 1
      },
      "status": 27.67,
      "dimension": {
        "CHECKED": 1,
        "width": 1,
        "dev": 1,
        "items": 1
      }
    }
  ]
}

"response_time_backend_url": {
  "service_code": 400,
  "created": true,
  "message_id": "**1*3*GO:233000",
  "service_id": 9,
  "error": "This is to intended_id if 'ERROR'",
  "trigger": "NO: The webhook can still retrieve."
}

"method_url": "POST"
```

---

## 응답 시간 제한 (견적)

- 개별 견적은 구매 및 결제가 이루어질 때 사용자 경험 내에서 고려됩니다.
- 예상되는 배송 견적 응답 시간은 3초입니다.
- 기한 내에 처리할 수 없는 경우 파트너는 오류를 반환하고 고객이 다시 시도하도록 허용할 수 있습니다.
- 3초를 초과하는 응답 시간은 허용되지 않으며 통합이 불가능합니다. 따라서 경험을 충족할 수 있도록 쿼리를 평가하고 최적화하십시오.

---

## 신뢰성 테이블

- 이 레이블은 배송비 계산 결과와 같으며, 배송비 견적 또는 CDI(취소) 수준 구조화에 문제가 있는 경우 발생합니다.

---

## 신뢰성을 위한 배송비 값 구성

- 자세한 기능: 판매자가 배송비를 요청하거나 판매자의 개별 채널을 수행하는 민주주의는 아래 설명된 단계를 따릅니다.
- 값을 혼동하지 마십시오. 액세스: 배송 설정 및 판매.

### 배송 설정

[토글 스위치가 있는 구성 옵션을 보여주는 인터페이스]

---

## 신뢰성 배송비 값 채우기

- 이 질문: 배송비 값을 채우기 위한 목록으로 구성됩니다. 배송비 값은 주별로 채워지거나 집중된 우편 번호를 삽입할 수 있습니다.

---

## 주별 단일 배송비 값

### 단계:
1. 주 선택
2. 단일 배송비 값 입력
3. 저장을 클릭합니다.

[주, 값 및 옵션 열이 있는 주 기반 배송 구성 테이블]

---

## 도시별 단일 배송비 값

### 단계:
1. 주 선택
2. 각 주의 도시 선택
3. 모든 도시 선택 옵션을 채우고 배송비 값을 알리고 싶은 도시를 선택합니다.
4. 단일 배송비 값 입력
5. 저장을 클릭합니다.

[주, 도시, 값 및 옵션 열이 있는 도시 기반 배송 구성 테이블]

---

## 판매자 물류 채널 활성화

[물류 채널 구성 설정을 보여주는 인터페이스]

---

**참고:** 모든 인터페이스 요소, 테이블 및 구성 화면은 원본 문서에 표시된 대로 구조적 컨텍스트에 보존됩니다.

---

# 판매자 물류 채널 활성화

## 배송 설정

### 배송 채널
Shop 통합 배송 채널 활성화

### 통합 배송 채널
이러한 유형의 채널을 활성화하면 Shopee의 물류 서비스를 사용하여 기계 액세서리를 배송하고 물류 허브에서 직접 패키지를 허용하는 데 동의합니다.

**판매자 물류** [채널 설정 토글: 켜짐]

### 감열식 프린터
Bluetooth를 통해 연결된 감열식 프린터를 사용하여 배송 가이드를 쉽게 인쇄하거나 라벨을 만들 수 있습니다. 현재 프린터는 감열식 프린터가 Android 장치에서 사용되었는지 확인할 수 있습니다. USB 가이드 인쇄 구성은 판매자가 사용할 수 있어야 합니다.

---

## 비상 배송비 값에 대한 배송 시간

비상 테이블이 트리거되면 배송 시간은 고정됩니다.
- 최소 기간: 10일
- 최대 기간: 20 영업일

---

## 주문 상태 및 추적

### 상태 업데이트 API 개발:

주문 상태 및 추적 업데이트는 추적 업데이트 API를 통해 이루어지며 Ship Confirm API(logistics_init_info)에서 반환된 추적 번호를 기반으로 수행됩니다. 여기서 전송되는 추적 번호를 알려야 합니다.

주문 추적 정보 업데이트에는 엔드포인트 v2 logistics_update_tracking_status(OpenAPI)를 사용해야 합니다. 사용 가능한 상태는 다음과 같습니다.

1. 주문 발송됨(logistics_pickup_done)
   - 발송됨 상태는 주문 상태 업데이트를 발송됨으로 완료할 때만 사용해야 합니다.

2. 주문 배송됨(logistics_delivery_done)
   - 배송됨 상태는 주문에 이미 발송됨 상태가 있는 경우에만 수신됩니다.

3. 배송 실패(logistics_delivery_failed)
   - 배송 실패 상태는 주문에 이미 발송됨 상태가 있는 경우에만 수신됩니다.

**중요:** 주문 배송됨 또는 배송 실패 상태를 보낸 후에는 상태 업데이트가 더 이상 허용되지 않아 최종 상태가 됩니다.

tracking_number 및 tracking_url 파라미터는 logistics_pickup_done에 대한 상태 업데이트에서만 보내야 합니다.

API 문서 링크 "https://open.shopee.com/documents/v2/v2.logistics.update_tracking_status"를 따르십시오.

---

## 채널 주문 흐름 및 상태

### 상태 흐름:

**1. 결제가 Open Seller API에 의해 확인되는 즉시 상태가 "UNPAID_TO_SHIP"으로 업데이트됩니다.**

↓

**2. 판매자는 NF를 생성하고 API "nrg_invoice_doc"를 통해 보내야 합니다.**

↓

**3. Irma는 발행된 NF를 확인하고 상태는 "READY_TO_SHIP" 배송(API "ship_order")으로 업데이트되고 상태는 "SHIPPED"로 업데이트됩니다.**

→ **시작**
→ **UNPAID**(1 및 2)
→ **READY_TO_SHIP**
→ **SHIPPED**(3)
→ **TO_CONFIRM_RECEIVE**(6)
→ **COMPLETED**(7)

**4. 주문이 물류 파트너에 의해 수집되는 즉시 판매자는 Ship Confirm API를 호출하고 상태 "logistic_pickup_done"을 업데이트하고 tracking_url을 사용할 수 없는 경우 tracking_number를 보낼 수도 있으며 SHIPPED 상태로 업데이트됩니다(상태가 업데이트됨).**

**5. 배송에 문제가 있는 경우 판매자는 주문 상태를 "logistic_delivery_failed"로 업데이트하여 주문을 취소하고("CANCELLED" 상태) Shopee API에 수동으로 반환할 수 있습니다.**

**6. 주문이 성공적으로 배송되면 판매자는 API "update_tracking_status"를 호출하고 상태를 "logistic_delivery_done"으로 업데이트하고 주문은 "TO_CONFIRM_RECEIVE" 상태가 됩니다.**

**7. 구매자가 수령을 확인하거나 7일 후 주문이 자동으로 확인되고 상태가 COMPLETED 상태로 업데이트됩니다.**

---

## 채널에 대한 수수료 및 배송 규칙

- 최대 수수료 할인은 6%이지만 판매자는 무료 배송 프로그램에 참여해야 합니다.
- 판매자는 고객에게 제공되는 할인에 공동 참여합니다.

| Shopee에서 지불한 금액 | 판매자가 지불한 금액 | Shopee에서 지불한 금액 제한 |
|------------------------|-------------------------|----------------------------------|
| 37% | 63% | R$ 7.30 |

---

## FAQ:

### 1. 모든 판매자가 판매자 물류 채널에 액세스할 수 있습니까?
아니요, 무료 배송 프로그램에 참여하고 구성이 있는 판매자만 계정 관리자를 찾으십시오.

### 2. 주문 견적을 식별하는 방법은 무엇입니까?
성공적인 견적으로 생성된 주문을 할 때 "service_code" 파라미터는 "shipping_carrier_order_detail" 파라미터에서 반환됩니다.

예:
- "service_code" 전송됨: "채널 X"
- "shipping_carrier" 파라미터 반환: "판매자 물류 - 채널 X"

예 2:
- 비상 테이블 활성화됨
- "shipping_carrier" 파라미터 반환: "판매자 물류"

### 3. APP Seller Logistics를 사용하여 OpenAPI를 호출할 수 있습니까?
아니요, APP 사용은 견적 API에만 사용됩니다.

더 많은 기술적 질문이 있는 경우 OP 티켓 플랫폼에서 티켓을 여십시오.
물류 채널 액세스에 대한 일반적인 질문은 계정 관리자에게 문의하십시오.

## 사용 사례

1. 판매자의 ERP 시스템을 물류 제공업체의 API와 통합합니다.
2. 전자 상거래 플랫폼에서 배송 옵션 및 요금을 구성합니다.
3. 배송 견적 생성 프로세스를 자동화합니다.
4. 배송 계산을 위한 제품 데이터 및 크기를 관리합니다.
5. 안전한 데이터 교환을 위해 API 요청을 인증합니다.

## 관련 API

- Seller Logística API
- REST API

---

## 원문 (English)

### Summary

This guide provides comprehensive information on integrating with the Seller Logistics API. It covers API configuration, authentication, request/response parameters, and implementation steps for logistics and e-commerce applications.

### Content

# API Integration & Logistics

## API Rest Console Logística

O Console da API Rest é responsável, principalmente, por:
- Expor os produtos (SKU) disponíveis para registro (criação de vendas) no Console da Loja
- Integrar a rota de simulação e freete, ao enviar suas informações para cada cotimização (ex: "skus uniforme do Cliente")

### URL Config: REST API

É interessante que permita de execução rodar para o Shpppr e cada de freete e passo relatado que tenha oculados ao completar da envenda da compra. Mas seu exemplo esta nos documentação por estabelecimento.

To executar usar permita de execução relatarmente enviar a Shpppr você precisará usar URL do código op-se um desenter na API.

## Requisitos para homologação/integração

Denote-se uma tajeta ser de integração, a integração da API têm consiste o atualização de dados no portal sender.
Shpppr o "Seller Logística API" e uma REST API no ambiente e lojista. Port seu possível para um exemplo e
Shpppr o deseenho do Controller da API no consiste de teste

Configurar o API no consiste de teste:
- Envio: cadastrada ao ambiente de seller (API POST)
- Permitir passar uma de um produto relatado, proporcionando flexibilidade na configuração
- A URL, têm uma estrutura específica, proporcionando flexibilidade na configuração
- Painel de: geração de vendação de response no Shpppr sem limitar de produto
- Teste de configuração: cadastrada ambiente pela Central de Vendedor
- Funcionalidade de envio: p/ desenvolvimento (ambiente de sandbox) sobrecolhendo (ambiente de produção) estabelecidos pela Shpppr da tófica
- Envio de códigos e taxas de freete por meio de um JSON ao Shpppr através da URL cadastrada, demonstrarão código de produto real do momento

[Additional sections continue with technical details about API configuration, Criar API Seller Logística, authentication, and implementation steps]

**Note:** The image contains a technical documentation page in Portuguese with multiple sections about API integration. Due to the image quality and complex layout with multiple screenshots and annotations, some text may be partially obscured or unclear. The content includes setup instructions, requirements, and step-by-step procedures for integrating with a logistics/seller API system.

---

# API Integration Guide

## 6. 404 errors/page configuration: ecommerce calls, added, Cabo info msg podđeni francesco uni "catálogo" e "password", podem pesquisar como vv post dados:

### Shop App Configuration

**Steps for Page Integration:**

1. Para "App ID" obtenha o ID proudly da sua aplicação caso estejador que não carece terne o acesso apenas a essa API à parte. Clique IT Acesso Desativado, preencheri informações para ativação.

---

### Payload Pattern Configuration

**Configuration Steps:**

- Access the Platform em Touch API at: com o escalar "Seller Logístico" diretamente o blog, API req posto. Deve isso premerer, precaregar ou tag conforme a lista de 5Pages finalmente a cookie de URL. Cada classe onde API ou pode estar conectado a uma URL que va Shopper.

**Important Notes:**

- Apenas com integração não "Seller" e "partner_id" Lise nova técn-le você não's qualidade via e-mail. Assim fim o integração.

- Restitua o "Post ID" Shopper responsáber alto bales, confirmando que a configuração foi feira e seu bater realizar a fixa na autenticação, ja será possível realiar os processos baixos are LIVE.

---

## Payload padrão da requisição e resposta de cotação

### Parameters da requisição (query)

#### API Segmentation

| Field | Example | HTTP Address |
|-------|---------|--------------|
| URL | https://api.fixxel | URL formado pelo ERP do seller |

#### Parameters Table

| Nome | Tipo | Exemplo | Descrição |
|------|------|---------|-----------|
| platform_id | int | 1 | O ID do aplicativo é distribuído após o parceiro ser integrado. Identifica aplicação no sistema. |
| timestamp | timestamp | 1670000000 | Indica o carimbo de data/hora da requisição. Obrigatório para evitar ataques de autenticação. Expira em 5 minutos. |
| token | string | a519&5e2017f8f4a9f55622701f1017be77a6f2a5a56e5e6550b95516fde3 | Assinatura gerada pela concatenação ("platform_id" + api_key hash_hmac) timestamp. É preciso fazer SHA256 do token gerado. Ver: https://www.google.com/amp/s/acervolima.com/funcao-hash-hmac-no-php/ |

---

### Body Request Parameters

| Nome | Tipo | Obrigatório | Exemplo | Descrição |
|------|------|-------------|---------|-----------|
| drop_id | int | True | 12234525 | Identificador único de cada solicitação |
| origin_zip_code | string | True | 12345000 | CEP do remetente com 8 dígitos, apenas números sem traços e hífens |
| destination_zip_code | string | True | 12346000 | CEP do comprador com 8 dígitos, apenas números sem traços e hífens |
| items | array | True | - | Lista de produtos |
| items_id | int | True | 12334678 | Identificação do item na Shoppe |
| sku | string | False | 8695-abc | SKU cadastrado pelo lojista do Shoppe |
| price_id | int | False | 12334678 | Identificação do preço no Shoppe. Obs: não é obrigatório no POST e, model, se informado é necessário enviar price, mas não deve ser retornado na resposta |
| category_id | int | True | 12344678 | Identificação da categoria do item. Obs: não é obrigatório no POST e, se informado é necessário enviar price, mas deve ser retornado necessariamente |
| quantity | int | True | 12344678 | Quantidade de items |
| price | float | False | - | - |
| dimensions | object | True | - | Dimensões do produto |
| length | int | True | 10 | Comprimento em centímetros |
| width | int | True | 10 | Largura em centímetros |
| height | int | True | 10 | Altura em centímetros |
| weight | int | True | 100 | Peso em gramas |

---

### Response Parameters / Resposta da Cotação (Item)

| Nome | Tipo | Obrigatório | Exemplo | Descrição |
|------|------|-------------|---------|-----------|
| id | int | True | - | Identificador da cotação realizada pelo seller |
| message | string | True | - | CEP do comprador |
| request_id | string | True | - | Identificador bd chamadas do API |

---

### Response Parameters / Resposta da Cotação (SUCCESS)

| Nome | Tipo | Obrigatório | Exemplo | Descrição |
|------|------|-------------|---------|-----------|
| quotation_id | int | True | 0912234000 | Identificador da cotação finalizada pelo seller |

---

*End of Document*

---

# API Documentation Extract

## Request Parameters

| Name | Type | Obligatory | Example | Description |
|------|------|------------|---------|-------------|
| message | string | True | | Notification message with product info |
| request_id | string | True | | API user identification |

## Response Parameters / Resposta da Consulta (response)

| Name | Type | Obligatory | Example | Description |
|------|------|------------|---------|-------------|
| request_id | string | True | 0912345678 | API user identification code |
| session_key_code | string | True | | Product notification code folder |
| packages | array | True | | Package list |
| dimensions | object | True | | Product dimensions |
| length | int | True | 10 | Length in centimeters |
| width | int | True | 10 | Width in centimeters |
| height | int | True | 10 | Height in centimeters |
| weight | int | True | 100 | Weight in grams |
| items | array | True | | Product item list |
| item_id | int | True | 12345678 | Shoppe item identification |
| model_id | int | False | 12345678 | Shoppe model identification |
| model_id | int | False | 12345678 | Classification model identification in Shoppe |
| category_id | int | False | 12345678 | Product item category in Shoppe |
| quantity_id | int | True | 2 | Item quantity |
| price | int | True | 100.4 | Product price |
| dimensions | object | True | | Product dimensions |
| length | int | True | 10 | Length in centimeters |
| width | int | True | 10 | Width in centimeters |
| height | int | True | 10 | Height in centimeters |
| weight | int | True | 100 | Weight in grams |
| quotations | array | True | | Quotation list |
| price | float | True | 100.4 | Free shipping comparison value |
| handling_time | int | True | 20 | Preparation/handling time for order (preparation time + separation time for order preparation or shipping/transport time) |
| shipping_time | int | True | 10 | Shipping/transport time data |
| promise_time | int | True | 30 | Promised preparation time + transport time |
| service_code | string | True | M1020 | Code identifying provider service (code parameter "shipping_carrier" vs "shipping_method" parameter in M1020) |

## Request Example (Ejemplo de requisição)

```bash
curl -i -H "GET" "https://api.freights/" \
-H "Content-Type: application/json" \
-H "x-session-user-type: API" \
-H "x-session-key: YOUR_SESSION_KEY" \
-d '{
    "req_id": "REQ123456",
    "category_id": "CATEGORY",
    "package_id": 0,
    "dimensions": {
        "width": 10,
        "length": 30,
        "height": 5,
        "weight": 100
    }
}'
```

## Success Response Example (Exemplo de retorno (sucesso))

**HTTP Status Code:** 200

```json
{
    "return_id": "123456789",
    "session_key_code": "FREESPACE",
    "packages": [
        {
            "category_id": "CATEGORY",
            "items": [
                {
                    "width": 10,
                    "length": 30,
                    "height": 5,
                    "weight": 100,
                    "price": [],
                    "item_id": [],
                    "model_id": [],
                    "quantity": []
                }
            ],
            "quotations": [
                {
                    "price": 35.50,
                    "handling_time": 2,
                    "shipping_time": 10,
                    "promise_time": 12,
                    "service_code": "M1020"
                }
            ]
        }
    ]
}
```

## Error Response Example (Exemplo de retorno (erro))

**HTTP Status Code:** 400

```json
{
    "request_id": "12345678910.012345",
    "message": "Invalid parameters",
    "code": 400
}
```

## Error Response for Late Error (Resposta padrão para erros após requisição / Modelo de erros)

**HTTP Status Code:** 500

```json
{
    "request_id": "NOVALIDREQUEST00REQUEST",
    "message": "INTERNAL ERROR",
    "code": 500
}
```

---

# API Documentation

## Relatório padrão para erros não mapeados na tabela de erros

- Usar HTTP status code 500

```json
{
  "request_id": "KeyGen42xDN5xAs3cEu6AzY",
  "error": "internal server error",
  "message": "internal server error"
}
```

---

## Códigos de retorno esperados

| HTTP status code | Error | Message | Area |
|-----------------|-------|---------|------|
| 200 | success API call, none Empty | success API call, none Empty | False |
| 403 | error_partner_id | Refers to the partner_id, id's query | False |
| 403 | error_partner_id | partner_id is invalid | False |
| 403 | error_login | The login is invalid | False |
| 403 | error_login | error login is invalid | False |
| 403 | error_timestamp | there is no timestamp in query | False |
| 403 | error_timestamp | error timestamp is invalid | False |
| 403 | error_timestamp | post timestamp is invalid | False |
| 403 | error_debug_id | The debug_id is invalid | False |
| 403 | invalid origin_zip_code | The origin_zip_code is invalid | False |
| 403 | invalid destination_zip_code | The destination_zip_code is invalid | False |
| 403 | invalid destination_city_code | The destination_city is invalid | False |
| 403 | invalid item_id | The item_id is invalid | False |
| 403 | invalid model_id | The model_id is invalid | False |
| 403 | invalid title | The title is not valid | False |
| 403 | error_category_id | The category_id is invalid | False |
| 403 | invalid quantity | The quantity is invalid | False |
| 403 | invalid price | The price is invalid | False |
| 403 | error_dimensions | The dimensions is invalid | False |
| 403 | invalid_weight | The weight is invalid | False |
| 403 | error_width | The width is invalid | False |
| 403 | error_height | The height is invalid | False |
| 500 | Internal system error | Internal system error | - |

---

## Validação nos Responses (com URL da Shopee)

Validação feita no response da chamada de "get_shipping_parameter" que retorna todos os campos vazios, sendo que eles só usam uma variável informada.

- Nos testes de integração dessa API, seria viável usar alguma das rotas/urls do carrier na qual ele possa validar/listar os campos de forma completa;
- Dentro do parâmetro "x-quotation" deve ser sim usado para enviar a URL com todos os parâmetros formatados.
- Seguir padrão do exemplo de URL na validação do formato de response

```
curl --request GET \
  --url 'https://partner-api-ur.test.api_mall.id/api/v1/shipping/CARRIER-TEST-1-534?' \
  --header 'Content-Type: application/json' \
  --header 'X-Partner-ID: PARTNER_ID' \
  --header 'debug-id: DEBUG_ID' \
  --header 'origin-zip-code: 01310-000' \
  --header 'partner-id: 2' \
  --header 'timestamp: 2024-05-07T14:30:00.000Z' \
  --header 'x-quotation: {"partner_id": 2, "timestamp": "2024-05-07T14:30:00.000Z", "origin": {"zip_code": "01310-000"}, "destination": {"zip_code": "04563-040", "city": "São Paulo"}, "items": [{"id": "MLB123456", "quantity": 2, "unit_price": 150.00, "dimensions": {"width": 10, "height": 5, "length": 15, "weight": 2}}]}' \
  --data ''
```

---

## Jasper Response com array de cotações (usado o parâmetro "validation_errors")

Criar array de respostas similares a estrutura padrão do array retornado pelo Jasper antes de avaliar o preço de cada carrier.

```json
{
  "api call id customer": [],
  "debug_id": "",
  "version": "Array de cotações Jasper",
  "to": {
    "zip_code": "",
    "address": {
      "sublocality": {"name": ""},
      "locality": {"name": ""},
      "administrative_area_level_2": {"name": ""},
      "administrative_area_level_1": {
        "short_name": "",
        "name": ""
      },
      "country": {
        "short_name": "",
        "name": ""
      }
    }
  },
  "from": {
    "zip_code": ""
  },
  "items": [{
    "id": "INTERNAL",
    "title": "MIRROR",
    "variation_id": 0,
    "category_id": "",
    "model_id": "",
    "seller": {
      "id": 0
    },
    "dimensions": {
      "category_id": "",
      "model_id": "",
      "length": 0,
      "width": 0,
      "height": 0,
      "weight": 0
    }
  }],
  "options": [{
    "id": 0,
    "name": "INTERNAL",
    "display": "MIRROR",
    "currency_id": "BRL",
    "list_cost": 0,
    "cost": 0,
    "type": "",
    "carrier_id": "",
    "shipping_method_id": 0,
    "estimated_delivery": {
      "date": "0",
      "type": "",
      "offset": {
        "date": "0"
      },
      "time_from": 0,
      "time_to": 0
    },
    "estimated_handling_limit": {
      "date": ""
    }
  }],
  "validation_errors": {
    "zip_code": "",
    "message": "Error code details"
  }
}
```

---

# API Integration Guide

## Trigger response with error mapping

```json
{
  "response": {
    "trigger": {
      "height": 1,
      "weight": 1,
      "width": 15
    }
  },
  "duration": 1,
  "grid": 56.5,
  "pending_time": 0,
  "sorting_time": 30,
  "service_time": 30,
  "service_order": "UNDEFINED"
}

"method_url": "POST"
```

> Trigger response com example de error a seleção completa.

```json
{
  "response": {
    "error": "There should be a valid string 'Error'."
  },
  "url": "https://api.onfleet.com/v2/tasks/:taskId/clone",
  "method": "POST",
  "headers": {
    "Content-Type": "application/json",
    "Authorization": "Bearer dGVzdGluZ0Bjb3VsZGJlYmV0dGVyQWVpdGhlcg=="
  },
  "webhook_url": "AMEND",
  "trigger": "CREATED"
}

"with_body": {
  "url_id": "CHECKED",
  "trigger": "CREATED",
  "webhook": "UNDEFINED",
  "type": "UPDATED",
  "time_id": "UPDATED",
  "items": [
    {
      "time_id": "UPDATED",
      "watch": 5,
      "dev": "UPDATED",
      "obj": {
        "dimension": "CREATED"
      },
      "field": {
        "width": 1,
        "dev": 1,
        "time": 1,
        "items": 1
      },
      "status": 27.67,
      "dimension": {
        "CHECKED": 1,
        "width": 1,
        "dev": 1,
        "items": 1
      }
    }
  ]
}

"response_time_backend_url": {
  "service_code": 400,
  "created": true,
  "message_id": "**1*3*GO:233000",
  "service_id": 9,
  "error": "This is to intended_id if 'ERROR'",
  "trigger": "NO: The webhook can still retrieve."
}

"method_url": "POST"
```

---

## Tempo de resposta limite (da cotação)

- Cotação individual a resposta é considerada dentro da experiência de usuário quando a compra e checkout são efetuadas.
- O tempo de resposta esperado para a cotação de frete é de 3 segundos.
- Caso não seja possível atender dentro do prazo, o parceiro pode retornar erro e permitir que o cliente tente novamente.
- Tempos de resposta superiores a 3s não serão permitidos, inviabilizando a integração. Dessa forma, avaliar e otimizar consultas para garantir que a experiência possa ser cumprida.

---

## Tabela de confiabilidade

- Este rótulos são como resultado apesar do cálculo de frete, caso ocorra algum problema na cotação de frete ou CDI (cancelamento) dos níveis na estruturação.

---

## Configuração do valor de frete para confiabilidade

- Para mais detalhes funções: A democracia que o vendedor solicita ou valores de frete e realizar o canal individual do vendedor, conforme passos descritos abaixo:
- Não confundir valores, acesso: Configurações de Envio e vender.

### Configuração de Envio

[Interface showing configuration options with toggle switches]

---

## Preenchimento do valor de frete da confiabilidade

- Esta pergunta: Será composto a lista para preenchimento de valor de frete. O valor de frete pode ser preenchido por estado ou pode inserir uma CEP focalizado.

---

## Valor de frete único por estado

### Steps:
1. Selecione os estados
2. Preencha o valor de frete únicos
3. Clique em Salvar

[Table showing state-based freight configuration with columns for State, Values, and Options]

---

## Valor de frete único por cidade

### Steps:
1. Selecione os estados
2. Selecione as cidades de cada estado
3. Preencha a opção Select All Cities e selecionar as cidades que deseja informar o valor de frete
4. Preencha o valor de frete únicos
5. Clique em Salvar

[Table showing city-based freight configuration with columns for State, City, Values, and Options]

---

## Habilitação do canal Logística do vendedor

[Interface showing logistics channel configuration settings]

---

**Note:** All interface elements, tables, and configuration screens are preserved in their structural context as shown in the original documentation.

---

# Habilitação do canal Logística do vendedor

## Configurações de Envio

### Canal de Envio
Ativar o Canal de Envio Integrado Shop

### Canais de Envio Integrado
Ao ativar esse tipo de canal, você aceita utilizar o serviço de logística da Shopee para realizar entregas com acessórios de uma máquina para você permitir o pacote diretamente em um hub logístico.

**Logística do vendedor** [Channel Setting Toggle: ON]

### Impressora Térmica
Você pode usar uma impressora térmica conectada via Bluetooth para imprimir facilmente guias de envio ou realizar etiquetas. A sua impressora do momento pode garantir que a impressora térmica foi usada em dispositivos Android. É importante que a configuração de impressão de USB Guia pode estar disponível para o vendedor.

---

## Prazo de entrega para valor de frete de contingência

O prazo de entrega será fixo, quando acionada a tabela de contingência:
- Prazo mínimo: 10 dias corridos
- Prazo máximo: 20 dias úteis

---

## Status dos pedidos e rastreamento

### Desenvolvimento API de atualização de status:

A atualização de status dos pedidos e rastreamento se dará por meio da API de atualização de tracking e será feita com base no número de rastreio, retornado pela API de Ship Confirm (logistics_init_info). No qual será enviado será obrigatório informar o número de rastreio.

Para atualização de informações de rastreio de pedido deve ser utilizado o endpoint v2 logistics_update_tracking_status (OpenAPI). Os status disponíveis são:

1. Pedido Enviado (logistics_pickup_done)
   - O status Enviado deve ser usado apenas na finalização da atualização do status do pedido para Enviado.

2. Pedido Entregue (logistics_delivery_done)
   - O status Entregue só será recebido, caso o pedido já possua status o status Enviado.

3. Falha na Entrega (logistics_delivery_failed)
   - O status Falha na Entrega só será recebido, caso o pedido já possua status o status Enviado.

**IMPORTANTE:** após envios dos status Pedido Entregue ou Falha na Entrega não será mais permitida atualização de status, tornando-os status finalizadores.

Vos parâmetros tracking_number e tracking_url deverão ser enviados apenas na atualização de status para logistics_pickup_done.

Seguir link para documentação da API "https://open.shopee.com/documents/v2/v2.logistics.update_tracking_status"

---

## Fluxo e status de pedidos do canal

### Fluxo de Status:

**1. Status atualiza para "UNPAID_TO_SHIP" logo que é pagamento é confirmado pela API Open Seller.**

↓

**2. O vendedor precisa gerar a NF e enviar via API "nrg_invoice_doc"**

↓

**3. Irma vira a NF emitida, o status atualiza para "READY_TO_SHIP" o envio (API "ship_order") e o status irá atualizar para "SHIPPED"**

→ **INICIO**
→ **UNPAID** (1 e 2)
→ **READY_TO_SHIP**
→ **SHIPPED** (3)
→ **TO_CONFIRM_RECEIVE** (6)
→ **COMPLETED** (7)

**4. Assim que o pedido é coletado pelo parceiro logístico, o vendedor chama a API de Ship Confirm e atualiza o status "logistic_pickup_done" e também pode enviar o tracking_number caso o tracking_url não esteja disponível, para irá atualizado para o status SHIPPED (o status será atualizado)**

**5. Caso tenha problema na entrega o vendedor atualiza o status do pedido para "logistic_delivery_failed", cancelando o pedido (status "CANCELLED") e possibilitando o retorno manual a API de Shopee.**

**6. Caso o pedido seja entregue com sucesso o vendedor chama a API "update_tracking_status" e atualiza o status para "logistic_delivery_done" o pedido irá o status de "TO_CONFIRM_RECEIVE"**

**7. Uma vez que o comprador confirme o recebimento ou após 7 dias o pedido é automaticamente confirmado e o status é atualizado para o status de COMPLETED**

---

## Regras de comissão e frete para o canal

- O desconto máximo de comissão é de 6% mas o vendedor esteja no Programa de Frete Grátis
- O vendedor terá uma coparticipação no desconto oferecido ao cliente.

| Valor pago pelo Shopee | Valor pago pelo vendedor | Limite do valor pago pela Shopee |
|------------------------|-------------------------|----------------------------------|
| 37% | 63% | R$ 7.30 |

---

## FAQ:

### 1. Todos os vendedores tem acesso ao canal Logística do Vendedor?
Não, apenas os vendedores que estão no Programa de Frete Grátis e possuem configurações procure seu Gerente de contas.

### 2. Como identificar uma cotação de pedido?
Quando voz faz um pedido criado com uma cotação de sucesso, o parâmetro "service_code" será retornado no parâmetro "shipping_carrier_order_detail"

Ex.:
- "service_code" enviado: "Canal X"
- retorno do parâmetro "shipping_carrier": "Logística do vendedor - Canal X"

Ex2:
- Tabela de contingência ativada
- retorno do parâmetro "shipping_carrier": "Logística do vendedor"

### 3. Posso usar o APP Seller Logistics para chamar a OpenAPI?
Não, o uso do APP só é usado para API de cotação.

Para mais dúvidas técnicas, abrir um ticket na Plataforma de Tickets da OP.
Dúvidas gerais sobre acesso ao canal logístico, falar com gerente de contas.

---

**문서 ID**: developer-guide.286
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/286
**처리 완료**: 2025-10-16T09:05:20
