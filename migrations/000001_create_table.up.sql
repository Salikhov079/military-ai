CREATE TABLE aistorage (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    text TEXT,
    reqtext text,
    user_id UUID,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now(),
    deleted_at BIGINT DEFAULT 0
);
