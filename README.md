# PRE-Test

Write a function that validates that a []string of files contains all the
expected files, and no unexpected files.  

The expected files include cluster specific and node specific files. 

```
# Cluster specific files
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_configs",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/volume_configs",
	"/tmp/diagBundle-71832245/HttpClientFile803827980/2021-03-04T11-10-45.659243121Z/volume_configs",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/cluster_config",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/namespace_partitions",	
	
# Node specific files for node: 83662212-59d5-44b6-a381-65bc7f4f0e7d
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/attached_volumes",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_83662212-59d5-44b6-a381-65bc7f4f0e7d/cached_licence",
	...

# Node specific files for node: 43287213-54xf-57s5-bnd5-sd5f45saf2351
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/attached_volumes",
	"/tmp/diagBundle-76461881/HttpClientFile095427323/2021-03-04T11-10-45.769125854Z/node_43287213-54xf-57s5-bnd5-sd5f45saf2351/cached_licence",
```