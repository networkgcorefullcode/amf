<!--
SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
Copyright 2019 free5GC.org

SPDX-License-Identifier: Apache-2.0
-->
[![Go Report Card](https://goreportcard.com/badge/github.com/omec-project/amf)](https://goreportcard.com/report/github.com/omec-project/amf)

# amf
It is a control plane function in the 5G core network. AMF supports termination
of NAS signalling, NAS ciphering & integrity protection, registration
management, connection management, mobility management, access authentication
and authorization, security context management.


## Repository Structure

Below is a high-level view of the repository and its main components:

```
.
├── amf.go
├── amfTest                     # Example test configurations for AMF setup and behavior
│   ├── amfcfg_with_custom_webui_url.yaml
│   └── amfcfg.yaml
├── amf_test.go
├── communication               # Service-Based Interface (SBI) definitions and routers
│   ├── api_individual_subscription_document.go
│   ├── api_individual_ue_context_document.go
│   ├── api_n1_n2_individual_subscription_document.go
│   ├── api_n1_n2_message_collection_document.go
│   ├── api_n1_n2_subscriptions_collection_for_individual_ue_contexts_document.go
│   ├── api_non_uen2_message_notification_individual_subscription_document.go
│   ├── api_non_uen2_messages_collection_document.go
│   ├── api_non_uen2_messages_subscriptions_collection_document.go
│   ├── api_subscriptions_collection_document.go
│   └── routers.go
├── consumer                    # Outgoing NF interactions
│   ├── am_policy.go
│   ├── communication.go
│   ├── nf_discovery.go
│   ├── nf_mangement.go
│   ├── nsselection.go
│   ├── sm_context.go
│   ├── subscriber_data_management.go
│   ├── ue_authentication.go
│   └── ue_context_management.go
├── context                     # Core runtime context: UE, RAN, sessions, timers
│   ├── 3gpp_types.go
│   ├── amf_ran.go
│   ├── amf_ue.go
│   ├── common_function.go
│   ├── context.go
│   ├── db.go
│   ├── ran_ue.go
│   ├── sm_context.go
│   ├── timer.go
│   └── transaction.go
├── Dockerfile              # Container build definitions
├── Dockerfile.fast
├── docs                        # Documentation and diagrams
│   └── images
│       ├── README-AMF.png
│       └── README-AMF.png.license
├── eventexposure               # 3GPP-defined event exposure services
│   ├── api_individual_subscription_document.go
│   ├── api_subscriptions_collection_document.go
│   └── routers.go
├── factory                     # Configuration loading and validation logic
│   ├── amf_config_test.go
│   ├── config.go
│   └── factory.go
├── gmm                         # 5G Mobility Management procedures and message handling
│   ├── handler.go
│   ├── init.go
│   ├── init_test.go
│   ├── message
│   │   ├── build.go
│   │   └── send.go
│   ├── mock_gmm.go
│   └── sm.go
├── go.mod
├── go.mod.license
├── go.sum
├── go.sum.license
├── httpcallback                # HTTP callback endpoints for notifications and subscriptions
│   ├── api_am_policy_control_update_notify.go
│   ├── api_dereg_notify.go
│   ├── api_n1_message_notify.go
│   ├── api_nf_subscribe_notify.go
│   ├── api_sm_context_status_notify.go
│   └── router.go
├── LICENSES
│   └── Apache-2.0.txt
├── location                    # Location management APIs
│   ├── api_individual_ue_context_document.go
│   └── routers.go
├── logger                      # Centralized logging for AMF
│   └── logger.go
├── Makefile                    # Build and test automation
├── metrics                     # Telemetry and Kafka-based metrics publication
│   ├── kafka.go
│   └── telemetry.go
├── msgtypes                    
│   └── ngapmsgtypes
│       └── ngapmsgtypes.go
├── mt
│   ├── api_ue_context_document.go
│   ├── api_ue_reach_ind_document.go
│   └── routers.go
├── nas                         # NAS message parsing, dispatching, and security handling
│   ├── dispatch.go
│   ├── handler.go
│   └── nas_security
│       └── security.go
├── ngap                        # NGAP protocol handlers and utilities
│   ├── dispatcher.go
│   ├── handler.go
│   ├── message
│   │   ├── build.go
│   │   ├── forward_ie.go
│   │   ├── send.go
│   │   └── send_test.go.test
│   ├── ngap_test.go
│   ├── service
│   │   └── service.go
│   └── util
│       └── ngap_util.go
├── NOTICE.txt
├── oam                         # OAM APIs for UE context operations
│   ├── api_purge_ue_context.go
│   ├── api_registered_ue_context.go
│   └── routers.go
├── producer                    # API handlers for incoming requests and callbacks
│   ├── callback
│   │   ├── n1n2message.go
│   │   ├── subscription.go
│   │   └── ue_context.go
│   ├── callback.go
│   ├── event_exposure.go
│   ├── location_info.go
│   ├── mt.go
│   ├── n1n2message.go
│   ├── oam.go
│   ├── oam_test.go
│   ├── subscription.go
│   └── ue_context.go
├── protos                      # gRPC service definitions and generated code
│   ├── sdcoreAmfServer
│   │   ├── server_grpc.pb.go
│   │   └── server.pb.go
│   └── server.proto
├── README.md
├── service                     # AMF initialization and server startup logic
│   ├── amf_server.go
│   └── init.go
├── Taskfile.yml                # Build and test automation
├── test-mirror.txt
├── util                        # Shared utilities, conversion helpers, test data
│   ├── convert.go
│   ├── init_context.go
│   ├── json.go
│   ├── mock.drsm.go
│   ├── search_nf_service.go
│   ├── test
│   │   ├── testAmfcfg2.yaml
│   │   └── testAmfcfg.yaml
│   └── util_func.go
├── VERSION
└── VERSION.license

33 directories, 117 files
```

## Configuration and Deployment

**Docker**

To build the container image:

```bash
task mod-start
task webconsole-ui
task docker-build-fast
```

**Kubernetes**

The standard deployment uses [Helm charts](https://charts.aetherproject.org) from the Aether project. The version of the Chart can be found in the OnRamp repository in the `vars/main.yml` file.


## Quick Navegation

| Path             | Description                                                      |
| ---------------- | ---------------------------------------------------------------- |
| `context/`       | Core runtime state of UEs, RANs, and SM contexts.                |
| `consumer/`      | Logic for inter-NF communication and discovery.                  |
| `producer/`      | Exposes 5G SBI APIs and manages incoming NF requests.            |
| `gmm/`           | Implements UE registration, service request, and mobility flows. |
| `ngap/`          | Manages NGAP signaling between AMF and gNB.                      |
| `nas/`           | Handles NAS encoding, decoding, and security.                    |
| `metrics/`       | Collects and publishes telemetry metrics.                        |
| `factory/`       | Loads and validates AMF configuration.                           |
| `eventexposure/` | Implements event exposure interfaces for observability.          |
| `logger/`        | Provides structured logging utilities across modules.            |


## AMF Block Diagram
![AMF Block Diagram](/docs/images/README-AMF.png)

AMF takes configuration from Configuration Service. Configuration is handled at
Network Slice level. Configuration (Network Slices) can be added, removed and
deleted. AMF has prometheus interface to export metrics. Metrics include
connected gNodeB's and its status.

## The SD-Core AMF currently supports the following functionalities:
- Termination of RAN CP interface (N2)
- Termination of NAS (N1), NAS ciphering and integrity protection
- Registration management
- Connection management
- Reachability management
- Mobility Management
- Provide transport for SM messages between UE and SMF
- Transparent proxy for routing SM messages
- Access Authentication
- Access Authorization

## Supported Procedures:
- Registration/Deregistration
- Registration update
- UE initiated Service request
- N2 Handover
- Xn handover
- PDU Establishment Request/Release
- Paging
- CN High Availibilty and Stateless session support
- AMF metrics are available via metricfunc on the 5g Grafana dashboard

## Upcoming Changes in AMF



Compliance of the 5G Network functions can be found at [5G Compliance](https://docs.sd-core.opennetworking.org/main/overview/3gpp-compliance-5g.html)

## Reach out to us thorugh

1. #sdcore-dev channel in [ONF Community Slack](https://onf-community.slack.com/)
2. Raise Github issues
