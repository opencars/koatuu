CREATE TABLE IF NOT EXISTS level1 (
    "id"   CHAR(2) NOT NULL,
    "name" TEXT    NOT NULL,
    PRIMARY KEY (id)
);

CREATE TYPE LEVEL2_KIND_T AS ENUM ('DISTRICT_CITY', 'DISTRICT', 'SPECIAL_CITY_REGION');

CREATE TABLE IF NOT EXISTS level2 (
    "id"        CHAR(5)       NOT NULL,
    "kind"      LEVEL2_KIND_T NOT NULL,
    "name"      TEXT          NOT NULL,
    "level1_id" CHAR(2)       NOT NULL REFERENCES level1("id"),
    PRIMARY KEY (id)
);

CREATE TYPE LEVEL3_KIND_T AS ENUM ('REGION_CITY', 'DISTRICT_CITY_REGION', 'CITY_URBAN_SETTLEMENT', 'REGION_URBAN_SETTLEMENT', 'CITY_REGION_URBAN_SETTLEMENT', 'CITY', 'REGION_SETTLEMENT', 'CITY_SETTLEMENT');

CREATE TABLE IF NOT EXISTS level3 (
    "id"          CHAR(8)       NOT NULL,
    "kind"        LEVEL3_KIND_T NOT NULL,
    "name"        TEXT          NOT NULL,
    "level1_id"   CHAR(2)       NOT NULL REFERENCES level1("id"),
    "level2_id"   CHAR(5)       NOT NULL REFERENCES level2("id"),
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS level4 (
    "id"          CHAR(10) NOT NULL,
    "name"        TEXT     NOT NULL,
    "level1_id"   CHAR(2)  NOT NULL REFERENCES level1("id"),
    "level2_id"   CHAR(5)  NOT NULL REFERENCES level2("id"),
    "level3_id"   CHAR(8)  NOT NULL REFERENCES level3("id"),
    PRIMARY KEY (id)
);