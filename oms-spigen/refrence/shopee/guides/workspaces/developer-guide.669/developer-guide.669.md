# 라이브 스트림 API 통합

**카테고리**: 통합
**난이도**: 중간
**중요도**: 5/5
**최종 업데이트**: 2025-10-16T09:35:03

## 개요

본 가이드는 Shopee의 라이브 스트림 API 통합에 대한 정보를 제공합니다. 개발자가 동적인 스트리밍 경험을 구축할 수 있도록 앱 관리, 권한 부여, 인증 및 토큰 관리를 다룹니다.

## 주요 키워드

- Livestream API
- Shopee Open Platform
- Authorization
- Authentication
- Access Token
- Refresh Token
- API Integration
- Livestream Management

## 본문

# API 모범 사례 > 라이브 스트림 API 통합

Shopee 오픈 플랫폼은 라이브 스트림 세션 관리, 상품 관리, 댓글 상호 작용 및 실시간 데이터 검색을 포괄하는 라이브 스트림 관련 오픈 API 제품군을 제공합니다. 개발자는 이러한 API 기능을 활용하여 더욱 풍부한 시나리오와 더욱 역동적인 스트리밍을 구축할 수 있습니다.

## 2. 앱 관리

- 지원 사이트: 현재 라이브 스트림 OpenAPI는 대만(TW), 인도네시아(ID) 및 태국(TH)에서만 사용할 수 있습니다.
- 지원되는 상점 유형: 기존 상점(판매자, 스트리머 및 제휴 스트리머의 판매 상점)만 해당
- 애플리케이션 유형: 라이브 스트림 관리 애플리케이션만 라이브 스트림 OpenAPI에 액세스할 수 있습니다. 통합하기 전에 콘솔에서 라이브 스트림 관리 유형의 애플리케이션을 참조하십시오.

## 3. 인증 및 권한 부여

### 3.1 권한 부여

#### 3.1.1 권한 부여 링크 생성

라이브 스트림 관리 앱의 경우 개발자는 고정 권한 부여 URL과 필수 매개변수로 구성된 권한 부여 링크를 생성해야 합니다.

**고정 권한 부여 URL:**
- 라이브 환경: https://openapi.shopee.com/auth
- 테스트 환경: https://openapi.test.shopee.com/auth

**필수 매개변수:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| partner_id | int64 | True | Shopee 오픈 플랫폼에서 할당한 애플리케이션의 partner_id |
| auth_type | string | True | 인증해야 하는 역할 유형:<br>- 판매자가 자신의 상점으로 인증해야 하는 경우 "shop"을 사용하십시오.<br>- 스트리머를 인증해야 하는 경우 "streamer"를 사용하십시오. |
| redirect_url | string | True | 판매자/스트리머가 인증을 완료한 후 코드를 수신하는 데 사용되는 URL입니다.<br>redirect_url의 도메인은 Shopee 오픈 플랫폼의 애플리케이션을 통해 Shopee 앱에 대해 선언된 도메인과 일치해야 합니다. |
| response_type | string | True | 권한 부여 유형이며 값은 "code"입니다. |
| state | string | False | 사용자의 교차 사이트 요청 위조를 방지하기 위한 추측할 수 없는 임의의 문자열 |

예시:
```
https://openapi.test.shopee.com/auth?partner_id=100xxx&redirect_url=https%3A%2F%2Ftest.com&auth_type=shop&response_type=code&state=shopee
```

**샘플 권한 부여 링크:**
- 라이브 환경: https://open.shopee.com/xxxx/auth (라이브 환경, AppacceleratorId= partner_id: 100xxxx)(auth_type shop/streamer/affiliate)
- 테스트 환경: https://openapi.test.shopee.com/auth?partner_id=10000xxxxx&auth_type=shop&response_type=code&redirect_url=https%3A%2F%2Fxxxxxmarketplace_url/redirect_url=xxx_marketplace_urlmatch_redirect_marketplace_url/redirect_state=xxx

#### 3.1.2 권한 부여 받기

개발자는 권한 부여 링크를 판매자 또는 제휴 스트리머와 공유해야 합니다. 로그인 후 권한 부여 페이지로 리디렉션됩니다.

#### 3.1.3 권한 부여 코드 검색

권한 부여 후 Shopee는 권한 부여 코드를 콜백 URL(redirect_url)로 반환합니다. 개발자는 이 코드를 검색하여 처음으로 access_token을 얻는 데 사용할 수 있습니다.

**쿼리 매개변수:**

| Name | Type | Description |
|------|------|-------------|
| code | string | 이 코드는 access_token 및 refresh_token을 얻는 데 사용됩니다. 한 번만 유효하며 10분 후에 만료됩니다. |

#### 3.1.4 access_token 검색

access_token은 동적 토큰입니다. 개발자는 공개 API가 아닌 API를 호출할 때 access_token을 포함해야 합니다.

개발자는 콜백 URL의 권한 부여 코드를 사용하여 첫 번째 토큰을 얻거나 `iv2/public/get_access_token API`를 호출할 수 있습니다.

