# Sandbox Testing V2 개발자 가이드

**카테고리**: 통합
**난이도**: 중간
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T08:18:47

## 개요

본 가이드는 개발자를 위해 Open Platform에서 제공하는 Sandbox V2 테스트 환경을 소개합니다. 테스트 계정 생성, 테스트 주문 생성, 테스트 계정 권한 부여를 다루며, 개발자가 격리된 환경에서 API 기능을 테스트할 수 있도록 지원합니다.

## 주요 키워드

- Sandbox
- 테스트
- 테스트 계정
- 테스트 주문
- 권한 부여
- Open API
- Seller Center
- 상품 관리
- 주문 처리

## 본문

```ko
# 샌드박스 테스트 V2

**시작하기 >** 샌드박스 테스트 V2

## 개요

샌드박스는 Open Platform에서 개발자에게 제공하는 격리된 테스트 환경입니다. 다양한 유형의 테스트 계정과 데이터를 제공합니다. 개발자는 샌드박스 환경에서 대부분의 API 기능 테스트를 완료할 수 있습니다. 샌드박스는 일부 기능 및 인터페이스에 대한 기능을 제공하지만 제품 관리, 주문 처리 등과 같은 대부분의 시나리오만 다룹니다.

**샌드박스 V2 지원 범위**

| 포털 | 기능 | 네트워크 |
|--------|----------|---------|
| 콘솔 | 테스트 상점 계정 생성 | |
| | 주문 | 테스트 주문 생성 |
| | | 테스트 데이터 푸시 |
| 판매자 센터 | 상품 | 글로벌 SKU 생성 및 관리 |
| | 주문 | 테스트 주문 게시 및 관리 |
| | 주문 | 주문 |
| | | 상점 주문 | 현재 영수증 인쇄는 지원되지 않으므로 Open API를 사용하여 인쇄하십시오. |
| Open API | 상품 | 모든 API | 콘솔의 API 테스트 도구를 통하거나 도메인 이름 https://openplatform.sandbox.test 를 사용하여 직접 호출할 수 있습니다. |
| | 글로벌 상품 | 모든 API | 상태 항상 켜짐: ON, https://openplatform.sandbox.test (단일 SKU)를 사용하십시오. |
| | 미디어 리소스 | 모든 API | |
| | 주문 | 모든 API | |
| | 물류 | 모든 API | |
| | 퍼스트 마일 | 모든 API | |
| | 상점 | 모든 API | |
| | 판매자 | 모든 API | |
| 푸시 서비스 | 일부 푸시 테스트 데이터 수신 지원, 자세한 내용은 2.3 참조 | |

> 추가 샌드박스 지원 기능이 필요한 경우 특정 사용 사례 및 요구 사항을 포함하여 고객 서비스에 문의하십시오.

이 문서는 주로 테스트 계정, 테스트 주문 생성 및 샌드박스 테스트 환경 배송 프로세스를 소개합니다.

---

## 1. 테스트 계정 생성

**빠른 콘솔:** 테스트 계정-샌드박스 v2를 선택하고 6에서 **테스트 계정 생성**을 클릭합니다.

### 1.1 로컬 및 크로스보더 테스트 상점

로컬 및 크로스보더 테스트 상점은 품목 카테고리, 사용 가능한 운송 채널 및 결제 방법 등 여러 측면에서 다릅니다. 서비스 시장에 해당하는 테스트 상점을 선택하십시오.

[“테스트 계정 생성” 버튼이 강조 표시된 계정 생성 인터페이스 스크린샷]

### 1.2 판매자

중국 판매자 센터(CNSC)에 저장하기 위해 개발자는 "판매자"를 선택하여 테스트 마스터 계정을 생성하고 바인딩할 수 있습니다.

[탐색 화살표가 있는 판매자 계정 생성 인터페이스 스크린샷]

---

## 2. 테스트 계정을 테스트 Partner_id에 권한 부여

### 2.1 사전 단계:

1. 다음을 생성합니다.

[여러 패널과 화살표가 있는 권한 부여 흐름 스크린샷]

2. 샌드박스 계정 생성

[샌드박스 계정 생성 인터페이스 스크린샷]

> 계정 유형을 선택하고 상점 또는 판매자 유형 테스트 계정을 만들 수 있습니다.

[계정 유형 선택 대화 상자 스크린샷]

---

### 2.2 상점 권한 부여 문서

**권한 부여 참고 사항:**

- 권한 부여를 위해 상점에는 특정 권한 부여 진입점에 대한 버튼이 있습니다(https://open.shopee.com/documents/v2/v2.push.get_config 에서 참조하거나 아래 단계를 따를 수 있음).

[여러 제품 패널과 화살표가 있는 권한 부여 인터페이스 스크린샷]

---

**참고:** 이 추출은 개발자 가이드 스크린샷에 나타나는 문서 구조, 제목, 표, 단계 및 참고 사항을 유지합니다. 모든 텍스트는 요청대로 영어로 보존되었습니다.

---

# 개발자 통합 가이드

## 권한 부여 흐름

### 테스트/Partner_id의 경우 특정 권한 부여 프로세스에 대해 아래 그림을 설정합니다. 권한 부여 관련 문서를 참조하거나 아래 단계를 따르십시오.

---

## 1단계: 권한 부여 흐름 설정

**권한 부여 링크를 사용하고 해당 `partner_id`를 입력합니다:**

```
https://open.sandbox.best.statestreet.com/oauth2/
```

```
https://web.sandbox.test.statestreet.com/page?id=wfm*oauth*authorize&response_type=code
```

---

## 권한 부여 페이지

### 샌드박스 Best Review에 로그인
**사용자 샌드박스 Best Review**

[권한 부여 버튼]

**참고:** 샌드박스 환경에 대한 고정 권한 부여 URL은 https://account.sandbox.best.statestreet.com 을 포함합니다(라이브 계정인 경우 판매자-계정 가능 로그인 URL을 확인하십시오. 그렇지 않으면 '계정/비밀번호 확인 실패' 오류가 보고됩니다.

---

## 2단계: 해당 샌드박스 계정을 입력하고 로그인합니다.

### 권한 부여
[2단계를 나타내는 확인 표시 아이콘]

- 파트너 앱이 결제 및 계정 사용에 동의
- 거래의 결제 정보를 수정할 수 없음
- 앱이 거래 정보를 읽을 수 있음
- 파트너 앱이 계정 잔액을 쿼리할 수 있음
- 거래 또는 결제의 결제 정보를 수정할 수 없음
- 완료된 결제의 거래 정보를 쿼리할 수 있음

---

## 3단계: 권한 부여를 클릭하여 성공 페이지로 이동합니다.

### 권한 부여
[확인 표시 아이콘]

**참고:** 권한 부여 후 애플리케이션 페이지로 돌아갑니다.

**권한 부여 성공!**

**리디렉션 URL:**
- https://

[뒤로] [확인]

---

## 2.3 판매자 권한 부여 문서

테스트/Partner_id의 경우 특정 권한 부여 프로세스에 대해 아래 그림을 설정합니다. 권한 부여 관련 문서를 참조하거나 아래 단계를 따르십시오.

---

## 권한 부여 단계 (판매자)

**권한 부여 링크를 사용하고 해당 `partner_id`를 입력합니다:**

```
https://open.sandbox.best.statestreet.com/oauth2/
```

```
https://web.sandbox.test.statestreet.com/page?id=wfm*oauth*authorize&response_type=code
```

판매자 권한 부여를 위해 "계정 신뢰"를 클릭합니다.

---

## 샌드박스 Best Review에 로그인
**샌드박스 Best Review**

[권한 부여 버튼 - 강조 표시됨]

**참고:** 샌드박스 환경에 대한 고정 권한 부여 URL은 https://account.sandbox.best.statestreet.com 을 포함합니다(라이브 계정인 경우 판매자-계정 가능 로그인 URL을 확인하십시오. 그렇지 않으면 '계정/비밀번호 확인 실패' 오류가 보고됩니다.

인증 코드는 "123456"입니다.

---

## 확인

**인증 코드 입력**

[확인 버튼]

---

## 4단계: 로그인 후 권한 부여가 필요한 상태를 확인합니다.

### 성공 페이지 플랫폼

**권한 부여/API/확인**

[확인 표시 아이콘]

**권한 부여/확인 - 완료됨 - 성공 - 플랫폼**

[여러 권한 부여 상태 항목 나열됨]

---

## 5단계: 권한 부여를 클릭하여 성공 페이지로 이동합니다.

### 권한 부여/API/확인

[확인 표시 아이콘]

**권한 부여 성공!**

**리디렉션 URL:**
- 샌드박스

[뒤로] [확인 - 강조 표시됨]

---

**참고:** 추출된 콘텐츠는 개발자 가이드 스크린샷의 계층 구조, 코드 예제, 단계, 참고 사항 및 경고를 유지합니다. 모든 텍스트는 원본 문서에 나타난 대로 영어로 보존되었습니다.

---

# 3. 샌드박스 테스트 프로세스

## 3.1 상점 계정

### 3.1.1 판매자 센터에 로그인

콘솔->**귀하의 계정-샌드박스** 페이지 선택->생성된 테스트 상점의 오른쪽에 있는 판매자 센터 로그인을 클릭하여 판매자 센터 페이지로 들어갑니다.

---

*판매자 센터 페이지는 아래와 같이 표시됩니다.*

---

## 3.1.2 테스트 상품 생성

### 3.1.2.1 판매자 센터 또는 Open API를 통해 상품을 생성하도록 선택할 수 있습니다.

다음은 판매자 센터를 통해 생성된 두 가지 테스트 상품입니다.

---

*필수 필드를 모두 채운 다음 기본 및 게시를 선택합니다.*

---

### 3.1.2.2 상품을 성공적으로 생성한 후 판매자 센터에서 상품(내 상품)을 볼 수 있습니다.

---

## 3.1.3 테스트 주문 생성

### 3.1.3.1 시뮬레이션 도구-테스트 주문 구매 페이지를 클릭하고 "테스트 주문 생성"을 클릭하여 지금 테스트하려는 주문 상태를 확인합니다.

---

"상점" 드롭다운 상자를 클릭하고 생성하려는 상점을 선택합니다.

---

"양식 선택"(주문 상태)을 선택한 후 생성해야 하는 상자를 선택하고 "확인"을 클릭합니다.

---

### 3.1.3.2 "배송 센터" 드롭다운 상자를 클릭하여 이행 채널을 선택하고 "확인"을 클릭하여 주문 생성을 완료합니다.

---

*참고: 주문 상태에 따라 사용 가능한 이행 채널 옵션이 다를 수 있습니다.*

---

# 주문 관리 가이드

## 3.1.2 "테스트 주문 생성" 완료

[이 섹션은 주문 세부 정보, 상태 및 작업을 포함한 여러 열이 있는 주문 목록 인터페이스를 보여주는 것 같습니다.]

---

## 3.1.4 생성된 주문 보기

판매자 센터에 들어가서 "내 주문"을 클릭하여 생성된 주문을 봅니다.

**참고:** Gossipay 페이지에서 주문을 생성한 후 다음 단계를 진행하기 전에 약 5분 정도 기다려야 합니다.

[주문 정보, 상품, 결제, 상태 및 작업에 대한 열이 있는 여러 주문을 보여주는 주문 목록 인터페이스]

---

## 3.1.5 배송

"배송 준비"를 클릭하고 배송 방법(픽업/드롭오프)을 선택합니다. 추적 번호가 자동으로 생성됩니다. 이 테스트에서 주문 상태는 **처리 중**입니다.

**참고:** 먼저 "배송 준비 중" 상태의 주문을 취소하십시오. "jet" 태그는 작동하지 않을 수 있습니다.

[배송 방법 선택 및 확인 버튼을 보여주는 배송 준비 대화 상자 인터페이스]

**지역별로 주문에 표시되는 물류 채널이 다르며 주문 배송 방법도 다릅니다. 국가와 배송 방법만 구분하십시오.**

---

## 3.1.6 양식 인쇄

현재 모든 양식 교차 목록/내 양식은 현재 지원되지 않습니다. 이 작업을 수행하려면 먼저 Gosipal 구매 목록 API를 먼저 사용해야 합니다.

[다음과 같은 배송 라벨/양식 표시:
- 바코드가 있는 주문 세부 정보
- 다음을 포함한 배송 정보:
  - 택배: J&T Express (MY)
  - 상품 이름/SKU
  - 목적지 주소
  - Shopee 브랜딩
  - 상품 이름, 옵션, 수량, 소계를 포함한 열이 있는 포장 목록 표]

**참고:** 생성된 주문은 주문이 배송(스캔/전송)된 후 주문 상태(주문 = 배송됨)를 반환하는 경우에만 배송됩니다.

---

## 3.1.7 주문

마지막 환경에서 배송을 완료한 후 콘솔 > 테스트 주문 페이지로 들어가서 주문을 조작합니다. 이것을 테스트하십시오.

[여러 열과 작업 버튼이 있는 주문 목록을 보여주는 주문 관리 인터페이스]

---

## 3.1.7.1 "픽업"을 클릭하면 주문 상태가 자동으로 "배송됨"으로 변경됩니다.

테스트 주문 상태가 "배송됨"으로 변경되면 교차 목록에서도 주문 상태가 "배송됨"으로 업데이트됩니다. /fetch/Logisticsorder_detail을 통해 확인할 수 있습니다. "픽업"을 다시 확인하여 주문을 "배송 완료"로 설정할 수도 있습니다.

[업데이트된 상태를 보여주는 유사한 주문 관리 인터페이스]

---

## 3.1.7.2 "배송"을 클릭하면 주문 배송이 완료되면 "수령 확인 중"으로 변경됩니다.

[배송 상태 업데이트를 보여주는 주문 목록 인터페이스]

---

# 개발자 통합 가이드

## 3.1.7.2 주문 배송 완료

**참고:** 데이터 배송 주문 배송이 완료되면 "수령 확인 중"으로 변경됩니다.
주문 결과가 6 또는 9이거나 "배송됨" 상태인 경우 "배송"을 클릭할 수 있습니다.

---

## 3.2 판매자 계정

결제 기능을 활용하기 위해 판매자는 "중국 판매자"를 선택하여 일치하는 계정 및 판매자, 판매자 및 판매자를 만들 수 있습니다.

**참고:** 중국 판매자 온라인의 RM 모든 에이전트는 중국 크로스보더 판매자를 위한 판매 물류로 간주됩니다. 중국의 판매자, FM 모든 에이전트 및 물류 등은 이를 통해 고려됩니다. PS 기본, C2C의 개설 지침 및 소개는 [웹사이트 링크]를 방문하십시오.

### 3.2.1 기본 설정

판매자 계정이 생성된 후 "판매자 센터 열기"를 통해 기본 계정에 로그인하고 기본 계정 및 채권에 대한 권한 부여를 완료하고 각 통화의 환율 변환 및 가격 조정 비율을 설정합니다. 또한 자세한 자습서에 따라 "예금 사업"이 위임된 경우 주문 "학습 센터"[제안된 기본 설정]를 참조하십시오.

**참고:** 결제 모드(OTP)를 사용해야 하는 경우 "[링크]**"를 입력하십시오.

[이것은 그림입니다: 필드와 "추가" 버튼이 있는 양식 인터페이스를 보여주는 스크린샷]

---

[이것은 그림입니다: 다양한 설정과 모달 대화 상자가 있는 "추가" 버튼이 있는 CNBC 플랫폼 인터페이스를 보여주는 스크린샷]

**참고:** CNBC에 로그인—문서/수출 기준 통화 편집 선택—팝업 창 설정 완료—확인을 클릭합니다.

플랫폼 기준 통화 단위 및 시장 환율 설정

[이것은 그림입니다: 환율 구성 인터페이스를 보여주는 스크린샷]

[이것은 그림입니다: 여러 통화 옵션이 있는 환율 설정의 또 다른 스크린샷]

**참고:** 플랫폼 통화의 시장 환율 설정—특정 환율 비율 입력—확인을 클릭—설정 완료 및 CNBC 닫기)

---

글로벌 상품 및 상점 상품 설정:

"판매자 센터" > "글로벌 상품" 페이지로 이동하면 글로벌 상품 및 상점 상품의 가격을 설정하는 팝업 창이 나타납니다.

[이것은 그림입니다: 상품 가격 구성 대화 상자 및 설정을 보여주는 여러 스크린샷]

---

판매 가격 조정 비율, 이벤트 서비스 요금 및 기타 매개변수를 입력합니다. 여러 상점에 적용할 수 있습니다(참고: "상점 설정" - "세금 설정"을 통해 설정해야 함). "다음"을 클릭합니다(업그레이드 기간 포함). 설정이 완료되었습니다.

[이것은 그림입니다: 가격 조정 및 세금 설정 인터페이스를 보여주는 스크린샷]

---

글로벌 상점 FM 배송 창고 설정:

"판매자 센터"로 이동하여 설정->가격 설정->배송 재무 설정을 선택하면 설정이 완료된 후 표준 프로세스에 따라 상품을 배송할 수 있습니다.

[이것은 그림입니다: 다양한 설정 및 옵션을 표시하는 테이블이 있는 배송 및 창고 구성 인터페이스를 보여주는 스크린샷]

---

# 개발자 가이드 - 글로벌 상품

## 개요
글로벌 HIN FM 채널은 설정을 유지합니다.

Shopline 관리 → 상품 → 글로벌 설정에서 글로벌 상품을 관리할 수 있으며 설정이 완료된 후 표준 프로세스에 따라 상품을 배송할 수 있습니다.

---

## 3.2.1 글로벌 상품 추가

### 3.2.1 사이드바에서 글로벌 상품 추가를 클릭하거나 글로벌 상품 페이지에서 글로벌 상품 추가를 클릭합니다.

[글로벌 상품 추가 버튼 위치를 보여주는 탐색]

---

## 3.2.2 글로벌 상품 세부 정보 추가

글로벌 상품 세부 정보 추가(필요에 따라 각 속성을 채우고 선택하십시오)

**참고:** 글로벌 상품 생성

[상품 세부 정보 입력 양식 인터페이스]

---

## 3.2.3 글로벌 상품 및 상점 상품

### 3.2.3 글로벌 상품 추가 및 게시

저장 및 게시를 클릭합니다.

[저장 및 게시 버튼을 보여주는 인터페이스]

---

### 3.2.3 게시 상점 선택

**참고:** 녹색 표시는 싱가포르 사이트에서 선택한 상점을 보여줍니다(지역 제한으로 인해 게시할 수 없는 상점 사이트는 회색으로 표시됩니다). 필요에 따라 적절한 글로벌 상품 게시 상점을 선택하십시오.

[상점 선택 인터페이스]

---

### 3. 상점 상품 정보를 확인하고 확인을 클릭하여 상점에 게시합니다.

[확인 대화 상자 인터페이스]

---

## 4. 상점 상품을 선택하여 출시된 상품을 봅니다.

[상품 보기 인터페이스]

---

**문서 종료**

---

# 테스트 주문 생성 문서

## 1단계: 출시된 상품을 저장할 판매 상품을 선택합니다.

**참고:** 출시된 상품을 저장할 판매 상품을 선택합니다.

---

## 2단계: 업데이트할 수 있는 품목의 재고, 가격 및 기타 속성을 수정하도록 선택합니다(MFRSKU).

**참고:** 업데이트할 수 있는 품목의 재고, 가격 및 기타 속성을 수정하도록 선택합니다(MFRSKU).

---

## 3단계: MFRGU 및 MFRSKU

- **등급 상품(MFRGU):** MFRGU는 글로벌 SKU로, 상위 상품, 가상 상품에만 해당될 수 있으며 여러 지역에 게시할 수 없으며 슈퍼 SKU라고도 합니다.
- **국가 품목(MFRSKU):** 제거 가능한 품목, 구매자에게 보이는 실제 품목

**참고:** MFRGU는 MFRSKU의 기본 정보를 직접 기록합니다. MFRSKU 및 MFRSKU의 강제 상환은 더 이상 여러 지역에서 동일한 상품의 상위 품목 정보를 변경하는 데 필요하지 않습니다. 판매자는 더 이상 여러 지역에서 동일한 상품의 상위 품목 정보를 변경할 필요가 없습니다. MFRGU 품목 유형을 기반으로 MFRSKU의 기본 정보와 시스템은 자동으로 새 MFRGU를 수정하거나 생성할 수 있지만 시스템은 새 MFRSKU를 생성하지 않습니다. MFRSKU는 판매자가 수동으로 작동합니다.

---

## 3.2.3단계: 글로벌 주문 배송

**참고:** 동일한 ID(로컬 상점)에 대한 다른 상점의 주문의 경우 1.1.3 테스트 주문 생성 및 3.1.4 생성된 주문 보기를 사용할 수 있습니다.

---

## 테스트 주문 생성

### 상점
**22512640 (크로스보더 - MY)**

### 품목
1. **80196051** (샌드위치 찾아보기, 테스트 글로벌 품목 01)
   - 수량: 1
   - 작업: + ×

2. **80196473** (샌드위치 테스트 글로벌 품목 02 찾아보기)
   - 수량: 1
   - 작업: + ×

### 품목 선택 (2/3)

### 배송 옵션
**27002**

**버튼:** 취소 | 생성

---

## 4단계: 해당 상점으로 전환한 후 생성된 주문 번호를 볼 수 있습니다.

**중요:** 주문을 생성한 후 Shopee의 "중국 창고에 대한 테스트 계정 업데이트"에서 귀하의 계정을 클릭하고 "판매자 센터 로그인" 버튼을 누릅니다.

**또한 참고:** 내 주문에 성공적으로 로그인한 후(오른쪽에서 생성한 주문에 대한 주문을 선택해야 함)

---

## 5단계: 해당 상점으로 전환한 후 해당 주문을 볼 수 있습니다.

**참고:** 주문 상태 흐름은 "배송 준비 중"이며 "배송 준비"로 작동할 수 있습니다. 다른 배송의 경우 다음 "글로벌 상품 배송 창고 설정의 순간"을 선택하십시오.

---

## 6단계: 주문을 선택하고 "배송 준비"를 클릭합니다.

---

## 7단계: 주문을 선택하고 "배송 준비"를 클릭하고 드롭오프 방법을 선택하고 "확인"을 클릭하여 주문을 배송합니다.

**참고:** (CB는 현재 드롭오프 방법을 일시 중단할 수 없습니다)

---

## 최종 참고 사항

검사 후 주문은 "개발자가 적절한 현재 ID를 얻을 수 있는 AWS 역할을 수행하는 주문을 생성합니다." 이는 "키"를 얻어 얻어야 합니다.

---

*원래 구조, 제목, 단계, 참고 사항 및 콘텐츠를 유지하면서 개발자 가이드 스크린샷에서 추출한 모든 텍스트.*

---

# 개발자 가이드 - 주문 이행 및 테스트

## 배송 후

f. 배송 후 주문은 TN과 AWB를 생성하며 개발자는 적절하게 얻을 수 있습니다(현재 AWB를 얻기 위해 API만 지원됨).

### 내 주문 인터페이스

**주문 관리 탭:**
- 모두
- 미결제
- 구매 예정 (2)
- 배송 중
- 완료됨
- 취소
- 반품/환불
- 배송 실패

**주문 세부 정보:**
- 주문 ID
- 주문 ID 입력
- 배송/채널: 모든 채널
- 작업: 적용, 재설정

**12개 주문**

**상품 정보:**
- 상품
- 주문
```

