# API 접근 도구

**카테고리**: 통합
**난이도**: medium
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T09:11:04

## 개요

본 가이드는 API 로그 및 푸시 로그 도구를 액세스하고 사용하는 방법에 대한 지침을 제공합니다. 로그 쿼리, API 관련 문제에 대한 티켓 개설, Shopee Open Platform 내에서 푸시 알림 문제 해결 단계를 자세히 설명합니다.

## 주요 키워드

- API Log
- Push Log
- API 접근
- 문제 해결
- 티켓 제출
- 로그 쿼리
- 오류 메시지
- Shopee Open Platform

## 본문

# API 액세스 도구

## 공지 | 콘솔

### 탐색
API 모범 사례 > Brand Open API 모범 사례... > Ferramentas de Log da Op...

---

## API 로그 액세스

[여기를 클릭](link)하여 도구에 액세스하십시오.

API 로그 도구에 대한 액세스 권한은 계약 기간 동안 Live 파트너에게만 독점적으로 부여됩니다. 시스템 속성을 수정하는 요청은 API 운영에 직접적인 영향을 미칠 수 있으므로 신중하게 수행해야 합니다. 적절하다고 판단되는 요청은 로그에 정확하게 전달될 수 있습니다.

1. 요청 및 응답 필드 각각에 대한 별도의 쿼리(분할 쿼리)를 타임스탬프 순으로 정렬하여 로그에 정확하게 보낼 수 있습니다.
2. 단어별 키워드 검색. 예를 들어 메시지 기록에서 order_sn 문자열을 검색할 수 있습니다.
3. 각 로그 항목의 전체 세부 정보를 볼 수 있습니다.
4. shop_id와 관련된 모든 관련 정보가 티켓 개설 양식에 자동으로 복사됩니다.
5. 문서 위에 마우스를 올리면 View API Document 페이지로 리디렉션되어 해당 페이지에서 왔는지 확인할 수 있습니다.
6. API 문서는 단일 페이지에서 제공됩니다.

---

## API 로그 도구를 통한 호출 열기

1) 파트너를 선택한 후 로그를 확인해야 합니다.
2) Request ID가 있는 경우 Request ID 필드를 사용하여 직접 검색합니다. Request ID가 없는 경우 키워드 필드를 사용하여 원하는 필드를 검색합니다.
3) API Path 필드를 사용하여 과거 로그를 쿼리할 수 있습니다.
4) 검색을 클릭하여 모든 로그 매개변수 스크리닝 페이지에서 원하는 결과를 찾을 수 있습니다.
5) 티켓 올리기를 클릭합니다.
6) 호출 개설 양식으로 리디렉션된 경우 다음과 같은 더 많은 관련 정보가 표시됩니다.

### 필수 필드:
1. Developer Email
2. L1 Question Category
3. L2 Question Category
4. Question Description
5. API Version
6. API Category
7. API Name
8. Request
9. Response
10. Request ID

데이터가 포함된 Question Description 필드의 존재 여부에 따라 달라집니다.

---

## Shopee Open Platform에 티켓 올리기

**문제 해결에 도움이 되도록 정확한 정보를 제출해 주십시오.**

### 양식 필드:

**\* Developer Email**
[입력 필드]

**\* L1 Question Category**
[드롭다운: Live API Issue]

**\* L2 Question Category**
[드롭다운: Product]

**\* Question Description**
[입력용 대형 텍스트 영역]

---

9) 필요한 경우 관련 파일을 첨부하여 티켓을 열 수 있습니다. 파일 섹션에 추가하기만 하면 됩니다.

**\* Developer Email**
[입력 필드]

**\* L1 Question Category**
[드롭다운: Live API Issue]

**\* L2 Question Category**
[드롭다운: Product]

**\* Question Description**
[텍스트 영역]

**파일**
[업로드하려면 클릭 (0/5)]

---

10) 이러한 작업을 수행한 후 제출을 클릭하고 Captcha를 완료합니다.
11) 호출이 성공적으로 제출됩니다.

---

## 푸시 로그 액세스

[여기를 클릭](link)하여 도구에 액세스하십시오.

1. 계약 기간 동안 Live 파트너에 해당하는 푸시 로그를 쿼리하여 푸시 요청이 성공적으로 수행되도록 할 수 있습니다.
2. 푸시 메시지가 성공적으로 전송된 경우 오류 메시지와 이유가 표시됩니다.

---

### 선택 방법:

**참고:** 콜백을 받지 못한 경우 오류 메시지 "Shopee has sent the push to your callback url, but we didn't get an response with 2xx code within the requested timeout seconds, please check your API/V2."

3) "The request sent to your callback url can't be completed successfully with 2xx code within the timeout seconds"와 같은 오류가 발생하면 API/V2를 확인하십시오.

