# API 모범 사례 - Brand tyson API 모범 사례...

**카테고리**: 모범 사례
**난이도**: medium
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T09:13:56

## 개요

이 가이드는 Brand Tyson API 사용에 대한 모범 사례를 제공하며, 간단한 등록 및 NF-a 배출량 계산에 대한 단계별 지침을 포함합니다. 데이터 유효성 검사, 송장 생성, 오류 처리, 주문 및 송장 관리를 위한 특정 endpoint 사용법을 다룹니다.

## 주요 키워드

- API 모범 사례
- 송장 생성
- NF-a 배출
- 데이터 유효성 검사
- endpoint 사용법
- 주문 관리
- 오류 처리
- API integration

## 본문

# API 모범 사례 - Brand tyson API 모범 사례... > 간단한 등록 단계별 가이드

## 간단한 등록 단계별 가이드

데이터는 "teplcode_colonial_failed" 파라미터의 상세 값을 사용하여 **ir_pelas_por_ender_modal** 엔드포인트를 통해 얻을 수 있습니다. 그러나 데이터 보호 정책만 확인하려면 다음 문서 링크를 참조하십시오: [link]

### 1. ICPF 유효성 검사기

다음 데이터는 "READY_TO_SHIP", "PROCESSED" 및 "DELIVERED" 상태와 NF-e의 "orderDate" 데이터를 포함하기 위한 반환 "fieldnames"로 사용됩니다. 다음 구조에 따라 "Name", "Sumcom"의 "Validated" 및 구매자의 "ICPF"는 NF-e에서 찾을 수 없으므로 어떤 주문 상태에서도 사용할 수 없습니다.

| Car Seller | MPDGE | WAITING_SELLER_HANDLING | PAYMENT | COMMITMENT | CANCELED | RETURNED_RECEIVER |
|------------|-------|-------------------------|---------|------------|----------|-------------------|
| Name | Mask | Mask | Mask | Mask | Mask | Mask |
| Surname | Mask | Mask | Mask | Mask | Mask | Mask |
| Address | Mask | Show | Mask | Mask | Mask | Show |
| CPF | Mask | Mask | Mask | Mask | Mask | Mask |

### 2. CNPJ 급여

다음 데이터는 구매자의 "fieldnames" "Name", "Delivered" 및 "CPF"가 "INVOICE_PENDING", "VALID_INVOICE", "IN_TRANSIT" 및 "RETURNED_TO_SELLER" 상태를 포함하기 위해 사용되며, NF-e에서 찾을 수 없으므로 어떤 주문 상태에서도 사용할 수 없습니다.

구매자의 "Validated"는 필요하지 않으므로 어떤 주문 상태에서도 사용할 수 없습니다.

| Car Seller | MPDGE | WAITING_SELLER_HANDLING | PAYMENT | COMMITMENT | CANCELED | RETURNED_RECEIVER |
|------------|-------|-------------------------|---------|------------|----------|-------------------|
| Name | Mask | Show | Mask | Mask | Mask | Show |
| Surname | Mask | Show | Mask | Mask | Mask | Show |
| Address | Mask | Show | Mask | Mask | Mask | Show |
| CPF | Mask | Show | Mask | Mask | Mask | Show |

이러한 데이터 자산에 대한 가장 유용한 작업은 프로세스의 고유한 리소스에 대한 데이터가 필수가 아닌 주문에서 수행되었습니다.

---

## NF-a 발급 계산

NF 발급 작업은 다음 계산 단계를 통해 **ir_tacimat_int_torme_detail** 엔드포인트를 통해 수행할 수 있습니다.

NF-a를 계산할 때 "virgint_price", "seller_discount", "fieldname_item_voucher_seller", "buyer_paid_shipping_fee"가 발견됩니다.

```json
{
  "total_calc(x)": "request GET 'https://partner.shopeemx.io.mx/api/v2/order/get_order_list'",
  "Redirect 200": {
    "parm": "Content-Type: application/json || Authorization: Bearer eyJkh3k...FWdL1...JMJN-Ws"
  },
  "fsp-frequet118 api.h2m55expemz.io.mx^59660-etc_price.0002718_JMN-Ws": {}
}
```

