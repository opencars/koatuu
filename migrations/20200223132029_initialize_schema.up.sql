CREATE TABLE IF NOT EXISTS level1 (
    "id"   CHAR(2) NOT NULL,
    "name" VARCHAR(255),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS level2 (
    "id"        CHAR(4) NOT NULL,
    "kind"      INTEGER NOT NULL,
    "level1_id" CHAR(2) NOT NULL REFERENCES level1("id"),
    "name"      VARCHAR(255),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS level3 (
    "id"          CHAR(10) NOT NULL,
    "kind"        INTEGER  NOT NULL,
    "level1_id"   CHAR(2)  NOT NULL REFERENCES level1("id"),
    "level2_id"   CHAR(4)  NOT NULL REFERENCES level2("id"),
    "name"        VARCHAR(255),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS level4 (
    "id"          CHAR(10) NOT NULL,
    "level1_id"   CHAR(2)  NOT NULL REFERENCES level1("id"),
    "level2_id"   CHAR(4)  NOT NULL REFERENCES level2("id"),
    "level3_id"   CHAR(4)  NOT NULL REFERENCES level3("id"),
    "name"        VARCHAR(255),
    PRIMARY KEY (id)
);