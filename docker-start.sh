#sudo docker build -t realworld-api .

#sudo docker run -d -p 8080:8080 realworld-api

#!/bin/bash

# 실행 중인 컨테이너의 이름 또는 ID 설정
CONTAINER_NAME="realworld-api"

# 도커 캐시 삭제
echo "Cleaning Docker Cache..."
sudo docker system prune -a -f

# 실행 중인 동일한 이름의 컨테이너가 있는지 확인
if [ $(sudo docker ps -q -f ancestor=${CONTAINER_NAME}) ]; then
	echo "qwejioqwjeoiqwejoiqwej"
	# 컨테이너가 실행 중이면 중지하고 삭제
	echo "Stopping and removing existing container..."
	sudo docker stop ${CONTAINER_NAME}
	sudo docker rm ${CONTAINER_NAME}
fi

echo ${CONTAINER_NAME}a

# Docker 이미지 빌드
echo "Building Docker image..."
sudo docker build -t realworld-api .

# 새로운 컨테이너 실행
echo "Running new container..."
sudo docker run --network="host" -d -p 8080:8080 --name ${CONTAINER_NAME} realworld-api

echo "Container ${CONTAINER_NAME} is running."
