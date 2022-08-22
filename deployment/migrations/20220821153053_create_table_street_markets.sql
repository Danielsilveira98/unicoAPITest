-- +goose Up
-- +goose StatementBegin
create table if not exists street_market (
  id uuid primary key not null,
  long float8,
  lat float8,
  sectcens VARCHAR(50),
  area VARCHAR(50),
  iddist INT,
  district VARCHAR(50),
  idsubth INT,
  subtownhall VARCHAR(50),
  region5 VARCHAR(50),
  region8 VARCHAR(50),
  name VARCHAR(50),
  register VARCHAR(50),
  street VARCHAR(50),
  number VARCHAR(50),
  neighborhood VARCHAR(50),
  addrextrainfo VARCHAR(250),
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop table street_market;

-- +goose StatementEnd
