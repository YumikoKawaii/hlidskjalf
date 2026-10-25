kubectl drain lima-cp1 --ignore-daemonsets --delete-emptydir-data

sudo apt-mark unhold kubeadm kubectl kubelet
echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.37/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list
sudo apt-get update
sudo apt-get install -y kubeadm='1.37.1-1.1'
sudo kubeadm upgrade plan
sudo kubeadm upgrade apply -y v1.37.1
sudo apt-get install -y kubectl='1.37.1-1.1' kubelet='1.37.1-1.1'
sudo apt-mark hold kubeadm kubectl kubelet

sudo systemctl daemon-reload
sudo systemctl restart kubelet

kubectl uncordon lima-cp1