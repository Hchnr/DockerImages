用go实现的简单的webservice: chart-server

主要用于发布helm所需要的charts 

## 本地部署&调试

1. 执行 `docker build --rm -t chart-server --build-arg SSH_KEY="$(cat ~/.ssh/id_rsa)" .`
2. 执行 `docker tag chart-server registry.cn-beijing.aliyuncs.com/shannonai/chart-server:v1.0.0`
3. 执行 `docker-compose up -d` OR `kubectl apply -f deployment.yaml`
4. container_id=`docker ps | grep chart-server | awk -F " " '{print $1}'`
5. docker logs $container_id
6. docker exec -it $container_id bash

## k8s 集群部署&调试
1. kubectl apply -f deployment.yaml
2. 可以使用k8s dashboard查看部署状态  

集群中需要有ssh秘钥文件对应的secret, ssh-key-shannon下还需要有：
1. id_rsa
2. id_rsa.pub
3. config 用于git clone时免于输入yes  
4. 执行 kubectl delete secret ssh-key-shannon -n monitoring
5. 执行 kubectl create secret generic ssh-key-shannon --from-file=ssh_rsa=./id_rsa --from-file=ssh_rsa.pub=./id_rsa.pub --from-file=config=./config -n monitoring


## chart-server 使用方法
1. GET https://chart.shannonai.com:4443/health 用于健康检测
2. GET https://chart.shannonai.com:4443/release 用于获取chart仓库内容
3. POST https://chart.shannonai.com:4443/update 用于更新chart仓库, POST请求，payload全部置于http.body中，通过Secret对payload加密，并把加密后的的摘要置于http.header["X-Hub-Signature"]下

## charts源码仓库webhook部署注意事项
1. [Content type]选择[application/json], 选择[application/x-www-form-urlencoded]可能导致secret认证失败
2. Webhook所用的Secret可以在./conf中配置
3. github.com上的Wenhook需要禁用[SSL verification]选项, 目前集群的https代理节点的证书不支持低版本ssl


