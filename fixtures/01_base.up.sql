CREATE OR REPLACE FUNCTION getSeqUUID(tableId INT, seqName VARCHAR(100))
  RETURNS UUID AS $$
BEGIN
  RETURN concat('00000000', lpad(TO_HEX(tableId), 4, '0'), '00000000', lpad(TO_HEX(nextval(seqName::regclass)),12, '0'))::UUID;
END;
$$ language 'plpgsql';

