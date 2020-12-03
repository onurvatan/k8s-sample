docker build . -t product-api:v1

docker run -p 8080:8080 -d -t product-api:v1 

kubectl cluster-info

cd deployment/product-api

kubectl apply -f pod.yaml

kubectl apply -f deployment.yaml

kubectl apply -f service.yaml


#kubectl port-forward product-api 8080:8080  

kubectl port-forward deployment/product-api 8080