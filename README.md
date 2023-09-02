# Password Manager

## Project Overview

This project is the beginning of my experience with **DevOps** workflows. The project consists of a three-tier **cloud-native** architecture with my **Gloang RESTful API**, a **Mongo** database, and a website hosted on **Apache**. These three components run as **Docker** containers orchestrated by **Docker-Compose**. 

## Technology

This uses a number of open-source projects and technologies to work properly:

- [Gorilla Mux] - HTTP Router to handle API calls
- [RS CORS] - Middleware to enable Cross_origin Resource Sharing (CORS)
- [XMLHttpRequest] - Web API that makes HTTP and XML calls

- [Docker] - Containerization and orchestration platform
- [Apache] - HTTP Webserver for hosting websites
- [MongoDB] - Non-Relational key-value store  


## Installation

**Password Manager** was built with [Docker](https://www.docker.com/) v24.0.2+ and needs the same to run smoothly so please ensure that you have done this as you follow along.

Verify the `Docker` installation with  with:

```sh
docker version
```

Download the `Password-Manager` `docker-compose.yaml` file from [here](https://github.com/anirudh-devanand/Password-Manager/blob/main/docker-compose.yaml) by pressing `Ctrl + Shift + S`.

Open the `terminal` application on your system and navigate to the folder that containes the downloaded `docker-compose.yaml` file.

 
 Run the following command in the said folder:
  ```sh
  docker-compose up
  ```
Check if the `PwdMngr` website is up and running by typing this into any web browser of your choice:
```sh
http://localhost:8082
```

## Finals Thoughts

This is still a work in process and as I learn and evolve so will this project.
In the meantime, I hope y'all like what I have done here.


.




[Gorilla Mux]:https://github.com/gorilla/mux
[RS CORS]:https://github.com/rs/cors
[XMLHttpRequest]: https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest
[Docker]: https://www.docker.com/
[Apache]: https://httpd.apache.org/
[MongoDB]: https://www.mongodb.com/
