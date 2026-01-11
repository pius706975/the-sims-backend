CREATE TABLE employee_types (
    employee_type_id VARCHAR(50) PRIMARY KEY NOT NULL,
    employee_type_name VARCHAR NOT NULL,
    created_at TIMESTAMP,
    created_by VARCHAR,
    updated_at TIMESTAMP,
    updated_by VARCHAR
);

CREATE TABLE employment_statuses (
    employment_status_id VARCHAR(50) PRIMARY KEY NOT NULL,
    employment_status_name VARCHAR NOT NULL,
    created_at TIMESTAMP,
    created_by VARCHAR,
    updated_at TIMESTAMP,
    updated_by VARCHAR
);

CREATE TABLE employees (
    employee_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    employee_number VARCHAR NOT NULL,
    full_name VARCHAR NOT NULL,
    gender VARCHAR NOT NULL,
    birth_place VARCHAR NOT NULL,
    birth_date DATE NOT NULL,
    religion VARCHAR NOT NULL,
    marital_status VARCHAR NOT NULL,
    address TEXT NOT NULL,
    phone VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    identify_card_number BIGINT NOT NULL,
    join_date DATE,
    end_date DATE,
    is_activated BOOLEAN DEFAULT TRUE,
    employee_type_id VARCHAR,
    employment_status_id VARCHAR,
    is_deleted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP,
    created_by VARCHAR,
    updated_at TIMESTAMP,
    updated_by VARCHAR,
    CONSTRAINT fk_employee_type FOREIGN KEY(employee_type_id) REFERENCES employee_types(employee_type_id),
    CONSTRAINT fk_employment_status FOREIGN KEY(employment_status_id) REFERENCES employment_statuses(employment_status_id)
);

CREATE TABLE positions (
    position_id VARCHAR(50) PRIMARY KEY NOT NULL,
    position_name VARCHAR NOT NULL,
    created_at TIMESTAMP,
    created_by VARCHAR,
    updated_at TIMESTAMP,
    updated_by VARCHAR
);

CREATE TABLE employee_positions (
    employee_positions_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    employee_id UUID NOT NULL,
    position_id VARCHAR NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP,
    created_by VARCHAR,
    updated_at TIMESTAMP,
    updated_by VARCHAR,
    CONSTRAINT fk_employee FOREIGN KEY(employee_id) REFERENCES employees(employee_id),
    CONSTRAINT fk_position FOREIGN KEY(position_id) REFERENCES positions(position_id)
);