# goapp
Simple web app which takes the scrap text entered by the user and saves it in mongo along with the timestamp. The scraps along with the timestamp is displayed to the user.The whole project is a demonstration of creating a dockerised webapp which has mongo, go code which is the controller and nginx which acts as the webserver to serve html.

There are 3 main components for this application, 

First part is to bring up a mongodb as a container, so to do that we define a service inside the docker-compose.yaml file named mongo. This service defines an image that needs to be used, along with the container name and port exposed by mongo.

Second is the main.go file inside api folder which has the logic to accept GET and POST http APIs on the endpoint /scraps. This is added by the gorilla mux rounter which is used to configure routes (http end point path for API). There is also a mongo driver used to connect to the mongo database which is used to save the scrap in the Scraps collection. There are two main methods, 
1. readScrap which reads all scraps present in the database and returns the result.
2. saveScrap which saves the scrap entered by the user along with the timestamp into the mongo collection.
To add this as a container, we define another service in docker-compose.yaml named api where we specifcy the port which the application listens, dependency which is the mongo container. The api folder contains a docker file which has steps to bring up the api container itself.

Third part is a webserver which can server html files, Nginx is used for this purpose. An index.html file is created which has the logic to display a text box and invoke the apis exposed by the api container to save and display scraps. This is created under web folder which also has a docker file which specify how to load the nginx as a container and uses the index.html to serve requests. A third service is added to docker-compose.yaml file named web which specifies the nginx image, port and volumes so as to load the index.html. It depends on the api container which is added as a dependency.

To run this application, first install docker (or docker-compose) and then check out the code, navigate to the src directory and run the command

docker-compose up --build

This will download all the required images and start the containers. Once it is started, the application can be accessed by the url http://localhost:8081 (or http://hostname:8081).

This is done by Jagadeesh P N (jagadeesh.2@iitj.ac.in, Roll No: MT19AIE246).


