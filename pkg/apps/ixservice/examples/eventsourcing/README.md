# Eventsourcing Demo

The eventstore demo uses CockroachDB as underlying database system. Follow the below mentioned steps to run the demo with CockroachDB running in a docker container:

1. Run steps 1 and 2 explained in [Start a local Cockroach Cluster](https://www.cockroachlabs.com/docs/stable/start-a-local-cluster-in-docker.html)

2. Once the docker container is set up and running execute
   
    ```docker exec roach1 ./cockroach user set test --insecure``` 
    
    This will create a new user ```test```

3. Next run
   
    ```docker exec roach1 ./cockroach sql --insecure -e 'CREATE DATABASE eventstore'``` 

    which will create a new database called ```eventstore```

4. Last execute

    ```docker exec roach1 ./cockroach sql --insecure -e 'GRANT ALL ON DATABASE eventstore TO test'``` 

    which grants the ```test``` user access to the ```eventstore``` database.

5. Run ```go run main.go``` and look at the log output which documents the steps in persisting initiated events.

6. Check if the data is properly stored in the database by executing the following commands:

    - ```docker exec -it roach1 ./cockroach sql --insecure```
    
    Once logged in to the database within the container execute:
    
    - ```use eventstore;```
    
    - ```select * from events;```

    - ```select * from snapshots;```


## Further reading

A more detailed explanation of the ixservice eventsourcing approach can be found under [ixstorage Eventsourcing Quick Start Guide](../../../../../docs/eventsourcing/README.md)