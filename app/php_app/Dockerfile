FROM php:7.4-cli
RUN pecl install redis-5.1.1 \
    && pecl install xdebug-2.8.1 \
    && docker-php-ext-enable redis xdebug

COPY . /app/php_app/
CMD [ "php", "index.php" ]

EXPOSE 8000
