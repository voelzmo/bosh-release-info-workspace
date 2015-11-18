## Set GOPATH

Source environment file directly
```
$ source .envrc
```

Or install [direnv](http://direnv.net/) and have this done automatically from now on
```
$ direnv allow
```

## Build

```
$ ./bin/build
```

## Run

```
$ ./out/bosh-release-info help
NAME:
   bosh-release-info - lists info about a boshrelease

USAGE:
   ./out/bosh-release-info [global options] command [command options] [arguments...]

VERSION:
   0.1

COMMANDS:
   package-list	lists all packages in this release
   file-list	lists all files in all packages in this release
   help, h	Shows a list of commands or help for one command
```

### List packages in a release

```
$ ./out/bosh-release-info package-list ~/workspace/bosh/release/dev_releases/bosh/bosh-217+dev.1.tgz
Info for release: bosh-217+dev.1.tgz

Release name: bosh

Packages: director, genisoimage, health_monitor, libpq, mysql, nats, nginx, postgres, powerdns, redis, registry, ruby
```

### List files grouped by type
Specify the option `--by-type`

*Note: This lists a few meaningful files grouped by their type. Some files contained in the release may be left out of this listing.*

```
$ /out/bosh-release-info file-list ~/Downloads/cf-release-225.tgz --by-type

Info for release 'cf', version '225'

=== Files for type 'License' (87) ===
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/Godeps/_workspace/src/github.com/cloudfoundry/noaa/LICENSE
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/Godeps/_workspace/src/github.com/cloudfoundry-incubator/cf-test-helpers/LICENSE
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/Godeps/_workspace/src/github.com/gorilla/websocket/LICENSE
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/Godeps/_workspace/src/github.com/onsi/ginkgo/LICENSE
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/Godeps/_workspace/src/github.com/onsi/gomega/LICENSE
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/Godeps/_workspace/src/github.com/onsi/gomega/matchers/support/goraph/MIT.LICENSE
cloud_controller_ng:./cloud_controller_ng/vendor/cache/vcap-concurrency-2a5b0179642c/LICENSE
cloud_controller_ng:./cloud_controller_ng/vendor/cache/cf-uaa-lib-b1e11235dc6c/LICENSE.TXT
cloud_controller_ng:./cloud_controller_ng/vendor/cache/delayed_job_sequel-ae606dd8dd72/LICENSE
collector:./collector/vendor/cache/vcap-common-4ad156de3fcb/LICENSE
collector:./collector/vendor/cache/cf-message-bus-b48eeb4357bd/LICENSE.txt
dea_next:./dea_next/go/src/github.com/cloudfoundry/gosteno/LICENSE
dea_next:./dea_next/go/src/github.com/howeyc/fsnotify/LICENSE
dea_next:./dea_next/go/src/launchpad.net/gocheck/LICENSE
dea_next:./dea_next/go/src/launchpad.net/goyaml/LICENSE
dea_next:./dea_next/go/src/launchpad.net/goyaml/LICENSE.libyaml
dea_next:./dea_next/vendor/cache/warden-181b550918c8/LICENSE
dea_next:./dea_next/vendor/cache/grape-b59999dcbcd1/LICENSE
etcd:./github.com/coreos/etcd/LICENSE
etcd:./github.com/coreos/etcd/Godeps/_workspace/src/google.golang.org/grpc/LICENSE
etcd:./github.com/coreos/etcd/Godeps/_workspace/src/github.com/boltdb/bolt/LICENSE
etcd:./github.com/coreos/etcd/Godeps/_workspace/src/github.com/golang/glog/LICENSE
etcd:./github.com/coreos/etcd/Godeps/_workspace/src/github.com/bradfitz/http2/LICENSE
etcd:./github.com/coreos/etcd/Godeps/_workspace/src/github.com/bgentry/speakeasy/LICENSE_WINDOWS
etcd:./github.com/coreos/etcd/Godeps/_workspace/src/github.com/google/btree/LICENSE
etcd:./github.com/coreos/etcd/Godeps/_workspace/src/github.com/prometheus/procfs/LICENSE
etcd:./github.com/coreos/etcd/Godeps/_workspace/src/github.com/codegangsta/cli/LICENSE
etcd:./github.com/coreos/etcd/Godeps/_workspace/src/github.com/jonboulle/clockwork/LICENSE
etcd:./github.com/coreos/etcd/Godeps/_workspace/src/golang.org/x/oauth2/LICENSE
hm9000:./hm9000/src/bitbucket.org/kardianos/osext/LICENSE
hm9000:./hm9000/src/code.google.com/p/snappy-go/LICENSE
hm9000:./hm9000/src/code.google.com/p/go.net/LICENSE
hm9000:./hm9000/src/code.google.com/p/go.tools/LICENSE
hm9000:./hm9000/src/code.google.com/p/goprotobuf/LICENSE
hm9000:./hm9000/src/code.google.com/p/gogoprotobuf/LICENSE
hm9000:./hm9000/src/github.com/gorilla/websocket/LICENSE
hm9000:./hm9000/src/github.com/cloudfoundry-incubator/cf-debug-server/LICENSE
hm9000:./hm9000/src/github.com/coreos/go-log/LICENSE
hm9000:./hm9000/src/github.com/coreos/go-etcd/LICENSE
hm9000:./hm9000/src/github.com/coreos/etcd/LICENSE
hm9000:./hm9000/src/github.com/coreos/go-systemd/LICENSE
hm9000:./hm9000/src/github.com/cloudfoundry/yagnats/LICENSE
hm9000:./hm9000/src/github.com/cloudfoundry/go_cfmessagebus/LICENSE
hm9000:./hm9000/src/github.com/cloudfoundry/loggregatorlib/LICENSE
hm9000:./hm9000/src/github.com/cloudfoundry/gunk/LICENSE
hm9000:./hm9000/src/github.com/cloudfoundry/storeadapter/LICENSE
hm9000:./hm9000/src/github.com/cloudfoundry/dropsonde/LICENSE
hm9000:./hm9000/src/github.com/cloudfoundry/dropsonde/dropsonde-protocol/LICENSE
hm9000:./hm9000/src/github.com/cloudfoundry/gosteno/LICENSE
hm9000:./hm9000/src/github.com/cloudfoundry/hm9000/LICENSE
hm9000:./hm9000/src/github.com/goji/httpauth/LICENSE
hm9000:./hm9000/src/github.com/ccding/go-config-reader/LICENSE
hm9000:./hm9000/src/github.com/ccding/go-logging/LICENSE
hm9000:./hm9000/src/github.com/pivotal-golang/lager/LICENSE
hm9000:./hm9000/src/github.com/onsi/ginkgo/LICENSE
hm9000:./hm9000/src/github.com/onsi/gomega/LICENSE
hm9000:./hm9000/src/github.com/onsi/gomega/matchers/support/goraph/MIT.LICENSE
hm9000:./hm9000/src/github.com/tedsuo/rata/LICENSE
hm9000:./hm9000/src/github.com/tedsuo/ifrit/LICENSE
hm9000:./hm9000/src/github.com/codegangsta/cli/LICENSE
hm9000:./hm9000/src/github.com/samuel/go-zookeeper/LICENSE
hm9000:./hm9000/src/github.com/apcera/nats/LICENSE
hm9000:./hm9000/src/tux21b.org/v1/gocql/LICENSE
nginx_newrelic_plugin:./nginx/newrelic_nginx_agent/LICENSE
routing-api:./github.com/cloudfoundry-incubator/routing-api/Godeps/_workspace/src/github.com/cloudfoundry/dropsonde/LICENSE
routing-api:./github.com/cloudfoundry-incubator/routing-api/Godeps/_workspace/src/github.com/cloudfoundry/gosteno/LICENSE
routing-api:./github.com/cloudfoundry-incubator/routing-api/Godeps/_workspace/src/github.com/cloudfoundry/storeadapter/LICENSE
routing-api:./github.com/cloudfoundry-incubator/routing-api/Godeps/_workspace/src/github.com/cloudfoundry-incubator/candiedyaml/LICENSE
routing-api:./github.com/cloudfoundry-incubator/routing-api/Godeps/_workspace/src/github.com/cloudfoundry-incubator/candiedyaml/libyaml-LICENSE
routing-api:./github.com/cloudfoundry-incubator/routing-api/Godeps/_workspace/src/github.com/cloudfoundry-incubator/cf-debug-server/LICENSE
routing-api:./github.com/cloudfoundry-incubator/routing-api/Godeps/_workspace/src/github.com/cloudfoundry-incubator/cf-lager/LICENSE
routing-api:./github.com/cloudfoundry-incubator/routing-api/Godeps/_workspace/src/github.com/cloudfoundry-incubator/trace-logger/LICENSE
routing-api:./github.com/cloudfoundry-incubator/routing-api/Godeps/_workspace/src/github.com/dgrijalva/jwt-go/LICENSE
routing-api:./github.com/cloudfoundry-incubator/routing-api/Godeps/_workspace/src/github.com/onsi/ginkgo/LICENSE
routing-api:./github.com/cloudfoundry-incubator/routing-api/Godeps/_workspace/src/github.com/onsi/gomega/LICENSE
routing-api:./github.com/cloudfoundry-incubator/routing-api/Godeps/_workspace/src/github.com/onsi/gomega/matchers/support/goraph/MIT.LICENSE
routing-api:./github.com/cloudfoundry-incubator/routing-api/Godeps/_workspace/src/github.com/pivotal-golang/clock/LICENSE
routing-api:./github.com/cloudfoundry-incubator/routing-api/Godeps/_workspace/src/github.com/pivotal-golang/lager/LICENSE
routing-api:./github.com/cloudfoundry-incubator/routing-api/Godeps/_workspace/src/github.com/tedsuo/ifrit/LICENSE
routing-api:./github.com/cloudfoundry-incubator/routing-api/Godeps/_workspace/src/github.com/tedsuo/rata/LICENSE
routing-api:./github.com/cloudfoundry-incubator/routing-api/LICENSE.txt
smoke-tests:./smoke-tests/Godeps/_workspace/src/github.com/onsi/gomega/matchers/support/goraph/MIT.LICENSE
smoke-tests:./smoke-tests/Godeps/_workspace/src/github.com/onsi/gomega/LICENSE
smoke-tests:./smoke-tests/Godeps/_workspace/src/github.com/onsi/ginkgo/LICENSE
warden:./warden/vendor/cache/common-6554ce1c83b7/LICENSE.TXT
warden:./warden/vendor/cache/membrane-b33b45c7b8a8/LICENSE
warden:./warden/vendor/cache/steno-e71a658f05b7/LICENSE

=== Files for type 'Rubygem' (273) ===
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/dora/vendor/cache/diff-lcs-1.2.5.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/dora/vendor/cache/json-1.8.1.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/dora/vendor/cache/rack-1.5.1.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/dora/vendor/cache/rack-protection-1.3.2.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/dora/vendor/cache/rack-test-0.6.2.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/dora/vendor/cache/rspec-2.14.1.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/dora/vendor/cache/rspec-core-2.14.7.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/dora/vendor/cache/rspec-expectations-2.14.4.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/dora/vendor/cache/rspec-mocks-2.14.4.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/dora/vendor/cache/sinatra-1.3.4.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/dora/vendor/cache/tilt-1.3.3.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/fuse-mount/vendor/cache/rack-1.5.2.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/fuse-mount/vendor/cache/rack-protection-1.5.3.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/fuse-mount/vendor/cache/sinatra-1.4.5.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/fuse-mount/vendor/cache/tilt-1.4.1.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/hello-routing/vendor/cache/daemons-1.2.3.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/hello-routing/vendor/cache/eventmachine-1.0.7.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/hello-routing/vendor/cache/rack-1.5.2.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/hello-routing/vendor/cache/thin-1.6.3.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/hello-world/vendor/cache/rack-1.5.2.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/loggregator-load-generator/vendor/cache/rack-1.5.2.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/loggregator-load-generator/vendor/cache/rack-protection-1.5.0.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/loggregator-load-generator/vendor/cache/sinatra-1.4.3.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/loggregator-load-generator/vendor/cache/tilt-1.4.1.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/ruby_simple/vendor/cache/rack-1.5.2.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/ruby_simple/vendor/cache/rack-protection-1.5.0.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/ruby_simple/vendor/cache/sinatra-1.4.3.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/ruby_simple/vendor/cache/tilt-1.4.1.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/service_broker/vendor/cache/colorize-0.7.5.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/service_broker/vendor/cache/diff-lcs-1.2.5.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/service_broker/vendor/cache/json-1.8.1.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/service_broker/vendor/cache/rack-1.5.2.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/service_broker/vendor/cache/rack-protection-1.5.2.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/service_broker/vendor/cache/rack-test-0.6.2.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/service_broker/vendor/cache/rspec-2.14.1.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/service_broker/vendor/cache/rspec-core-2.14.7.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/service_broker/vendor/cache/rspec-expectations-2.14.4.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/service_broker/vendor/cache/rspec-mocks-2.14.4.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/service_broker/vendor/cache/sinatra-1.4.4.gem
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/service_broker/vendor/cache/tilt-1.4.1.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/rake-10.4.2.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/CFPropertyList-2.3.1.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/i18n-0.7.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/json-1.8.3.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/minitest-5.8.2.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/thread_safe-0.3.5.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/tzinfo-1.2.2.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/activesupport-4.1.13.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/builder-3.2.2.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/erubis-2.7.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/actionview-4.1.13.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/rack-1.6.4.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/rack-test-0.6.2.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/actionpack-4.1.13.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/activemodel-4.1.13.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/addressable-2.3.6.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/allowy-1.0.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/awesome_print-1.2.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/backports-3.6.4.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/beefcake-1.0.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/columnize-0.9.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/byebug-4.0.5.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/eventmachine-1.0.8.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/daemons-1.2.3.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/json_pure-1.8.2.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/thin-1.6.3.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/nats-0.5.1.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/yajl-ruby-1.2.1.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/cf-message-bus-0.3.4.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/multi_json-1.11.2.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/clockwork-1.1.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/cloudfront-signer-3.0.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/coderay-1.1.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/cookiejar-0.3.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/delayed_job-4.0.4.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/sequel-4.21.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/em-socksify-0.3.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/http_parser.rb-0.6.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/em-http-request-1.1.2.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/excon-0.45.4.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fission-0.5.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/msgpack-0.5.9.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fluent-logger-0.4.9.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/formatador-0.2.5.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/mime-types-2.6.1.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/net-ssh-2.9.2.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/net-scp-1.2.1.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-core-1.32.1.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/mini_portile-0.6.2.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/nokogiri-1.6.6.2.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-xml-0.1.2.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-atmos-0.1.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-json-1.0.2.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/ipaddress-0.8.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-aws-0.7.6.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/inflecto-0.0.2.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-brightbox-0.9.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-dynect-0.0.2.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-ecloud-0.3.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-google-0.0.7.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-local-0.2.1.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-powerdns-0.1.1.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-profitbricks-0.0.5.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-radosgw-0.0.4.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-riakcs-0.1.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-sakuracloud-1.0.1.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-serverlove-0.1.2.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-softlayer-0.4.7.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-storm_on_demand-0.1.1.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-terremark-0.1.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-vmfusion-0.1.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-voxel-0.1.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/fog-1.34.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/httpclient-2.6.0.1.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/loggregator_emitter-5.0.1.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/membrane-1.1.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/method_source-0.8.2.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/multipart-post-2.0.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/mysql2-0.3.20.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/netaddr-1.5.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/newrelic_rpm-3.12.0.288.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/pg-0.16.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/posix-spawn-0.3.9.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/slop-3.6.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/pry-0.10.1.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/pry-byebug-3.1.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/rack-protection-1.5.3.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/thor-0.19.1.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/railties-4.1.13.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/rfc822-0.1.5.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/tilt-1.4.1.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/sinatra-1.4.6.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/sinatra-contrib-1.4.2.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/squash_ruby-2.0.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/statsd-ruby-1.2.1.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/steno-1.2.4.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/unf_ext-0.0.6.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/unf-0.1.4.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/uuidtools-2.1.5.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/vmstat-2.1.0.gem
cloud_controller_ng:./cloud_controller_ng/vendor/cache/vcap_common-4.0.2.gem
collector:./collector/vendor/cache/uuidtools-2.1.3.gem
collector:./collector/vendor/cache/daemons-1.1.9.gem
collector:./collector/vendor/cache/posix-spawn-0.3.8.gem
collector:./collector/vendor/cache/squash_ruby-1.4.0.gem
collector:./collector/vendor/cache/multipart-post-2.0.0.gem
collector:./collector/vendor/cache/nats-0.5.0.beta.16.gem
collector:./collector/vendor/cache/hashie-1.2.0.gem
collector:./collector/vendor/cache/eventmachine-1.0.3.gem
collector:./collector/vendor/cache/yajl-ruby-1.2.1.gem
collector:./collector/vendor/cache/rack-1.5.2.gem
collector:./collector/vendor/cache/addressable-2.3.5.gem
collector:./collector/vendor/cache/aws-sdk-resources-2.0.39.gem
collector:./collector/vendor/cache/vmstat-2.1.0.gem
collector:./collector/vendor/cache/em-http-request-1.1.2.gem
collector:./collector/vendor/cache/grape-0.2.1.1.gem
collector:./collector/vendor/cache/rack-mount-0.8.3.gem
collector:./collector/vendor/cache/jmespath-1.0.2.gem
collector:./collector/vendor/cache/builder-3.2.2.gem
collector:./collector/vendor/cache/http_parser.rb-0.6.0.gem
collector:./collector/vendor/cache/json_pure-1.8.1.gem
collector:./collector/vendor/cache/membrane-0.0.2.gem
collector:./collector/vendor/cache/fluent-logger-0.4.6.gem
collector:./collector/vendor/cache/msgpack-0.5.5.gem
collector:./collector/vendor/cache/aws-sdk-core-2.0.39.gem
collector:./collector/vendor/cache/httpclient-2.3.4.1.gem
collector:./collector/vendor/cache/em-socksify-0.3.0.gem
collector:./collector/vendor/cache/multi_xml-0.5.4.gem
collector:./collector/vendor/cache/multi_json-1.11.0.gem
collector:./collector/vendor/cache/httparty-0.11.0.gem
collector:./collector/vendor/cache/json-1.7.7.gem
collector:./collector/vendor/cache/cookiejar-0.3.0.gem
collector:./collector/vendor/cache/mime-types-2.1.gem
collector:./collector/vendor/cache/aws-sdk-2.0.39.gem
collector:./collector/vendor/cache/thin-1.6.3.gem
collector:./collector/vendor/cache/vcap-concurrency-0.1.0.gem
collector:./collector/vendor/cache/steno-1.1.0.gem
dea_next:./dea_next/vendor/cache/rake-10.0.3.gem
dea_next:./dea_next/vendor/cache/i18n-0.7.0.gem
dea_next:./dea_next/vendor/cache/json-1.8.1.gem
dea_next:./dea_next/vendor/cache/minitest-5.5.1.gem
dea_next:./dea_next/vendor/cache/thread_safe-0.3.4.gem
dea_next:./dea_next/vendor/cache/tzinfo-1.2.2.gem
dea_next:./dea_next/vendor/cache/activesupport-4.2.0.gem
dea_next:./dea_next/vendor/cache/addressable-2.3.5.gem
dea_next:./dea_next/vendor/cache/descendants_tracker-0.0.4.gem
dea_next:./dea_next/vendor/cache/ice_nine-0.11.1.gem
dea_next:./dea_next/vendor/cache/axiom-types-0.1.1.gem
dea_next:./dea_next/vendor/cache/beefcake-1.0.0.gem
dea_next:./dea_next/vendor/cache/builder-3.2.2.gem
dea_next:./dea_next/vendor/cache/coercible-1.0.0.gem
dea_next:./dea_next/vendor/cache/cookiejar-0.3.0.gem
dea_next:./dea_next/vendor/cache/daemons-1.2.3.gem
dea_next:./dea_next/vendor/cache/eventmachine-1.0.8.gem
dea_next:./dea_next/vendor/cache/em-socksify-0.2.1.gem
dea_next:./dea_next/vendor/cache/http_parser.rb-0.5.3.gem
dea_next:./dea_next/vendor/cache/em-http-request-1.0.3.gem
dea_next:./dea_next/vendor/cache/em-synchrony-1.0.3.gem
dea_next:./dea_next/vendor/cache/equalizer-0.0.9.gem
dea_next:./dea_next/vendor/cache/ffi-1.0.11.gem
dea_next:./dea_next/vendor/cache/msgpack-0.4.7.gem
dea_next:./dea_next/vendor/cache/yajl-ruby-1.2.1.gem
dea_next:./dea_next/vendor/cache/fluent-logger-0.4.5.gem
dea_next:./dea_next/vendor/cache/hashie-3.4.0.gem
dea_next:./dea_next/vendor/cache/multi_json-1.10.1.gem
dea_next:./dea_next/vendor/cache/multi_xml-0.5.5.gem
dea_next:./dea_next/vendor/cache/rack-1.6.4.gem
dea_next:./dea_next/vendor/cache/rack-accept-0.4.5.gem
dea_next:./dea_next/vendor/cache/rack-mount-0.8.3.gem
dea_next:./dea_next/vendor/cache/virtus-1.0.4.gem
dea_next:./dea_next/vendor/cache/httpclient-2.3.4.1.gem
dea_next:./dea_next/vendor/cache/json_pure-1.8.2.gem
dea_next:./dea_next/vendor/cache/loggregator_emitter-5.0.1.gem
dea_next:./dea_next/vendor/cache/membrane-1.1.0.gem
dea_next:./dea_next/vendor/cache/mime-types-2.1.gem
dea_next:./dea_next/vendor/cache/mini_portile-0.6.0.gem
dea_next:./dea_next/vendor/cache/multipart-post-2.0.0.gem
dea_next:./dea_next/vendor/cache/thin-1.6.3.gem
dea_next:./dea_next/vendor/cache/nats-0.5.1.gem
dea_next:./dea_next/vendor/cache/nokogiri-1.6.2.1.gem
dea_next:./dea_next/vendor/cache/posix-spawn-0.3.8.gem
dea_next:./dea_next/vendor/cache/squash_ruby-1.4.0.gem
dea_next:./dea_next/vendor/cache/steno-1.1.0.gem
dea_next:./dea_next/vendor/cache/sys-filesystem-1.0.0.gem
dea_next:./dea_next/vendor/cache/uuidtools-2.1.3.gem
dea_next:./dea_next/vendor/cache/vmstat-2.0.0.gem
dea_next:./dea_next/vendor/cache/vcap_common-4.0.3.gem
nats:./nats/vendor/cache/daemons-1.1.9.gem
nats:./nats/vendor/cache/eventmachine-1.0.3.gem
nats:./nats/vendor/cache/yajl-ruby-1.2.1.gem
nats:./nats/vendor/cache/thin-1.6.2.gem
nats:./nats/vendor/cache/rack-1.5.2.gem
nats:./nats/vendor/cache/nats-0.5.0.beta.14.gem
nats:./nats/vendor/cache/json_pure-1.8.1.gem
nginx_newrelic_plugin:./nginx/newrelic_nginx_agent/vendor/cache/newrelic_plugin-1.3.1.gem
nginx_newrelic_plugin:./nginx/newrelic_nginx_agent/vendor/cache/json-1.8.3.gem
nginx_newrelic_plugin:./nginx/newrelic_nginx_agent/vendor/cache/daemons-1.2.3.gem
route_registrar:./route_registrar/vendor/cache/eventmachine-1.0.7.gem
route_registrar:./route_registrar/vendor/cache/daemons-1.2.3.gem
route_registrar:./route_registrar/vendor/cache/json_pure-1.8.2.gem
route_registrar:./route_registrar/vendor/cache/rack-1.6.4.gem
route_registrar:./route_registrar/vendor/cache/thin-1.6.3.gem
route_registrar:./route_registrar/vendor/cache/nats-0.5.1.gem
route_registrar:./route_registrar/vendor/cache/vcap-concurrency-0.1.0.gem
route_registrar:./route_registrar/vendor/cache/yajl-ruby-1.2.1.gem
route_registrar:./route_registrar/vendor/cache/cf-message-bus-0.3.4.gem
route_registrar:./route_registrar/vendor/cache/msgpack-0.5.12.gem
route_registrar:./route_registrar/vendor/cache/fluent-logger-0.5.0.gem
route_registrar:./route_registrar/vendor/cache/steno-1.2.4.gem
ruby-2.1.7:./ruby-2.1.7/bundler-1.9.4.gem
ruby-2.2.3:./ruby-2.2.3/bundler-1.10.6.gem
smoke-tests:./smoke-tests/assets/ruby_simple/vendor/cache/rack-protection-1.5.0.gem
smoke-tests:./smoke-tests/assets/ruby_simple/vendor/cache/tilt-1.4.1.gem
smoke-tests:./smoke-tests/assets/ruby_simple/vendor/cache/rack-1.5.2.gem
smoke-tests:./smoke-tests/assets/ruby_simple/vendor/cache/sinatra-1.4.3.gem
warden:./warden/vendor/cache/rake-0.9.2.2.gem
warden:./warden/vendor/cache/beefcake-1.0.0.gem
warden:./warden/vendor/cache/diff-lcs-1.1.3.gem
warden:./warden/vendor/cache/eventmachine-1.0.3.gem
warden:./warden/vendor/cache/posix-spawn-0.3.6.gem
warden:./warden/vendor/cache/hashie-1.2.0.gem
warden:./warden/vendor/cache/multi_json-1.3.6.gem
warden:./warden/vendor/cache/multi_xml-0.5.1.gem
warden:./warden/vendor/cache/rack-1.4.1.gem
warden:./warden/vendor/cache/rack-mount-0.8.3.gem
warden:./warden/vendor/cache/grape-0.2.1.gem
warden:./warden/vendor/cache/pidfile-0.3.0.gem
warden:./warden/vendor/cache/rspec-core-2.11.0.gem
warden:./warden/vendor/cache/rspec-expectations-2.11.1.gem
warden:./warden/vendor/cache/rspec-mocks-2.11.1.gem
warden:./warden/vendor/cache/rspec-2.11.0.gem
warden:./warden/vendor/cache/yajl-ruby-1.2.1.gem
warden:./warden/vendor/cache/common-6554ce1c83b7/vcap_logging/pkg/vcap_logging-1.0.0.gem

=== Files for type 'Archive' (39) ===
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/dora.zip
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/fuse-mount/httpfs2-0.1.5.tar.gz
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/logging-route-service.zip
acceptance-tests:./github.com/cloudfoundry/cf-acceptance-tests/assets/security_group_buildpack.zip
buildpack_binary:./binary-buildpack/binary_buildpack-cached-v1.0.1.zip
buildpack_go:./go-buildpack/go_buildpack-cached-v1.6.3.zip
buildpack_java:./java-buildpack/java-buildpack-v3.3.1.zip
buildpack_java_offline:./java-buildpack/java-buildpack-offline-v3.3.1.zip
buildpack_nodejs:./nodejs-buildpack/nodejs_buildpack-cached-v1.5.2.zip
buildpack_php:./php-buildpack/php_buildpack-cached-v4.2.1.zip
buildpack_python:./python-buildpack/python_buildpack-cached-v1.5.1.zip
buildpack_ruby:./ruby-buildpack/ruby_buildpack-cached-v1.6.8.zip
buildpack_staticfile:./staticfile-buildpack/staticfile_buildpack-cached-v1.2.2.zip
cli:./cli/cf-linux-amd64.tgz
consul:./consul/consul-template_0.9.0_linux_amd64.tar.gz
consul:./consul/consul_0.5.2_linux_amd64.zip
etcd:./github.com/coreos/etcd/Godeps/_workspace/src/github.com/prometheus/client_golang/text/testdata/protobuf.gz
etcd:./github.com/coreos/etcd/Godeps/_workspace/src/github.com/prometheus/client_golang/text/testdata/text.gz
golang1.4:./golang/go1.4.3.linux-amd64.tar.gz
haproxy:./haproxy/pcre-8.37.tar.gz
haproxy:./haproxy/haproxy-1.5.14.tar.gz
libpq:./postgres/postgresql-9.0.3.tar.gz
mysqlclient-5.5:./mysqlclient-5.5/libmysqlclient-dev_5.5.44.tar.gz
nginx:./nginx/headers-more-v0.25.tgz
nginx:./nginx/nginx-upload-module-2.2.tar.gz
nginx:./nginx/nginx-1.8.0.tar.gz
nginx:./nginx/pcre-8.37.tar.gz
nginx_newrelic_plugin:./nginx/newrelic_nginx_agent.tar.gz
postgres-9.4.2:./postgres/postgresql-9.4.2.tar.gz
postgres:./postgres/postgres-9.0.3-1.i386.tar.gz
postgres:./postgres/postgres-9.0.3-1.amd64.tar.gz
rootfs_cflinuxfs2:./rootfs/cflinuxfs2-1.17.0.tar.gz
rtr:./rtr/rtr-linux-amd64.tgz
ruby-2.1.7:./ruby-2.1.7/ruby-2.1.7.tar.gz
ruby-2.1.7:./ruby-2.1.7/yaml-0.1.6.tar.gz
ruby-2.2.3:./ruby-2.2.3/ruby-2.2.3.tar.gz
ruby-2.2.3:./ruby-2.2.3/rubygems-2.4.8.tgz
uaa:./openjdk/openjdk-1.8.0_60-x86_64-trusty-jre.tar.gz
uaa:./uaa/apache-tomcat-7.0.61.tar.gz


```

### List all files in all packages in a release
Just list all files contained in all packages.

```
$ ./out/bosh-release-info file-list ~/workspace/bosh/release/dev_releases/bosh/bosh-217+dev.1.tgz

Info for release: bosh-217+dev.1.tgz

Release name: bosh

Files for package 'director': ./
./bosh/
./packaging
./bosh/bosh-director/
./bosh/REVISION
./bosh/bosh-director/vendor/
./bosh/bosh-director/vendor/cache/
./bosh/bosh-director/vendor/cache/aws-sdk-1.60.2.gem
./bosh/bosh-director/vendor/cache/aws-sdk-v1-1.60.2.gem
./bosh/bosh-director/vendor/cache/bcrypt-ruby-3.0.1.gem
./bosh/bosh-director/vendor/cache/blobstore_client-1.0000.0.gem
./bosh/bosh-director/vendor/cache/bosh-core-1.0000.0.gem
./bosh/bosh-director/vendor/cache/bosh-director-1.0000.0.gem
./bosh/bosh-director/vendor/cache/bosh-director-core-1.0000.0.gem
./bosh/bosh-director/vendor/cache/bosh-registry-1.0000.0.gem
./bosh/bosh-director/vendor/cache/bosh-template-1.0000.0.gem
./bosh/bosh-director/vendor/cache/bosh_aws_cpi-2.1.0.gem
./bosh/bosh-director/vendor/cache/bosh_common-1.0000.0.gem
./bosh/bosh-director/vendor/cache/bosh_cpi-1.0000.0.gem
./bosh/bosh-director/vendor/cache/bosh_openstack_cpi-2.1.0.gem
./bosh/bosh-director/vendor/cache/bosh_vcloud_cpi-0.11.0.gem
./bosh/bosh-director/vendor/cache/bosh_vsphere_cpi-2.1.0.gem
./bosh/bosh-director/vendor/cache/builder-3.1.4.gem
./bosh/bosh-director/vendor/cache/cf-uaa-lib-3.2.1.gem
./bosh/bosh-director/vendor/cache/CFPropertyList-2.3.2.gem
./bosh/bosh-director/vendor/cache/daemons-1.2.3.gem
./bosh/bosh-director/vendor/cache/eventmachine-1.0.3.gem
./bosh/bosh-director/vendor/cache/excon-0.45.4.gem
./bosh/bosh-director/vendor/cache/ffi-1.9.10.gem
./bosh/bosh-director/vendor/cache/fission-0.5.0.gem
./bosh/bosh-director/vendor/cache/fog-1.34.0.gem
./bosh/bosh-director/vendor/cache/fog-atmos-0.1.0.gem
./bosh/bosh-director/vendor/cache/fog-aws-0.7.6.gem
./bosh/bosh-director/vendor/cache/fog-brightbox-0.9.0.gem
./bosh/bosh-director/vendor/cache/fog-core-1.32.1.gem
./bosh/bosh-director/vendor/cache/fog-dynect-0.0.2.gem
./bosh/bosh-director/vendor/cache/fog-ecloud-0.3.0.gem
./bosh/bosh-director/vendor/cache/fog-google-0.1.0.gem
./bosh/bosh-director/vendor/cache/fog-json-1.0.2.gem
./bosh/bosh-director/vendor/cache/fog-local-0.2.1.gem
./bosh/bosh-director/vendor/cache/fog-powerdns-0.1.1.gem
./bosh/bosh-director/vendor/cache/fog-profitbricks-0.0.5.gem
./bosh/bosh-director/vendor/cache/fog-radosgw-0.0.4.gem
./bosh/bosh-director/vendor/cache/fog-riakcs-0.1.0.gem
./bosh/bosh-director/vendor/cache/fog-sakuracloud-1.3.3.gem
./bosh/bosh-director/vendor/cache/fog-serverlove-0.1.2.gem
./bosh/bosh-director/vendor/cache/fog-softlayer-0.4.7.gem
./bosh/bosh-director/vendor/cache/fog-storm_on_demand-0.1.1.gem
./bosh/bosh-director/vendor/cache/fog-terremark-0.1.0.gem
./bosh/bosh-director/vendor/cache/fog-vmfusion-0.1.0.gem
./bosh/bosh-director/vendor/cache/fog-voxel-0.1.0.gem
./bosh/bosh-director/vendor/cache/fog-xml-0.1.2.gem
./bosh/bosh-director/vendor/cache/formatador-0.2.5.gem
./bosh/bosh-director/vendor/cache/gibberish-1.4.0.gem
./bosh/bosh-director/vendor/cache/httpclient-2.4.0.gem
./bosh/bosh-director/vendor/cache/inflecto-0.0.2.gem
./bosh/bosh-director/vendor/cache/ipaddress-0.8.0.gem
./bosh/bosh-director/vendor/cache/json-1.8.3.gem
./bosh/bosh-director/vendor/cache/json_pure-1.8.1.gem
./bosh/bosh-director/vendor/cache/little-plugger-1.1.4.gem
./bosh/bosh-director/vendor/cache/logging-1.8.2.gem
./bosh/bosh-director/vendor/cache/membrane-1.1.0.gem
./bosh/bosh-director/vendor/cache/mime-types-2.6.2.gem
./bosh/bosh-director/vendor/cache/mini_portile-0.6.2.gem
./bosh/bosh-director/vendor/cache/mono_logger-1.1.0.gem
./bosh/bosh-director/vendor/cache/multi_json-1.11.2.gem
./bosh/bosh-director/vendor/cache/mysql2-0.3.11.gem
./bosh/bosh-director/vendor/cache/nats-0.5.0.beta.12.gem
./bosh/bosh-director/vendor/cache/net-scp-1.1.2.gem
./bosh/bosh-director/vendor/cache/net-ssh-2.9.2.gem
./bosh/bosh-director/vendor/cache/net-ssh-gateway-1.2.0.gem
./bosh/bosh-director/vendor/cache/netaddr-1.5.0.gem
./bosh/bosh-director/vendor/cache/nokogiri-1.6.6.2.gem
./bosh/bosh-director/vendor/cache/pg-0.15.1.gem
./bosh/bosh-director/vendor/cache/rack-1.6.4.gem
./bosh/bosh-director/vendor/cache/rack-protection-1.5.3.gem
./bosh/bosh-director/vendor/cache/rack-test-0.6.2.gem
./bosh/bosh-director/vendor/cache/rake-10.3.2.gem
./bosh/bosh-director/vendor/cache/redis-3.0.3.gem
./bosh/bosh-director/vendor/cache/redis-namespace-1.3.1.gem
./bosh/bosh-director/vendor/cache/resque-1.25.2.gem
./bosh/bosh-director/vendor/cache/resque-backtrace-0.0.1.gem
./bosh/bosh-director/vendor/cache/rest-client-1.6.7.gem
./bosh/bosh-director/vendor/cache/rufus-scheduler-2.0.24.gem
./bosh/bosh-director/vendor/cache/semi_semantic-1.1.0.gem
./bosh/bosh-director/vendor/cache/sequel-3.43.0.gem
./bosh/bosh-director/vendor/cache/sinatra-1.4.6.gem
./bosh/bosh-director/vendor/cache/sinatra-contrib-1.4.2.gem
./bosh/bosh-director/vendor/cache/sys-filesystem-1.1.4.gem
./bosh/bosh-director/vendor/cache/thin-1.5.1.gem
./bosh/bosh-director/vendor/cache/thread_safe-0.3.5.gem
./bosh/bosh-director/vendor/cache/tilt-1.4.1.gem
./bosh/bosh-director/vendor/cache/tzinfo-1.2.2.gem
./bosh/bosh-director/vendor/cache/vegas-0.1.11.gem
./bosh/bosh-director/vendor/cache/yajl-ruby-1.2.1.gem

Files for package 'genisoimage': ./
./genisoimage/
./genisoimage/cdrkit-1.1.11.tar.gz
./packaging

Files for package 'health_monitor': ./
./bosh/
./packaging
./syslog_event_forwarder/
./syslog_event_forwarder/setup_syslog_event_forwarder.sh
./bosh/bosh-monitor/
./bosh/bosh-monitor/vendor/
./bosh/bosh-monitor/vendor/cache/
./bosh/bosh-monitor/vendor/cache/addressable-2.3.7.gem
./bosh/bosh-monitor/vendor/cache/aws-sdk-1.60.2.gem
./bosh/bosh-monitor/vendor/cache/aws-sdk-v1-1.60.2.gem
./bosh/bosh-monitor/vendor/cache/bosh-monitor-1.0000.0.gem
./bosh/bosh-monitor/vendor/cache/cf-uaa-lib-3.2.1.gem
./bosh/bosh-monitor/vendor/cache/daemons-1.2.3.gem
./bosh/bosh-monitor/vendor/cache/dogapi-1.6.0.gem
./bosh/bosh-monitor/vendor/cache/em-http-request-0.3.0.gem
./bosh/bosh-monitor/vendor/cache/escape_utils-1.1.0.gem
./bosh/bosh-monitor/vendor/cache/eventmachine-1.0.3.gem
./bosh/bosh-monitor/vendor/cache/httpclient-2.4.0.gem
./bosh/bosh-monitor/vendor/cache/json-1.8.3.gem
./bosh/bosh-monitor/vendor/cache/json_pure-1.8.1.gem
./bosh/bosh-monitor/vendor/cache/little-plugger-1.1.4.gem
./bosh/bosh-monitor/vendor/cache/logging-1.8.2.gem
./bosh/bosh-monitor/vendor/cache/mini_portile-0.6.2.gem
./bosh/bosh-monitor/vendor/cache/multi_json-1.11.2.gem
./bosh/bosh-monitor/vendor/cache/nats-0.5.0.beta.12.gem
./bosh/bosh-monitor/vendor/cache/nokogiri-1.6.6.2.gem
./bosh/bosh-monitor/vendor/cache/rack-1.6.4.gem
./bosh/bosh-monitor/vendor/cache/rack-protection-1.5.3.gem
./bosh/bosh-monitor/vendor/cache/rack-test-0.6.2.gem
./bosh/bosh-monitor/vendor/cache/sinatra-1.4.6.gem
./bosh/bosh-monitor/vendor/cache/sinatra-contrib-1.4.2.gem
./bosh/bosh-monitor/vendor/cache/thin-1.5.1.gem
./bosh/bosh-monitor/vendor/cache/tilt-1.4.1.gem
./bosh/bosh-monitor/vendor/cache/yajl-ruby-1.2.1.gem

Files for package 'libpq': ./
./postgres/
./postgres/postgresql-9.0.23.tar.gz
./packaging
./config/
./config/config.sub
./config/config.guess

Files for package 'mysql': ./
./mysql/
./mysql/server-5.1.62-rel13.3-435-Linux-x86_64.tar.gz
./packaging

Files for package 'nats': ./
./nats/
./nats/Gemfile
./nats/Gemfile.lock
./nats/vendor/
./nats/vendor/cache/
./nats/vendor/cache/rack-1.5.2.gem
./nats/vendor/cache/yajl-ruby-1.1.0.gem
./nats/vendor/cache/eventmachine-1.0.3.gem
./nats/vendor/cache/thin-1.5.1.gem
./nats/vendor/cache/json_pure-1.8.1.gem
./nats/vendor/cache/nats-0.5.0.beta.12.gem
./nats/vendor/cache/daemons-1.1.9.gem
./packaging

Files for package 'nginx': ./
./nginx/
./nginx/headers-more-v0.25.tgz
./nginx/nginx-1.4.5.tar.gz
./nginx/nginx-upload-module-2.2.tar.gz
./nginx/pcre-8.37.tar.gz
./packaging

Files for package 'postgres': ./
./postgres/
./postgres/postgresql-9.0.23.tar.gz
./packaging

Files for package 'powerdns': ./
./packaging
./powerdns/
./powerdns/pdns-static_3.3.1-1_amd64.deb

Files for package 'redis': ./
./redis/
./redis/redis-3.0.2.tar.gz
./packaging

Files for package 'registry': ./
./bosh/
./packaging
./bosh/bosh-registry/
./bosh/bosh-registry/vendor/
./bosh/bosh-registry/vendor/cache/
./bosh/bosh-registry/vendor/cache/aws-sdk-1.60.2.gem
./bosh/bosh-registry/vendor/cache/aws-sdk-v1-1.60.2.gem
./bosh/bosh-registry/vendor/cache/bosh-registry-1.0000.0.gem
./bosh/bosh-registry/vendor/cache/bosh_common-1.0000.0.gem
./bosh/bosh-registry/vendor/cache/bosh_cpi-1.0000.0.gem
./bosh/bosh-registry/vendor/cache/builder-3.1.4.gem
./bosh/bosh-registry/vendor/cache/CFPropertyList-2.3.2.gem
./bosh/bosh-registry/vendor/cache/daemons-1.2.3.gem
./bosh/bosh-registry/vendor/cache/eventmachine-1.0.3.gem
./bosh/bosh-registry/vendor/cache/excon-0.45.4.gem
./bosh/bosh-registry/vendor/cache/fission-0.5.0.gem
./bosh/bosh-registry/vendor/cache/fog-1.34.0.gem
./bosh/bosh-registry/vendor/cache/fog-atmos-0.1.0.gem
./bosh/bosh-registry/vendor/cache/fog-aws-0.7.6.gem
./bosh/bosh-registry/vendor/cache/fog-brightbox-0.9.0.gem
./bosh/bosh-registry/vendor/cache/fog-core-1.32.1.gem
./bosh/bosh-registry/vendor/cache/fog-dynect-0.0.2.gem
./bosh/bosh-registry/vendor/cache/fog-ecloud-0.3.0.gem
./bosh/bosh-registry/vendor/cache/fog-google-0.1.0.gem
./bosh/bosh-registry/vendor/cache/fog-json-1.0.2.gem
./bosh/bosh-registry/vendor/cache/fog-local-0.2.1.gem
./bosh/bosh-registry/vendor/cache/fog-powerdns-0.1.1.gem
./bosh/bosh-registry/vendor/cache/fog-profitbricks-0.0.5.gem
./bosh/bosh-registry/vendor/cache/fog-radosgw-0.0.4.gem
./bosh/bosh-registry/vendor/cache/fog-riakcs-0.1.0.gem
./bosh/bosh-registry/vendor/cache/fog-sakuracloud-1.3.3.gem
./bosh/bosh-registry/vendor/cache/fog-serverlove-0.1.2.gem
./bosh/bosh-registry/vendor/cache/fog-softlayer-0.4.7.gem
./bosh/bosh-registry/vendor/cache/fog-storm_on_demand-0.1.1.gem
./bosh/bosh-registry/vendor/cache/fog-terremark-0.1.0.gem
./bosh/bosh-registry/vendor/cache/fog-vmfusion-0.1.0.gem
./bosh/bosh-registry/vendor/cache/fog-voxel-0.1.0.gem
./bosh/bosh-registry/vendor/cache/fog-xml-0.1.2.gem
./bosh/bosh-registry/vendor/cache/formatador-0.2.5.gem
./bosh/bosh-registry/vendor/cache/inflecto-0.0.2.gem
./bosh/bosh-registry/vendor/cache/ipaddress-0.8.0.gem
./bosh/bosh-registry/vendor/cache/json-1.8.3.gem
./bosh/bosh-registry/vendor/cache/little-plugger-1.1.4.gem
./bosh/bosh-registry/vendor/cache/logging-1.8.2.gem
./bosh/bosh-registry/vendor/cache/membrane-1.1.0.gem
./bosh/bosh-registry/vendor/cache/mime-types-2.6.2.gem
./bosh/bosh-registry/vendor/cache/mini_portile-0.6.2.gem
./bosh/bosh-registry/vendor/cache/multi_json-1.11.2.gem
./bosh/bosh-registry/vendor/cache/mysql2-0.3.11.gem
./bosh/bosh-registry/vendor/cache/net-scp-1.1.2.gem
./bosh/bosh-registry/vendor/cache/net-ssh-2.9.2.gem
./bosh/bosh-registry/vendor/cache/net-ssh-gateway-1.2.0.gem
./bosh/bosh-registry/vendor/cache/nokogiri-1.6.6.2.gem
./bosh/bosh-registry/vendor/cache/pg-0.15.1.gem
./bosh/bosh-registry/vendor/cache/rack-1.6.4.gem
./bosh/bosh-registry/vendor/cache/rack-protection-1.5.3.gem
./bosh/bosh-registry/vendor/cache/rack-test-0.6.2.gem
./bosh/bosh-registry/vendor/cache/semi_semantic-1.1.0.gem
./bosh/bosh-registry/vendor/cache/sequel-3.43.0.gem
./bosh/bosh-registry/vendor/cache/sinatra-1.4.6.gem
./bosh/bosh-registry/vendor/cache/sinatra-contrib-1.4.2.gem
./bosh/bosh-registry/vendor/cache/thin-1.5.1.gem
./bosh/bosh-registry/vendor/cache/tilt-1.4.1.gem
./bosh/bosh-registry/vendor/cache/yajl-ruby-1.2.1.gem

Files for package 'ruby': ./
./ruby/
./ruby/ruby-2.1.7.tar.gz
./ruby/bundler-1.10.6.gem
./ruby/rubygems-2.4.8.tgz
./ruby/yaml-0.1.5.tar.gz
./packaging
./config/
./config/config.sub
./config/config.guess
```
