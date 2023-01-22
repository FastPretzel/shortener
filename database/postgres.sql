DROP TABLE IF EXISTS links;
CREATE TABLE links (
    short_link VARCHAR(10) PRIMARY KEY,
    orig_link TEXT UNIQUE
);
