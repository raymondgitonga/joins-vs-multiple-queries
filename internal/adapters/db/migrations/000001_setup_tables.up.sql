CREATE TABLE IF NOT EXISTS product_catalog
(
    id          INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    name        varchar,
    price       int,
    category_id int
);

CREATE TABLE IF NOT EXISTS product_category
(
    id   INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    name varchar
);

CREATE TABLE IF NOT EXISTS product_quantity
(
    id         INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    quantity   int,
    catalog_id int
);

CREATE TABLE IF NOT EXISTS product_catalog_category
(
   id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
   catalog_id int,
   category_id int
);

ALTER TABLE product_catalog
    ADD FOREIGN KEY (category_id) REFERENCES product_category (id);

ALTER TABLE product_quantity
    ADD FOREIGN KEY (catalog_id) REFERENCES product_catalog (id);

ALTER TABLE product_catalog_category
    ADD FOREIGN KEY (catalog_id) REFERENCES product_catalog (id);

ALTER TABLE product_catalog_category
    ADD FOREIGN KEY (category_id) REFERENCES product_category (id);