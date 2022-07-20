
## Caddy 
```
docker build --platform linux/x86_64 -f caddy.production.dockerfile -t tcharlezin/micro-caddy-production:1.0.0 .
docker push tcharlezin/micro-caddy-production:1.0.0
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

## View the list of stack
```
sudo docker stack ps myapp
sudo docker node ls
```