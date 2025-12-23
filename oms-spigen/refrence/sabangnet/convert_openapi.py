"""OpenAPI YAML을 JSON으로 변환"""
import yaml
import json
from pathlib import Path

yaml_path = Path("refrence/sabangnet/restful_specs/openapi.yaml")
json_path = Path("refrence/sabangnet/restful_specs/openapi.json")

# YAML 읽기
with open(yaml_path, 'r', encoding='utf-8') as f:
    data = yaml.safe_load(f)

# JSON 저장
with open(json_path, 'w', encoding='utf-8') as f:
    json.dump(data, f, ensure_ascii=False, indent=2)

print(f"[SUCCESS] {yaml_path} -> {json_path}")
print(f"[INFO] OpenAPI 스펙 생성 완료")
print(f"  - YAML: {yaml_path}")
print(f"  - JSON: {json_path}")
print(f"\n[TIP] Swagger UI에서 확인하려면:")
print(f"  https://editor.swagger.io/ 에서 파일을 열어보세요")

