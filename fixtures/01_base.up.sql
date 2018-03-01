CREATE TYPE t_account_status AS ENUM ('ENABLE', 'DISABLE', 'PENDING', 'PASSWORD_CHANGE');

CREATE OR REPLACE FUNCTION getSeqUUID(tableId INT, seqName VARCHAR(100))
  RETURNS UUID AS $$
BEGIN
  RETURN concat(lpad(TO_HEX(tableId), 16, '0'), lpad(TO_HEX(nextval(seqName::regclass)),16, '0'))::UUID;
END;
$$ language 'plpgsql';

