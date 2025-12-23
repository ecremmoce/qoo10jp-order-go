# KRSC API 통합 가이드

**카테고리**: 통합
**난이도**: medium
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T08:28:17

## 개요

본 가이드는 개발자 등록, 앱 생성, API 테스트, 환경 설정을 포함하여 KRSC API와 통합하는 방법에 대한 지침을 제공합니다. KRSC로 업그레이드하고, 상점을 인증하고, GlobalProduct API를 사용하는 데 필요한 단계를 다룹니다.

## 주요 키워드

- KRSC API
- API 통합
- 개발자 등록
- 앱 생성
- API 테스팅
- Sandbox 환경
- 상점 인증
- GlobalProduct API
- Merchant API
- KBSC 업그레이드

## 본문

# KRSC API 통합 가이드

## 공지 | 콘솔

### 시작하기 > KRSC API 통합 가이드

중요: KRSC 토큰 및 KRSC kbsp 마케팅 KSP 개발자(ISV)는 KRSC 업그레이드에 주의해야 하지만 제품 API는 ubrsc로 마이그레이션해야 합니다. 개발자 중심 KRSC API는 일반적으로 기능 API이지만 탐색할 수 없는 기술 API도 포함됩니다. KRSC 업그레이드는 API 호출 또는 제품 관련 모듈을 사용할 수 없습니다.

#### FAQ:

