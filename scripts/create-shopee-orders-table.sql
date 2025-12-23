-- Shopee Orders Table
-- This table stores order information collected from Shopee API

CREATE TABLE IF NOT EXISTS shopee_orders (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    
    -- Order identification
    order_sn VARCHAR(255) NOT NULL UNIQUE,
    platform_account_id VARCHAR(255) NOT NULL,
    
    -- Order status and timestamps
    order_status VARCHAR(100) NOT NULL,
    create_time TIMESTAMPTZ NOT NULL,
    update_time TIMESTAMPTZ NOT NULL,
    
    -- Buyer information
    buyer_user_id BIGINT,
    buyer_username VARCHAR(255),
    
    -- Recipient information
    recipient_name VARCHAR(255),
    recipient_phone VARCHAR(100),
    recipient_address TEXT,
    recipient_district VARCHAR(255),
    recipient_city VARCHAR(255),
    recipient_state VARCHAR(255),
    recipient_country VARCHAR(100),
    recipient_zipcode VARCHAR(50),
    
    -- Order financial information
    total_amount DECIMAL(12, 2),
    currency VARCHAR(10),
    payment_method VARCHAR(100),
    
    -- Shipping information
    shipping_carrier VARCHAR(255),
    tracking_number VARCHAR(255),
    
    -- Items (stored as JSONB for flexibility)
    items_json JSONB,
    
    -- Metadata
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

-- Indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_shopee_orders_order_sn ON shopee_orders(order_sn);
CREATE INDEX IF NOT EXISTS idx_shopee_orders_platform_account_id ON shopee_orders(platform_account_id);
CREATE INDEX IF NOT EXISTS idx_shopee_orders_order_status ON shopee_orders(order_status);
CREATE INDEX IF NOT EXISTS idx_shopee_orders_create_time ON shopee_orders(create_time);
CREATE INDEX IF NOT EXISTS idx_shopee_orders_buyer_username ON shopee_orders(buyer_username);
CREATE INDEX IF NOT EXISTS idx_shopee_orders_account_create_time ON shopee_orders(platform_account_id, create_time DESC);

-- Composite index for common queries
CREATE INDEX IF NOT EXISTS idx_shopee_orders_account_status ON shopee_orders(platform_account_id, order_status);

-- Updated timestamp trigger
CREATE OR REPLACE FUNCTION update_shopee_orders_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_shopee_orders_updated_at
    BEFORE UPDATE ON shopee_orders
    FOR EACH ROW
    EXECUTE FUNCTION update_shopee_orders_updated_at();

-- Comments for documentation
COMMENT ON TABLE shopee_orders IS 'Stores order information collected from Shopee API';
COMMENT ON COLUMN shopee_orders.order_sn IS 'Shopee order serial number (unique identifier)';
COMMENT ON COLUMN shopee_orders.platform_account_id IS 'Shop ID or account identifier';
COMMENT ON COLUMN shopee_orders.items_json IS 'Order items stored as JSON array';

