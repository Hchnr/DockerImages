用go实现的简单的webservice: simplehttpserver

主要用于调试, 此项目可以同时启动 consul, traefik, registrator, simplehttpserver

执行`docker-compose up -d` 之前要先执行 `docker build --rm -t simplehttpserver .`

还可以执行`docker-compose up -d --scale simplehttpserver=3` 对simplehttpserver进行扩展