**알림**

요청이 필요합니다:

```json
{
  "error": "",
  "message": "",
  "response": {},
  "buyer": {
    "id": "...",
    "total_calc(x)": ""
  },
  "config_list": "",
  "order_sn": "23302022222222",
  "order_sn": "23302022222222",
  "order_sn": "23302022222222",
  "request_id": "?????????"
}
```

---

## NF-a가 필요한 주문 확인

1. **ir_pelas_por_ender_modal** 엔드포인트를 호출합니다. 응답에서 "optional "valid_status" 파라미터는 "INVOICE_PENDING" 값이 됩니다. 따라서 모든 주문 번호에 대해 추가 NF-a를 진행하려면 좋은 주문이 필요합니다.

응답 데이터는 cURL 형식입니다:

```bash
curl --location --request GET "https://partner.shopeemx.io.mx/api/v2/order/get_order_list?order_list=abc.def||
ghilmno.pqr.stu.v.wx.yz.ABCD.EFG.HIJ.KLM.NOP.QRST.UV.WXYZ...[]^_`...
partnerId=2920000X&signature(x)=20^59660-etc_price.0002718_JMN-Ws"
--header "Content-Type: application/json||
--header "Authorization: Bearer eyJkh3k...FWdL1...JMJN-Ws"
```

**알림**

NF-a를 확인하고 발급된 주문은 되돌릴 수 없습니다 (계산은 디렉터리에 있음).

이 "response_optional_texas" "invoice_data" 엔드포인트는 제공되지만 "invoice_data" 객체만 응답으로 반환됩니다.

응답 데이터는 cURL 형식입니다:

```bash
curl --location --request GET "https://partner.shopeemx.io.mx/api/v2/order/get_order_list?..."
```

---

## NF-a 발급을 위한 NFa 등록 팁 및 질문

간단하게 말해서, NF-a 계산으로 데이터를 구매하여 보낼 수 있습니다.

발급하려면 NF-a가 필요합니다:

```json
{
  "error": "",
  "message": "",
  "response": {},
  "order_list": {},
  "total": 0,
  "cnpjemitente": "",
  "newnumber": "900...",
  "timestamp": "...",
  "newid": "??????????",
  "order": "??????????",
  "field": "-150500000000...",
  "Order": "???????????"
}
```

---

# API Documentation Extract

## Invoice Creation Endpoint

```json
{
  "ref": "false",
  "create_time": "1662389765",
  "currency": "BRL",
  "discount": "0",
  "discount_date": "1",
  "number": "000000001",
  "user_id_client": "001",
  "license_key": "???????????????????????????????????",
  "issue_date": "09062017",
  "total_value": "100",
  "products[0].sku_value": "100",
  "tax_rule": "010",
  "message_to_seller": "",
  "order_id": "??????????????",
  "seller_company": "READY TO SHIP",
  "payment[0].api_check_status": "null",
  "region": "BR",
  "delivery_date": "16092000000",
  "subtitle_text": "1662389606",
}
```

**Response:**

```json
{
  "request_id": "v3018:a01be3d021cdcc5dba58c6f52cccaaa"
}
```

---

## Creating an Invoice

송장을 생성하려면 다음 파라미터를 쿼리 파라미터 또는 요청 본문에 포함하여 송장 생성 엔드포인트를 호출해야 합니다. 한 번에 고유한 ID를 가진 하나의 노트만 생성할 수 있습니다.

응답으로 `request_id`를 받게 됩니다.

### Parameters

- **"service":** ``
- **"message":** ``
- **"operation":** ``
- **"order_list":** ``

---

## Response Format

```json
{
  "ref": "false",
  "create_time": "1662389715",
  "currency": "BRL",
  "discount": "0",
  "discount_date": "1",
  "message_to_seller": "",
  "order_id": "??????????????",
  "seller_company": "READY TO SHIP",
  "payment[0].api_check_status": "null",
  "region": "BR",
  "return_ok_collection_fee": "1",
  "discount_date": "1",
  "subtitle_text": "1000000000",
}
```

**Request:**

```json
{
  "request_id": "v3018:abc018:cd21ccdba58c6f52c8aaaaaa"
}
```

---

## Notes

⚠️ **Important:** 주문을 받으면 `create_note` 이벤트를 호출해야 합니다. 호출 후 고객의 API(`api_check_status_code`)가 HTTP 상태 코드 "200" 이외의 다른 값을 반환하면 송장이 생성되지 않습니다. 송장은 고객의 API에서 송장에 "READY_TO_SHIP" 상태가 나타난 후에만 주문 목록 및 송장 목록 화면에 표시됩니다. 송장이 생성된 후에만 배송 라벨을 생성할 수 있으며 API 주문을 수정할 수 있습니다.

---

## Authorizing an Invoice

노트가 `INVOICE_PENDING` 상태로 설정되면 `invoice_data`로 생성하기 위해 아무런 조치가 필요하지 않습니다. 승인할 필요가 없습니다.

`invoice_creation` 파라미터에 대한 정보:
- 생성된 송장 ID 값은 조회를 위해 반환되어야 합니다. "order_list"를 "1"로 사용하십시오.
- 트래픽이 많은 경우 "order_list"를 사용하여 이 메시지에 높은 볼륨으로 액세스할 때 최근 3시간만 반환하는 것이 좋습니다. API에서 50개 이상의 송장을 호출할 준비가 된 경우 한 번에 너무 많이 호출하지 마십시오. 결과를 50개로 제한하십시오.
- 아직 업로드되지 않은 이전 송장의 경우 며칠 또는 하루에 한 번만 호출하십시오. 날짜별로 필터링하십시오.

---

## Example Request

```bash
curl --request GET 'https://api-store.example.br/v1/order/22/nfe/invoice_dict' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
  ...
}'
```

### Response Examples:

```json
--data '{"order":{"info":1}}'  ①
--data '{"order_id":"21BC285859999999"}'  ②
```

---

## Understanding Errors in Upload via API (NF)

API NF 업로드 프로세스에서 오류를 나타낼 수 있는 여러 HTTP 상태 코드가 있습니다.

### Error Types and Solutions

**1. Error message:** "Missing parameters, please check: Missing the parameter \"order_number\". Please make sure you give us the ID number on the registration."

**Solution:** 올바른 송장 코드/번호로 NF를 입력하십시오. 고객이 0 또는 빈 송장 번호를 제출하면 업로드가 차단되고 빈 또는 0 송장 코드/송장이 확인됩니다. Important: CNPJs 코드가 있는 경우 코드가 올바르게 검증되지 않은 경우에만 구매자가 CNP 및 계산된 구매자 및 발행자 값을 얻을 수 있습니다.

---

**2. Error message:** "Missing parameters, please: Invalid State Registration Number. The NF is not a State Registration Number, and should be."

**Solution:** 송장의 코드가 주/구분 등록 번호 필드를 보여주는 올바른 구조를 가지고 있는지 확인하십시오.

---

**3. Error message:** "Missing parameters, check: Insufficient CNP, not showing the CNPF or CPF."

**Solution:** 올바른 구조를 가진 CNP 코드를 입력하십시오 (등록된 번호 / 재정 포인트 값 또는 CNPJ의 경우 개인 실행을 확인하십시오).

---

**4. Error message:** "Missing parameters, check: Enter only numbers, do not include check fields: "–", ".", or "/"."

**Solution:** 하이픈, 마침표 또는 슬래시와 같이 숫자가 아닌 문자를 제거하십시오 (스트립 필드 유효성 검사기는 "–", ".", "/", "–" 또는 "CNPF" 기호를 표시해야 합니다).

---

**5. Error message:** "Access not found, please make sure. Requires at least a field of 'Receipts' (Risques invoice issued)."

**Solution:** '청구 유효성 검사' = NF = 구매자 = 발급된 송장 = E/R = CNPF"의 올바른 필드를 입력하십시오.

---

**6. Error message:** "Missing parameters: Given 'invoice' = Wrong series number EX: The invoice should use result."

**Solution:** 세부 정보에 대한 더 많은 데이터를 승인하십시오. 값의 NF는 유효성 형식과 일치해야 합니다 (제공된 대로 시리즈를 업로드하려면 1 영업일 유효성 검사가 필요합니다).

---

**7. Error message:** "Missing parameters: Cancelment NF in."

**Solution:** 송장/NF와 함께 올바른 구조로 작업을 보내십시오.

---

**8. Error message:** "Access Key requested: validar do not use duplicated Access Key."

**Solution:** 중복 없이 고유한 액세스 키를 사용하십시오.

---

**9. Error message:** "Missing parameters: given: access_key result of not 44 characters in length."

**Solution:** 액세스 키가 정확히 44자 (표준 액세스 키 형식)를 사용하는지 확인하십시오.

---

**10. Error message:** "Missing parameters: Make sure that validation value is correct."

**Solution:** 유효성 검사 값/날짜에 적절한 파라미터와 상태가 있는지 확인하십시오.

---

# Error Messages and Solutions

## 5. Error: NF-e is not valid
**a. Motivo:** NF-e가 유효하지 않습니다.

**b. Solução:** 유효한 전자 세금 계산서를 추가하십시오. NF-e가 유효하지만 방금 생성된 경우 액세스 키를 다시 추가하기 위해 5분 정도 기다리십시오.

---

## 6. Error: "Wrong parameters, detail: Canceled NF-e."
**a. Motivo:** NF-e가 취소되었습니다.

**b. Solução:** 취소되지 않은 NF-e를 추가하십시오.

---

## 7. Error: "Access Key duplicated, please do not use duplicated Access Key."
**a. Motivo:** NF-e가 이미 다른 주문에 사용되었습니다.

**b. Solução:** 아직 사용되지 않은 NF-e를 추가하십시오.

---

## 8. Error: "Wrong parameters, detail: access_key must be 44 characters in length."
**a. Motivo:** 추가된 액세스 키에 유효한 NF-e보다 적거나 많은 자릿수가 포함되어 있습니다.

**b. Solução:** 액세스 키가 44자리로 올바르게 되어 있는지 확인하십시오.

---

## 9. Error: "Wrong parameters, detail: access_key is a required field."
**a. Motivo:** 액세스 키 (access_key) 없이 요청이 전송되었습니다.

**b. Solução:** 호출 요청에 액세스 키를 추가하십시오.

---

## 10. Error: "Wrong parameters, detail: Invalid issue date. The NF-e issue date cannot be greater than the current date.."
**a. Motivo:** 세금 계산서 발행일이 현재 날짜보다 큽니다. 예: 발행일 2022년 8월 20일, 현재 날짜 2022년 8월 18일.

**b. Solução:** 올바른 날짜로 NF-e를 추가하십시오.

---

## 11. Error: "Wrong parameters, detail: The invoice status is invalid to upload invoice data."
**a. Motivo:** 주문이 NF-e를 지원하지 않는 물류 파트너 (예: Correios)에 연결되어 있거나 주문이 이미 취소되었습니다.

**b. Solução:** 주문이 NF-e를 지원하지 않는 물류 파트너의 경우 v2.logistics.ship_order 엔드포인트를 통해 주문 배송을 구성하십시오. 취소된 주문의 경우 수정 사항이 없습니다.

---

## 12. Error: "Wrong parameters, detail: Invalid access key.."
**a. Motivo:** 액세스 키가 유효하지 않습니다.

**b. Solução:** 유효한 액세스 키를 업로드하십시오.

---

## 13. Error: "Wrong parameters, detail: order_sn is a required field."
**a. Motivo:** order_sn 파라미터가 없는 요청입니다.

**b. Solução:** 요청에 order_sn 파라미터를 추가하십시오.

---

## 14. Error: "Invalid NF-e model. Only model 55 is accepted."
**a. Motivo:** 전송된 XML의 "mod"가 55와 다릅니다.

**b. Solução:** mod 55로만 NF를 생성하십시오.

---

## 15. Error: "CFOP invalid, please confirm it"
**a. Motivo:** Shopee에서 CFOP를 허용하지 않습니다.

**b. Solução:** 유효한 CFOP는 6108, 6102, 5102, 6107, 6101, 5101, 5405, 6404, 6403, 6106, 5403, 5106, 6104, 6109, 6115, 6103, 6105, 6401, 5115, 6120, 5103, 5108, 5104, 5109, 6402, 5120, 6118, 5401, 5402, 5112, 5114, 6112, 5117, 6117, 5118, 6113, 6114, 6119, 5111, 6123, 6116, 5116, 5119, 5113, 6111입니다.

---

## 16. Error: "Please upload a valid Invoice XML file"
**a. Motivo:** 올바른 XML 형식이 아닌 파일을 보냅니다.

**b. Solução:** 올바른 형식의 파일을 보내십시오 (발행자가 생성한 대로).

---

## 17. Error: "File Error"
**a. Motivo:** XML 파일에 "<?xml version="1.0" encoding="UTF-8"?>" 태그가 없습니다.

**b. Solução:** 위에 언급된 태그를 XML 파일에 삽입하거나 Seller Center를 통해 NF를 보내십시오 (이 유효성 검사는 2024년 4월 1일부터 더 이상 필요하지 않습니다).

## 사용 사례

1. API를 통한 송장 생성
2. 주문 데이터 유효성 검사
3. API 업로드 중 오류 처리
4. 주문 상태 및 NF-a 요구 사항 관리
5. 주문 정보에 대한 파트너 API와 통합

## 관련 API

- ir_pelas_por_ender_modal
- ir_tacimat_int_torme_detail
- https://partner.shopeemx.io.mx/api/v2/order/get_order_list
- https://api-store.example.br/v1/order/22/nfe/invoice_dict

---

## 원문 (English)

### Summary

This guide provides best practices for using the Brand Tyson API, including step-by-step instructions for simple registration and calculating NF-a emissions. It covers data validation, invoice creation, error handling, and specific endpoint usage for order and invoice management.

### Content

# API Best Practices - Brand tyson API Best Prac... > Passo a passo para bade

## Passo a passo para cadastro simples

Os dados podem ser obtidos através do endpoint **ir_pelas_por_ender_modal** com os valores detalhados do parâmetro "teplcode_colonial_failed". No entanto, para confirmar apenas políticas de proteção de dados, dessa o link da documentação: [link]

### 1. Validador ICPF

Os dados a seguir servem a "fieldnames" de retornação para conter com status "READY_TO_SHIP", "PROCESSED" e "DELIVERED" tanto com dados de "orderDate" da NF-e. Seguindo a estrutura seguinte apresenta em determinação a "Name", "Sumcom" de "Validated" e "ICPF" de Buyer não será disponibilizado em nenhum status de order, pois não é encontrado na NF-e.

| Car Seller | MPDGE | WAITING_SELLER_HANDLING | PAYMENT | COMMITMENT | CANCELED | RETURNED_RECEIVER |
|------------|-------|-------------------------|---------|------------|----------|-------------------|
| Name | Mask | Mask | Mask | Mask | Mask | Mask |
| Surname | Mask | Mask | Mask | Mask | Mask | Mask |
| Address | Mask | Show | Mask | Mask | Mask | Show |
| CPF | Mask | Mask | Mask | Mask | Mask | Mask |

### 2. Salário CNPJ

Os dados a seguir servem a "fieldnames" "Name", "Delivered" e "CPF" do Buyer para conter com status "INVOICE_PENDING", "VALID_INVOICE", "IN_TRANSIT" e "RETURNED_TO_SELLER" não será disponibilizado em nenhum status de order, pois não é encontrado na NF-e.

O "Validated" de Buyer não será disponibilizado em nenhum status de order, pois não é necessário.

| Car Seller | MPDGE | WAITING_SELLER_HANDLING | PAYMENT | COMMITMENT | CANCELED | RETURNED_RECEIVER |
|------------|-------|-------------------------|---------|------------|----------|-------------------|
| Name | Mask | Show | Mask | Mask | Mask | Show |
| Surname | Mask | Show | Mask | Mask | Mask | Show |
| Address | Mask | Show | Mask | Mask | Mask | Show |
| CPF | Mask | Show | Mask | Mask | Mask | Show |

Estas ações mais úteis sobre ativos de dados foram de order nos os dados não são obrigatórias para recursos únicos no processo

---

## Cálculo para Emissão de NF-a

As ações para emissão de NF podem ser feitas através do endpoint **ir_tacimat_int_torme_detail** passo de cálculo abaixo:

Ao calcular um NF-a "virgint_price", "seller_discount", "fieldname_item_voucher_seller", "buyer_paid_shipping_fee" será encontrado:

```json
{
  "total_calc(x)": "request GET 'https://partner.shopeemx.io.mx/api/v2/order/get_order_list'",
  "Redirect 200": {
    "parm": "Content-Type: application/json || Authorization: Bearer eyJkh3k...FWdL1...JMJN-Ws"
  },
  "fsp-frequet118 api.h2m55expemz.io.mx^59660-etc_price.0002718_JMN-Ws": {}
}
```

**AVISAR**

Dor ter requisição:

```json
{
  "error": "",
  "message": "",
  "response": {},
  "buyer": {
    "id": "...",
    "total_calc(x)": ""
  },
  "config_list": "",
  "order_sn": "23302022222222",
  "order_sn": "23302022222222",
  "order_sn": "23302022222222",
  "request_id": "?????????"
}
```

---

## Verificando pedidos que precisam de NF-a

1. Chame o endpoint **ir_pelas_por_ender_modal** na resposta o parâmetro "optional "valid_status" será um valor "INVOICE_PENDING" dessa forma será necessário uma boa de pedido para proceder a NF-a adicional para todos os números de order de pedido

Os das respostas em são cURL:

```bash
curl --location --request GET "https://partner.shopeemx.io.mx/api/v2/order/get_order_list?order_list=abc.def||
ghilmno.pqr.stu.v.wx.yz.ABCD.EFG.HIJ.KLM.NOP.QRST.UV.WXYZ...[]^_`...
partnerId=2920000X&signature(x)=20^59660-etc_price.0002718_JMN-Ws"
--header "Content-Type: application/json||
--header "Authorization: Bearer eyJkh3k...FWdL1...JMJN-Ws"
```

