# KV-Sync

KV-Sync allows to sync key value pairs defined in Nomad job template's [meta stanza](https://www.nomadproject.io/docs/job-specification/meta) with Consul key vault. This continuously listen to Nomad's job events and update Consul key vault. Currently this listens to `JobRegistered` and `JobDeregistered` events.

### Quick start

Run the docker container:
```
$ isurusiri91/kvsync -n <NOMAD_HOST_URL> -c <CONSUL_HOST_URL>
```
