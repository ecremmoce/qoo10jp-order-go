"""엑셀 구조 상세 분석"""
import pandas as pd
import json

excel_path = "refrence/사방넷 API 가이드.xlsx"
excel_file = pd.ExcelFile(excel_path)

for sheet_name in excel_file.sheet_names:
    print(f"\n{'='*60}")
    print(f"시트: {sheet_name}")
    print('='*60)
    
    df = pd.read_excel(excel_file, sheet_name=sheet_name, header=None)
    
    print(f"\n[처음 15행 미리보기]")
    print(df.head(15).to_string())
    
    print(f"\n[시트 정보]")
    print(f"- 총 행 수: {len(df)}")
    print(f"- 총 열 수: {len(df.columns)}")


