# BoilerPlate test app

This is a POC of a Go boilerPlate app where only need to create new struct with your logic and api endpoint.

Remove the database logic that you dont going to use.

## How to use

## Logging

The log in the terminal allways will work, you can add more logs passing a channel that accept the LogEntry model. Its can be select logging in terminal only, no matter if have channel or not, using:

```
log.Get().SetLocal(true)
```

The logging have 2 level info and error, can be call from any place using

```
log.Get().Print(log.ErrorLevel, message string)
```

## Creating new models

Copy one of the model, sql or nosql, and modify the struct to much your requirement.
In the defined method you can write the logic that you want.

To use this new model, you must create a new instance an pass it to the controller, and if you use sql database to the database instance creation.

```
    //Config database SQL
	stuff := new(stuff.Stuff)
    nModel: new(your.NewModel)
	db := sql.CreatePostgres(stuff,nModel)
	defer func() {
		db.Close()
	}()

	//SQL controller
	go controller.New("3003", db, stuff,nModel)
```

- In SQL CreatePostgres function will automigrate your model and create the new table for it.
- In NoSQL(mongo) CreateMongo will use the name of the model, using reflection, as a collection name.

## Getting started

For test the need a mongo and a postgres database plus nats as a broker, eveerything is posible to deploy using docker:

Go to build folder and run:

```
docker compose up -d
```

This will deploy the mongo and postgres database plus nats

## Create swagger documentation

You need to install swag and export the variable with this commands:

```sh
go install github.com/swaggo/swag/cmd/swag@latest
export PATH=$(go env GOPATH)/bin:$PATH
```

Run this to generate the documentation:

```sh
swag init -g ./cmd/main.go -o ./docs
```

## How to test the app

To do the test of the project type

```sh
go test ./... -cover
```

## Folder structure

### Build

In build you can find the doccker compose, the docker build for the app and the example env file.

In the subfolder there is the scripts and configuration:

- mongo-initdb.d : script in javascript for iniziallizate the mongo database, in this case add a few documents to the other collection.
- nats-conf : configuration file for nats broker.
- postgres-initdb.d : SQL script for the postgres database in this case create a new user, but you can create new databases, tables, whatever you want.

### cmd

Folder with the main file of the project.

- Inside client folder there is a small nats client to send data to the main project

### docs

Documentation generate with swagger

### src

#### api

Here is the controller, the core of the application, start the api and

#### config

Contain the configuration for the project, is a singleton.

#### handler

##### broker

Contain package for the broker in this case nats.

##### database

Packages with the logic for connect to database, in this case sql and nosql.

##### log

Singleton that manage the log of the project.

##### validator

Validator for the model

#### model

Contain the base model of the project.

- base_nosql : Base model for struct that use a nosql database.
- base_sql : Base model for struct that use a sql database.
- broker_msg : Model for the async broker communication.
- const : Constant for the project, for now errors.
- controller : Interface for the controller (api).
- model : Interface for the model.
- request: Structs for the request and response of the project.

Both models other and stuff have the same structure:

- model_api : Contain the definition of the endpoints
- model_create : Swagger documentation and call to controller create.
- model_delete : Swagger documentation and call to controller deleete.
- model_find : Swagger documentation and call to controller find.
- model_list : Swagger documentation and call to controller list.
- model_update : Swagger documentation and call to controller update.
- model : Definition of the model and method of the interface

##### other

Example model for nosql database

##### stuff

Example model for sql database

## Improvements

Add async logic for event driven architecture (remove gin from the controller?)
