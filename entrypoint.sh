#!/bin/bash

/opt/mssql/bin/sqlservr &


# waiting for the sqlserver to run
sleep 30

/opt/mssql-tools18/bin/sqlcmd -S localhost -U sa -P "$MSSQL_SA_PASSWORD" -i init.sql -C
wait