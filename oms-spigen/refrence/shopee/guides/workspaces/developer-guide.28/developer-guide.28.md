# CNSC API 통합 가이드

**카테고리**: 통합
**난이도**: medium
**중요도**: 5/5
**최종 업데이트**: 2025-10-16T08:25:49

## 개요

본 가이드는 CNSC (Cross-border New Seller Center) API 통합에 대한 지침을 제공합니다. 개발자 등록, 앱 생성, API 테스트, 판매자/샵 인증을 다루며, CNSC로 업그레이드된 샵에 대한 특정 고려 사항을 포함합니다.

## 주요 키워드

- CNSC
- API integration
- Shopee Open Platform
- Developer registration
- App creation
- API testing
- Merchant authorization
- Shop authorization
- Global Product API
- Order APIs

## 본문

# CNSC API 통합 가이드

## 시작하기 - CNSC API 통합 가이드

중요: CNSC 판매자와 FBP 개발자(ISV)를 관리하는 CNSC 쇼핑몰은 CNSC 업그레이드에 주의해야 하며, 주문 API를 마이그레이션해야 하는 API를 제공합니다. 개발자가 사용해야 하는 API에는 글로벌 상품 API, 글로벌 주문 API, 물류 API, 금융 API, 퍼스트마일 물류 API가 포함되며, 판매자 API는 변경할 수 없습니다. CNSC로 업그레이드된 쇼핑몰은 오류 없이 주문을 관리하기 위해 이전 API 호출을 사용할 수 없습니다.

또는 쇼핑몰이 CNSC로 업그레이드되었는지 어떻게 판단할 수 있습니까?
API를 통해 CB 지역을 조회할 수 있습니다. 반환 매개변수가 "CNSC"이면 해당 쇼핑몰은 CNSC 지역에 속합니다. 매개변수 cbcshipType_id가 반환되지 않으면 매개변수 merchant_ext를 쿼리합니다. CB 유형이 판매자의 경우 "CNSC"(cbcshipType)이면 CNSC로 업그레이드된 새 쇼핑몰입니다.

Q: 제 시스템에 상품 관련 기능, 즉 상품 생성 및 상품 업데이트 기능이 없고 일부 애플리케이션만 지원하는 경우 어떻게 해야 합니까?
A: 시스템에 상품 관리 기능이 없으면 사용할 API가 없습니다.

---

## 2. CNSC API 통합 방법

### 2.1 개발자로 등록 (이미 개발자 계정이 있는 경우 이 단계를 건너뛸 수 있습니다.)

a. "가입"을 클릭하고 "계약"을 읽고 Shopee Open Platform 계정을 등록하여 https://open.shopee.com 에 접속합니다.

[스크린샷은 "Documentation English SIGN UP" 탐색 및 이메일/비밀번호 필드가 있는 가입 양식이 있는 Shopee Open Platform 홈페이지를 보여줍니다.]

b. 인증 이메일을 받게 됩니다. 이메일을 확인하고 비밀번호를 설정하십시오.

---

**Open Platform**

### 개발자님께,

Shopee Open Platform 계정을 성공적으로 생성했습니다.

이메일 주소를 확인하려면 다음 버튼을 클릭하십시오.

**[이메일 확인]**

감사합니다.

Shopee Open Platform

---

c. 계정으로 로그인합니다.

[스크린샷은 이메일/비밀번호 필드가 있는 Shopee Open Platform 로그인 양식을 보여줍니다.]

---

### 2.2 계정 정보 완료 (이미 승인된 개발자 계정이 있는 경우 이 단계를 건너뛸 수 있습니다.)

a. 등록하려는 개발자 계정의 계정 유형은 앱 유형이어야 합니다(쇼핑몰 유형은 Seller Center 및 Open API를 사용하는 개별 판매자의 계정입니다). 개발자의 유형에 따라 다른 유형의 앱이 표시됩니다. 자세한 내용은 앱 통합 페이지를 참조하십시오.

b. 선택한 쇼핑몰 유형에 따라 적절한 프로필을 작성하십시오. 정보는 Shopee 플랫폼과 일치해야 합니다.

---

### 2.3 앱 생성

Shopee Channels를 사용하려면 먼저 앱을 만들어야 합니다. 승인된 판매자로서 개발자 프로필에서 앱을 만들 수 있습니다.