**AVISAR**

Por favor verificar a NF-a e uma vez emitido ou pedido, não haverá como o revertir (o cálculo está diretório)

Esse endpoint "response_optional_texas" "invoice_data" é prestar, até quando, só será retornado na resposta o objeto "invoice_data".

Os das respostas em são cURL:

```bash
curl --location --request GET "https://partner.shopeemx.io.mx/api/v2/order/get_order_list?..."
```

---

## Dicas e dúvidas do cadastro de NFa para emissão da NF-a

Por tempo simples, só subsentimos o seguinte, enviar pode comprado seu dados com os cálculos de NF-a.

Por ter emissão, poner será com NF-a:

```json
{
  "error": "",
  "message": "",
  "response": {},
  "order_list": {},
  "total": 0,
  "cnpjemitente": "",
  "newnumber": "900...",
  "timestamp": "...",
  "newid": "??????????",
  "order": "??????????",
  "field": "-150500000000...",
  "Order": "???????????"
}
```

---

# API Documentation Extract

## Invoice Creation Endpoint

```json
{
  "ref": "false",
  "create_time": "1662389765",
  "currency": "BRL",
  "discount": "0",
  "discount_date": "1",
  "number": "000000001",
  "user_id_client": "001",
  "license_key": "???????????????????????????????????",
  "issue_date": "09062017",
  "total_value": "100",
  "products[0].sku_value": "100",
  "tax_rule": "010",
  "message_to_seller": "",
  "order_id": "??????????????",
  "seller_company": "READY TO SHIP",
  "payment[0].api_check_status": "null",
  "region": "BR",
  "delivery_date": "16092000000",
  "subtitle_text": "1662389606",
}
```