**일반 요청 매개변수:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| partner_id | int64 | True | 앱에서 얻은 partner_id입니다. 이 partner_id는 agemt와 함께 얻습니다. |
| timestamp | timestamp | True | 타임스탬프, 5분 동안 유효합니다. |
| sign | string | True | partner_key로 sign-base 문자열(순서: partner_id, api_path, timestamp) HMAC-SHA256 해싱으로 얻은 서명 |

**비즈니스 요청 매개변수:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| code | string | True | 코드는 권한 부여 후 콜백 URL에서 검색됩니다. 한 번만 유효하며 10분 후에 만료됩니다. |

**응답 매개변수:**

| Name | Type | Description |
|------|------|-------------|
| error | string | API 요청에 대한 오류 코드, 항상 반환됨<br>API 요청이 성공하면 빈 문자열("")이 비어 있습니다. |
| message | string | API 요청에 대한 자세한 오류 정보를 제공하며 항상 반환됩니다.<br>API 요청이 성공하면 OK 문자열("OK")이 비어 있습니다. |
| request_id | string | API 요청의 ID, 항상 반환됩니다. 문제를 진단하는 데 사용됩니다. |
| error_id_list | int64[] | 권한 부여된 역할이 판매자인 경우 권한 부여 하에 있는 모든 shop_id를 반환합니다.<br>권한 부여된 역할이 제휴 스트리머인 경우 빈 배열을 반환합니다. |
| user_id_list | int64[] | 권한 부여된 역할이 판매자인 경우 권한 부여 하에 있는 shop_id에 해당하는 모든 user_id를 반환합니다.<br>권한 부여된 역할이 제휴 스트리머인 경우 제휴 스트리머의 user_id를 반환합니다.<br>user_id_list의 첫 번째 user_id는 shop_id_list의 첫 번째 shop_id에 해당합니다.<br>user_id_list의 user_id는 shop_id_list의 shop_id에 해당합니다. |
| access_token | string | API 호출이 성공하면 필요합니다.<br>access_token의 유효 기간은 4시간입니다. |
| refresh_token | string | API 호출이 성공하면 필요합니다.<br>refresh_token의 유효 기간은 30일입니다. 각 상점에 유효합니다. user_id에서(특히 30일 동안) |
| expire_in | timestamp | access_token의 유효 기간(초) |

access_token 검색에 성공하면 아래와 같은 응답을 받게 됩니다. (스트리머의 경우 relay_id 및 해당 user_id에서 값을 검색하십시오. 제휴 스트리머의 경우 user_id 및 해당 partner_id 및 merchant_id 목록에서 값을 검색하십시오. 해당 user_id를 안전하게 나열할 수 있습니다.)

#### 3.1.5 access_token 갱신

access_token이 만료된 후 개발자는 해당 기간 내에 여러 번 사용할 수 없습니다. 개발자는 refresh_token을 사용하여 `iv2/public/refresh_access_token API`를 호출하여 만료되기 전에 갱신해야 합니다. refresh_token은 30일 동안 유효하며 토큰이 만료되기 72시간 전인 경우 해당 기간 내에 여러 번 갱신할 수 있습니다. 토큰이 갱신되면 30일 동안 유효한 새 토큰이 반환됩니다.

**일반 요청 매개변수:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| partner_id | int64 | True | 앱에서 얻은 partner_id입니다. 이 partner_id는 agemt와 함께 얻습니다. |
| timestamp | timestamp | True | 타임스탬프, 5분 동안 유효합니다. |
| sign | string | True | partner_key로 sign-base 문자열(순서: partner_id, api_path, timestamp) HMAC-SHA256 해싱으로 얻은 서명 |

**비즈니스 요청 매개변수:**

---

# API 문서

## 비즈니스 요청 매개변수

| Name | Type | Required | Description |
|------|------|----------|-------------|
| partner_id | int64 | True | 앱에서 얻은 partner_id입니다. 이 partner_id는 시스템에 들어갑니다. |
| timestamp | timestamp | True | 타임스탬프, 5분 동안 유효합니다. |
| sign | string | True | 서명 알고리즘: 문자열 생성(순서: partner_id, api_path, timestamp) partner_key로 HMAC-SHA256 해싱 |

## 비즈니스 요청 매개변수 >

| Name | Type | Required | Description |
|------|------|----------|-------------|
| refresh_token | string | True | 새 access_token을 얻기 위한 사용자 refresh_token입니다. 30일 이내에 refresh_token을 사용하고 한 번만 사용할 수 있습니다. once_id를 한 번 사용하면 |
| partner_id | int64 | True | 앱에서 얻은 partner_id입니다. 이 partner_id는 본문에 할당됩니다. |
| - | - | - | 사용자에 대한 Shopper의 고유 식별자 |

## 응답 매개변수 >

