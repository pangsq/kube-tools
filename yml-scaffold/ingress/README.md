# Ingress

https://kubernetes.io/docs/concepts/services-networking/ingress/

https://kubernetes.io/docs/concepts/services-networking/ingress-controllers/

```
internet  ->   service  ->  ingress controller(nginx/traefik/istio/kong等等) -> ingress(rules/backend) -> services
```

在集群中使用ingress需要创建三种资源：

1. Service，有条件用LoadBalance，没条件就用Nodeport或者ingress controller的pod直接使用host network
2. Ingress Controller，运行ingress服务的实例，控制器一般用Deployment（随机应变），目前流行的有nginx/traefik/istio/kong等；在部署的时候一般是控制器、configmap、serviceaccount、RBAC等结合
3. Ingress，配置路由规则、负载均衡、限速、认证、TLS等等，包含rules/backend

## ingress控制器

- nginx - https://kubernetes.github.io/ingress-nginx/

## yaml

- ingress
    - [example-ingress.yaml](example-ingress.yaml)
    - [basic-auth-ingress.yaml](basic-auth-ingress.yaml)
- controller
    - [nginx-ingress-controller.yaml](nginx-ingress-controller) 
        <details><summary>来源</summary>
        https://github.com/kubernetes/ingress-nginx/blob/master/deploy/static/provider/baremetal/deploy.yaml
        </details>

    - [traefik-ingress-controller.yaml](traefik-ingress-controller.yaml)
        <details><summary>来源</summary>
        https://docs.traefik.io/user-guides/crd-acme/
        </details>