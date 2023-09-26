#!/usr/bin/env bash

set -euo pipefail

CONTAINER_NAME=faceto-ai
IMAGE_NAME=faceto-ai
IMAGE_VERSION=v1.0.0
IMAGE_NAME_VERSION=${IMAGE_NAME}:${IMAGE_VERSION}

# current dir
MOUNT_DIR=$(pwd)/logs

run_main() {
    echo "1) Generate code"
    echo "2) Build and run server"
    echo "3) Build and run cron"
    echo "4) Build image"
    echo "5) Run container"
    read -r -p "Please select an option: " option
    case $option in
        1)
            generate_code
            ;;
        2)
            build_main
            run_server
            ;;
        3)
            build_main
            run_cron
            ;;
        4)
            build_image
            ;;
        5)
            run_docker
            ;;
        *)
            echo "Invalid option"
            exit 1
            ;;
    esac
}

generate_code() {
    go generate ./internal/data/ent --feature sql/upsert
    make all
}

build_main() {
    make build
}

run_server() {
    ./bin/faceto-ai -conf ./configs/config.yaml -env prod
}

run_cron() {
    #./bin/task cron -conf ./configs/config.yaml -env prod
}

build_image() {
    # build image
    docker build -t ${IMAGE_NAME_VERSION} .
}

run_docker() {
    docker stop ${CONTAINER_NAME} || true
    docker rm ${CONTAINER_NAME} || true
    # run container
    docker run --name ${CONTAINER_NAME} -d --env-file .env.local -p 8001:8001 -v ${MOUNT_DIR}:/app/logs/ ${IMAGE_NAME_VERSION}
}

run_main