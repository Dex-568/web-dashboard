FROM mcr.microsoft.com/mssql/server:2022-latest

ENV ACCEPT_EULA=Y

USER root
COPY init.sql /init.sql
COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

USER mssql
EXPOSE 1433

CMD ["/entrypoint.sh"]