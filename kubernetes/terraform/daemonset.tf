resource "kubernetes_namespace" "kt_ns" {
  metadata {
    name = local.namespace_name
  }
}

resource "kubernetes_service_account" "kt_sa" {
  metadata {
    name = local.resource_unified_name
    namespace = kubernetes_namespace.kt_ns.id
  }
}

resource "kubernetes_role" "kt_rl" {
  metadata {
    name = local.resource_unified_name
    namespace = kubernetes_namespace.kt_ns.id
  }
  rule {
    api_groups = [ "k8s.sers.dev" ]
    resources = [ "ktbans" ]
    verbs = [ "get", "watch", "list", "delete" ]
  }
}

resource "kubernetes_role_binding" "kt_rb" {
  metadata {
    name = local.resource_unified_name
    namespace = kubernetes_namespace.kt_ns.id
  }
  role_ref {
    api_group = "rbac.authorization.k8s.io"
    kind      = "Role"
    name      = local.resource_unified_name
  }
  subject {
    kind = "ServiceAccount"
    name = local.resource_unified_name
    namespace = kubernetes_namespace.kt_ns.id
  }
}

resource "kubernetes_daemonset" "kt_ds" {
  metadata {
    name = local.resource_unified_name
    namespace = kubernetes_namespace.kt_ns.id
  }
  spec {
    selector {
      match_labels = {
        app = "kubetables"
      }
    }
    template {
      metadata {
        labels = {
          app = "kubetables"
        }
      }
      spec {
        host_network = true 
        service_account_name = kubernetes_service_account.kt_sa.metadata[0].name
        toleration {
          key = "node-role.kubernetes.io/master"
          operator = "Exists"
          effect = "NoSchedule"
        }
        container {
          name = "kubetables"
          image = "${local.kubetables_image_name}:${local.kubetables_image_tag}"
          security_context {
            capabilities {
              add = [ "NET_ADMIN" ]
            }
          }
          resources {
            limits = {
              memory = "200Mi"
            }
            requests = {
              memory = "200Mi"
              cpu = "100m"
            }
          }
        }
      }
    }
  }
  
}
