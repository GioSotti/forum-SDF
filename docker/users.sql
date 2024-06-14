CREATE TABLE forum_users (
    id SERIAL PRIMARY KEY,
    pseudo VARCHAR(50) NOT NULL,
    mail VARCHAR(100) NOT NULL,
    mot_de_passe VARCHAR(100) NOT NULL
);