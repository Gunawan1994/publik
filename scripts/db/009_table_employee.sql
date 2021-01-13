BEGIN;

-- CREATE SEQUENCE "id" -------------------------------
CREATE SEQUENCE "public"."id_employee"
INCREMENT 1
MINVALUE 1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
-- -------------------------------------------------------------

-- CREATE TABLE "employee" -------------------------------------
CREATE TABLE "public"."employee" (
    "id" Bigint DEFAULT nextval('id_employee'::regclass) NOT NULL,
    "full_name" Character Varying( 255 ) NOT NULL,
    "address" Character Varying( 255 ) NOT NULL,
	"dob" Timestamp Without Time Zone NOT NULL,
	"role_id" Bigint NOT NULL,
	"role_name" Text NOT NULL,
	"salary" Bigint NOT NULL,
    PRIMARY KEY ( "id" ) );
 ;
 -----------------------------------------------------------


COMMIT;