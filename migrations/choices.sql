-- -------------------------------------------------------------
-- TablePlus 4.5.0(396)
--
-- https://tableplus.com/
--
-- Database: driving-license-examination
-- Generation Time: 2565-01-22 10:48:05.0080
-- -------------------------------------------------------------


DROP TABLE IF EXISTS "public"."choices";
-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: indices, triggers. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS choices_id_seq;

-- Table Definition
CREATE TABLE "public"."choices" (
    "id" int4 NOT NULL DEFAULT nextval('choices_id_seq'::regclass),
    "image" varchar NOT NULL DEFAULT 'some_path'::character varying,
    "detail" text NOT NULL,
    "is_correct" bool NOT NULL DEFAULT false,
    "question_id" int4 NOT NULL,
    CONSTRAINT "choices_question_id_fkey" FOREIGN KEY ("question_id") REFERENCES "public"."questions"("id"),
    PRIMARY KEY ("id")
);

