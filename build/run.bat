@echo off
REM Batch script to start Docker Compose services

REM Start default Docker Compose services
docker-compose up -d

REM Start Docker Compose services for the database
docker-compose -f ./docker-compose-db.yml up -d

REM Start Docker Compose services for logging (Grafana, Loki, Promtail)
docker-compose -f ./docker-compose-log.yml up -d
