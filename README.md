# k8s-CRD-Controller

To run this project

```bash
    go build -o statefulcontroller .
   ./statefulcontroller -kubeconfig=$HOME/.kube/config
   kubectl create -f artifacts/example/fooDB.yaml
   kubectl apply -f artifacts/example/example-fooDB.yaml
```