a. 기존 API 유형은 V2 API를 지원하지 않습니다. 기존 앱이 Original 유형인 경우 설명서에 따라 앱 업그레이드 프로세스를 완료하십시오.

b. 앱을 만들려면 다음 경로를 통해 새로 만들 수 있습니다.
각 페이지의 왼쪽 상단에서 기사에 따라 권한을 확인하십시오.

c. 앱을 만든 후 다음 경로를 통해 새로 만들 수 있습니다.
API 유형으로 Open API를 선택할 수 있습니다. 생성 시 V2를 선택하십시오. 자세한 내용은 정보를 읽어보십시오.

---

### 2.4 앱 출시

앱 상태를 온라인으로 제출해야 합니다. "Go Live"를 클릭하여 "파트너" 또는 "콘솔" 또는 "앱 목록" 또는 "개발/샌드박스 선택" 등에서 라이브 환경 API 사용을 요청합니다.

[스크린샷은 앱 ID, 앱 키, 테스트 쇼핑몰 등을 포함한 문서, 영어 및 앱 세부 정보가 있는 앱 관리 인터페이스를 보여줍니다.]

**기본 정보**
- App ID: [redacted]
- App Key: [키 관리 버튼 표시]
- App secret: [숨김]
- Test API URL: [URL 표시]
- Redirect URL (callback): https://
- Status: Development
- Inventory
- Test shop link: https://
- API valid code: [코드 표시]
- Your return (self): Singapore
- Testing (self): Global Product

**앱 API는 24시간 후에 자동으로 루틴을 완료하고 API 상태가 온라인 상태로 변경되며 앱이 프로덕션에 자동으로 배포됩니다.**

---

# API 테스트 가이드

## 개요
- **APP**: xxx
- **Testing**: [link]
- **API**: xxx 및 xxx
- **API call xxx**: xxx
- **App Service Region**: Singapore
- **Live Consultation**: xxx

APP는 24시간 후에 자동으로 모드를 완료하고 APP 상태가 온라인 상태로 변경되며 개발자는 라이브 환경의 라이브 프리미엄 및 파트너 도움말을 확인할 수 있습니다.

---

## 2.5 API 테스트 시작

### 1단계: 테스트 계정 생성
1) 테스트를 시작하려면 Cassie에서 중국 판매자를 만들어야 합니다.

**테스트 계정 생성 양식:**
- **Merchant ID**: [필드]
- **Merchant Name**: [필드]
- **Merchant Status**: Organization OP (Suspended)
- **Account Status**: Active
- **API Key**: [redacted]

**테스트 계정 생성 대화 상자:**
- Local Time
- Local Date
- Cross-Border Shop ☐ China Domestic ☐
- Location: [드롭다운]
- Industry: [드롭다운]
- Username: [필드]
- Password: [필드]
- Confirm Password: [필드]
- Submit button

---

### 2단계: 로그인 및 구성
2) https://seller-test-admin.tuyacn.cn 에 로그인합니다. 우편번호는 123456입니다.
   * 참고: 사용자 지역 전환 또는 체크인 시간 초과 중에는 확인이 필요합니다.

**中国发起跨境商家联系中心**
如果您联系在异中国的分销

**验证码登录/注册表单:**
- 账号登录
- 手机登录显示
- 输入账号/手机号
- 输入密码
- [password] 필드
- 登录 button

---

### 3단계: 통화 설정
3) 통화 쇼핑몰 유형을 설정합니다.
   * 참고: CNSC 샌드박스에서 테스트할 때 모든 CNY 통화를 사용할 수 있습니다.
   * 실제 CNSC(라이브 환경)에서는 통화 옵션이 테스트 및 지원에 대해 USD 또는 CNY입니다(향후 ABC).

**店铺管理后台跨境商家统筹工具选择页面**

货运管理界面选择:
- 国内店铺类型: [표시된 옵션]
- 币种选择框 (통화 선택)
- 确认 (확인) button

---

### 4단계: 마켓 정보 가져오기
4) 마켓 정보를 가져옵니다. 잘못된 번호를 받으면 스크린샷에 올바른 번호 범위를 표시합니다.