**Q: 상점이 KRSC로 업그레이드되었는지 어떻게 판단합니까?**
A: 상점 OS 관리 API를 열 수 있습니다. 매개변수 timezhan_at이 보이면 (응답 timezhan_at을 엽니다. 비어 있지 않은 경우 상점이 KRSC로 업그레이드되지 않았음을 의미합니다.

**Q: 시스템에 제품 관련 기능이 없는 경우 제품 생성 및 제품 업데이트 기능 인터페이스를 포함하여 스크립트 조정을 수행하려면 어떻게 사용해야 합니까?**
A: KRSC를 업그레이드한 후 제품 API가 비활성화되므로 사용할 수 있습니다.

---

## 2. KRSC API 통합 방법

### 2.1 개발자로 등록 (이미 개발자 계정이 있는 경우 이 단계를 건너뛸 수 있습니다.)

a. "가입"을 클릭하고 "계약"을 읽고 포털에서 Shopee Open Platform 계정을 등록하십시오: https://open.shopee.com

**Shopee Open Platform**
- 이메일/비밀번호 필드가 있는 가입 페이지
- 등록 양식 인터페이스

b. 확인 이메일을 받게 됩니다. 이메일을 확인하고 비밀번호를 설정하십시오.

---

### Open Platform

**개발자님께,**

Shopee Open Platform 계정을 성공적으로 생성했습니다.

이메일 주소를 확인하려면 버튼을 클릭하십시오.

[이메일 확인]

감사합니다.
Shopee Open Platform

---

c. 계정으로 로그인

**Shopee Open Platform**
- 로그인
- 이메일/비밀번호 로그인 양식

---

### 2.2 계정 정보 완료 (이미 승인된 개발자 계정이 있는 경우 이 단계를 건너뛸 수 있습니다.)

프로세스가 완료되면 다음을 통해 개발자 계정 정보를 완료하십시오: Open Platform에 로그인 -> 콘솔 -> 앱 관리 -> [Your App] -> 개발자 프로필 완료

**개발자 유형에 따라 다른 유형의 상점을 사용합니다.** 자세한 내용은 앱 관리 페이지를 참조하십시오.

a. 선택한 개발자 계정 유형에 따라 적절한 프로필을 완료하십시오. 정보는 Shopee 플랫폼에서 검토해야 합니다.

---

### 2.3 앱 만들기

Shopee OpenAPI를 호출하려면 먼저 앱을 만들어야 합니다. 개발자 프로필이 승인되면 앱을 만들 수 있습니다.

**참고:**
- Original API Type은 V2.0 api를 지원하지 않습니다. 기존 앱이 Original 유형인 경우 앱 업그레이드를 완료하십시오.
- 앱을 업그레이드하려면 이 경로를 통해 하나를 만들 수 있습니다.
- Audit 유형의 앱 페이지의 경우 기사에 따라 SPI 권한을 확인하십시오.
- 앱을 만든 후 이 경로를 통해 이동할 수 있습니다.
- 앱 목록에서 Public type V2.0을 클릭하거나 새 앱을 만듭니다. 정보를 완료한 후 제출하십시오.

---

### 2.4 앱 출시

앱 상태가 온라인 상태인지 제출해야 합니다. "Go Live"를 클릭하여 Shopee에서 라이브 (상업용) API 사용을 요청하십시오 -> 콘솔 -> 앱 목록 -> [Select Shop] -> Go-live 요청 또는 Go-live 요청 취소 -> Go live.

**Shopee Open Platform 콘솔**

- 앱 관리
  - 앱 목록 → #1
  - 아직 URL 주소를 제출하지 않음
  - URL을 연결해야 합니다.
  - 사용자를 리디렉션할 수 있는 리디렉션 URL을 추가하십시오...

**도구**
- 로그

**앱 키**
- Key ID
- Test Shops
- Test Shop
- Test_URL

---

### 기본 정보

**요청 중**
- 다음을 추적할 수 있습니다...
- URI를 실행할 수 있습니다.
- 테스트 계정 시작 → 싱가포르
- 앱이 이제...

---

**APP ID는 24시간 후에 자동으로 모듈을 완료하고 APP 상태가 Go Online 상태로 변경되며 API를 프로덕션 및 테스트 목적으로 정상적으로 호출할 수 있습니다.**

---

# 개발자 가이드 - API 테스트 및 환경 설정

## 개요

**App ID:** [ID Number]

**상태 정보:**
- App ID: [ID]
- Sandbox: running
- App in prod: live
- In prod since: [Date]
- App Secret Status: [Status]
- Live Capabilities: [Details]

---

## 2.5 API 테스트 시작

`USE CHNC Sandbox`가 준비되었습니다. KPEC Devs는 KPDC와 기능이 유사하므로 CHNC Sandbox를 사용하여 API를 실행할 수 있습니다.

테스트를 시작하려면 Cpaxide에 `chisa.passportapi.account`가 있어야 합니다.

### 계정 설정 단계

1. **계정 가져오기**
   - User Name: [Enter Name Here]
   - Referrer URL: [URL field]
   - Merchant ID: 100474A

2. **테스트 계정 만들기**
   - 계정 유형을 선택하십시오:
     - 옵션 드롭다운 메뉴 표시:
       - [Account type selection]
       - Create button

3. **로그인 프로세스**
   - Login https://sandbox-test.static.example.cn
   - 앱은 123456입니다.
   - **Note**: 영어이거나 필요한 경우 병합된 계정 테스트/비밀번호/계정으로 신청하십시오.

---

## 3. 欢迎来到中国服务中心
## 设置开发者账户并开始测试

[양식 필드와 확인 대화 상자가 있는 서비스 센터 인터페이스를 보여주는 스크린샷]

**Instructions:**
4. 로그인 후 통화를 가져옵니다.
   - PC Web CHNC 샌드박스 페이지에서 모든 CNY 통화는 실제 KPEC (라이브 환경)에서 사용할 수 있으며 CNY 단위는 JON입니다.

---

## 4. 连接与数据映射配置中心

[구성 인터페이스 스크린샷]

**Steps:**
- 시장 정보를 설정합니다. 잘못된 숫자를 설정하면 스크린샷에 표시된 올바른 숫자 범위를 알려드립니다.

---

## 5. 언어 선택

5) 이제 **Merchant's setting page**에서 인터페이스 언어를 영어로 설정할 수 있습니다.

**Language Settings:**
- Merchant setting
- MchInfo
- Interface language: [화살표 표시기가 있는 드롭다운 선택]

---

## 6. 상태 업데이트 페이지

6) "Status Info" 페이지를 확인하십시오. 설정을 확인하고 "Start to upgrade to Global Env" 버튼을 클릭하십시오.

[업그레이드 버튼이 있는 "Status Info" 페이지를 보여주는 스크린샷]

