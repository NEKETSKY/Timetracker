FROM webdevops/liquibase:postgres

USER root

WORKDIR /liquibase

COPY ./migrator.sh .
RUN chmod +x ./migrator.sh

COPY ./migrations .

ENTRYPOINT ["./migrator.sh"]

CMD ["update"]