package main

// Write a function that validates that a []string of files contains all the
// expected files, and no unexpected files.
// Expected files for a given node are the nodeFiles.
// Expected files for a cluster are the clusterFiles
func verifyBundle(bundleFiles []string, nodeIDs []string) error {
	return nil
}

// nodeLevelFiles contains a slice of files that are expected to be found for
// each node_ID in the diagnostic bundle.
// For instance, given two nodes: node_A... and node_B...
// We would expect to find node_A/cached_licence AND node_B/cached_licence files.
var nodeFiles = []string{
	"attached_volumes",
	"cached_licence",
	"capacity",
	"crashes",
	"dataplane_config",
	"dataplane_crashes",
	"dmesg_output",
	"env_config",
	"local_time",
	"local_volumes",
	"log_latest",
	"lshw_output",
	"nfs_logs",
	"peer_health",
	"rotated_logs",
	"version",
	"metrics",
	"metrics/api_metrics.csv",
	"metrics/csi_backend_metrics.csv",
	"metrics/csi_controller_metrics.csv",
	"metrics/csi_identity_metrics.csv",
	"metrics/csi_node_metrics.csv",
	"metrics/dataplane_dfs_metrics.csv",
	"metrics/dataplane_director_metrics.csv",
	"metrics/dataplane_fs_metrics.csv",
	"metrics/dataplane_metrics.csv",
	"metrics/dataplane_rdb_metrics.csv",
	"metrics/dataplane_supervisor_metrics.csv",
	"metrics/state_tracker_metrics.csv",
	"metrics/store_metrics.csv",
}

// clusterLevelFiles contains a slice of files that are expected to be found in
// a diagnostic bundle. These files are only expected to be found once per
// bundle, regardless of the node count.
var clusterFiles = []string{
	"cluster_config",
	"namespace_partitions",
	"node_configs",
	"volume_configs",
}

// files contains a slice of walked file paths from a StorageOS diagnostic bundle.
var files = []string{
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/attached_volumes",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/cached_licence",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/capacity",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/crashes",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/dataplane_config",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/dataplane_crashes",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/dmesg_output",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/env_config",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/local_time",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/local_volumes",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/log_latest",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/lshw_output",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/metrics",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/metrics/api_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/metrics/csi_backend_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/metrics/csi_controller_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/metrics/csi_identity_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/metrics/csi_node_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/metrics/dataplane_dfs_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/metrics/dataplane_director_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/metrics/dataplane_fs_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/metrics/dataplane_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/metrics/dataplane_rdb_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/metrics/dataplane_supervisor_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/metrics/state_tracker_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/metrics/store_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/nfs_logs",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/peer_health",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/rotated_logs",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/version",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_configs",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/volume_configs",
	"/tmp/diagBundle-71832245/HttpClientFile803827980/2021-03-04T11-10-45.659243121Z/volume_configs",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/cluster_config",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/namespace_partitions",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/attached_volumes",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/cached_licence",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/capacity",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/crashes",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/dataplane_config",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/dataplane_crashes",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/dmesg_output",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/env_config",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/local_time",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/local_volumes",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/log_latest",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/lshw_output",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/metrics",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/metrics/api_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/metrics/csi_backend_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/metrics/csi_controller_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/metrics/csi_identity_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/metrics/csi_node_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/metrics/dataplane_dfs_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/metrics/dataplane_director_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/metrics/dataplane_fs_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/metrics/dataplane_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/metrics/dataplane_rdb_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/metrics/dataplane_supervisor_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/metrics/state_tracker_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/metrics/store_metrics.csv",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/nfs_logs",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/peer_health",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/rotated_logs",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/version",
}
