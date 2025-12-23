# API 호출

**카테고리**: 시작하기
**난이도**: 중간
**중요도**: 5/5
**최종 업데이트**: 2025-10-16T08:05:09

## 개요

본 가이드는 Shopee Open API v2.0에 대한 API 호출 방법에 대한 정보를 제공합니다. API 도메인, 요청 메서드, 프로토콜 및 파라미터(GET 및 POST 요청에 대한 공통 및 요청 파라미터 포함)를 다룹니다.

## 주요 키워드

- API 호출
- Shopee Open API
- API 도메인
- 요청 메서드
- API 프로토콜
- 요청 파라미터
- GET
- POST

## 본문

# API 호출

**최종 업데이트: 2025-09-22**

지원 언어: English / 简体中文 / 繁體中文 / Português / ไทย

⚠️ **참고:** 이 가이드는 Shopee Open API v2.0에 대한 API 호출에만 적용됩니다.

## API 도메인

사용 가능한 도메인은 3개입니다.

### 프로덕션 환경
- https://openplatform.shopee.cn/ — 중국 본토 근처에 서비스를 배포한 개발자용
- https://openplatform.shopee.com.br/ — 미국 근처에 서비스를 배포한 개발자용
- https://partner.shopeemobile.com/ — 싱가포르 근처에 서비스를 배포한 개발자용

### 샌드박스 환경
- https://partner.test-stable.shopeemobile.com/ — 모든 개발자용

Open API에 액세스하는 서버 위치에 따라 올바른 도메인 이름을 선택하십시오.

## API 요청 메서드

현재 Open API는 GET 및 POST의 두 가지 요청 메서드만 제공합니다.

## API 프로토콜

대부분의 API는 HTTP/JSON을 사용합니다. 파일 업로드를 위한 API와 같이 특정 API의 경우 HTTP/FORM을 사용합니다.

## API 요청 파라미터

API 문서에는 다음과 같은 두 가지 유형의 요청 파라미터가 있습니다.

1. 공통 파라미터 (Common parameter)
2. 요청 파라미터 (Request parameter)

GET 유형 API의 경우 이러한 두 파라미터가 동시에 존재하거나 공통 파라미터만 존재할 수 있습니다.

POST 유형 API의 경우 이러한 두 파라미터가 동시에 존재합니다.

---

## 탐색 링크 (오른쪽 사이드바)

- API 도메인
- API 요청 메서드
- API 프로토콜
- API 요청 파라미터
- Open API 유형
- 서명 계산
- API 요청 샘플
- API 응답 파라미터
- API 기능
- API 관련 문제

---

## 왼쪽 사이드바 - 목차

### 시작하기
- 소개
- 개발자 계정 등록
- 앱 관리
- **API 호출** (현재 페이지)
- 푸시 메커니즘 알림
- 인증 및 권한 부여
- 샌드박스 테스트 V2
- 서비스 파트너 프로그램
- V2.0 API 호출 흐름
- CNSC API 통합 가이드
- KRSC API 통합 가이드
- V2.0 데이터 정의

### API 모범 사례
- 제품 생성 지침
- 제품 생성 준비
- 제품 생성
- 글로벌 제품 생성
- 글로벌 제품 게시
- 옵션 관리
- 제품 기본 정보 관리
- 재고 및 가격 관리
- 주문 관리
- First Mile 바인딩

## 사용 사례

1. Shopee의 Open API와 통합
2. Shopee에 API 요청하기
3. API 요청 메서드 이해하기
4. API 파라미터 처리하기
5. 올바른 API 도메인 선택하기

## 관련 API

- Shopee Open API v2.0

---

## 원문 (English)

### Summary

This guide provides information on making API calls for Shopee Open API v2.0. It covers API domains, request methods, protocols, and parameters, including common and request parameters for GET and POST requests.

### Content

# API calls

**Last Updated: 2025-09-22**

Language Supported: English / 简体中文 / 繁體中文 / Português / ไทย

⚠️ **Note:** This guide only applies to making API calls for Shopee Open API v2.0.

## API domains

There are 3 domains available:

### Production environment
- https://openplatform.shopee.cn/ — for developers who deployed their services near Chinese Mainland.
- https://openplatform.shopee.com.br/ — for developers who deployed their services near US.
- https://partner.shopeemobile.com/ — for developer who deployed their services near SG.

### Sandbox environment
- https://partner.test-stable.shopeemobile.com/ — for All developers

Select the correct domain name, based on the server location you're accessing the Open API from.

## API request methods

Currently, Open API only provides two request methods: GET and POST.

## API Protocol

HTTP/JSON for most APIs. HTTP/FORM for some certain APIs, for example, the API for uploading files.

## API request parameters

In the API document, you will see two types of request parameters:

1. Common parameter
2. Request parameter

For GET-type APIs, these two parameters may exist at the same time, or only the common parameter will exist.

For POST-type APIs, these two parameters will exist at the same time.

---

## Navigation Links (Right Sidebar)

- API domains
- API request methods
- API Protocol
- API request parameters
- Types of Open API
- Signature calculation
- API request sample
- API response parameters
- API functions
- Problems with API

---

## Left Sidebar - Tables of Content

### Getting Started
- Introduction
- Developer account registration
- App management
- **API calls** (current page)
- Push Mechanism notifications
- Authorization and Authentication
- Sandbox Testing V2
- Service Partner Program
- V2.0 API Call Flow
- CNSC API Integration Guide
- KRSC API Integration Guide
- V2.0 Data Definition

### API Best Practices
- Guidelines for Creating Product
- Product creation preparation
- Creating product
- Creating global product
- Publishing global product
- Variant management
- Product base info management
- Stock & Price Management
- Order Management
- First Mile Binding

---

**문서 ID**: developer-guide.16
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/16
**처리 완료**: 2025-10-16T08:05:09
