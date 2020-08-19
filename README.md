# Service_Order_System
It is a service order system implemented using Golang
This program allows <b>users</b> to create <b>service Orders</b> to a company who give mantain and support of television by <b>technicians</b> asistense <br>
# Enviroment
This project was completly developed in windows 10

# Dependecies
go 1.5 <br>
mysql 3.4 <br>
gqlgen v0.11.3 <br>

# Docker image
You can build the image by:
`docker build -t golang-ubuntu:1151604 .`
and start it by:
`docker run -p 8000:8000 golang-ubuntu:1151604 `

# Executing
you can run the server by: <br>
`go run server.go` <br>

<b>Be sure the database is mysql and is running on the port 3306</b>

Also, you can build an image of the server by: <br>
`go build server.go`

# Usage
Once executed you can request from the `localhost:8000/graphql`<br>

<b>you must request by the graphql standars, if you use postman, you can import [this](https://github.com/cybernuki/Service_Order_System/blob/master/Service%20Order%20System.postman_collection.json) file to see some examples of petitions </b><br>
Also, you can look at the [schema](https://github.com/cybernuki/Service_Order_System/blob/master/schema.graphql) file in order to have some documentation of the entities

## Entities
You can register <b>Users</b> using the <b>createUser</b> mutation<br>
You can register <b>Technicians</b> using the <b>createTechnician</b> mutation<br>
You can register <b>Orders</b> using the <b>createOrders</b> mutation<br>
This last one requires that you are a registed <b>User</b> and be logged. <br>
In order to log as <b>User</b> you must use the <b>login</b> mutation, it will retreive a token. Once got it, require the <b>createOrders</b> adding the "Authorization" header