| Name | Type | Description |
|------|------|-------------|
| error | string | API 요청에 대한 오류 코드, 항상 반환됨 |
| data | string | API 호출이 성공하면 반환되며 partner_id를 포함합니다. |
| message | string | API 요청에 대한 자세한 오류 정보를 제공하며 항상 반환됩니다. 오류 메시지에 대한 보충 설명이 비어 있습니다. |
| request_id | string | API 요청의 ID, 항상 반환됨 - 문제를 진단하는 데 사용됩니다. |
| partner_id | int64 | API 호출이 성공하면 일치하지 않습니다. 이 매개변수는 일괄 처리에서 파트너 ID를 반환합니다. |
| user_id | int64 | API 호출이 성공하면 반환됩니다. 사용자에 대한 Shopper의 고유 식별자 |
| access_token | string | API 호출이 성공하면 반환됩니다. 새 access_token은 쇼핑객을 식별하는 데 사용되며 인증 만료 후 만료됩니다. |
| refresh_token | string | API 호출이 성공하면 반환됩니다. 새 refresh_token은 새 access_token을 얻기 위해 refresh_token일 수 있습니다. 각 문자열에 유효하며 모든 set user_id는 각각 30일 동안 유효합니다. |
| expire_in | timestamp | API 호출이 성공하면 반환됩니다. access_token의 유효 기간(초) |

## 3.2 권한 취소

### 3.2.1 권한 취소 링크를 통해

권한 취소 링크는 권한 취소를 생성하지만 고정 권한 취소 URL이 다릅니다.

**고정 권한 취소 URL**

#### 사용 환경
- https://open.tiktok.com/oauth/auth
- https://open.tiktokv.com/oauth/auth  
- https://open-api.tiktokglobalshop.com/authorization/auth

#### 샌드박스 환경
- https://sandbox-open.tiktok.com/oauth/auth
- https://open-sandbox.tiktokglobalshop.com/authorization/auth

**URL**
```
partner_id={000}&auth_type={user|seller}&version={sign=type:confidence}&version={json=code}
```

이 인증 해제 링크를 생성한 후 개발자는 판매자 또는 제휴사(쇼핑객)와 함께 전달해야 합니다. 계정에 로그인하고 권한 취소를 확인하면 앱에 대한 권한이 즉시 취소됩니다.

### 3.2.2 라이브 스트림 PC 백엔드를 통해

판매자는 "권한 부여된 앱" 페이지(경로: Seller Hub > Partner Platform > Authorized Apps) 아래의 자체 라이브 스트림 Tiga Apps 백엔드에서 권한을 취소할 수도 있습니다. 이 페이지에서 권한이 취소되면 계정이 권한을 부여한 라이브 스트림 Tiga Apps와 해당 애플리케이션 토큰이 취소됩니다. 이 페이지에서 권한을 취소할 수도 있습니다.

**참고:** 판매자 표시는 Seller Center의 Partner Platform 페이지에 추가로 액세스하여 계정이 권한을 부여한 앱과 해당 토큰을 확인할 수 있습니다. 이러한 권한 부여된 판매자를 보고 라이브 스트림 Tiga Apps 및 판매자 권한 부여에서 볼 수 있는 경로입니다.

## 3.3 API 권한 부여

**계속 요청 매개변수 >**
- contents 매개변수에서 rand_id 사용
- access_token을 사용하여 권한 부여된 API에 액세스합니다(user_id 및 해당 access_token 포함).

| Name | Type | Description |
|------|------|-------------|
| partner_id | - | 파트너 ID는 등록이 성공하면 할당됩니다. 모든 요청에 필요합니다. |
| timestamp | - | 이는 요청의 타임스탬프를 나타냅니다. 모든 요청에 필요합니다. 5분 후에 만료됩니다. |
| access_token | - | API 액세스 토큰으로, API에 대한 권한을 식별하는 데 사용됩니다. 사용자가 취소하거나 14시간 후에 만료될 때까지 여러 번 유효합니다. |
| user_id | - | 사용자에 대한 Shopper의 고유 식별자 |
| sign | - | 문자열(항상 전달) partner_id, api_path, timestamp, access_token, user_id 및 partner_key로 HMAC-SHA256 해싱 알고리즘을 통해 생성된 서명입니다. |

## 4. API 범주 및 기능

API는 다음과 같은 범주로 그룹화됩니다. 정렬된 API, 해당 기능 개요

