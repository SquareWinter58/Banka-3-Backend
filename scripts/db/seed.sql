INSERT INTO permissions (name)
VALUES
    ('admin'),
    ('trade_stocks'),
    ('view_stocks'),
    ('manage_contracts'),
    ('manage_insurance')
ON CONFLICT (name) DO NOTHING;

-- default admin (password: "Admin123!" hashed with placeholder, replace with real hash)
INSERT INTO employees (
    first_name, last_name, date_of_birth, gender, email,
    phone_number, address, username, password, salt_password,
    position, department, active
)
VALUES (
    'Admin', 'Admin', '1990-01-01', 'M', 'admin@banka.raf',
    '+381600000000', 'N/A', 'admin', '\x00'::BYTEA, '\x00'::BYTEA,
    'Administrator', 'IT', true
)
ON CONFLICT (email) DO NOTHING;

-- give admin user the admin permission
INSERT INTO employee_permissions (employee_id, permission_id)
SELECT e.id, p.id
FROM employees e, permissions p
WHERE e.email = 'admin@banka.raf' AND p.name = 'admin'
ON CONFLICT DO NOTHING;

-- test client
INSERT INTO clients (
    first_name, last_name, date_of_birth, gender, email,
    phone_number, address, password, salt_password
)
VALUES (
    'Petar', 'Petrovic', '1990-05-20', 'M', 'petar@primer.raf',
    '+381645555555', 'Njegoseva 25', '\x00'::BYTEA, '\x00'::BYTEA
)
ON CONFLICT (email) DO NOTHING;
