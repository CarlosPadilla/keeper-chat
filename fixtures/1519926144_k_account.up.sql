CREATE TYPE t_account_status AS ENUM ('ENABLED', 'DISABLED', 'PENDING', 'PASSWORD_CHANGE');

CREATE TABLE k_account (
  id     UUID NOT NULL PRIMARY KEY DEFAULT getSeqUUID(1, 'k_account_id_seq'),
  name   VARCHAR(100) NOT NULL,
  email  VARCHAR(100) NOT NULL,
  status t_account_status NOT NULL DEFAULT 'PENDING'
);
CREATE SEQUENCE k_account_id_seq OWNED BY k_account.id;