4) 기타 오류 메시지: "Receiving other error."

---

### 키워드 검색:

단어별로 검색할 수 있습니다. 예를 들어 푸시 메시지에서 order_sn 문자열을 검색하여 푸시 로그를 찾을 수 있습니다.

엔터티는 모든 필수 필드( "\*"로 표시된 필드)를 설명할 수 있습니다.

### 푸시 로그 필드:
1. Partner ID
2. Push Mechanism
3. Send Time (UTC+03:00) - 데이터가 전송된 날짜 이후 (최근 7일 동안 사용 가능)

특정 키워드 필드를 검색하려면 키워드 필드(예: order_sn 또는 shop_id)를 사용하십시오.

---

**문서 종료**

## 사용 사례

1. API 통합 문제 디버깅
2. 푸시 알림 실패 문제 해결
3. 과거 API 로그 쿼리
4. API 문제에 대한 지원 티켓 제출
5. 성공적인 푸시 요청 확인

## 관련 API

- Live API
- API/V2

---

## 원문 (English)

### Summary

This guide provides instructions on how to access and use the API Log and Push Log tools. It details the steps for querying logs, opening tickets for API-related issues, and troubleshooting push notification problems within the Shopee Open Platform.

### Content

# API Access Tools

## Announcement | Console

### Navigation
API Best Practices > Brand Open API Best Prac... > Ferramentas de Log da Op...

---

## API Log Access

Access the tool by [clicking here](link)

Access to the API Log tool is the exclusive right of the Live Partner during the contract period. Requests to modify system attributes must be made with caution, as they may directly impact the operation of the API. Requests deemed appropriate can be forwarded accurately to the log.

1. A separate query (split query) for each of the request and response fields, sorted by timestamp, can be sent accurately to the log.
2. Keyword search by word. For example, you can search for the string of order_sn in the message record.
3. You can view the complete details of each log entry.
4. All relevant information related to the shop_id will be automatically copied to the ticket opening form.
5. Hovering over the documentation will redirect the user to the View API Document page, where you can check if it's from this page.
6. API documentation is provided on a single page.

---

## Opening a Call Through the API Log Tool

1) After selecting a partner, you must verify the logs.
2) If you have a Request ID, search directly using the Request ID field. If you don't have a Request ID, use the Keywords field to search for the desired field.
3) The API Path field can be used to query historical logs.
4) You can click on Search to find the desired results on all log parameter screening pages.
5) Click on Raise a Ticket.
6) If you've been redirected to our call opening form, you'll see more relevant information, such as:

### Required Fields:
1. Developer Email
2. L1 Question Category
3. L2 Question Category
4. Question Description
5. API Version
6. API Category
7. API Name
8. Request
9. Response
10. Request ID

Based on the presence of the Question Description field with its data.

---

## Raise Ticket to Shopee Open Platform

**Please submit accurate information to help us troubleshoot the issue.**

### Form Fields:

**\* Developer Email**
[Input field]

**\* L1 Question Category**
[Dropdown: Live API Issue]

**\* L2 Question Category**
[Dropdown: Product]

**\* Question Description**
[Large text area for input]

---

9) If necessary, you can attach relevant files to open the ticket, just add them in the File section:

**\* Developer Email**
[Input field]

**\* L1 Question Category**
[Dropdown: Live API Issue]

**\* L2 Question Category**
[Dropdown: Product]

**\* Question Description**
[Text area]

**File**
[Click to Upload (0/5)]

---

10) After performing these actions, click Submit and complete the Captcha.
11) Your call will be successfully submitted.

---

## Push Log Access

Access the tool by [clicking here](link)

1. You can query the Push logs corresponding to your Live Partner during the contract period, so that push requests are successful.
2. If the push message was sent successfully, the error message and reason will be displayed.

---

### How the Selection Works:

**Note:** If you don't receive a callback, the error message "Shopee has sent the push to your callback url, but we didn't get an response with 2xx code within the requested timeout seconds, please check your API/V2."

3) If you receive an error like "The request sent to your callback url can't be completed successfully with 2xx code within the timeout seconds", please check your API/V2.

4) Other error messages: "Receiving other error."

---

### Keyword Search:

You can search by word. For example, you can search for the string of order_sn in the push message to look for the Push log.

Entities can describe all mandatory fields (fields marked with "\*"):

### Push Log Fields:
1. Partner ID
2. Push Mechanism
3. Send Time (UTC+03:00) - after the date the data was sent (available from the last 7 days)

To search for a specific keyword field, use the Keywords field, for example order_sn or shop_id.

---

**End of Document**

---

**문서 ID**: developer-guide.381
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/381
**처리 완료**: 2025-10-16T09:11:04
