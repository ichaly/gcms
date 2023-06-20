CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
  -ldflags "-w -s \
  -X 'github.com/ichaly/gcms/core.Version=0.1.0' \
  -X 'github.com/ichaly/gcms/core.GitHash=$(git rev-parse --short HEAD)' \
  -X 'github.com/ichaly/gcms/core.BuildTime=$(date '+%Y-%m-%d %H:%M:%S')'" \
  -o main/dist/app main/main.go

docker build -t gcms:latest .

rm -rf main/dist

docker login -u 15210203617 -p docker123 registry.cn-qingdao.aliyuncs.com
docker tag yugong:latest registry.cn-qingdao.aliyuncs.com/ichaly/gcms:latest
docker push registry.cn-qingdao.aliyuncs.com/ichaly/gcms:latest