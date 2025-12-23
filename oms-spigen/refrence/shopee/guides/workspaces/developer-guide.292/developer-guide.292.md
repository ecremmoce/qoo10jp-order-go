# Passo a Passo Logística API

**카테고리**: 모범 사례
**난이도**: medium
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T09:09:56

## 개요

이 가이드는 Logistics API의 단계별 안내를 제공하며, Shopee의 물류 서비스와 통합하기 위한 호출 흐름 및 모범 사례를 자세히 설명합니다. 배송 문서 생성, 추적, 다양한 배송 채널 처리 등 다양한 시나리오를 다룹니다.

## 주요 키워드

- Logistics API
- 배송
- 추적
- 배송 문서
- 배송 채널
- 주문 관리
- Shopee API
- v2 API

## 본문

# Passo a Passo Logística API

## API 모범 사례 > Brand Coop API 모범 사례... > Passo a Passo Logística API

### 흐름도

```
v2.logistics.get_shipping_parameter
          ↓
      Drop Off ←→ Pick Up
          ↓
   v2.logistics.ship_order
          ↓
v2.logistics.get_tracking_number
          ↓
v2.logistics.get_airway_document_parameter
          ↓
v2.logistics.create_shipping_document
          ↓
v2.logistics.get_shipping_document_result
          ↓
v2.logistics.download_shipping_document
          ↓
         Fim
```

### 참고 및 정보 상자

**ℹ️ 참고 1:**
배송을 구성하기 위한 정보를 수집합니다. 판매자가 확인 없이 주문을 로컬에서 처리할 수 있도록 모드 주소를 수집합니다.

**ℹ️ 참고 2:**
이 API를 호출한 후 주문 상태가 자동으로 "READY_TO_SHIP"(Shopee에 게시)으로 업데이트됩니다. "v2.order.get_order_list"를 사용하여 "PROCESSED" 목록을 호출합니다.

**ℹ️ 참고 3:**
"PROCESSED" 상태에서만 더 일찍 호출할 수 있습니다. 실패한 시도(API에 오류가 발생)가 있는 경우 판매자의 주문에는 영향을 미치지 않습니다.

**ℹ️ 참고 4:**
일반 AWB 및 Thermal AWB 유형이 있을 수 있습니다. 귀사의 물류 회사에서 배송 문서를 받을 필요가 없는 경우 문서 획득 단계를 건너뛰고 배송 문서 다운로드 단계로 바로 이동할 수 있습니다.

**ℹ️ 참고 5:**
TrackingNo가 생성된 후에 라벨을 만들 수 있습니다.

**ℹ️ 참고 6:**
이 API가 모든 배송 문서의 상태가 표시되는 상태를 반환하면 모든 문서를 다운로드할 수 있습니다.

**ℹ️ 참고 7:**
6단계에서 결과가 제공되지 않으면 다운로드할 CDF를 얻기 위해 "v2.logistics.get_shipping_document_result"를 계속 호출합니다. TrackingNo,_init가 반환되지 않으면 이 값을 null로 설정할 수 있습니다.

---

## v2에서 배송 채널과 관련된 호출 흐름은 다음과 같습니다.

### 1. Shopee는 판매자를 위해 배송 채널을 활성화합니다(예: Correios, 모든 판매자에게 공개, Total Express 등).

### 2. 스토어에서 활성화된 물류 채널을 확인하려면

1. 엔드포인트 v2.logistics.get_channel_list를 호출합니다.

### 3. 판매자는 제공할 수 있도록 해제된 각 품목에 대한 물류 채널을 얻어야 합니다(Shopee 물류).

1. 제품의 배송 채널을 활성화/비활성화하려면 v2.product.edit, v2.product.update_price의 품목에서 logistic_info 내의 logistic_info.logistic_id를 준비해야 합니다.

### 4. 판매자는 구매자가 선택한 배송 채널로 주문을 받습니다.

### 5. 배송 준비(ship_order)를 하기 전에 주문에 대한 NF-e를 알려야 하는지 여부를 확인하려면 다음을 수행할 수 있습니다.

1. v2.order.get_order_detail 호출을 사용합니다. 각 호출에는 "order_status" "INVOICE_PENDING", "READY_TO_SHIP" 매개변수가 있습니다.
2. NF-e 정보를 제공하려면 v2.order.add_invoice_data 호출을 사용합니다.
3. 5단계를 확인한 후 v2.logistics.get_shipping_parameter, v2.logistics.ship_order를 "invoice_data" 매개변수와 함께 사용합니다.

