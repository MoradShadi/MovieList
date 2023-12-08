FROM postgres:latest

ENV POSTGRES_DB=movie_watchlist_db
ENV POSTGRES_USER=morad
ENV POSTGRES_PASSWORD=password


EXPOSE 5432