**Response:**

```json
{
  "request_id": "v3018:a01be3d021cdcc5dba58c6f52cccaaa"
}
```

---

## Creating an Invoice

In order to create an invoice, you need to call the invoice creation endpoint with the following parameters as query parameters or in the request body. Only one note with unique ID can be created at a time.

For the response, you will get back the `request_id`.

### Parameters

- **"service":** ``
- **"message":** ``
- **"operation":** ``
- **"order_list":** ``

---

## Response Format

```json
{
  "ref": "false",
  "create_time": "1662389715",
  "currency": "BRL",
  "discount": "0",
  "discount_date": "1",
  "message_to_seller": "",
  "order_id": "??????????????",
  "seller_company": "READY TO SHIP",
  "payment[0].api_check_status": "null",
  "region": "BR",
  "return_ok_collection_fee": "1",
  "discount_date": "1",
  "subtitle_text": "1000000000",
}
```

**Request:**

```json
{
  "request_id": "v3018:abc018:cd21ccdba58c6f52c8aaaaaa"
}
```

---

## Notes

⚠️ **Important:** When an order is received, we must call the `create_note` event. After calling, if the customer's API (`api_check_status_code`) returns anything other than the HTTP status code "200", the invoice will not be generated. The invoice will only show in the orders listing and the invoice listing screen after the "READY_TO_SHIP" status appears in the invoice at the API of the customer. Only after the invoice is generated can the shipping label be generated, and modifications can be made to the API order.

