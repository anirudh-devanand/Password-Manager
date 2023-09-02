# Password Manager
---
## Project Overview
---
This project is the beginning of my experience with **DevOps** workflows. The project consists of a three-tier **cloud-native** architecture with my **Gloang RESTful API**, a **Mongo** database, and a website hosted on **Apache** at the cores. These three components runs as **Docker** containers orchestrated by **Docker-Compose**. 

## Technology
---
This uses a number of open-source projects and technologies to work properly:

- [Gorrlia Mux] - HTTP Router to handle API calls
- [RS CORS] - Middleware to enable Cross_origin Resource Sharing (CORS)
- [XMLHttpRequest] - Web API that makes HTTP and XML calls

- [Docker] - Containerization and orchestration platform
- [Apache] - HTTP Webserver for hosting websites
- [MongoDB] - Non-Relational key-value store  
   
## Installation
---
**Password Manager** was built with [Docker](https://www.docker.com/) v24.0.2+ and needs the same to run smoothly so please ensure that you have done this as you follow along.

Verify the `Docker` installation with  with:

```sh
docker version
```

Download the `Password-Manager` `docker-compose.yaml` file from [here](https://github.com/anirudh-devanand/Password-Manager/blob/main/docker-compose.yaml) by pressing `Ctrl + Shift + S`.

Open the `docker-compose.yaml`  file with a text editor of your choice and change the following line:
```sh
HOST: 192.168.2.141 #Replace '192.168.2.141' IP with your Private IPv4 Address
```
You can find your Private IPv4 on **Windows** either [manually](https://support.microsoft.com/en-us/windows/find-your-ip-address-in-windows-f21a9bbc-c582-55cd-35e0-73431160a1b9) or by running the command:
```sh
ipconfig
```
In the output, under `Wireless LAN Adapter`, and under the `IPv4 Address` section find a number that looks like **192.168.x.x** or **10.x.x.x**. 

On **Unix-like** systems run the command:    
 ```sh
$ ifconfig
 ```
 In the output, under `wlan0`, and next to the `inet` section, find a number that looks like **192.168.x.x** or **10.x.x.x**. 
 
 Once you have your `Private IPv4`, replace the aforementioned line and save the `docker-compose.yaml` file.
 
 Run the following command in the folder that contains `docker-compose.yaml`:
  ```sh
  docker-compose up
  ```
Check if the `PwdMngr` website is up and running by typing this into any web browser of your choice:
```sh
http://localhost:8082
```

## Finals Thoughts
---
This is still a work in process and as I learn and evolve so will this project.
In the meantime, I hope y'all like what I have done here.


.




[Gorrlia Mux]:https://github.com/gorilla/mux
[RS CORS]:https://github.com/rs/cors
[XMLHttpRequest]: https://developer.mozilla.org/en-US/docs/Web/API/XMLHttpRequest
[Docker]: https://www.docker.com/
[Apache]: https://httpd.apache.org/
[MongoDB]: https://www.mongodb.com/
