# Ads Open API 사용 규칙 및 안내

**카테고리**: 통합
**난이도**: medium
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T08:59:08

## 개요

본 가이드는 Ads Open API 사용에 대한 규칙 및 지침을 제공합니다. 데이터 보호, 예산 관리, 키워드 검색, 캠페인 데이터, 광고 생성 및 ROAS 최적화와 같은 주제를 다룹니다. 또한 문제 해결 팁 및 API 권한 정보도 포함되어 있습니다.

## 주요 키워드

- Ads Open API
- 광고
- 예산 관리
- 키워드
- 캠페인 데이터
- 상품 광고
- ROAS
- 입찰 방식

## 본문

# Ads Open API 사용 규칙

ads/storefront_api를 사용할 때, 판매자와 구매자 모두에게 공정하고 평등한 마켓플레이스를 제공하기 위해 Shopee가 정한 다음 규칙을 준수해야 합니다.

1. 광고 데이터를 처리할 때 플랫폼의 **데이터 보호 정책**을 준수해야 합니다.
2. S 내의 모든 항목에 대해 Shopee의 공식 협력 CPS 제한은 플랫폼 사용자와 함께 사용됩니다. 제휴 마케팅 용도로 신규 사용자/구매자에게 사용할 수 있는 경우 지불 금액은 합리적이어야 하며 다른 용도로 사용할 수 없습니다.
3. 데이터는 상점 운영 및 공식 Shopee 협력 프로젝트 이외의 다른 목적으로 사용될 수 있으므로 이 API를 통해 데이터에 액세스할 때 플랫폼 판매자의 동의를 얻어야 하며 정보를 공개적으로 사용할 수 있어야 합니다.
4. 어떠한 방식으로도 귀하, 당사의 이름 및 링크를 촉매제로 사용해서는 안 되며, Shopee로부터 경고 이메일을 받게 됩니다. 제때 수정하지 못하면 권한을 잃을 수 있습니다.

---

# Ads Open API 안내

## 광고 예산 잔액 및 예산 정보 검색

- **지갑 잔액**으로 광고 계정을 충전할 수 있으며 완료되면 이메일로 알려드립니다.
- 사용하기 전에 이 API를 통해 잔액 상태를 변환하여 광고를 실행하기에 충분한 잔액인지 확인하십시오.
- 잔액이 부족하면 일일 입찰가, 가져오기 및 가격 최적화 기능 활성화 여부와 같은 상점 광고 설정으로 돌아가야 합니다.

### MyLite/Shopee-Budgetbalada-wallet-top-up

[예산 잔액 및 "잔액 충전" 버튼이 강조 표시된 충전 옵션을 보여주는 인터페이스 요소]

**캠페인, 단일** - 잔액 알림 활성화

[잔액 알림 설정이 있는 "Shopee Ads Shop Level Settings" 대화 상자를 보여주는 스크린샷]

---

## 3. 추천 키워드 및 제품 데이터 검색

<CT ads_get_recommended_keywords_by_item> API를 사용하여 항목에 대한 추천 키워드 목록과 선택적 검색 키워드를 검색할 수 있습니다. 검색 정보는 키워드 인기 점수, 검색량 및 제안 입찰가도 볼 수 있습니다.

- <CT ads_get_recommended_products> 광고 캠페인에 추가할 수 있는 항목 목록입니다. 또한 항목이 추천되는지 여부를 정의하고 항목 홍보에 실패하면 추천 SDG 상점 수준 광고 특정 캠페인 목록이 표시됩니다.

## 제품 선택

[제품, 상태, 조회수, 제품 이름 열이 있는 제품 선택 인터페이스를 보여주는 표]

---

## 키워드 추가

[키워드, 노출수(최근 30일), 인기 등 열이 있는 키워드 데이터를 보여주는 표]

---

## 4. 캠페인 수준 데이터 검색

<CT ads_get_campaign_list> API를 사용하여 특정 상점과 연결된 모든 캠페인 목록을 검색합니다.

API를 사용하여 모든 캠페인 목록을 가져옵니다.

