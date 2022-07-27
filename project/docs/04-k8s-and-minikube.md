# Starting minikube
minikube start --nodes=2

minikube status

minikube stop

minikube dashboard

kubectl get pods

kubectl get pods -A

kubectl apply -f k8s

kubectl apply -f k8s/broker.yml

kubectl get svc

kubectl get deployments

kubectl delete deployments broker-service mongo rabbitmq

kubectl delete svc broker-service mongo rabbitmq

# Start as LoadBalancer to expose the cluster
kubectl expose deployment broker-service --type=LoadBalancer --port=8080 --target-port=8080
# Starting tunnel minikube to cluster
minikube tunnel
# Stop the tunnel and turn off load balancer of broker-service
kubectl delete svc broker-service


# Enable minikube ingress
minikube addons enable ingress
