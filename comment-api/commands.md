docker build . -t comment-api:v1

cd deployment/comment-api

kubectl apply -f pod.yaml

kubectl apply -f deployment.yaml

kubectl apply -f service.yaml

kubectl port-forward deployment/comment-api 8081