- <CT ads_get_product_campaign_data> 및 <CT ads_get_product_campaign_daily_performance> 이 API를 사용하여 특정 캠페인에 대한 자세한 설정을 검색하고 <CT ads_get_product_keyword_performance> 및 <CT ads_get_product_campaign_daily_performance> 이러한 API를 사용하여 광고 지출, 노출수, 클릭수, 전환수 등과 같은 캠페인 수준의 성과 데이터를 얻습니다.

### API_CREATE_1

[노출수, 클릭수, 비용 등과 같은 지표를 보여주는 성과 데이터 표]

### 전환수

[전환 추적 인터페이스]

### 성과 차트

[노출수, 클릭수, 직접 주문 등과 같은 데이터 포인트가 있는 시간 경과에 따른 성과 지표를 보여주는 차트]

---

## 5. 광고 제품 광고 만들기

<CT ads_create_product_ads> API를 사용하여 광고를 만듭니다. (상점 생성 검색)

---

## 제품

**예산 광고 스토어프론트**

최고의 제품 자동 선택

- 상점의 스토어프론트 광고에 제품 추가
- 판매자 센터에서 광고 검색

---

## 6. 일일 최대 및 수동 제품 광고 만들기

- <CT ads_create_manual_product_ads>

일일 최대 광고를 만들려면 bidding_method:manual을 전달합니다.

수동 광고를 만들려면 bidding_method:manual을 전달합니다.

### 요약

[예산 설정을 보여주는 인터페이스]

### 제품 수동 관리

[수동 제품 관리 옵션을 보여주는 인터페이스]

### 입찰

["CPC 모델로 설정" 및 "수동으로 개별적으로" 라디오 버튼이 있는 입찰 옵션을 보여주는 인터페이스]

---

## 기본 설정

수동 광고를 만들 때 bidding_method:manual을 사용하는 경우 <CT ads_set_environment_item_list> 및 <CT ads_set_recommended_keywords_list>를 사용하여 광고에 대한 추천 제품 및 키워드를 얻을 수 있습니다. 그런 다음 예산 및 광고 티어 권장 사항 데이터를 입력하여 생성 API에 전달합니다.

---

## 광고 예산 및 ROAs 검색

- ROAs를 사용하여 광고 예산 검색
- <CT ads_get_recommended_roi_target> CPC Max 광고를 만들 때 ROAs(광고 지출 수익) 값을 가져옵니다.

---

## 기본 설정

[텍스트가 있는 경고 상자: "❗ CPC 모델로 설정 - 더 나은 노출을 위해 키워드 입찰가를 자동으로 조정합니다."]

**알림: CPC Max 광고를 만들 때 ROAs 목표 값을 API에 전달하십시오(광고 지출 수익).**

---

# 7. 광고 예산 및 ROAS 검색

## API 엔드포인트

- **v2.ads.get_product_recommended_roi_target**: GMV Max 광고를 만들 때 ROAS(광고 지출 수익) 값을 가져옵니다.

- **v2.ads.get_create_product_ad_budget_suggestion**: 제품 광고를 만들기 위한 예산 제안을 검색합니다.

---

## 기본 설정

### 예산
- **무제한**
- **일일 예산 설정**: Rp 0
  - 제안: Rp7,046

### 기간
- **종료 날짜 없음**
- **시작/종료 날짜 설정**: 📅 21/04 - 21/05 (GMT+7)

---

## 입찰

### 입찰 방법

#### GMV Max
Shopee가 시간이 지남에 따라 최적의 성과를 위해 쇼핑객에게 광고를 전달하는 가장 좋은 방법을 찾도록 하십시오.

**ROAS 목표 설정**

**GMV Max 자동 입찰**
Shopee는 건전한 ROAS를 유지하면서 일일 예산 내에서 증가하거나 감소합니다.

#### GMV Max 사용자 지정 ROAS 🅘
광고의 수익성을 높이기 위해 목표 ROAS(목표)를 설정합니다. 광고 학습 7일 후 목표 ROAS의 86%-126%를 달성하십시오.

