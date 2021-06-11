# Manage your go app 1 pod to use the local storage created before

### Vamos pegar o yaml feito para o go app 1 e modificar ele para usar nosso volume e também colocar uma variavel de ambiente para dizer qual é o nosso arquivo request

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
    env:
    - name: REQUEST_FILE_NAME
      value: "/var/goapp1/upa_cavalinho.json"
    resources:
      limits:
        memory: "128Mi"
        cpu: "200m"
    ports:
      - containerPort: 8080

    volumeMounts:
      - mountPath: "/var/goapp1"
        name: local-pv
  volumes:
    - name: local-pv
      persistentVolumeClaim:
        claimName: local-pvc
```

## Faça o deploy dessa nova versão do pod

*Se você estiver com o pod antigo rodando ainda, delete ele*
```bash
kubectl delete pod golang-pod1
```

## Verifique se está tudo funcionando como deveria
```bash
$ kubectl get pods # verifique se o status está como running
$ kubectl exec --stdin --tty golang-pod1 -- /bin/bash # entre no bash
$ curl localhost:8080/ready # valide o nosso helthcheck
{"status": "Ok"}
$ curl localhost:8080/write # valide se nosso pod consegue escrever no arquivo 2x
{"request_status": "created"}

# se você chamou o /write 2x, ao rodar o /read você deve ver algo como isso
$ curl localhost:8080/read
{"requests":[{"request_id":1,"date":1623419896},{"request_id":2,"date":1623419906}]}
```

## legal né ? hehe Bora ver se está funcionando tudo como esperado. Entre no worker-1 e verifique se nosso arquivo está lá com o conteúdo esperado

```bash
$ vagrant ssh worker-1 # entre no worker
$ cat /mnt/disks/ssd1/upa_cavalinho.json
{
 "requests": [
  {
   "request_id": 1,
   "date": 1623419896
  },
  {
   "request_id": 2,
   "date": 1623419906
  }
 ]
}
# nome maneiro do arquivo não é mesmo ? de nada :)
```
