## Mande o kubernetes fazer o deploy do seu pod

### Deploy python app
1º realize o deploy no k8s

*python-app-pod.yaml :*
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: python-pod
  labels:
    name: python-app
spec:
  containers:
  - name: python-app
    image: jhonatanlteodoro/app_python:latest
    resources:
      limits:
        memory: "128Mi"
        cpu: "200m"
    ports:
      - containerPort: 5000
```
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

*golang-app1-pod.yaml :*
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: golang-pod1
  labels:
    name: golang-app
spec:
  containers:
  - name: golang-app1
    image: jhonatanlteodoro/go_app1:latest
    resources:
      limits:
        memory: "128Mi"
        cpu: "200m"
    ports:
      - containerPort: 8080
```
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

*golang-app2-pod.yaml :*
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: golang-pod2
  labels:
    name: golang-app
spec:
  containers:
  - name: golang-app2
    image: jhonatanlteodoro/go_app2:latest
    resources:
      limits:
        memory: "128Mi"
        cpu: "200m"
    ports:
      - containerPort: 8080

```
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
