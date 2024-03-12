# Utilizamos la imagen base de MySQL 5.7
FROM mysql:5.7

# Definimos las variables de entorno para la configuración de MySQL
ENV MYSQL_ROOT_PASSWORD=supersecurepassword1234
ENV MYSQL_DATABASE=users_db
ENV MYSQL_USER=userdb
ENV MYSQL_PASSWORD=supersecurepassworduser

# Exponemos el puerto 3306 para que se pueda acceder a la base de datos MySQL
EXPOSE 3306

# Copiamos el script SQL al contenedor
COPY script.sql /docker-entrypoint-initdb.d/

# El script.sql se ejecutará automáticamente al iniciar el contenedor

CMD ["mysqld"]
