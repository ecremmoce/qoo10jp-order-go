"""
사방넷 API 가이드 고급 분석기
- 헤더/데이터 구조 정확히 파싱
- XML 노드 구조 추출
- 필수/선택 필드 구분
- 기능별 분류 및 JSON 스키마 생성
"""

import pandas as pd
import json
from pathlib import Path
from typing import Dict, List, Any, Optional
import re

class AdvancedSabangnetAnalyzer:
    def __init__(self, excel_path: str):
        self.excel_path = excel_path
        self.base_dir = Path("refrence/sabangnet")
        self.apis_dir = self.base_dir / "apis"
        self.analysis_dir = self.base_dir / "analysis"
        
        # 기능별 폴더 매핑
        self.feature_mapping = {
            "주문수집": "orders",
            "송장등록": "invoices",
            "클레임수집": "claims",
            "문의사항 수집": "inquiries",
            "상품등록&수정": "products",
            "상품요약수정": "products",
            "추가상품등록&수정": "products",
            "상품 쇼핑몰별 DATA 수정": "products",
            "쇼핑몰 코드 조회": "reference",
            "상품정보고시 코드 조회": "reference",
            "카테고리": "reference"
        }
        
    def create_folder_structure(self):
        """기능별 폴더 구조 생성"""
        folders = [
            self.analysis_dir,
            self.apis_dir / "orders" / "schemas",
            self.apis_dir / "orders" / "examples",
            self.apis_dir / "invoices" / "schemas",
            self.apis_dir / "invoices" / "examples",
            self.apis_dir / "claims" / "schemas",
            self.apis_dir / "claims" / "examples",
            self.apis_dir / "inquiries" / "schemas",
            self.apis_dir / "products" / "schemas",
            self.apis_dir / "reference" / "schemas",
            self.base_dir / "restful_specs"
        ]
        
        for folder in folders:
            folder.mkdir(parents=True, exist_ok=True)
            print(f"[CREATED] {folder}")
    
    def parse_sheet_structure(self, df: pd.DataFrame, sheet_name: str) -> Dict[str, Any]:
        """시트 구조를 상세히 파싱"""
        structure = {
            "sheet_name": sheet_name,
            "feature": self.feature_mapping.get(sheet_name, "other"),
            "description": [],
            "api_endpoint": None,
            "xml_root": None,
            "header_fields": [],
            "data_fields": [],
            "authentication": {}
        }
        
        current_section = None
        node_depth = []
        
        for idx, row in df.iterrows():
            # 모든 컬럼 값 가져오기
            values = [str(v) if pd.notna(v) else None for v in row]
            
            # 설명 부분 추출 (- 로 시작하는 줄)
            if values[1] and isinstance(values[1], str) and values[1].startswith('-'):
                structure["description"].append(values[1])
            
            # API 엔드포인트 URL 추출
            if values[1] and 'https://' in str(values[1]):
                structure["api_endpoint"] = values[1]
            
            # XML 노드 구조 파싱
            if values[2]:
                node_name = str(values[2]).strip()
                
                # 루트 노드 감지
                if node_name and not node_name.startswith('/'):
                    if idx > 10 and not structure["xml_root"]:  # 헤더 이후
                        if node_name.isupper() and '_' in node_name:
                            structure["xml_root"] = node_name
                            current_section = "root"
                    
                    # HEADER 섹션
                    if node_name == "HEADER":
                        current_section = "header"
                    
                    # DATA 섹션
                    elif node_name == "DATA":
                        current_section = "data"
                    
                    # ITEM 섹션
                    elif node_name == "ITEM":
                        current_section = "item"
                    
                    # 필드 정보 추출
                    elif current_section in ["header", "data", "item"]:
                        field_info = {
                            "name": node_name,
                            "description": values[3] if len(values) > 3 and values[3] else "",
                            "required": values[4] == 'Y' if len(values) > 4 else False,
                            "note": values[5] if len(values) > 5 and values[5] else "",
                            "section": current_section
                        }
                        
                        if current_section == "header":
                            structure["header_fields"].append(field_info)
                            
                            # 인증 관련 필드 추출
                            if 'AUTH' in node_name or 'ID' in node_name:
                                structure["authentication"][node_name] = field_info
                        else:
                            structure["data_fields"].append(field_info)
        
        return structure
    
    def convert_xml_to_json_schema(self, xml_structure: Dict[str, Any]) -> Dict[str, Any]:
        """XML 구조를 JSON 스키마로 변환"""
        json_schema = {
            "type": "object",
            "title": xml_structure["sheet_name"],
            "description": " ".join(xml_structure["description"][:3]) if xml_structure["description"] else "",
            "required": [],
            "properties": {}
        }
        
        # 데이터 필드를 JSON 스키마로 변환
        for field in xml_structure["data_fields"]:
            # XML 필드명을 camelCase로 변환
            json_field_name = self.xml_to_camel_case(field["name"])
            
            json_schema["properties"][json_field_name] = {
                "type": "string",  # 기본 타입, 실제로는 더 정교하게 분석 필요
                "description": field["description"],
                "xml_name": field["name"]  # 원본 XML 필드명 매핑
            }
            
            if field["required"]:
                json_schema["required"].append(json_field_name)
        
        return json_schema
    
    def xml_to_camel_case(self, xml_name: str) -> str:
        """XML 필드명을 camelCase로 변환"""
        # 예: COMPANY_GOODS_CD -> companyGoodsCd
        parts = xml_name.lower().split('_')
        if not parts:
            return xml_name.lower()
        
        return parts[0] + ''.join(word.capitalize() for word in parts[1:])
    
    def generate_rest_endpoint(self, sheet_name: str) -> Dict[str, Any]:
        """시트 이름을 기반으로 RESTful 엔드포인트 설계"""
        endpoint_mapping = {
            "주문수집": {
                "method": "POST",
                "path": "/api/v1/orders/collect",
                "summary": "주문 정보 수집",
                "description": "사방넷 호환 주문 데이터 수신"
            },
            "송장등록": {
                "method": "POST",
                "path": "/api/v1/invoices",
                "summary": "송장 정보 등록",
                "description": "주문 송장번호 및 택배사 정보 등록"
            },
            "클레임수집": {
                "method": "POST",
                "path": "/api/v1/claims/collect",
                "summary": "클레임 정보 수집",
                "description": "취소/반품/교환 등 클레임 데이터 수신"
            },
            "문의사항 수집": {
                "method": "POST",
                "path": "/api/v1/inquiries/collect",
                "summary": "문의사항 수집",
                "description": "고객 문의사항 데이터 수신"
            }
        }
        
        return endpoint_mapping.get(sheet_name, {
            "method": "POST",
            "path": f"/api/v1/{sheet_name.lower()}",
            "summary": sheet_name,
            "description": ""
        })
    
    def analyze_all_sheets(self) -> Dict[str, Any]:
        """모든 시트 분석"""
        print("\n" + "="*60)
        print("[ANALYSIS] 사방넷 API 가이드 상세 분석 시작")
        print("="*60)
        
        excel_file = pd.ExcelFile(self.excel_path)
        all_structures = {}
        
        for sheet_name in excel_file.sheet_names:
            print(f"\n[PARSE] {sheet_name}")
            
            # header=None으로 읽어서 원본 구조 유지
            df = pd.read_excel(excel_file, sheet_name=sheet_name, header=None)
            
            # 구조 파싱
            structure = self.parse_sheet_structure(df, sheet_name)
            all_structures[sheet_name] = structure
            
            print(f"  - 설명: {len(structure['description'])}줄")
            print(f"  - XML 루트: {structure['xml_root']}")
            print(f"  - 헤더 필드: {len(structure['header_fields'])}개")
            print(f"  - 데이터 필드: {len(structure['data_fields'])}개")
            
            # 기능별로 저장
            self.save_feature_analysis(structure)
        
        return all_structures
    
    def save_feature_analysis(self, structure: Dict[str, Any]):
        """기능별 분석 결과 저장"""
        feature = structure["feature"]
        sheet_name = structure["sheet_name"]
        
        # 마크다운 문서 생성
        md_content = self.generate_feature_markdown(structure)
        md_path = self.apis_dir / feature / f"{sheet_name.replace('&', '_').replace(' ', '_')}.md"
        
        with open(md_path, 'w', encoding='utf-8') as f:
            f.write(md_content)
        
        print(f"  [SAVED] {md_path}")
        
        # XML 스키마 저장
        schema_data = {
            "xml_root": structure["xml_root"],
            "header_fields": structure["header_fields"],
            "data_fields": structure["data_fields"],
            "authentication": structure["authentication"]
        }
        
        schema_path = self.apis_dir / feature / "schemas" / f"{sheet_name.replace('&', '_').replace(' ', '_')}_xml.json"
        with open(schema_path, 'w', encoding='utf-8') as f:
            json.dump(schema_data, f, ensure_ascii=False, indent=2)
        
        # JSON 스키마 생성
        json_schema = self.convert_xml_to_json_schema(structure)
        json_schema_path = self.apis_dir / feature / "schemas" / f"{sheet_name.replace('&', '_').replace(' ', '_')}_json.json"
        
        with open(json_schema_path, 'w', encoding='utf-8') as f:
            json.dump(json_schema, f, ensure_ascii=False, indent=2)
    
    def generate_feature_markdown(self, structure: Dict[str, Any]) -> str:
        """기능별 마크다운 문서 생성"""
        lines = []
        
        lines.append(f"# {structure['sheet_name']}\n")
        lines.append(f"**분류**: {structure['feature']}\n\n")
        
        # 설명
        if structure["description"]:
            lines.append("## 개요\n\n")
            for desc in structure["description"]:
                lines.append(f"{desc}\n\n")
        
        # API 엔드포인트
        if structure["api_endpoint"]:
            lines.append("## 사방넷 원본 엔드포인트\n\n")
            lines.append(f"```\n{structure['api_endpoint']}\n```\n\n")
        
        # RESTful 변환
        rest_endpoint = self.generate_rest_endpoint(structure["sheet_name"])
        if rest_endpoint:
            lines.append("## RESTful API 변환\n\n")
            lines.append(f"**Method**: `{rest_endpoint.get('method', 'POST')}`\n\n")
            lines.append(f"**Path**: `{rest_endpoint.get('path', '')}`\n\n")
            lines.append(f"**Summary**: {rest_endpoint.get('summary', '')}\n\n")
        
        # 인증
        if structure["authentication"]:
            lines.append("## 인증 (Authentication)\n\n")
            lines.append("### 사방넷 XML 헤더\n\n")
            lines.append("| 필드명 | 설명 | 필수 |\n")
            lines.append("|--------|------|------|\n")
            for field_name, field_info in structure["authentication"].items():
                lines.append(f"| {field_name} | {field_info['description']} | {'Y' if field_info['required'] else 'N'} |\n")
            lines.append("\n")
            
            lines.append("### RESTful HTTP 헤더\n\n")
            lines.append("```http\n")
            lines.append("Authorization: Bearer {API_KEY}\n")
            lines.append("X-Company-ID: {COMPANY_ID}\n")
            lines.append("Content-Type: application/json\n")
            lines.append("```\n\n")
        
        # 헤더 필드
        if structure["header_fields"]:
            lines.append("## 헤더 필드 (Header Fields)\n\n")
            lines.append("| 필드명 | 설명 | 필수 | 비고 |\n")
            lines.append("|--------|------|------|------|\n")
            for field in structure["header_fields"]:
                required = "Y" if field["required"] else "N"
                note = field["note"] or "-"
                lines.append(f"| {field['name']} | {field['description']} | {required} | {note} |\n")
            lines.append("\n")
        
        # 데이터 필드
        if structure["data_fields"]:
            lines.append("## 데이터 필드 (Data Fields)\n\n")
            lines.append("| XML 필드명 | JSON 필드명 | 설명 | 필수 | 비고 |\n")
            lines.append("|-----------|-------------|------|------|------|\n")
            for field in structure["data_fields"][:50]:  # 처음 50개만
                json_name = self.xml_to_camel_case(field['name'])
                required = "Y" if field["required"] else "N"
                note = field["note"] or "-"
                desc = field["description"] or "-"
                lines.append(f"| {field['name']} | {json_name} | {desc} | {required} | {note} |\n")
            
            if len(structure["data_fields"]) > 50:
                lines.append(f"\n*... 외 {len(structure['data_fields']) - 50}개 필드*\n")
            lines.append("\n")
        
        # XML 예제
        lines.append("## XML 구조 예제\n\n")
        lines.append("```xml\n")
        if structure["xml_root"]:
            lines.append(f"<{structure['xml_root']}>\n")
            lines.append("  <HEADER>\n")
            for field in structure["header_fields"][:3]:
                lines.append(f"    <{field['name']}>{field['description']}</{field['name']}>\n")
            lines.append("  </HEADER>\n")
            lines.append("  <DATA>\n")
            lines.append("    <ITEM>\n")
            for field in structure["data_fields"][:5]:
                lines.append(f"      <{field['name']}>값</{field['name']}>\n")
            lines.append("    </ITEM>\n")
            lines.append("  </DATA>\n")
            lines.append(f"</{structure['xml_root']}>\n")
        lines.append("```\n\n")
        
        # JSON 예제
        lines.append("## JSON 구조 예제\n\n")
        lines.append("```json\n")
        json_example = {}
        for field in structure["data_fields"][:5]:
            json_name = self.xml_to_camel_case(field['name'])
            json_example[json_name] = "값"
        lines.append(json.dumps(json_example, ensure_ascii=False, indent=2))
        lines.append("\n```\n")
        
        return "".join(lines)
    
    def run(self):
        """전체 분석 실행"""
        # 폴더 구조 생성
        self.create_folder_structure()
        
        # 모든 시트 분석
        all_structures = self.analyze_all_sheets()
        
        # 전체 요약 저장
        summary = {
            "total_sheets": len(all_structures),
            "by_feature": {},
            "structures": {}
        }
        
        for sheet_name, structure in all_structures.items():
            feature = structure["feature"]
            if feature not in summary["by_feature"]:
                summary["by_feature"][feature] = []
            summary["by_feature"][feature].append(sheet_name)
            
            # 구조 정보만 저장 (간소화)
            summary["structures"][sheet_name] = {
                "feature": feature,
                "xml_root": structure["xml_root"],
                "header_count": len(structure["header_fields"]),
                "data_count": len(structure["data_fields"]),
                "has_auth": len(structure["authentication"]) > 0
            }
        
        # 분석 디렉토리로 이동
        with open(self.analysis_dir / "detailed_analysis.json", 'w', encoding='utf-8') as f:
            json.dump(summary, f, ensure_ascii=False, indent=2)
        
        print("\n" + "="*60)
        print("[COMPLETE] 분석 완료")
        print("="*60)
        print(f"총 시트: {summary['total_sheets']}")
        print(f"기능별 분류:")
        for feature, sheets in summary["by_feature"].items():
            print(f"  - {feature}: {len(sheets)}개")
        
        return summary


if __name__ == "__main__":
    analyzer = AdvancedSabangnetAnalyzer("refrence/사방넷 API 가이드.xlsx")
    result = analyzer.run()