### 배송 옵션
**27002**

**버튼:** 취소 | 생성

---

## 4단계: 해당 쇼핑몰로 전환 후 생성된 주문 번호를 확인할 수 있습니다.

**중요:** 주문 생성 후 Shopee의 "Test Account Update for China Warehouse" 아래에 있는 계정을 클릭하고 "Login seller Center" 버튼을 누르십시오.

**참고:** My Order에 성공적으로 로그인한 후 (생성한 주문에 대한 주문을 선택해야 함)

---

## 5단계: 해당 쇼핑몰로 전환 후 해당 주문을 확인할 수 있습니다.

**참고:** 주문 상태 흐름이 "To Ship"인 경우 "Arrange Shipment"로 작동할 수 있습니다. 다른 배송의 경우 다음을 선택하십시오. "For Me Moment of the global goods Shipment Warehouse Setups is Complete"

---

## 6단계: 주문을 선택하고 "Arrange Shipment"를 클릭합니다.

---

## 7단계: 주문을 선택하고 "Arrange Shipment"를 클릭하고 Drop-off 방법을 선택한 다음 "Confirm"을 클릭하여 주문을 배송합니다.

**참고:** (현재 CB는 suspend Drop-off 방법을 클릭할 수 없습니다.)

---

## 최종 참고 사항

