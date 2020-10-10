FROM centos:7
COPY golang-like-service /
COPY config.toml /
EXPOSE 8080
RUN chmod +x golang-like-service
CMD ["/bin/bash","-l","-c","./golang-like-service"]
