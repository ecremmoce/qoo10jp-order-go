# 새로운 주문 식별

**카테고리**: 통합
**난이도**: 중간
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T09:15:50

## 개요

이 가이드는 TikTok Open Platform API를 사용하여 새로운 주문을 식별하고 주문 상태를 관리하는 방법에 대한 정보를 제공합니다. 다양한 주문 상태, 취소 프로세스 및 주문 관리와 관련된 자주 묻는 질문을 다룹니다.

## 주요 키워드

- 주문 관리
- 주문 상태
- OpenAPI
- Seller Center
- API
- 취소
- 추적 번호
- Push Marketplace

## 본문

# 새 주문 식별

현재 API를 통해 새 주문을 식별하는 방법이 있습니다.

## 단계

Push Marketplace를 사용하려면 OAuth를 위한 경로도 생성해야 합니다. 두 URL은 API에서 제공하는 콜백 URL의 주문을 사용하며, 이 Push는 API 요구 사항에 따라 응답을 구성해야 합니다.

1.  `ORDER` 상태를 호출해야 합니다.
2.  새로운 `POST1` 메서드를 포함합니다.

현재 Push Marketplace를 얻을 수 있습니다.

API `"v2_post_set_push_ordem"`은 Push를 알리고, Push는 선택되고 호출자의 상태는 "'ready_status'"입니다.

요청에 따라 입력합니다:

```json
"seller_url" : [
  "v1/po_orders"
]

"apiback_url" : "http://1-38.10"
```

Open Platform에서 Open Platform > Consoles > App Management > Push Marketplace > Set Posts R app > app details에서 지침을 통해 정보를 입력합니다.

## GET 요청 가져오기

`get_order_list`를 얻으려면 특정 기간의 모든 `orders_id`를 추적하는 두 개의 응답이 필요합니다.

선택적 매개변수 `"response_optional_fields=order_status"`를 선택하고 주문 데이터를 관리하고 기간 목록을 선택하여 요청합니다.

### API 호출 예시

```http
curl --OauthID --request GET
'https://open.tiktokapis.com/v1/fulfillment/api_order_list?
app_key=12&access_token=USER_ACCESS_TOKEN&shop_id=12&create_time_from=1609862012&create_time_to=1609862012&update_time_from=1609862012&update_time_to=1609862012&
response_optional_fields=order_status&cursor=ROW%3D1%3BINDEX%3D0&page_size=10'
```

### 응답

```json
{
  "error": "",
  "message": "",
  "timestamp": 0,
  "api": true,
  "next": "",
  "cargo_line": [],
  "fee_id": "INITSTORED",
  "order_status": "READY_TO_SHIP",
  "seller_id": "INITSTORED",
  "order_id": "TEMPLATE_SELLER",
  "related_id": "?????????????"
}
```

---

# OpenAPI의 주문 상태

셀러가 액세스할 수 있는 6가지 상태 주문은 다음과 같습니다.

## 1. UNPAID
생성할 수 있도록 승인된 주문 및 결제 유효성을 검사하기 전:
일반적으로 신용 카드 주문은 결제보다 유효성 검사가 더 크지 않으므로 이러한 상태를 표시하지 않습니다.

## 2. ON_HOLD
요청된 주문은 `v1_get_order_list` API 또는 정보가 제공된 orders 네임스페이스에만 표시됩니다.

매개변수:
- AWAITING_SHIPMENT
- READY_TO_SHIP
- 결제 확인을 위해 보류 중이며, 주문 API에서 반환된 데이터 주문 양식 내에서 진행됩니다 (예: 결제가 유효성 검사 중).

## 3. AWAITING_SHIPMENT
다른 새로운 상태는 조직에서 사용할 수 있도록 허용합니다 (`v2_ship_getlist`). 일반화된 주문을 사용하고 결제 확인 상태를 알리고 사용자가 게시할 수 있도록 허용합니다 (예: 결제 문제가 있는 주문).

## 4. PROCESSED
`v2_ship_order`에서 이미 배송된 주문 및 `v2_ship_order`의 데이터는 추적 번호로 관리됩니다.

