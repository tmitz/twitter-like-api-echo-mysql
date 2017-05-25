docker run -d --rm --name test-mysql -p 13306:3306 -v $(pwd)/docker/mysql/conf.d:/etc/mysql/conf.d -v $(pwd)/docker/mysql/sql:/docker-entrypoint-initdb.d/ -e MYSQL_ALLOW_EMPTY_PASSWORD=yes mysql:5.7
