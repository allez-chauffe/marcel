FROM hypriot/rpi-alpine-scratch

COPY backend /backend/

ENV MARCEL_LOG_FILE=/backend/logs/backend.log
WORKDIR /backend

ENTRYPOINT ["/backend/backend"]
EXPOSE 8090