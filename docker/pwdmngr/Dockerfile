#Using an Ubuntu base image
FROM ubuntu:latest

#Copying pwdmngr.exe to the container
COPY . /bin

#Make pwdmngr.sh executable
RUN chmod +x /bin/pwdmngr.sh

CMD ["pwdmngr.sh"]