**ROAS 옵션:**
- ⭕ **ROAS = 3.0** - 유사한 광고의 30%보다 경쟁력이 높습니다.
- **ROAS = 4.4** - 유사한 광고의 50%보다 경쟁력이 높습니다.
- **ROAS = 5.3** - 유사한 광고의 70%보다 경쟁력이 높습니다.
- **값 설정**

> 최상의 결과를 얻으려면 게시 후 처음 14일 동안 ROAS 목표 설정을 수정하지 마십시오. [자세히 알아보기](link)

#### 수동
모든 입찰 설정을 수동으로 구성합니다.

---

## 문제 해결

문제가 발생하면 [개발자 가이드](link)를 확인하고 [FAQ](link)를 검색하는 것이 좋습니다.

문제가 발생하면 Open Platform Console에 로그인하여 [티켓 제출](link)하십시오.

---

## Ads service API 권한

📍 https://open.shopee.cn/faq/381

---

**참고**: ⓘ 입찰 방법에 대해 알아보십시오.

## 사용 사례

1. 프로그래밍 방식으로 광고 캠페인 관리
2. 광고에 대한 추천 키워드 검색
3. API를 통한 상품 광고 생성
4. 광고 예산 및 ROAS 최적화
5. 광고 데이터를 외부 시스템에 통합

## 관련 API

- ads_get_recommended_keywords_by_item
- ads_get_recommended_products
- ads_get_campaign_list
- ads_get_product_campaign_data
- ads_get_product_campaign_daily_performance
- ads_get_product_keyword_performance
- ads_create_product_ads
- ads_create_manual_product_ads
- ads_set_environment_item_list
- ads_set_recommended_keywords_list
- ads_get_recommended_roi_target
- v2.ads.get_product_recommended_roi_target
- v2.ads.get_create_product_ad_budget_suggestion

---

## 원문 (English)

### Summary

This guide provides rules and instructions for using the Ads Open API. It covers topics such as data protection, budget management, keyword retrieval, campaign data, ad creation, and ROAS optimization. The guide also includes troubleshooting tips and API permission information.

### Content

# Ads Open API Usage Rules

When using the ads/storefront_api, you need to abide by the following rules as have listed as how Shopee can create a fair and equal marketplace for both sellers and buyers:

1. When processing advertising data, you need to comply with the platform's **data protection policy**.
2. With all items in S, Shopee's official cooperation CPS limit is used with platform users. If can be used to the new user/buyer for affiliate marketing use, the payment amount must be reasonable and cannot be used for other purposes.
3. Since data may be used for any purpose other than shop operations and official Shopee cooperation projects, when accessing data through this API, you need to obtain the consent of the platform seller, and be able to use the information publicly.
4. In no way/should a catalyst: you, our names and links, you will receive a warning email from Shopee. If you fail to modify it in time, you may lose your privileges.

---

# Ads Open API Instruction

## Balance Ad Budget and Retrieve Budget Information

- You can top up your ad account with **wallet balance** and we will notify you via email once done.
- Your credits, before use, convert their balance status through this API to ensure that the balance is sufficient to run the advertisements.
- If your balance is insufficient, you need to return to where the shop's ad settings, such as whether daily bid, get and price optimization features are enabled.

### MyLite/Shopee-Budgetbalada-wallet-top-up

[Interface elements showing budget balance and top-up options with "Top-up balance" button highlighted]

**Campaign, Single** - Enable balance notifications

[Screenshot showing "Shopee Ads Shop Level Settings" dialog with balance notification settings]

---

## 3. Retrieve Recommended Keywords and Product Data

You can use <CT ads_get_recommended_keywords_by_item> API to retrieve a list of recommended keywords for item, along with optional search keywords. Search info also view keyword popular scores, search volume, and suggested bids.

- <CT ads_get_recommended_products> Item list that can be added to advertising campaigns. We also define if an item is recommended, and if item fails to promote, it will display a list of recommended SDGs shop-level ads specific campaigns.

## Select Products

[Table showing product selection interface with columns for Product, Status, Views, Product Name]

---

## Add Keywords

[Table showing keyword data with columns for Keywords, Exposure (Last 30 D), Popularity, etc.]

---

## 4. Retrieve Campaign Level Data

