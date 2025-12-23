# API 모범 사례 - AUTO PARTS: COMP/VPS

**카테고리**: 모범 사례
**난이도**: medium
**중요도**: 4/5
**최종 업데이트**: 2025-10-16T09:08:11

## 개요

본 가이드는 AUTO PARTS 호환성 API (COMP/VPS) 사용을 위한 모범 사례를 제공합니다. 브랜드, 모델, 연도 및 버전 ID를 포함한 차량 호환성을 위한 데이터 구조를 다룹니다. 또한 차량 목록 및 호환성 세부 정보를 검색하기 위한 특정 엔드포인트를 자세히 설명합니다.

## 주요 키워드

- API
- 모범 사례
- 자동차 부품
- 호환성
- 차량 데이터
- Endpoints
- Data Structure
- brand_id
- model_id
- year_id
- version_id

## 본문

# API 모범 사례 - 자동차 부품: COMP/VPS

## 소개

호환성 유형은 특정 위치와 호환되는 차량 목록을 제공합니다.

호환성 항목은 엔진 제품의 작동 목록을 제공합니다. (부품 요구 사항에 필수적이며, 위치는 GRFIN 생성이며 더 이상 생성되지 않으며, 항목에서 소비하여 차량 목록에 생성됩니다.)

위 내용을 고려하여, 고객(또는 판매자)이 적응 가능한 항목을 나열하는 경우, vocle 플랫폼의 새로운 응답을 실험하는 것이 좋습니다.

## 데이터 구조

차량 항목은 다음 형식으로 저장됩니다.

| 파라미터 | 의미 | 예시 |
|-----------|-------------|---------|
| brand_id | 특정 브랜드의 ID | 1270 |
| brand_name | 이 brand_id와 관련된 브랜드 이름 | "Chevrolet" |
| model_id | 특정 모델의 ID | 9950 |
| model_name | model_id와 관련된 모델 이름 | "Chevrolet" |
| year_id | 특정 연도의 ID | 3712 |
| year_name | 주어진 year_id와 관련된 연도 | "1972" |
| version_id | 차량의 각 변형 ID | 5607 |
| version_name | version_id와 연결된 버전 이름 | "Pickup" |

참고: hyperProduct의 결과에 따라 위에 나열된 파라미터의 ID를 사용하는 것이 중요합니다 (brand -> model -> year -> version)

## 참고 사항: productnfo

**vp_product_with_api_vehicle_list**

- 초기 엔진으로서, adaptivo 데이터베이스에서 가능한 모든 차량 목록을 가져오는 역할을 담당하는 엔드포인트입니다. 이를 보려면 플러그인된 요소가 페이지에서 생성되며, 페이지당 최대 100대의 차량을 볼 수 있습니다. 브랜드 필터에서 가능하게 할 수 있습니다.
- 요청 반환 예시:

```json
"price" : 0,
"items" : [
  {
    "response" : {
      "https://api.mercadolibre.com/items/MLM572819345?attributes=id,title,price",
      "id": "MLM572819345",
      "brand_id": 1270,
      "brand_name": "Chevrolet",
      "model_id": 9950,
      "model_name": "Chevrolet",
      "year_id": 3712,
      "year_name": "1972",
      "version_id": 5607,
      "version_name": "Pickup",
      "brand_id": 1270,
      "brand_name": "Chevrolet",
      "model_id": 9950,
      "model_name": "Chevrolet",
      "year_id": 3712,
      "year_name": "1972",
      "version_id": 5607,
      "version_name": "Pickup"
    },
    "item_id": null,
    "body": null,
  }
]
```

## vp_product_api_vehicle_list_by_compatibility_detail

- 이 엔드포인트에서는 각 차량 요소(브랜드 / 모델 / 연도) 및 버전에 대한 세부 정보를 식별할 수 있습니다. 이 플러그인된 응답은 "compatibility_details"이며, 여기서 "model_id" 모델을 지정하여 시각화할 수 있습니다. 따라서 model_id 사용 시 데이터베이스에 있는 ID 정보를 필터링할 수 있으며, API에 나열된 모든 정보를 페이징할 수 있습니다.
- 엔드포인트 호출 예시:

