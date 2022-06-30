// Code generated by "enumer -type ResourceType -addprefix azurerm_ -transform snake -linecomment"; DO NOT EDIT.

package azurerm

import (
	"fmt"
	"strings"
)

const _ResourceTypeName = "azurerm_resource_groupazurerm_virtual_machineazurerm_windows_virtual_machineazurerm_linux_virtual_machineazurerm_virtual_machine_extensionazurerm_windows_virtual_machine_scale_setazurerm_linux_virtual_machine_scale_setazurerm_virtual_machine_scale_set_extensionazurerm_virtual_networkazurerm_availability_setazurerm_managed_diskazurerm_imageazurerm_subnetazurerm_network_interfaceazurerm_network_security_groupazurerm_application_gatewayazurerm_application_security_groupazurerm_network_ddos_protection_planazurerm_firewallazurerm_local_network_gatewayazurerm_nat_gatewayazurerm_network_profileazurerm_network_security_ruleazurerm_public_ipazurerm_public_ip_prefixazurerm_routeazurerm_route_tableazurerm_virtual_network_gatewayazurerm_virtual_network_gateway_connectionazurerm_virtual_network_peeringazurerm_web_application_firewall_policyazurerm_virtual_hubazurerm_virtual_hub_bgp_connectionazurerm_virtual_hub_connectionazurerm_virtual_hub_ipazurerm_virtual_hub_route_tableazurerm_virtual_hub_security_partner_providerazurerm_lbazurerm_lb_backend_address_poolazurerm_lb_ruleazurerm_lb_outbound_ruleazurerm_lb_nat_ruleazurerm_lb_nat_poolazurerm_lb_probeazurerm_virtual_desktop_host_poolazurerm_virtual_desktop_application_groupazurerm_logic_app_workflowazurerm_logic_app_trigger_customazurerm_logic_app_action_customazurerm_container_registryazurerm_container_registry_webhookazurerm_kubernetes_clusterazurerm_kubernetes_cluster_node_poolazurerm_storage_accountazurerm_storage_queueazurerm_storage_shareazurerm_storage_tableazurerm_storage_blobazurerm_mariadb_configurationazurerm_mariadb_databaseazurerm_mariadb_firewall_ruleazurerm_mariadb_serverazurerm_mariadb_virtual_network_ruleazurerm_mysql_configurationazurerm_mysql_databaseazurerm_mysql_firewall_ruleazurerm_mysql_serverazurerm_mysql_virtual_network_ruleazurerm_postgresql_configurationazurerm_postgresql_databaseazurerm_postgresql_firewall_ruleazurerm_postgresql_serverazurerm_postgresql_virtual_network_ruleazurerm_mssql_elasticpoolazurerm_mssql_databaseazurerm_mssql_firewall_ruleazurerm_mssql_serverazurerm_mssql_server_security_alert_policyazurerm_mssql_server_vulnerability_assessmentazurerm_mssql_virtual_machineazurerm_mssql_virtual_network_ruleazurerm_redis_cacheazurerm_redis_firewall_ruleazurerm_dns_zoneazurerm_dns_a_recordazurerm_dns_aaaa_recordazurerm_dns_caa_recordazurerm_dns_cname_recordazurerm_dns_mx_recordazurerm_dns_ns_recordazurerm_dns_ptr_recordazurerm_dns_srv_recordazurerm_dns_txt_recordazurerm_private_dns_zoneazurerm_private_dns_a_recordazurerm_private_dns_aaaa_recordazurerm_private_dns_cname_recordazurerm_private_dns_mx_recordazurerm_private_dns_ptr_recordazurerm_private_dns_srv_recordazurerm_private_dns_txt_recordazurerm_private_dns_zone_virtual_network_linkazurerm_policy_definitionazurerm_policy_remediationazurerm_policy_set_definitionazurerm_key_vaultazurerm_key_vault_access_policyazurerm_application_insightsazurerm_application_insights_api_keyazurerm_application_insights_analytics_itemazurerm_log_analytics_workspaceazurerm_log_analytics_linked_serviceazurerm_log_analytics_datasource_windows_performance_counterazurerm_log_analytics_datasource_windows_eventazurerm_monitor_action_groupazurerm_monitor_activity_log_alertazurerm_monitor_autoscale_settingazurerm_monitor_log_profileazurerm_monitor_metric_alert"

var _ResourceTypeIndex = [...]uint16{0, 22, 45, 76, 105, 138, 179, 218, 261, 284, 308, 328, 341, 355, 380, 410, 437, 471, 507, 523, 552, 571, 594, 623, 640, 664, 677, 696, 727, 769, 800, 839, 858, 892, 922, 944, 975, 1020, 1030, 1061, 1076, 1100, 1119, 1138, 1154, 1187, 1228, 1254, 1286, 1317, 1343, 1377, 1403, 1439, 1462, 1483, 1504, 1525, 1545, 1574, 1598, 1627, 1649, 1685, 1712, 1734, 1761, 1781, 1815, 1847, 1874, 1906, 1931, 1970, 1995, 2017, 2044, 2064, 2106, 2151, 2180, 2214, 2233, 2260, 2276, 2296, 2319, 2341, 2365, 2386, 2407, 2429, 2451, 2473, 2497, 2525, 2556, 2588, 2617, 2647, 2677, 2707, 2752, 2777, 2803, 2832, 2849, 2880, 2908, 2944, 2987, 3018, 3054, 3114, 3160, 3188, 3222, 3255, 3282, 3310}

