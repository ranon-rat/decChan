rm -rf db/database.db
touch db/database.db
cat ./db/init.sql | sqlite3 ./db/database.db