**마켓 정보 인터페이스:**
- Market configuration settings
- Location: [필드]
- Language: [필드]  
- Currency: [필드]

---

### 5단계: 인터페이스 언어
5) 그런 다음 판매자 설정 페이지에서 인터페이스 언어를 영어로 설정할 수 있습니다.

**판매자 설정 페이지:**
- Profile settings
- Language selection 드롭다운
- "Start to upgrade to Global BG" 버튼 강조 표시

---

### 6단계: 공식 시작 페이지
6) 공식 시작 페이지를 확인합니다. 설정을 완료하고 "Start to upgrade to Global BG" 버튼을 클릭하십시오.

**Shop Connection Level Global BG 및 Shop BG 대화 상자:**

이 대화 상자에는 여러 섹션이 있습니다.
- Shop connection status information
- Configuration requirements
- API connection details
- Terms and conditions text
- Confirmation checkboxes
- Action buttons at bottom

**표시되는 주요 섹션:**
- Shop information verification
- API configuration requirements  
- Terms of service agreement
- Confirm and proceed options

---

## 참고 사항 및 경고

⚠️ **중요**: 
- 테스트 환경은 우편번호 123456을 사용합니다.
- 통화 옵션은 환경(샌드박스 vs 라이브)에 따라 다릅니다.
- API 테스트 전에 모든 설정 단계를 완료하십시오.
- 진행하기 전에 마켓 구성을 확인하십시오.

---

## 다음 단계

이러한 설정 단계를 완료한 후 샌드박스 환경에서 API 통합 및 테스트를 진행할 수 있습니다.

---

# 개발자 가이드 - Shopee Open Platform

## 2.6 판매자 및 쇼핑몰 인증

쇼핑몰 계정이 CNSC로 업그레이드된 후에는 여러 쇼핑몰이 판매자에 속하므로 이 [FAQ](link)를 통해 판매자와 쇼핑몰 간의 관계에 대해 자세히 알아볼 수 있습니다.

따라서 GlobalProduct API를 정상적으로 호출하려면 쇼핑몰을 다시 인증하고 인증 시 하위 계정 페이지를 전환하십시오.

### Shopee Openplatform APP 인증을 위해 로그인

**인터페이스 요소:**
- Country selector: SG
- Email / Phone / Username 필드
- Password 필드
- Forgot Password ? 링크
- Log In button
- OR separator
- **Switch to Sub Account** 버튼 (화살표로 강조 표시)

Language selector: English

---

### 인증 지침

인증에는 기본 계정을 사용하십시오. 하위 계정은 인증을 완료할 수 없습니다. 기본 계정의 형식은 다음과 같습니다. **" main

계정에 성공적으로 로그인하면 인증 페이지가 표시됩니다. 먼저 인증하려는 쇼핑몰을 확인한 다음 Auth Merchant를 클릭하십시오. Auth Merchant를 확인하지 않으면 관련 글로벌 상품 API를 호출하거나 판매자 정보를 얻을 수 없습니다. 일부 쇼핑몰을 확인하지 않으면 API를 통해 관련 쇼핑몰에 상품을 게시할 수 없습니다. 따라서 완전히 관리하려는 판매자와 쇼핑몰을 확인하십시오.

**인증 페이지 요소:**
- 영어 언어 선택기가 있는 Shopee Open Platform 헤더
- 판매자 정보 섹션("Auth Merchant" 버튼을 가리키는 화살표로 강조 표시)
- 다음을 보여주는 쇼핑몰 목록:
  - [PH] rec_account: [ID info]
  - [PH] rec_account: [ID info]  
  - [PH] rec_account: [ID info]
- "Auth Merchant" 확인란 및 버튼 (강조 표시)
- 유사한 형식의 추가 쇼핑몰 항목
- 하단의 Reject 및 Allow 버튼
- "Reject all (0/3)" 옵션

---

### 콜백 정보

인증에 성공하면 콜백 주소는 쇼핑몰 ID 대신 main_account_id를 반환합니다. Main_account_id는 나중에 AccessToken을 얻는 데 사용됩니다.

그런 다음 이 [문서](link)에 따라 토큰을 새로 고치십시오.

GetAccessToken API를 통해 해당 시점에 성공적으로 인증된 모든 판매자 ID 및 쇼핑몰 ID 목록을 얻을 수 있습니다.

