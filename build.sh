#GOOS=darwin
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
  -ldflags "-w -s \
  -X 'github.com/ichaly/gcms/core.Version=v0.1.0' \
  -X 'github.com/ichaly/gcms/core.GitHash=$(git rev-parse --short HEAD)' \
  -X 'github.com/ichaly/gcms/core.BuildTime=$(date '+%Y-%m-%d %H:%M:%S')'" \
  -o main/dist/app main/main.go

#删除所有旧镜像
#docker rmi -f $(docker images | grep "yugong" | awk '{print $3}')

#登录到阿里云镜像中心
docker login -u 15210203617 -p docker123 registry.cn-qingdao.aliyuncs.com

#编译镜像
docker buildx build --platform linux/amd64,linux/arm64 -t registry.cn-qingdao.aliyuncs.com/ichaly/gcms:latest \
  -t registry.cn-qingdao.aliyuncs.com/ichaly/gcms:v0.1.0-$(git rev-parse --short HEAD) . --push

#删除Golang编译目录
rm -rf main/dist