const _ResourceTypeLowerName = "azurerm_resource_groupazurerm_virtual_machineazurerm_windows_virtual_machineazurerm_linux_virtual_machineazurerm_virtual_machine_extensionazurerm_windows_virtual_machine_scale_setazurerm_linux_virtual_machine_scale_setazurerm_virtual_machine_scale_set_extensionazurerm_virtual_networkazurerm_availability_setazurerm_managed_diskazurerm_imageazurerm_subnetazurerm_network_interfaceazurerm_network_security_groupazurerm_application_gatewayazurerm_application_security_groupazurerm_network_ddos_protection_planazurerm_firewallazurerm_local_network_gatewayazurerm_nat_gatewayazurerm_network_profileazurerm_network_security_ruleazurerm_public_ipazurerm_public_ip_prefixazurerm_routeazurerm_route_tableazurerm_virtual_network_gatewayazurerm_virtual_network_gateway_connectionazurerm_virtual_network_peeringazurerm_web_application_firewall_policyazurerm_virtual_hubazurerm_virtual_hub_bgp_connectionazurerm_virtual_hub_connectionazurerm_virtual_hub_ipazurerm_virtual_hub_route_tableazurerm_virtual_hub_security_partner_providerazurerm_lbazurerm_lb_backend_address_poolazurerm_lb_ruleazurerm_lb_outbound_ruleazurerm_lb_nat_ruleazurerm_lb_nat_poolazurerm_lb_probeazurerm_virtual_desktop_host_poolazurerm_virtual_desktop_application_groupazurerm_logic_app_workflowazurerm_logic_app_trigger_customazurerm_logic_app_action_customazurerm_container_registryazurerm_container_registry_webhookazurerm_kubernetes_clusterazurerm_kubernetes_cluster_node_poolazurerm_storage_accountazurerm_storage_queueazurerm_storage_shareazurerm_storage_tableazurerm_storage_blobazurerm_mariadb_configurationazurerm_mariadb_databaseazurerm_mariadb_firewall_ruleazurerm_mariadb_serverazurerm_mariadb_virtual_network_ruleazurerm_mysql_configurationazurerm_mysql_databaseazurerm_mysql_firewall_ruleazurerm_mysql_serverazurerm_mysql_virtual_network_ruleazurerm_postgresql_configurationazurerm_postgresql_databaseazurerm_postgresql_firewall_ruleazurerm_postgresql_serverazurerm_postgresql_virtual_network_ruleazurerm_mssql_elasticpoolazurerm_mssql_databaseazurerm_mssql_firewall_ruleazurerm_mssql_serverazurerm_mssql_server_security_alert_policyazurerm_mssql_server_vulnerability_assessmentazurerm_mssql_virtual_machineazurerm_mssql_virtual_network_ruleazurerm_redis_cacheazurerm_redis_firewall_ruleazurerm_dns_zoneazurerm_dns_a_recordazurerm_dns_aaaa_recordazurerm_dns_caa_recordazurerm_dns_cname_recordazurerm_dns_mx_recordazurerm_dns_ns_recordazurerm_dns_ptr_recordazurerm_dns_srv_recordazurerm_dns_txt_recordazurerm_private_dns_zoneazurerm_private_dns_a_recordazurerm_private_dns_aaaa_recordazurerm_private_dns_cname_recordazurerm_private_dns_mx_recordazurerm_private_dns_ptr_recordazurerm_private_dns_srv_recordazurerm_private_dns_txt_recordazurerm_private_dns_zone_virtual_network_linkazurerm_policy_definitionazurerm_policy_remediationazurerm_policy_set_definitionazurerm_key_vaultazurerm_key_vault_access_policyazurerm_application_insightsazurerm_application_insights_api_keyazurerm_application_insights_analytics_itemazurerm_log_analytics_workspaceazurerm_log_analytics_linked_serviceazurerm_log_analytics_datasource_windows_performance_counterazurerm_log_analytics_datasource_windows_eventazurerm_monitor_action_groupazurerm_monitor_activity_log_alertazurerm_monitor_autoscale_settingazurerm_monitor_log_profileazurerm_monitor_metric_alert"

