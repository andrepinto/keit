#!/bin/sh

/usr/local/bin/main &
/usr/local/bin/envoy -c /keit/envoy.json --service-cluster keit