**Details shown:**
- Trade Commission Payment (USD) 2011 and Shop Info
- 다양한 구성 필드 및 설정
- 자세한 약관 텍스트가 있는 확인 대화 상자
- 업그레이드 프로세스 및 요구 사항을 설명하는 여러 단락

---

## 7. 수수료 지불 정보

[결제 인터페이스를 보여주는 스크린샷]

**Trade Commission Payment (USD) 2011 and Shop Info**

표시되는 필드:
- Currency
- CHNC number: [입력 필드]
- Legal entity: [화살표 표시기가 있는 드롭다운 선택]

---

## 참고 사항 및 경고

- API는 24시간 후에 자동으로 모드를 완료하고 API 상태가 온라인 상태로 변경되므로 개발자는 라이브 결제 API 및 라이브 환경에 대한 추가 도움말을 완료합니다.

- 업그레이드를 진행하기 전에 모든 구성 설정이 올바른지 확인하십시오.

- 프로덕션으로 이동하기 전에 샌드박스 환경에서 철저히 테스트하십시오.

- Merchant ID 및 자격 증명을 안전하게 추적하십시오.

---

*개발자 가이드 발췌 종료*

---

# 글로벌 SKU와 Shop SKU 간의 가격 변환

## 대화 상자에 표시되는 단계

**Merchant/Shop Name:** 여기에 표시되는 제한 없음

**Singapore:**
- 1 SGD = 0.00000

**Philippines:**
- 1 PHP = 0.00000

*Successful/UPDATED (중요한 라인 정보)*

**셀 파악:**

---

