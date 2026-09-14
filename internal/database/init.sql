CREATE TABLE IF NOT EXISTS Components(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    overview TEXT,
    price NUMERIC(10,2) NOT NULL,
    quantity INT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS Computers(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    overview TEXT,
    price NUMERIC(10,2) NOT NULL
);

CREATE TABLE IF NOT EXISTS Computer_Components(
    Computer_id INT REFERENCES Computers(id),
    Component_id INT REFERENCES Components(id),
    quantity INT DEFAULT 1,
    PRIMARY KEY(Computer_id, Component_id)
);