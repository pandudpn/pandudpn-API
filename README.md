# API

Project for handle my website [www.pandudpn.id](https://www.pandudpn.id).

## Installation

1. Clone Project

Clone this project into your local

```bash
$ git clone git@github.com:pandudpn/pandudpn-api.git && cd pandudpn-api
```

2. Install Dependencies

After you in root project, and you will see file `Makefile`. if you see, please run this script

```bash
$ make tools
```

If you already success install the dependencies, then you need to input the value of environment file `.env` with your own value.

3. Migrate database

Run this script

```bash
$ make migrate-up
```

## Architecture

This project using Architecture *__The Clean Architecture__* from __Uncle Bob__. From the deep first we have

1. Model / Entity

In this layer, we will declare all data structure from database, request or response. And datastore (repository).

2. Usecase (business logic)

All of specific business logic will be here. These use cases orchestrate the flow of data to and from entities, and direct those entities to use their datastore or datastructure business rules to achieve the goals of the use case.

3. Controller

Controller will convert the data from the form most convenient for entities and use cases.

4. Frameworks and Driver (Adapter)
