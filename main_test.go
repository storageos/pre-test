package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifyBundle(t *testing.T) {
	bundleStructs := []struct {
		Name    string
		files   []string
		NodeIDs []string
		wantErr bool
	}{
		{
			Name: "All expected files",
			files: []string{
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/cluster_config",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/namespace_partitions",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/attached_volumes",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/cached_licence",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/capacity",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/crashes",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/dataplane_config",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/dataplane_crashes",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/dmesg_output",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/env_config",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/local_time",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/local_volumes",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/log_latest",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/lshw_output",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/api_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/csi_backend_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/csi_controller_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/csi_identity_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/csi_node_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/dataplane_dfs_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/dataplane_director_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/dataplane_fs_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/dataplane_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/dataplane_rdb_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/dataplane_supervisor_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/state_tracker_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/store_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/nfs_logs",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/peer_health",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/rotated_logs",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/version",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_configs",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/volume_configs",
			},
			NodeIDs: []string{"240c0131-6cae-4951-a3a7-33933eef2fed"},
			wantErr: false,
		},
		{
			Name: "All expected files - two nodes",
			files: []string{
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/cluster_config",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/namespace_partitions",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/attached_volumes",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/cached_licence",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/capacity",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/crashes",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/dataplane_config",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/dataplane_crashes",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/dmesg_output",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/env_config",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/local_time",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/local_volumes",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/log_latest",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/lshw_output",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/api_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/csi_backend_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/csi_controller_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/csi_identity_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/csi_node_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/dataplane_dfs_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/dataplane_director_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/dataplane_fs_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/dataplane_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/dataplane_rdb_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/dataplane_supervisor_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/state_tracker_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/store_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/nfs_logs",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/peer_health",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/rotated_logs",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/version",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/attached_volumes",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/cached_licence",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/capacity",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/crashes",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/dataplane_config",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/dataplane_crashes",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/dmesg_output",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/env_config",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/local_time",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/local_volumes",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/log_latest",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/lshw_output",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/metrics",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/metrics/api_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/metrics/csi_backend_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/metrics/csi_controller_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/metrics/csi_identity_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/metrics/csi_node_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/metrics/dataplane_dfs_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/metrics/dataplane_director_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/metrics/dataplane_fs_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/metrics/dataplane_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/metrics/dataplane_rdb_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/metrics/dataplane_supervisor_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/metrics/state_tracker_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/metrics/store_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/nfs_logs",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/peer_health",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/rotated_logs",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_9866dd68-1029-497e-a15e-db9e6d621423/version",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_configs",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/volume_configs",
			},
			NodeIDs: []string{
				"240c0131-6cae-4951-a3a7-33933eef2fed",
				"9866dd68-1029-497e-a15e-db9e6d621423",
			},
			wantErr: false,
		},
		{
			Name: "Unexpected cluster- and node-specific files",
			files: []string{
				"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z",
				"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d",
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
				"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/and_this",
				"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_configs",
				"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/volume_configs",
				"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/cluster_config",
				"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/namespace_partitions",
				"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/attached_volumes",
				"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/hey_what_is_this",
			},
			NodeIDs: []string{"83662212-59d5-44b6-a381-65bc7f4f0e7d"},
			wantErr: true,
		},
		{
			Name: "Missing cluster- and node-specific files",
			files: []string{
				"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z",
				"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d",
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
				// "/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/peer_health",
				"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/rotated_logs",
				"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/version",
				"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_configs",
				"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/volume_configs",
				// "/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/cluster_config",
				"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/namespace_partitions",
				"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/attached_volumes",
			},
			NodeIDs: []string{"83662212-59d5-44b6-a381-65bc7f4f0e7d"},
			wantErr: true,
		},
		{
			Name: "Unexpected Node Files",
			files: []string{
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/cluster_config",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/namespace_partitions",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/attached_volumes",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/cached_licence",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/capacity",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/crashes",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/dataplane_config",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/dataplane_crashes",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/dmesg_output",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/env_config",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/local_time",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/local_volumes",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/log_latest",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/lshw_output",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/api_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/csi_backend_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/csi_controller_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/csi_identity_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/csi_node_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/dataplane_dfs_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/dataplane_director_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/dataplane_fs_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/dataplane_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/dataplane_rdb_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/dataplane_supervisor_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/state_tracker_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/metrics/store_metrics.csv",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/nfs_logs",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/peer_health",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/rotated_logs",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-6cae-4951-a3a7-33933eef2fed/version",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_240c0131-9xk1-4159-k7z3-33933eef2fed/version",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/node_configs",
				"/tmp/diagBundle-90976411/HttpClientFile771489167/2021-09-03T15-56-35.195000555Z/volume_configs",
			},
			NodeIDs: []string{"240c0131-6cae-4951-a3a7-33933eef2fed"},
			wantErr: false,
		},
	}

	for _, tt := range bundleStructs {
		t.Run(tt.Name, func(st *testing.T) {
			err := verifyBundle(tt.files, tt.NodeIDs)
			if tt.wantErr {
				require.Error(st, err, "wanted error but got none")
				return
			}

			require.NoError(st, err, "wanted no error but got err")
		})

	}
}
