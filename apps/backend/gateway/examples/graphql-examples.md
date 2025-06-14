# GraphQL API Examples for Zplus SaaS

This document provides examples of GraphQL queries, mutations, and subscriptions for the Zplus SaaS multi-tenant platform.

## Authentication

All GraphQL requests require proper authentication and tenant context.

### Headers Required
```http
Authorization: Bearer YOUR_JWT_TOKEN
X-Tenant-ID: your-tenant-slug
Content-Type: application/json
```

## Basic Queries

### Get Current User Information
```graphql
query GetCurrentUser {
  me {
    id
    email
    firstName
    lastName
    tenantId
    status
    roles {
      id
      name
      permissions {
        id
        name
        resource
        action
      }
    }
    createdAt
    updatedAt
  }
}
```

### Get System Information (Admin Only)
```graphql
query GetSystemInfo {
  systemInfo {
    version
    environment
    tenantCount
    uptime
  }
}
```

### List Users with Filtering
```graphql
query GetUsers($filter: UserFilter, $pagination: Pagination) {
  users(filter: $filter, pagination: $pagination) {
    edges {
      node {
        id
        email
        firstName
        lastName
        status
        roles {
          name
        }
      }
      cursor
    }
    pageInfo {
      hasNextPage
      hasPreviousPage
      startCursor
      endCursor
    }
    totalCount
  }
}
```

Variables:
```json
{
  "filter": {
    "status": "ACTIVE",
    "search": "john"
  },
  "pagination": {
    "first": 10,
    "after": null
  }
}
```

## CRM Queries

### List Customers
```graphql
query GetCustomers($filter: CustomerFilter, $pagination: Pagination) {
  customers(filter: $filter, pagination: $pagination) {
    edges {
      node {
        id
        name
        email
        phone
        company
        status
        tags
        createdBy {
          firstName
          lastName
        }
        createdAt
      }
    }
    pageInfo {
      hasNextPage
      endCursor
    }
    totalCount
  }
}
```

### Get Single Customer
```graphql
query GetCustomer($id: ID!) {
  customer(id: $id) {
    id
    name
    email
    phone
    address
    company
    status
    tags
    notes
    createdBy {
      firstName
      lastName
      email
    }
    createdAt
    updatedAt
  }
}
```

## HRM Queries

### List Employees
```graphql
query GetEmployees($filter: EmployeeFilter) {
  employees(filter: $filter) {
    edges {
      node {
        id
        employeeId
        firstName
        lastName
        email
        position
        department {
          name
        }
        status
        hireDate
      }
    }
  }
}
```

### Get Departments
```graphql
query GetDepartments {
  departments {
    id
    name
    description
    manager {
      firstName
      lastName
    }
    employees {
      id
      firstName
      lastName
      position
    }
  }
}
```

## POS/Inventory Queries

### List Products
```graphql
query GetProducts($filter: ProductFilter) {
  products(filter: $filter) {
    edges {
      node {
        id
        sku
        name
        description
        price
        cost
        stock
        category {
          name
        }
        status
        images
      }
    }
  }
}
```

## Mutations

### Create Customer
```graphql
mutation CreateCustomer($input: CreateCustomerInput!) {
  createCustomer(input: $input) {
    id
    name
    email
    phone
    company
    status
    createdAt
  }
}
```

Variables:
```json
{
  "input": {
    "name": "John Doe",
    "email": "john.doe@example.com",
    "phone": "+1-555-0123",
    "company": "ACME Corp",
    "tags": ["prospect", "enterprise"]
  }
}
```

### Update Customer
```graphql
mutation UpdateCustomer($id: ID!, $input: UpdateCustomerInput!) {
  updateCustomer(id: $id, input: $input) {
    id
    name
    email
    status
    updatedAt
  }
}
```

### Create Employee
```graphql
mutation CreateEmployee($input: CreateEmployeeInput!) {
  createEmployee(input: $input) {
    id
    employeeId
    firstName
    lastName
    email
    position
    department {
      name
    }
    hireDate
  }
}
```

Variables:
```json
{
  "input": {
    "employeeId": "EMP001",
    "firstName": "Jane",
    "lastName": "Smith",
    "email": "jane.smith@acme.com",
    "position": "Software Engineer",
    "departmentId": "1",
    "salary": 75000,
    "hireDate": "2024-01-15T00:00:00Z"
  }
}
```

### Create Product
```graphql
mutation CreateProduct($input: CreateProductInput!) {
  createProduct(input: $input) {
    id
    sku
    name
    price
    stock
    category {
      name
    }
  }
}
```

## Subscriptions

### Real-time Notifications
```graphql
subscription GetNotifications {
  notifications {
    id
    type
    title
    message
    read
    createdAt
  }
}
```

### Live Statistics
```graphql
subscription GetLiveStats {
  liveStats {
    tenantId
    timestamp
    totalUsers
    activeUsers
    totalCustomers
    newCustomersToday
    totalEmployees
    totalProducts
    salesToday
  }
}
```

### Module Activity Feeds
```graphql
subscription GetCRMActivity {
  crmActivity {
    id
    type
    entity
    entityId
    user {
      firstName
      lastName
    }
    description
    timestamp
  }
}
```

## Error Handling

GraphQL responses include errors in a standardized format:

```json
{
  "data": null,
  "errors": [
    {
      "message": "authentication required",
      "path": ["me"],
      "extensions": {
        "code": "UNAUTHENTICATED"
      }
    }
  ]
}
```

Common error codes:
- `UNAUTHENTICATED`: User not authenticated
- `FORBIDDEN`: User lacks required permissions
- `NOT_FOUND`: Resource not found
- `INVALID_INPUT`: Input validation failed
- `TENANT_MISMATCH`: User doesn't belong to current tenant

## Pagination

All list queries support cursor-based pagination:

```graphql
query GetUsersWithPagination {
  users(pagination: { first: 10, after: "cursor_string" }) {
    edges {
      node {
        id
        email
      }
      cursor
    }
    pageInfo {
      hasNextPage
      hasPreviousPage
      startCursor
      endCursor
    }
    totalCount
  }
}
```

## Filtering

Most list queries support filtering:

```json
{
  "filter": {
    "search": "john doe",
    "status": "ACTIVE",
    "dateRange": {
      "from": "2024-01-01T00:00:00Z",
      "to": "2024-12-31T23:59:59Z"
    }
  }
}
```