검사 후 주문은 "AWS 역할을 수행하기 위해 수행해야 할 작업으로 생성되며, 개발자는 "키"를 획득하여 적절한 현재 ID를 얻을 수 있습니다."

---

# 개발자 가이드 - 주문 처리 및 테스트

## 배송 후

f. 배송 후 주문은 TN과 AWB를 생성하며, 개발자는 적절하게 획득할 수 있습니다 (현재 AWB 획득에는 API만 지원됨).

### 내 주문 인터페이스

**주문 관리 탭:**
- 전체
- 미결제
- 구매 예정 (2)
- 배송 중
- 완료됨
- 취소
- 반품/환불
- 배송 실패

**주문 상세 정보:**
- 주문 ID
- 주문 ID 입력
- 배송/채널: 모든 채널
- 액션: 적용, 초기화

**12개의 주문**

**상품 정보:**
- 상품
- 총 주문 금액
- 상태
- Counddons
- 배송 채널
- 액션

**주문 예시:**
- local_man.my 🔥
- 주문 ID: 230708004254745
- Shoppe Seedbox Text Global Item 01
- 총 주문 금액: RM146.80 (2x Line Partner)
- 상태: To Ship (READY for canal is south in engineer.)
- 액션:
  - 📋 배송 상세 정보 보기
  - 인쇄/상세 정보
  - 🖨️ 운송장 인쇄

