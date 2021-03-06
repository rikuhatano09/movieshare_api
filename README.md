# movieshare_api

Repository for backend of project "MovieShare".

## MovieShare API

An awesome API for project "MovieShare".

### Requirements

See modules list in `go.mod`.

### Setup

1. Start application

    ```sh
    docker-compose build
    docker-compose up --detach
    ```

    Server starts on localhost.
    See "http://127.0.0.1:8000/".


## MovieShare Database

### Structure

- movieshare_database/data

  Raw data to preserve in a container.

- movieshare_database/definition

  Initialization DDLs.  
  They should be always synced with production database schema.

## Usage

- Initialize (Reinitialize)

    ```sh
    rm -rfi movieshare_database/data/*
    docker-compose up --detach --build
    ```

- Start

    ```sh
    docker-compose up --detach
    ```

- Stop

    ```sh
    docker-compose down
    ```

- Destroy

    ```sh
    rm -rfi movieshare_database/data/*
    docker-compose down --rmi all --volumes
    ```

