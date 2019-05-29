kubectl delete -f deploy/operator.yaml
kubectl delete -f deploy/role_binding.yaml
kubectl delete -f deploy/role.yaml
kubectl delete -f deploy/service_account.yaml
kubectl delete -f deploy/crds/apistatus_v1alpha1_apistatus_cr.yaml
kubectl delete -f deploy/crds/apistatus_v1alpha1_apistatus_crd.yaml
