BEGIN;

-- CREATE SEQUENCE "id_sel_seq" -------------------------------
CREATE SEQUENCE "public"."id_role"
INCREMENT 1
MINVALUE 1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
-- -------------------------------------------------------------

-- CREATE TABLE "role" -------------------------------------
CREATE TABLE "public"."role" (
    "id" Bigint DEFAULT nextval('id_role'::regclass) NOT NULL,
    "name_role" Character Varying( 255 ) NOT NULL,
    PRIMARY KEY ( "id" ) );
 ;
 -----------------------------------------------------------

INSERT INTO "public"."role" ("id", "name_role") VALUES
(1, 'Manager');
INSERT INTO "public"."role" ("id", "name_role") VALUES
(2, 'Cashier');
INSERT INTO "public"."role" ("id", "name_role") VALUES
(3, 'Staff');
INSERT INTO "public"."role" ("id", "name_role") VALUES
(4, 'Receptionist');
INSERT INTO "public"."role" ("id", "name_role") VALUES
(5, 'Security');

COMMIT;