docker login --username=lun3322@aliyun.com registry.cn-beijing.aliyuncs.com

docker tag hong-post:0.1 registry.cn-beijing.aliyuncs.com/lun3322/hong-post:0.1
docker tag hong-user-center:0.1 registry.cn-beijing.aliyuncs.com/lun3322/hong-user-center:0.1
docker tag hong-login:0.1 registry.cn-beijing.aliyuncs.com/lun3322/hong-login:0.1

docker push registry.cn-beijing.aliyuncs.com/lun3322/hong-post:0.1
docker push registry.cn-beijing.aliyuncs.com/lun3322/hong-user-center:0.1
docker push registry.cn-beijing.aliyuncs.com/lun3322/hong-login:0.1