---

## 주문 처리 프로세스

g. 마지막으로 테스트 주문 페이지로 돌아가서 "Pickup" 및 "Deliver"를 클릭하여 시뮬레이션을 완료하여 후속 처리를 완료합니다.

### 상태 흐름:
- "Pickup"을 클릭하면 주문이 **"SHIPPED"** 상태로 변경됩니다.
- "Deliver"를 클릭하면 주문이 **"TO_CONFIRM_RECEIVE"** 상태로 변경됩니다.

### 테스트 주문 콘솔

**탐색 메뉴:**
- 앱 목록
- 푸시 메커니즘
- 테스트
  - 테스트 계정 샌드박스 xx
  - 테스트 주문
  - API 테스트 도구
  - 결제 API
  - 푸시 로그
  - API 호출 통계

**테스트 주문 검색:**
- 주문 SN 검색
- 주문 SN: 23070804254740

**주문 목록 테이블:**

| 주문 SN | 상품 ID | 상태 | 업데이트 시간 | 쇼핑몰 ID | 액션 |
|----------|---------|--------|-------------|---------|--------|
| 23070804254740 | 80186024 s_1_80186039 s_1 | PROCESSED | 09-07-2025 10:54:40 | 231323045 (Casa Shopin - MY) | 픽업 상세 정보 삭제 |
| | 80186024 s_1_80186039 s_1 | PROCESSED | 09-07-2025 10:54:40 | 231940165 (Luna - MY) | 픽업 상세 정보 삭제 |
| | 80186024 s_1_80186039 s_1 | PROCESSED | 08-07-2025 10:54:40 | 231940165 (Luna - MY) | 픽업 상세 정보 삭제 |
| | 80186057 s_1_80186088 s_1 | INVALID | 29-06-2025 09:15:45 | 231632388 (Luna - MY) | 픽업 상세 정보 삭제 |
| | 80186024 s_1_80186039 s_1 | READY_TO_SHIP | 25-06-2025 15:33:23 | 231324245 (Luna - SS) | 픽업 상세 정보 삭제 |
| | 80186024 s_1_80186039 s_1 | READY_TO_SHIP | 25-06-2025 16:47:17 | 231324245 (Luna - SS) | 픽업 상세 정보 삭제 |
| | 80186024 s_1_80186039 s_1 | SHIPPED | 25-06-2025 16:59:26 | 231424404 (Luna - SS) | 픽업 배송 삭제 |

