FROM golang:1.13

COPY marcel /usr/bin/

WORKDIR /var/lib/marcel/
ENTRYPOINT ["marcel"]
EXPOSE 8090
