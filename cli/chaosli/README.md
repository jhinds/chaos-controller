# Chaosli

The Chaos Controller CLI, `chaosli`, is meant to help create a more user friendly and digestible chaos-controller experience. The CLI pairs with your Disruption's definition, giving you information to give you a better sense of what is going on, and what will happen when you apply new disruptions. With features like explain and validation, it gives users, not only a better understanding, but a better experience with the controller.  


#### Table of Contents
---
- Validate
- Explain
- Create

#### Validate
---
Usage: `go run chaosli/main.go validate --path <path to disruption file>`

Description: Validates your disruption file (location defined by `--path`) to make sure everything is correctly formatted, and all options have valid configurations.

Example:

```
# Running Validate on a Configuration that has no Disruptions specified
# Note; We removed the network portion from the actual example for this example, normally
# network_delay.yaml in examples has a network disruption specified

$ go run chaosli/main.go validate --path=../examples/network_delay.yaml
Error: cannot apply an empty disruption - at least one of Network, DNS, DiskPressure, NodeFailure, CPUPressure fields is needed
```

#### Explain
---
Usage: `go run chaosli/main.go explain --path <path to disruption file>`

Description: Prints out a summary of the disruption (location defined by `--path`).

Example:

```
$ go run chaosli/main.go explain --path=../examples/network_delay.yaml
This Disruption...
=======================================================================================================================================
🧰 has the following metadata  ...
	ℹ️  has DryRun set to false meaning this disruption WILL run.
	ℹ️  will be run on the Pod level, so everything in this disruption is scoped at this level.
	ℹ️  has the following selectors which will be used to target pods
		🎯  app=demo-curl
	ℹ️  is going to target 1 pod(s) (either described as a percentage of total pods or actual number of them).
=======================================================================================================================================
💉 injects a network disruption ...
	💥 applies network failures on outgoing traffic.
		💣 applies a packet delay of 1000 ms.
			💣 applies a jitter of 1000 ms to the delay value to add randomness to the delay.
=======================================================================================================================================
```

#### Creation
---
Usage: `go run chaosli/main.go create --path <path to output generated disruption file>`

Description: User friendly input process that helps you create your disruptions from scratch answering simple questions.