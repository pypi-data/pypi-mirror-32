import json

none = "d3043820717d74d9a17694c176d39733"


# region Elastigroup
class Elastigroup:
    def __init__(self, name=none, description=none, capacity=none, strategy=none,
                 compute=none, scaling=none,
                 scheduling=none, multai=none,
                 third_parties_integration=none):
        self.name = name
        self.description = description
        self.capacity = capacity
        self.strategy = strategy
        self.scaling = scaling
        self.scheduling = scheduling
        self.multai = multai
        self.third_parties_integration = third_parties_integration
        self.compute = compute


# endregion

# region Strategy
class Strategy:
    def __init__(self, availability_vs_cost=none, risk=none, utilize_reserved_instances=none,
                 fallback_to_od=none,
                 on_demand_count=none, draining_timeout=none,
                 spin_up_time=none, lifetime_period=none, signals=none,
                 scaling_strategy=none, persistence=none):
        self.risk = risk
        self.utilize_reserved_instances = utilize_reserved_instances
        self.fallback_to_od = fallback_to_od
        self.on_demand_count = on_demand_count
        self.availability_vs_cost = availability_vs_cost
        self.draining_timeout = draining_timeout
        self.spin_up_time = spin_up_time
        self.lifetime_period = lifetime_period
        self.spin_up_time = spin_up_time
        self.signals = signals
        self.scaling_strategy = scaling_strategy
        self.persistence = persistence


class Signal:
    def __init__(self, name=none, timeout=none):
        self.name = name
        self.timeout = timeout


class ScalingStrategy:
    def __init__(self, terminate_at_end_of_billing_hour):
        self.terminate_at_end_of_billing_hour = terminate_at_end_of_billing_hour


class Persistence:
    def __init__(self, should_persist_block_devices=none, should_persist_root_device=none,
                 should_persist_private_ip=none):
        self.should_persist_block_devices = should_persist_block_devices
        self.should_persist_root_device = should_persist_root_device
        self.should_persist_private_ip = should_persist_private_ip


# endregion

# region Capacity
class Capacity:
    def __init__(self, minimum=none, maximum=none, target=none, unit=none):
        self.minimum = minimum
        self.maximum = maximum
        self.target = target
        self.unit = unit


# endregion

# region Scaling
class Scaling:
    def __init__(self, up=none, down=none, target=none):
        self.up = up
        self.down = down
        self.target = target


class ScalingPolicyDimension:
    def __init__(self, name=none, value=none):
        self.name = name
        self.value = value


class ScalingPolicyAction:
    def __init__(self, type=none, adjustment=none, min_target_capacity=none,
                 max_target_capacity=none, target=none,
                 minimum=none,
                 maximum=none):
        self.type = type
        self.adjustment = adjustment
        self.min_target_capacity = min_target_capacity
        self.max_target_capacity = max_target_capacity
        self.target = target
        self.minimum = minimum
        self.maximum = maximum


class ScalingPolicy:
    def __init__(self, namespace=none, metric_name=none, statistic=none,
                 evaluation_periods=none, period=none,
                 threshold=none,
                 cooldown=none, action=none, unit=none, operator=none,
                 dimensions=none, policy_name=none):
        self.policy_name = policy_name
        self.namespace = namespace
        self.metric_name = metric_name
        self.dimensions = dimensions
        self.statistic = statistic
        self.evaluation_periods = evaluation_periods
        self.period = period
        self.threshold = threshold
        self.cooldown = cooldown
        self.action = action
        self.unit = unit
        self.operator = operator


class TargetTrackingPolicy:
    def __init__(self, namespace=none, metric_name=none, statistic=none,
                 cooldown=none, target=none, unit=none,
                 dimensions=none, policy_name=none, source=none):
        self.policy_name = policy_name
        self.namespace = namespace
        self.source = source
        self.metric_name = metric_name
        self.dimensions = dimensions
        self.statistic = statistic
        self.unit = unit
        self.cooldown = cooldown
        self.target = target


# endregion

# region Scheduling
class Scheduling:
    def __init__(self, tasks=none):
        self.tasks = tasks


class ScheduledTask:
    def __init__(self, task_type=none, scale_target_capacity=none,
                 scale_min_capacity=none,
                 scale_max_capacity=none, batch_size_percentage=none, grace_period=none,
                 adjustment=none,
                 adjustment_percentage=none, is_enabled=none,
                 frequency=none, cron_expression=none, ):
        self.is_enabled = is_enabled
        self.frequency = frequency
        self.cron_expression = cron_expression
        self.task_type = task_type
        self.scale_target_capacity = scale_target_capacity
        self.scale_min_capacity = scale_min_capacity
        self.scale_max_capacity = scale_max_capacity
        self.batch_size_percentage = batch_size_percentage
        self.grace_period = grace_period
        self.adjustment = adjustment
        self.adjustment_percentage = adjustment_percentage


# endregion

