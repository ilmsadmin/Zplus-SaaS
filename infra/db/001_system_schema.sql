-- System level tables
CREATE SCHEMA IF NOT EXISTS system;

-- Tenants table
CREATE TABLE system.tenants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(100) UNIQUE NOT NULL,
    domain VARCHAR(255),
    subdomain VARCHAR(100),
    plan_id UUID,
    status VARCHAR(50) DEFAULT 'active',
    settings JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Plans table
CREATE TABLE system.plans (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2),
    features JSONB DEFAULT '{}',
    max_users INTEGER,
    max_storage BIGINT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Modules table
CREATE TABLE system.modules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL,
    description TEXT,
    version VARCHAR(20) DEFAULT '1.0.0',
    enabled BOOLEAN DEFAULT true,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Tenant modules (which modules are enabled for which tenant)
CREATE TABLE system.tenant_modules (
    tenant_id UUID REFERENCES system.tenants(id) ON DELETE CASCADE,
    module_id UUID REFERENCES system.modules(id) ON DELETE CASCADE,
    enabled BOOLEAN DEFAULT true,
    configuration JSONB DEFAULT '{}',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    PRIMARY KEY (tenant_id, module_id)
);

-- Insert default modules
INSERT INTO system.modules (name, description) VALUES
('auth', 'Authentication and Authorization'),
('crm', 'Customer Relationship Management'),
('hrm', 'Human Resource Management'), 
('pos', 'Point of Sale'),
('file', 'File Management'),
('payment', 'Payment and Subscriptions');

-- Insert default plan
INSERT INTO system.plans (name, description, price, max_users, max_storage) VALUES
('Basic', 'Basic plan with essential features', 29.99, 10, 1073741824),  -- 1GB
('Pro', 'Professional plan with advanced features', 99.99, 50, 5368709120), -- 5GB
('Enterprise', 'Enterprise plan with unlimited features', 299.99, -1, -1); -- Unlimited