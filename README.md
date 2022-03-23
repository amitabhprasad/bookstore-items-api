# bookstore-items-api
book store items API
- Uses gorilla/mux as http server instead of gin/gonic

##

    (Http engine)
router (gorilla/mux) --> Items Controller --> Item service --> Item (domain)
        |                              |
      User                       Elastic Search

- single instance of elastic search can manage around 50,000 request per min
- to scale this to 1 million request per min we can add more instance and it can grow vertically
- normally elastic search is not used as primary DB, this is because if we have bad configuraiton in 
elasticsearch cluster then we stand risk of loosing some data.
- Ideal configuraiton is to have DB (SQL or No-SQL) as persistence and use elastic search as search engine for faster search.


--- Install ES using docker ---
- https://www.elastic.co/guide/en/elasticsearch/reference/current/docker.html
- docker network create elastic
  - 41e3563e96f72237fb3b9652c4ed035b2dd1293fe0546516685248bbe6569880
- docker run --name es01 --net elastic -p 9200:9200 -p 9300:9300 --volume /Users/amitabhprasad/my-app-data/bookstore-app/elasticsearch/data:/usr/share/elasticsearch/data -it -d docker.elastic.co/elasticsearch/elasticsearch:8.1.0 

- reset password 
  - docker exec -it es01 /bin/bash
     bin/elasticsearch-reset-password -u elastic
  - password: F0=HQFdS=kOx3Y0nilFT
    "cluster_name" : "docker-cluster",
    "cluster_uuid" : "evULS7ofQ4aR0fC0hBovKw",

- chown -R elasticsearch:elasticsearch for the mounted directory if needed
- create user elasticsearch
- sysctl -w vm.max_map_count=262144
- If you want to set this permanently, you need to edit /etc/sysctl.conf and set vm.max_map_count to 262144.

curl -k -u  elastic http://localhost:9200

- shards 
  - number of replicas of the shards are normally placed in different nodes
  - number of shards and the replicas
