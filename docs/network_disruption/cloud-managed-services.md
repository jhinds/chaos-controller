# Network disruption: Specifying cloud managed services

## Why

Large cloud services providers are using wide IP ranges. Hostnames used to identify those services are resolving with some IPs of that range, and resolved IPs can change between each DNS request. Applying a network disruption using those hostnames only doesn’t work well since retrying the resolution of such hostname would return new IPs (not disrupted) and the disruption would be ineffective.

Available cloud providers are:
- AWS

### Cloud Provider Manager

The service will pull and parse the IP Ranges from the available cloud providers every x minutes/hours, defined in the chaos-controller configuration:

```
cloudProviders:
    pullInterval: "1d"
```

On the creation of the chaos pod, the chaos-controller will then use those ip ranges for the Network Disruption and transform it into a Host Network Disruption.

### Example


```
apiVersion: chaos.datadoghq.com/v1beta1
kind: Disruption
metadata:
  name: network-cloud
  namespace: chaos-demo
spec:
  level: pod
  selector:
    app: demo-cirl
  count: 1
  network:
    cloud:
      aws:
        - service: "S3"
          flow: "egress" # available are egress or ingress. Optional
          protocol: "tcp" # available are tcp or udp. Optional
    delay: 1000 # delay (in milliseconds) to add to outgoing packets, 10% of jitter will be added by default
    delayJitter: 5 # (optional) add X % (1-100) of delay as jitter to delay (+- X% ms to original delay), defaults to 10%
```

## AWS

Available services are:
```
 DYNAMODB, ROUTE53, ROUTE53_RESOLVER, EBS, CODEBUILD, API_GATEWAY, WORKSPACES_GATEWAYS, EC2_INSTANCE_CONNECT, CHIME_VOICECONNECTOR, GLOBALACCELERATOR, CHIME_MEETINGS, CLOUDFRONT_ORIGIN_FACING, AMAZON_APPFLOW, KINESIS_VIDEO_STREAMS, EC2, CLOUDFRONT, ROUTE53_HEALTHCHECKS_PUBLISHING, CLOUD9, ROUTE53_HEALTHCHECKS, S3, AMAZON_CONNECT
```

We do not support using the service "AMAZON" (from the ip ranges file) as it's a combination of all ip ranges from all services and more miscellaneous ips; the number of ip ranges being too much from this, it's not possible for us to filter all of them at once.

We are using the URL **https://ip-ranges.amazonaws.com/ip-ranges.json** to pull all the IP Ranges of AWS.