CREATE TABLE measures (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    quantity FLOAT(24) NOT NULL,
    unit_coast FLOAT(24) NOT NULL,
    measure INTEGER REFERENCES measures (id)
);