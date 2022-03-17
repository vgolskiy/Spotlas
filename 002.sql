CREATE OR REPLACE FUNCTION get_domain(uri text)
RETURNS text AS
$$
DECLARE tmp text;
BEGIN
    SELECT split_part(split_part(regexp_replace(uri,'^http(.)?://(www(.?)\.)?', ''), '/', 1), '?', 1)
    INTO tmp;

    IF tmp LIKE '%.%.%.%' THEN
        RETURN regexp_replace(tmp, '[^.]*\.*\.*', '');
    END IF;

    RETURN tmp;
END;
$$ LANGUAGE 'plpgsql';