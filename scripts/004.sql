CREATE OR REPLACE VIEW domains_aggregated AS
    SELECT id, name, website, coordinates, description, rating, qty
    FROM (
             SELECT id, name, website, coordinates, description, rating,
                    count(*) OVER (PARTITION BY get_domain(website)) qty
             FROM "MY_TABLE"
             WHERE website IS NOT NULL
               AND length(website) > 0
         ) a
    WHERE qty > 1;