package gethwrappers

//go:generate go run ./generation/wrap.go automation VerifiableLoadUpkeep verifiable_load_upkeep_wrapper
//go:generate go run ./generation/wrap.go automation VerifiableLoadStreamsLookupUpkeep verifiable_load_streams_lookup_upkeep_wrapper
//go:generate go run ./generation/wrap.go automation StreamsLookupUpkeep streams_lookup_upkeep_wrapper
//go:generate go run ./generation/wrap.go automation StreamsLookupCompatibleInterface streams_lookup_compatible_interface
//go:generate go run ./generation/wrap.go automation AutomationConsumerBenchmark automation_consumer_benchmark
//go:generate go run ./generation/generate_automation/wrap.go AutomationRegistrar2_1 AutomationRegistrar automation_registrar_wrapper2_1
//go:generate go run ./generation/generate_automation/wrap.go KeeperRegistry2_1 KeeperRegistry keeper_registry_wrapper_2_1
//go:generate go run ./generation/generate_automation/wrap.go KeeperRegistryLogicA2_1 KeeperRegistryLogicA keeper_registry_logic_a_wrapper_2_1
//go:generate go run ./generation/generate_automation/wrap.go KeeperRegistryLogicB2_1 KeeperRegistryLogicB keeper_registry_logic_b_wrapper_2_1
//go:generate go run ./generation/wrap.go automation IKeeperRegistryMaster i_keeper_registry_master_wrapper_2_1
//go:generate go run ./generation/wrap.go automation IAutomationV21PlusCommon i_automation_v21_plus_common

//go:generate go run ./generation/wrap.go automation ILogAutomation i_log_automation
//go:generate go run ./generation/wrap.go automation AutomationForwarderLogic automation_forwarder_logic
//go:generate go run ./generation/wrap.go automation LogTriggeredStreamsLookup log_triggered_streams_lookup_wrapper
//go:generate go run ./generation/wrap.go automation DummyProtocol dummy_protocol_wrapper

//go:generate go run ./generation/wrap.go automation KeeperConsumerPerformance keeper_consumer_performance_wrapper
//go:generate go run ./generation/wrap.go automation PerformDataChecker perform_data_checker_wrapper
//go:generate go run ./generation/wrap.go automation UpkeepCounter upkeep_counter_wrapper
//go:generate go run ./generation/wrap.go automation UpkeepPerformCounterRestrictive upkeep_perform_counter_restrictive_wrapper
