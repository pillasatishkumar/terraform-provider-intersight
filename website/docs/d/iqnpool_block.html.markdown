---
subcategory: "iqnpool"
layout: "intersight"
page_title: "Intersight: intersight_iqnpool_block"
description: |-
  A block of contiguous IQNs that are part of a pool.
---

# Data Source: intersight_iqnpool_block
A block of contiguous IQNs that are part of a pool.
## Argument Reference
The following arguments can be used to get data of already created objects in Intersight appliance:
* `free_block_count`:(int) Free IDs that can be allocated in this block. 
* `moid`:(string) The unique identifier of this Managed Object instance. 
* `next_id_allocator`:(int) Moving counter to allocate the next identifier. 
