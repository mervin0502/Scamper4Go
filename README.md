# Scamper4Go

Scamper4Go is a pure-Golang parsing library for the Warts format. Warts is an extensible binary format produced by Scamper, an Internet measurement tool from CAIDA, to store measurement results such as traceroutes and pings.

## Example

```
    package main

    import (
        "mervin.me/Scamper4Go/extract/filter"
    )
	srcFiles := []string{
		"../data/warts/2017.gz",
		"../data/warts2/2018.warts",
	}
	filter.TopologyDump("../data/warts-2007.txt", nil, srcFiles...)
    
```