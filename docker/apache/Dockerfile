#Using an Apache base image
FROM httpd:2.4

#Adding default servername to Apache config files
RUN echo ServerName localhost >> /usr/local/apache2/conf/httpd.conf

#Copying index.html to the container's webserver directory
COPY index.html /usr/local/apache2/htdocs
