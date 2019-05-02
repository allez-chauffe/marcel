FROM alpine

COPY marcel /usr/bin/

ENV MARCEL_LOG_FILE=/var/log/marcel/marcel.log
WORKDIR /var/lib/marcel/
ENTRYPOINT ["marcel"]
EXPOSE 8090
