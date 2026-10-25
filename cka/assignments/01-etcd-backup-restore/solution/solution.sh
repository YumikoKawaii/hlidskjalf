sudo -i
export KUBECONFIG=/etc/kubernetes/admin.conf
export ETCDCTL_CACERT=/etc/kubernetes/pki/etcd/ca.crt
export ETCDCTL_CERT=/etc/kubernetes/pki/etcd/healthcheck-client.crt
export ETCDCTL_KEY=/etc/kubernetes/pki/etcd/healthcheck-client.key
etcdctl snapshot save /opt/backup/etcd-snap.db
etcdutl snapshot status /opt/backup/etcd-snap.db
kubectl create namespace canary
etcdutl snapshot restore /opt/backup/etcd-snap.db --data-dir /var/lib/etcd-restore
sed -i 's|path: /var/lib/etcd$|path: /var/lib/etcd-restore|' /etc/kubernetes/manifests/etcd.yaml
crictl stop $(crictl ps --name kube-apiserver -q)