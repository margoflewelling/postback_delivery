FROM php:7.4-apache
RUN apt-get update -qq && apt-get install -y
RUN pecl install redis-5.1.1 \
&& pecl install xdebug-2.8.1 \
&& docker-php-ext-enable redis xdebug

COPY . /var/www/html

EXPOSE 80