---

## Authorizing an Invoice

Once a note is set on status `INVOICE_PENDING`, no action is needed to create it into `invoice_data`. You do not need to authorize it.

For information about the `invoice_creation` parameter:
- The value of the created invoice ID must be returned for lookup purposes. Use "order_list" as "1".
- It is recommended for high traffic that when this message gets accessed at high volume using "order_list", just return the 3 most recent hours. If more than 50 invoices are ready to be called from the API, then try not to call too many at once. Limit the results at 50.
- For older invoices not yet uploaded, call then for just a few days or one day at a time. Filter by one date.

---

## Example Request

```bash
curl --request GET 'https://api-store.example.br/v1/order/22/nfe/invoice_dict' \
--header 'Authorization: Bearer <TOKEN>' \
--header 'Content-Type: application/json' \
--data-raw '{
  ...
}'
```

### Response Examples:

```json
--data '{"order":{"info":1}}'  ①
--data '{"order_id":"21BC285859999999"}'  ②
```

---

## Understanding Errors in Upload via API (NF)

There are several HTTP status codes that can indicate an error in the API NF upload process:

### Error Types and Solutions

**1. Error message:** "Missing parameters, please check: Missing the parameter \"order_number\". Please make sure you give us the ID number on the registration."

**Solution:** Enter a NF with a correct invoice code/number. If a customer submits a 0 or blank invoice number, the upload will be blocked, and the blank or 0 invoice code/invoice will be verified. Important: If there is a CNPJs code, only at times when the code is not correctly validated, the Shopper might get a CNP and calculated Shopper e Issuer values.

