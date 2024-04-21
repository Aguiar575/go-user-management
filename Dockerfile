FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=root

EXPOSE 3306

VOLUME /Users/arthuraguiar/ContainerVolumes/mysql

CMD ["mysqld"]
