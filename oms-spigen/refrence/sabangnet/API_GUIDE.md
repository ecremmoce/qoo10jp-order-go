# 사방넷 API 가이드
**추출 시간**: 2025-10-16T16:15:07.729074
**총 시트 수**: 11
**총 API 수**: 431

## 📑 목차
- [상품등록&수정](#상품등록&수정)
- [쇼핑몰 코드 조회](#쇼핑몰-코드-조회)
- [상품정보고시 코드 조회](#상품정보고시-코드-조회)
- [상품요약수정](#상품요약수정)
- [추가상품등록&수정](#추가상품등록&수정)
- [카테고리](#카테고리)
- [주문수집](#주문수집)
- [송장등록](#송장등록)
- [클레임수집](#클레임수집)
- [상품 쇼핑몰별 DATA 수정](#상품-쇼핑몰별-data-수정)
- [문의사항 수집](#문의사항-수집)

---

## 상품등록&수정

**행 수**: 138 | **열 수**: 6

### 📋 컬럼 구조

| 컬럼명 | 데이터 타입 |
|--------|-------------|
| Unnamed: 0 | object |
| Unnamed: 1 | object |
| Unnamed: 2 | object |
| Unnamed: 3 | object |
| Unnamed: 4 | object |
| Unnamed: 5 | object |

### 🔌 API 목록 (136개)

#### API #0

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 상품등록 & 수정을 하기 위한 API로 자체상품코드 기준으로 Insert 또는 Update 가 실행됩니다. |

---

#### API #1

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 사방넷 메뉴 상품관리 >사방넷 상품조회수정 메뉴에서 해당 데이터 확인이 가능합니다. |

---

#### API #2

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 자체상품코드 (COMPAYNY_GOODS_CD) 기준 값으로 등록 또는 수정 진행합니다. |

---

#### API #3

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 웹에서 접근 가능한 URL 또는 IP 주소가 존재 해야 합니다. |

---

#### API #4

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 권장 데이터 수 : 1,000건 |

---

#### API #5

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - DATA 에 ITEM 항목은 CDATA를 사용합니다. (문자열 경우 <![CDATA[ ]]> 사용 필수) |

---

#### API #7

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | * 등록접근주소 : https://sbadminXX.sabangnet.co.kr/RTL_API/xml_goods_info.html?xml_url=상품xml주소 |

---

#### API #8

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | * 상품XML주소 : 아래와 같은 내용으로 구성된 xml파일의 위치경로 (예 : http://www.abc.co.kr/aa.xml ) |

---

#### API #10

**API명**: No

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | No |
| Unnamed: 1 | 구분 |
| Unnamed: 2 | NODE |
| Unnamed: 3 | NODE 설명 |
| Unnamed: 4 | 필수 |
| Unnamed: 5 | 비고 |

---

#### API #11

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | SABANG_GOODS_REGI |

---

#### API #12

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | HEADER |

---

#### API #13

**API명**: 1

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 1 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_COMPAYNY_ID |
| Unnamed: 3 | 사방넷 로그인아이디 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷 어드민 로그인 아이디입력 |

---

#### API #14

**API명**: 2

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 2 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_AUTH_KEY |
| Unnamed: 3 | 사방넷에서 발급한 인증키 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷에서 발급 받은 인증키(만약 받지 않았다면 사방넷에서 인증키 발급 요청) |

---

#### API #15

**API명**: 3

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 3 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_DATE |
| Unnamed: 3 | 전송일자 |
| Unnamed: 4 | Y |
| Unnamed: 5 | YYYYMMDD |

---

#### API #16

**API명**: 4

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 4 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_GOODS_CD_RT |
| Unnamed: 3 | 자체코드 반환여부 |
| Unnamed: 5 | 상품등록/수정 성공시 결과에서 자체코드를 표시합니다. Y:반환, NULL:없음 |

---

#### API #17

**API명**: 5

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 5 |
| Unnamed: 1 | Header |
| Unnamed: 2 | RESULT_TYPE |
| Unnamed: 3 | 결과 메시지 타입 |
| Unnamed: 5 | 미 입력시 기존 HTML형식의 결과값 형태, XML 입력시 XML 형태로 결과값을 출력합니다. |

---

#### API #18

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /HEADER |

---

#### API #19

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | DATA |

---

#### API #20

**API명**: 6

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 6 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | GOODS_NM |
| Unnamed: 3 | 상품명 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 한글기준 50자리까지 사용가능하며 , HTML 태그 사용은 불가합니다. |

---

#### API #21

**API명**: 7

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 7 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | GOODS_KEYWORD |
| Unnamed: 3 | 상품약어 |
| Unnamed: 5 | 간략한 상품명으로써 택배송장 출력과 물류 담당자의 빠른 인식을 위하여 사용할 수 있습니다. ( 단, "NULL"이면 수정안함) |

---

#### API #22

**API명**: 8

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 8 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | MODEL_NM |
| Unnamed: 3 | 모델명 |
| Unnamed: 5 | 상품의 모델명을 정확히 기재합니다. (30자리까지) |

---

#### API #23

**API명**: 9

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 9 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | MODEL_NO |
| Unnamed: 3 | 모델No |
| Unnamed: 5 | 상품의 모델No.를 정확히 기재합니다. ( 30자리까지 ) |

---

#### API #24

**API명**: 10

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 10 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | BRAND_NM |
| Unnamed: 3 | 브랜드명 |
| Unnamed: 5 | 브랜드명을 기재합니다. |

---

#### API #25

**API명**: 11

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 11 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | COMPAYNY_GOODS_CD |
| Unnamed: 3 | 자체상품코드 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 자사에서 사용하는 상품코드를 기재합니다. ( 30자리까지 ) |

---

#### API #26

**API명**: 12

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 12 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | GOODS_SEARCH |
| Unnamed: 3 | 사이트검색어 |
| Unnamed: 5 | 쇼핑몰로 상품정보 전송시 사용될 사이트검색어를 콤마(,)로 구분하여 입력합니다.( 단, "NULL"이면 수정안함) |

---

#### API #27

**API명**: 13

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 13 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | GOODS_GUBUN |
| Unnamed: 3 | 상품구분 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 상품의 구분을 숫자로 입력합니다. 1.위탁상품 2.제조상품 3.사입상품 4.직영상품 |

---

#### API #28

**API명**: 14

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 14 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CLASS_CD1 |
| Unnamed: 3 | 대분류코드 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷에 등록된 대분류코드를 입력합니다.( 단, "NULL"이면 수정안함) |

---

#### API #29

**API명**: 15

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 15 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CLASS_CD2 |
| Unnamed: 3 | 중분류코드 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷에 등록된 중분류코드를 입력합니다.( 단, "NULL"이면 수정안함) |

---

#### API #30

**API명**: 16

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 16 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CLASS_CD3 |
| Unnamed: 3 | 소분류코드 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷에 등록된 소분류코드를 입력합니다.( 단, "NULL"이면 수정안함) |

---

#### API #31

**API명**: 17

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 17 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CLASS_CD4 |
| Unnamed: 3 | 세분류코드 |
| Unnamed: 5 | 사방넷에 등록된 세분류코드를 입력합니다.( 단, "NULL"이면 수정안함) |

---

#### API #32

**API명**: 18

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 18 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PARTNER_ID |
| Unnamed: 3 | 매입처ID |
| Unnamed: 5 | 매입처의 ID를 기재합니다.(대/소문자 정확히 구분해야 함) |

---

#### API #33

**API명**: 19

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 19 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | DPARTNER_ID |
| Unnamed: 3 | 물류처ID |
| Unnamed: 5 | 물류처의 ID를 기재합니다.(대/소문자 정확히 구분해야 함) |

---

#### API #34

**API명**: 20

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 20 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | MAKER |
| Unnamed: 3 | 제조사 |
| Unnamed: 5 | 제조회사의 명칭을 정확히 기재합니다. ( 30자리까지 ) |

---

#### API #35

**API명**: 21

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 21 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | ORIGIN |
| Unnamed: 3 | 원산지(제조국) |
| Unnamed: 4 | Y |
| Unnamed: 5 | 예:중국,사방넷 원산지 표를 참고하시어 표에 기재되어 있는 원산지 명으로 기입해주세요.
원산지가 등록되어 있지 않는 경우 온라인 문의로 요청하시기 바랍니다 ( "NULL" 이거나 ... |

---

#### API #36

**API명**: 22

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 22 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | MAKE_YEAR |
| Unnamed: 3 | 생산연도 |
| Unnamed: 5 | 상품이 생산된 년도를 숫자 4자리로 입력합니다. 예 : 2019 |

---

#### API #37

**API명**: 23

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 23 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | MAKE_DM |
| Unnamed: 3 | 제조일자 |
| Unnamed: 5 | 상품이 제조된 일자를 숫자 8자리로 입력합니다. 예 : 20190820 |

---

#### API #38

**API명**: 24

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 24 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | GOODS_SEASON |
| Unnamed: 3 | 시즌 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 계절의 구분을 숫자로 입력합니다. 1.봄 2.여름 3.가을 4.겨울 5.FW 6.SS 7.해당없음 ( 단, "NULL"이면 수정안함) |

---

#### API #39

**API명**: 25

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 25 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | SEX |
| Unnamed: 3 | 남녀구분 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 남여구분을 숫자로 입력합니다. 1.남성용 2.여성용 3.공용 4.해당없음 ( 단, "NULL"이면 수정안함) |

---

#### API #40

**API명**: 26

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 26 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | STATUS |
| Unnamed: 3 | 상품상태 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 상품의 공급상태에 대한 구분코드를 기재합니다. 1.대기중 2.공급중 3.일시중지 4.완전품절 5.미사용 6.삭제 |

---

#### API #41

**API명**: 27

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 27 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | DELIV_ABLE_REGION |
| Unnamed: 3 | 판매지역 |
| Unnamed: 5 | 판매가능지역을 숫자로 입력합니다. 1.전국 2.전국(도서제외) 3.수도권 4.기타 |

---

#### API #42

**API명**: 28

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 28 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | TAX_YN |
| Unnamed: 3 | 세금구분 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 과세여부를 숫자로 입력합니다. 1.과세 2.면세 3.자료없음 4.비과세 |

---

#### API #43

**API명**: 29

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 29 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | DELV_TYPE |
| Unnamed: 3 | 배송비구분 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 배송비 구분을 숫자로 입력합니다. 1.무료 2.착불 3.선결제 4.착불/선결제 |

---

#### API #44

**API명**: 30

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 30 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | DELV_COST |
| Unnamed: 3 | 배송비 |
| Unnamed: 5 | 배송비를 숫자로 입력합니다. 첫글자는 반드시 '(ENTER좌측Key)로 시작해야하며 숫자사이에 콤마(,)가 들어가면 안됩니다. |

---

#### API #45

**API명**: 31

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 31 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | BANPUM_AREA |
| Unnamed: 3 | 반품지구분 |
| Unnamed: 5 | 매입처의 복수의 반품지중 해당하는 순서를 기재합니다. 예 : 1, 공백일경우 기본주소가 적용됩니다. |

---

#### API #46

**API명**: 32

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 32 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | GOODS_COST |
| Unnamed: 3 | 원가 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 입력시 첫글자는 반드시 ( ‘ ) 아포스트로피(ENTER좌측Key)로 시작해야 하며 숫자 사이에 ( , ) 콤마가 들어가면 안됩니다. |

---

#### API #47

**API명**: 33

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 33 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | GOODS_PRICE |
| Unnamed: 3 | 판매가 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 입력시 첫글자는 반드시 ( ‘ ) 아포스트로피(ENTER좌측Key)로 시작해야 하며 숫자 사이에 ( , ) 콤마가 들어가면 안됩니다. |

---

#### API #48

**API명**: 34

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 34 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | GOODS_CONSUMER_PRICE |
| Unnamed: 3 | TAG가(소비자가) |
| Unnamed: 4 | Y |
| Unnamed: 5 | 입력시 첫글자는 반드시 ( ‘ ) 아포스트로피(ENTER좌측Key)로 시작해야 하며 숫자 사이에 ( , ) 콤마가 들어가면 안됩니다. |

---

#### API #49

**API명**: 35

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 35 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CHAR_1_NM |
| Unnamed: 3 | 옵션제목(1) |
| Unnamed: 5 | 예 : 색상 / 옵션이 없을경우엔는 단품 이라고 입력합니다. |

---

#### API #50

**API명**: 36

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 36 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CHAR_1_VAL |
| Unnamed: 3 | 옵션상세명칭(1) |
| Unnamed: 5 | 옵션항목명1^^재고수량^^추가금액^^별칭^^EA^^옵션공급상태^^옵션모음전여부^^옵션모음전연결상품코드^^안전재고수량
예) 블루^^재고수량^^추가금액^^별칭^^EA,옐로우^^재고수량... |

---

#### API #51

**API명**: 37

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 37 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CHAR_2_NM |
| Unnamed: 3 | 옵션제목(2) |
| Unnamed: 5 | 예 : 사이즈 |

---

#### API #52

**API명**: 38

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 38 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CHAR_2_VAL |
| Unnamed: 3 | 옵션상세명칭(2) |
| Unnamed: 5 | 예 : 44,55,66,77 (각 항목은 반드시 콤마로 구분함,기술된 순서대로 일련된 코드 부여됨) |

---

#### API #53

**API명**: 39

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 39 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH |
| Unnamed: 3 | 대표이미지 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_700.jpg |

---

#### API #54

**API명**: 40

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 40 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH1 |
| Unnamed: 3 | 종합몰(JPG)이미지 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_700.jpg (종합몰(JPG)이미지 (500x500 ~ 700x70... |

---

#### API #55

**API명**: 41

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 41 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH2 |
| Unnamed: 3 | 부가이미지2 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_700.jpg (이미지2) |

---

#### API #56

**API명**: 42

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 42 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH3 |
| Unnamed: 3 | 부가이미지3 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_700.jpg (11번가목록이미지 (300*300)) |

---

#### API #57

**API명**: 43

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 43 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH4 |
| Unnamed: 3 | 부가이미지4 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_700.jpg (이미지4) |

---

#### API #58

**API명**: 44

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 44 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH5 |
| Unnamed: 3 | 부가이미지5 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_700.jpg (이미지5) |

---

#### API #59

**API명**: 45

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 45 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH6 |
| Unnamed: 3 | 부가이미지6 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_700.jpg (쇼핑몰 추가이미지1) |

---

#### API #60

**API명**: 46

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 46 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH7 |
| Unnamed: 3 | 부가이미지7 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_700.jpg (쇼핑몰 추가이미지2) |

---

#### API #61

**API명**: 47

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 47 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH8 |
| Unnamed: 3 | 부가이미지8 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_700.jpg (쇼핑몰 추가이미지3) |

---

#### API #62

**API명**: 48

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 48 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH9 |
| Unnamed: 3 | 부가이미지9 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_700.jpg (쇼핑몰 추가이미지4) |

---

#### API #63

**API명**: 49

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 49 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH10 |
| Unnamed: 3 | 부가이미지10 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_700.jpg (쇼핑몰 추가이미지5) |

---

#### API #64

**API명**: 50

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 50 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH11 |
| Unnamed: 3 | 부가이미지11 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_701.jpg (카페24,메이크샵,고도몰,패플등 리스트이미지) |

---

#### API #65

**API명**: 51

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 51 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH12 |
| Unnamed: 3 | 부가이미지12 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_702.jpg (이미지12) |

---

#### API #66

**API명**: 52

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 52 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH13 |
| Unnamed: 3 | 부가이미지13 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_703.jpg (이미지13) |

---

#### API #67

**API명**: 53

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 53 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH14 |
| Unnamed: 3 | 부가이미지14 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_703.jpg (쇼핑몰 추가이미지6) |

---

#### API #68

**API명**: 54

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 54 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH15 |
| Unnamed: 3 | 부가이미지15 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_703.jpg (쇼핑몰 추가이미지7) |

---

#### API #69

**API명**: 55

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 55 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH16 |
| Unnamed: 3 | 부가이미지16 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_703.jpg (쇼핑몰 추가이미지8) |

---

#### API #70

**API명**: 56

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 56 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH17 |
| Unnamed: 3 | 부가이미지17 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_703.jpg (쇼핑몰 추가이미지9) |

---

#### API #71

**API명**: 57

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 57 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH18 |
| Unnamed: 3 | 부가이미지18 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_703.jpg (쇼핑몰 추가이미지10) |

---

#### API #72

**API명**: 58

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 58 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH19 |
| Unnamed: 3 | 부가이미지19 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_703.jpg (쇼핑몰 추가이미지11) |

---

#### API #73

**API명**: 59

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 59 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH20 |
| Unnamed: 3 | 부가이미지20 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_703.jpg (쇼핑몰 추가이미지12) |

---

#### API #74

**API명**: 60

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 60 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH21 |
| Unnamed: 3 | 부가이미지21 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_703.jpg (쇼핑몰 추가이미지13) |

---

#### API #75

**API명**: 61

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 61 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH22 |
| Unnamed: 3 | 부가이미지22 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_703.jpg (쇼핑몰 추가이미지14) |

---

#### API #76

**API명**: 62

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 62 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH23 |
| Unnamed: 3 | 인증서이미지 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_703.jpg (인증서이미지) |

---

#### API #77

**API명**: 63

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 63 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMG_PATH24 |
| Unnamed: 3 | 수입면장이미지 |
| Unnamed: 5 | 예 : http://gs4333.CO.KR/product_image/a0000769/200907/image20_703.jpg (수입면장이미지) |

---

#### API #78

**API명**: 64

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 64 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | GOODS_REMARKS |
| Unnamed: 3 | 상품상세설명 |
| Unnamed: 5 | 상품상세(HTML)을 기재합니다. |

---

#### API #79

**API명**: 65

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 65 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CERTNO |
| Unnamed: 3 | 인증번호 |
| Unnamed: 5 | 전기용품, 유아 안전용품 등 안전검사를 거쳐야 하는 상품의 경우 해당기관에서 부여한 인증번호를 입력합니다 |

---

#### API #80

**API명**: 66

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 66 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | AVLST_DM |
| Unnamed: 3 | 인증유효 시작일 |
| Unnamed: 5 | 숫자8자리 입력하세요 예:20190820 |

---

#### API #81

**API명**: 67

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 67 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | AVLED_DM |
| Unnamed: 3 | 인증유효 마지막일 |
| Unnamed: 5 | 숫자8자리 입력하세요 예:20190820 |

---

#### API #82

**API명**: 68

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 68 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | ISSUEDATE |
| Unnamed: 3 | 발급일자 |
| Unnamed: 5 | 숫자8자리 입력하세요 예:20190820 |

---

#### API #83

**API명**: 69

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 69 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CERTDATE |
| Unnamed: 3 | 인증일자 |
| Unnamed: 5 | 숫자8자리 입력하세요 예:20190820 |

---

#### API #84

**API명**: 70

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 70 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CERT_AGENCY |
| Unnamed: 3 | 인증기관 |
| Unnamed: 5 | 예 : 한국기업인증연구원 |

---

#### API #85

**API명**: 71

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 71 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CERTFIELD |
| Unnamed: 3 | 인증분야 |
| Unnamed: 5 | 예 : 규격인증 |

---

#### API #86

**API명**: 72

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 72 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | MATERIAL |
| Unnamed: 3 | 식품재료/원산지 |
| Unnamed: 5 | ※식품의 재료와 원산지 구분은 /(슬러시)로 표기하며 추가 입력 시 ,(콤마)로 구분하여 추가할 재료와 원산지를 입력합니다.
판매식품 돼지갈비 예: 갈비/호수산,양념/국내산 |

---

#### API #87

**API명**: 73

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 73 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | STOCK_USE_YN |
| Unnamed: 3 | 재고관리사용여부 |
| Unnamed: 5 | 재고관리 사용여부를 Y or N로 입력합니다.
Y입력시 [재고관리] 메뉴에서 해당상품에 대한 입/출고가 가능하며, 쇼핑몰에 상품연동시 재고수량으로 연동됩니다.
N입력시 [재고관리]... |

---

#### API #88

**API명**: 74

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 74 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | OPT_TYPE |
| Unnamed: 3 | 옵션수정여부 |
| Unnamed: 4 | 2 |
| Unnamed: 5 | 상품수정시 등록된 옵션의 내용을 모두 지우고 새로 등록하는 옵션입니다.
9: 옵션의 내용을 지우지 않는다. 사방넷 재고관리를 이용하는 업체인경우 옵션의 내용을 지우게 되면 기존에 ... |

---

#### API #89

**API명**: 75

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 75 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_EDIT_YN |
| Unnamed: 3 | 속성수정여부 |
| Unnamed: 5 | 속성정보 수정여부를 Y or N로 입력합니다.
Y입력시 속성정보(속성분류코드, 속성값)를 수정 처리합니다. |

---

#### API #90

**API명**: 76

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 76 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP1_CD |
| Unnamed: 3 | 속성분류코드 |
| Unnamed: 5 | 속성분류코드를 숫자 3자리 형식으로 입력합니다.
속성분류코드는 상품속성코드 조회 API나 사방넷 상품관리화면의 속성분류표를 참고하시기 바랍니다.
예: 의류는 001을 입력합니다. |

---

#### API #91

**API명**: 77

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 77 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL1 |
| Unnamed: 3 | 속성값1 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번1인 속성명에 해당하는 속성값을 입력합니다.
속성값(1 ~ 20)은 입력순서대로 처리되므로, 속성순번에 주의하시기 바랍니다.(속성값이 없을 경우, 공란으... |

---

#### API #92

**API명**: 78

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 78 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL2 |
| Unnamed: 3 | 속성값2 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번2인 속성명에 해당하는 속성값을 입력합니다.
예 : 의류 001의 속성명2은 색상이며, 속성값2에 레드,블루 등을 입력합니다. |

---

#### API #93

**API명**: 79

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 79 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL3 |
| Unnamed: 3 | 속성값3 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번3인 속성명에 해당하는 속성값을 입력합니다.
예 : 의류 001의 속성명3은 치수이며, 속성값3에 S,M,L 등을 입력합니다. |

---

#### API #94

**API명**: 80

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 80 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL4 |
| Unnamed: 3 | 속성값4 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번4인 속성명에 해당하는 속성값을 입력합니다.
예 : 의류 001의 속성명4은 제조사(수입자/병행수입)이며, 속성값4에 나이키 등을 입력합니다. |

---

#### API #95

**API명**: 81

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 81 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL5 |
| Unnamed: 3 | 속성값5 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번5인 속성명에 해당하는 속성값을 입력합니다.
예 : 의류 001의 속성명5은 제조국이며, 속성값5에 베트남 등을 입력합니다. |

---

#### API #96

**API명**: 82

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 82 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL6 |
| Unnamed: 3 | 속성값6 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번6인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #97

**API명**: 83

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 83 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL7 |
| Unnamed: 3 | 속성값7 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번7인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #98

**API명**: 84

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 84 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL8 |
| Unnamed: 3 | 속성값8 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번8인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #99

**API명**: 85

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 85 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL9 |
| Unnamed: 3 | 속성값9 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번9인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #100

**API명**: 86

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 86 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL10 |
| Unnamed: 3 | 속성값10 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번10인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #101

**API명**: 87

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 87 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL11 |
| Unnamed: 3 | 속성값11 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번11인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #102

**API명**: 88

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 88 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL12 |
| Unnamed: 3 | 속성값12 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번12인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #103

**API명**: 89

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 89 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL13 |
| Unnamed: 3 | 속성값13 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번13인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #104

**API명**: 90

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 90 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL14 |
| Unnamed: 3 | 속성값14 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번14인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #105

**API명**: 91

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 91 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL15 |
| Unnamed: 3 | 속성값15 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번15인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #106

**API명**: 92

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 92 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL16 |
| Unnamed: 3 | 속성값16 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번16인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #107

**API명**: 93

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 93 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL17 |
| Unnamed: 3 | 속성값17 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번17인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #108

**API명**: 94

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 94 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL18 |
| Unnamed: 3 | 속성값18 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번18인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #109

**API명**: 95

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 95 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL19 |
| Unnamed: 3 | 속성값19 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번19인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #110

**API명**: 96

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 96 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL20 |
| Unnamed: 3 | 속성값20 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번20인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #111

**API명**: 97

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 97 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL21 |
| Unnamed: 3 | 속성값21 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번21인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #112

**API명**: 98

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 98 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL22 |
| Unnamed: 3 | 속성값22 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번22인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #113

**API명**: 99

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 99 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL23 |
| Unnamed: 3 | 속성값23 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번23인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #114

**API명**: 100

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 100 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL24 |
| Unnamed: 3 | 속성값24 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번24인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #115

**API명**: 101

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 101 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL25 |
| Unnamed: 3 | 속성값25 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번25인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #116

**API명**: 102

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 102 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL26 |
| Unnamed: 3 | 속성값26 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번26인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #117

**API명**: 103

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 103 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL27 |
| Unnamed: 3 | 속성값27 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번27인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #118

**API명**: 104

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 104 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL28 |
| Unnamed: 3 | 속성값28 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번28인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #119

**API명**: 105

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 105 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL29 |
| Unnamed: 3 | 속성값29 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번29인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #120

**API명**: 106

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 106 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL30 |
| Unnamed: 3 | 속성값30 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번30인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #121

**API명**: 107

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 107 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL31 |
| Unnamed: 3 | 속성값31 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번31인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #122

**API명**: 108

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 108 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL32 |
| Unnamed: 3 | 속성값32 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번32인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #123

**API명**: 109

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 109 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP_VAL33 |
| Unnamed: 3 | 속성값33 |
| Unnamed: 5 | 속성분류코드에 따른 속성순번33인 속성명에 해당하는 속성값을 입력합니다. |

---

#### API #124

**API명**: 110

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 110 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PACK_CODE_STR |
| Unnamed: 3 | 추가상품그룹코드 |
| Unnamed: 5 | 사방넷에 입력되어 구성된 추가상품의 그룹을 기재합니다. 예 : G001,G004,G201 (7개의 그룹이 입력 가능함) |

---

#### API #125

**API명**: 111

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 111 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | GOODS_NM_EN |
| Unnamed: 3 | 영문 상품명 |
| Unnamed: 5 | 영문 100자리까지 사용가능하며 , HTML 태그 사용은 불가합니다. |

---

#### API #126

**API명**: 112

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 112 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | GOODS_NM_PR |
| Unnamed: 3 | 출력 상품명 |
| Unnamed: 5 | 한글기준 50자리까지 사용가능하며 , HTML 태그 사용은 불가합니다. |

---

#### API #127

**API명**: 113

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 113 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | GOODS_REMARKS2 |
| Unnamed: 3 | 추가 상품상세설명_1 |
| Unnamed: 5 | 상품 추가상세(HTML)을 기재합니다. (단, "DEL" 입력시 저장된 추가상세설명1을 삭제합니다.) |

---

#### API #128

**API명**: 114

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 114 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | GOODS_REMARKS3 |
| Unnamed: 3 | 추가 상품상세설명_2 |
| Unnamed: 5 | 상품 추가상세(HTML)을 기재합니다. (단, "DEL" 입력시 저장된 추가상세설명2를 삭제합니다.) |

---

#### API #129

**API명**: 115

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 115 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | GOODS_REMARKS4 |
| Unnamed: 3 | 추가 상품상세설명_3 |
| Unnamed: 5 | 상품 추가상세(HTML)을 기재합니다. (단, "DEL" 입력시 저장된 추가상세설명3을 삭제합니다.) |

---

#### API #130

**API명**: 116

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 116 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | IMPORTNO |
| Unnamed: 3 | 수입신고번호 |
| Unnamed: 5 | 상품 수입신고번호를 기재합니다. (12345-12-123456U) |

---

#### API #131

**API명**: 117

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 117 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | GOODS_COST2 |
| Unnamed: 3 | 원가2 |
| Unnamed: 5 | 원가2는 상품송신,주문매핑,매출집계,정산등에 이용되지 않으며, 관리상 참고를 위한 가격입니다. |

---

#### API #132

**API명**: 118

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 118 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | ORIGIN2 |
| Unnamed: 3 | 원산지 상세지역 |
| Unnamed: 5 | 원산지 상세 정보를 입력하세요. |

---

#### API #133

**API명**: 119

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 119 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | EXPIRE_DM |
| Unnamed: 3 | 유효일자 |
| Unnamed: 5 | 숫자8자리 입력하세요 예:20190820 |

---

#### API #134

**API명**: 120

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 120 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | SUPPLY_SAVE_YN |
| Unnamed: 3 | 합포제외여부 |
| Unnamed: 5 | 합포장 제외 여부를 Y or N로 입력하세요. "Y" 입력시 합포장 제외 항목에 체크됩니다. |

---

#### API #135

**API명**: 121

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 121 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | DESCRITION |
| Unnamed: 3 | 관리자메모 |
| Unnamed: 5 | 관리자 메모 내용을 입력하세요 |

---

#### API #136

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /DATA |

---

#### API #137

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /SABANG_GOODS_REGI |

---


## 쇼핑몰 코드 조회

**행 수**: 15 | **열 수**: 6

### 📋 컬럼 구조

| 컬럼명 | 데이터 타입 |
|--------|-------------|
| Unnamed: 0 | object |
| Unnamed: 1 | object |
| Unnamed: 2 | object |
| Unnamed: 3 | object |
| Unnamed: 4 | object |
| Unnamed: 5 | object |

### 🔌 API 목록 (12개)

#### API #0

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 사방넷 메뉴 기본정보 >쇼핑몰관리(지원) 에 세팅 되어 있는 거래중인 쇼핑몰을 조회 할 수 있습니다. |

---

#### API #1

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 웹에서 접근 가능한 URL 또는 IP 주소가 존재 해야 합니다. |

---

#### API #3

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | * 등록접근주소 : https://sbadminXX.sabangnet.co.kr/RTL_API/xml_mall_info.html?xml_url=쇼핑몰코드조회xml주소 |

---

#### API #4

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | * XML주소 : 아래와 같은 내용으로 구성된 xml파일의 위치경로 (예 : http://www.abc.co.kr/aa.xml ) |

---

#### API #7

**API명**: No

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | No |
| Unnamed: 1 | 구분 |
| Unnamed: 2 | NODE |
| Unnamed: 3 | NODE 설명 |
| Unnamed: 4 | 필수 |
| Unnamed: 5 | 비고 |

---

#### API #8

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | SABANG_MALL_LIST |

---

#### API #9

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | HEADER |

---

#### API #10

**API명**: 1

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 1 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_COMPAYNY_ID |
| Unnamed: 3 | 사방넷 로그인아이디 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷 어드민 로그인 아이디입력 |

---

#### API #11

**API명**: 2

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 2 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_AUTH_KEY |
| Unnamed: 3 | 사방넷에서 발급한 인증키 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷에서 발급 받은 인증키(만약 받지 않았다면 사방넷에서 인증키 발급 요청) |

---

#### API #12

**API명**: 3

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 3 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_DATE |
| Unnamed: 3 | 전송일자 |
| Unnamed: 4 | Y |
| Unnamed: 5 | YYYYMMDD |

---

#### API #13

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /HEADER |

---

#### API #14

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /SABANG_MALL_LIST |

---


## 상품정보고시 코드 조회

**행 수**: 18 | **열 수**: 6

### 📋 컬럼 구조

| 컬럼명 | 데이터 타입 |
|--------|-------------|
| Unnamed: 0 | object |
| Unnamed: 1 | object |
| Unnamed: 2 | object |
| Unnamed: 3 | object |
| Unnamed: 4 | object |
| Unnamed: 5 | object |

### 🔌 API 목록 (16개)

#### API #0

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 상품정보고시 항목 등록을 위해 제공하는 API 입니다. |

---

#### API #1

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 상품관리 > 사방넷상품속성관리 - 도움자료 에서 속성코드 전체 리스트를 다운 받을 수 있습니다. |

---

#### API #2

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 웹에서 접근 가능한 URL 또는 IP 주소가 존재 해야 합니다. |

---

#### API #4

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | * https://sbadminXX.sabangnet.co.kr/RTL_API/xml_goods_prop_code_info.html?xml_url=상품_속성코드_조회_xml주소 |

---

#### API #5

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | * XML주소 : 아래와 같은 내용으로 구성된 xml파일의 위치경로 (예 : http://www.abc.co.kr/aa.xml ) |

---

#### API #7

**API명**: No

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | No |
| Unnamed: 1 | 구분 |
| Unnamed: 2 | NODE |
| Unnamed: 3 | NODE 설명 |
| Unnamed: 4 | 필수 |
| Unnamed: 5 | 비고 |

---

#### API #8

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | SABANG_GOODS_PROP_CODE_INFO_LIST |

---

#### API #9

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | HEADER |

---

#### API #10

**API명**: 1

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 1 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_COMPAYNY_ID |
| Unnamed: 3 | 사방넷 로그인아이디 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷 어드민 로그인 아이디입력 |

---

#### API #11

**API명**: 2

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 2 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_AUTH_KEY |
| Unnamed: 3 | 사방넷에서 발급한 인증키 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷에서 발급 받은 인증키(만약 받지 않았다면 사방넷에서 인증키 발급 요청) |

---

#### API #12

**API명**: 3

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 3 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_DATE |
| Unnamed: 3 | 전송일자 |
| Unnamed: 4 | Y |
| Unnamed: 5 | YYYYMMDD |

---

#### API #13

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /HEADER |

---

#### API #14

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | DATA |

---

#### API #15

**API명**: 4

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 4 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PROP1_CD |
| Unnamed: 3 | 속성분류코드 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 상품 속성분류코드를 입력합니다. 예 : 의류는 001을 입력합니다. |

---

#### API #16

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /DATA |

---

#### API #17

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /SABANG_GOODS_PROP_CODE_INFO_LIST |

---


## 상품요약수정

**행 수**: 31 | **열 수**: 6

### 📋 컬럼 구조

| 컬럼명 | 데이터 타입 |
|--------|-------------|
| Unnamed: 0 | object |
| Unnamed: 1 | object |
| Unnamed: 2 | object |
| Unnamed: 3 | object |
| Unnamed: 4 | object |
| Unnamed: 5 | object |

### 🔌 API 목록 (29개)

#### API #0

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 상품등록 & 수정 API 항목의 축소판으로 수정이 빈번히 일어나는 항목를 수정하기 위한 API 입니다. |

---

#### API #1

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 재고송신을 위하여 주로 사용됩니다. |

---

#### API #2

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 옵션 수정은 가능하지만 등록 또는 삭제는 불가능합니다. 등록 또는 삭제를 원하실 경우 상품등록 & 수정 API 로 진행합니다. |

---

#### API #3

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 웹에서 접근 가능한 URL 또는 IP 주소가 존재 해야 합니다. |

---

#### API #4

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 기준값은 자체상품코드(COMPAYNY_GOODS_CD) 와 옵션상세명칭을 기준으로 수정이 진행됩니다. |

---

#### API #5

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 옵션을 여러개 수정을 진행 할 경우 SKU_INFO 안에 SKU_VALUE 를 옵션의 갯수 만큼 입력합니다. |

---

#### API #7

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | * 등록접근주소 : https://sbadminXX.sabangnet.co.kr/RTL_API/xml_goods_info2.html?xml_url=상품xml주소 |

---

#### API #8

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | * XML주소 : 아래와 같은 내용으로 구성된 xml파일의 위치경로 (예 : http://www.abc.co.kr/aa.xml ) |

---

#### API #10

**API명**: No

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | No |
| Unnamed: 1 | 구분 |
| Unnamed: 2 | NODE |
| Unnamed: 3 | NODE 설명 |
| Unnamed: 4 | 필수 |
| Unnamed: 5 | 비고 |

---

#### API #11

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | SABANG_GOODS_REGI |

---

#### API #12

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | HEADER |

---

#### API #13

**API명**: 1

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 1 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_COMPAYNY_ID |
| Unnamed: 3 | 사방넷 로그인아이디 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷 어드민 로그인 아이디입력 |

---

#### API #14

**API명**: 2

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 2 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_AUTH_KEY |
| Unnamed: 3 | 사방넷에서 발급한 인증키 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷에서 발급 받은 인증키(만약 받지 않았다면 사방넷에서 인증키 발급 요청) |

---

#### API #15

**API명**: 3

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 3 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_DATE |
| Unnamed: 3 | 전송일자 |
| Unnamed: 4 | Y |
| Unnamed: 5 | YYYYMMDD |

---

#### API #16

**API명**: 4

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 4 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_GOODS_CD_RT |
| Unnamed: 3 | 자체코드 반환여부 |
| Unnamed: 5 | 상품등록/수정 성공시 결과에서 자체코드를 표시합니다. Y:반환, NULL:없음 |

---

#### API #17

**API명**: 5

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 5 |
| Unnamed: 1 | Header |
| Unnamed: 2 | RESULT_TYPE |
| Unnamed: 3 | 결과 메시지 타입 |
| Unnamed: 5 | 미 입력시 기존 HTML형식의 결과값 형태, XML 입력시 XML 형태로 결과값을 출력합니다. |

---

#### API #18

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /HEADER |

---

#### API #19

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | DATA |

---

#### API #20

**API명**: 6

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 6 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | GOODS_NM |
| Unnamed: 3 | 상품명 |
| Unnamed: 5 | 한글기준 50자리까지 사용가능하며 , HTML 태그 사용은 불가합니다. |

---

#### API #21

**API명**: 7

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 7 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | COMPAYNY_GOODS_CD |
| Unnamed: 3 | 자체상품코드 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 자사에서 사용하는 상품코드를 기재합니다. ( 30자리까지 ) |

---

#### API #22

**API명**: 8

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 8 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | STATUS |
| Unnamed: 3 | 상품상태 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 상품의 공급상태에 대한 구분코드를 기재합니다. 1.대기중 2.공급중 3.일시중지 4.완전품절 5.미사용 6.삭제 |

---

#### API #23

**API명**: 9

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 9 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | GOODS_COST |
| Unnamed: 3 | 원가 |
| Unnamed: 5 | 입력시 첫글자는 반드시 ( ‘ ) 아포스트로피(ENTER좌측Key)로 시작해야 하며 숫자 사이에 ( , ) 콤마가 들어가면 안됩니다. |

---

#### API #24

**API명**: 10

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 10 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | GOODS_PRICE |
| Unnamed: 3 | 판매가 |
| Unnamed: 5 | 입력시 첫글자는 반드시 ( ‘ ) 아포스트로피(ENTER좌측Key)로 시작해야 하며 숫자 사이에 ( , ) 콤마가 들어가면 안됩니다. |

---

#### API #25

**API명**: 11

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 11 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | GOODS_CONSUMER_PRICE |
| Unnamed: 3 | TAG가(소비자가) |
| Unnamed: 5 | 입력시 첫글자는 반드시 ( ‘ ) 아포스트로피(ENTER좌측Key)로 시작해야 하며 숫자 사이에 ( , ) 콤마가 들어가면 안됩니다. |

---

#### API #26

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | SKU_INFO |

---

#### API #27

**API명**: 12

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 12 |
| Unnamed: 2 | SKU_VALUE |
| Unnamed: 3 | 수정할 옵션내용 |
| Unnamed: 5 | 옵션상세명칭^^재고수량^^추가금액^^별칭^^EA^^옵션공급상태^^옵션모음전여부^^옵션모음전연결상품코드^^바코드^^안전재고수량
예) 블루:55^^99^^0
1.등록되어 있는옵션수만큼... |

---

#### API #28

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /SKU_INFO |

---

#### API #29

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /DATA |

---

#### API #30

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /SABANG_GOODS_REGI |

---


## 추가상품등록&수정

**행 수**: 34 | **열 수**: 6

### 📋 컬럼 구조

| 컬럼명 | 데이터 타입 |
|--------|-------------|
| Unnamed: 0 | object |
| Unnamed: 1 | object |
| Unnamed: 2 | object |
| Unnamed: 3 | object |
| Unnamed: 4 | object |
| Unnamed: 5 | object |

### 🔌 API 목록 (32개)

#### API #0

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 상품관리 > 사방넷추가상품관리 메뉴의 기능을 api로 제공하는 부분입니다. |

---

#### API #1

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 사방넷에서 사용하는 추가상품코드 (G코드) 기준으로 수정이 진행 됩니다. |

---

#### API #2

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 추가상품코드를 미입력시 신규 등록합니다. |

---

#### API #3

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 추가상품조회 기능은 제공하지 않습니다. |

---

#### API #4

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 웹에서 접근 가능한 URL 또는 IP 주소가 존재 해야 합니다. |

---

#### API #5

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - |

---

#### API #6

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | * 등록접근주소 : https://sbadminXX.sabangnet.co.kr/RTL_API/xml_package_info.html?xml_url=xml주소 |

---

#### API #7

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | * XML주소 : 아래와 같은 내용으로 구성된 xml파일의 위치경로 (예 : http://www.abc.co.kr/aa.xml ) |

---

#### API #10

**API명**: No

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | No |
| Unnamed: 1 | 구분 |
| Unnamed: 2 | NODE |
| Unnamed: 3 | NODE 설명 |
| Unnamed: 4 | 필수 |
| Unnamed: 5 | 비고 |

---

#### API #11

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | SABANG_PACKAGE_REGI |

---

#### API #12

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | HEADER |

---

#### API #13

**API명**: 1

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 1 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_COMPAYNY_ID |
| Unnamed: 3 | 사방넷 로그인아이디 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷 어드민 로그인 아이디입력 |

---

#### API #14

**API명**: 2

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 2 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_AUTH_KEY |
| Unnamed: 3 | 사방넷에서 발급한 인증키 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷에서 발급 받은 인증키(만약 받지 않았다면 사방넷에서 인증키 발급 요청) |

---

#### API #15

**API명**: 3

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 3 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_DATE |
| Unnamed: 3 | 전송일자 |
| Unnamed: 4 | Y |
| Unnamed: 5 | YYYYMMDD |

---

#### API #16

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /HEADER |

---

#### API #17

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | DATA |

---

#### API #18

**API명**: 4

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 4 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PACKAGE_NM |
| Unnamed: 3 | 추가상품그룹명 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 추가상품의 그룹명을 입력합니다.HTML 태그 사용은 불가합니다. |

---

#### API #19

**API명**: 5

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 5 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PACKAGE_CODE |
| Unnamed: 3 | 추가상품그룹코드 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 자사에서 사용하는 추가상품코드를 기재합니다. 미입력시 신규 등록되며, 입력시에는 수정됩니다.(ex:G002) |

---

#### API #20

**API명**: 6

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 6 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PACK_DELIVERY |
| Unnamed: 3 | 배송처 |
| Unnamed: 5 | 배송처를 숫자로 입력합니다. 1:자사, 2:타사 |

---

#### API #21

**API명**: 7

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 7 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | SALE_TYPE |
| Unnamed: 3 | 판매형태 |
| Unnamed: 5 | 판매형태를 숫자로 입력합니다. 1:수탁, 2:사입, 3:기타 |

---

#### API #22

**API명**: 8

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 8 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | MALL_ID |
| Unnamed: 3 | 쇼핑몰CODE |
| Unnamed: 5 | 쇼핑몰코드를 입력합니다. shop으로 시작해야 합니다. 사방넷 메뉴 [기본정보-쇼핑몰관리(지원)]에서 확인 가능합니다. 미입력시 전체몰로 적용됩니다. |

---

#### API #23

**API명**: 9

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 9 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PARTNER_ID |
| Unnamed: 3 | 매입처명 |
| Unnamed: 5 | 매입처의 ID를 기재합니다.(대/소문자 정확히 구분해야 함) |

---

#### API #24

**API명**: 10

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 10 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PACKAGE_GUBUN |
| Unnamed: 3 | 그룹구분 |
| Unnamed: 5 | G:ESM/11번가 기본옵션용, M:ESM/11번가 기본옵션 + 다른몰추가옵션, 미입력시는 일반추가상품으로 적용됩니다. |

---

#### API #25

**API명**: 11

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 11 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | MEMO |
| Unnamed: 3 | 관리자메모 |

---

#### API #26

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | PACK_INFO |

---

#### API #27

**API명**: 12

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 12 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | P_NAME |
| Unnamed: 3 | 명세명(사이트 선택명칭) |
| Unnamed: 4 | Y |
| Unnamed: 5 | 특수문자 / : , | * < >는 사용불가합니다. 또한 쇼핑몰마다 제한하는 특수문자가 존재하므로 특수문자 이용시 주의바랍니다. |

---

#### API #28

**API명**: 13

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 13 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | P_SALE_COST |
| Unnamed: 3 | 판매가격 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 입력시 첫글자는 반드시 ( ‘ ) 아포스트로피(ENTER좌측Key)로 시작해야 하며 숫자 사이에 ( , ) 콤마가 들어가면 안됩니다. |

---

#### API #29

**API명**: 14

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 14 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | P_SALE_STATUS |
| Unnamed: 3 | 상태 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 추가상품의 상태를 숫자로 입력합니다. 2:판매, 4:품절, 6:미사용 |

---

#### API #30

**API명**: 15

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 15 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | P_PRODUCT_ID |
| Unnamed: 3 | 연결할 사방넷상품코드 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷에 등록된 품번코드와 단품코드를 입력합니다. 품번코드-단품코드 형식, 재고를 사용하는 경우 123456-0001, 재고 관리하지 않을 경우 000000-0000 |

---

#### API #31

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /PACK_INFO |

---

#### API #32

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /DATA |

---

#### API #33

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /SABANG_PACKAGE_REGI |

---


## 카테고리

**행 수**: 18 | **열 수**: 6

### 📋 컬럼 구조

| 컬럼명 | 데이터 타입 |
|--------|-------------|
| Unnamed: 0 | object |
| Unnamed: 1 | object |
| Unnamed: 2 | object |
| Unnamed: 3 | object |
| Unnamed: 4 | object |
| Unnamed: 5 | object |

### 🔌 API 목록 (13개)

#### API #0

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 사방넷 메뉴 기본정보 >카테고리관리 에 등록된 카테고리를 조회하는 API 입니다. |

---

#### API #1

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 카테고리 조회 API에서 조회된 코드를 상품에 등록하여 사용합니다. |

---

#### API #2

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 웹에서 접근 가능한 URL 또는 IP 주소가 존재 해야 합니다. |

---

#### API #4

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | * https://sbadminXX.sabangnet.co.kr/RTL_API/xml_category_info.html?xml_url=카테고리조회xml주소 |

---

#### API #5

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | * XML주소 : 아래와 같은 내용으로 구성된 xml파일의 위치경로 (예 : http://www.abc.co.kr/aa.xml ) |

---

#### API #10

**API명**: No

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | No |
| Unnamed: 1 | 구분 |
| Unnamed: 2 | NODE |
| Unnamed: 3 | NODE 설명 |
| Unnamed: 4 | 필수 |
| Unnamed: 5 | 비고 |

---

#### API #11

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | SABANG_CATEGORY_LIST |

---

#### API #12

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | HEADER |

---

#### API #13

**API명**: 1

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 1 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_COMPAYNY_ID |
| Unnamed: 3 | 사방넷 로그인아이디 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷 어드민 로그인 아이디입력 |

---

#### API #14

**API명**: 2

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 2 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_AUTH_KEY |
| Unnamed: 3 | 사방넷에서 발급한 인증키 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷에서 발급 받은 인증키(만약 받지 않았다면 사방넷에서 인증키 발급 요청) |

---

#### API #15

**API명**: 3

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 3 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_DATE |
| Unnamed: 3 | 전송일자 |
| Unnamed: 4 | Y |
| Unnamed: 5 | YYYYMMDD |

---

#### API #16

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /HEADER |

---

#### API #17

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /SABANG_CATEGORY_LIST |

---


## 주문수집

**행 수**: 33 | **열 수**: 6

### 📋 컬럼 구조

| 컬럼명 | 데이터 타입 |
|--------|-------------|
| Unnamed: 0 | object |
| Unnamed: 1 | object |
| Unnamed: 2 | object |
| Unnamed: 3 | object |
| Unnamed: 4 | object |
| Unnamed: 5 | object |

### 🔌 API 목록 (31개)

#### API #0

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 사방넷에 이미 수집되 있는 주문 데이터를 가져가기 위해 사용합니다. |

---

#### API #1

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 주문관리 > 주문서수집(자동) 메뉴를 통하여 주문을 수집 후 주문관리 > 주문서확정관리 메뉴에서 확정처리 한 데이터에 한하여 데이터를 받아갈 수 있습니다. |

---

#### API #2

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - API 요청 시 조건에 해당하는 신규주문 을 주문확인 상태로 변경됩니다. |

---

#### API #3

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 요청할 때마다 조건에 해당하는 모든 데이터를 보내주기에 중복 체크 작업이 필요합니다. |

---

#### API #4

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 사방넷 API에서 유일하게 UTF-8 과 EUC-KR 두가지를 지원 합니다. (UTF-8 사용 시 별도 항목 추가) |

---

#### API #5

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 정렬 조건은 수취인명, 우편번호, 주소, 사방넷 주문번호로 오름차순으로 정렬됩니다. |

---

#### API #6

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 웹에서 접근 가능한 URL 또는 IP 주소가 존재 해야 합니다. |

---

#### API #8

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | * 등록접근주소 : https://sbadminXX.sabangnet.co.kr/RTL_API/xml_order_info.html?xml_url=xml주소 |

---

#### API #9

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | * XML주소 : 아래와 같은 내용으로 구성된 xml파일의 위치경로 (예 : http://www.abc.co.kr/aa.xml ) |

---

#### API #11

**API명**: No

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | No |
| Unnamed: 1 | 구분 |
| Unnamed: 2 | NODE |
| Unnamed: 3 | NODE 설명 |
| Unnamed: 4 | 필수 |
| Unnamed: 5 | 비고 |

---

#### API #12

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | SABANG_ORDER_LIST |

---

#### API #13

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | HEADER |

---

#### API #14

**API명**: 1

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 1 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_COMPAYNY_ID |
| Unnamed: 3 | 사방넷 로그인아이디 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷 어드민 로그인 아이디입력 |

---

#### API #15

**API명**: 2

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 2 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_AUTH_KEY |
| Unnamed: 3 | 사방넷에서 발급한 인증키 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷에서 발급 받은 인증키(만약 받지 않았다면 사방넷에서 인증키 발급 요청) |

---

#### API #16

**API명**: 3

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 3 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_DATE |
| Unnamed: 3 | 전송일자 |
| Unnamed: 4 | Y |
| Unnamed: 5 | YYYYMMDD |

---

#### API #17

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /HEADER |

---

#### API #18

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | DATA |

---

#### API #19

**API명**: 4

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 4 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | ORD_ST_DATE |
| Unnamed: 3 | 검색 시작일자 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷에 수집된 주문의 수집일자 기준의 검색조건 |

---

#### API #20

**API명**: 5

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 5 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | ORD_ED_DATE |
| Unnamed: 3 | 검색 마지막일자 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷에 수집된 주문의 수집일자 기준의 검색조건 |

---

#### API #21

**API명**: 6

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 6 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | ORD_FIELD |
| Unnamed: 3 | 출력필드 리스트 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 출력할 필드를 |의 구분자로 입력합니다.필드리스트는 아래 Field List 표를 참고하여 작성합니다. |

---

#### API #22

**API명**: 7

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 7 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | JUNG_CHK_YN2 |
| Unnamed: 3 | 정산대조확인여부 |
| Unnamed: 5 | Y:확인완료, N:미확인, NULL:전체 |

---

#### API #23

**API명**: 8

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 8 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | ORDER_ID |
| Unnamed: 3 | 주문번호(쇼핑몰) |
| Unnamed: 5 | 쇼핑몰 주문번호를 입력합니다. |

---

#### API #24

**API명**: 9

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 9 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | MALL_ID |
| Unnamed: 3 | 쇼핑몰CODE |
| Unnamed: 5 | 쇼핑몰코드를 입력합니다. shop 또는 chop 로 시작하는 코드를 입력합니다.
기본정보 >쇼핑몰관리(지원) 또는 기본정보 >쇼핑몰관리(일반) 메뉴에서 코드 확인 가능합니다. |

---

#### API #25

**API명**: 10

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 10 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | ORDER_STATUS |
| Unnamed: 3 | 주문상태 |
| Unnamed: 5 | 숫자로 입력합니다. 첫글자는 반드시 '(ENTER좌측Key)로 시작해야합니다. |

---

#### API #26

**API명**: 11

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 11 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | LANG |
| Unnamed: 3 | encoding타입 선택 |
| Unnamed: 5 | UTF-8 타입으로 출력 원할 경우 입력합니다. UTF-8 외에 값이 들어오면 기본값 euc-kr로 출력됩니다. |

---

#### API #27

**API명**: 12

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 12 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | PARTNER_ID |
| Unnamed: 3 | 매입처ID |
| Unnamed: 5 | 매입처 아이디명을 입력합니다. |

---

#### API #28

**API명**: 13

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 13 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | MALL_USER_ID |
| Unnamed: 3 | 쇼핑몰ID |
| Unnamed: 5 | 해당 몰에 해당하는 쇼핑몰 아이디를 입력합니다. MALL_ID 입력시 사용가능합니다. |

---

#### API #29

**API명**: 14

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 14 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | DPARTNER_ID |
| Unnamed: 3 | 물류처ID |
| Unnamed: 5 | 물류처 아이디명을 입력합니다. |

---

#### API #30

**API명**: 15

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 15 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | ACNT_REGS_SRNO |
| Unnamed: 3 | 계정등록순번 |
| Unnamed: 5 | 쇼핑몰로그인ID를 구분하는 고유 번호를 입력합니다. |

---

#### API #31

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /DATA |

---

#### API #32

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /SABANG_ORDER_LIST |

---


## 송장등록

**행 수**: 26 | **열 수**: 6

### 📋 컬럼 구조

| 컬럼명 | 데이터 타입 |
|--------|-------------|
| Unnamed: 0 | object |
| Unnamed: 1 | object |
| Unnamed: 2 | object |
| Unnamed: 3 | object |
| Unnamed: 4 | object |
| Unnamed: 5 | object |

### 🔌 API 목록 (23개)

#### API #0

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 사방넷 주문 중 주문확인, 출고대기, 교환발송준비 상태의 주문에 운송장번호와 택배사 코드를 입력합니다. |

---

#### API #1

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 주문확인 상태의 주문에 운송장을 입력할 경우 출고대기로 상태가 변경됩니다. |

---

#### API #2

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 출고대기 상태이나 쇼핑몰에 이미 송장이 전송 완료된 경우와 강제완료 상태는 송장정보 수정이 불가능합니다. |

---

#### API #3

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 웹에서 접근 가능한 URL 또는 IP 주소가 존재 해야 합니다. |

---

#### API #4

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 사방넷 주문번호 기준으로 반영됩니다. |

---

#### API #6

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | * 등록접근주소 : https://sbadminXX.sabangnet.co.kr/RTL_API/xml_order_invoice.html?xml_url=xml주소 |

---

#### API #7

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | * XML주소 : 아래와 같은 내용으로 구성된 xml파일의 위치경로 (예 : http://www.abc.co.kr/aa.xml ) |

---

#### API #10

**API명**: No

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | No |
| Unnamed: 1 | 구분 |
| Unnamed: 2 | NODE |
| Unnamed: 3 | NODE 설명 |
| Unnamed: 4 | 필수 |
| Unnamed: 5 | 비고 |

---

#### API #11

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | SABANG_INV_REGI |

---

#### API #12

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | HEADER |

---

#### API #13

**API명**: 1

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 1 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_COMPAYNY_ID |
| Unnamed: 3 | 사방넷 로그인아이디 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷 어드민 로그인 아이디입력 |

---

#### API #14

**API명**: 2

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 2 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_AUTH_KEY |
| Unnamed: 3 | 사방넷에서 발급한 인증키 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷에서 발급 받은 인증키(만약 받지 않았다면 사방넷에서 인증키 발급 요청) |

---

#### API #15

**API명**: 3

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 3 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_DATE |
| Unnamed: 3 | 전송일자 |
| Unnamed: 4 | Y |
| Unnamed: 5 | YYYYMMDD |

---

#### API #16

**API명**: 4

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 4 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_INV_EDIT_YN |
| Unnamed: 3 | 송장정보 수정여부 |
| Unnamed: 5 | 주문확인 or 교환발송준비 주문인 경우에만, 송장등록 처리가 가능합니다.
출고대기 주문의 송장정보(택배사코드/운송장번호)를 강제 수정하고 싶은 경우에, Y를 입력합니다. (Y:수정... |

---

#### API #17

**API명**: 5

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 5 |
| Unnamed: 1 | Header |
| Unnamed: 2 | RESULT_TYPE |
| Unnamed: 3 | 결과 메시지 타입 |
| Unnamed: 5 | 미 입력시 기존 HTML형식의 결과값 형태, XML 입력시 XML 형태로 결과값을 출력합니다. |

---

#### API #18

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /HEADER |

---

#### API #19

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | DATA |

---

#### API #20

**API명**: 6

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 6 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | SABANGNET_IDX |
| Unnamed: 3 | 주문번호(사방넷) |
| Unnamed: 4 | Y |

---

#### API #21

**API명**: 7

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 7 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | TAK_CODE |
| Unnamed: 3 | 택배사코드 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 택배사 코드를 조회 하여 입력합니다. |

---

#### API #22

**API명**: 8

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 8 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | TAK_INVOICE |
| Unnamed: 3 | 송장번호 |
| Unnamed: 4 | Y |

---

#### API #23

**API명**: 9

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 9 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | DELV_HOPE_DATE |
| Unnamed: 3 | 배송희망일자 |
| Unnamed: 5 | 배송예정일을 입력합니다.(공백허용) |

---

#### API #24

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /DATA |

---

#### API #25

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /SABANG_INV_REGI |

---


## 클레임수집

**행 수**: 71 | **열 수**: 6

### 📋 컬럼 구조

| 컬럼명 | 데이터 타입 |
|--------|-------------|
| Unnamed: 0 | object |
| Unnamed: 1 | object |
| Unnamed: 2 | object |
| Unnamed: 3 | object |
| Unnamed: 4 | object |
| Unnamed: 5 | object |

### 🔌 API 목록 (65개)

#### API #0

**API명**: - 사방넷에 이미 수집되 있는 클레임 데이터를 가져가기 위해 사용합니다.

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | - 사방넷에 이미 수집되 있는 클레임 데이터를 가져가기 위해 사용합니다. |

---

#### API #1

**API명**: - 자동 수집된 클레임 건에 한해서 지원되며, 사방넷에서 수동으로 입력된 클레임 내용은 API를 통해 지원되지 않습니다.

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | - 자동 수집된 클레임 건에 한해서 지원되며, 사방넷에서 수동으로 입력된 클레임 내용은 API를 통해 지원되지 않습니다. |

---

#### API #2

**API명**: - 클레임 수집 내용만 조회 할 수 있는 API 입니다. 처리는 사방넷 또는 쇼핑몰에서 직접 수동으로 진행해야 합니다.

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | - 클레임 수집 내용만 조회 할 수 있는 API 입니다. 처리는 사방넷 또는 쇼핑몰에서 직접 수동으로 진행해야 합니다. |

---

#### API #3

**API명**: - 웹에서 접근 가능한 URL 또는 IP 주소가 존재 해야 합니다.

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | - 웹에서 접근 가능한 URL 또는 IP 주소가 존재 해야 합니다. |

---

#### API #4

**API명**: - 요청할 때마다 조건에 해당하는 모든 데이터를 보내주기에 중복 체크 작업이 필요합니다.

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | - 요청할 때마다 조건에 해당하는 모든 데이터를 보내주기에 중복 체크 작업이 필요합니다. |

---

#### API #6

**API명**: * 등록접근주소 : https://sbadminXX.sabangnet.co.kr/RTL_API/xml_clm_info.html?xml_url=xml주소

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | * 등록접근주소 : https://sbadminXX.sabangnet.co.kr/RTL_API/xml_clm_info.html?xml_url=xml주소 |

---

#### API #7

**API명**: * XML주소 : 아래와 같은 내용으로 구성된 xml파일의 위치경로 (예 : http://www.abc.co.kr/aa.xml )

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | * XML주소 : 아래와 같은 내용으로 구성된 xml파일의 위치경로 (예 : http://www.abc.co.kr/aa.xml ) |

---

#### API #10

**API명**: No

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | No |
| Unnamed: 1 | 구분 |
| Unnamed: 2 | NODE |
| Unnamed: 3 | NODE 설명 |
| Unnamed: 4 | 필수 |
| Unnamed: 5 | 비고 |

---

#### API #11

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | SABANG_ORDER_LIST |

---

#### API #12

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | HEADER |

---

#### API #13

**API명**: 1

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 1 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_COMPAYNY_ID |
| Unnamed: 3 | 사방넷 로그인아이디 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷 어드민 로그인 아이디입력 |

---

#### API #14

**API명**: 2

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 2 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_AUTH_KEY |
| Unnamed: 3 | 사방넷에서 발급한 인증키 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷에서 발급 받은 인증키(만약 받지 않았다면 사방넷에서 인증키 발급 요청) |

---

#### API #15

**API명**: 3

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 3 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_DATE |
| Unnamed: 3 | 전송일자 |
| Unnamed: 4 | Y |
| Unnamed: 5 | YYYYMMDD |

---

#### API #16

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /HEADER |

---

#### API #17

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | DATA |

---

#### API #18

**API명**: 4

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 4 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CLM_ST_DATE |
| Unnamed: 3 | 검색 시작일자 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷에 수집된 클레임의 수집일자 기준의 검색조건 |

---

#### API #19

**API명**: 5

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 5 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CLM_ED_DATE |
| Unnamed: 3 | 검색 마지막일자 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷에 수집된 클레임의 수집일자 기준의 검색조건 |

---

#### API #20

**API명**: 6

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 6 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CLM_FIELD |
| Unnamed: 3 | 출력필드 리스트 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 출력할 필드를 | 의 구분자로 입력합니다.필드리스트는 아래 Field List 표를 참고하여 작성합니다. |

---

#### API #21

**API명**: 7

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 7 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | LANG |
| Unnamed: 3 | encoding타입 선택 |
| Unnamed: 5 | UTF-8 타입으로 출력 원할 경우 입력합니다. UTF-8 외에 값이 들어오면 기본값 euc-kr로 출력됩니다. |

---

#### API #22

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /DATA |

---

#### API #23

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /SABANG_ORDER_LIST |

---

#### API #26

**API명**: Field List

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | Field List |

---

#### API #27

**API명**: - Request 전문에 CLM_FIELD 항목에 아래 항목 중 받고 싶은 항목을 | (세로바) 기준으로 입력합니다.

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | - Request 전문에 CLM_FIELD 항목에 아래 항목 중 받고 싶은 항목을 | (세로바) 기준으로 입력합니다. |

---

#### API #28

**API명**: - CLAME_STATUS_GUBUN 항목은 쇼핑몰에서 수집된 구분값이 입력 됩니다. 텍스트로 수집되기에 사용 하는 쇼핑몰마다 다를 수 있습니다.

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | - CLAME_STATUS_GUBUN 항목은 쇼핑몰에서 수집된 구분값이 입력 됩니다. 텍스트로 수집되기에 사용 하는 쇼핑몰마다 다를 수 있습니다. |

---

#### API #30

**API명**: 필드명

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 필드명 |
| Unnamed: 1 | 필드설명 |
| Unnamed: 2 | 비고 |

---

#### API #31

**API명**: IDX

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | IDX |
| Unnamed: 1 | 주문번호(사방넷) |
| Unnamed: 2 | 주문 등록(수집시) 사방넷에서 부여되는 유니크번호 |

---

#### API #32

**API명**: ORDER_ID

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | ORDER_ID |
| Unnamed: 1 | 주문번호(쇼핑몰) |
| Unnamed: 2 | 주문수집시 쇼핑몰의 주문번호 |

---

#### API #33

**API명**: MALL_ID

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | MALL_ID |
| Unnamed: 1 | 쇼핑몰명 |
| Unnamed: 2 | 수집된 주문의 쇼핑몰명 |

---

#### API #34

**API명**: MALL_USER_ID

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | MALL_USER_ID |
| Unnamed: 1 | 쇼핑몰ID |
| Unnamed: 2 | 수집된 주문의 쇼핑몰 로그인아이디 |

---

#### API #35

**API명**: ORDER_STATUS

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | ORDER_STATUS |
| Unnamed: 1 | 주문상태 |
| Unnamed: 2 | 사방넷 주문상태 |

---

#### API #36

**API명**: USER_ID

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | USER_ID |
| Unnamed: 1 | 주문자ID |

---

#### API #37

**API명**: USER_NAME

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | USER_NAME |
| Unnamed: 1 | 주문자명 |

---

#### API #38

**API명**: USER_TEL

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | USER_TEL |
| Unnamed: 1 | 주문자전화번호 |

---

#### API #39

**API명**: USER_CEL

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | USER_CEL |
| Unnamed: 1 | 주문자핸드폰번호 |

---

#### API #40

**API명**: USER_EMAIL

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | USER_EMAIL |
| Unnamed: 1 | 주문자이메일주소 |

---

#### API #41

**API명**: RECEIVE_TEL

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | RECEIVE_TEL |
| Unnamed: 1 | 수취인전화번호 |

---

#### API #42

**API명**: RECEIVE_CEL

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | RECEIVE_CEL |
| Unnamed: 1 | 수취인핸드폰번호 |

---

#### API #43

**API명**: RECEIVE_EMAIL

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | RECEIVE_EMAIL |
| Unnamed: 1 | 수취인이메일주소 |

---

#### API #44

**API명**: DELV_MSG

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | DELV_MSG |
| Unnamed: 1 | 배송메세지 |
| Unnamed: 2 | 쇼핑몰에서 수집된 배송메세지 |

---

#### API #45

**API명**: RECEIVE_NAME

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | RECEIVE_NAME |
| Unnamed: 1 | 수취인명 |

---

#### API #46

**API명**: RECEIVE_ZIPCODE

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | RECEIVE_ZIPCODE |
| Unnamed: 1 | 수취인우편번호 |

---

#### API #47

**API명**: RECEIVE_ADDR

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | RECEIVE_ADDR |
| Unnamed: 1 | 수취인주소 |

---

#### API #48

**API명**: TOTAL_COST

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | TOTAL_COST |
| Unnamed: 1 | 주문금액 |
| Unnamed: 2 | 주문자가 쇼핑몰에 주문한 주문 총금액(상품의 판매단가*수량) |

---

#### API #49

**API명**: PAY_COST

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | PAY_COST |
| Unnamed: 1 | 결제금액 |
| Unnamed: 2 | 주문자가 쇼핑몰에 결재한 실제금액 |

---

#### API #50

**API명**: ORDER_DATE

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | ORDER_DATE |
| Unnamed: 1 | 주문일자 |

---

#### API #51

**API명**: PARTNER_ID

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | PARTNER_ID |
| Unnamed: 1 | 매입처명 |

---

#### API #52

**API명**: DPARTNER_ID

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | DPARTNER_ID |
| Unnamed: 1 | 물류처명 |

---

#### API #53

**API명**: MALL_PRODUCT_ID

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | MALL_PRODUCT_ID |
| Unnamed: 1 | 상품코드(쇼핑몰) |

---

#### API #54

**API명**: PRODUCT_ID

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | PRODUCT_ID |
| Unnamed: 1 | 품번코드(사방넷) |

---

#### API #55

**API명**: SKU_ID

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | SKU_ID |
| Unnamed: 1 | 단품코드(사방넷) |

---

#### API #56

**API명**: P_PRODUCT_NAME

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | P_PRODUCT_NAME |
| Unnamed: 1 | 상품명(확정) |

---

#### API #57

**API명**: P_SKU_VALUE

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | P_SKU_VALUE |
| Unnamed: 1 | 옵션(확정) |

---

#### API #58

**API명**: PRODUCT_NAME

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | PRODUCT_NAME |
| Unnamed: 1 | 상품명(수집) |

---

#### API #59

**API명**: SALE_COST

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | SALE_COST |
| Unnamed: 1 | 판매가(수집) |
| Unnamed: 2 | 쇼핑몰에 판매하는 상품의 1개 금액 |

---

#### API #60

**API명**: MALL_WON_COST

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | MALL_WON_COST |
| Unnamed: 1 | 공급단가 |
| Unnamed: 2 | 쇼핑몰에서 업체에게 정산해야할 상품 1개 금액 |

---

#### API #61

**API명**: WON_COST

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | WON_COST |
| Unnamed: 1 | 원가 |
| Unnamed: 2 | 상품등록시 입력한 상품원가 |

---

#### API #62

**API명**: SKU_VALUE

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | SKU_VALUE |
| Unnamed: 1 | 옵션(수집) |

---

#### API #63

**API명**: SALE_CNT

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | SALE_CNT |
| Unnamed: 1 | 수량 |

---

#### API #64

**API명**: COMPAYNY_GOODS_CD

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | COMPAYNY_GOODS_CD |
| Unnamed: 1 | 자체상품코드 |
| Unnamed: 2 | 상품등록시에 입력한 자사 상품코드 |

---

#### API #65

**API명**: CLAME_STATUS_GUBUN

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | CLAME_STATUS_GUBUN |
| Unnamed: 1 | 클레임 구분 |
| Unnamed: 2 | 교환/반품/취소 구분 |

---

#### API #66

**API명**: CLAME_CONTENT

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | CLAME_CONTENT |
| Unnamed: 1 | 클레임 내용 |
| Unnamed: 2 | 쇼핑몰에 접수된 클레임 내용 |

---

#### API #67

**API명**: CLAME_INS_DATE

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | CLAME_INS_DATE |
| Unnamed: 1 | 클레임 접수일자 |
| Unnamed: 2 | 쇼핑몰에 접수된 클레임 접수(요청)일자 또는 사방넷의 클레임 수집일자 |

---

#### API #68

**API명**: CLAME_REG_DATE

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | CLAME_REG_DATE |
| Unnamed: 1 | 클레임 수집일자 |
| Unnamed: 2 | 사방넷의 클레임 수집일자 |

---

#### API #69

**API명**: CL_IDX

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | CL_IDX |
| Unnamed: 1 | 클레임 번호 |
| Unnamed: 2 | 클레임 수집시 사방넷에서 부여되는 유니크번호 |

---

#### API #70

**API명**: MALL_ORDER_ID

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | MALL_ORDER_ID |
| Unnamed: 1 | 부주문번호 |

---


## 상품 쇼핑몰별 DATA 수정

**행 수**: 37 | **열 수**: 6

### 📋 컬럼 구조

| 컬럼명 | 데이터 타입 |
|--------|-------------|
| Unnamed: 0 | object |
| Unnamed: 1 | object |
| Unnamed: 2 | object |
| Unnamed: 3 | object |
| Unnamed: 4 | object |
| Unnamed: 5 | object |

### 🔌 API 목록 (33개)

#### API #0

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 품번안의 상품명, 판매가, 상세설명 등 고정값이 아닌 쇼핑몰마다 다르게 입력해서 송신해야 하는 경우 사용합니다. |

---

#### API #1

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 쇼핑몰별 재고분할 경우 자체상품코드 기준으로 입력된 값의 합이 100보다 적게 입력해야 합니다. 기존값 + 추가 입력값이 100이 넘을 경우 실패됩니다. |

---

#### API #2

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 쇼핑몰코드는 사방넷 메뉴 기본정보 > 쇼핑몰관리(지원) 메뉴를 참고하거나 쇼핑몰코드 조회 API를 참고해서 입력합니다. |

---

#### API #3

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | - 웹에서 접근 가능한 URL 또는 IP 주소가 존재 해야 합니다. |

---

#### API #5

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | * 등록접근주소 : https://sbadminXX.sabangnet.co.kr/RTL_API/xml_goods_info3.html?xml_url=xml주소 |

---

#### API #6

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 1 | * XML주소 : 아래와 같은 내용으로 구성된 xml파일의 위치경로 (예 : http://www.abc.co.kr/aa.xml ) |

---

#### API #10

**API명**: No

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | No |
| Unnamed: 1 | 구분 |
| Unnamed: 2 | NODE |
| Unnamed: 3 | NODE 설명 |
| Unnamed: 4 | 필수 |
| Unnamed: 5 | 비고 |

---

#### API #11

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | SABANG_GOODS_REGI |

---

#### API #12

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | HEADER |

---

#### API #13

**API명**: 1

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 1 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_COMPAYNY_ID |
| Unnamed: 3 | 사방넷 로그인아이디 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷 어드민 로그인 아이디입력 |

---

#### API #14

**API명**: 2

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 2 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_AUTH_KEY |
| Unnamed: 3 | 사방넷에서 발급한 인증키 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷에서 발급 받은 인증키(만약 받지 않았다면 사방넷에서 인증키 발급 요청) |

---

#### API #15

**API명**: 3

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 3 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_DATE |
| Unnamed: 3 | 전송일자 |
| Unnamed: 4 | Y |
| Unnamed: 5 | YYYYMMDD |

---

#### API #16

**API명**: 4

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 4 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_GOODS_CD_RT |
| Unnamed: 3 | 자체코드 반환여부 |
| Unnamed: 5 | 상품등록/수정 성공시 결과에서 자체코드를 표시합니다. Y:반환, NULL:없음 |

---

#### API #17

**API명**: 5

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 5 |
| Unnamed: 1 | Header |
| Unnamed: 2 | RESULT_TYPE |
| Unnamed: 3 | 결과 메시지 타입 |
| Unnamed: 5 | 미 입력시 기존 HTML형식의 결과값 형태, XML 입력시 XML 형태로 결과값을 출력합니다. |

---

#### API #18

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /HEADER |

---

#### API #19

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | DATA |

---

#### API #20

**API명**: 6

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 6 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | MALL_CODE |
| Unnamed: 3 | 쇼핑몰CODE |
| Unnamed: 4 | Y |
| Unnamed: 5 | 쇼핑몰 코드를 기재합니다. (사방넷메뉴 A>2 쇼핑몰관리(지원) 메뉴 참조) |

---

#### API #21

**API명**: 7

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 7 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | COMPAYNY_GOODS_CD |
| Unnamed: 3 | 자체상품코드 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 자사에서 사용하는 상품코드를 기재합니다. ( 30자리까지 ) |

---

#### API #22

**API명**: 8

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 8 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | MALL_GOODS_PRICE |
| Unnamed: 3 | 판매가 |
| Unnamed: 5 | 입력시 첫글자는 반드시 ( ‘ ) 아포스트로피(ENTER좌측Key)로 시작해야 하며 숫자 사이에 ( , ) 콤마가 들어가면 안됩니다. |

---

#### API #23

**API명**: 9

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 9 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | MALL_PROD_NAME |
| Unnamed: 3 | 상품명 |
| Unnamed: 5 | 한글기준 50자리까지 사용가능하며 , HTML 태그 사용은 불가합니다. |

---

#### API #24

**API명**: 10

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 10 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | MALL_PROD_DESC |
| Unnamed: 3 | 상품상세설명 |
| Unnamed: 5 | 상품상세(HTML)을 기재합니다. |

---

#### API #25

**API명**: 11

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 11 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | MALL_PROP1_CD |
| Unnamed: 3 | 속성분류코드 |
| Unnamed: 5 | 속성분류코드를 숫자 3자리 형식으로 입력합니다.
속성분류코드는 상품속성코드 조회 API나 사방넷 상품관리화면의 속성분류표를 참고하시기 바랍니다.
예: 의류는 001을 입력합니다. |

---

#### API #26

**API명**: 12

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 12 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | BUINFO_ID3 |
| Unnamed: 3 | 부가정보Ⅱ코드 |
| Unnamed: 5 | 사방넷메뉴 쇼핑몰 관리>> 쇼핑몰 부가정보Ⅱ 메뉴에서 등록한 부가정보Ⅱ 코드값을 입력합니다. 예) B0000001 |

---

#### API #27

**API명**: 13

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 13 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | MALL_STOCK_RATE |
| Unnamed: 3 | 재고분할 퍼센트 |
| Unnamed: 5 | 숫자로 입력합니다. 품번에 설정된 쇼핑몰별 재고분할 합이 100 이하로 입력되어야 합니다. |

---

#### API #28

**API명**: 14

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 14 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CERTNO |
| Unnamed: 3 | 인증번호 |
| Unnamed: 5 | 인증일자를 입력합니다. 8자리로 입력합니다. 예)20190901 |

---

#### API #29

**API명**: 15

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 15 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | ISSUEDATE |
| Unnamed: 3 | 발급일자 |
| Unnamed: 5 | 발급일자를 입력합니다. 8자리로 입력합니다. 예)20190901 |

---

#### API #30

**API명**: 16

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 16 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CERTDATE |
| Unnamed: 3 | 인증일자 |
| Unnamed: 5 | 인증일자를 입력합니다. 8자리로 입력합니다. 예)20190901 |

---

#### API #31

**API명**: 17

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 17 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | AVLST_DM |
| Unnamed: 3 | 인증유효 시작일 |
| Unnamed: 5 | 유효기간 시작일을 입력합니다. 8자리로 입력합니다. 예)20190901 |

---

#### API #32

**API명**: 18

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 18 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | AVLED_DM |
| Unnamed: 3 | 인증유효 마지막일 |
| Unnamed: 5 | 유효기간 종료일을 입력합니다. 8자리로 입력합니다. 예)20190901 |

---

#### API #33

**API명**: 19

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 19 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CERT_AGENCY |
| Unnamed: 3 | 인증기관 |
| Unnamed: 5 | 인증검사기관을 입력합니다. |

---

#### API #34

**API명**: 20

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 20 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CERTFIELD |
| Unnamed: 3 | 인증분야 |
| Unnamed: 5 | 인증분야를 입력합니다. |

---

#### API #35

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /DATA |

---

#### API #36

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /SABANG_GOODS_REGI |

---


## 문의사항 수집

**행 수**: 45 | **열 수**: 6

### 📋 컬럼 구조

| 컬럼명 | 데이터 타입 |
|--------|-------------|
| Unnamed: 0 | object |
| Unnamed: 1 | object |
| Unnamed: 2 | object |
| Unnamed: 3 | object |
| Unnamed: 4 | object |
| Unnamed: 5 | object |

### 🔌 API 목록 (41개)

#### API #0

**API명**: - 사방넷에 이미 수집되어 있는 쇼핑몰의 문의 사항을 수집하기 위한 API 입니다.

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | - 사방넷에 이미 수집되어 있는 쇼핑몰의 문의 사항을 수집하기 위한 API 입니다. |

---

#### API #1

**API명**: - 웹에서 접근 가능한 URL 또는 IP 주소가 존재 해야 합니다.

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | - 웹에서 접근 가능한 URL 또는 IP 주소가 존재 해야 합니다. |

---

#### API #3

**API명**: * 등록접근주소 : https://sbadminXX.sabangnet.co.kr/RTL_API/xml_cs_info.html?xml_url=xml주소

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | * 등록접근주소 : https://sbadminXX.sabangnet.co.kr/RTL_API/xml_cs_info.html?xml_url=xml주소 |

---

#### API #4

**API명**: * XML주소 : 아래와 같은 내용으로 구성된 xml파일의 위치경로 (예 : http://www.abc.co.kr/aa.xml )

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | * XML주소 : 아래와 같은 내용으로 구성된 xml파일의 위치경로 (예 : http://www.abc.co.kr/aa.xml ) |

---

#### API #6

**API명**: No

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | No |
| Unnamed: 1 | 구분 |
| Unnamed: 2 | NODE |
| Unnamed: 3 | NODE 설명 |
| Unnamed: 4 | 필수 |
| Unnamed: 5 | 비고 |

---

#### API #7

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | SABANG_CS_LIST |

---

#### API #8

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | HEADER |

---

#### API #9

**API명**: 1

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 1 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_COMPAYNY_ID |
| Unnamed: 3 | 사방넷 로그인아이디 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷 어드민 로그인 아이디입력 |

---

#### API #10

**API명**: 2

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 2 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_AUTH_KEY |
| Unnamed: 3 | 사방넷에서 발급한 인증키 |
| Unnamed: 4 | Y |
| Unnamed: 5 | 사방넷에서 발급 받은 인증키(만약 받지 않았다면 사방넷에서 인증키 발급 요청) |

---

#### API #11

**API명**: 3

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 3 |
| Unnamed: 1 | Header |
| Unnamed: 2 | SEND_DATE |
| Unnamed: 3 | 전송일자 |
| Unnamed: 4 | Y |
| Unnamed: 5 | YYYYMMDD |

---

#### API #12

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /HEADER |

---

#### API #13

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | DATA |

---

#### API #14

**API명**: 4

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 4 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CS_ST_DATE |
| Unnamed: 3 | 검색 시작일자 |
| Unnamed: 4 | Y |
| Unnamed: 5 | YYYYMMDD : 수집일자 기준 검색 시작일 |

---

#### API #15

**API명**: 5

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 5 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CS_ED_DATE |
| Unnamed: 3 | 검색 마지막일자 |
| Unnamed: 4 | Y |
| Unnamed: 5 | YYYYMMDD : 수집일자 기준 검색 종료일 |

---

#### API #16

**API명**: 6

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 6 |
| Unnamed: 1 | ITEM |
| Unnamed: 2 | CS_STATUS |
| Unnamed: 3 | 검색 처리구분 |
| Unnamed: 5 | 001:신규접수, 002:답변저장, 003:답변전송, 004:강제완료, NULL:전체 |

---

#### API #17

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /DATA |

---

#### API #18

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 2 | /SABANG_CS_LIST |

---

#### API #20

**API명**: Field List

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | Field List |

---

#### API #21

**API명**: - 호출 후 리턴 되는 항목 리스트입니다.

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | - 호출 후 리턴 되는 항목 리스트입니다. |

---

#### API #22

**API명**: - CS_GUBUN 항목은 고정값이 아닌 쇼핑몰에서 주는 텍스트 항목입니다.

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | - CS_GUBUN 항목은 고정값이 아닌 쇼핑몰에서 주는 텍스트 항목입니다. |

---

#### API #24

**API명**: 필드명

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | 필드명 |
| Unnamed: 1 | 필드설명 |
| Unnamed: 2 | 비고 |

---

#### API #25

**API명**: NUM

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | NUM |
| Unnamed: 1 | 사방넷 일련번호 |
| Unnamed: 2 | 문의 등록(수집시) 사방넷에서 부여되는 유니크번호 |

---

#### API #26

**API명**: MALL_ID

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | MALL_ID |
| Unnamed: 1 | 쇼핑몰명 |
| Unnamed: 2 | 수집된 문의사항의 쇼핑몰명 |

---

#### API #27

**API명**: MALL_USER_ID

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | MALL_USER_ID |
| Unnamed: 1 | 쇼핑몰ID |
| Unnamed: 2 | 수집된 문의사항의 쇼핑몰 로그인아이디 |

---

#### API #28

**API명**: CS_STATUS

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | CS_STATUS |
| Unnamed: 1 | 처리구분 |
| Unnamed: 2 | 001:신규접수, 002:답변저장, 003:답변전송, 004:강제완료 |

---

#### API #29

**API명**: REG_DM

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | REG_DM |
| Unnamed: 1 | 수집일자 |
| Unnamed: 2 | 수집일자 |

---

#### API #30

**API명**: ORDER_ID

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | ORDER_ID |
| Unnamed: 1 | 주문번호(쇼핑몰) |
| Unnamed: 2 | 쇼핑몰 주문번호 |

---

#### API #31

**API명**: PRODUCT_ID

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | PRODUCT_ID |
| Unnamed: 1 | 품번코드(사방넷) |

---

#### API #32

**API명**: MALL_PROD_ID

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | MALL_PROD_ID |
| Unnamed: 1 | 상품코드(쇼핑몰) |
| Unnamed: 2 | 쇼핑몰 상품코드 |

---

#### API #33

**API명**: PRODUCT_NM

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | PRODUCT_NM |
| Unnamed: 1 | 상품명 |

---

#### API #34

**API명**: SUBJECT

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | SUBJECT |
| Unnamed: 1 | 문의제목 |

---

#### API #35

**API명**: CNTS

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | CNTS |
| Unnamed: 1 | 문의내용 |

---

#### API #36

**API명**: INS_NM

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | INS_NM |
| Unnamed: 1 | 작성자 |
| Unnamed: 2 | 질문자 |

---

#### API #37

**API명**: INS_DM

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | INS_DM |
| Unnamed: 1 | 고객등록일자 |
| Unnamed: 2 | 질문일자 |

---

#### API #38

**API명**: RPLY_CNTS

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | RPLY_CNTS |
| Unnamed: 1 | 답변 및 안내 |

---

#### API #39

**API명**: UPD_NM

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | UPD_NM |
| Unnamed: 1 | 답변자 |

---

#### API #40

**API명**: UPD_DM

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | UPD_DM |
| Unnamed: 1 | 답변저장일자 |

---

#### API #41

**API명**: SEND_DM

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | SEND_DM |
| Unnamed: 1 | 답변전송일자 |

---

#### API #42

**API명**: CS_GUBUN

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | CS_GUBUN |
| Unnamed: 1 | 문의구분 |
| Unnamed: 2 | 문의구분 : 상품문의, Q&A,긴급메시지,... |

---

#### API #43

**API명**: COMPAYNY_GOODS_CD

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | COMPAYNY_GOODS_CD |
| Unnamed: 1 | 자체상품코드 |
| Unnamed: 2 | 품번코드가 존재 할 경우 해당 품번의 자체상품코드 |

---

#### API #44

**API명**: MALL_CODE

**상세 정보**:

| 항목 | 값 |
|------|----|
| Unnamed: 0 | MALL_CODE |
| Unnamed: 1 | 쇼핑몰 코드 |
| Unnamed: 2 | 수집된 CS의 쇼핑몰 코드 |

---


