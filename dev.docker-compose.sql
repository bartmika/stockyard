SELECT citus_set_coordinator_host('citus_coordinator', 5432);
SELECT citus_add_node('citus_worker1', 5432);
SELECT citus_add_node('citus_worker2', 5432);
SELECT rebalance_table_shards();
