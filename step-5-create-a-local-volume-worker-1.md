# Manage your cluster to be able retain data using some "local disk"

## Accesse o worker-1 e depois crie um caminho para usarmos depois
```bash
$ vagrant ssh worker-1

$ sudo mkdir -p /mnt/disks/ssd1
$ sudo vim /mnt/disks/ssd1/index.html
# Escreve qualquer coisa legal :) e salva o arquivo
```

## Crie um storage class para seu persistent volume e persistent volume claim

*local_storage_class.yaml :*
```yaml
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: local-storage
provisioner: kubernetes.io/no-provisioner
volumeBindingMode: WaitForFirstConsumer
```
*rode o comando*
```bash
kubectl create -f local_storage_class.yaml
```

## Crie um persistent volume

*pv_local_storage.yaml :*
```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: local-pv
spec:
  capacity:
    storage: 10Gi
  volumeMode: Filesystem
  accessModes:
  - ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain
  storageClassName: local-storage
  local:
    path: /mnt/disks/ssd1
  nodeAffinity:
    required:
      nodeSelectorTerms:
      - matchExpressions:
        - key: kubernetes.io/hostname
          operator: In
          values:
          - worker-1
```
*rode o comando*
```bash
kubectl create -f pv_local_storage.yaml
```

## Crie um persistent volume claim

*pvc_local_storage.yaml :*
```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: local-pvc
spec:
  accessModes:
    - ReadWriteOnce
  volumeMode: Filesystem
  resources:
    requests:
      storage: 2Gi
  storageClassName: local-storage
```
*rode o comando*
```bash
kubectl create -f pvc_local_storage.yaml
```

# Why all of that? describe...

## Testando nosso volume em um pod do nginx

*nginx_pod.yaml :*
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: mypod
spec:
  containers:
    - name: mynginx
      image: nginx
      volumeMounts:
      - mountPath: "/usr/share/nginx/html"
        name: mypd
  volumes:
    - name: mypd
      persistentVolumeClaim:
        claimName: local-pvc
```
*rode o comando*
```bash
kubectl create -f nginx_pod.yaml
```

### Se você fez tudo certinho, ao entrar no seu pod e rodar um curl apontando a localhost:80 você vai ver o conteudo do seu index custom.
#Exemplo aqui
