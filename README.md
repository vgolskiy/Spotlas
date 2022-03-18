**Spotlas task for spots search**

1. Created function that get domins from URLs (/scripts/002.sql). Additionally made view with resulting domain names aggregation to find groups with more than 1 spot inside (/scripts/004.sql).

-	To verify result you can use query:

	``` SELECT * FROM domains_aggregated; ```

2. Created endpoint to get spots groups inside square/circle areas around given point with given area radius. Searching areas are defined with postgreSQL functions using built-in postGIS functions (/scripts/003.sql).

Working with repository:
-	To clone repository use:

	``` https://github.com/vgolskiy/Spotlas.git ```

-	To build an application, please, run the following commands sequence inside the repository root directory:

	a.	``` docker-compose up -d postgres ```

	b.	``` docker-compose up -d spotlas ```
	
-	To verify results you can use Postman requests collection: https://www.postman.com/orange-moon-322189/workspace/spotlas/collection/16303804-28f310c9-8b6e-4fc0-b4b7-d75318868e45?action=share&creator=16303804
