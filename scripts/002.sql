CREATE OR REPLACE FUNCTION get_domain(url text)
RETURNS text AS
$$
DECLARE tmp text;
BEGIN
    SELECT split_part(split_part(regexp_replace(url,'^http(.)?://(www(.?)\.)?', ''), '/', 1), '?', 1)
    INTO tmp;

    IF tmp LIKE '%.%.%' AND tmp NOT SIMILAR TO '%(.co.|.com.|.uk.|.gov.|.go.|.org.|.ac.|.in.|.bel.|.ca.|.net.)%' THEN
        RETURN regexp_replace(tmp, '[^.]*\.*\.*', '');
    END IF;

    RETURN tmp;
END;
$$ LANGUAGE 'plpgsql';