## deploy Statefulset
1. Switch to correct context of your kubernetes cluster
2. Replace `DOCKER_REPO` with your docker repository in `litecoin-sts.yaml`
3. `kubectl apply -f litecoin-sts.yaml`