# region Multai
class Multai:
    def __init__(self, token=none, balancers=none):
        self.token = token
        self.balancers = balancers


class MultaiLoadBalancer:
    def __init__(self, project_id=none, balancer_id=none, target_set_id=none, az_awareness=none,
                 auto_weight=none):
        self.project_id = project_id
        self.balancer_id = balancer_id
        self.target_set_id = target_set_id
        self.az_awareness = az_awareness
        self.auto_weight = auto_weight


# endregion

# region ThirdPartyIntegrations
class Rancher:
    def __init__(self, access_key=none, secret_key=none, master_host=none):
        self.access_key = access_key
        self.secret_key = secret_key
        self.master_host = master_host


class Mesosphere:
    def __init__(self, api_server=none):
        self.api_server = api_server


class ElasticBeanstalk:
    def __init__(self, environment_id=none, deployment_preferences=none):
        self.environment_id = environment_id
        self.deployment_preferences = deployment_preferences


class DeploymentPreferences:
    def __init__(self, automatic_roll=none, batch_size_percentage=none, grace_period=none,
                 strategy=none):
        self.automatic_roll = automatic_roll
        self.batch_size_percentage = batch_size_percentage
        self.grace_period = grace_period
        self.strategy = strategy


class BeanstalkDeploymentStrategy:
    def __init__(self, action=none, should_drain_instances=none):
        self.action = action
        self.should_drain_instances = should_drain_instances


class EcsConfiguration:
    def __init__(self, cluster_name=none):
        self.cluster_name = cluster_name


class KubernetesConfiguration:
    def __init__(self, api_server=none, token=none):
        self.api_server = api_server
        self.token = token


class RightScaleConfiguration:
    def __init__(self, account_id=none, refresh_token=none):
        self.account_id = account_id
        self.refresh_token = refresh_token


class OpsWorksConfiguration:
    def __init__(self, layer_id=none):
        self.layer_id = layer_id


class ChefConfiguration:
    def __init__(self, chef_server=none, organization=none, user=none, pem_key=none,
                 chef_version=none):
        self.chef_server = chef_server
        self.organization = organization
        self.user = user
        self.pem_key = pem_key
        self.chef_version = chef_version


class ThirdPartyIntegrations:
    def __init__(self, rancher=none, mesosphere=none, elastic_beanstalk=none, ecs=none,
                 kubernetes=none, right_scale=none,
                 ops_works=none, chef=none):
        self.rancher = rancher
        self.mesosphere = mesosphere
        self.elastic_beanstalk = elastic_beanstalk
        self.ecs = ecs
        self.kubernetes = kubernetes
        self.right_scale = right_scale
        self.ops_works = ops_works
        self.chef = chef


# endregion

# region Compute
class Compute:
    def __init__(self, launch_specification=none, instance_types=none, product=none,
                 availability_zones=none,
                 elastic_ips=none,
                 ebs_volume_pool=none):
        self.elastic_ips = elastic_ips
        self.instance_types = instance_types
        self.availability_zones = availability_zones
        self.product = product
        self.ebs_volume_pool = ebs_volume_pool
        self.launch_specification = launch_specification


class AvailabilityZone:
    def __init__(self, name=none, subnet_id=none, subnet_ids=none, placement_group_name=none):
        self.name = name
        self.subnet_id = subnet_id
        self.subnet_ids = subnet_ids
        self.placement_group_name = placement_group_name


class InstanceTypes:
    def __init__(self, ondemand=none, spot=none, weights=none):
        self.ondemand = ondemand
        self.spot = spot
        self.weights = weights


class Weight:
    def __init__(self, instance_type=none, weighted_capacity=none):
        self.instance_type = instance_type
        self.weighted_capacity = weighted_capacity


class EbsVolume:
    def __init__(self, device_name=none, volume_ids=none):
        self.device_name = device_name
        self.volume_ids = volume_ids


class LaunchSpecification:
    def __init__(self, security_group_ids=none, image_id=none, monitoring=none,
                 health_check_type=none,
                 load_balancers_config=none,
                 health_check_grace_period=none, health_check_unhealthy_duration_before_replacement=none,
                 ebs_optimized=none, tenancy=none, iam_role=none, key_pair=none,
                 user_data=none, shutdown_script=none,
                 block_device_mappings=none,
                 network_interfaces=none, tags=none):
        self.load_balancers_config = load_balancers_config
        self.health_check_type = health_check_type
        self.health_check_grace_period = health_check_grace_period
        self.health_check_unhealthy_duration_before_replacement = health_check_unhealthy_duration_before_replacement
        self.security_group_ids = security_group_ids
        self.monitoring = monitoring
        self.ebs_optimized = ebs_optimized
        self.image_id = image_id
        self.tenancy = tenancy
        self.iam_role = iam_role
        self.key_pair = key_pair
        self.user_data = user_data
        self.shutdown_script = shutdown_script
        self.block_device_mappings = block_device_mappings
        self.network_interfaces = network_interfaces
        self.tags = tags


