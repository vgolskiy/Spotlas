CREATE OR REPLACE FUNCTION getSpotsInCircle(longitude float8, latitude float8, radius float8)
RETURNS SETOF "MY_TABLE"
AS $$
BEGIN
    DROP TABLE IF EXISTS spots;
    CREATE TEMP TABLE spots AS
    SELECT id, name, website, coordinates, description, rating,
           st_distance(st_makepoint(longitude, latitude)::geography,
               coordinates::geography) dist
    FROM "MY_TABLE"
    WHERE coordinates::geography && st_buffer(
        st_makepoint(longitude, latitude)::geography, radius, 'quad_segs=8')
    ORDER BY dist ASC;
    RETURN QUERY ((SELECT id, name, website, coordinates, description, rating
    FROM spots
    WHERE dist < 50
    ORDER BY rating DESC)
    UNION ALL
    SELECT id, name, website, coordinates, description, rating
    FROM spots
    WHERE dist >= 50);
END
$$ LANGUAGE  'plpgsql';


CREATE OR REPLACE FUNCTION getSpotsInSquare(longitude float8, latitude float8, minLng float8, minLat float8, maxLng float8, maxLat float8)
RETURNS SETOF "MY_TABLE"
AS $$
BEGIN
    DROP TABLE IF EXISTS spots;
    CREATE TEMP TABLE spots AS
    SELECT id, name, website, coordinates, description, rating,
           st_distance(st_makepoint(longitude, latitude)::geography,
               coordinates::geography) dist
    FROM "MY_TABLE"
    WHERE coordinates::geometry && ST_Envelope(
        ST_GeomFromText(CONCAT('LINESTRING(', minLng, ' ', minLat, ',', maxLng, ' ', maxLat, ')'), 4326))
    ORDER BY dist ASC;
    RETURN QUERY ((SELECT id, name, website, coordinates, description, rating
    FROM spots
    WHERE dist < 50
    ORDER BY rating DESC)
    UNION ALL
    SELECT id, name, website, coordinates, description, rating
    FROM spots
    WHERE dist >= 50);
END
$$ LANGUAGE  'plpgsql';
