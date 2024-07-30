./docker-compose up -d
./docker-compose -f ./docker-compose-db.yml up -d
./docker-compose -f ./docker-compose-log.yml up -d
./docker-compose -f ./docker-compose-broker.yml up -d