Use the <CT ads_get_campaign_list> API to retrieve the list of all campaigns associated with a specific shop.

Use the API to get the list of all the campaigns.

- <CT ads_get_product_campaign_data> and <CT ads_get_product_campaign_daily_performance> Use this API to retrieve detailed settings for a specific campaign and <CT ads_get_product_keyword_performance> and <CT ads_get_product_campaign_daily_performance> Use these APIs to obtain campaign-level performance data, such as ad spend, impressions, clicks, conversions, etc.

### API_CREATE_1

[Performance data table showing metrics like Exposure, Click, Cost, etc.]

### Conversions

[Conversion tracking interface]

### Performance Chart

[Chart showing performance metrics over time with data points for Exposure, Click, Direct Order, etc.]

---

## 5. Create Ads Product Ads

Use <CT ads_create_product_ads> API to create ads. (Retrieve shop creation)

---

## Products

**Budget Ad Storefront**

Automatically select best products(s)

- Add products to the shop's storefront ads
- Retrieve the ad in your Seller Center

---

## 6. Create Daily Max and Manual Product Ads

- <CT ads_create_manual_product_ads>

To create daily max ads, pass bidding_method:manual

To create Manual ads, pass bidding_method:manual

### Summary

[Interface showing budget settings]

### Manually manage products

[Interface showing manual product management options]

### Bidding

[Interface showing bidding options with "Set by CPC Model" and "Manually Individually" radio buttons]

---

## Basic Settings

When creating manual ads, bidding_method:manual, you can use <CT ads_set_environment_item_list> and <CT ads_set_recommended_keywords_list> to get suggested products and keywords for advertising. Then, input the budget and ad-tier-recommendations data to pass to the creation API.

---

## Retrieve Ad Budget and ROAs

- Retrieve ad budget use ROAs
- <CT ads_get_recommended_roi_target> Get the ROAs (Return on Ad Spend) value when creating CPC Max ads.

---

## Basic Settings

[Warning box with text: "❗ SET BY CPC MODEL - Automatically adjust keyword bidding for better exposure"]

**Notice: When creating CPC Max ads, PASS THE ROAs target value to the API (Return on Ad Spend).**

---

# 7. Retrieve Ad Budget and ROAS

## API Endpoints

- **v2.ads.get_product_recommended_roi_target**: Get the ROAS (Return on Ad Spend) value when creating GMV Max ads.

- **v2.ads.get_create_product_ad_budget_suggestion**: Retrieve budget suggestions for creating product ads.

---

## Basic Settings

### Budget
- **Unlimited**
- **Set daily budget**: Rp 0
  - Suggested: Rp7,046

### Duration
- **No end date**
- **Set start/end dates**: 📅 21/04 - 21/05 (GMT+7)

---

## Bidding

### Bidding Method

#### GMV Max
Let Shopee find the best way to deliver your ads to shoppers for optimal performance over time.

**Set your ROAS target**

**GMV Max Auto Bidding**
Shopee increases or decreases within your daily budget while maintaining a healthy ROAS.

#### GMV Max Custom ROAS 🅘
Set a Target ROAS (goal) to enhance your ad's profitability. Achieve 86%-126% of your Target ROAS after 7 days of ad learning.

**ROAS Options:**
- ⭕ **ROAS = 3.0** - Better competitiveness than 30% of similar ads
- **ROAS = 4.4** - Better competitiveness than 50% of similar ads  
- **ROAS = 5.3** - Better competitiveness than 70% of similar ads
- **Set a value**

> For best results, avoid modifying your ROAS target setting in the first 14 days after publishing. [Learn more](link)

#### Manual
Configure all bidding settings manually.

---

## Troubleshooting

If you encounter problems, it is recommended to check the [Developer Guide](link) and search for [FAQ](link).

If you encounter problems, please log in to the Open Platform Console to [Raise a ticket](link).

---

## Ads service API Permission

📍 https://open.shopee.cn/faq/381

---

**Note**: ⓘ Learn about bidding methods

---

**문서 ID**: developer-guide.277
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/277
**처리 완료**: 2025-10-16T08:59:08
