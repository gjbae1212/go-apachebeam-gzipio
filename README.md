# go-apachebeam-gzipio

<p align="left">
<a href="https://hits.seeyoufarm.com"/><img src="https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https%3A%2F%2Fgithub.com%2Fgjbae1212%2Fgo-apachebeam-gzipio"/></a>
<a href="https://goreportcard.com/badge/github.com/gjbae1212/go-apachebeam-gzipio"><img src="https://goreportcard.com/badge/github.com/gjbae1212/go-apachebeam-gzipio" alt="Go Report Card"/></a>
<a href="/LICENSE"><img src="https://img.shields.io/badge/license-MIT-GREEN.svg" alt="license" /></a> 
</p>

## Overview
This project is to transform for reading and writing gzip files in apache beam.  
And it is using to official Go SDK on apache beam.
 
## Getting Started
```go
import (
   gzipio "github.com/gjbae1212/go-apachebeam-gzipio"
)

p, s := beam.NewPipelineWithRoot()
// Read files
gzipio.Read(s, "read-path")

// Write file
gzipio.Write(s, "write-path", pcollection)
```

## License
This project is licensed under the MIT License

