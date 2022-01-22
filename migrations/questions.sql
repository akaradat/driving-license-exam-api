-- -------------------------------------------------------------
-- TablePlus 4.5.0(396)
--
-- https://tableplus.com/
--
-- Database: driving-license-examination
-- Generation Time: 2565-01-22 10:05:57.8930
-- -------------------------------------------------------------


DROP TABLE IF EXISTS "public"."questions";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS questions_id_seq;

-- Table Definition
CREATE TABLE "public"."questions" (
    "id" int4 NOT NULL DEFAULT nextval('questions_id_seq'::regclass),
    "image" varchar NOT NULL DEFAULT 'some_path'::character varying,
    "detail" text NOT NULL,
    PRIMARY KEY ("id")
);

