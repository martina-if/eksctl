---
title: "Road Map"
weight: 20
---

## Project Roadmap

### Developer use-case (0.2.0)

It should suffice to install a cluster for development with just a single command. Here are some examples:

To create a cluster with default configuration (2 `m5.large` nodes), run:

```
eksctl create cluster
```

The developer may choose to pre-configure popular addons, e.g.:

- Weave Net: `eksctl create cluster --networking weave`
- Helm: `eksctl create cluster --addons helm`
- AWS CI tools (CodeCommit, CodeBuild, ECR): `eksctl create cluster --addons aws-ci`
- Jenkins X: `eksctl create cluster --addons jenkins-x`
- AWS CodeStar: `eksctl create cluster --addons aws-codestar`
- Weave Scope and Flux: `eksctl create cluster --addons weave-scope,weave-flux`

It should be possible to combine any or all of these addons.

It would also be possible to add any of the addons after cluster was created with `eksctl create addons`.

### Manage EKS the GitOps way (0.3.0)

Just like `kubectl`, `eksctl` aims to be compliant with GitOps model, and can be used as part of a GitOps toolkit!

For example, `eksctl apply --cluster-config prod-cluster.yaml` will manage cluster state declaratively.

And `eksctld` will be a controller inside of one cluster that can manage multiple other clusters based on Kubernetes Cluster API definitions (CRDs).


## Get in touch

[Create an issue](https://github.com/weaveworks/eksctl/issues/new), or login to [Weave Community Slack (#eksctl)][slackchan] ([signup][slackjoin]).

[slackjoin]: https://slack.weave.works/
[slackchan]: https://weave-community.slack.com/messages/CAYBZBWGL/

> ***Logo Credits***
>
> *Original Gophers drawn by [Ashley McNamara](https://twitter.com/ashleymcnamara), unique E, K, S, C, T & L Gopher identities had been produced with [Gopherize.me](https://gopherize.me/).*
