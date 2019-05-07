FROM hypriot/rpi-alpine-scratch

COPY marcel /usr/bin/

WORKDIR /var/lib/marcel/
EXPOSE 8090
ENTRYPOINT ["marcel"]
