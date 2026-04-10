output "alb_dns_name" {
  description = "DNS del Application Load Balancer"
  value       = aws_lb.main.dns_name
}

output "ecs_cluster_name" {
  description = "Nombre del cluster ECS"
  value       = aws_ecs_cluster.main.name
}

output "ecs_service_name" {
  description = "Nombre del servicio ECS"
  value       = aws_ecs_service.api.name
}

output "cloudwatch_log_group" {
  description = "Grupo de logs en CloudWatch"
  value       = aws_cloudwatch_log_group.api.name
}
