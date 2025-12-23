"""
사방넷 API Python 클라이언트

사용 예제:
    from sabangnet_client import SabangnetClient
    
    client = SabangnetClient(
        api_key="your_api_key",
        secret_key="your_secret_key",
        endpoint="https://api.sabangnet.co.kr"
    )
    
    # 쇼핑몰 코드 조회
    result = client.get_shop_codes()
    
    # 상품 등록
    product_data = {
        "COMPANY_GOODS_CD": "PROD001",
        "GOODS_NM": "테스트 상품",
        # ... 기타 필드
    }
    result = client.register_product(product_data)
"""

import requests
import json
import hashlib
import time
from typing import Dict, List, Optional, Any
from urllib.parse import urljoin
import logging

# 로깅 설정
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


class SabangnetAPIException(Exception):
    """사방넷 API 예외"""
    def __init__(self, message: str, status_code: Optional[int] = None, response: Optional[Dict] = None):
        self.message = message
        self.status_code = status_code
        self.response = response
        super().__init__(self.message)


class SabangnetClient:
    """사방넷 API 클라이언트"""
    
    def __init__(
        self,
        api_key: str,
        secret_key: str,
        endpoint: str = "https://api.sabangnet.co.kr",
        timeout: int = 30,
        verify_ssl: bool = True
    ):
        """
        Args:
            api_key: 사방넷 API 키
            secret_key: 사방넷 시크릿 키
            endpoint: API 엔드포인트 URL
            timeout: 요청 타임아웃 (초)
            verify_ssl: SSL 인증서 검증 여부
        """
        self.api_key = api_key
        self.secret_key = secret_key
        self.endpoint = endpoint.rstrip('/')
        self.timeout = timeout
        self.verify_ssl = verify_ssl
        self.session = requests.Session()
        
    def _generate_signature(self, params: Dict[str, Any]) -> str:
        """API 서명 생성 (구현 필요)"""
        # 실제 사방넷 API의 서명 알고리즘에 맞게 구현 필요
        # 예시: MD5 또는 SHA256 해시
        data_str = json.dumps(params, sort_keys=True)
        signature = hashlib.md5(f"{data_str}{self.secret_key}".encode()).hexdigest()
        return signature
    
    def _build_headers(self) -> Dict[str, str]:
        """요청 헤더 생성"""
        return {
            "Content-Type": "application/json",
            "Accept": "application/json",
            "X-API-Key": self.api_key,
            "User-Agent": "SabangnetPythonClient/1.0"
        }
    
    def _request(
        self,
        method: str,
        path: str,
        params: Optional[Dict] = None,
        data: Optional[Dict] = None,
        **kwargs
    ) -> Dict[str, Any]:
        """
        API 요청 실행
        
        Args:
            method: HTTP 메소드 (GET, POST, PUT, DELETE 등)
            path: API 경로
            params: URL 파라미터
            data: 요청 바디 데이터
            
        Returns:
            API 응답 딕셔너리
            
        Raises:
            SabangnetAPIException: API 요청 실패 시
        """
        url = urljoin(self.endpoint, path.lstrip('/'))
        headers = self._build_headers()
        
        # 서명 추가 (필요한 경우)
        if data:
            signature = self._generate_signature(data)
            headers["X-Signature"] = signature
        
        try:
            logger.info(f"[REQUEST] {method} {url}")
            if data:
                logger.debug(f"[DATA] {json.dumps(data, ensure_ascii=False)[:200]}")
            
            response = self.session.request(
                method=method,
                url=url,
                params=params,
                json=data,
                headers=headers,
                timeout=self.timeout,
                verify=self.verify_ssl,
                **kwargs
            )
            
            logger.info(f"[RESPONSE] Status: {response.status_code}")
            
            # 응답 확인
            if response.status_code >= 400:
                try:
                    error_data = response.json()
                except:
                    error_data = {"error": response.text}
                
                raise SabangnetAPIException(
                    f"API 요청 실패: {response.status_code}",
                    status_code=response.status_code,
                    response=error_data
                )
            
            # JSON 응답 파싱
            try:
                return response.json()
            except ValueError:
                return {"raw_response": response.text}
                
        except requests.RequestException as e:
            logger.error(f"[ERROR] {str(e)}")
            raise SabangnetAPIException(f"네트워크 오류: {str(e)}")
    
    # ========================================
    # 쇼핑몰 코드 조회 API
    # ========================================
    
    def get_shop_codes(self) -> Dict[str, Any]:
        """
        쇼핑몰 코드 조회
        사방넷에 등록된 거래중인 쇼핑몰 목록을 조회합니다.
        
        Returns:
            쇼핑몰 코드 목록
        """
        return self._request("GET", "/api/v1/shop/codes")
    
    def get_shop_info(self, shop_code: str) -> Dict[str, Any]:
        """
        특정 쇼핑몰 정보 조회
        
        Args:
            shop_code: 쇼핑몰 코드
            
        Returns:
            쇼핑몰 상세 정보
        """
        return self._request("GET", f"/api/v1/shop/{shop_code}")
    
    # ========================================
    # 상품 등록 & 수정 API
    # ========================================
    
    def register_product(self, product_data: Dict[str, Any]) -> Dict[str, Any]:
        """
        상품 등록 또는 수정
        자체상품코드(COMPANY_GOODS_CD) 기준으로 Insert 또는 Update 실행
        
        Args:
            product_data: 상품 정보 딕셔너리
                - COMPANY_GOODS_CD: 자체상품코드 (필수)
                - GOODS_NM: 상품명
                - ... 기타 필드
                
        Returns:
            등록/수정 결과
        """
        if "COMPANY_GOODS_CD" not in product_data:
            raise ValueError("COMPANY_GOODS_CD는 필수 항목입니다.")
        
        return self._request("POST", "/api/v1/product/register", data=product_data)
    
    def update_product(self, company_goods_cd: str, update_data: Dict[str, Any]) -> Dict[str, Any]:
        """
        상품 정보 수정
        
        Args:
            company_goods_cd: 자체상품코드
            update_data: 수정할 정보
            
        Returns:
            수정 결과
        """
        update_data["COMPANY_GOODS_CD"] = company_goods_cd
        return self._request("PUT", f"/api/v1/product/{company_goods_cd}", data=update_data)
    
    def get_product(self, company_goods_cd: str) -> Dict[str, Any]:
        """
        상품 정보 조회
        
        Args:
            company_goods_cd: 자체상품코드
            
        Returns:
            상품 정보
        """
        return self._request("GET", f"/api/v1/product/{company_goods_cd}")
    
    def delete_product(self, company_goods_cd: str) -> Dict[str, Any]:
        """
        상품 삭제
        
        Args:
            company_goods_cd: 자체상품코드
            
        Returns:
            삭제 결과
        """
        return self._request("DELETE", f"/api/v1/product/{company_goods_cd}")
    
    # ========================================
    # 상품 요약 수정 API
    # ========================================
    
    def update_product_summary(
        self,
        company_goods_cd: str,
        summary_data: Dict[str, Any]
    ) -> Dict[str, Any]:
        """
        상품 요약 정보 수정
        
        Args:
            company_goods_cd: 자체상품코드
            summary_data: 요약 정보 (재고, 가격 등)
            
        Returns:
            수정 결과
        """
        return self._request(
            "PATCH",
            f"/api/v1/product/{company_goods_cd}/summary",
            data=summary_data
        )
    
    # ========================================
    # 카테고리 API
    # ========================================
    
    def get_categories(self, shop_code: Optional[str] = None) -> Dict[str, Any]:
        """
        카테고리 목록 조회
        
        Args:
            shop_code: 쇼핑몰 코드 (선택)
            
        Returns:
            카테고리 목록
        """
        params = {"shop_code": shop_code} if shop_code else None
        return self._request("GET", "/api/v1/category", params=params)
    
    def get_category_tree(self, shop_code: str) -> Dict[str, Any]:
        """
        카테고리 트리 구조 조회
        
        Args:
            shop_code: 쇼핑몰 코드
            
        Returns:
            카테고리 트리
        """
        return self._request("GET", f"/api/v1/category/tree/{shop_code}")
    
    # ========================================
    # 주문 수집 API
    # ========================================
    
    def collect_orders(
        self,
        start_date: str,
        end_date: str,
        shop_code: Optional[str] = None
    ) -> Dict[str, Any]:
        """
        주문 수집
        
        Args:
            start_date: 시작일 (YYYY-MM-DD)
            end_date: 종료일 (YYYY-MM-DD)
            shop_code: 쇼핑몰 코드 (선택)
            
        Returns:
            주문 목록
        """
        params = {
            "start_date": start_date,
            "end_date": end_date
        }
        if shop_code:
            params["shop_code"] = shop_code
            
        return self._request("GET", "/api/v1/order/collect", params=params)
    
    def get_order(self, order_no: str) -> Dict[str, Any]:
        """
        주문 상세 조회
        
        Args:
            order_no: 주문번호
            
        Returns:
            주문 상세 정보
        """
        return self._request("GET", f"/api/v1/order/{order_no}")
    
    # ========================================
    # 송장 등록 API
    # ========================================
    
    def register_invoice(self, invoice_data: Dict[str, Any]) -> Dict[str, Any]:
        """
        송장 등록
        
        Args:
            invoice_data: 송장 정보
                - order_no: 주문번호
                - delivery_company: 택배사
                - invoice_no: 송장번호
                
        Returns:
            등록 결과
        """
        return self._request("POST", "/api/v1/invoice/register", data=invoice_data)
    
    def update_invoice(self, order_no: str, invoice_data: Dict[str, Any]) -> Dict[str, Any]:
        """
        송장 정보 수정
        
        Args:
            order_no: 주문번호
            invoice_data: 수정할 송장 정보
            
        Returns:
            수정 결과
        """
        return self._request("PUT", f"/api/v1/invoice/{order_no}", data=invoice_data)
    
    # ========================================
    # 클레임 수집 API
    # ========================================
    
    def collect_claims(
        self,
        start_date: str,
        end_date: str,
        claim_type: Optional[str] = None
    ) -> Dict[str, Any]:
        """
        클레임 수집
        
        Args:
            start_date: 시작일 (YYYY-MM-DD)
            end_date: 종료일 (YYYY-MM-DD)
            claim_type: 클레임 유형 (취소, 반품, 교환 등)
            
        Returns:
            클레임 목록
        """
        params = {
            "start_date": start_date,
            "end_date": end_date
        }
        if claim_type:
            params["claim_type"] = claim_type
            
        return self._request("GET", "/api/v1/claim/collect", params=params)
    
    def process_claim(self, claim_no: str, action: str, reason: Optional[str] = None) -> Dict[str, Any]:
        """
        클레임 처리
        
        Args:
            claim_no: 클레임 번호
            action: 처리 액션 (승인, 거부 등)
            reason: 처리 사유
            
        Returns:
            처리 결과
        """
        data = {"action": action}
        if reason:
            data["reason"] = reason
            
        return self._request("POST", f"/api/v1/claim/{claim_no}/process", data=data)
    
    # ========================================
    # 문의사항 수집 API
    # ========================================
    
    def collect_inquiries(
        self,
        start_date: str,
        end_date: str,
        answered: Optional[bool] = None
    ) -> Dict[str, Any]:
        """
        문의사항 수집
        
        Args:
            start_date: 시작일 (YYYY-MM-DD)
            end_date: 종료일 (YYYY-MM-DD)
            answered: 답변 여부 (True: 답변완료, False: 미답변)
            
        Returns:
            문의사항 목록
        """
        params = {
            "start_date": start_date,
            "end_date": end_date
        }
        if answered is not None:
            params["answered"] = str(answered).lower()
            
        return self._request("GET", "/api/v1/inquiry/collect", params=params)
    
    def answer_inquiry(self, inquiry_no: str, answer: str) -> Dict[str, Any]:
        """
        문의사항 답변
        
        Args:
            inquiry_no: 문의 번호
            answer: 답변 내용
            
        Returns:
            답변 등록 결과
        """
        return self._request(
            "POST",
            f"/api/v1/inquiry/{inquiry_no}/answer",
            data={"answer": answer}
        )
    
    # ========================================
    # 상품정보고시 코드 조회 API
    # ========================================
    
    def get_product_notice_codes(self) -> Dict[str, Any]:
        """
        상품정보고시 코드 조회
        
        Returns:
            상품정보고시 코드 목록
        """
        return self._request("GET", "/api/v1/product/notice-codes")
    
    # ========================================
    # 유틸리티 메소드
    # ========================================
    
    def test_connection(self) -> bool:
        """
        API 연결 테스트
        
        Returns:
            연결 성공 여부
        """
        try:
            self.get_shop_codes()
            logger.info("✓ API 연결 테스트 성공")
            return True
        except Exception as e:
            logger.error(f"✗ API 연결 테스트 실패: {e}")
            return False
    
    def close(self):
        """세션 종료"""
        self.session.close()
    
    def __enter__(self):
        """컨텍스트 매니저 진입"""
        return self
    
    def __exit__(self, exc_type, exc_val, exc_tb):
        """컨텍스트 매니저 종료"""
        self.close()


# ========================================
# 사용 예제
# ========================================

if __name__ == "__main__":
    # 클라이언트 초기화
    client = SabangnetClient(
        api_key="YOUR_API_KEY",
        secret_key="YOUR_SECRET_KEY",
        endpoint="https://api.sabangnet.co.kr"
    )
    
    try:
        # 연결 테스트
        client.test_connection()
        
        # 쇼핑몰 코드 조회
        shops = client.get_shop_codes()
        print("쇼핑몰 목록:", shops)
        
        # 상품 등록
        product_data = {
            "COMPANY_GOODS_CD": "TEST001",
            "GOODS_NM": "테스트 상품",
            "SALE_PRICE": 10000,
            "STOCK_CNT": 100
        }
        result = client.register_product(product_data)
        print("상품 등록 결과:", result)
        
        # 주문 수집
        orders = client.collect_orders("2025-10-01", "2025-10-16")
        print("주문 수집 결과:", orders)
        
    except SabangnetAPIException as e:
        print(f"API 오류: {e.message}")
        if e.response:
            print(f"응답: {e.response}")
    finally:
        client.close()


