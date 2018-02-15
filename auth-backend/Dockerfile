FROM hypriot/rpi-alpine-scratch

COPY auth-backend /auth-backend/

ENV MARCEL_LOG_FILE=/backend/logs/auth-backend.log
WORKDIR /auth-backend

ENTRYPOINT ["./auth-backend"]
EXPOSE 8091