| 필수 | 응답 |
|----------|----------|
| compatibility_details="Brand" | |

```json
{
  "response": [
    {
      "model_id": ["v_type"]
    },
    {
      "model_id": ["v_attr"]
    }
  ],
  "total": null,
  {
    "brand_id": [],
  },
  {
    "brand_name": []
  },
  "FIAT": "",
  {
    "Result": {
    }
  },
  "FIAT": "",
  {
    "Chevrolet": ""
  }
}
```

**compatibility_brand_id=1234&compatibility_details="Model"**

```json
{
  "response": [
    {
      "model_id": ["v_type"]
    },
    {
      "model_id": ["v_attr"]
    }
  ],
  "total": null,
  {
    "by-list": "",
  },
  {
    "avto_moto"
  },
  "Onix-bx": "",
  {
    "Onix-bx"
  },
  "Classic": "",
  {
    "Dualr-plof-bx"
  },
  {
    "Dualr-classic"
  },
  "Blockbuster-bx": ""
}
```

## 확장된 엔드포인트 검색

### vp_product_with_auto

- 이 엔드포인트는 특정 호환 필터와 호환되는 모든 제품 목록을 얻는 데 사용됩니다. 또한:

**vp_product_list_by_[BRAND]_AND_item=(> [BRAND_ITEM]_AND_v= [> [BRAND]_autom_item**

처음 두 엔드포인트에서 위치는 이러한 모든 모델에 대한 호환성 정보를 전달합니다.

다음 예시 데이터를 고려하십시오.

다음 데이터 (franchise, os v. [BRAND] AND_item=(> [BRAND]_autom_item)을 고려하십시오.

```json
{
  "model_id": ["v_type"],
  "category_id": ["MLM1747"],
  "results": [
    {
      "brand_id": 1270,
      "model_year": "",
      "series_id": "Pickup",
      "series_id": 5607
    },
    {
      "sales": 0,
      "year_id": 3950,
      "year_filter": ""
    }
  ]
}
```

---

# 호환성 정보 문서

## 데이터 구조

```python
"compatibility_info": {
  "vehicle_info_list": [
    {
      "brand_id": 5770,
      "model_id": 5911,
      "year_id": 5590,
      "version_id": 5912
    },
    {
      "brand_id": 5508,
      "model_id": 5509,
      "year_id": 5515
    },
    {
      "brand_id": 5770,
      "model_id": 5905
    }
  ]
},
```

## 중요 고려 사항

- **특정 제품이 특정 연도의 모든 버전과 호환되는 경우, 브랜드, 모델 및 연도에 대한 ID만 제공하면 됩니다. 시스템은 해당 연도의 모든 버전을 호환성 목록에 삽입해야 하는 것으로 이해합니다.**

```python
{
  "brand_id": 5508,
  "model_id": 5509,
  "year_id": 5515
}
```

- **마찬가지로, 특정 모델의 모든 연도의 모든 버전이 제품과 호환되는 경우, 브랜드 및 모델에 대한 ID만 제공하면 됩니다.**

```python
{
  "brand_id": 5770,
  "model_id": 5905
}
```

- **v2 product update_item은 기존 항목에 호환성 정보가 없는 경우, 호환성 정보를 추가하는 데 정상적으로 사용할 수 있습니다. 이 경우, 위의 예와 같이 호환성 목록을 제공하면 됩니다.**

- **주의: v2 product update_item을 사용하여 목록에 이미 차량이 있는 항목에 호환성을 추가하는 경우, 항목에 현재 있는 ID 목록 + 새로운 호환성 ID를 제공해야 합니다. 그렇지 않으면 기존 ID 목록이 update item 호출에서 제공된 새로운 ID로 덮어쓰여집니다.**

**새로운 엔드포인트 사용에 대한 의문 사항이 있는 경우, 티켓을 열어 주시면 새로운 기능과 올바르게 통합하는 데 도움을 드리겠습니다.**

## 사용 사례

1. 특정 자동차 부품에 호환되는 차량 목록 검색.
2. 호환성을 위한 차량 세부 정보 (브랜드, 모델, 연도, 버전) 식별.
3. 특정 기준에 따른 차량 호환성 데이터 필터링.
4. 기존 항목에 호환성 정보 추가.

## 관련 API

- vp_product_with_api_vehicle_list
- vp_product_api_vehicle_list_by_compatibility_detail
- vp_product_with_auto
- v2 product update_item

---

## 원문 (English)

### Summary

This guide provides best practices for using the AUTO PARTS compatibility API (COMP/VPS). It covers data structures for vehicle compatibility, including brand, model, year, and version IDs. The guide also details specific endpoints for retrieving vehicle lists and compatibility details.

### Content

# API Best Practices - AUTO PARTS: COMP/VPS

## Introduction

Os tipos com compatibilidade proverão a lista de veículos compatíveis com certa posição.

Os itens com compatibilidade proverão a lista operante um produto de motor (obrigatório para o requisito da peça a ser a posição é criação a GRFIN já criação que não vão mais, trará consome um em item a um terão a criar no terão lista do com seu veículo.

Considerando acima, et a lista clientes (ou sellers) listalisem com itens de adaptáveis, recomendamos físeremos
que experimentos os novos response na plataforma de vocle.

## Estrutura dos dados

Os itens dos veículos são armazenados na seguinte forma

| Parâmetro | Significado | Exemplo |
|-----------|-------------|---------|
| brand_id | Id da marca relativa específica | 1270 |
| brand_name | O nome da marca relacionada a esta brand_id | "Chevrolet" |
| model_id | O id de um modelo específico | 9950 |
| model_name | O nome do modelo relacionado ao model_id | "Chevrolet" |
| year_id | O id de um ano específico | 3712 |
| year_name | O ano relacionado a um dado year_id | "1972" |
| version_id | A id de cada variação de um veículo | 5607 |
| version_name | O nome da versão associada ao version_id | "Pickup" |

Observe: É importante utilizar os ids dos parâmetros listados acima segundo-resultão do hyperProduct
(brand -> model -> year -> version

## Notas: productnfo

**vp_product_with_api_vehicle_list**

- Como o motor inicial, é o endpoint responsável por trazer toda a lista de veículos possíveis na base de dados do
  adaptivo. Para ver ele, o elemento pluginado será criado da página page, além que pode ver no máximo 100
  veículos por página. É possível fazer com ele será possível no filtro de marca.
- Exemplos de retorno da requisição:

```json
"price" : 0,
"items" : [
  {
    "response" : {
      "https://api.mercadolibre.com/items/MLM572819345?attributes=id,title,price",
      "id": "MLM572819345",
      "brand_id": 1270,
      "brand_name": "Chevrolet",
      "model_id": 9950,
      "model_name": "Chevrolet",
      "year_id": 3712,
      "year_name": "1972",
      "version_id": 5607,
      "version_name": "Pickup",
      "brand_id": 1270,
      "brand_name": "Chevrolet",
      "model_id": 9950,
      "model_name": "Chevrolet",
      "year_id": 3712,
      "year_name": "1972",
      "version_id": 5607,
      "version_name": "Pickup"
    },
    "item_id": null,
    "body": null,
  }
]
```

## vp_product_api_vehicle_list_by_compatibility_detail

- Neste endpoint é possível identificar as obterais para cada um dos elementos de veículo (marca / modelo / ano)
  e versão. Este response pluginado é o "compatibility_details" onde você especificar o modelo do "model_id" para
  visualização. Assim pode-se filtrar as informações os ids informados na base, no uso de model_id / o parâmetro além todos
  informações em um paging na API listado.
- Exemplos de chamada da endpoint:

| Required | Response |
|----------|----------|
| compatibility_details="Brand" | |

```json
{
  "response": [
    {
      "model_id": ["v_type"]
    },
    {
      "model_id": ["v_attr"]
    }
  ],
  "total": null,
  {
    "brand_id": [],
  },
  {
    "brand_name": []
  },
  "FIAT": "",
  {
    "Result": {
    }
  },
  "FIAT": "",
  {
    "Chevrolet": ""
  }
}
```

**compatibility_brand_id=1234&compatibility_details="Model"**

```json
{
  "response": [
    {
      "model_id": ["v_type"]
    },
    {
      "model_id": ["v_attr"]
    }
  ],
  "total": null,
  {
    "by-list": "",
  },
  {
    "avto_moto"
  },
  "Onix-bx": "",
  {
    "Onix-bx"
  },
  "Classic": "",
  {
    "Dualr-plof-bx"
  },
  {
    "Dualr-classic"
  },
  "Blockbuster-bx": ""
}
```

## Búsqueda con endpoint extendidos

### vp_product_with_auto

- Este endpoint es utilizado para obtener el listado de todos los productos que sean compatibles con determinados filtros compatibles. Además también:

**vp_product_list_by_[BRAND]_AND_item=(> [BRAND_ITEM]_AND_v= [> [BRAND]_autom_item**

Nos dois primeiros endpoints a posição passa será informações da compatibilidade para todos esses / modelos.

Considere o seguinte dado de exemplo:

Dado el siguiente para (franchise, os v. [BRAND] AND_item=(> [BRAND]_autom_item

```json
{
  "model_id": ["v_type"],
  "category_id": ["MLM1747"],
  "results": [
    {
      "brand_id": 1270,
      "model_year": "",
      "series_id": "Pickup",
      "series_id": 5607
    },
    {
      "sales": 0,
      "year_id": 3950,
      "year_filter": ""
    }
  ]
}
```

---

# Compatibility Info Documentation

## Data Structure

```python
"compatibility_info": {
  "vehicle_info_list": [
    {
      "brand_id": 5770,
      "model_id": 5911,
      "year_id": 5590,
      "version_id": 5912
    },
    {
      "brand_id": 5508,
      "model_id": 5509,
      "year_id": 5515
    },
    {
      "brand_id": 5770,
      "model_id": 5905
    }
  ]
},
```

## Important Considerations

- **Se um determinado produto é compatível com todas as versões de um determinado ano, basta informar os ids para brand, model e year; o sistema entenderá que todas as versões daquele ano deverão ser inseridas na lista de compatibilidade;**

```python
{
  "brand_id": 5508,
  "model_id": 5509,
  "year_id": 5515
}
```

- **De maneira similar, se todas as versões de todos os anos de um dado modelo forem compatíveis com o produto, basta informar os ids para brand e model;**

```python
{
  "brand_id": 5770,
  "model_id": 5905
}
```

- **O v2 product update_item poderá ser usado normalmente para adicionar as informações de compatibilidade em itens existentes que não a possuam. Nesse caso, basta informar a lista de compatibilidade conforme exemplos acima.**

- **ATENÇÃO: Ao usar o v2 product update_item para adicionar compatibilidade em um item que já possui veículos em sua lista, você deverá informar a lista de ids existente atualmente no item + os novos ids de compatibilidade. Caso contrário, a lista de ids existente será sobreposta pelos novos ids informados na chamada do update item.**

**Se restar alguma dúvida no uso dos novos endpoints, basta abrir um ticket e iremos te auxiliar a integrar corretamente com a nova funcionalidade.**

---

**문서 ID**: developer-guide.378
**플랫폼**: shopee
**URL**: https://open.shopee.com/developer-guide/378
**처리 완료**: 2025-10-16T09:08:11