### 6. 노드 v2.logistics.ship_order에 대한 배송 및 드롭을 수행하려면

1. 구매자가 선택한 배송 채널에 따라 판매자는 주문을 게시하거나 수거를 기다립니다.
2. 수거(pickup) 또는 게시(dropoff)인지 확인하려면 v2.logistics.get_shipping_parameter를 사용합니다.
3. Dropoff: 판매자는 물류에 보내야 합니다. v2.logistics.ship_order를 사용한 후
4. dropoff인 경우(Correios를 통해 dropoff로 가는 경우) info_needed가 비어 있으면("dropoff" []), ship_order 요청은 다음과 같아야 합니다.

```json
{
  "order_sn": "12CR12DCXXXXX",
  
  "dropoff": []
}
```

### 8. pickup인 경우(pickup이 아닌 다른 모든 루프, Coleta Correios 포함) info_needed에 "pickup"("address_id","pickup_time_id")이 있는 경우 ship_order 요청은 다음과 같아야 합니다.

```json
{
  "order_sn": "12CR12DCXXXXX",
  
  "pickup": []
}
```

**매개변수가 있는 예:**

```json
"address_id":10,box_address_list.address_id | in_get_address_list.n.ne.id do_self_ship(log_parameter)

"pickup_time_id":10,box_address_list.time_slot_list.pickup_time_id do_self_ship(log_parameter)
```

---

## 9. 그런 다음 추적 코드를 얻으려면 엔드포인트 v2.logistics.get_tracking_number를 호출하면 됩니다.

### 10. 라벨에 사용할 수 있는 매개변수를 얻으려면 v2.logistics.get_shipping_document_parameter를 사용합니다.

1. 라벨에 사용할 수 있는 매개변수를 얻으려면 v2.logistics.get_shipping_document_parameter(유형 "NORMAL_AIR_WAYBILL")는 라벨이 각각 pdf 또는 .zip 형식으로 생성될 수 있음을 반환합니다(.txt 형식의 라벨이 포함되어 있으며 .ZPL 형식은 컨텍스트 언어임).

### 11. Shopee에 하나 이상의 주문에 대한 라벨 생성을 요청하려면

1. v2.logistics.create_shipping_document를 호출합니다.
2. "tracking_number" 매개변수를 보내야 합니다.

### 12. 라벨이 이미 생성되었는지 아니면 아직 처리 중인지 확인하려면

1. v2.logistics.get_shipping_document_resultHolder는 응답 "status" "READY"로 라벨을 얻을 수 있습니다.

### 13. Shopee 프로세스에 대한 라벨(AIR WAYBILL)을 얻으려면

1. v2.logistics.download_shipping_document

---

## 주문 호출 흐름에 대한 자세한 내용은 추적 배송 및 추적 번호 가져오기 및 항공 운송장 인쇄 섹션의 설명서를 참조하십시오.

