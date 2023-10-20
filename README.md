# PizzaDelivery

Pizza delivery application
This is a pizza delivery application implemented using golang. This app has:

-> two microservices (one that operates on users and one on pizzeria)
-> docker that keeps two separate database running
Prerequisites to run the project:

Go
The chosen programming language for this project

https://go.dev/doc/install

Docker
Used to run the database containers.
Each microservice has its own database container.

https://docs.docker.com

golang-migrate
Used for DB schema migrations.
Must install the library:

brew install golang-migrate

