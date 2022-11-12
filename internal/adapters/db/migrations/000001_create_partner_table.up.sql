CREATE TABLE IF NOT EXISTS partner
(
    id       int PRIMARY KEY,
    name     varchar,
    location point,
    material jsonb,
    radius   int,
    rating   decimal
);

CREATE SEQUENCE IF NOT EXISTS partner_seq START 1 INCREMENT 1 MINVALUE 1 OWNED BY partner.id;