class LoadBalancersConfig:
    def __init__(self, load_balancers=none):
        self.load_balancers = load_balancers


class LoadBalancer:
    def __init__(self, type=none, arn=none, name=none):
        self.name = name
        self.arn = arn
        self.type = type


class IamRole:
    def __init__(self, name=none, arn=none):
        self.name = name
        self.arn = arn


class BlockDeviceMapping:
    def __init__(self, device_name=none, ebs=none, no_device=none, virtual_name=none):
        self.device_name = device_name
        self.ebs = ebs
        self.no_device = no_device
        self.virtual_name = virtual_name


class EBS:
    def __init__(self, delete_on_termination=none, encrypted=none, iops=none, snapshot_id=none,
                 volume_size=none,
                 volume_type=none):
        self.delete_on_termination = delete_on_termination
        self.encrypted = encrypted
        self.iops = iops
        self.snapshot_id = snapshot_id
        self.volume_size = volume_size
        self.volume_type = volume_type


class Tag:
    def __init__(self, tag_key=none, tag_value=none):
        self.tag_key = tag_key
        self.tag_value = tag_value


class NetworkInterface:
    def __init__(self, delete_on_termination=none, device_index=none, description=none,
                 secondary_private_ip_address_count=none,
                 associate_public_ip_address=none, groups=none, network_interface_id=none,
                 private_ip_address=none,
                 private_ip_addresses=none, subnet_id=none,
                 associate_ipv6_address=none):
        self.description = description
        self.device_index = device_index
        self.secondary_private_ip_address_count = secondary_private_ip_address_count
        self.associate_public_ip_address = associate_public_ip_address
        self.delete_on_termination = delete_on_termination
        self.groups = groups
        self.network_interface_id = network_interface_id
        self.private_ip_address = private_ip_address
        self.private_ip_addresses = private_ip_addresses
        self.subnet_id = subnet_id
        self.associate_ipv6_address = associate_ipv6_address


class PrivateIpAddress:
    def __init__(self, private_ip_address=none, primary=none):
        self.private_ip_address = private_ip_address
        self.primary = primary


# endregion

class Roll:
    def __init__(self, batch_size_percentage=none, grace_period=none, health_check_type=none, strategy=none):
        self.batch_size_percentage = batch_size_percentage
        self.grace_period = grace_period
        self.health_check_type = health_check_type
        self.strategy = strategy


class DetachConfiguration:
    def __init__(self, instances_to_detach=none, should_terminate_instances=none, draining_timeout=none,
                 should_decrement_target_capacity=none):
        self.instances_to_detach = instances_to_detach
        self.should_terminate_instances = should_terminate_instances
        self.draining_timeout = draining_timeout
        self.should_decrement_target_capacity = should_decrement_target_capacity


class StatefulDeallocation:
    def __init__(self, should_delete_images=none, should_delete_network_interfaces=none, should_delete_volumes=none,
                 should_delete_snapshots=none):
        self.should_delete_images = should_delete_images
        self.should_delete_network_interfaces = should_delete_network_interfaces
        self.should_delete_volumes = should_delete_volumes
        self.should_delete_snapshots = should_delete_snapshots


class ElastigroupCreationRequest:
    def __init__(self, elastigroup):
        self.group = elastigroup

    def toJSON(self):
        return json.dumps(self, default=lambda o: o.__dict__,
                          sort_keys=True, indent=4)


class ElastigroupDeletionRequest:
    def __init__(self, stateful_deallocation):
        self.stateful_deallocation = stateful_deallocation

    def toJSON(self):
        return json.dumps(self, default=lambda o: o.__dict__,
                          sort_keys=True, indent=4)


class ElastigroupUpdateRequest:
    def __init__(self, elastigroup):
        self.group = elastigroup

    def toJSON(self):
        return json.dumps(self, default=lambda o: o.__dict__,
                          sort_keys=True, indent=4)


class ElastigroupRollRequest:
    def __init__(self, group_roll):
        self.batch_size_percentage = group_roll.batch_size_percentage
        self.grace_period = group_roll.grace_period
        self.health_check_type = group_roll.health_check_type
        self.strategy = group_roll.strategy

    def toJSON(self):
        return json.dumps(self, default=lambda o: o.__dict__,
                          sort_keys=True, indent=4)


class ElastigroupDetachInstancesRequest:
    def __init__(self, detach_configuration):
        self.should_decrement_target_capacity = detach_configuration.should_decrement_target_capacity
        self.draining_timeout = detach_configuration.draining_timeout
        self.instances_to_detach = detach_configuration.instances_to_detach
        self.should_terminate_instances = detach_configuration.should_terminate_instances

    def toJSON(self):
        return json.dumps(self, default=lambda o: o.__dict__, sort_keys=True, indent=4)