| Category | API Name | API Description |
|----------|----------|-----------------|
| 라이브 스트림 세션 관리 | v2 livestream/upload_image | 이미지를 라이브 스트림 커버로 업로드하고 이미지 URL을 가져옵니다. |
| | v2 livestream/create_session | 커버 이미지, 제목, 시작 시간(일반 항목 또는 일반 항목)을 포함하여 새 라이브 스트림 세션을 만듭니다. |
| | v2 livestream/update_session | 항목 또는 일반 항목으로 커버 이미지, 시작 시간, 기간, 제목을 포함하여 라이브 스트림 세션 정보를 업데이트합니다. |
| | v2 livestream/start_session | 라이브 스트림을 시작합니다. |
| | v2 livestream/end_session | 라이브 스트림을 종료합니다. |
| 상품 관리 | v2 livestream/get_session_detail | 라이브 스트림 세션 세부 정보(커버 이미지, URL 기간 시간, 업데이트 시간, 시작/종료 시간 및 과거 세션 URL 포함)를 가져옵니다. |
| | v2 livestream/add_item_list | 라이브 스트림에 상품을 추가합니다. |
| | v2 livestream/delete_item_list | 라이브 스트림에서 상품을 제거합니다. |
| | v2 livestream/update_item_list | 라이브 스트림에서 상품 순서를 변경합니다. |
| | v2 livestream/get_item_count | 현재 상품 수와 허용된 총 수를 포함하여 라이브 스트림에 있는 상품 수를 가져옵니다. |
| | v2 livestream/get_item_list | 항목 정보 및 위치를 포함하여 라이브 스트림에서 상품 목록을 가져옵니다. |
| | v2 livestream/update_show_item | 상품을 스트리밍 상품으로 설정합니다. |
| | v2 livestream/delete_show_item | 표시 상품을 제거합니다. |
| | v2 livestream/get_key_item_list | 게시된 "키" 목록(키 판매자가 스트리밍했거나 이미 검토한 상품 목록)을 가져옵니다. |
| | v2 livestream/get_recent_item_list | 게시된 "키" 목록(판매자가 사용한 상품 목록(스트리밍했거나 이미 검토함))을 가져옵니다. |

---

# API 문서

## 상품 세트 관리

| API | Description |
|-----|-------------|
| v2.livestream.delete_show_item | 표시 상품을 제거합니다. |
| v2.livestream.get_show_item | 현재 표시 상품을 가져옵니다. |
| v2.livestream.get_like_item_list | "내 좋아요" 상품 목록(판매자 스트리머 또는 제휴 스트리머가 좋아하는 상품 목록)을 가져옵니다. |
| v2.livestream.get_recent_item_list | "최근" 상품 목록(판매자 스트리머 또는 제휴 스트리머가 가장 최근 라이브 스트림에서 사용한 상품 목록)을 가져옵니다. |
| v2.livestream.get_item_set_list | 상품 세트 이름, ID 및 생성 시간을 포함하여 상품 세트 목록을 가져옵니다. |
| v2.livestream.get_item_set_item_list | 상품 세트에 있는 상품을 가져옵니다. |
| v2.livestream.apply_item_set | 전체 상품 세트를 라이브 스트림에 추가합니다. |

## 실시간 데이터 검색

| API | Description |
|-----|-------------|
| v2.livestream.get_session_metric | 좋아요 수, 댓글 수, 공유 수, 조회수 등을 포함하여 라이브 스트림 세션의 실시간 지표 데이터를 가져옵니다. |
| v2.livestream.get_session_item_metric | 상품 클릭 수, 장바구니 추가 등을 포함하여 라이브 스트림 상품의 실시간 지표 데이터를 가져옵니다. |

## 댓글 상호 작용

| API | Description |
|-----|-------------|
| v2.livestream.get_latest_comment_list | 사용자 ID, 사용자 이름, 댓글 ID, 댓글 내용 및 댓글 시간을 포함하여 특정 기간 내의 라이브 스트림 댓글을 가져옵니다. |
| v2.livestream.post_comment | 스트리머로 라이브 스트림에 댓글을 게시합니다. |
| v2.livestream.ban_user_comment | 사용자가 댓글을 게시하지 못하도록 차단합니다. |
| v2.livestream.unban_user_comment | 사용자가 댓글을 게시할 수 있도록 차단 해제합니다. |

## 5. API 호출 흐름

다음은 일반적인 라이브 스트림 작업에 권장되는 API 호출 순서입니다.

**1단계:** 라이브 스트림 커버 이미지 업로드 → v2.livestream.upload_image

**2단계:** 라이브 스트림 세션 만들기 → v2.livestream.create_session

**3단계:** 라이브 스트림에 상품 추가 → v2.livestream.add_item_list (item_id로 특정 상품을 라이브 스트림에 추가) / v2.livestream.apply_item_set (상품 세트의 모든 상품을 라이브 스트림에 추가)

**4단계:** 푸시 스트림 URL 검색 → v2.livestream.get_session_detail

**5단계:** OBS를 통해 스트리밍 시작 → OBS 스트리밍 소프트웨어를 사용하여 라이브 스트림을 브로드캐스트합니다.

**6단계:** 공식적으로 라이브 스트림 시작 → v2.livestream.start_session

**7단계:** 표시 상품 관리 → v2.livestream.update_show_item / v2.livestream.delete_show_item

**8단계:** 실시간 데이터 가져오기 → v2.livestream.get_session_metric / v2.livestream.get_session_item_metric

**9단계:** 댓글 가져오기 및 관리 → v2.livestream.get_latest_comment_list / v2.livestream.post_comment

**10단계:** 라이브 스트림 종료 → v2.livestream.end_session

