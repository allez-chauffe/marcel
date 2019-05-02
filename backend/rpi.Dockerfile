FROM hypriot/rpi-alpine-scratch

COPY marcel /usr/bin/

ENV MARCEL_LOG_FILE=/var/log/marcel/marcel.log
WORKDIR /var/lib/marcel/
EXPOSE 8090
ENTRYPOINT ["marcel"]
