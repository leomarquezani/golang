FROM iron/base 

EXPOSE 6767

ADD rest-api-linux-amd64 /

ENTRYPOINT [ "./rest-api-linux-amd64" ]