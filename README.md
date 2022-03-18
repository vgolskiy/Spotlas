**Spotlas task for spots search**

1. Created function that get domins from URLs (/scripts/002.sql). Additionally made resulting domain names aggregation to find groups with more than 1 spot inside (task/script.sql).
2. Created endpoint to get spots groups inside square/circle areas around given point with given area radius.
-	Searching areas are defined with postgreSQL functions using built-in postGIS functions (/scripts/003.sql). 
-	To verify work you can use Postman requests collection: https://www.postman.com/orange-moon-322189/workspace/spotlas/collection/16303804-28f310c9-8b6e-4fc0-b4b7-d75318868e45?action=share&creator=16303804