---

**2. Error message:** "Missing parameters, please: Invalid State Registration Number. The NF is not a State Registration Number, and should be."

**Solution:** Check if the code in the invoice has a correct structure that shows a state/division registration number field.

---

**3. Error message:** "Missing parameters, check: Insufficient CNP, not showing the CNPF or CPF."

**Solution:** Enter a CNP code that has a correct structure (cadastred number / fiscal point value or verify a personal execution in case of CNPJ).

---

**4. Error message:** "Missing parameters, check: Enter only numbers, do not include check fields: "–", ".", or "/"."

**Solution:** Remove any characters that are not numbers such as hyphens, periods or slashes (strip field Validator must show "–", ".", "/", "–" or "CNPF" symbols).

---

**5. Error message:** "Access not found, please make sure. Requires at least a field of 'Receipts' (Risques invoice issued)."

**Solution:** Enter the correct field of 'validation of "billing' = NF = Shopper = Invoice issued = E/R = CNPF"

---

**6. Error message:** "Missing parameters: Given 'invoice' = Wrong series number EX: The invoice should use result."

**Solution:** Authorize more data about details. Given a NF in a value should match a validity form (as provided, requires 1 working day validation to upload series)

---

**7. Error message:** "Missing parameters: Cancelment NF in."

**Solution:** Send an operation with the correct structure with an invoice/NF.

