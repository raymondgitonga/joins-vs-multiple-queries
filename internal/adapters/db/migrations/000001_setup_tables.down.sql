ALTER TABLE IF EXISTS product_quantity
    DROP CONSTRAINT catalog_id;
ALTER TABLE IF EXISTS product_catalog
    DROP CONSTRAINT category_id
ALTER TABLE IF EXISTS product_catalog_category
    DROP CONSTRAINT category_id;
ALTER TABLE IF EXISTS product_catalog_category
    DROP CONSTRAINT catalog_id;
DROP TABLE IF EXISTS product_category;
DROP TABLE IF EXISTS product_quantity;
DROP TABLE IF EXISTS product_catalog;