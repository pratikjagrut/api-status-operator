kubectl apply -f deploy/operator.yaml
kubectl apply -f deploy/role_binding.yaml
kubectl apply -f deploy/role.yaml
kubectl apply -f deploy/service_account.yaml
#kubectl apply -f deploy/crds/apistatus_v1alpha1_apistatus_cr.yaml
#kubectl apply -f deploy/crds/apistatus_v1alpha1_apistatus_crd.yaml