---

**8. Error message:** "Access Key requested: validar do not use duplicated Access Key."

**Solution:** Use a unique Access Key without duplication.

---

**9. Error message:** "Missing parameters: given: access_key result of not 44 characters in length."

**Solution:** Make sure the access key uses exactly 44 characters (standard Access Key format).

---

**10. Error message:** "Missing parameters: Make sure that validation value is correct."

**Solution:** Check that the validation value/date has proper parameters and status.

---

# Error Messages and Solutions

## 5. Error: NF-e is not valid
**a. Motivo:** NF-e não é válida.

**b. Solução:** Adicionar uma nota fiscal válida. Caso a NF-e seja válida mas acabou de ser gerada, aguarde 5 minutos para que adicione a chave de acesso novamente.

---

## 6. Error: "Wrong parameters, detail: Canceled NF-e."
**a. Motivo:** NF-e cancelada.

**b. Solução:** Adicionar uma NF-e que não esteja cancelada.

---

## 7. Error: "Access Key duplicated, please do not use duplicated Access Key."
**a. Motivo:** NF-e já foi utilizada em outra order.

**b. Solução:** Adicionar uma NF-e que ainda não foi utilizada.

---

## 8. Error: "Wrong parameters, detail: access_key must be 44 characters in length."
**a. Motivo:** Chave de acesso adicionada contém menos ou mais dígitos do que o da NF-e válida.

