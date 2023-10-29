#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <directory_name>"
    exit 1
fi

BASE_DIR=$1

# Create base directory
mkdir -p "apps/$BASE_DIR"
cd "apps/$BASE_DIR"
touch .env

mkdir test
touch test/.gitkeep
mkdir config
touch config/.gitkeep
mkdir cmd
touch cmd/main.go

# domain
mkdir -p app/domain/model
mkdir -p app/domain/repository
mkdir -p app/domain/service
touch app/domain/.gitkeep
touch app/domain/model/.gitkeep
touch app/domain/repository/.gitkeep
touch app/domain/service/.gitkeep

# infra
mkdir -p app/infra/dao/query
mkdir -p app/infra/db/ddl
mkdir -p app/infra/db/migrations
mkdir -p app/infra/db/seed
mkdir -p app/infra/dto
touch app/infra/.gitkeep
touch app/infra/dao/.gitkeep
touch app/infra/dao/query/.gitkeep
touch app/infra/db/ddl/.gitkeep
touch app/infra/db/migrations/.gitkeep
touch app/infra/db/seed/.gitkeep
touch app/infra/dto/.gitkeep

# ui
mkdir -p app/ui/grpc
touch app/ui/.gitkeep
touch app/ui/grpc/.gitkeep

# usecase
mkdir -p app/usecase/query
touch app/usecase/.gitkeep
touch app/usecase/query/.gitkeep
