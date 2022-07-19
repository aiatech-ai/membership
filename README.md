# membership-go

In this project, we use following tech stacks:
----
1. Migration database: in order to create/alter db, column by command line, avoid changing database manualy.
2. Docker: in order to install database  easier.
3. Jaeger to tracing
4. Kibana to collect log

Install
------
Make sure install Docker & make in your develop enviroment before do clone the repo:
1. Install Docker Desktop: [Download](https://www.docker.com/products/docker-desktop/)
2. Install make tool: `sudo apt install make`
3. Install go-migrate: [Detail](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)   
- MAC: `brew install golang-migrate`
- UBUNTU: 
```
curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash

sudo apt-get update

sudo apt-get install migrate
 ```
4. Clone the repo:    
`git clone git@github.com:shinystarvn/shinymembership.git`
5. Go to the folder:   
`cd shinymembership`
6. Checkout to new branch:   
`git checkout -b mytest`
7. Start docker images:   
`make init`
8. Migrate database:   
`make up`
9. Init all library:   
`go mod tidy`