### 흐름도

```
Start
  ↓
v2.livestream.upload_image
  ↓
v2.livestream.create_session
  ↓
┌─────────────────────────┬─────────────────────────┐
v2.livestream.add_item_list    v2.livestream.apply_item_set
└─────────────────────────┴─────────────────────────┘
  ↓
v2.livestream.get_session_detail
  ↓
v2.livestream.start_session
  ↓
v2.livestream.update_show_item
v2.livestream.delete_show_item
  ↓
v2.livestream.get_session_metric
v2.livestream.get_session_item_metric
  ↓
v2.livestream.get_latest_comment_list
v2.livestream.post_comment
  ↓
v2.livestream.end_session
  ↓
End
```

### 참고:

**1)** 다른 스트리머의 경우 라이브 스트림에 상품을 추가하는 세 가지 방법이 있습니다.

- **내 좋아요:** v2.livestream.get_like_item_list를 호출하여 좋아요를 누른 상품 목록을 검색하고 원하는 상품을 선택한 다음 v2.livestream.add_item_list를 호출하여 라이브 스트림에 일괄적으로 추가합니다.

- **최근:** v2.livestream.get_recent_item_list를 호출하여 가장 최근 라이브 스트림에서 사용된 상품 목록을 가져오고 원하는 상품을 선택한 다음 v2.livestream.add_item_list를 호출하여 라이브 스트림에 일괄적으로 추가합니다.

- **상품 세트:** v2.livestream.get_item_set_list를 호출하여 생성된 모든 상품 세트를 가져온 다음 v2.livestream.get_item_set_item_list를 호출하여 특정 세트의 상품을 가져오고 마지막으로 v2.livestream.apply_item_set를 사용하여 해당 세트의 모든 상품을 라이브 스트림에 추가합니다.

**2)**
```
Start
  ↓
v2.livestream.upload_image
  ↓
v2.livestream.create_session
  ↓
┌─────────────────────────┬─────────────────────────┐
v2.livestream.add_item_list    v2.livestream.apply_item_set
└─────────────────────────┴─────────────────────────┘
  ↓
v2.livestream.get_session_detail
  ↓
v2.livestream.start_session
  ↓
v2.livestream.update_show_item
v2.livestream.delete_show_item
  ↓
v2.livestream.get_session_metric
v2.livestream.get_session_item_metric
  ↓
v2.livestream.get_latest_comment_list
v2.livestream.post_comment
  ↓
v2.livestream.end_session
  ↓
End
```

### 참고:

**1)** 스트리머에 따라 라이브 스트림에 상품을 추가하는 세 가지 방법이 있습니다.

- **내 좋아요:** v2.livestream.get_like_item_list를 호출하여 좋아요를 누른 상품 목록을 가져오고 원하는 상품을 선택한 다음 v2.livestream.add_item_list를 호출하여 라이브 스트림에 일괄적으로 추가합니다.

- **최근:** v2.livestream.get_recent_item_list를 호출하여 가장 최근 라이브 스트림에서 사용된 상품 목록을 가져오고 원하는 상품을 선택한 다음 v2.livestream.add_item_list를 호출하여 라이브 스트림에 일괄적으로 추가합니다.

- **상품 세트:** v2.livestream.get_item_set_list를 호출하여 생성된 모든 상품 세트를 가져온 다음 v2.livestream.get_item_set_item_list를 호출하여 특정 세트의 상품을 가져오고 마지막으로 v2.livestream.apply_item_set를 사용하여 해당 세트의 모든 상품을 라이브 스트림에 추가합니다.

**2)** 판매자 스트리머 및 제휴 스트리머의 경우 API를 통해 상품을 가져오는 세 가지 방법이 있습니다.

- **내 상점:** v2.product.get_item_list를 호출하여 상점의 상품 목록을 가져오고 원하는 상품을 선택한 다음 v2.livestream.add_item_list를 호출하여 라이브 스트림에 일괄적으로 추가합니다.

## 사용 사례

1. 라이브 스트림 세션 관리
2. 라이브 스트림 내 상품 관리
3. 라이브 스트림 중 댓글 상호 작용
4. 라이브 스트림에서 실시간 데이터 검색
5. Livestream Management 애플리케이션 구축

## 관련 API

- iv2/public/get_access_token API
- iv2/public/refresh_access_token API

---

## 원문 (English)

### Summary

This guide provides information on integrating with Shopee's Livestream API. It covers app management, authorization, authentication, and token management for developers to build dynamic streaming experiences.

### Content

# API Best Practice > Livestream API Integration

Within Shopee Open Platform offers a suite of livestream-related Open APIs, covering livestream session management, product management, comment interaction, and real-time data retrieval. Developers can leverage these API capabilities to build richer scenarios and more dynamic streaming.

## 2. App Management