func (i ResourceType) String() string {
	if i < 0 || i >= ResourceType(len(_ResourceTypeIndex)-1) {
		return fmt.Sprintf("ResourceType(%d)", i)
	}
	return _ResourceTypeName[_ResourceTypeIndex[i]:_ResourceTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ResourceTypeNoOp() {
	var x [1]struct{}
	_ = x[ResourceGroup-(0)]
	_ = x[VirtualMachine-(1)]
	_ = x[WindowsVirtualMachine-(2)]
	_ = x[LinuxVirtualMachine-(3)]
	_ = x[VirtualMachineExtension-(4)]
	_ = x[WindowsVirtualMachineScaleSet-(5)]
	_ = x[LinuxVirtualMachineScaleSet-(6)]
	_ = x[VirtualMachineScaleSetExtension-(7)]
	_ = x[VirtualNetwork-(8)]
	_ = x[AvailabilitySet-(9)]
	_ = x[ManagedDisk-(10)]
	_ = x[Image-(11)]
	_ = x[Subnet-(12)]
	_ = x[NetworkInterface-(13)]
	_ = x[NetworkSecurityGroup-(14)]
	_ = x[ApplicationGateway-(15)]
	_ = x[ApplicationSecurityGroup-(16)]
	_ = x[NetworkDdosProtectionPlan-(17)]
	_ = x[Firewall-(18)]
	_ = x[LocalNetworkGateway-(19)]
	_ = x[NatGateway-(20)]
	_ = x[NetworkProfile-(21)]
	_ = x[NetworkSecurityRule-(22)]
	_ = x[PublicIP-(23)]
	_ = x[PublicIPPrefix-(24)]
	_ = x[Route-(25)]
	_ = x[RouteTable-(26)]
	_ = x[VirtualNetworkGateway-(27)]
	_ = x[VirtualNetworkGatewayConnection-(28)]
	_ = x[VirtualNetworkPeering-(29)]
	_ = x[WebApplicationFirewallPolicy-(30)]
	_ = x[VirtualHub-(31)]
	_ = x[VirtualHubBgpConnection-(32)]
	_ = x[VirtualHubConnection-(33)]
	_ = x[VirtualHubIP-(34)]
	_ = x[VirtualHubRouteTable-(35)]
	_ = x[VirtualHubSecurityPartnerProvider-(36)]
	_ = x[Lb-(37)]
	_ = x[LbBackendAddressPool-(38)]
	_ = x[LbRule-(39)]
	_ = x[LbOutboundRule-(40)]
	_ = x[LbNatRule-(41)]
	_ = x[LbNatPool-(42)]
	_ = x[LbProbe-(43)]
	_ = x[VirtualDesktopHostPool-(44)]
	_ = x[VirtualDesktopApplicationGroup-(45)]
	_ = x[LogicAppWorkflow-(46)]
	_ = x[LogicAppTriggerCustom-(47)]
	_ = x[LogicAppActionCustom-(48)]
	_ = x[ContainerRegistry-(49)]
	_ = x[ContainerRegistryWebhook-(50)]
	_ = x[KubernetesCluster-(51)]
	_ = x[KubernetesClusterNodePool-(52)]
	_ = x[StorageAccount-(53)]
	_ = x[StorageQueue-(54)]
	_ = x[StorageShare-(55)]
	_ = x[StorageTable-(56)]
	_ = x[StorageBlob-(57)]
	_ = x[MariadbConfiguration-(58)]
	_ = x[MariadbDatabase-(59)]
	_ = x[MariadbFirewallRule-(60)]
	_ = x[MariadbServer-(61)]
	_ = x[MariadbVirtualNetworkRule-(62)]
	_ = x[MysqlConfiguration-(63)]
	_ = x[MysqlDatabase-(64)]
	_ = x[MysqlFirewallRule-(65)]
	_ = x[MysqlServer-(66)]
	_ = x[MysqlVirtualNetworkRule-(67)]
	_ = x[PostgresqlConfiguration-(68)]
	_ = x[PostgresqlDatabase-(69)]
	_ = x[PostgresqlFirewallRule-(70)]
	_ = x[PostgresqlServer-(71)]
	_ = x[PostgresqlVirtualNetworkRule-(72)]
	_ = x[MssqlElasticpool-(73)]
	_ = x[MssqlDatabase-(74)]
	_ = x[MssqlFirewallRule-(75)]
	_ = x[MssqlServer-(76)]
	_ = x[MssqlServerSecurityAlertPolicy-(77)]
	_ = x[MssqlServerVulnerabilityAssessment-(78)]
	_ = x[MssqlVirtualMachine-(79)]
	_ = x[MssqlVirtualNetworkRule-(80)]
	_ = x[RedisCache-(81)]
	_ = x[RedisFirewallRule-(82)]
	_ = x[DNSZone-(83)]
	_ = x[DNSARecord-(84)]
	_ = x[DNSAaaaRecord-(85)]
	_ = x[DNSCaaRecord-(86)]
	_ = x[DNSCnameRecord-(87)]
	_ = x[DNSMxRecord-(88)]
	_ = x[DNSNsRecord-(89)]
	_ = x[DNSPtrRecord-(90)]
	_ = x[DNSSrvRecord-(91)]
	_ = x[DNSTxtRecord-(92)]
	_ = x[PrivateDNSZone-(93)]
	_ = x[PrivateDNSARecord-(94)]
	_ = x[PrivateDNSAaaaRecord-(95)]
	_ = x[PrivateDNSCnameRecord-(96)]
	_ = x[PrivateDNSMxRecord-(97)]
	_ = x[PrivateDNSPtrRecord-(98)]
	_ = x[PrivateDNSSrvRecord-(99)]
	_ = x[PrivateDNSTxtRecord-(100)]
	_ = x[PrivateDNSZoneVirtualNetworkLink-(101)]
	_ = x[PolicyDefinition-(102)]
	_ = x[PolicyRemediation-(103)]
	_ = x[PolicySetDefinition-(104)]
	_ = x[KeyVault-(105)]
	_ = x[KeyVaultAccessPolicy-(106)]
	_ = x[ApplicationInsights-(107)]
	_ = x[ApplicationInsightsAPIKey-(108)]
	_ = x[ApplicationInsightsAnalyticsItem-(109)]
	_ = x[LogAnalyticsWorkspace-(110)]
	_ = x[LogAnalyticsLinkedService-(111)]
	_ = x[LogAnalyticsDatasourceWindowsPerformanceCounter-(112)]
	_ = x[LogAnalyticsDatasourceWindowsEvent-(113)]
	_ = x[MonitorActionGroup-(114)]
	_ = x[MonitorActivityLogAlert-(115)]
	_ = x[MonitorAutoscaleSetting-(116)]
	_ = x[MonitorLogProfile-(117)]
	_ = x[MonitorMetricAlert-(118)]
}

var _ResourceTypeValues = []ResourceType{ResourceGroup, VirtualMachine, WindowsVirtualMachine, LinuxVirtualMachine, VirtualMachineExtension, WindowsVirtualMachineScaleSet, LinuxVirtualMachineScaleSet, VirtualMachineScaleSetExtension, VirtualNetwork, AvailabilitySet, ManagedDisk, Image, Subnet, NetworkInterface, NetworkSecurityGroup, ApplicationGateway, ApplicationSecurityGroup, NetworkDdosProtectionPlan, Firewall, LocalNetworkGateway, NatGateway, NetworkProfile, NetworkSecurityRule, PublicIP, PublicIPPrefix, Route, RouteTable, VirtualNetworkGateway, VirtualNetworkGatewayConnection, VirtualNetworkPeering, WebApplicationFirewallPolicy, VirtualHub, VirtualHubBgpConnection, VirtualHubConnection, VirtualHubIP, VirtualHubRouteTable, VirtualHubSecurityPartnerProvider, Lb, LbBackendAddressPool, LbRule, LbOutboundRule, LbNatRule, LbNatPool, LbProbe, VirtualDesktopHostPool, VirtualDesktopApplicationGroup, LogicAppWorkflow, LogicAppTriggerCustom, LogicAppActionCustom, ContainerRegistry, ContainerRegistryWebhook, KubernetesCluster, KubernetesClusterNodePool, StorageAccount, StorageQueue, StorageShare, StorageTable, StorageBlob, MariadbConfiguration, MariadbDatabase, MariadbFirewallRule, MariadbServer, MariadbVirtualNetworkRule, MysqlConfiguration, MysqlDatabase, MysqlFirewallRule, MysqlServer, MysqlVirtualNetworkRule, PostgresqlConfiguration, PostgresqlDatabase, PostgresqlFirewallRule, PostgresqlServer, PostgresqlVirtualNetworkRule, MssqlElasticpool, MssqlDatabase, MssqlFirewallRule, MssqlServer, MssqlServerSecurityAlertPolicy, MssqlServerVulnerabilityAssessment, MssqlVirtualMachine, MssqlVirtualNetworkRule, RedisCache, RedisFirewallRule, DNSZone, DNSARecord, DNSAaaaRecord, DNSCaaRecord, DNSCnameRecord, DNSMxRecord, DNSNsRecord, DNSPtrRecord, DNSSrvRecord, DNSTxtRecord, PrivateDNSZone, PrivateDNSARecord, PrivateDNSAaaaRecord, PrivateDNSCnameRecord, PrivateDNSMxRecord, PrivateDNSPtrRecord, PrivateDNSSrvRecord, PrivateDNSTxtRecord, PrivateDNSZoneVirtualNetworkLink, PolicyDefinition, PolicyRemediation, PolicySetDefinition, KeyVault, KeyVaultAccessPolicy, ApplicationInsights, ApplicationInsightsAPIKey, ApplicationInsightsAnalyticsItem, LogAnalyticsWorkspace, LogAnalyticsLinkedService, LogAnalyticsDatasourceWindowsPerformanceCounter, LogAnalyticsDatasourceWindowsEvent, MonitorActionGroup, MonitorActivityLogAlert, MonitorAutoscaleSetting, MonitorLogProfile, MonitorMetricAlert}

var _ResourceTypeNameToValueMap = map[string]ResourceType{
	_ResourceTypeName[0:22]:           ResourceGroup,
	_ResourceTypeLowerName[0:22]:      ResourceGroup,
	_ResourceTypeName[22:45]:          VirtualMachine,
	_ResourceTypeLowerName[22:45]:     VirtualMachine,
	_ResourceTypeName[45:76]:          WindowsVirtualMachine,
	_ResourceTypeLowerName[45:76]:     WindowsVirtualMachine,
	_ResourceTypeName[76:105]:         LinuxVirtualMachine,
	_ResourceTypeLowerName[76:105]:    LinuxVirtualMachine,
	_ResourceTypeName[105:138]:        VirtualMachineExtension,
	_ResourceTypeLowerName[105:138]:   VirtualMachineExtension,
	_ResourceTypeName[138:179]:        WindowsVirtualMachineScaleSet,
	_ResourceTypeLowerName[138:179]:   WindowsVirtualMachineScaleSet,
	_ResourceTypeName[179:218]:        LinuxVirtualMachineScaleSet,
	_ResourceTypeLowerName[179:218]:   LinuxVirtualMachineScaleSet,
	_ResourceTypeName[218:261]:        VirtualMachineScaleSetExtension,
	_ResourceTypeLowerName[218:261]:   VirtualMachineScaleSetExtension,
	_ResourceTypeName[261:284]:        VirtualNetwork,
	_ResourceTypeLowerName[261:284]:   VirtualNetwork,
	_ResourceTypeName[284:308]:        AvailabilitySet,
	_ResourceTypeLowerName[284:308]:   AvailabilitySet,
	_ResourceTypeName[308:328]:        ManagedDisk,
	_ResourceTypeLowerName[308:328]:   ManagedDisk,
	_ResourceTypeName[328:341]:        Image,
	_ResourceTypeLowerName[328:341]:   Image,
	_ResourceTypeName[341:355]:        Subnet,
	_ResourceTypeLowerName[341:355]:   Subnet,
	_ResourceTypeName[355:380]:        NetworkInterface,
	_ResourceTypeLowerName[355:380]:   NetworkInterface,
	_ResourceTypeName[380:410]:        NetworkSecurityGroup,
	_ResourceTypeLowerName[380:410]:   NetworkSecurityGroup,
	_ResourceTypeName[410:437]:        ApplicationGateway,
	_ResourceTypeLowerName[410:437]:   ApplicationGateway,
	_ResourceTypeName[437:471]:        ApplicationSecurityGroup,
	_ResourceTypeLowerName[437:471]:   ApplicationSecurityGroup,
	_ResourceTypeName[471:507]:        NetworkDdosProtectionPlan,
	_ResourceTypeLowerName[471:507]:   NetworkDdosProtectionPlan,
	_ResourceTypeName[507:523]:        Firewall,
	_ResourceTypeLowerName[507:523]:   Firewall,
	_ResourceTypeName[523:552]:        LocalNetworkGateway,
	_ResourceTypeLowerName[523:552]:   LocalNetworkGateway,
	_ResourceTypeName[552:571]:        NatGateway,
	_ResourceTypeLowerName[552:571]:   NatGateway,
	_ResourceTypeName[571:594]:        NetworkProfile,
	_ResourceTypeLowerName[571:594]:   NetworkProfile,
	_ResourceTypeName[594:623]:        NetworkSecurityRule,
	_ResourceTypeLowerName[594:623]:   NetworkSecurityRule,
	_ResourceTypeName[623:640]:        PublicIP,
	_ResourceTypeLowerName[623:640]:   PublicIP,
	_ResourceTypeName[640:664]:        PublicIPPrefix,
	_ResourceTypeLowerName[640:664]:   PublicIPPrefix,
	_ResourceTypeName[664:677]:        Route,
	_ResourceTypeLowerName[664:677]:   Route,
	_ResourceTypeName[677:696]:        RouteTable,
	_ResourceTypeLowerName[677:696]:   RouteTable,
	_ResourceTypeName[696:727]:        VirtualNetworkGateway,
	_ResourceTypeLowerName[696:727]:   VirtualNetworkGateway,
	_ResourceTypeName[727:769]:        VirtualNetworkGatewayConnection,
	_ResourceTypeLowerName[727:769]:   VirtualNetworkGatewayConnection,
	_ResourceTypeName[769:800]:        VirtualNetworkPeering,
	_ResourceTypeLowerName[769:800]:   VirtualNetworkPeering,
	_ResourceTypeName[800:839]:        WebApplicationFirewallPolicy,
	_ResourceTypeLowerName[800:839]:   WebApplicationFirewallPolicy,
	_ResourceTypeName[839:858]:        VirtualHub,
	_ResourceTypeLowerName[839:858]:   VirtualHub,
	_ResourceTypeName[858:892]:        VirtualHubBgpConnection,
	_ResourceTypeLowerName[858:892]:   VirtualHubBgpConnection,
	_ResourceTypeName[892:922]:        VirtualHubConnection,
	_ResourceTypeLowerName[892:922]:   VirtualHubConnection,
	_ResourceTypeName[922:944]:        VirtualHubIP,
	_ResourceTypeLowerName[922:944]:   VirtualHubIP,
	_ResourceTypeName[944:975]:        VirtualHubRouteTable,
	_ResourceTypeLowerName[944:975]:   VirtualHubRouteTable,
	_ResourceTypeName[975:1020]:       VirtualHubSecurityPartnerProvider,
	_ResourceTypeLowerName[975:1020]:  VirtualHubSecurityPartnerProvider,
	_ResourceTypeName[1020:1030]:      Lb,
	_ResourceTypeLowerName[1020:1030]: Lb,
	_ResourceTypeName[1030:1061]:      LbBackendAddressPool,
	_ResourceTypeLowerName[1030:1061]: LbBackendAddressPool,
	_ResourceTypeName[1061:1076]:      LbRule,
	_ResourceTypeLowerName[1061:1076]: LbRule,
	_ResourceTypeName[1076:1100]:      LbOutboundRule,
	_ResourceTypeLowerName[1076:1100]: LbOutboundRule,
	_ResourceTypeName[1100:1119]:      LbNatRule,
	_ResourceTypeLowerName[1100:1119]: LbNatRule,
	_ResourceTypeName[1119:1138]:      LbNatPool,
	_ResourceTypeLowerName[1119:1138]: LbNatPool,
	_ResourceTypeName[1138:1154]:      LbProbe,
	_ResourceTypeLowerName[1138:1154]: LbProbe,
	_ResourceTypeName[1154:1187]:      VirtualDesktopHostPool,
	_ResourceTypeLowerName[1154:1187]: VirtualDesktopHostPool,
	_ResourceTypeName[1187:1228]:      VirtualDesktopApplicationGroup,
	_ResourceTypeLowerName[1187:1228]: VirtualDesktopApplicationGroup,
	_ResourceTypeName[1228:1254]:      LogicAppWorkflow,
	_ResourceTypeLowerName[1228:1254]: LogicAppWorkflow,
	_ResourceTypeName[1254:1286]:      LogicAppTriggerCustom,
	_ResourceTypeLowerName[1254:1286]: LogicAppTriggerCustom,
	_ResourceTypeName[1286:1317]:      LogicAppActionCustom,
	_ResourceTypeLowerName[1286:1317]: LogicAppActionCustom,
	_ResourceTypeName[1317:1343]:      ContainerRegistry,
	_ResourceTypeLowerName[1317:1343]: ContainerRegistry,
	_ResourceTypeName[1343:1377]:      ContainerRegistryWebhook,
	_ResourceTypeLowerName[1343:1377]: ContainerRegistryWebhook,
	_ResourceTypeName[1377:1403]:      KubernetesCluster,
	_ResourceTypeLowerName[1377:1403]: KubernetesCluster,
	_ResourceTypeName[1403:1439]:      KubernetesClusterNodePool,
	_ResourceTypeLowerName[1403:1439]: KubernetesClusterNodePool,
	_ResourceTypeName[1439:1462]:      StorageAccount,
	_ResourceTypeLowerName[1439:1462]: StorageAccount,
	_ResourceTypeName[1462:1483]:      StorageQueue,
	_ResourceTypeLowerName[1462:1483]: StorageQueue,
	_ResourceTypeName[1483:1504]:      StorageShare,
	_ResourceTypeLowerName[1483:1504]: StorageShare,
	_ResourceTypeName[1504:1525]:      StorageTable,
	_ResourceTypeLowerName[1504:1525]: StorageTable,
	_ResourceTypeName[1525:1545]:      StorageBlob,
	_ResourceTypeLowerName[1525:1545]: StorageBlob,
	_ResourceTypeName[1545:1574]:      MariadbConfiguration,
	_ResourceTypeLowerName[1545:1574]: MariadbConfiguration,
	_ResourceTypeName[1574:1598]:      MariadbDatabase,
	_ResourceTypeLowerName[1574:1598]: MariadbDatabase,
	_ResourceTypeName[1598:1627]:      MariadbFirewallRule,
	_ResourceTypeLowerName[1598:1627]: MariadbFirewallRule,
	_ResourceTypeName[1627:1649]:      MariadbServer,
	_ResourceTypeLowerName[1627:1649]: MariadbServer,
	_ResourceTypeName[1649:1685]:      MariadbVirtualNetworkRule,
	_ResourceTypeLowerName[1649:1685]: MariadbVirtualNetworkRule,
	_ResourceTypeName[1685:1712]:      MysqlConfiguration,
	_ResourceTypeLowerName[1685:1712]: MysqlConfiguration,
	_ResourceTypeName[1712:1734]:      MysqlDatabase,
	_ResourceTypeLowerName[1712:1734]: MysqlDatabase,
	_ResourceTypeName[1734:1761]:      MysqlFirewallRule,
	_ResourceTypeLowerName[1734:1761]: MysqlFirewallRule,
	_ResourceTypeName[1761:1781]:      MysqlServer,
	_ResourceTypeLowerName[1761:1781]: MysqlServer,
	_ResourceTypeName[1781:1815]:      MysqlVirtualNetworkRule,
	_ResourceTypeLowerName[1781:1815]: MysqlVirtualNetworkRule,
	_ResourceTypeName[1815:1847]:      PostgresqlConfiguration,
	_ResourceTypeLowerName[1815:1847]: PostgresqlConfiguration,
	_ResourceTypeName[1847:1874]:      PostgresqlDatabase,
	_ResourceTypeLowerName[1847:1874]: PostgresqlDatabase,
	_ResourceTypeName[1874:1906]:      PostgresqlFirewallRule,
	_ResourceTypeLowerName[1874:1906]: PostgresqlFirewallRule,
	_ResourceTypeName[1906:1931]:      PostgresqlServer,
	_ResourceTypeLowerName[1906:1931]: PostgresqlServer,
	_ResourceTypeName[1931:1970]:      PostgresqlVirtualNetworkRule,
	_ResourceTypeLowerName[1931:1970]: PostgresqlVirtualNetworkRule,
	_ResourceTypeName[1970:1995]:      MssqlElasticpool,
	_ResourceTypeLowerName[1970:1995]: MssqlElasticpool,
	_ResourceTypeName[1995:2017]:      MssqlDatabase,
	_ResourceTypeLowerName[1995:2017]: MssqlDatabase,
	_ResourceTypeName[2017:2044]:      MssqlFirewallRule,
	_ResourceTypeLowerName[2017:2044]: MssqlFirewallRule,
	_ResourceTypeName[2044:2064]:      MssqlServer,
	_ResourceTypeLowerName[2044:2064]: MssqlServer,
	_ResourceTypeName[2064:2106]:      MssqlServerSecurityAlertPolicy,
	_ResourceTypeLowerName[2064:2106]: MssqlServerSecurityAlertPolicy,
	_ResourceTypeName[2106:2151]:      MssqlServerVulnerabilityAssessment,
	_ResourceTypeLowerName[2106:2151]: MssqlServerVulnerabilityAssessment,
	_ResourceTypeName[2151:2180]:      MssqlVirtualMachine,
	_ResourceTypeLowerName[2151:2180]: MssqlVirtualMachine,
	_ResourceTypeName[2180:2214]:      MssqlVirtualNetworkRule,
	_ResourceTypeLowerName[2180:2214]: MssqlVirtualNetworkRule,
	_ResourceTypeName[2214:2233]:      RedisCache,
	_ResourceTypeLowerName[2214:2233]: RedisCache,
	_ResourceTypeName[2233:2260]:      RedisFirewallRule,
	_ResourceTypeLowerName[2233:2260]: RedisFirewallRule,
	_ResourceTypeName[2260:2276]:      DNSZone,
	_ResourceTypeLowerName[2260:2276]: DNSZone,
	_ResourceTypeName[2276:2296]:      DNSARecord,
	_ResourceTypeLowerName[2276:2296]: DNSARecord,
	_ResourceTypeName[2296:2319]:      DNSAaaaRecord,
	_ResourceTypeLowerName[2296:2319]: DNSAaaaRecord,
	_ResourceTypeName[2319:2341]:      DNSCaaRecord,
	_ResourceTypeLowerName[2319:2341]: DNSCaaRecord,
	_ResourceTypeName[2341:2365]:      DNSCnameRecord,
	_ResourceTypeLowerName[2341:2365]: DNSCnameRecord,
	_ResourceTypeName[2365:2386]:      DNSMxRecord,
	_ResourceTypeLowerName[2365:2386]: DNSMxRecord,
	_ResourceTypeName[2386:2407]:      DNSNsRecord,
	_ResourceTypeLowerName[2386:2407]: DNSNsRecord,
	_ResourceTypeName[2407:2429]:      DNSPtrRecord,
	_ResourceTypeLowerName[2407:2429]: DNSPtrRecord,
	_ResourceTypeName[2429:2451]:      DNSSrvRecord,
	_ResourceTypeLowerName[2429:2451]: DNSSrvRecord,
	_ResourceTypeName[2451:2473]:      DNSTxtRecord,
	_ResourceTypeLowerName[2451:2473]: DNSTxtRecord,
	_ResourceTypeName[2473:2497]:      PrivateDNSZone,
	_ResourceTypeLowerName[2473:2497]: PrivateDNSZone,
	_ResourceTypeName[2497:2525]:      PrivateDNSARecord,
	_ResourceTypeLowerName[2497:2525]: PrivateDNSARecord,
	_ResourceTypeName[2525:2556]:      PrivateDNSAaaaRecord,
	_ResourceTypeLowerName[2525:2556]: PrivateDNSAaaaRecord,
	_ResourceTypeName[2556:2588]:      PrivateDNSCnameRecord,
	_ResourceTypeLowerName[2556:2588]: PrivateDNSCnameRecord,
	_ResourceTypeName[2588:2617]:      PrivateDNSMxRecord,
	_ResourceTypeLowerName[2588:2617]: PrivateDNSMxRecord,
	_ResourceTypeName[2617:2647]:      PrivateDNSPtrRecord,
	_ResourceTypeLowerName[2617:2647]: PrivateDNSPtrRecord,
	_ResourceTypeName[2647:2677]:      PrivateDNSSrvRecord,
	_ResourceTypeLowerName[2647:2677]: PrivateDNSSrvRecord,
	_ResourceTypeName[2677:2707]:      PrivateDNSTxtRecord,
	_ResourceTypeLowerName[2677:2707]: PrivateDNSTxtRecord,
	_ResourceTypeName[2707:2752]:      PrivateDNSZoneVirtualNetworkLink,
	_ResourceTypeLowerName[2707:2752]: PrivateDNSZoneVirtualNetworkLink,
	_ResourceTypeName[2752:2777]:      PolicyDefinition,
	_ResourceTypeLowerName[2752:2777]: PolicyDefinition,
	_ResourceTypeName[2777:2803]:      PolicyRemediation,
	_ResourceTypeLowerName[2777:2803]: PolicyRemediation,
	_ResourceTypeName[2803:2832]:      PolicySetDefinition,
	_ResourceTypeLowerName[2803:2832]: PolicySetDefinition,
	_ResourceTypeName[2832:2849]:      KeyVault,
	_ResourceTypeLowerName[2832:2849]: KeyVault,
	_ResourceTypeName[2849:2880]:      KeyVaultAccessPolicy,
	_ResourceTypeLowerName[2849:2880]: KeyVaultAccessPolicy,
	_ResourceTypeName[2880:2908]:      ApplicationInsights,
	_ResourceTypeLowerName[2880:2908]: ApplicationInsights,
	_ResourceTypeName[2908:2944]:      ApplicationInsightsAPIKey,
	_ResourceTypeLowerName[2908:2944]: ApplicationInsightsAPIKey,
	_ResourceTypeName[2944:2987]:      ApplicationInsightsAnalyticsItem,
	_ResourceTypeLowerName[2944:2987]: ApplicationInsightsAnalyticsItem,
	_ResourceTypeName[2987:3018]:      LogAnalyticsWorkspace,
	_ResourceTypeLowerName[2987:3018]: LogAnalyticsWorkspace,
	_ResourceTypeName[3018:3054]:      LogAnalyticsLinkedService,
	_ResourceTypeLowerName[3018:3054]: LogAnalyticsLinkedService,
	_ResourceTypeName[3054:3114]:      LogAnalyticsDatasourceWindowsPerformanceCounter,
	_ResourceTypeLowerName[3054:3114]: LogAnalyticsDatasourceWindowsPerformanceCounter,
	_ResourceTypeName[3114:3160]:      LogAnalyticsDatasourceWindowsEvent,
	_ResourceTypeLowerName[3114:3160]: LogAnalyticsDatasourceWindowsEvent,
	_ResourceTypeName[3160:3188]:      MonitorActionGroup,
	_ResourceTypeLowerName[3160:3188]: MonitorActionGroup,
	_ResourceTypeName[3188:3222]:      MonitorActivityLogAlert,
	_ResourceTypeLowerName[3188:3222]: MonitorActivityLogAlert,
	_ResourceTypeName[3222:3255]:      MonitorAutoscaleSetting,
	_ResourceTypeLowerName[3222:3255]: MonitorAutoscaleSetting,
	_ResourceTypeName[3255:3282]:      MonitorLogProfile,
	_ResourceTypeLowerName[3255:3282]: MonitorLogProfile,
	_ResourceTypeName[3282:3310]:      MonitorMetricAlert,
	_ResourceTypeLowerName[3282:3310]: MonitorMetricAlert,
}

var _ResourceTypeNames = []string{
	_ResourceTypeName[0:22],
	_ResourceTypeName[22:45],
	_ResourceTypeName[45:76],
	_ResourceTypeName[76:105],
	_ResourceTypeName[105:138],
	_ResourceTypeName[138:179],
	_ResourceTypeName[179:218],
	_ResourceTypeName[218:261],
	_ResourceTypeName[261:284],
	_ResourceTypeName[284:308],
	_ResourceTypeName[308:328],
	_ResourceTypeName[328:341],
	_ResourceTypeName[341:355],
	_ResourceTypeName[355:380],
	_ResourceTypeName[380:410],
	_ResourceTypeName[410:437],
	_ResourceTypeName[437:471],
	_ResourceTypeName[471:507],
	_ResourceTypeName[507:523],
	_ResourceTypeName[523:552],
	_ResourceTypeName[552:571],
	_ResourceTypeName[571:594],
	_ResourceTypeName[594:623],
	_ResourceTypeName[623:640],
	_ResourceTypeName[640:664],
	_ResourceTypeName[664:677],
	_ResourceTypeName[677:696],
	_ResourceTypeName[696:727],
	_ResourceTypeName[727:769],
	_ResourceTypeName[769:800],
	_ResourceTypeName[800:839],
	_ResourceTypeName[839:858],
	_ResourceTypeName[858:892],
	_ResourceTypeName[892:922],
	_ResourceTypeName[922:944],
	_ResourceTypeName[944:975],
	_ResourceTypeName[975:1020],
	_ResourceTypeName[1020:1030],
	_ResourceTypeName[1030:1061],
	_ResourceTypeName[1061:1076],
	_ResourceTypeName[1076:1100],
	_ResourceTypeName[1100:1119],
	_ResourceTypeName[1119:1138],
	_ResourceTypeName[1138:1154],
	_ResourceTypeName[1154:1187],
	_ResourceTypeName[1187:1228],
	_ResourceTypeName[1228:1254],
	_ResourceTypeName[1254:1286],
	_ResourceTypeName[1286:1317],
	_ResourceTypeName[1317:1343],
	_ResourceTypeName[1343:1377],
	_ResourceTypeName[1377:1403],
	_ResourceTypeName[1403:1439],
	_ResourceTypeName[1439:1462],
	_ResourceTypeName[1462:1483],
	_ResourceTypeName[1483:1504],
	_ResourceTypeName[1504:1525],
	_ResourceTypeName[1525:1545],
	_ResourceTypeName[1545:1574],
	_ResourceTypeName[1574:1598],
	_ResourceTypeName[1598:1627],
	_ResourceTypeName[1627:1649],
	_ResourceTypeName[1649:1685],
	_ResourceTypeName[1685:1712],
	_ResourceTypeName[1712:1734],
	_ResourceTypeName[1734:1761],
	_ResourceTypeName[1761:1781],
	_ResourceTypeName[1781:1815],
	_ResourceTypeName[1815:1847],
	_ResourceTypeName[1847:1874],
	_ResourceTypeName[1874:1906],
	_ResourceTypeName[1906:1931],
	_ResourceTypeName[1931:1970],
	_ResourceTypeName[1970:1995],
	_ResourceTypeName[1995:2017],
	_ResourceTypeName[2017:2044],
	_ResourceTypeName[2044:2064],
	_ResourceTypeName[2064:2106],
	_ResourceTypeName[2106:2151],
	_ResourceTypeName[2151:2180],
	_ResourceTypeName[2180:2214],
	_ResourceTypeName[2214:2233],
	_ResourceTypeName[2233:2260],
	_ResourceTypeName[2260:2276],
	_ResourceTypeName[2276:2296],
	_ResourceTypeName[2296:2319],
	_ResourceTypeName[2319:2341],
	_ResourceTypeName[2341:2365],
	_ResourceTypeName[2365:2386],
	_ResourceTypeName[2386:2407],
	_ResourceTypeName[2407:2429],
	_ResourceTypeName[2429:2451],
	_ResourceTypeName[2451:2473],
	_ResourceTypeName[2473:2497],
	_ResourceTypeName[2497:2525],
	_ResourceTypeName[2525:2556],
	_ResourceTypeName[2556:2588],
	_ResourceTypeName[2588:2617],
	_ResourceTypeName[2617:2647],
	_ResourceTypeName[2647:2677],
	_ResourceTypeName[2677:2707],
	_ResourceTypeName[2707:2752],
	_ResourceTypeName[2752:2777],
	_ResourceTypeName[2777:2803],
	_ResourceTypeName[2803:2832],
	_ResourceTypeName[2832:2849],
	_ResourceTypeName[2849:2880],
	_ResourceTypeName[2880:2908],
	_ResourceTypeName[2908:2944],
	_ResourceTypeName[2944:2987],
	_ResourceTypeName[2987:3018],
	_ResourceTypeName[3018:3054],
	_ResourceTypeName[3054:3114],
	_ResourceTypeName[3114:3160],
	_ResourceTypeName[3160:3188],
	_ResourceTypeName[3188:3222],
	_ResourceTypeName[3222:3255],
	_ResourceTypeName[3255:3282],
	_ResourceTypeName[3282:3310],
}

// ResourceTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ResourceTypeString(s string) (ResourceType, error) {
	if val, ok := _ResourceTypeNameToValueMap[s]; ok {
		return val, nil
	}
	s = strings.ToLower(s)
	if val, ok := _ResourceTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ResourceType values", s)
}

// ResourceTypeValues returns all values of the enum
func ResourceTypeValues() []ResourceType {
	return _ResourceTypeValues
}

// ResourceTypeStrings returns a slice of all String values of the enum
func ResourceTypeStrings() []string {
	strs := make([]string, len(_ResourceTypeNames))
	copy(strs, _ResourceTypeNames)
	return strs
}

// IsAResourceType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ResourceType) IsAResourceType() bool {
	for _, v := range _ResourceTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