---

h. 처리가 완료되면 주문은 사용자가 조작할 필요가 없으며, 일정 시간이 지나면 "COMPLETED" 상태로 변경되어 처리가 완료됩니다.

---

## 3.2.5 CNSC 테스트 가능 인터페이스

CNSC가 중점을 두는 인터페이스는 상품 관리와 관련이 있으며, 테스트 가능한 인터페이스에는 **Merchant GlobalProduct** 및 **MediaSpace**가 포함됩니다. 모든 인터페이스가 다운로드되며 다른 테스트는 일반 상점과 다르지 않습니다.

---

## 3.3 푸시 메커니즘

콘솔-> **푸시 메커니즘 페이지**를 선택하고 상태가 **Developing**인 APP을 선택한 다음 푸시 설정을 입력합니다.

### 푸시 메커니즘 구성

**탐색:**
- 푸시 메커니즘 > 푸시 설정

**테스트 푸시 구성 설정:**

- 콜백 URL 설정
  - https://your.domain.com/

- 테스트 푸시 키 설정

**이벤트 수신기:**

**푸시 테스트 목록:**

| 상품 푸시 | 이벤트 | 액션 |
|--------------|-------|--------|
| 상품 푸시 | shopee_item_push | 0 | 테스트 데이터 입력 |
| 상품 푸시 | shopee_notify_markup_push | 0 | 테스트 데이터 입력 |
| 주문 푸시 | order_status_push | 0 | 테스트 데이터 입력 |
| 주문 푸시 | order_address_push | 0 | 테스트 데이터 입력 |
| 주문 푸시 | wholesale_favorited_order_push | 12 | 테스트 데이터 입력 |
| 선택적 푸시 | promotion_main | 0 | 테스트 데이터 입력 |
| 마케팅 푸시 | promotion_update_push | 0 | 테스트 데이터 입력 |
| | shop_authorization.push | 0 | 테스트 데이터 입력 |
| Shopin 푸시 | shop_authorization_cancelled_push | 0 | 테스트 데이터 입력 |
| 상태 푸시 | | 13 | 테스트 데이터 입력 |

---

## 중요 참고 사항

샌드박스 환경의 푸시 메커니즘은 프로덕션 환경과 다릅니다. 더 이상 푸시를 트리거하기 위해 관련 작업을 사용할 필요가 없습니다. 테스트 콜백 URL을 입력하고 "확인 및 저장"을 클릭하여 확인을 완료합니다. 해당 푸시 메커니즘 후 "푸시 테스트 데이터"를 클릭하기만 하면 테스트 데이터를 받을 수 있습니다.

## 사용 사례

1. 배포 전 API 통합 테스트
2. 주문 처리 워크플로우 시뮬레이션
3. 테스트 환경에서 상품 목록 관리
4. API 액세스를 위한 테스트 계정 권한 부여
5. 푸시 서비스 통합 테스트

## 관련 API

- Product APIs
- Global Product APIs
- Media Resources APIs
- Order APIs
- Logistics APIs
- First Mile APIs
- Shop APIs
- Merchant APIs

---

## 원문 (English)

### Summary

This guide introduces the Sandbox V2 testing environment provided by Open Platform for developers. It covers creating test accounts, test orders, and authorizing test accounts, enabling developers to test API functions in an isolated environment.

### Content

# Sandbox Testing V2

**Getting Started >** Sandbox Testing V2

## Overview

The Sandbox is a isolated testing environment provided by Open Platform to developers. It provides various types of test accounts and data. Developers can complete testing of most API functions in the sandbox environment. The sandbox provides features for some functions and interfaces, but only covers most scenarios, such as product management, order processing, etc.

**Sandbox V2 support range**