---

### 중요 사항

**이전에 시스템에 인증된 CNSC 쇼핑몰을 업그레이드하고 별도로 저장해야 하는 경우:** 인증을 완료하면 처음 얻은 새로 고침 토큰 및 액세스 토큰을 현재 인증된 판매자 ID 및 쇼핑몰 ID에서 공유할 수 있습니다. 그러나 RefreshAccessToken API를 호출하여 다른 쇼핑몰 ID에 대해 다른 새로 고침 토큰을 얻으면 다른 새로 고침 토큰 및 액세스 토큰이 반환되므로 토큰을 별도로 보관하십시오.

판매자 ID와 쇼핑몰 ID 간의 관계를 얻으려면 [get_shop_list_by_merchant](link)를 호출하십시오.
그런 다음 판매자 API [get_merchant_info](link)를 호출하여 각 판매자 정보를 얻고 [get_shop_info](link)를 호출하여 각 쇼핑몰 정보를 얻을 수 있습니다.

---

## 3. 요약

**주요 사항:**

1) 앱을 만들거나 기존 앱을 사용하는 경우 앱 유형이 V.2 GlobalProduct API를 호출할 수 있는지 확인하십시오. (유형: public_app/private_shop_app)

2) GlobalProduct API를 호출하기 전에 쇼핑몰을 CNSC로 업그레이드해야 합니다.

3) 판매자 사용자가 프로덕션 환경에서 CNSC API를 사용할 수 있도록 APP 상태가 **Online**인지 확인해야 합니다. 앱을 라이브로 클릭하면 Open Platform에서 상품 정보를 받기 시작합니다. APP는 24시간 후에 자동으로 승인됩니다.

4) 판매자는 CNSC에 로그인하여 CNSC 상품 업그레이드가 성공하기 전에 관련 설정을 완료해야 합니다.

5) 상품이 성공적으로 업그레이드된 경우에만 Global Product API를 정상적으로 호출할 수 있습니다.

6) 인증이 완료된 후 토큰을 새로 고치십시오([문서](link) 기준).

---

### 지원

이 문서에서 다루지 않는 기술 연결 문제가 있는 경우 [티켓 시스템](link)을 통해 Shopee Open Platform에 문의하십시오.

## 사용 사례

1. Shopee의 CNSC API와 새로운 애플리케이션 통합
2. 기존 애플리케이션을 CNSC 지원하도록 마이그레이션
3. CNSC 샵의 상품 및 주문 데이터 관리
4. CNSC 판매자를 위한 샵 인증 자동화
5. CNSC 샌드박스 환경에서 API 통합 테스트

## 관련 API

- Global Product API
- Global Order API
- Logistics API
- Finance API
- First-mile Logistics API

---

## 원문 (English)

### Summary

This guide provides instructions for integrating with the CNSC (Cross-border New Seller Center) API. It covers developer registration, app creation, API testing, and merchant/shop authorization, with specific considerations for shops upgraded to CNSC.

### Content

# CNSC API Integration Guide

## Getting Started - CNSC API Integration Guide

Important: CNSC Sellers and CNSC shops managing FBP developers (ISVs) need to pay attention are upgrade to CNSC need to pay attention and we provide API needs to migrate order APIs. The APIs you need to be developers include Global Product API, Global Order API, Logistics API, Finance API, First-mile Logistics API and Merchant APIs are not changeable. CNSC upgrades shops cannot use old API calls to manage orders without error.

Or how to judge whether a shop has been upgraded to CNSC?
You can inquire into its CB region by API, if the return parameter is "CNSC", its Shop belongs to CNSC region, if no return the parameter cbcshipType_id, then query the parameter merchant_ext, if CB type is "CNSC"(cbcshipType) for the merchant. Otherwise this is new shop upgraded to CNSC.

Q:If my system does not have product-related function, that is, it does not include product creation and product update function, only supports to some applications?
A:If your system does not have product management functions, there is no API use.

---

## 2. How to integrate CNSC API

### 2.1 Register as a developer (If you already have a developer account, you can skip this step)

a. Click "Sign up" and read "Agreement", and register a Shopee Open Platform account to enter https://open.shopee.com

