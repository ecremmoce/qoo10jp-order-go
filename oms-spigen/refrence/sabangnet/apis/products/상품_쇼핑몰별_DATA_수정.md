# 상품 쇼핑몰별 DATA 수정
**분류**: products

## 개요

- 품번안의 상품명, 판매가, 상세설명 등 고정값이 아닌 쇼핑몰마다 다르게 입력해서 송신해야 하는 경우 사용합니다.

- 쇼핑몰별 재고분할 경우 자체상품코드 기준으로 입력된 값의 합이 100보다 적게 입력해야 합니다. 기존값 + 추가 입력값이 100이 넘을 경우 실패됩니다.

- 쇼핑몰코드는 사방넷 메뉴 기본정보 > 쇼핑몰관리(지원) 메뉴를 참고하거나 쇼핑몰코드 조회 API를 참고해서 입력합니다.

- 웹에서 접근 가능한 URL 또는 IP 주소가 존재 해야 합니다.

## 사방넷 원본 엔드포인트

```
* 등록접근주소 : https://sbadminXX.sabangnet.co.kr/RTL_API/xml_goods_info3.html?xml_url=xml주소
```

## RESTful API 변환

**Method**: `POST`

**Path**: `/api/v1/상품 쇼핑몰별 data 수정`

**Summary**: 상품 쇼핑몰별 DATA 수정

## 인증 (Authentication)

### 사방넷 XML 헤더

| 필드명 | 설명 | 필수 |
|--------|------|------|
| SEND_COMPAYNY_ID | 사방넷 로그인아이디 | Y |
| SEND_AUTH_KEY | 사방넷에서 발급한 인증키 | Y |

### RESTful HTTP 헤더

```http
Authorization: Bearer {API_KEY}
X-Company-ID: {COMPANY_ID}
Content-Type: application/json
```

## 헤더 필드 (Header Fields)

| 필드명 | 설명 | 필수 | 비고 |
|--------|------|------|------|
| SEND_COMPAYNY_ID | 사방넷 로그인아이디 | Y | 사방넷 어드민 로그인 아이디입력 |
| SEND_AUTH_KEY | 사방넷에서 발급한 인증키 | Y | 사방넷에서 발급 받은 인증키(만약 받지 않았다면 사방넷에서 인증키 발급 요청) |
| SEND_DATE | 전송일자 | Y | YYYYMMDD |
| SEND_GOODS_CD_RT | 자체코드 반환여부 | N | 상품등록/수정 성공시 결과에서 자체코드를 표시합니다. Y:반환, NULL:없음 |
| RESULT_TYPE | 결과 메시지 타입 | N | 미 입력시 기존 HTML형식의 결과값 형태, XML 입력시 XML 형태로 결과값을 출력합니다. |

## 데이터 필드 (Data Fields)

| XML 필드명 | JSON 필드명 | 설명 | 필수 | 비고 |
|-----------|-------------|------|------|------|
| MALL_CODE | mallCode | 쇼핑몰CODE | Y | 쇼핑몰 코드를 기재합니다. (사방넷메뉴 A>2 쇼핑몰관리(지원) 메뉴 참조) |
| COMPAYNY_GOODS_CD | compaynyGoodsCd | 자체상품코드 | Y | 자사에서 사용하는 상품코드를 기재합니다. ( 30자리까지 ) |
| MALL_GOODS_PRICE | mallGoodsPrice | 판매가 | N | 입력시 첫글자는 반드시 ( ‘ ) 아포스트로피(ENTER좌측Key)로 시작해야 하며 숫자 사이에 ( , ) 콤마가 들어가면 안됩니다. |
| MALL_PROD_NAME | mallProdName | 상품명 | N | 한글기준 50자리까지 사용가능하며 , HTML 태그 사용은 불가합니다. |
| MALL_PROD_DESC | mallProdDesc | 상품상세설명 | N | 상품상세(HTML)을 기재합니다. |
| MALL_PROP1_CD | mallProp1Cd | 속성분류코드 | N | 속성분류코드를 숫자 3자리 형식으로 입력합니다.
속성분류코드는 상품속성코드 조회 API나 사방넷 상품관리화면의 속성분류표를 참고하시기 바랍니다.
예: 의류는 001을 입력합니다. |
| BUINFO_ID3 | buinfoId3 | 부가정보Ⅱ코드 | N | 사방넷메뉴 쇼핑몰 관리>> 쇼핑몰 부가정보Ⅱ 메뉴에서 등록한 부가정보Ⅱ 코드값을 입력합니다. 예) B0000001 |
| MALL_STOCK_RATE | mallStockRate | 재고분할 퍼센트 | N | 숫자로 입력합니다. 품번에 설정된 쇼핑몰별 재고분할 합이 100 이하로 입력되어야 합니다. |
| CERTNO | certno | 인증번호 | N | 인증일자를 입력합니다. 8자리로 입력합니다. 예)20190901 |
| ISSUEDATE | issuedate | 발급일자 | N | 발급일자를 입력합니다. 8자리로 입력합니다. 예)20190901 |
| CERTDATE | certdate | 인증일자 | N | 인증일자를 입력합니다. 8자리로 입력합니다. 예)20190901 |
| AVLST_DM | avlstDm | 인증유효 시작일 | N | 유효기간 시작일을 입력합니다. 8자리로 입력합니다. 예)20190901 |
| AVLED_DM | avledDm | 인증유효 마지막일 | N | 유효기간 종료일을 입력합니다. 8자리로 입력합니다. 예)20190901 |
| CERT_AGENCY | certAgency | 인증기관 | N | 인증검사기관을 입력합니다. |
| CERTFIELD | certfield | 인증분야 | N | 인증분야를 입력합니다. |

## XML 구조 예제

```xml
<SABANG_GOODS_REGI>
  <HEADER>
    <SEND_COMPAYNY_ID>사방넷 로그인아이디</SEND_COMPAYNY_ID>
    <SEND_AUTH_KEY>사방넷에서 발급한 인증키</SEND_AUTH_KEY>
    <SEND_DATE>전송일자</SEND_DATE>
  </HEADER>
  <DATA>
    <ITEM>
      <MALL_CODE>값</MALL_CODE>
      <COMPAYNY_GOODS_CD>값</COMPAYNY_GOODS_CD>
      <MALL_GOODS_PRICE>값</MALL_GOODS_PRICE>
      <MALL_PROD_NAME>값</MALL_PROD_NAME>
      <MALL_PROD_DESC>값</MALL_PROD_DESC>
    </ITEM>
  </DATA>
</SABANG_GOODS_REGI>
```

## JSON 구조 예제

```json
{
  "mallCode": "값",
  "compaynyGoodsCd": "값",
  "mallGoodsPrice": "값",
  "mallProdName": "값",
  "mallProdDesc": "값"
}
```
