package core

const usersDDL = `
CREATE TABLE IF NOT EXISTS users
(
    id      INTEGER PRIMARY KEY AUTOINCREMENT,
    name    TEXT    NOT NULL,
    login   TEXT    NOT NULL UNIQUE,
    password TEXT NOT NULL,
    telNumber int NOT NULL
);`

const accountsDDL = `
CREATE TABLE IF NOT EXISTS accounts
(
    id    INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER NOT NULL UNIQUE,
    score INTEGER
);`

const usersInitialData = `INSERT INTO users
VALUES (1, 'Vasya', 'vasya', 'secret', 1),
       (2, 'Petya', 'petya', 'secret', 2),
       (3, 'Vanya', 'vanya', 'secret', 3),
       (4, 'Masha', 'masha', 'secret', 4),
       (5, 'Dasha', 'dasha', 'secret', 5),
       (6, 'Sasha', 'sasha', 'secret', 6)
       ON CONFLICT DO NOTHING;`

const accountsInitialData = `INSERT INTO accounts(id, user_id, score)
VALUES 	(1, 1, 200),
		(2, 1, 100),
		(3, 2, 1100),
		(4, 3, 1010),
		(5, 4, 100),
		(6, 5, 0),
		(7, 6, 1
       ON CONFLICT DO NOTHING;`

const loginSQL = `SELECT login, password FROM users WHERE login = ?`
const getAllAccountsByIdUser = `SELECT score FROM accounts WHERE id = ?;`
//const insertSaleSQL = `INSERT INTO sales(manager_id, product_id, price, qty) VALUES (:manager_id, :product_id, :price, :qty);`