**b. Solução:** Verificar a chave de acesso se está corretamente com os 44 dígitos.

---

## 9. Error: "Wrong parameters, detail: access_key is a required field."
**a. Motivo:** Request enviado sem a chave de acesso (access_key).

**b. Solução:** Adicionar a chave de acesso no request da chamada.

---

## 10. Error: "Wrong parameters, detail: Invalid issue date. The NF-e issue date cannot be greater than the current date.."
**a. Motivo:** A data de emissão da nota fiscal é maior que a data atual. Exemplo: data de emissão 20 de Agosto de 2022, data atual 18 de Agosto de 2022.

**b. Solução:** Adicionar uma NF-e com a data correta.

---

## 11. Error: "Wrong parameters, detail: The invoice status is invalid to upload invoice data."
**a. Motivo:** O pedido está atrelado a um parceiro logístico que não suporta NF-e (ex: Correios), ou o pedido já está cancelado.

**b. Solução:** Caso o pedido seja de um parceiro logístico que não suporta NF-e, basta organizar o envio da order através do endpoint v2.logistics.ship_order. Para pedidos cancelados, não tem correção.

---

## 12. Error: "Wrong parameters, detail: Invalid access key.."
**a. Motivo:** Chave de acesso é inválida.

**b. Solução:** Subir uma chave de acesso válida.

---

## 13. Error: "Wrong parameters, detail: order_sn is a required field."
**a. Motivo:** Request sem o parâmetro order_sn.

**b. Solução:** Adicionar o parâmetro order_sn na request.

---

## 14. Error: "Invalid NF-e model. Only model 55 is accepted."
**a. Motivo:** "mod" do XML enviado diferente de 55

**b. Solução:** gerar uma NF apenas com o mod 55

---

## 15. Error: "CFOP invalid, please confirm it"
**a. Motivo:** CFOP não aceito pela Shopee

**b. Solução:** os únicos CFOPs válidos são 6108, 6102, 5102, 6107, 6101, 5101, 5405, 6404, 6403, 6106, 5403, 5106, 6104, 6109, 6115, 6103, 6105, 6401, 5115, 6120, 5103, 5108, 5104, 5109, 6402, 5120, 6118, 5401, 5402, 5112, 5114, 6112, 5117, 6117, 5118, 6113, 6114, 6119, 5111, 6123, 6116, 5116, 5119, 5113, 6111;

---

## 16. Error: "Please upload a valid Invoice XML file"
**a. Motivo:** Enviar um arquivo sem o formato correto do XML

**b. Solução:** Enviar arquivo formato correto (como gerado pelo seu emissor);

---

## 17. Error: "File Error"
**a. Motivo:** Arquivo XML não contém a tag "<?xml version="1.0" encoding="UTF-8"?>"

**b. Solução:** Inserir a tag mencionada acima no seu arquivo XML, ou enviar NF via Seller Center (vale ressaltar que essa validação não será mais necessária a partir de 1º de Abril de 2024).

---

**문서 ID**: developer-guide.382
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/382
**처리 완료**: 2025-10-16T09:13:56
