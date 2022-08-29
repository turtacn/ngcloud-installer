# k8OS
k3OS is a Linux distribution designed to remove as much OS maintenance
as possible in a Kubernetes cluster.  It is specifically designed to only
have what is needed to run k8s. Additionally
the OS is designed to be managed by kubectl once a cluster is bootstrapped.
Nodes only need to join a cluster and then all aspects of the OS can be managed
from Kubernetes. Both k3OS and k3s upgrades are handled by the k3OS operator.

