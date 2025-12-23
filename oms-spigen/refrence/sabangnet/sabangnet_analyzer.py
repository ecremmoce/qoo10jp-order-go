"""
사방넷 API 가이드 분석 및 문서화 도구
엑셀 파일에서 API 정보를 추출하여 JSON과 마크다운으로 변환
"""

import pandas as pd
import json
from pathlib import Path
from datetime import datetime
from typing import Dict, List, Any
import re

class SabangnetAPIAnalyzer:
    def __init__(self, excel_path: str):
        self.excel_path = excel_path
        self.output_dir = Path("refrence/sabangnet")
        self.output_dir.mkdir(parents=True, exist_ok=True)
        self.api_data = {}
        
    def read_excel(self) -> Dict[str, pd.DataFrame]:
        """엑셀 파일의 모든 시트를 읽어옴"""
        try:
            excel_file = pd.ExcelFile(self.excel_path)
            sheets = {}
            print(f"[INFO] 발견된 시트: {excel_file.sheet_names}")
            
            for sheet_name in excel_file.sheet_names:
                df = pd.read_excel(excel_file, sheet_name=sheet_name)
                sheets[sheet_name] = df
                print(f"  [OK] {sheet_name}: {len(df)} rows, {len(df.columns)} columns")
            
            return sheets
        except Exception as e:
            print(f"[ERROR] 엑셀 파일 읽기 오류: {e}")
            return {}
    
    def analyze_sheet_structure(self, df: pd.DataFrame, sheet_name: str) -> Dict[str, Any]:
        """시트 구조 분석"""
        structure = {
            "sheet_name": sheet_name,
            "total_rows": len(df),
            "total_columns": len(df.columns),
            "columns": list(df.columns),
            "sample_data": df.head(3).to_dict('records') if len(df) > 0 else [],
            "column_types": {col: str(df[col].dtype) for col in df.columns}
        }
        return structure
    
    def extract_api_info(self, sheets: Dict[str, pd.DataFrame]) -> Dict[str, Any]:
        """API 정보 추출"""
        api_info = {
            "metadata": {
                "source": "사방넷 API 가이드",
                "extracted_at": datetime.now().isoformat(),
                "total_sheets": len(sheets)
            },
            "sheets": {},
            "apis": []
        }
        
        for sheet_name, df in sheets.items():
            print(f"\n[ANALYZE] {sheet_name} 분석 중...")
            
            # 시트 구조 분석
            structure = self.analyze_sheet_structure(df, sheet_name)
            api_info["sheets"][sheet_name] = structure
            
            # API 엔드포인트 추출 시도
            apis = self.parse_api_endpoints(df, sheet_name)
            if apis:
                api_info["apis"].extend(apis)
                print(f"  [OK] {len(apis)}개의 API 발견")
        
        return api_info
    
    def parse_api_endpoints(self, df: pd.DataFrame, sheet_name: str) -> List[Dict[str, Any]]:
        """API 엔드포인트 파싱"""
        apis = []
        
        # 일반적인 API 정보 컬럼명들
        possible_columns = {
            'api_name': ['API명', 'API 이름', 'API Name', 'name', '명칭', 'API'],
            'endpoint': ['엔드포인트', 'Endpoint', 'URL', 'Path', '경로'],
            'method': ['메소드', 'Method', 'HTTP Method', 'HTTP 메소드'],
            'description': ['설명', 'Description', '상세', '비고', 'Desc'],
            'parameters': ['파라미터', 'Parameters', '매개변수', 'Params'],
            'response': ['응답', 'Response', '리턴', 'Return'],
            'example': ['예제', 'Example', '샘플', 'Sample']
        }
        
        # 컬럼 매핑
        column_map = {}
        for key, possible_names in possible_columns.items():
            for col in df.columns:
                if any(name.lower() in str(col).lower() for name in possible_names):
                    column_map[key] = col
                    break
        
        print(f"  [MAP] 매핑된 컬럼: {column_map}")
        
        # 데이터 추출
        for idx, row in df.iterrows():
            # 빈 행 스킵
            if row.isna().all():
                continue
            
            api_entry = {
                "id": f"{sheet_name}_{idx}",
                "sheet": sheet_name,
                "row_index": int(idx)
            }
            
            # 매핑된 컬럼에서 데이터 추출
            for key, col_name in column_map.items():
                value = row.get(col_name)
                if pd.notna(value):
                    api_entry[key] = str(value).strip()
            
            # 모든 컬럼 데이터 저장 (원본)
            api_entry["raw_data"] = {}
            for col in df.columns:
                value = row[col]
                if pd.notna(value):
                    api_entry["raw_data"][str(col)] = str(value).strip()
            
            if api_entry.get("raw_data"):  # 데이터가 있는 경우만 추가
                apis.append(api_entry)
        
        return apis
    
    def save_json(self, data: Dict[str, Any], filename: str):
        """JSON 파일로 저장"""
        output_path = self.output_dir / filename
        with open(output_path, 'w', encoding='utf-8') as f:
            json.dump(data, f, ensure_ascii=False, indent=2)
        print(f"[SAVED] JSON 저장: {output_path}")
        return output_path
    
    def generate_markdown_docs(self, api_info: Dict[str, Any]) -> str:
        """마크다운 문서 생성"""
        md_content = []
        
        # 헤더
        md_content.append("# 사방넷 API 가이드\n")
        md_content.append(f"**추출 시간**: {api_info['metadata']['extracted_at']}\n")
        md_content.append(f"**총 시트 수**: {api_info['metadata']['total_sheets']}\n")
        md_content.append(f"**총 API 수**: {len(api_info['apis'])}\n\n")
        
        # 목차
        md_content.append("## 📑 목차\n")
        for sheet_name in api_info['sheets'].keys():
            md_content.append(f"- [{sheet_name}](#{sheet_name.lower().replace(' ', '-')})\n")
        md_content.append("\n---\n\n")
        
        # 시트별 상세 정보
        for sheet_name, sheet_info in api_info['sheets'].items():
            md_content.append(f"## {sheet_name}\n\n")
            md_content.append(f"**행 수**: {sheet_info['total_rows']} | ")
            md_content.append(f"**열 수**: {sheet_info['total_columns']}\n\n")
            
            # 컬럼 정보
            md_content.append("### 📋 컬럼 구조\n\n")
            md_content.append("| 컬럼명 | 데이터 타입 |\n")
            md_content.append("|--------|-------------|\n")
            for col, dtype in sheet_info['column_types'].items():
                md_content.append(f"| {col} | {dtype} |\n")
            md_content.append("\n")
            
            # 해당 시트의 API 목록
            sheet_apis = [api for api in api_info['apis'] if api['sheet'] == sheet_name]
            if sheet_apis:
                md_content.append(f"### 🔌 API 목록 ({len(sheet_apis)}개)\n\n")
                
                for api in sheet_apis:
                    md_content.append(f"#### API #{api['row_index']}\n\n")
                    
                    # 기본 정보
                    if 'api_name' in api:
                        md_content.append(f"**API명**: {api['api_name']}\n\n")
                    if 'endpoint' in api:
                        md_content.append(f"**엔드포인트**: `{api['endpoint']}`\n\n")
                    if 'method' in api:
                        md_content.append(f"**메소드**: `{api['method']}`\n\n")
                    if 'description' in api:
                        md_content.append(f"**설명**: {api['description']}\n\n")
                    
                    # 원본 데이터 테이블
                    if api.get('raw_data'):
                        md_content.append("**상세 정보**:\n\n")
                        md_content.append("| 항목 | 값 |\n")
                        md_content.append("|------|----|\n")
                        for key, value in api['raw_data'].items():
                            # 긴 값은 줄바꿈 처리
                            if len(str(value)) > 100:
                                value = str(value)[:100] + "..."
                            md_content.append(f"| {key} | {value} |\n")
                        md_content.append("\n")
                    
                    md_content.append("---\n\n")
            
            md_content.append("\n")
        
        return "".join(md_content)
    
    def save_markdown(self, content: str, filename: str):
        """마크다운 파일 저장"""
        output_path = self.output_dir / filename
        with open(output_path, 'w', encoding='utf-8') as f:
            f.write(content)
        print(f"[SAVED] 마크다운 저장: {output_path}")
        return output_path
    
    def run_analysis(self):
        """전체 분석 실행"""
        print("=" * 60)
        print("[START] 사방넷 API 가이드 분석 시작")
        print("=" * 60)
        
        # 1. 엑셀 읽기
        sheets = self.read_excel()
        if not sheets:
            print("[ERROR] 엑셀 파일을 읽을 수 없습니다.")
            return None
        
        # 2. API 정보 추출
        print("\n" + "=" * 60)
        print("[EXTRACT] API 정보 추출 중...")
        print("=" * 60)
        api_info = self.extract_api_info(sheets)
        
        # 3. JSON 저장
        print("\n" + "=" * 60)
        print("[SAVE] 데이터 저장 중...")
        print("=" * 60)
        json_path = self.save_json(api_info, "api_guide_full.json")
        
        # 4. 마크다운 생성
        md_content = self.generate_markdown_docs(api_info)
        md_path = self.save_markdown(md_content, "API_GUIDE.md")
        
        # 5. 요약 정보 생성
        summary = {
            "total_sheets": len(sheets),
            "total_apis": len(api_info['apis']),
            "sheets": list(sheets.keys()),
            "output_files": {
                "json": str(json_path),
                "markdown": str(md_path)
            }
        }
        self.save_json(summary, "analysis_summary.json")
        
        print("\n" + "=" * 60)
        print("[COMPLETE] 분석 완료!")
        print("=" * 60)
        print(f"[DIR] 출력 디렉토리: {self.output_dir}")
        print(f"[INFO] 총 시트 수: {summary['total_sheets']}")
        print(f"[INFO] 총 API 수: {summary['total_apis']}")
        print(f"[FILES] 생성된 파일:")
        print(f"  - {json_path}")
        print(f"  - {md_path}")
        print(f"  - {self.output_dir / 'analysis_summary.json'}")
        
        return api_info


if __name__ == "__main__":
    analyzer = SabangnetAPIAnalyzer("refrence/사방넷 API 가이드.xlsx")
    result = analyzer.run_analysis()

