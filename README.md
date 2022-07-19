# membership-go

In this project, we use following tech stacks:
----
1. Migration database: in order to create/alter db, column by command line, avoid changing database manualy.
2. Docker: in order to install database  easier.
3. Jaeger to tracing
4. Kibana to collect log

Install
------
1. Install Docker Desktop: [Download](https://www.docker.com/products/docker-desktop/)
2. Clone the repo:    
`git clone git@github.com:shinystarvn/shinymembership.git`
3. Go to the folder:   
`cd shinymembership`
4. Checkout to new branch:   
`git checkout -b mytest`
5. Start docker images:   
`make init`
6. Migrate database:   
`make migrateup`
7. Init all library:   
`go mod tidy`

