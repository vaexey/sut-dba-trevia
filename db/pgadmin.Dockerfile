FROM dpage/pgadmin4

USER pgadmin
COPY ./pgpass /
COPY ./servers.json /pgadmin4/servers.json

USER root
RUN chown pgadmin /pgpass
RUN chmod 0600 /pgpass

USER pgadmin
ENTRYPOINT ["/entrypoint.sh"]