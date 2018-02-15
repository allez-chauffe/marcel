FROM tobi312/rpi-nginx

WORKDIR /var/www/html

ADD build .
EXPOSE 80