| Portal | Features | Network |
|--------|----------|---------|
| Console | Create test shop account | |
| | Order | Create test order |
| | | Push test data |
| Seller Center | Product | Create and manage the global SKU |
| | Order | Publish and manage test orders |
| | Order | Order |
| | | Shop Order | Printing of receipts is not supported at the moment, please use Open API to print |
| Open API | Product | All APIs | Can through the API Test Tools in the Console, or call it yourself using the domain name https://openplatform.sandbox.test |
| | Global Product | All APIs | Status always on: ON, please use https://openplatform.sandbox.test (single-sku) |
| | Media Resources | All APIs | |
| | Order | All APIs | |
| | Logistics | All APIs | |
| | First Mile | All APIs | |
| | Shop | All APIs | |
| | Merchant | All APIs | |
| Push Service | Supports receiving some push test data, see details 2.3 | |

> If you require additional sandbox support features, please contact customer service with specific use-case and requirements in your environment.

This document mainly introduces the process of creating test accounts, test orders and shipping the Sandbox test environment.

---

## 1. Create a test account

**Quick Console:** Select Test Account-Sandbox v2, on at 6 **Create a test account**

### 1.1 Local and cross-border test stores

Local and cross-border test stores differ in many aspects such as item categories, available transportation channels, and payment methods, etc. Please select a test store corresponding to the service market.

[Screenshot showing account creation interface with "Create Test Account" button highlighted]

### 1.2 Merchant

For storing in China Seller Center (CNSC), developers can choose "Merchant" to create a test master account and bind it.

[Screenshot showing merchant account creation interface with navigation arrows]

---

## 2. Authorize the test account to the test Partner_id

### 2.1 Preliminary steps:

1. Create an

[Screenshot showing authorization flow with multiple panels and arrows]

2. Create Sandbox Account

[Screenshot showing sandbox account creation interface]

> You can choose the account type and create a shop or merchant type test account

[Screenshot showing account type selection dialog]

---

### 2.2 Shop authorization document

**Authorization Notes:**

- For the authorization, the shop has the button for the specific authorization entrance (Can be referencable from https://open.shopee.com/documents/v2/v2.push.get_config
or follow the steps below.

[Screenshot showing authorization interface with multiple product panels and arrows]

---

**Note:** This extraction maintains the document structure, headings, tables, steps, and notes as they appear in the developer guide screenshot. All text has been preserved in English as requested.

---

# Developer Integration Guide

## Authorization Flow

### For the test/Partner_id, set the figure below for the specific authorization process Call be referenced Authorization related documents, or follow the steps below

---

## Step 1: Authorization Flow Setup

**Use the authorization link and fill in the corresponding `partner_id`:**

```
https://open.sandbox.best.statestreet.com/oauth2/
```

```
https://web.sandbox.test.statestreet.com/page?id=wfm*oauth*authorize&response_type=code
```

---

## Authorization Page

### Login to Sandbox Best Review
**User Sandbox Best Review**

[AUTHORIZE button]

**Note:** Fixed authorization URL for sandbox environment including https://account.sandbox.best.statestreet.com (please confirm Merchant-Accountable a login url if live account. Otherwise, the error 'Account/Password Verification Failed' will be reported.

---

## Step 2: Fill in the corresponding Sandbox Account and log in

### Authorization
[Checkmark icon indicating step 2]

- Partner App consent to payment and to use account
- Payment information of transaction cannot be modified
- App can read transaction information
- Partner App can query account balance
- Payment information of transaction or payment cannot be
- Transaction information of completed payment can be queried

---

## Step 3: Click Authorization to jump to the success page

### Authorization
[Checkmark icon]

**Note:** After authorization, return to the application page

**Authorization successful!**

**Redirect URL:**
- https://

[Back] [Confirm]

---

## 2.3 Merchant authorization document

For the test/Partner_id, set the figure below for the specific authorization process Call be referenced Authorization related documents, or follow the steps below

---

## Authorization Steps (Merchant)

**Use the authorization link and fill in the corresponding `partner_id`:**

```
https://open.sandbox.best.statestreet.com/oauth2/
```

```
https://web.sandbox.test.statestreet.com/page?id=wfm*oauth*authorize&response_type=code
```

Click "trust account" for merchant authorization.

---

## Login to Sandbox Best Review
**Sandbox Best Review**

[AUTHORIZE button - highlighted]

**Note:** Fixed authorization URL for sandbox environment including https://account.sandbox.best.statestreet.com (please confirm Merchant-Accountable a login url if live account. Otherwise, the error 'Account/Password Verification Failed' will be reported.

The verification code is "123456"

---

## Verification

**Enter verification code**

[CONFIRM button]

---

## Step 4: After logging in, check the status that require authorization

### Success Page Platform

**Authorization/API/Confirm**

[Checkmark icon]

**Authorization/Confirm - Completed - Success - Platform**

[Multiple authorization status items listed]

---

## Step 5: Click Authorization to jump to the success page

### Authorization/API/Confirm

[Checkmark icon]

**Authorization successful!**

**Redirect URL:**
- sandbox

[Back] [Confirm - highlighted]

---

**Note:** The extracted content maintains the hierarchical structure, code examples, steps, notes, and warnings from the developer guide screenshots. All text has been preserved in English as it appeared in the original document.

---

# 3. Sandbox testing process

## 3.1 Shop account

### 3.1.1 Log in to the Seller Center

Click Console->Select**Your Account-Sandbox** page->click Login Seller Center on the right side of the created test store to enter the Seller Center page.

---

*The Seller Center page is displayed as shown below:*

---

## 3.1.2 Create test products

### 3.1.2.1 You can choose to create products from SellerCenter or through Open API.

The following two test products created through Seller Center:

---

*Fill in all required fields, then select Base and Publish.*

---

### 3.1.2.2 After successfully creating the product, the seller center can view the Products (My Products).

---

## 3.1.3 Create test order

### 3.1.3.1 Click Simulation tools-Buy Test Order page, click "Create Test Order" to view the order status you want to test now.

---

Click the "Shop" drop-down box and select the shop you want to create.

---

Once "Select form" (order status), select the Box that needs to be created, and click "Confirm"

---

### 3.1.3.2 Click the "Shipping Central" drop-down box to select the fulfillment channel, and click "Confirm" to complete the order creation.

---

*Note: Different order statuses may have different fulfillment channel options available.*

---

# Order Management Guide

## 3.1.2 "test order creation" completed

[This section appears to show an order list interface with multiple columns including order details, status, and actions]

---

## 3.1.4 View created orders

Enter the seller centre and click "My Order" to view the created orders.

**Note:** After creating an order on the Gossipay page, you need to wait for about 5 minutes before proceeding to the next step.

[Order list interface showing multiple orders with columns for order information, products, payment, status, and operations]

---

## 3.1.5 Shipping

Click "Arrange Shipment" and select the shipping method (pickup/dropoff). A tracking number will be automatically generated. In this test, the order status is **PROCESSED**.

**Note:** Please cancel the order with "To Ship" status first. The "jet" tag may not be able to operate.

[Shipping arrangement dialog interface showing shipping method selection and confirmation buttons]

**The logistics channels displayed for orders in different regions are different, and the order delivery methods are also different, only countries, only delivery methods, remember to distinguish.**

---

## 3.1.6 Print the form

Currently, all the forms Crosslisting/my of the form is not currently supported This operation. If you need to print a form, first you need to first Gosipal Purchase list API

[Shipping label/form showing:
- Order details with barcode
- Shipping information including:
  - COURIER: J&T Express (MY)
  - PRODUCT NAME/SKU
  - Destination address
  - Shopee branding
  - Packing List table with columns for Product Name, Option, Qty, Subtotal]

**Note:** The created orders will only be shipped after your order is shipped (scanned/sent) and returns the order status (order = SHIPPED)

---

## 3.1.7 Order

After completing the shipment in the last environment, you can enter the Console > Test Order page to operate the order. Test this.

[Order management interface showing order list with multiple columns and operation buttons]

---

## 3.1.7.1 Click "Pickup" and the order status will automatically change to "SHIPPED"

After the test order status changes to "SHIPPED", the order status will also be updated to "SHIPPED" on Crosslisting. You can check it via /fetch/Logisticsorder_detail. You can also check "Pickup" again to set the order to "DELIVERED"

[Similar order management interface showing updated status]

---

## 3.1.7.2 Click "Deliver When the order delivery is completed, it will change to "TO_CONFIRM_RECEIVE"

[Order list interface showing delivery status updates]

---

# Developer Integration Guide

## 3.1.7.2 Order Delivery Completion

**Note:** Data Deliver When the order delivery is completed, it will change to "TO_CONFIRM_RECEIVE". 
After the order result is 6 or 9 or "SHIPPED" status where you can click "Deliver".

---

## 3.2 Merchant Account

To utilize the payment functions, merchants can choose "China Merchant" to create a matching account and sellers, merchants and sellers.

**Note:** RM All agents of China Seller Online is a sales logistics considered for Chinese cross-border sellers. Sellers in China, FM All agents and logistics etc. in considered to through it. PS base, opening instructions and introduction in the C2C, please visit [website link].

### 3.2.1 Basic settings

After the Merchant account is created, log in to the main account through "open Seller Center", complete the authorization of the main account and bonds, and set the exchange rate conversion and price adjustment percentage of each currency. Additionally, if the "deposit business" is delegated according to detailed tutorials, please refer to the order "Learning Center" [suggested basic settings].

**Note:** If you need to use the settlement mode (OTP), please enter "[link]**"

[THIS IS FIGURE: Screenshot showing a form interface with fields and an "ADD" button]

---

[THIS IS FIGURE: Screenshot showing a CNBC platform interface with various settings and an "Add" button with a modal dialog]

**Note:** Log in to CNBC—Select the document/export base currency edit—Complete pop-up window settings—Click OK

Set the platform base currency unit and market exchange rate

[THIS IS FIGURE: Screenshot showing exchange rate configuration interface]

[THIS IS FIGURE: Another screenshot of exchange rate settings with multiple currency options]

**Note:** Set the market exchange rate of the platform currency—fill in the specific exchange rate ratio—Click OK—complete the setting and close CNBC)