## 5. SHIPPED
주문은 이미 전체 물류 환자에게 배송되었습니다.

## 6. TO_CONFIRM_RECEIVE
주문이 접수될 때

## 7. COMPLETED
이 상태는 항목이 최종적으로 수신된 경우입니다. 구매자 또는 반품, 개혁 또는 새로운 상태의 셀러 목록입니다.

## 8. IN_CANCEL
구매자 또는 주문 반품 요청을 나타냅니다.

주문이 취소되면 구매자가 주문을 표시할 때, 셀러가 주문을 거부할 때, 시스템이 OpenAPI 및 Seller Center 주문을 비활성화하는 경우 [documentation link] 테이블에 의해 보호됩니다.

---

## 상태 매핑

| OpenAPI | SellerCenter |
|---------|--------------|
| UNPAID | - |
| READY_TO_SHIP | A enviar (a iniciar) |

---

# 9 - CANCELLED

주문이 취소되면 구매자가 주문을 취소하거나, 셀러가 주문을 준비하지 않거나, 배송 중 주문이 손실된 경우 발생할 수 있습니다.

OpenAPI와 Seller Center 간의 상태 비교는 다음과 같습니다 ([Documentação](link)에서 찾을 수도 있습니다).

| OpenAPI | SellerCenter |
|---------|--------------|
| UNPAID | Não Pago |
| READY_TO_SHIP | A enviar (A enviar) |
| IN_CANCEL | Cancelado (A responder) |
| CANCELLED | Cancelado (Cancelado) |
| SHIPPED | A enviar (Processado) |
| TO_RETURN | Enviado |
| TO_CONFIRM_RECEIVE | Devolução Reembolso |
| COMPLETED | Shipping (Shipped) |

## 주문 취소

셀러 측에서 주문을 취소하려면 API [v2 cancel_order](link)를 사용하십시오. 몇 가지 규칙이 있습니다.

1 - 배송 준비 및 `tracking_number` 생성 후에는 주문을 취소할 수 없습니다.

2 - 취소에 사용할 수 있는 4가지 이유가 있습니다.
   1. OUT_OF_STOCK: 재고가 소진된 경우 (이유가 항상 이 경우 `"item_list":["item_id", "model_id"]` 배열을 보내야 함);
   2. CUSTOMER_REQUEST: 구매자가 요청하고 더 이상 귀하의 승인 없이 취소할 수 없는 경우;
   3. UNDELIVERABLE_AREA: 요청된 주소가 배송 지역 내에 없는 경우;
   4. COD_NOT_SUPPORTED: 현금 배송은 브라질에서 사용할 수 없습니다.

주문 취소에 대한 요청 본문의 예는 다음과 같습니다.

```java
{
  "order_sn": "2012030005ACB",
  
  "cancel_reason": "OUT_OF_STOCK",
  
  "item_list": [{
    
    "item_id": 1680783,
    
    "model_id": 327890123
    
  }]
  
}
```

오류 메시지가 표시되면:

"`Can not cancel this order.`"는 주문이 이미 "`tracking_number`"가 생성되었거나 주문이 이미 취소되었기 때문에 취소할 수 없기 때문입니다. 이 두 가지 상황과 다른 상황이 발생하면 [Plataforma de Tickets](link)를 통해 문의하십시오.

구매자 측에서 주문 취소 요청을 하려면 API [v2 handle_buyer_cancellation](link)를 사용하여 취소 요청을 수락하거나 거부하십시오.

## 자주 묻는 질문 (FAQ)

1 - 브라질 통합에 사용되지 않는 몇 가지 API는 다음과 같습니다.

```
"v2.order.split_order"
"v2.order.unsplit_order"
"v2.order.get_buyer_invoice_info"
"v2.order.get_pending_buyer_invoice_order_list"
"v2.order.upload_invoice_doc"
"v2.order.download_invoice_doc"
```

2 - 지난 주에 재고가 0으로 설정되었는데 왜 주문을 생성할 수 있었습니까?

