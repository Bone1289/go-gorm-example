# go-gorm-example

#Docker
docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres
user: postgres
password: mysecretpassword

docker run -p 80:80 -e 'PGADMIN_DEFAULT_EMAIL=user@domain.com' -e 'PGADMIN_DEFAULT_PASSWORD=SuperSecret' -d dpage/pgadmin4