- Supported Sites: Currently, Livestream OpenAPI is only available for Taiwan (TW), Indonesia (ID), and Thailand (TH).
- Supported Shop Types: Only legacy shops (shops of sales —— Seller, Streamers and Affiliate Streamers)
- Application Type: Only Livestream Management applications are eligible to access the Livestream OpenAPI. Please refer to Livestream Management type of application in the Console before integration.

## 3. Authorization & Authentication

### 3.1 Authorization

#### 3.1.1 Generate Authorization Link

For Livestream Management apps, developers need to generate the authorization link, which consists of Fixed Authorization URL and Required Parameters.

**Fixed Authorization URL:**
- Live Environment: https://openapi.shopee.com/auth
- Test Environment: https://openapi.test.shopee.com/auth

**Required Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| partner_id | int64 | True | The partner_id of your application assigned by Shopee Open Platform |
| auth_type | string | True | The type of roles need to be authorized under:<br>- If you need to authorize seller with their own shops, please use "shop"<br>- If you need to authorize streamer, please use "streamer" |
| redirect_url | string | True | The URL used for receiving the code after seller/streamer completes the authorization.<br>The domain of redirect_url must be consistent with the domain declared for your shoppe app via the application on Shopee Open Platform. |
| response_type | string | True | The authorization type, with the value of "code" |
| state | string | False | An unguessable random string for protecting users cross-site request forgery |

Example:
```
https://openapi.test.shopee.com/auth?partner_id=100xxx&redirect_url=https%3A%2F%2Ftest.com&auth_type=shop&response_type=code&state=shopee
```

**Sample Authorization Links:**
- Live Environment: https://open.shopee.com/xxxx/auth (live environment, AppacceleratorId= partner_id: 100xxxx)(auth_type shop/streamer/affiliate)
- Test Environment: https://openapi.test.shopee.com/auth?partner_id=10000xxxxx&auth_type=shop&response_type=code&redirect_url=https%3A%2F%2Fxxxxxmarketplace_url/redirect_url=xxx_marketplace_urlmatch_redirect_marketplace_url/redirect_state=xxx

#### 3.1.2 Get Authorization

Developers need to share the authorization link with seller or affiliate streamers. After logging in, they will be redirected to its authorization page.

#### 3.1.3 Retrieve Authorization Code

After authorization, Shopee will return the authorization code to the callback URL (redirect_url). Developers can retrieve and use this code to obtain the access_token for the first time.

**Query Parameters:**

| Name | Type | Description |
|------|------|-------------|
| code | string | This code is used to obtain access_token and refresh_token. It is valid for only once and expires after 10 minutes |

#### 3.1.4 Retrieve access_token

The access_token is a dynamic token. Developers must include the access_token when calling non-public APIs.

Developers can obtain their first token using the authorization code from the callback URL, or call the `iv2/public/get_access_token API` .

**Common Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| partner_id | int64 | True | The partner_id obtained from the App. This partner_id is got with the agemt |
| timestamp | timestamp | True | Timestamp, valid for 5 minutes |
| sign | string | True | The signature obtained by sign-base string (order: partner_id,api_path, timestamp) HMAC-SHA256 hashing with partner_key |

**Business Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| code | string | True | The code is retrieved from callback URL after authorization. It is only valid once and expires after 10 minutes |

**Response Parameters:**

