FROM postgres:12

RUN apt-get update
RUN apt-get -y install pgagent

RUN apt-get -y install postgis postgresql-12-postgis-3 && \
apt-get install postgresql-12-postgis-3-scripts