가능한 시나리오는 두 가지가 있습니다. 하나는 셀러가 캠페인을 위해 별도로 예약된 재고를 가지고 있다는 것이고, 다른 하나는 이전 주문이 취소되어 재고가 복원되었다는 것입니다 (취소된 모든 주문은 재고를 복원합니다).

## 사용 사례

1. 주문 목록 검색
2. 주문 상태 업데이트
3. 주문 취소
4. 판매자 시스템과의 주문 관리 통합
5. 구매자 취소 요청 처리

## 관련 API

- v2_post_set_push_ordem
- v1_get_order_list
- v2_ship_getlist
- v2_ship_order
- v2 cancel_order
- v2 handle_buyer_cancellation
- v2.order.split_order
- v2.order.unsplit_order
- v2.order.get_buyer_invoice_info
- v2.order.get_pending_buyer_invoice_order_list
- v2.order.upload_invoice_doc
- v2.order.download_invoice_doc

---

## 원문 (English)

### Summary

This guide provides information on how to identify new orders and manage order statuses using the TikTok Open Platform APIs. It covers various order statuses, cancellation processes, and frequently asked questions related to order management.

### Content

# Identifying a New Order

Currently there are ways to identify a new order via APIs

## Steps

Para usar o Push Marketplace devem também ser uma rota para fins OAuth ser criada, as duas Urls end usuaria o Pedido do a API informada callback URL, e essa Push também precisa configurar uma resposta de acordo must requisitos da API

1. Voce deve também chamar o status ORDER
2. Incluir um outro novo método POST1

Atualmente é possível obter o Push Marketplace

são API "v2_post_set_push_ordem" informando prata Push dentro um efetivado, o Push para elegido e criados um status dos callers é "o 'ready_status'"

Digitar de acordo as request:

```json
"seller_url" : [
  "v1/po_orders"
]

"apiback_url" : "http://1-38.10"
```

In Open Platform, Você logo va Open Platform > Consoles > App Management > Push Marketplace > Set Posts R app > app details da informar via instrução

## Obtaining a GET Request

Para obter a get_order_list, two response track todos os orders_id da pedque a periodo especifico

Selecionar o parâmetro opcional "response_optional_fields=order_status" e requisitar também gerenciado os dados dos orders, como os seleção period lista.

### API Call Example

```http
curl --OauthID --request GET
'https://open.tiktokapis.com/v1/fulfillment/api_order_list?
app_key=12&access_token=USER_ACCESS_TOKEN&shop_id=12&create_time_from=1609862012&create_time_to=1609862012&update_time_from=1609862012&update_time_to=1609862012&
response_optional_fields=order_status&cursor=ROW%3D1%3BINDEX%3D0&page_size=10'
```

### Response

```json
{
  "error": "",
  "message": "",
  "timestamp": 0,
  "api": true,
  "next": "",
  "cargo_line": [],
  "fee_id": "INITSTORED",
  "order_status": "READY_TO_SHIP",
  "seller_id": "INITSTORED",
  "order_id": "TEMPLATE_SELLER",
  "related_id": "?????????????"
}
```

---

# Order Status on OpenAPI

As of 6 status orders as que o sellers podem be accessed:

## 1. UNPAID
Order orders que autorizam de ser criadas e antes têm ao possível validar e pagamento:
Normalmente orders de Status não têm apresentar esses status, pois os orders do Cartão de Crédito não validation em maiores do que pagamento.

## 2. ON_HOLD
Pedido solicitar aparece apenas API do v1_get_order_list ou o namespace orders onde dos informado acerca

Parameters:
- AWAITING_SHIPMENT
- READY_TO_SHIP
- Hold for confirmado de pagamento, go within form orders order be dados return order API (e g., pagamento está sendo validado).

## 3. AWAITING_SHIPMENT
Outros novos status whilst permite uma Organização de uso (v2_ship_getlist) combinação usar o normalizado order e sinalizado e permitir status do confirmação de pagamento, pois dentro destes permite n Post o User (e.g., orders com problemas de pagamento).

## 4. PROCESSED
Orders que ja foram entregues no v2_ship_order e dados no v2_ship_order a gerida to tracking Number

