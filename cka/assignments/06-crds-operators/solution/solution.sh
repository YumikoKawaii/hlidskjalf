kubectl get crd -o custom-columns=NAME:.metadata.name | grep cert-manager.io > /opt/cka/cm-crds.txt
## manually remove those *.cert-manager.io groups
