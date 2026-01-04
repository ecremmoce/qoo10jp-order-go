-- Qoo10JP Order Collector V2 - Supabase Tables
-- 주문 수집 시스템 v2 테이블 스키마
-- shopee-order-go의 shopee_orders_v2 구조를 참조하여 설계

-- ============================================
-- 1. qoo10jp_shops_v2: 상점 계정 관리
-- ============================================
CREATE TABLE IF NOT EXISTS qoo10jp_shops_v2 (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    seller_id VARCHAR NOT NULL UNIQUE,           -- Qoo10JP 판매자 ID
    shop_name VARCHAR,                           -- 상점명
    api_id VARCHAR NOT NULL,                     -- API ID (인증용)
    certification_key TEXT NOT NULL,             -- Certification Key (암호화 저장)
    region VARCHAR DEFAULT 'JP',                 -- 지역 (JP, SG 등)
    is_active BOOLEAN DEFAULT true,              -- 활성 상태
    last_sync_at TIMESTAMP WITH TIME ZONE,       -- 마지막 동기화 시간
    platform_account_id VARCHAR,                 -- 플랫폼 계정 연결 ID
    token_expire_at TIMESTAMP WITH TIME ZONE,    -- 토큰 만료 시간 (해당되는 경우)
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 인덱스
CREATE INDEX IF NOT EXISTS idx_qoo10jp_shops_v2_seller_id ON qoo10jp_shops_v2(seller_id);
CREATE INDEX IF NOT EXISTS idx_qoo10jp_shops_v2_is_active ON qoo10jp_shops_v2(is_active);
CREATE INDEX IF NOT EXISTS idx_qoo10jp_shops_v2_platform_account_id ON qoo10jp_shops_v2(platform_account_id);

-- ============================================
-- 2. orders_qoo10jp_v2: 주문 정보 (확장 필드)
-- ============================================
CREATE TABLE IF NOT EXISTS orders_qoo10jp_v2 (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    
    -- 주문 기본 정보
    order_no VARCHAR NOT NULL,                   -- Qoo10JP 주문번호
    pack_no VARCHAR,                             -- 패키지 번호
    seller_id VARCHAR NOT NULL,                  -- 판매자 ID
    region VARCHAR DEFAULT 'JP',                 -- 지역
    
    -- 상태 정보
    qoo10_status VARCHAR,                        -- Qoo10JP 공식 상태 (ShippingStatus)
    oms_status VARCHAR DEFAULT 'NEW',            -- OMS 내부 상태 (NEW, PROCESSING, SHIPPED, COMPLETED, CANCELLED)
    payment_status VARCHAR,                      -- 결제 상태
    
    -- 날짜 정보
    order_date TIMESTAMP WITH TIME ZONE,         -- 주문일
    payment_date TIMESTAMP WITH TIME ZONE,       -- 결제일
    shipping_date TIMESTAMP WITH TIME ZONE,      -- 발송일
    delivered_date TIMESTAMP WITH TIME ZONE,     -- 배송완료일
    order_create_time TIMESTAMP WITH TIME ZONE,  -- 주문 생성 시간
    order_update_time TIMESTAMP WITH TIME ZONE,  -- 주문 업데이트 시간
    
    -- 금액 정보
    order_price DECIMAL(12, 2),                  -- 주문 가격
    total_amount DECIMAL(12, 2),                 -- 총 금액
    discount DECIMAL(12, 2) DEFAULT 0,           -- 할인 금액
    seller_discount DECIMAL(12, 2) DEFAULT 0,    -- 판매자 할인
    shipping_rate DECIMAL(12, 2) DEFAULT 0,      -- 배송비
    settle_price DECIMAL(12, 2),                 -- 정산 금액
    currency VARCHAR DEFAULT 'JPY',              -- 통화
    
    -- 구매자 정보
    buyer_id VARCHAR,                            -- 구매자 ID
    buyer_name VARCHAR,                          -- 구매자 이름
    buyer_kana VARCHAR,                          -- 구매자 이름 (가나)
    buyer_email VARCHAR,                         -- 구매자 이메일
    buyer_phone VARCHAR,                         -- 구매자 전화번호
    buyer_mobile VARCHAR,                        -- 구매자 휴대폰
    
    -- 수령인 정보
    recipient_name VARCHAR,                      -- 수령인 이름
    recipient_kana VARCHAR,                      -- 수령인 이름 (가나)
    recipient_phone VARCHAR,                     -- 수령인 전화번호
    recipient_mobile VARCHAR,                    -- 수령인 휴대폰
    recipient_zipcode VARCHAR,                   -- 수령인 우편번호
    recipient_address TEXT,                      -- 수령인 전체 주소
    recipient_address1 VARCHAR,                  -- 수령인 주소1
    recipient_address2 VARCHAR,                  -- 수령인 주소2
    
    -- 배송 정보
    shipping_method VARCHAR,                     -- 배송 방법
    delivery_company VARCHAR,                    -- 배송사
    tracking_number VARCHAR,                     -- 송장번호 (트래킹 번호)
    packing_no VARCHAR,                          -- 포장 번호
    seller_delivery_no VARCHAR,                  -- 판매자 배송번호
    
    -- 결제 정보
    payment_method VARCHAR,                      -- 결제 방법
    
    -- 상품 정보 (단일 상품인 경우 - 호환성)
    item_no VARCHAR,                             -- 상품 번호
    item_title VARCHAR,                          -- 상품명
    seller_item_code VARCHAR,                    -- 판매자 상품코드
    option_name VARCHAR,                         -- 옵션명
    option_code VARCHAR,                         -- 옵션코드
    order_qty INTEGER DEFAULT 1,                 -- 주문 수량
    
    -- 플래그 및 메타데이터
    flag VARCHAR,                                -- 플래그
    fulfillment_type VARCHAR,                    -- 풀필먼트 타입
    notes TEXT,                                  -- 메모
    label_printed_at TIMESTAMP WITH TIME ZONE,   -- 라벨 출력 시간
    
    -- 원본 데이터
    raw_data JSONB,                              -- API 응답 원본 저장
    
    -- 시스템 정보
    synced_at TIMESTAMP WITH TIME ZONE,          -- 마지막 동기화 시간
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    
    -- 유니크 제약조건
    CONSTRAINT uq_orders_qoo10jp_v2_order_pack UNIQUE (order_no, pack_no)
);

-- 인덱스
CREATE INDEX IF NOT EXISTS idx_orders_qoo10jp_v2_order_no ON orders_qoo10jp_v2(order_no);
CREATE INDEX IF NOT EXISTS idx_orders_qoo10jp_v2_pack_no ON orders_qoo10jp_v2(pack_no);
CREATE INDEX IF NOT EXISTS idx_orders_qoo10jp_v2_seller_id ON orders_qoo10jp_v2(seller_id);
CREATE INDEX IF NOT EXISTS idx_orders_qoo10jp_v2_qoo10_status ON orders_qoo10jp_v2(qoo10_status);
CREATE INDEX IF NOT EXISTS idx_orders_qoo10jp_v2_oms_status ON orders_qoo10jp_v2(oms_status);
CREATE INDEX IF NOT EXISTS idx_orders_qoo10jp_v2_order_date ON orders_qoo10jp_v2(order_date);
CREATE INDEX IF NOT EXISTS idx_orders_qoo10jp_v2_created_at ON orders_qoo10jp_v2(created_at);
CREATE INDEX IF NOT EXISTS idx_orders_qoo10jp_v2_tracking_number ON orders_qoo10jp_v2(tracking_number);

-- ============================================
-- 3. order_items_qoo10jp_v2: 주문 상품 정보
-- ============================================
CREATE TABLE IF NOT EXISTS order_items_qoo10jp_v2 (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    
    -- 연관 정보
    order_no VARCHAR NOT NULL,                   -- 주문번호 (orders_qoo10jp_v2 참조)
    pack_no VARCHAR,                             -- 패키지 번호
    seller_id VARCHAR NOT NULL,                  -- 판매자 ID
    
    -- 상품 기본 정보
    item_no VARCHAR,                             -- Qoo10 상품번호
    item_code VARCHAR,                           -- 판매자 상품코드
    item_name VARCHAR,                           -- 상품명
    item_title VARCHAR,                          -- 상품 타이틀
    
    -- 옵션 정보
    option_name VARCHAR,                         -- 옵션명
    option_code VARCHAR,                         -- 옵션코드
    
    -- 수량 및 가격
    quantity INTEGER DEFAULT 1,                  -- 수량
    unit_price DECIMAL(12, 2),                   -- 단가
    original_price DECIMAL(12, 2),               -- 원가
    discounted_price DECIMAL(12, 2),             -- 할인가
    total_price DECIMAL(12, 2),                  -- 총 가격
    
    -- 상품 상태
    item_status VARCHAR,                         -- 상품 상태
    
    -- 이미지
    image_url VARCHAR,                           -- 상품 이미지 URL
    
    -- 원본 데이터
    raw_data JSONB,                              -- 상품 원본 데이터
    
    -- 시스템 정보
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 인덱스
CREATE INDEX IF NOT EXISTS idx_order_items_qoo10jp_v2_order_no ON order_items_qoo10jp_v2(order_no);
CREATE INDEX IF NOT EXISTS idx_order_items_qoo10jp_v2_pack_no ON order_items_qoo10jp_v2(pack_no);
CREATE INDEX IF NOT EXISTS idx_order_items_qoo10jp_v2_seller_id ON order_items_qoo10jp_v2(seller_id);
CREATE INDEX IF NOT EXISTS idx_order_items_qoo10jp_v2_item_no ON order_items_qoo10jp_v2(item_no);
CREATE INDEX IF NOT EXISTS idx_order_items_qoo10jp_v2_item_code ON order_items_qoo10jp_v2(item_code);

-- ============================================
-- 4. qoo10jp_sync_jobs: 동기화 작업 추적
-- ============================================
CREATE TABLE IF NOT EXISTS qoo10jp_sync_jobs (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    
    -- 작업 식별
    seller_id VARCHAR NOT NULL,                  -- 판매자 ID
    job_type VARCHAR NOT NULL DEFAULT 'order_collection', -- 작업 타입
    
    -- 기간 정보
    start_date TIMESTAMP WITH TIME ZONE,         -- 수집 시작일
    end_date TIMESTAMP WITH TIME ZONE,           -- 수집 종료일
    
    -- 상태 정보
    status VARCHAR NOT NULL DEFAULT 'pending',   -- pending, running, completed, failed, cancelled
    
    -- 진행 정보
    total_collected INTEGER DEFAULT 0,           -- 수집된 총 건수
    total_saved INTEGER DEFAULT 0,               -- 저장된 건수
    total_updated INTEGER DEFAULT 0,             -- 업데이트된 건수
    total_skipped INTEGER DEFAULT 0,             -- 건너뛴 건수
    total_failed INTEGER DEFAULT 0,              -- 실패한 건수
    progress_percentage DECIMAL(5, 2) DEFAULT 0, -- 진행률 (%)
    
    -- 시간 정보
    started_at TIMESTAMP WITH TIME ZONE,         -- 시작 시간
    completed_at TIMESTAMP WITH TIME ZONE,       -- 완료 시간
    duration_ms BIGINT,                          -- 소요 시간 (밀리초)
    
    -- 결과 정보
    error_message TEXT,                          -- 에러 메시지
    result_summary JSONB,                        -- 결과 요약
    
    -- 시스템 정보
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- 인덱스
CREATE INDEX IF NOT EXISTS idx_qoo10jp_sync_jobs_seller_id ON qoo10jp_sync_jobs(seller_id);
CREATE INDEX IF NOT EXISTS idx_qoo10jp_sync_jobs_status ON qoo10jp_sync_jobs(status);
CREATE INDEX IF NOT EXISTS idx_qoo10jp_sync_jobs_job_type ON qoo10jp_sync_jobs(job_type);
CREATE INDEX IF NOT EXISTS idx_qoo10jp_sync_jobs_created_at ON qoo10jp_sync_jobs(created_at);

-- 중복 실행 방지를 위한 부분 유니크 인덱스
CREATE UNIQUE INDEX IF NOT EXISTS idx_qoo10jp_sync_jobs_active_job 
ON qoo10jp_sync_jobs(seller_id, job_type) 
WHERE status IN ('pending', 'running');

-- ============================================
-- 5. RLS (Row Level Security) 정책
-- ============================================

-- qoo10jp_shops_v2
ALTER TABLE qoo10jp_shops_v2 ENABLE ROW LEVEL SECURITY;
CREATE POLICY "Allow all operations for authenticated users" ON qoo10jp_shops_v2
    FOR ALL USING (true);

-- orders_qoo10jp_v2
ALTER TABLE orders_qoo10jp_v2 ENABLE ROW LEVEL SECURITY;
CREATE POLICY "Allow all operations for authenticated users" ON orders_qoo10jp_v2
    FOR ALL USING (true);

-- order_items_qoo10jp_v2
ALTER TABLE order_items_qoo10jp_v2 ENABLE ROW LEVEL SECURITY;
CREATE POLICY "Allow all operations for authenticated users" ON order_items_qoo10jp_v2
    FOR ALL USING (true);

-- qoo10jp_sync_jobs
ALTER TABLE qoo10jp_sync_jobs ENABLE ROW LEVEL SECURITY;
CREATE POLICY "Allow all operations for authenticated users" ON qoo10jp_sync_jobs
    FOR ALL USING (true);

-- ============================================
-- 6. 권한 부여
-- ============================================
GRANT ALL ON qoo10jp_shops_v2 TO authenticated;
GRANT SELECT ON qoo10jp_shops_v2 TO anon;

GRANT ALL ON orders_qoo10jp_v2 TO authenticated;
GRANT SELECT ON orders_qoo10jp_v2 TO anon;

GRANT ALL ON order_items_qoo10jp_v2 TO authenticated;
GRANT SELECT ON order_items_qoo10jp_v2 TO anon;

GRANT ALL ON qoo10jp_sync_jobs TO authenticated;
GRANT SELECT ON qoo10jp_sync_jobs TO anon;

-- ============================================
-- 7. 통계 뷰
-- ============================================
CREATE OR REPLACE VIEW qoo10jp_order_statistics AS
SELECT 
    DATE(order_date) as date,
    seller_id,
    COUNT(*) as total_orders,
    COUNT(*) FILTER (WHERE oms_status = 'NEW') as new_orders,
    COUNT(*) FILTER (WHERE oms_status = 'PROCESSING') as processing_orders,
    COUNT(*) FILTER (WHERE oms_status = 'SHIPPED') as shipped_orders,
    COUNT(*) FILTER (WHERE oms_status = 'COMPLETED') as completed_orders,
    COUNT(*) FILTER (WHERE oms_status = 'CANCELLED') as cancelled_orders,
    SUM(total_amount) as total_amount,
    AVG(total_amount) as avg_order_amount
FROM orders_qoo10jp_v2
GROUP BY DATE(order_date), seller_id
ORDER BY date DESC, seller_id;

GRANT SELECT ON qoo10jp_order_statistics TO authenticated;
GRANT SELECT ON qoo10jp_order_statistics TO anon;

-- ============================================
-- 8. 동기화 작업 통계 뷰
-- ============================================
CREATE OR REPLACE VIEW qoo10jp_sync_job_statistics AS
SELECT 
    DATE(created_at) as date,
    seller_id,
    COUNT(*) as total_jobs,
    COUNT(*) FILTER (WHERE status = 'completed') as successful_jobs,
    COUNT(*) FILTER (WHERE status = 'failed') as failed_jobs,
    AVG(duration_ms) as avg_duration_ms,
    SUM(total_collected) as total_orders_collected,
    SUM(total_saved) as total_orders_saved
FROM qoo10jp_sync_jobs
GROUP BY DATE(created_at), seller_id
ORDER BY date DESC;

GRANT SELECT ON qoo10jp_sync_job_statistics TO authenticated;
GRANT SELECT ON qoo10jp_sync_job_statistics TO anon;

-- ============================================
-- 9. 헬퍼 함수
-- ============================================

-- 최근 동기화 작업 조회
CREATE OR REPLACE FUNCTION get_recent_qoo10jp_sync_jobs(
    p_seller_id VARCHAR DEFAULT NULL,
    p_limit INTEGER DEFAULT 10
)
RETURNS TABLE (
    id UUID,
    seller_id VARCHAR,
    job_type VARCHAR,
    status VARCHAR,
    total_collected INTEGER,
    total_saved INTEGER,
    duration_ms BIGINT,
    started_at TIMESTAMP WITH TIME ZONE,
    completed_at TIMESTAMP WITH TIME ZONE,
    error_message TEXT,
    created_at TIMESTAMP WITH TIME ZONE
) 
LANGUAGE sql
AS $$
    SELECT 
        sj.id,
        sj.seller_id,
        sj.job_type,
        sj.status,
        sj.total_collected,
        sj.total_saved,
        sj.duration_ms,
        sj.started_at,
        sj.completed_at,
        sj.error_message,
        sj.created_at
    FROM qoo10jp_sync_jobs sj
    WHERE (p_seller_id IS NULL OR sj.seller_id = p_seller_id)
    ORDER BY sj.created_at DESC
    LIMIT p_limit;
$$;

GRANT EXECUTE ON FUNCTION get_recent_qoo10jp_sync_jobs TO authenticated;
GRANT EXECUTE ON FUNCTION get_recent_qoo10jp_sync_jobs TO anon;

-- 활성 동기화 작업 확인
CREATE OR REPLACE FUNCTION has_active_qoo10jp_sync_job(
    p_seller_id VARCHAR,
    p_job_type VARCHAR DEFAULT 'order_collection'
)
RETURNS BOOLEAN
LANGUAGE sql
AS $$
    SELECT EXISTS (
        SELECT 1 
        FROM qoo10jp_sync_jobs 
        WHERE seller_id = p_seller_id 
          AND job_type = p_job_type
          AND status IN ('pending', 'running')
    );
$$;

GRANT EXECUTE ON FUNCTION has_active_qoo10jp_sync_job TO authenticated;
GRANT EXECUTE ON FUNCTION has_active_qoo10jp_sync_job TO anon;