[Screenshot shows Shopee Open Platform homepage with "Documentation English SIGN UP" navigation and Sign up form with email/password fields]

b. You will receive a verification email, please verify your email and set a password.

---

**Open Platform**

### Dear developer,

You have successfully created your Shopee Open Platform Account.

To verify your email address, please click the button:

**[Verify my email]**

Best regards,

Shopee Open Platform

---

c. Login with your account

[Screenshot shows Shopee Open Platform Sign in form with email/password fields]

---

### 2.2 Complete account information (If you already have an approved developer account, you can skip this step)

a. Please note that the account type of developer account you are registering for should be App type (Shop type is individual Seller's account using Seller Center and Open APIs). Different types of developer's will see different types of Apps. For details, please refer to the App integration page.

b. Complete the appropriate profile according to the type of shop you choose. The information needs to match the Shopee platform.

---

### 2.3 Create an app

To use Shopee Channels, you will need to create an App first. You can create apps from the developer profile as approved Seller.

a. Original API Type does not support V2 API. If your existing app is Original type, please complete the app upgrade process according to documentations

b. To create an app, you can create new through this path:
In upper left on each, please check the the permission according to the article.

c. After create an app, you can create new through this path
You can select Open API as API type. V2 when creating. Please read the information for further information.

---

### 2.4 Go live app

You must submit your app status is online. Click "Go Live" to request to use the Live environment APIs under "partners" or "Console" or "App List" or "Select Development/sandbox" or...

[Screenshot shows app management interface with Documentation, English, and app details including App ID, App Key, Test shop, etc.]

**Basic Information**
- App ID: [redacted]
- App Key: [shows key management buttons]
- App secret: [hidden]
- Test API URL: [URL shown]
- Redirect URL (callback): https://
- Status: Development
- Inventory
- Test shop link: https://
- API valid code: [code shown]
- Your return (self): Singapore
- Testing (self): Global Product

**App API are automatic ally complete the routine after 24 hours and the API status will change to the Online status, and the app will be automatically deployed in production.**

---

# API Testing Guide

## Overview
- **APP**: xxx
- **Testing**: [link]
- **API**: xxx and xxx
- **API call xxx**: xxx
- **App Service Region**: Singapore
- **Live Consultation**: xxx

APP will automatically complete the modes after 24 hours and the APP status will change to the Online status, and the developers can check the Live Premium to and Partner help of the Live Environment.

---

## 2.5 Start API testing

### Step 1: Create Test Account
1) To start testing, you should create a China merchant on Cassie.

**Test Account Creation Form:**
- **Merchant ID**: [field]
- **Merchant Name**: [field]
- **Merchant Status**: Organization OP (Suspended)
- **Account Status**: Active
- **API Key**: [redacted]

**Create Test Account Dialog:**
- Local Time
- Local Date
- Cross-Border Shop ☐ China Domestic ☐
- Location: [dropdown]
- Industry: [dropdown]
- Username: [field]
- Password: [field]
- Confirm Password: [field]
- Submit button

---

### Step 2: Login and Configuration
2) Login https://seller-test-admin.tuyacn.cn. Please note that zip is 123456.
   * Note: During user region transition or check-in timeout, it will require verification.

**中国发起跨境商家联系中心**
如果您联系在异中国的分销

**验证码登录/注册表单:**
- 账号登录
- 手机登录显示
- 输入账号/手机号
- 输入密码
- [password] field
- 登录 button

---

### Step 3: Currency Setup
3) Set the currency shop type:
   * Note: When testing in the CNSC sandbox, any CNY currency is available.
   * In the actual CNSC (Live environment), the currency options will be USD or CNY for test and support (ABC in the future).

**店铺管理后台跨境商家统筹工具选择页面**

货运管理界面选择:
- 国内店铺类型: [options shown]
- 币种选择框 (Currency selection)
- 确认 (Confirm) button

---

### Step 4: Get Market Info
4) Get the market info. If you get the wrong number, we will prompt the correct range of the number on the screenshot.

**Market Information Interface:**
- Market configuration settings
- Location: [field]
- Language: [field]  
- Currency: [field]

---

### Step 5: Interface Language
5) Then, you can set the interface language in Merchant setting page to English.

**Merchant Setting Page:**
- Profile settings
- Language selection dropdown
- "Start to upgrade to Global BG" button highlighted

