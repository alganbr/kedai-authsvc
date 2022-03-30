CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS access_token
(
    id uuid NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_id bigint NOT NULL,
    expires bigint NOT NULL
);