CREATE TABLE "clients" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "account_id" int,
  "contact_email" varchar UNIQUE NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "api_keys" (
  "id" bigserial PRIMARY KEY,
  "client_id" bigint,
  "api_key" varchar UNIQUE NOT NULL,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "addresses" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "address" varchar NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "shipments" (
  "id" bigserial PRIMARY KEY,
  "client_id" bigint,
  "from_address_id" bigint,
  "to_address_id" bigint,
  "shipper_id" bigint,
  "shipment_code" varchar UNIQUE,
  "fee" int NOT NULL,
  "status" varchar CHECK (status IN ('created','picked','in_transit','delivered','failed')),
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "shipment_status_logs" (
  "id" bigserial PRIMARY KEY,
  "shipment_id" bigint,
  "status" varchar NOT NULL,
  "note" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "shippers" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "phone" varchar NOT NULL,
  "active" boolean DEFAULT true,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "accounts" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "role" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

COMMENT ON COLUMN "shipments"."status" IS 'pending, accepted, in_transit, delivered, canceled';

ALTER TABLE "clients" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");

ALTER TABLE "api_keys" ADD FOREIGN KEY ("client_id") REFERENCES "clients" ("id");

ALTER TABLE "shipments" ADD FOREIGN KEY ("client_id") REFERENCES "clients" ("id");

ALTER TABLE "shipments" ADD FOREIGN KEY ("from_address_id") REFERENCES "addresses" ("id");

ALTER TABLE "shipments" ADD FOREIGN KEY ("to_address_id") REFERENCES "addresses" ("id");

ALTER TABLE "shipments" ADD FOREIGN KEY ("shipper_id") REFERENCES "shippers" ("id");

ALTER TABLE "shipment_status_logs" ADD FOREIGN KEY ("shipment_id") REFERENCES "shipments" ("id");