---

### Step 6: Official Start Page
6) Check the official start page. Please finish the setting and click the button "Start to upgrade to Global BG."

**Shop Connection Level Global BG and Shop BG Dialog:**

This dialog contains multiple sections:
- Shop connection status information
- Configuration requirements
- API connection details
- Terms and conditions text
- Confirmation checkboxes
- Action buttons at bottom

**Key sections visible:**
- Shop information verification
- API configuration requirements  
- Terms of service agreement
- Confirm and proceed options

---

## Notes and Warnings

⚠️ **Important**: 
- Test environment uses zip code 123456
- Currency options depend on environment (sandbox vs live)
- Complete all setup steps before API testing
- Verify market configuration before proceeding

---

## Next Steps

After completing these setup steps, you can proceed with API integration and testing in the sandbox environment.

---

# Developer Guide - Shopee Open Platform

## 2.6 Merchant & Shop Authorization

Because after the shop account is upgraded to CNSC, multiple shops will belong to merchants, you can learn more about the relationship between Merchant and shop through this [FAQ](link).

Therefore, in order for you to call the GlobalProduct API normally, please re-authorize your shops and switch the sub account page when authorization.

### Login to Authorize Shopee Openplatform APP

**Interface elements:**
- Country selector: SG
- Email / Phone / Username field
- Password field
- Forgot Password ? link
- Log In button
- OR separator
- **Switch to Sub Account** button (highlighted with arrow)

Language selector: English

---

### Authorization Instructions

Please use your main account for authorization. Sub-accounts cannot finish the authorization. The account format of the main account is: **" main

After the account is successfully logged in, you will see the Authorization page, first check the shop you want to authorize and then click Auth Merchant. If you do not check Auth Merchant, you will not be able to call the relevant Global Product API or to obtain the Merchant information. If some shops are not checked, the product cannot be published to the related shop through the API. So please make sure you check the Merchant and Shop that you want manage completely.

**Authorization Page Elements:**
- Shopee Open Platform header with English language selector
- Merchant information section (highlighted with arrow pointing to "Auth Merchant" button)
- Shop list showing:
  - [PH] rec_account: [ID info]
  - [PH] rec_account: [ID info]  
  - [PH] rec_account: [ID info]
- "Auth Merchant" checkbox and button (highlighted)
- Additional shop entries with similar format
- Reject and Allow buttons at bottom
- "Reject all (0/3)" option

---

### Callback Information

After the authorization is successful, the callback address will return the main_account_id instead of shop id. Main_account_id will be used to obtain AccessToken later.

Then please refresh the token according to this [document](link).

Through the GetAccessToken API, you will get a list of all merchant ids and shop ids that have been successfully authorized at that time.

---

### Important Notes

**If you are upgrading your CNSC shops which have been authorized to your system before, and you need to store them separately:** When you completing the authorization, the initial refresh token and access token you obtained can be shared by the currently authorized merchant id and shop id. But then you call the RefreshAccessToken API to obtain different refresh token for different shop ids, will return different refresh tokens and access token, so please keep their tokens separately.

If you want to obtain the relationship between the merchant id and the shop id, please call the [get_shop_list_by_merchant](link).
Then you can get each merchant information by calling the seller API: [get_merchant_info](link) and call [get_shop_info](link) to get each shop information.

---

## 3. Summary

**Key Points:**

1) When you create an app, or use an existing app, please make sure that your app type can call the V.2 GlobalProduct API. (type: public_app/private_shop_app)

2) The shop needs to be upgraded to CNSC before you call GlobalProduct API

3) It is necessary to ensure that the APP status is **Online**, so that seller users can use the CNSC API in the production environment. Clicking Live the App, Open Platform will start to receive product information. APP will automatically be approved after 24 hours.

4) The seller needs to log in to CNSC and complete the relevant settings before the CNSC product upgrade is successful.

5) Only when the product is successfully upgraded, can the Global Product API be called normally.

6) Please refresh your token after authorization is finished (based on [document](link))

---

### Support

If you have any technical connection problems that are not covered by this document, please contact Shopee Open Platform through the [ticket system](link).

---

**문서 ID**: developer-guide.28
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/28
**처리 완료**: 2025-10-16T08:25:49
