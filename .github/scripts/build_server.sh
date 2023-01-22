#!/bin/bash
cd ./blog-backend

cd ./app/admin/service
make build

cd ../../../app/comment/service
make build

cd ../../../app/content/service
make build

cd ../../../app/file/service
make build

cd ../../../app/user/service
make build
