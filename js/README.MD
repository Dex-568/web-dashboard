# Web Based Dashboard

## Installation

```
docker build -t my-mssql-init .

docker run -e "ACCEPT_EULA=Y" -e "MSSQL_SA_PASSWORD=<PASSWORD>" -p 1433:1433 my-mssql-init
