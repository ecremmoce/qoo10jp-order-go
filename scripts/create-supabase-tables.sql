-- Qoo10JP Order Scheduler - Supabase Tables
-- Create tables for job results and scheduler state

-- Job results table to track execution history
CREATE TABLE IF NOT EXISTS job_results (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    job_id VARCHAR NOT NULL,
    start_time TIMESTAMP WITH TIME ZONE NOT NULL,
    end_time TIMESTAMP WITH TIME ZONE NOT NULL,
    duration_ms BIGINT NOT NULL,
    orders_count INTEGER DEFAULT 0,
    success BOOLEAN NOT NULL DEFAULT false,
    error_msg TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create index for better query performance
CREATE INDEX IF NOT EXISTS idx_job_results_job_id ON job_results(job_id);
CREATE INDEX IF NOT EXISTS idx_job_results_created_at ON job_results(created_at);
CREATE INDEX IF NOT EXISTS idx_job_results_success ON job_results(success);

-- Add RLS (Row Level Security) policies if needed
ALTER TABLE job_results ENABLE ROW LEVEL SECURITY;

-- Create a policy that allows all operations for authenticated users
CREATE POLICY "Allow all operations for authenticated users" ON job_results
    FOR ALL USING (auth.role() = 'authenticated');

-- Create a policy that allows read access for anonymous users (optional)
CREATE POLICY "Allow read access for anonymous users" ON job_results
    FOR SELECT USING (true);

-- Grant necessary permissions
GRANT ALL ON job_results TO authenticated;
GRANT SELECT ON job_results TO anon;

-- Create a view for job statistics
CREATE OR REPLACE VIEW job_statistics AS
SELECT 
    DATE(created_at) as date,
    COUNT(*) as total_jobs,
    COUNT(*) FILTER (WHERE success = true) as successful_jobs,
    COUNT(*) FILTER (WHERE success = false) as failed_jobs,
    AVG(duration_ms) as avg_duration_ms,
    SUM(orders_count) as total_orders_processed
FROM job_results
GROUP BY DATE(created_at)
ORDER BY date DESC;

-- Grant access to the view
GRANT SELECT ON job_statistics TO authenticated;
GRANT SELECT ON job_statistics TO anon;

-- Create function to get recent job results
CREATE OR REPLACE FUNCTION get_recent_jobs(limit_count INTEGER DEFAULT 10)
RETURNS TABLE (
    id UUID,
    job_id VARCHAR,
    start_time TIMESTAMP WITH TIME ZONE,
    end_time TIMESTAMP WITH TIME ZONE,
    duration_ms BIGINT,
    orders_count INTEGER,
    success BOOLEAN,
    error_msg TEXT,
    created_at TIMESTAMP WITH TIME ZONE
) 
LANGUAGE sql
AS $$
    SELECT 
        jr.id,
        jr.job_id,
        jr.start_time,
        jr.end_time,
        jr.duration_ms,
        jr.orders_count,
        jr.success,
        jr.error_msg,
        jr.created_at
    FROM job_results jr
    ORDER BY jr.created_at DESC
    LIMIT limit_count;
$$;

-- Grant execute permission on the function
GRANT EXECUTE ON FUNCTION get_recent_jobs TO authenticated;
GRANT EXECUTE ON FUNCTION get_recent_jobs TO anon;












