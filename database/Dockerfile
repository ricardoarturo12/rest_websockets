FROM postgres:10.3

# configuración de la base de datos 1 por que corre primero ese archivo
# copia el archivo dentro de ese directorio
COPY up.sql /docker-entrypoint-initdb.d/1.sql

CMD ["postgres"]