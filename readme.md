# Backend for face recognition project

## Goals of the backend

We would like to introduce the following features in the project :

* Face authorization by Rest API (in Golang)
* RethinkDb usage for the NOSQL database
* Kafka as the multi-broken as a publish-subscripe pattern

Clients should be able to run on any platform, being a mobile ap, a website,
or an embedded system.

## RethinkDb

RethinkDb is a NoSQL database with an open-source base code that is maintained by a large community of contributers. RethinkDb advantage compared to more famous NoSQL databases is the support of real-time queries implemented by changefeeds,
which are based on a publish-subscribe pattern allowing instant notifications of any table change. This is particularly useful for uses in push notifications or very latency limited applications.
RethinkDb offers a GeoJSON support for geo-queries and geolocation oriented data manipulation, as well as scientific measurement handling.
RethinkDb lack of popularity comes from the fact that the project has been abandoned multiple times before being restarted, but in no way this limit the power of RethinkDb as the community is constantly maintaining diverse repository and implementations
for each language.

Our database plan is the following :

    -------------------
    | squeezedb |
    :------------------:
    | employees |
    -------------------
    |    devices    |
    -------------------


## REST API

### GoLang

Go is a recent language develloped by the one of the father of the C language. It is a compiled, garbage collected language whose syntax is very close to the C syntax, while maintaining a simplistic approach. Many principles have been avoided to provide simplicity of coding. 
Classes are replaced by analog to C, structures. Interfaces provide an abstraction of the object oriented interfaces while introducing differences with them.
GoLang has in mind multi-threading, by the use of goroutines and channel based communication. Goroutines does not involve parallelism but asynchronous execution. 
Those concepts makes it very easy to implement fast Go softwares.

Go is particularly known for the ease that comes with it when develloping web applications.
Some lines of code are enough to build a working web server and process requests.
Further discussion on go concepts will be available in this documentation.

#### The net/http library

The net/http library is the core http library for the go language. Each submodule is based on it and extends it's functionnalities. The gorilla library is a great example.
Using the net/http library has been made very simple by the go developers, who have chosen to focus on modularity and simplicity rather than complexity. This does not mean it displays less features than
any other native http library from other languages.

#### Business logic of the REST API :

This project contains seven folders, which will be described in the following :
* /common : This folder allows a global setup of the program parameters by reading the provided json files
containing the port to be used, and diverse paths such as the main c++ program path and the database configuration.
It also allows to setup RethinkDb by providing a singleton session to be used during runtime.
It includes utils functions to be used by others submodules.

* /controllers :  Since we are using the MVC approach (Model-View-Controller), we must provide the controllers of the program.
The controllers allow to implement the behavior of the program.
We provide three controllers : the device controller, allowing to register and authenticate a device by providing a Json Web Token (JWT).
The employee controller allows various actions over the employees, such as authorizing, registering, deleting, updating. It includes an endpoint allowing
to save the images of an employee by encoding and decoding images.

* /data :  The data subfolder is the core of our program. All the business logic is implemented there. Each endpoint behavior gets it's result from the functions contained
in the folder. We provide this folder with some essentials tests in order to allow status-checking.
Tests are executed using a mock database provided by the goRethink package.

* /keys : This folder contains the keys used for the JWT generation.

* /models : The models folder simply provide the json models to provide to a REST command. The models are to be discussed in later subtopics.

* /routers : The routers folder allows routing of the different endpoint to their programmed behavior implemented in the controller folder.
It is important to precise that all the different endpoints are using negroni in order to implement the JWT authorization.

#### Endpoints description

* /device endpoints :

| Endpoint      | Method | Authorization | Comments                                                                                                                                                                                               |
|---------------|--------|---------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| /register     |  POST  |       No      | Allows to register a device. The json data provided should look like the following :   {    "data" : {              "deviceName":"name",               "deviceMac":"44:55:55:55:55:55"             } } |
| /authenticate |  POST  |       No      | Allows to authenticate a device. Consumes the same JSON format than previously.                                                                                                                        |

* /employees endpoints : (all require a JWT passed as a bearer authorization header)

|  Endpoint | Method | Authorization | Comments                                                                                                                                                                                                                                                                                                                                                                                                                                                   |
|:---------:|:------:|---------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| /register |  POST  |      YES      |  Registering an employee. Data should contains a "data" json object providing the firstName, lastName, email, authStarting and authEnding as strings seconds timestamp.                                                                                                                                                                                                                                                                                    |
| /update   | UPDATE |      YES      | Updating an employee. Email should be the same as it is used to retrieve the employee data.  Any other information should follow the previous json format.                                                                                                                                                                                                                                                                                                 |
| /face     |  POST  |      YES      |  Retrieving authorization status. Code 200 if the employee is authorized, 403 if the JWT is not valid or employee is not authorized. Json format should be as previously described, except the firstName and lastName are  mandatory, while other fields are not.                                                                                                                                                                                          |
| /list     |   GET  |      YES      | Retrieving list of employees in the database as a JSON array.                                                                                                                                                                                                                                                                                                                                                                                              |
| /image    |  POST  |      YES      |  Saving an image to the training-images folder of the main c++ program. The JSON format should contains uniquely   the following fields :  timestamp, name, filename and data. Name should be of the following format : firstName-LastName and match an employee in the database. Filename is the name under which the file will be saved. Return HTTP code 200 if everything went fine, 500 otherwise as the error would be from unmarshalling or saving. |
| /remove   |  POST  |      YES      | Remove an employee from the database. This endpoint consumes the same JSON format as the previous endpoints, except the email is the only mandatory field.                                                                                                                                                                                                                                                                                                 |
| /logs     |   GET  |      YES      | Get the log file monitoring the allowed and refused access by a camera. Produces a json array of string.                                                                                                                                                                                                                                                                                                                                                   |
| /logs     | DELETE |      YES      | Empty and delete the log file containing the monitoring data. Return the 200 HTTP code if success. Put in mind that the code behind executes a bash script so more prone to errors.                                                                                                                                                                                                                                                                        |


* /network endpoints :

| Endpoint  | Methods | Authorization |  Comments                                                                                                                                                                                                                                             |
|-----------|---------|---------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| /train    |   POST  |      YES      | Send the training signal to the classifier. This endpoint executes a bash script. If any errors, the bash script is to be checked. It consumes a JSON string composed of a "data" value containing a json object with the "train" boolean attribute.  |
| /pictures |  DELETE |      YES      |                                                                                   Delete the pictures inside the training folder. This endpoint removes ALL images.                                                                                   |

* Example of a request to /employees/register :
POST - { "data"  : {
                "firstName" : "Naif",
                "lastName" : "Mehanna",
                "email" : "email1@example.com",
                "authStarting" : "1154898621",
                "authEnding" : "1155748964"
                }
         }