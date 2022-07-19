# Creating images and pushing

## Frontend
```
docker build -f frontend-service.dockerfile -t tcharlezin/front-end:1.0.1 .
docker push tcharlezin/front-end:1.0.1
```

## Logger
```
docker build -f logger-service.dockerfile -t tcharlezin/logger-service:1.0.0 .
docker push tcharlezin/logger-service:1.0.0
```

## Broker
```
docker build -f broker-service.dockerfile -t tcharlezin/broker-service:1.0.0 .
docker push tcharlezin/broker-service:1.0.0
```

## Authentication
```
docker build -f authentication-service.dockerfile -t tcharlezin/authentication-service:1.0.0 .
docker push tcharlezin/authentication-service:1.0.0
```

## Listener
```
docker build -f listener-service.dockerfile -t tcharlezin/listener-service:1.0.0 .
docker push tcharlezin/listener-service:1.0.0
```

## Mail
```
docker build -f mail-service.dockerfile -t tcharlezin/mail-service:1.0.0 .
docker push tcharlezin/mail-service:1.0.0
```


# Initializing docker swarm
```
docker swarm init
```

## Generate something like
```
docker swarm join --token SWMTKN-1-6ceqc3z37ylwk4i6lvablja79qist50vkyz1ux0s2bmnv9rhke-7cbis3pnr5ol8prhlxm24jb82 192.168.65.3:2377
```

## To get the ID for include workers
```
docker swarm join-token worker
```

## To get the ID for include manager
```
docker swarm join-token manager
```

# Deploy Docker Swarm
```
docker stack deploy -c swarm.yml --resolve-image never {myapp}
```

## Show all docker services
```
docker service ls
```

## Scaling a service
```
docker service scale {myapp_listener-service}={2}
```

## Updating a service to a new version
```
docker service scale myapp_logger-service=2
docker service update --image tcharlezin/logger-service:1.0.1 myapp_logger-service
```

## Stop all services running
```
docker stack rm myapp
```

## Leaving the swarm
```
docker swarm leave
docker swarm leave --force
```