-- Tenant schema template
-- This will be used to create schemas for each tenant
-- Schema name: tenant_{tenant_id}

-- Users table (tenant-specific)
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL, -- Reference to system.tenants(id)
    email VARCHAR(255) NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    avatar VARCHAR(500),
    status VARCHAR(50) DEFAULT 'active', -- active, inactive, suspended
    last_login_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    
    UNIQUE(tenant_id, email)
);

-- Permissions table (system-wide, can be referenced across tenants)
CREATE TABLE permissions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) UNIQUE NOT NULL,
    resource VARCHAR(100) NOT NULL,
    action VARCHAR(50) NOT NULL,
    description TEXT,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Roles table (tenant-specific)
CREATE TABLE roles (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    is_system_role BOOLEAN DEFAULT false,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    
    UNIQUE(tenant_id, name)
);

-- User roles mapping
CREATE TABLE user_roles (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    role_id UUID REFERENCES roles(id) ON DELETE CASCADE,
    assigned_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    PRIMARY KEY (user_id, role_id)
);

-- Role permissions mapping
CREATE TABLE role_permissions (
    role_id UUID REFERENCES roles(id) ON DELETE CASCADE,
    permission_id UUID REFERENCES permissions(id) ON DELETE CASCADE,
    assigned_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    PRIMARY KEY (role_id, permission_id)
);

-- Customers table (for CRM module)
CREATE TABLE customers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    phone VARCHAR(50),
    address TEXT,
    company VARCHAR(255),
    status VARCHAR(50) DEFAULT 'lead', -- lead, prospect, active, inactive, churned
    tags JSONB DEFAULT '[]',
    notes TEXT,
    created_by UUID REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE
);

-- Departments table (for HRM module)
CREATE TABLE departments (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    manager_id UUID REFERENCES employees(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    
    UNIQUE(tenant_id, name)
);

-- Employees table (for HRM module)
CREATE TABLE employees (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    employee_id VARCHAR(50) NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone VARCHAR(50),
    department_id UUID REFERENCES departments(id),
    position VARCHAR(100) NOT NULL,
    salary DECIMAL(10,2),
    hire_date DATE NOT NULL,
    status VARCHAR(50) DEFAULT 'active', -- active, on_leave, terminated, resigned
    manager_id UUID REFERENCES employees(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    
    UNIQUE(tenant_id, employee_id),
    UNIQUE(tenant_id, email)
);

-- Product categories table (for POS module)  
CREATE TABLE product_categories (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    parent_id UUID REFERENCES product_categories(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    
    UNIQUE(tenant_id, name)
);

-- Products table (for POS module)
CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID NOT NULL,
    sku VARCHAR(100) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10,2) NOT NULL,
    cost DECIMAL(10,2),
    stock INTEGER NOT NULL DEFAULT 0,
    category_id UUID REFERENCES product_categories(id),
    images JSONB DEFAULT '[]',
    status VARCHAR(50) DEFAULT 'active', -- active, inactive, out_of_stock, discontinued
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    deleted_at TIMESTAMP WITH TIME ZONE,
    
    UNIQUE(tenant_id, sku)
);

-- Add indexes for better performance
CREATE INDEX idx_users_tenant_id ON users(tenant_id);
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_status ON users(status);
CREATE INDEX idx_roles_tenant_id ON roles(tenant_id);
CREATE INDEX idx_customers_tenant_id ON customers(tenant_id);
CREATE INDEX idx_customers_status ON customers(status);
CREATE INDEX idx_employees_tenant_id ON employees(tenant_id);
CREATE INDEX idx_employees_status ON employees(status);
CREATE INDEX idx_products_tenant_id ON products(tenant_id);
CREATE INDEX idx_products_status ON products(status);

-- Insert default permissions
INSERT INTO permissions (name, resource, action, description) VALUES
('tenants:read', 'tenants', 'read', 'Read tenants'),
('tenants:write', 'tenants', 'write', 'Write tenants'),
('users:read', 'users', 'read', 'Read users'),
('users:write', 'users', 'write', 'Write users'),
('customers:read', 'customers', 'read', 'Read customers'),
('customers:write', 'customers', 'write', 'Write customers'),
('employees:read', 'employees', 'read', 'Read employees'),
('employees:write', 'employees', 'write', 'Write employees'),
('products:read', 'products', 'read', 'Read products'),
('products:write', 'products', 'write', 'Write products'),
('roles:read', 'roles', 'read', 'Read roles'),
('roles:write', 'roles', 'write', 'Write roles');

-- Insert default roles (these will be created for each tenant)
-- Note: tenant_id will need to be replaced with actual tenant ID when creating tenant schema
INSERT INTO roles (tenant_id, name, description, is_system_role) VALUES
('00000000-0000-0000-0000-000000000000', 'tenant_admin', 'Tenant Administrator with full access', true),
('00000000-0000-0000-0000-000000000000', 'manager', 'Manager with limited administrative access', true),
('00000000-0000-0000-0000-000000000000', 'employee', 'Employee with basic access', true),
('00000000-0000-0000-0000-000000000000', 'user', 'Regular user with read-only access', true);