[https://open.shopee.com/developer/public/v2](https://open.shopee.com/developer/public/v2)

---

## FAQ

### 모든 배송은 Shopee와 직접 연결된 배송 채널(운송업체)을 통해 이루어집니다.

A: 아니요, Shopee는 무료이며 판매자가 지불합니다.

---

### 모든 배송 라벨은 Shopee에서 생성하거나 판매자 또는 통합업체/ERP에서 라벨을 생성해서는 안 됩니까?

A: 라벨이 있는 경우 라벨은 v2.logistics.download_shipping_document를 통해 Shopee에서 얻어야 합니다.

---

### 각 판매자는 단일 배송 채널을 가지고 있습니다. 즉, 다른 판매자는 다른 배송 채널에 액세스할 수 있습니다.

A: 예, 각 판매자는 Correios에 대한 액세스 권한을 가지고 있으며 판매자 2는 Loggi에만 액세스할 수 있습니다.

---

### Shopee는 언제든지 판매자의 배송 채널 액세스를 활성화하거나 비활성화할 수 있습니다.

A: 예

---

### 판매자는 Shopee Geren Invoice Issue를 통해 원하는 제한으로 NF-e를 입력할 수 있습니다.

A: 예

---

### NF-e 데이터를 받으면 다음을 평가합니다.

1. 주문 총액(예: item_price * qty)이 주문서의 총액과 동일한지 확인합니다. 최소 키가 item_price에 잘못된 정보인 경우
2. 전송된 품목(sku만 해당)이 주문에 있는지 확인합니다(다른 품목은 무시됨).
3. NF-e 번호와 판매자가 Shopee 등록(CNPJ 및 UF) 또는 NF-e의 다른 필드에 제공한 데이터 간의 관계가 null인지 확인합니다(NF-e 번호만 유효성 검사됨).
4. NF-e 날짜가 주문 결제 날짜 이후인지 확인합니다(완료된 경우).

---

### 참고 자료:

1. V1: https://open.shopee.com/documents?module=63
2. V2: https://open.shopee.com/documents?version=2

## 사용 사례

1. 배송 프로세스 자동화
2. 배송 문서 생성
3. 주문 상태 추적
4. 다양한 배송 업체와 통합
5. 배송 채널 관리

## 관련 API

- v2.logistics.get_shipping_parameter
- v2.logistics.ship_order
- v2.logistics.get_tracking_number
- v2.logistics.get_airway_document_parameter
- v2.logistics.create_shipping_document
- v2.logistics.get_shipping_document_result
- v2.logistics.download_shipping_document
- v2.logistics.get_channel_list
- v2.product.edit
- v2.product.update_price
- v2.order.get_order_detail
- v2.order.add_invoice_data

---

## 원문 (English)

### Summary

This guide provides a step-by-step walkthrough of the Logistics API, detailing the flow of calls and best practices for integrating with Shopee's logistics services. It covers various scenarios, including shipping document generation, tracking, and handling different shipping channels.

### Content

# Passo a Passo Logística API

## API Best Practices > Brand Coop API Best Prac... > Passo a Passo Logística API

### Flow Diagram

```
v2.logistics.get_shipping_parameter
          ↓
      Drop Off ←→ Pick Up
          ↓
   v2.logistics.ship_order
          ↓
v2.logistics.get_tracking_number
          ↓
v2.logistics.get_airway_document_parameter
          ↓
v2.logistics.create_shipping_document
          ↓
v2.logistics.get_shipping_document_result
          ↓
v2.logistics.download_shipping_document
          ↓
         Fim
```

### Notes and Information Boxes

**ℹ️ Note 1:**
Colete informações para organizar o envio. Endereço dos modes. Dose da ou vendedor que o pedido possa ser local enrico sem fazer verificações.

**ℹ️ Note 2:**
Depois de chamar essa API, o status do order para atualizará automaticamente para "READY_TO_SHIP" (Pós no Shopee). Use "v2.order.get_order_list" para chamar a lista de "PROCESSED".

**ℹ️ Note 3:**
Pode ser chamado apenas mais cedo num o status o "PROCESSED". Se houver uma tentativa mal-sucedida (api terá um erro), não afectará o pedido do vendedor.

**ℹ️ Note 4:**
Podem ser dos tipos Normal AWB e Thermal AWB. Se a sua empresa de logistica não tiver que receber documento de envio, pode pular a etapa de obtenção documento e ir diretamente para a etapa de download de documento de envio.

**ℹ️ Note 5:**
A Etiqueta do pode ser criada depois que o TrackingNo for generated.

**ℹ️ Note 6:**
Retorna quando essa API retornar o status de todos os documentos de envio forem exibidos, você poderá baixar todos os documentos.

**ℹ️ Note 7:**
Se a etapa 6 não fornecer resultado, continue a chamar "v2.logistics.get_shipping_document_result" para obter o CDF para download. Se os TrackingNo,_init não são retornados, pode definir esse valor como nulo.

---

## O fluxo de chamadas na v2 opti relaciona aos canais de envio é conforme:

### 1. A Shopee habilita o canal de envio para o vendedor (ex: Correios, aberto a todos os vendedores, Total Express, etc)

### 2. Para verificar o canal logístico habilitado em sua loja

1. Busca chamar o endpoint: v2.logistics.get_channel_list

### 3. O vendedor deve obter o canal logístico para cada item de acordo a que fez liberado para seller p a prover (shopee logistics)

1. Para habilitar/desabilitar os canais de envio do produto, v2.product.edit, item no v2.product.update_price, você prontado de logistic_info.logistic_id dentro de logistic_info.

### 4. O vendedor recebe o pedido ja com o canal de envio selecionado pelo comprador

### 5. Para saber se é preciso informar a NF-e para o pedido antes de fazer a organização de envio (ship_order), você pode:

1. Use a chamada v2.order.get_order_detail, Cada com o params "order_status" "INVOICE_PENDING", "READY_TO_SHIP"
2. Para informar a NF-e usa a chamada v2.order.add_invoice_data
3. Após validar a etapa 5, expromo v2.logistics.get_shipping_parameter, v2.logistics.ship_order com param "invoice_data"

### 6. Para fazer o ship e o drop para o nó v2.logistics.ship_order

1. Dependendo do canal de envio selecionado pelo comprador, o vendedor irá postar o pedido ou esperar a coleta
2. Para saber se o coleta (pickup) ou postagem (dropoff), v2.logistics.get_shipping_parameter
3. Dropoff: o vendedor precisa enviar para logística. Depois de usar v2.logistics.ship_order
4. Se for dropoff (mas se formos via Correios como dropoff), caso o info_needed for vazio ("dropoff" []), o request do ship_order deverá ser:

```json
{
  "order_sn": "12CR12DCXXXXX",
  
  "dropoff": []
}
```

### 8. Se for pickup (loop todos os outres não pickup, inclusive Coleta Correios) caso o info_needed tenha "pickup" ("address_id","pickup_time_id"), a request do ship_order deverá ser:

```json
{
  "order_sn": "12CR12DCXXXXX",
  
  "pickup": []
}
```

**Example with parameters:**

```json
"address_id":10,box_address_list.address_id | in_get_address_list.n.ne.id do_self_ship(log_parameter)

"pickup_time_id":10,box_address_list.time_slot_list.pickup_time_id do_self_ship(log_parameter)
```

---

## 9. Depois disso, para obter o código de rastreio basta chamar o endpoint: v2.logistics.get_tracking_number

### 10. Para obter os possíveis params para a etiqueta: v2.logistics.get_shipping_document_parameter

1. Para obter os possíveis params para a etiqueta: v2.logistics.get_shipping_document_parameter (type "NORMAL_AIR_WAYBILL") retorna que a etiqueta pode ser gerada respectivamente em formato: pdf ou .zip (contém a etiqueta em .txt, tal em formato .ZPL que é uma linguagem de contexto)

### 11. Para pedir para a Shopee gerar a etiqueta de um ou mais pedidos:

1. Chamar v2.logistics.create_shipping_document
2. E preciso enviar o parametro "tracking_number"

### 12. Para saber se a etiqueta já foi gerada ou ainda está processada

1. v2.logistics.get_shipping_document_resultHolder com response "status" "READY" ja podendo ter a etiqueta obtida

### 13. Para obter a etiqueta (AIR WAYBILL) para esse processo para Shopee:

1. v2.logistics.download_shipping_document

---

## Mais detalhes sobre o fluxo de chamadas de pedidos pode ser encontrado em nossa documentação da seção Tracking Shipment & Get TrackingNo & Print AirwayBill:

[https://open.shopee.com/developer/public/v2](https://open.shopee.com/developer/public/v2)

---

## FAQ

### Todos os envios são feitos com canais de envio (transportadoras) direto com a Shopee

R: Não, a Shopee é livre e pago pelo vendedor

---

### Todas as etiquetas de envio são geradas pela Shopee, ou vendedor ou Integradora/ERP nunca deve gerar a etiquetai?

R: Se houver uma etiqueta, a etiqueta precisa se obida da Shopee através da v2.logistics.download_shipping_document

---

### Cada vendedor possui um único canal de envio, ou seja, diferentes vendedores podem ter acesso a diferentes canais de envio

R: Sim, cada vendedor é um acesso no Correios, e o vendedor 2 possui acesso somente a Loggi

---

### A Shopee pode habilitar ou desabilitar o acesso de um vendedor a um canal de envio a qualquer momento:

R: Sim

---

### O vendedor pode entrar a NF-e cade limite que deseja pela Shopee Geren Invoice Issue

R: Sim

---

### Quando recebemos os dados da NF-e, avaliamos se:

1. Total do pedido (p.e.item_price * qty) é igual Na nota pedido, caso o minima chave seja informações Falsas em item_price
2. Os items enviados (apenas sku) estão no pedido (de outro items sera ignorado)
3. A relacao de numero de NF-e e os dados que o vendedor informou no cadastro da Shopee (CNPJ e UF), ou nos outros campos da NF-e nulos, (apenas tal que numero da NF-e sera validado)
4. Se a data da NF-e é posterior a data de pagamento do pedido (se completado)

---

### Referencias:

1. V1: https://open.shopee.com/documents?module=63
2. V2: https://open.shopee.com/documents?version=2

---

**문서 ID**: developer-guide.292
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/292
**처리 완료**: 2025-10-16T09:09:56
