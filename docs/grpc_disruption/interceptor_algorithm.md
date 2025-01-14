# gRPC Disruption Interceptor

When the interceptor recognizes a query's endpoint as one which is actively getting disrupted, the interceptor generates a random integer from `0` to `100`, and consults a slice with length <= 100 to figure out what `alteration` to apply to a query response. This mapping is populated by the percentage odds a user configured for each alteration. Currently, we support two `alteration`s:

1. return a gRPC error code (such as `NotFound` or `PermissionDenied`)
2. return an empty response (`emptypb.Empty`)

You can see an example below of a mapping that does not define all 100% of possible requests below.

## gRPC Disruption - Algorithm Examples

### Multiple alterations with defined query percentages

The following is a complex example to illustrate the algorithm's behavior.

```
spec:
  grpc:
    endpoints:
      - endpoint: /chaosdogfood.ChaosDogfood/order
        override: "{}"
        queryPercent: 5
      - endpoint: /chaosdogfood.ChaosDogfood/order
        error: NOT_FOUND
        queryPercent: 5
      - endpoint: /chaosdogfood.ChaosDogfood/order
        error: PERMISSION_DENIED
        queryPercent: 15
```

For the above specs, the calculated slice would look something like:

```
[
    0  -> Override: {}
    1  -> Override: {}
    2  -> Override: {}
    3  -> Override: {}
    4  -> Override: {}
    5  -> Error: NOT_FOUND
    6  -> Error: NOT_FOUND
    7  -> Error: NOT_FOUND
    8  -> Error: NOT_FOUND
    9  -> Error: NOT_FOUND
    10 -> Error: PERMISSION_DENIED
    11 -> Error: PERMISSION_DENIED
    12 -> Error: PERMISSION_DENIED
    13 -> Error: PERMISSION_DENIED
    .. -> Error: ...
    22 -> Error: PERMISSION_DENIED
    23 -> Error: PERMISSION_DENIED
    24 -> Error: PERMISSION_DENIED
]
```

In this case, we would return an Override with empty results for 5% of queries, a `NOT_FOUND` error for 5% of queries, and return `PERMISSION_DENIED` for 15% of queries.

### Multiple alterations with some undefined query percentages

We may also be provided with a configuration where some set of `queryPercent`s (query percentages) are defined, but not all..

```
spec:
  grpc:
    endpoints:
      - endpoint: /chaosdogfood.ChaosDogfood/order
        override: "{}"
        queryPercent: 25
      - endpoint: /chaosdogfood.ChaosDogfood/order
        error: NOT_FOUND
      - endpoint: /chaosdogfood.ChaosDogfood/order
        error: PERMISSION_DENIED
```

As in the previous case, all alterations with a defined `queryPercent` are allocated upfront. The algorithm keeps track of alterations which do not yet have `queryPercent`s assigned, and splits the remaining (unconfigured) queries equally amongst these unassigned alterations.

```
[
    0   -> Override: {}
    1   -> Override: {}
    2   -> Override: {}
    3   -> Override: {}
    4   -> Override: {}
    5   -> Override: {}
    6   -> Override: {}
    ..  -> Override: ..
    22  -> Override: {}
    23  -> Override: {}
    24  -> Override: {}
    25  -> Error: NOT_FOUND
    26  -> Error: NOT_FOUND
    26  -> Error: NOT_FOUND
    ..  -> Error: ...
    61  -> Error: NOT_FOUND
    62  -> Error: NOT_FOUND
    63  -> Error: PERMISSION_DENIED
    64  -> Error: PERMISSION_DENIED
    65  -> Error: PERMISSION_DENIED
    ..  -> Error: ...
    99  -> Error: PERMISSION_DENIED
    100 -> Error: PERMISSION_DENIED
]
```

You cannot specify query percentages for a single endpoint which sum to over `100%`.

### Simpler case of undefined query percentages

You may have noted that the second example appears a tad complex. The intuition behind this design is to support the case where a user wants to quickly define a disruption which errors on all queries (replicating a bad roll out). For one error, the algorithm returns an error every time. For two errors, the algorithm returns half of the queries with the first error and half with the other.

```
spec:
  grpc:
    endpoints:
      - endpoint: /chaosdogfood.ChaosDogfood/order
        error: NOT_FOUND
      - endpoint: /chaosdogfood.ChaosDogfood/order
        error: PERMISSION_DENIED
```

Rather than constraining the user in how they mix and match this simple configuration style with the explicit `spec.gprc.endpoints[x].queryPercent` field, the current implementation would simply do its best to apply of the configurations.

```
[
    0   -> Override: {}
    1   -> Override: {}
    2   -> Override: {}
    3   -> Override: {}
    4   -> Override: {}
    5   -> Override: {}
    6   -> Override: {}
    ..  -> Override: ..
    22  -> Override: {}
    23  -> Override: {}
    24  -> Override: {}
    25  -> Error: NOT_FOUND
    26  -> Error: NOT_FOUND
    26  -> Error: NOT_FOUND
    ..  -> Error: ...
    61  -> Error: NOT_FOUND
    62  -> Error: NOT_FOUND
    63  -> Error: PERMISSION_DENIED
    64  -> Error: PERMISSION_DENIED
    65  -> Error: PERMISSION_DENIED
    ..  -> Error: ...
    99  -> Error: PERMISSION_DENIED
    100 -> Error: PERMISSION_DENIED
]
```

## Design Implications

### Setting 0 as query percentage

It does not make sense for a user to set `queryPercent: 0`, and if a user tries to do so, they will see the error applied to all unclaimed queries. This is because Kubebuilder sets omitted `int`s to `0`, but chaos-controller interprets an omitted `queryPercent` to mean "apply all".

### Many errors, but very few slots remaining

When an even split across remaining points is not possible. For example, if 9% of queries are unaccounted for, and there are 6 different errors to assign to the mapping, the `pctPerAlt` (describing the mapping each query should be assign) would be `9 / 6` which is `1` in integer division. The final mapping would look like:
```
{
	..  -> ...
	90  -> ERROR_X
	91  -> ERROR_X
	92  -> ERROR_1
	93  -> ERROR_2
	94  -> ERROR_3
	95  -> ERROR_4
	96  -> ERROR_5
	97  -> ERROR_6
	98  -> ERROR_6
	99  -> ERROR_6
	100 -> ERROR_6
}
```
Note that the final alteration (in this case `ERROR_6`, always covers the remaining `Percent`s up to and including 100. This can result in a very weird proportions where there are not a lot of `Percent`s left.

Although these outcomes are unintuitive and therefore not very user-friendly, these disruptions are typically used to return a maximum of two or three alterations at a time (to simulate an already degraded state), so we designed the algorithm such that those common cases are intuitive and only reject configurations where the implementation definitely goes against the users' intentions (such as alterations that are defined but never apply).