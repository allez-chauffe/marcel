FROM nginx

WORKDIR /usr/share/nginx/html

ADD build .
COPY nginx.conf /etc/nginx/nginx.conf

EXPOSE 80