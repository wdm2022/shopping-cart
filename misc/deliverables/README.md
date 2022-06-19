# Web-scale Data Management project

## Authors

* Konrad Ponichtera
* Krzysztof Baran
* Rahim Klabér
* Melchior Oudemans
* Koen Hagen
* Dan Plămădeală

## Deployment requirements

1. Configured Kubernetes context
   * Configured user should have permissions to manage resources of the cluster
2. [Helm](https://helm.sh/)
3. (Optionally) [k9s](https://k9scli.io/) text user interface for Kubernetes - for monitoring and resilience testing

## Deployment guide

The system is delivered as a Helm Chart from the repository, hosted on GitHub.
The repository can be added with the following command:

```shell
helm repo add shopping-cart https://wdm2022.github.io/shopping-cart/
```

Then, the system can be installed with:

```shell
helm upgrade --install shopping-cart shopping-cart/shopping-cart
```

The Chart will display in its deployment notes the information on how to connect to the API gateway of the system,
depending on how the system was deployed:
* URL if Ingress was used
* port exposed on one of the cluster's workers if `NodePort` service type was used
* address and port of the load balancer if `LoadBalancer` service type was used
* commands to execute locally to perform port forwarding to the Kubernetes' service if `ClusterIP` service type was used (default)

The default installation will deploy one instance of each microservice and three single-replica MongoDB clusters.
This is preferable for Kubernetes clusters with limited resources (e.g. minikube), but does not offer high availability.
In order to deploy the system with multiple replicas of each microservice and MongoDB clusters with two replicas, execute:

```shell
helm upgrade --install shopping-cart shopping-cart/shopping-cart --values ha-values.yml
```

If the cluster has the Ingress controller configured, it is possible to expose the API gateway through it conveniently
by setting the desired hostname in the _ha-ingress-values.yml_ file and deploying it:

```shell
helm upgrade --install shopping-cart shopping-cart/shopping-cart --values ha-ingress-values.yml
```

The number of replicas can be adjusted in the _ha-values.yml_ and _ha-ingress-values.yml_ files.

## Testing resilience

Our recommended tool for testing the resilience of the system is [k9s](https://k9scli.io/),
a text user interface client (TUI) for Kubernetes.

It gives an accessible insight into the cluster resources and allows to monitor the status of individual Pods
as well as their CPU and memory usage.
It also allows deleting Pods to check how the system behaves when there are failing components in the system.
Moreover, it shows how quickly the cluster recreates these components.