## 5. SHIPPED
Orders ja foram all entregue do paciente logístico

## 6. TO_CONFIRM_RECEIVE
When orders de são ao recebimento

## 7. COMPLETED
Esta status é when an item reçu final, qual-que ou a Lista do buyer, em that classa de devolução, reforma ou Seller de uma novos status

## 8. IN_CANCEL
Indica quando buyer ou a solicitação de devolução do pedido

Quando a pedido é Cancelado, quais pode orders quando o buyer indicar o pedido, quando o Seller clica negativa o pedido, além sistema lugar de-ativando o pedido do order OpenAPI e Seller Center se enquadram por escudadas pela tabela [documentation link]

---

## Status Mapping

| OpenAPI | SellerCenter |
|---------|--------------|
| UNPAID | - |
| READY_TO_SHIP | A enviar (a iniciar) |

---

# 9 - CANCELLED

Quando a order é Cancelada, que pode ocorrer quando o buyer cancela o pedido, quando o Seller não organiza o pedido ou quando o pedido é perdido durante a entrega.

Segue comparação de status entre OpenAPI e Seller Center (que também pode ser encontrada na nossa [Documentação](link)):

| OpenAPI | SellerCenter |
|---------|--------------|
| UNPAID | Não Pago |
| READY_TO_SHIP | A enviar (A enviar) |
| IN_CANCEL | Cancelado (A responder) |
| CANCELLED | Cancelado (Cancelado) |
| SHIPPED | A enviar (Processado) |
| TO_RETURN | Enviado |
| TO_CONFIRM_RECEIVE | Devolução Reembolso |
| COMPLETED | Shipping (Shipped) |

## Cancelamento de pedidos

Para cancelamento de pedidos do lado do Seller, utilize a API [v2 cancel_order](link), algumas regras:

1 - Não é possível cancelar uma order após a organização de envio e geração do tracking_number;

2 - Existem 4 motivos disponíveis para cancelamento:
   1. OUT_OF_STOCK: Para quando o estoque fiver acabado (sempre que esse for o motivo é necessário o envio do array "item_list":["item_id", "model_id"] );
   2. CUSTOMER_REQUEST: Quando o Buyer solicita e não consegue mais cancelar sem sua autorização;
   3. UNDELIVERABLE_AREA: Quando o endereço solicitado não está dentro da sua área de entrega;
   4. COD_NOT_SUPPORTED: Cash in Delivery, não disponível para o Brasil.

Segue exemplo de request body para cancelamento de order:

```java
{
  "order_sn": "2012030005ACB",
  
  "cancel_reason": "OUT_OF_STOCK",
  
  "item_list": [{
    
    "item_id": 1680783,
    
    "model_id": 327890123
    
  }]
  
}
```

Caso receba a mensagem de erro:

"Can not cancel this order." é porque o pedido não pode ser cancelado por já ter sido gerado seu "tracking_number", ou porque a order já foi cancelada. Qualquer situação diferente dessas duas, por gentileza nos acionar via [Plataforma de Tickets](link).

Já para solicitação de cancelamento de pedido do lado do buyer, utilize a API [v2 handle_buyer_cancellation](link) para aceitar ou rejeitar a solicitação de cancelamento.

## Dúvidas frequentes (FAQ)

1 - Seguem algumas APIs que não são usadas para integrações no Brasil:

```
"v2.order.split_order"
"v2.order.unsplit_order"
"v2.order.get_buyer_invoice_info"
"v2.order.get_pending_buyer_invoice_order_list"
"v2.order.upload_invoice_doc"
"v2.order.download_invoice_doc"
```

2 - Por que foi possível criar uma order se o estoque foi zerado semana passada?

Existem dois cenários possíveis, um deles é que o Seller tem um estoque separado e reservado para campanhas, a outra possibilidade é que uma order antiga foi cancelada e restaurou o estoque (toda order cancelada restaura o estoque

---

**문서 ID**: developer-guide.383
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/383
**처리 완료**: 2025-10-16T09:15:50
