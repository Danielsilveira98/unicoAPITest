-- +goose Up
-- +goose StatementBegin
create table if not exists street_market (
  id uuid primary key not null,
  long float8 NOT NULL,
  lat float8 NOT NULL,
  sectcens VARCHAR(50) NOT NULL,
  area VARCHAR(50) NOT NULL,
  iddist VARCHAR(50) NOT NULL,
  district VARCHAR(50) NOT NULL,
  idsubth VARCHAR(50) NOT NULL,
  subtownhall VARCHAR(50) NOT NULL,
  region5 VARCHAR(50) NOT NULL,
  region8 VARCHAR(50) NOT NULL,
  name VARCHAR(50) NOT NULL,
  register VARCHAR(50) NOT NULL,
  street VARCHAR(50) NOT NULL,
  number VARCHAR(50) NOT NULL,
  neighborhood VARCHAR(50) NOT NULL,
  addrextrainfo VARCHAR(250) NOT NULL,
  createdat TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop table street_market;

-- +goose StatementEnd
