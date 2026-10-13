CREATE TABLE IF NOT EXISTS Roles(
    id SERIAL PRIMARY KEY,
    title VARCHAR(10)
);

INSERT INTO Roles(title) VALUES('Owner');
INSERT INTO Roles(title) VALUES('Admin');
INSERT INTO Roles(title) VALUES('Seller');
INSERT INTO Roles(title) VALUES('Buyer');

CREATE TABLE IF NOT EXISTS Users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    balance NUMERIC(10,2) DEFAULT 0,
    Role_id INT REFERENCES Roles(id)
);

CREATE TABLE IF NOT EXISTS Components(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    overview TEXT,
    price NUMERIC(10,2) NOT NULL,
    quantity INT DEFAULT 0,
    User_id INT REFERENCES Users(id)
);

CREATE TABLE IF NOT EXISTS Computers(
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    overview TEXT,
    price NUMERIC(10,2) NOT NULL,
    User_id INT REFERENCES Users(id)
);

CREATE TABLE IF NOT EXISTS Computer_Components(
    id SERIAL PRIMARY KEY,
    Computer_id INT REFERENCES Computers(id),
    Component_id INT REFERENCES Components(id),
    User_id INT REFERENCES Users(id),
    quantity INT DEFAULT 1
);
