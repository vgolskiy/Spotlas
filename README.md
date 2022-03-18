**Spotlas task for spots search**

1. Created function that get domins from URLs (/scripts/002.sql). Additionally made view with resulting domain names aggregation to find groups with more than 1 spot inside (/scripts/004.sql).

-	To verify result you can use query:

	``` SELECT * FROM domains_aggregated; ```

2. Created endpoint to get spots groups inside square/circle areas around given point with given area radius. Searching areas are defined with postgreSQL functions using built-in postGIS functions (/scripts/003.sql).

Working with repository:
-	To clone repository use:

	``` git clone https://github.com/vgolskiy/Spotlas.git Spotlas```

-	To build an application, please, run the following commands sequence inside the repository root directory:

	a.	``` docker-compose up -d postgres ```

	b.	``` docker-compose up -d spotlas ```
	
-	To verify results you can use Postman requests collection (Spotlas.postman_collection.json)