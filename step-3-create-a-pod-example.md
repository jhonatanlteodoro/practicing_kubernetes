## Creating a simple deployment for our apps


-# yaml definition here!!!


## Mande o kubernetes fazer o deploy to seu pod

### Deploy python app
1º realize o deploy no k8s
```bash
$ kubectl apply -f python-app-pod.yaml
```

2º Verifique se está tudo certo
```bash
$ kubectl get pods #check the status
$ kubectl exec --stdin --tty python-pod -- /bin/bash #access the terminal
$ curl localhost:5000 #check if the server is ok
```

### Deploy golang app 1
1º realize o deploy no k8s
```bash
$ kubectl apply -f golang-app1-pod.yaml
```

2º Verifique se está tudo certo
```bash
$ kubectl get pods #check the status
$ kubectl exec --stdin --tty golang-pod1 -- /bin/bash #access the terminal
$ curl localhost:8080 #check if the server is ok
```

### Deploy golang app 2
1º realize o deploy no k8s
```bash
$ kubectl apply -f golang-app2-pod.yaml
```

2º Verifique se está tudo certo
```bash
$ kubectl get pods #check the status
$ kubectl exec --stdin --tty golang-pod2 -- /bin/bash #access the terminal
$ curl localhost:8080 #check if the server is ok
```

<br>

## Se todos os passos anteriores foram feitos com sucesso, se você der uma olhada nos pods teremos algo semelhante a isso como resultado
```bash
$ kubectl get pods
NAME          READY   STATUS    RESTARTS   AGE
golang-pod1   1/1     Running   0          3m8s
golang-pod2   1/1     Running   0          52s
python-pod    1/1     Running   0          26m
```