---

Global product and store product settings:

Go to "Seller Center" > "Global Products" page, and it a pop-up window will pop up to set the prices of global products and store products.

[THIS IS FIGURE: Multiple screenshots showing product pricing configuration dialogs and settings]

---

Fill in the sale price adjustment ratio, event service rate and other parameters. It can be applied to multiple stores (note: Make sure to set through the "store Setup" - "Tax Setup" in click "Next" (upgrading period included). Its complete the settings.

[THIS IS FIGURE: Screenshot showing price adjustment and tax setup interface with a data table]

---

Global store FM shipment warehouse settings:

Go to "Seller Center", select Settings->Price Settings->Shipping Financial Settings, and then you can ship the product according to the standard process after the settings are completed.

[THIS IS FIGURE: Screenshot showing shipping and warehouse configuration interface with a table displaying various settings and options]

---

# Developer Guide - Global Products

## Overview
Global HIN FM channel maintains settings.

You can manage Global Products from the Shopline Admin → Products → Global Settings, and then you can ship your products according to the standard process after the settings are completed.

---

## 3.2.1 Add global products

### 3.2.1 Click Add Global Product in the sidebar or click Add Global Product on the Global Product Page

[Navigation showing Add Global Product button locations]

---

## 3.2.2 Add global product details

Add global product details (please fill in and select each attribute according to your needs)

**Note:** Create global product(s)

[Form interface showing product details entry]

---

## 3.2.3 Global products and store products

### 3.2.3 Add and publish global products

Click Save and Publish

[Interface showing Save and Publish buttons]

---

### 3.2.3 Select publishing store

**Note:** The green bellow shows a store selected from the Singapore site (due to region restrictions, shop sites that cannot be published will be grayed out). Please select the appropriate global product publishing store as needed.

[Store selection interface]

---

### 3. Confirm the store product information and click Confirm to publish it in a store

[Confirmation dialog interface]

---

## 4. Select the store product to view the released product

[Product viewing interface]

---

**End of Documentation**

---

# Create Test Order Documentation

## Step 1: Select the sales product to save the released product

**Note:** Select the sales product to save the released product.

---

## Step 2: Select to modify the inventory, price and other attributes of the item that can be updated (MFRSKU)

**Note:** Select to modify the inventory, price and other attributes of the item that can be updated (MFRSKU)

---

## Step 3: MFRGU and MFRSKU

- **Grade product (MFRGU):** MFRGU is Global sku, which can only correspond to a parent product, a virtual product, which can not be published to multiple regional, also known as Super Sku
- **Country item (MFRSKU):** Removable item, Real item visible to buyers

**Note:** MFRGU will directly record the basic information of MFRSKU. The mandatory redemption of MFRSKU and MFRSKU is no longer required to change parent-item information of the same product in multiple regions in the item. Sellers no longer need to change parent-item information of the same product in multiple regional. Based on the MFRGU item type, the basic information of the MFRSKU and the system can automatically modify or create a new MFRGU, but note the system will not create a new MFRSKU. MFRSKU is operated manually by the seller.

---

## Step 3.2.3: Shipment of global orders

**Note:** For orders from different shops to the same id (Local Shop), you can use to 1.1.3 Creating Test Orders and 3.1.4 Viewing Created Orders.

---

## Create Test Order

### Shop
**22512640 (Cross-Border - MY)**

### Item(s)
1. **80196051** (Browse Sandwich, Test Global Item 01)
   - Quantity: 1
   - Actions: + ×

2. **80196473** (Browse Sandwich Test Global Item 02)
   - Quantity: 1
   - Actions: + ×

### Select Item (2/3)

### Shipping Option
**27002**

**Buttons:** Cancel | Create

---

## Step 4: After switching to the corresponding shop, you can see the order number created

**Important:** After creating the order, click on your account under "Test Account Update for China Warehouse" on Shopee and press the button "Login seller Center"

**Also note:** After successfully log-in to My Order (right must select the order for the order you created)

---

## Step 5: After switching to the corresponding shop, you can see the corresponding order

**Note:** The order status flow is "To Ship" can be operated to "Arrange Shipment". For other shipment please select next "For Me Moment of the global goods Shipment Warehouse Setups is Complete"

---

## Step 6: Select your order and click "Arrange Shipment"

---

## Step 7: Select your order, click "Arrange Shipment", select the Drop-off method, and click "Confirm" to ship your order

**Note:** (CB cannot click suspend Drop-off method at this time)

---

## Final Note

After inspection, the order generates a "to do as to act as AWS, which the developer can obtain an appropriate currently id" which is supposed to be obtained by obtaining the "key".

---

*All text extracted from the developer guide screenshot maintaining original structure, headings, steps, notes, and content.*

---

# Developer Guide - Order Fulfillment and Testing

## After Shipment

f. After shipment, the order generates a TN as well as an AWB, which the developer can obtain as appropriate (currently only the API is supported for obtaining the AWB).

### My Orders Interface

**Order Management Tabs:**
- All
- Unpaid
- To Buy (2)
- Shipping
- Completed
- Cancellation
- Return/Refund
- Failed Delivery

**Order Details:**
- Order ID
- Input order ID
- Shipping/Channel: All Channels
- Actions: Apply, Reset

**12 Orders**

**Product Information:**
- Product(s)
- Order Total
- Status
- Counddons
- Shipping Channel
- Actions

**Example Order:**
- local_man.my 🔥
- Order ID: 230708004254745
- Shoppe Seedbox Text Global Item 01
- Order Total: RM146.80 (2x Line Partner)
- Status: To Ship (READY for canal is south in engineer.)
- Actions: 
  - 📋 View Shipping Details
  - Print/Details
  - 🖨️ Print Waybill

---

## Order Fulfillment Process

g. Finally, return to the Test Order page and complete the simulation by clicking "Pickup" and "Deliver" to complete the subsequent fulfillment.

### Status Flow:
- After clicking "Pickup", the order will flow to **"SHIPPED"** status.
- After clicking "Deliver", the order will flow to **"TO_CONFIRM_RECEIVE"** status.

### Test Order Console

**Navigation Menu:**
- App List
- Push Mechanism
- Test
  - Test Account Sandbox xx
  - Test Order
  - API Test Tool
  - Payment API
  - Push Log
  - API Calls Statistics

**Test Order Search:**
- Search Order SN
- Order SN: 23070804254740

**Order List Table:**

| Order SN | Item ID | Status | Update Time | Shop ID | Action |
|----------|---------|--------|-------------|---------|--------|
| 23070804254740 | 80186024 s_1_80186039 s_1 | PROCESSED | 09-07-2025 10:54:40 | 231323045 (Casa Shopin - MY) | Pickup Details Delete |
| | 80186024 s_1_80186039 s_1 | PROCESSED | 09-07-2025 10:54:40 | 231940165 (Luna - MY) | Pickup Details Delete |
| | 80186024 s_1_80186039 s_1 | PROCESSED | 08-07-2025 10:54:40 | 231940165 (Luna - MY) | Pickup Details Delete |
| | 80186057 s_1_80186088 s_1 | INVALID | 29-06-2025 09:15:45 | 231632388 (Luna - MY) | Pickup Details Delete |
| | 80186024 s_1_80186039 s_1 | READY_TO_SHIP | 25-06-2025 15:33:23 | 231324245 (Luna - SS) | Pickup Details Delete |
| | 80186024 s_1_80186039 s_1 | READY_TO_SHIP | 25-06-2025 16:47:17 | 231324245 (Luna - SS) | Pickup Details Delete |
| | 80186024 s_1_80186039 s_1 | SHIPPED | 25-06-2025 16:59:26 | 231424404 (Luna - SS) | Pickup Deliver Delete |

---

h. After the fulfillment is completed, the order does not need to be operated by you, and will flow to the status of "COMPLETED" after a certain period of time, so that the fulfillment is completed.

---

## 3.2.5 CNSC testable interface

The interfaces that CNSC focuses on are related to commodity management, and testable interfaces include **Merchant GlobalProduct** and **MediaSpace**. All interfaces are downloaded, and other tests are no different from ordinary stores.

---

## 3.3 Push Mechanism

Click Console-> select the **Push Mechaniam page**, select the APP with the status of **Developing**, and enter Set Push.

### Push Mechanism Configuration

**Navigation:**
- Push Mechanism > Set Push

**Set Test Push Configuration:**

- Set Call Back URL
  - https://your.domain.com/

- Set Test Push Key

**Event Receivers:**

**Push Test List:**

| Product Push | Event | Action |
|--------------|-------|--------|
| Product Push | shopee_item_push | 0 | Enter Test Data |
| Product Push | shopee_notify_markup_push | 0 | Enter Test Data |
| Order Push | order_status_push | 0 | Enter Test Data |
| Order Push | order_address_push | 0 | Enter Test Data |
| Order Push | wholesale_favorited_order_push | 12 | Enter Test Data |
| Selective Push | promotion_main | 0 | Enter Test Data |
| Marketing Push | promotion_update_push | 0 | Enter Test Data |
| | shop_authorization.push | 0 | Enter Test Data |
| Shopin Push | shop_authorization_cancelled_push | 0 | Enter Test Data |
| Status Push | | 13 | Enter Test Data |

---

## Important Notes

The Push Mechanism in the Sandbox environment is different from the production environment. It is no longer necessary to use related operations to trigger the push. Enter the Test Call Back URL and click "Verify and Save" to complete the verification. Just click "Push Test Data" after the corresponding Push Mechaniam to receive the test data.

---

**문서 ID**: developer-guide.644
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/644
**처리 완료**: 2025-10-16T08:18:47