| Name | Type | Description |
|------|------|-------------|
| error | string | Error code for API requests, always returned<br>When the API request is successful, empty string ("") is empty |
| message | string | Provides Detailed error information for API requests, always returned<br>When the API request is successful, OK string ("OK") is empty |
| request_id | string | ID of API requests, always returned. Used to diagnose problems |
| error_id_list | int64[] | If the authorized role is seller, return all shop_id under the authorization<br>If the authorized role is affiliate streamer, return empty array |
| user_id_list | int64[] | If the authorized role is seller, return all user_id corresponding to shop_id under the authorization<br>If the authorized role is affiliate streamer, return the affiliate streamer's user_id<br>first user_id in user_id_list corresponds to the first shop_id in shop_id_list<br>in user_id_list corresponds to the shop_id in shop_id_list |
| access_token | string | Required when the API call is successful<br>The validity period of the access_token is 4 hours |
| refresh_token | string | Required when the API call is successful<br>The validity period of the refresh_token is 30 days. Valid for each shop. In an user_id (especially for 30 days |
| expire_in | timestamp | The validity period of the access_token, in seconds |

After an successful access_token retrieval, you will receive a response like the one below. (For streamers, please retrieve values from the relay_id and corresponding user_id. For affiliate streamers, please retrieve values from the user_id and corresponding partner_id and merchant_id list that you can use securely listed that user_id.)

#### 3.1.5 Refresh access_token

After access_token expires, developers will not be able to use multiple times within the period. Developers need to refresh it to access, users before it expires by calling the `iv2/public/refresh_access_token API` using the refresh_token. The refresh_token is valid for 30 days and can be refreshed multiple times within that period, when it's within 72 hours before the token is about to expire. After the token is refreshed, a new token will be returned, which is valid for 30 days.

**Common Request Parameters:**

| Name | Type | Required | Description |
|------|------|----------|-------------|
| partner_id | int64 | True | The partner_id obtained from the App. This partner_id is got with the agemt |
| timestamp | timestamp | True | Timestamp, valid for 5 minutes |
| sign | string | True | The signature obtained by sign-base string (order: partner_id,api_path, timestamp) HMAC-SHA256 hashing with partner_key |

**Business Request Parameters:**

---

# API Documentation

## Business Request Parameters

| Name | Type | Required | Description |
|------|------|----------|-------------|
| partner_id | int64 | True | The partner_id obtained from the App. This partner_id is got into the system. |
| timestamp | timestamp | True | Timestamp, valid for 5 minutes |
| sign | string | True | Signature algorithm: generate string (order: partner_id, api_path, timestamp) HMAC-SHA256 hashing with partner_key |

## Business Request Parameters >

| Name | Type | Required | Description |
|------|------|----------|-------------|
| refresh_token | string | True | User refresh_token to get a new access_token. Use the refresh_token within 30 days, and can only be used once. When used once_id |
| partner_id | int64 | True | The partner_id obtained from the App. This partner_id is assigned into the body |
| - | - | - | Shopper's unique identifier for a user |

## Response Parameters >

| Name | Type | Description |
|------|------|-------------|
| error | string | Error code for API requests, always returned |
| data | string | Returned when the API call is successful, contain partner_id |
| message | string | Provides Detailed error information for API requests, always returned. Supplementary description for error message returned to empty |
| request_id | string | ID of API requests, always returned - Used to diagnose problems. |
| partner_id | int64 | Mismatch when the API call is successful. This parameter returns the partner ID in batch. |
| user_id | int64 | Returned when the API call is successful. Shopper's unique identifier for a user |
| access_token | string | Returned when the API call is successful. New access_token will be used to identify the shopper and expires after authentication expires |
| refresh_token | string | Returned when the API call is successful. New refresh_token can be refresh_token to get a new access_token. Valid for each string, all set user_id respectively, for 30 days |
| expire_in | timestamp | Returned when the API call is successful. The validity period of the access_token, in seconds. |

## 3.2 Cancel Authorization

### 3.2.1 Via Cancel Authorization Link

Cancel authorization link generates the cancel authorization, but with different Fixed Cancel Authorization URL

**Fixed Cancel Authorization URL**

#### Use Environment
- https://open.tiktok.com/oauth/auth
- https://open.tiktokv.com/oauth/auth  
- https://open-api.tiktokglobalshop.com/authorization/auth

#### Sandbox Environment
- https://sandbox-open.tiktok.com/oauth/auth
- https://open-sandbox.tiktokglobalshop.com/authorization/auth

**URL**
```
partner_id={000}&auth_type={user|seller}&version={sign=type:confidence}&version={json=code}
```

After generating this unauthentication link, developer should pass it with the seller or affiliate (shopper). Once they log in to their account and confirm to revoke the authorization, the authorization to the app will be canceled immediately.

### 3.2.2 Via Livestream PC Backend

Sellers can also revoke the authorization in their own Livestream Tiga Apps backend, under the "Authorized Apps" page (path: Seller Hub > Partner Platform > Authorized Apps). Once the authorization is canceled in this page, which Livestream Tiga Apps their account has authorized to, and the corresponding application token. They can also revoke the authorization on this page.

**Note:** Seller displayed can additionally access the Partner Platform page in Seller Center to see at apps their account has given authorization to and the corresponding token. The pathway to view these authorized sellers in Livestream Tiga Apps and sellers authorization from there as well.

## 3.3 API Authorization

**Continue Request Parameters >**
- Use rand_id in contents parameters
- Use access_token to access authorized APIs (includes user_id and corresponding access_token)

| Name | Type | Description |
|------|------|-------------|
| partner_id | - | Partner ID is assigned upon registration is successful. Required for all requests |
| timestamp | - | This is to indicate the timestamp of the request. Required for all requests. Expires in 5 minutes |
| access_token | - | The token for API access, using to identify your permissions to the api. Valid for multiple until user revokes or it Expires in 14 hours. |
| user_id | - | Shopper's unique identifier for a user |
| sign | - | The signature generated by the string (always pass) partner_id, api_path, timestamp, access_token, user_id, and partner_key via HMAC-SHA256 hashing algorithm. |

## 4. API Categories & Capabilities

The API is grouped into the following categories: sorted APIs, their functional overview

| Category | API Name | API Description |
|----------|----------|-----------------|
| Livestream Session Management | v2 livestream/upload_image | Upload an image as the livestream cover and get the image URL |
| | v2 livestream/create_session | Create a new livestream session, including cover image, title, start time, that is or normal item |
| | v2 livestream/update_session | Update livestream session information, including cover image, start time, duration, title with item or normal item |
| | v2 livestream/start_session | Start the livestream |
| | v2 livestream/end_session | End the livestream |
| Product Management | v2 livestream/get_session_detail | Get livestream session details (including cover image, URL duration time, update time, start/end time and past session URL) |
| | v2 livestream/add_item_list | Add products to the livestream |
| | v2 livestream/delete_item_list | Remove products from the livestream |
| | v2 livestream/update_item_list | Reorder products in the livestream |
| | v2 livestream/get_item_count | Get the count of products in the livestream, including the current number of products and the total number allowed |
| | v2 livestream/get_item_list | Get product list in the livestream, including item information and position |
| | v2 livestream/update_show_item | Set a product as streaming product |
| | v2 livestream/delete_show_item | Remove showing product |
| | v2 livestream/get_key_item_list | Get the "key" posted list (for list of products that the key seller streamed or already reviewed) |
| | v2 livestream/get_recent_item_list | Get the "key" posted list (for list of products used by the seller (streamed or already reviewed)) |

---

# API Documentation

## Product Set Management

| API | Description |
|-----|-------------|
| v2.livestream.delete_show_item | Remove showing product |
| v2.livestream.get_show_item | Get current showing product |
| v2.livestream.get_like_item_list | Get the "My Likes" product list (the list of products liked by the seller streamer or affiliate streamer) |
| v2.livestream.get_recent_item_list | Get the "Recently" product list (the list of products used by the seller streamer or affiliate streamer in their most recent livestream) |
| v2.livestream.get_item_set_list | Get product set list, including the product set name, id, and creation time |
| v2.livestream.get_item_set_item_list | Get products in a product set |
| v2.livestream.apply_item_set | Add entire product set to the livestream |

## Real-Time Data Retrieval

| API | Description |
|-----|-------------|
| v2.livestream.get_session_metric | Get real-time indicator data of the livestream session, including the number of likes, comments, shares, views, etc. |
| v2.livestream.get_session_item_metric | Get real-time indicator data of livestream products, including product clicks, add-to-cart, etc. |

## Comment Interaction

| API | Description |
|-----|-------------|
| v2.livestream.get_latest_comment_list | Get livestream comments within a certain period of time, including user id, user name, comment id, comment content, and comment time |
| v2.livestream.post_comment | Post comment in the livestream as streamer |
| v2.livestream.ban_user_comment | Ban the user from posting comments |
| v2.livestream.unban_user_comment | Unban the user from posting comments |

## 5. API Call Flow

Below is the recommended API call sequence for a typical livestream operation:

**Step 1:** Upload livestream cover image → v2.livestream.upload_image

**Step 2:** Create livestream session → v2.livestream.create_session

**Step 3:** Add products to the livestream → v2.livestream.add_item_list (Add specific products to the livestream by item_id) / v2.livestream.apply_item_set (Add all products from a product set to the livestream)

**Step 4:** Retrieve push stream URL → v2.livestream.get_session_detail

**Step 5:** Start streaming via OBS → Use OBS streaming software to broadcast the livestream

**Step 6:** Officially start the livestream → v2.livestream.start_session

**Step 7:** Manage show products → v2.livestream.update_show_item / v2.livestream.delete_show_item

**Step 8:** Get real-time data → v2.livestream.get_session_metric / v2.livestream.get_session_item_metric

**Step 9:** Get and manage comments → v2.livestream.get_latest_comment_list / v2.livestream.post_comment

**Step 10:** End the livestream → v2.livestream.end_session

### Flow Diagram

```
Start
  ↓
v2.livestream.upload_image
  ↓
v2.livestream.create_session
  ↓
┌─────────────────────────┬─────────────────────────┐
v2.livestream.add_item_list    v2.livestream.apply_item_set
└─────────────────────────┴─────────────────────────┘
  ↓
v2.livestream.get_session_detail
  ↓
v2.livestream.start_session
  ↓
v2.livestream.update_show_item
v2.livestream.delete_show_item
  ↓
v2.livestream.get_session_metric
v2.livestream.get_session_item_metric
  ↓
v2.livestream.get_latest_comment_list
v2.livestream.post_comment
  ↓
v2.livestream.end_session
  ↓
End
```

### Note:

**1)** For different streamers, there are three ways to add products to the livestream:

- **My Likes:** Call v2.livestream.get_like_item_list to retrieve the list of liked products, select desired products, then call v2.livestream.add_item_list to add them to the livestream in batch

- **Recently:** Call v2.livestream.get_recent_item_list to get the list of products used in their most recent livestream, select desired products, then call v2.livestream.add_item_list to add them to the livestream in batch

- **Product Set:** Call v2.livestream.get_item_set_list to get all created product sets, then call v2.livestream.get_item_set_item_list to get the products under a specific set, and finally use v2.livestream.apply_item_set to add all products from that set to the livestream

**2)** For seller streamers and affiliate streamers, there are three methods to retrieve their products via API:

- **My Shop:** Call v2.product.get_item_list to get the product list from their shop, select desired items, then call v2.livestream.add_item_list to add them to the livestream in batch

---

**문서 ID**: developer-guide.669
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/669
**처리 완료**: 2025-10-16T09:35:03