그런 다음 open api 테스트를 시작할 수 있습니다. 자세한 내용은 [Sandbox Testing](#) 기사를 확인하십시오.

6. 판매자가 KBSC에 로그인하여 각 상점, 제품에 대한 업그레이드를 완료하면 `v2.shop.get_price_info`는 `mslau_upgraded_status:UPGRADED`를 반환합니다.

c. [Global Product API](#) 및 [Merchant API](#)를 테스트하십시오. [여기](#)에서 API 호출 흐름을 확인할 수 있습니다.

---

## 2.6 Merchant & Shop 권한 부여

상점 계정이 KBSC로 업그레이드된 후 여러 상점이 판매자에게 속하므로 이 [FAQ](#)를 통해 Merchant와 상점 간의 관계에 대해 자세히 알아볼 수 있습니다. 따라서 GlobalProduct API를 정상적으로 호출하려면 상점에 권한을 부여하고 권한을 부여할 때 **하위 계정 페이지를 전환**하십시오.

---

# Shopee Openplatform APP에 권한을 부여하기 위해 로그인

**SG ▼** | Email / Phone / Username

**Password**

[비밀번호를 잊으셨습니까?](#)

**[로그인]**

**OR**

**[하위 계정으로 전환]** ← (이것을 가리키는 화살표)

---

**English ▼**

*권한 부여에는 기본 계정을 사용하십시오. 하위 계정은 권한 부여를 완료할 수 없습니다. 기본 계정 형식은 다음과 같습니다. ** main

---

계정이 성공적으로 로그인되면 권한 부여 페이지가 표시됩니다. 먼저 권한을 부여할 상점을 확인한 다음 Auto Auth Merchant를 클릭하십시오. Auto Merchant를 확인하지 않으면 관련 Global Product API를 호출하거나 Merchant 정보를 얻을 수 없습니다. 일부 상점이 선택되지 않은 경우 API를 통해 제품을 관련 상점에 게시할 수 없습니다. 따라서 관리하는 Merchant 및 상점을 완전히 확인하십시오.

[“Auth Merchant” 버튼이 강조 표시되고 “Return all info” 섹션이 표시된 권한 부여 인터페이스를 보여주는 스크린샷]

---

권한 부여가 성공하면 콜백 주소는 shop id 대신 main_account_id를 반환합니다. Main_account_id는 나중에 AccessToken을 얻는 데 사용됩니다.

그런 다음 이 [문서](#)에 따라 토큰을 새로 고치십시오.

GetAccesstoken API를 통해 해당 시점에 성공적으로 권한이 부여된 모든 판매자 ID 및 상점 ID 목록을 얻을 수 있습니다.

**각 판매자 및 각 상점에 대한 토큰은 독립적이며 별도로 저장해야 합니다. Relating-MerchatId 권한 부여를 호출할 때 회사에서 권한을 부여한 판매자 ID 및 shop_id에서 얻은 초기 새로 고침 토큰 및 액세스 토큰을 KnownBy하고 RefreshAccess token API를 호출하면 다른 판매자 ID 또는 다른 상점 ID가 다른 새로 고침 토큰 및 액세스 토큰을 반환하므로 토큰을 별도로 보관하십시오.**

---

판매자 ID와 상점 ID 간의 관계를 얻으려면 `get_shop_list_by_merchant` API를 호출하십시오. `get_merchant_info`를 호출하여 각 판매자 정보를 얻고 `get_shop_info`를 호출하여 각 상점 정보를 얻을 수 있습니다.

---

## 3. 요약

1) 앱을 만들거나 기존 앱을 사용할 때 앱 유형이 v2.0 GlobalProduct API를 호출할 수 있는지 확인하십시오.

2) 상점은 판매자 정보를 제공하기 위해 판매자의 권한이 필요합니다.

3) APP 상태를 호출하기 위한 필수 조건은 판매자 사용자가 프로덕션 환경에서 KBSC/PI를 사용할 수 있다는 것입니다. 상점은 KBSC/PI 판매자 계정 정보를 가질 수 있습니다. APP는 24시간 후에 자동으로 승인됩니다.

4) 판매자는 KBSC에 로그인하여 KBSC 제품 업그레이드가 성공하기 전에 관련 설정을 완료해야 합니다. 제품이 성공적으로 업그레이드된 경우에만 Global Product API를 정상적으로 호출할 수 있습니다.

5) 여기에서 더 많은 GBSC API FAQ를 배울 수 있습니다: [GBSC API FAQ](#)

---

이 문서에서 다루지 않은 기술 연결 문제가 있는 경우 [티켓 시스템](#)을 통해 Shopee Open Platform에 문의하십시오.

## 사용 사례

1. 상품 관리를 위해 Shopee의 KRSC 시스템과 통합.
2. Shopee Open Platform에서 개발자 계정 및 애플리케이션 설정.
3. 라이브 환경으로 전환하기 전에 샌드박스 환경에서 API 호출 테스트.
4. GlobalProduct API 액세스를 위한 상점 및 판매자 인증.
5. KRSC 기능을 지원하도록 기존 앱 업그레이드.

## 관련 API

- Shop OS admin API
- Global Product API
- Merchant API
- v2.shop.get_price_info

---

## 원문 (English)

### Summary

This guide provides instructions for integrating with the KRSC API, including developer registration, app creation, API testing, and environment setup. It covers the steps required to upgrade to KRSC, authorize shops, and use the GlobalProduct API.

### Content

# KRSC API Integration Guide

## Announcement | Console

### Getting Started > KRSC API Integration Guide

Important: KRSC tokens and KRSC kbsp marketing KSP developers (ISVs) need to pay attention one upgrade to KRSC need to pay attention but the product API needs to migrate ubrsc only. The APIs you need to be developer-focused KRSC APIs are generally functional APIs but are also include technical APIs are not navigable. KRSC upgrades which cannot use API calls or Product-facing related modules.

#### FAQ:

**Q: How to judge whether a shop has been upgraded to KRSC?**
A: You can open the shop OS admin API. If you can see the parameter timezhan_at, then (open the response timezhan_at. Except if it is empty, it means that the store has not been upgraded to KRSC.

**Q: If the system does not have product-related function, how to use it can include product creation and product update function interface to do script adjustment?**
A: After upgrading KRSC, the product API will be disabled, so we can use.

---

## 2. How to integrate KRSC API

### 2.1 Register as a developer (If you already have a developer account, you can skip this step)

a. Click "Sign up" and read "Agreement" and register a Shopee Open Platform account at portal: https://open.shopee.com

**Shopee Open Platform**
- Sign up page with email/password fields
- Registration form interface

b. You will receive a verification email, please verify your email and set a password.

---

### Open Platform

**Dear developer,**

You have successfully created your Shopee Open Platform Account.

To verify your email address, please click the button.

[Verify my email]

Best regards,
Shopee Open Platform

---

c. Login with your account

**Shopee Open Platform**
- Sign in
- Email/Password login form

---

### 2.2 Complete account information (If you already have an approved developer account, you can skip this step)

After the process has finished, please complete your developer account info via: Login to Open Platform -> Console -> App Management -> [Your App] -> Complete the developer profile

**Different types of developers will use different types of shops.** For details, please refer to App management page.

a. Complete the appropriate profile according to the of developer account you choose. The information needs to be reviewed by the Shopee platform.

---

### 2.3 Create an app

To call Shopee OpenAPI, you need to create an App first. You can create apps with the developer profile is approved.

**Notes:**
- Original API Type does not support V2.0 api. If your existing app is Original type, please complete the app upgrade.
- To upgrade the app, you can create one through this path.
- To Audit type of app page, please check the SPI permission according to the article.
- After create an app, you can jump through this path.
- Click Public type V2.0 on the App List or Create a new app. After completing the information, Submit.

---

### 2.4 Go live app

You must submit your app status is online. Click "Go Live" to request to use the Live (commercial) APIs from Shopee -> Console -> App List -> [Select Shop] -> Request Go-live or Cancel Go-live request -> Go live.

**Shopee Open Platform Console**

- App Management
  - App List → #1
  - Not yet submit URL address
  - You must link URL
  - Add redirect URL to which users can be redirected...

**Tools**
- Log

**App key**
- Key ID
- Test Shops
- Test Shop
- Test_URL

---

### Basic Information

**Requesting**
- You can track this...
- URI can run
- Test account begin → Singapore
- App can now...

---

**APP IDs automatically complete the module after 24 hours and the APP status will change to Go Online status, and the API can be called normally for production and testing purposes.**

---

# Developer Guide - API Testing and Environment Setup

## Overview

**App ID:** [ID Number]

**Status Information:**
- App ID: [ID]
- Sandbox: running
- App in prod: live
- In prod since: [Date]
- App Secret Status: [Status]
- Live Capabilities: [Details]

---

## 2.5 Start API testing

A `USE CHNC Sandbox` is ready. KPEC Devs can use CHNC Sandbox to run API since the function is similar to KPDC.

To start testing, you should have a `chisa.passportapi.account` on Cpaxide.

### Account Setup Steps

1. **Get Account**
   - User Name: [Enter Name Here]
   - Referrer URL: [URL field]
   - Merchant ID: 100474A

2. **Create Test Account**
   - Please select account type:
     - Options dropdown menu showing:
       - [Account type selection]
       - Create button

3. **Login Process**
   - Login https://sandbox-test.static.example.cn
   - Please note that app is 123456
   - **Note**: Apply with merged account test/password/account if English or if needed

---

## 3. 欢迎来到中国服务中心
## 设置开发者账户并开始测试

[Screenshot showing service center interface with form fields and confirmation dialog]

**Instructions:**
4. Get the currency after login.
   - On the page of PC Web CHNC sandbox, any CNY currency is available to the actual KPEC (Live environment), the unit of CNY is JON.

---

## 4. 连接与数据映射配置中心

[Configuration interface screenshot]

**Steps:**
- Set the market info. If you set the wrong number, we will prompt the correct range of the number at the screenshot shown.

---

## 5. Language Selection

5) Now, you can set the interface language in **Merchant's setting page** to English.

**Language Settings:**
- Merchant setting
- MchInfo
- Interface language: [Dropdown selection with arrow indicator]

---

## 6. Status Update Page

6) Check the "Status Info" page. Please check the setting and click the button "Start to upgrade to Global Env".

[Screenshot showing "Status Info" page with upgrade button]

**Details shown:**
- Trade Commission Payment (USD) 2011 and Shop Info
- Various configuration fields and settings
- Confirmation dialog with detailed terms and conditions text
- Multiple paragraphs explaining upgrade process and requirements

---

## 7. Commission Payment Information

[Screenshot showing payment interface]

**Trade Commission Payment (USD) 2011 and Shop Info**

Fields displayed:
- Currency
- CHNC number: [Input field]
- Legal entity: [Dropdown selection with arrow indicator]

---

## Notes and Warnings

- API will automatically complete the modes after 24 hours and the API status will change to the Online status, so the developer will finish the Live Payment API and Further help of the Live environment.

- Ensure all configuration settings are correct before proceeding with the upgrade.

- Test thoroughly in sandbox environment before moving to production.

- Keep track of your Merchant ID and credentials securely.

---

*End of Developer Guide Extract*

---

# Price Conversion between Global SKU and Shop SKU

## Steps displayed in dialog box

**Merchant/Shop Name:** No limits displayed here

**Singapore:**
- 1 SGD = 0.00000

**Philippines:**
- 1 PHP = 0.00000

*Successful/UPDATED (Critical line info)*

**Figure out the cell:**

---

Then you can start to test open api. More details you can check [Sandbox Testing](#) article.

6. When the seller logs in to KBSC and completes the upgrade for each shops, products, `v2.shop.get_price_info` will return `mslau_upgraded_status:UPGRADED`

c. Test [Global Product API](#) and [Merchant API](#). You can check the API call flow from [here](#)

---

## 2.6 Merchant & Shop Authorization

Because after the shop account is upgraded to KBSC, multiple shops will belong to merchants, you can learn more about the relationship between Merchant and shop through this [FAQ](#). Therefore, in order for you to call the GlobalProduct API normally, please to authorize your shops and **switch the sub account page** when authorizing.

---

# Login to Authorize Shopee Openplatform APP

**SG ▼** | Email / Phone / Username

**Password**

[Forgot Password ?](#)

**[Log In]**

**OR**

**[Switch to Sub Account]** ← (Arrow pointing to this)

---

**English ▼**

*Please use your main account for authorization. Sub-accounts cannot finish the authorization. The account format of the main account is: ** main

---

After the account is successfully logged in, you will see the Authorization page, first check the shop you want to authorize and then click Auto Auth Merchant. If you do not check Auto Merchant, you will not be able to call the relevant Global Product API or to obtain the Merchant information. If some shops are not checked, the product cannot be published to the related shop through the API. So please make sure you check the Merchant and shop that you manage completely.

[Screenshot showing authorization interface with "Auth Merchant" button highlighted and "Return all info" section marked]

---

After the authorization is successful, the callback address will return the main_account_id instead of shop id Main_account_id will be used to obtain AccessToken later.

Then please refresh the token according to this [document](#)

Through the GetAccesstoken API, you will get a list of all merchant ids and shop ids that have been successfully authorized at that time.

**Note that the tokens for each merchant and each shop are independent and you need to store them separately. When you call relating-MerchatId authorization, the initial refresh token and access token, you obtained from the KnownBy the company authorized merchant id and shop_id, and then you call the RefreshAccess token API, and different merchant ids or different shop ids will return different refresh tokens and access token, so please keep their tokens separately.**

---

If you want to obtain the relationship between the merchant id and the shop id, please call the `get_shop_list_by_merchant` API. You can get each merchant information by calling `get_merchant_info` and call `get_shop_info` to get each shop information

---

## 3. Summary

1) When you create an app, or use an existing app, please make sure that your app type can call the v2.0 GlobalProduct API

2) The shop needs the seller's authorization to provide the seller's information

3) The necessary condition for calling the APP status is that the seller users can use the KBSC/PI in the production environment. The shops are allowed to have KBSC/PI seller account information. APP will automatically be approved after 24 hours.

4) The seller needs to log into KBSC and complete the relevant settings before the KBSC product upgrade is successful. Only when the product is successfully upgraded, can the Global Product API be called normally.

5) You can learn more GBSC API FAQ from here: [GBSC API FAQ](#)

---

If you have any technical connection problems that are not covered by this document, please contact Shopee Open Platform through the [ticket system](#).

---

**문서 ID**: developer-guide.29
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/29
**처리 완료**: 2025-10-16T08:28:17
