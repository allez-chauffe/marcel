FROM tobi312/rpi-nginx

WORKDIR /var/www/html

COPY . .

EXPOSE 80