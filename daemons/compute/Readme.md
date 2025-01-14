# Data Movement

The nnf-dm service is a compute resident daemon that provides an interface for submitting copy offload requests to the Rabbit.

## Build

The nnf-dm service resides in [./server](./server) directory and can be built using `go build -o nnf-dm`

## Installation

Copy the nnf-dm daemon to a suitable location such as `/usr/bin`. After defining the necessary [environmental variables](#environmental-variables), install the daemon by running `nnf-dm install`. Verify the daemon is running using `systemctl status nnf-dm`

### Authentication

NNF software defines a Kubernetes service account for communicating data movement requests to the kubeapi server. The token file and certificate file can be downloaded by running `cert-load.sh` and verified with `cert-test.sh` The token file and certificate file must be provided to the nnf-dm daemon.

### Customizing the Configuration

Installing the nnf-dm daemon will create a default configuration located at `/etc/systemd/system/nnf-dm.service`. The default configuration created on installation is sparse as the use of environmental variables is assumed. If desired, one can edit the configuration with the [command line options](#command-line-options). An example is show below.

```conf
[Unit]
Description=Near-Node Flash (NNF) Data Movement Service

[Service]
PIDFile=/var/run/nnf-dm.pid
ExecStartPre=/bin/rm -f /var/run/nnf-dm.pid
ExecStart=/usr/bin/nnf-dm \
   --kubernetes-service-host=172.0.0.1 \
   --kubernetes-service-port=7777 \
   --nnf-data-movement-service-token-file=/path/to/service.token \
   --nnf-data-movement-service-cert-file=/path/to/cert.file \
   --nnf-node-name=rabbit-01
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

### Configuration Overrides

systemd can be used to create an override file that is useful for non-standard behavior of the data movement service. Run `systemctl edit nnf-dm` to create or edit the override file.

For example, running the nnf-dm service in simulation mode is achieved with the following (the empty ExecStart is required)

```conf
[Service]
ExecStart=
ExecStart=/root/nnf-dm --simulated
```

## Environmental Variables

| Name | Description |
| ---- | ----------- |
| `KUBERNETES_SERVICE_HOST` | Specifies the kubernetes service host |
| `KUBERNETES_SERVICE_PORT` | Specifies the kubernetes service port |
| `NODE_NAME` | Specifies the node name of this host |
| `NNF_NODE_NAME` | Specifies the NNF node name connected to this host |
| `NNF_DATA_MOVEMENT_SERVICE_TOKEN_FILE` | Specifies the path to the NNF Data Movement service token file |
| `NNF_DATA_MOVEMENT_SERVICE_CERT_FILE`  | Specifies the path to the NNF Data Movement service certificate file |

## Command Line Options

Command line options are defined below. Most default to their corresponding environmental variables of the same name in upper-case with hyphens replaced by underscores

| Option | Default | Description |
| ------ | ------- | ----------- |
| `--socket=[PATH]` | `/var/run/nnf-dm.sock` | Specifies the listening socket for the NNF Data Movement daemon |
| `--kubernetes-service-host=[HOST]` | `KUBERNETES_SERVICE_HOST` | Specifies the kubernetes service host |
| `--kubernetes-service-port=[PORT]` | `KUBERNETES_SERVICE_PORT` | Specifies the kubernetes service port |
| `--node-name=[NAME]` | `NODE_NAME` | Specifies the node name of this host |
| `--nnf-node-name=[NAME]` | `NNF_NODE_NAME` | Specifies the NNF node name connected to this host |
| `--nnf-data-movement-service-token-file=[PATH]` | `NNF_DATA_MOVEMENT_SERVICE_TOKEN_FILE` | Specifies the path to the NNF Data Movement service token file |
| `--nnf-data-movement-service-cert-file=[PATH]` | `NNF_DATA_MOVEMENT_SERVICE_CERT_FILE`  | Specifies the path to the NNF Data Movement service certificate file |

# Data Movement Interface

The nnf-dm service uses Protocol Buffers to define a set of APIs for initiating, querying, canceling and deleting data movement requests. The definitions for these can be found in [./api/datamovement.proto](./api/datamovement.proto) file. Consulting this file should be done in addition to the context in this section to ensure the latest API definitions are used.

## Common fields for all requests

Each nnf-dm request contains a Workflow message containing the Name and Namespace fields. WLM must provide these values to the user application that are then provided to the data-movement API. The recommended implementation is to use environmental variables.

### Workflow

| Field | Type | Description |
| ----- | ---- | ----------- |
| `name` | string | Name of the workflow that is associated with this data movement request. WLM must provide this as an environmental variable (i.e. `DW_WORKFLOW_NAME`) |
| `namespace` | string | Namespace of the workflow that is associated with this data movement request. WLM must provide this as an environmental variable (i.e. `DW_WORKFLOW_NAMESPACE`) |

## Creating a data movement request

### Create Request

| Field | Type | Description |
| ----- | ---- | ----------- |
| `workflow` | Workflow | The name and namespace of the initiating workflow |
| `source` | string | The source file or directory |
| `destination` | string | The destination file or directory |
| `dryrun` | bool | If True, the rsync copy operation should evaluate the inputs but not perform the copy |

### Create Response

| Field | Type | Description |
| ----- | ---- | ----------- |
| `uid` | string | The unique identifier for the created data movement resource if `Status` is Success |
| `status` | Status | Status of the data movement create request |
| `message` | string | String providing detailed message pertaining to the current status of the create request |

#### Create Response `Status`

| Value | Name | Description |
| ----- | ---- | ----------- |
| 0 | Success | The data movement resource created successfully |
| 1 | Failed | The data movement resource failed to create. See `message` field for more information |

## Query the status of a data movement request

### Status Request

| Field | Type | Description |
| ----- | ---- | ----------- |
| `workflow` | Workflow | The name and namespace of the initiating workflow |
| `uid` | string | The unique identifier for the created data movement resource |
| `maxWaitTime` | int64 | The maximum time in seconds to wait for completion of the data movement resource. Negative values imply an indefinite wait |

### Status Response

| Field | Type | Description |
| ----- | ---- | ----------- |
| `state` | State | Current state of the data movement request |
| `status` | Status | Current status of the data movement request |
| `message` | string | String providing detailed message pertaining to the current state/status of the request |
| `commandStatus` | CommandStatus | Status of the data movement command. Includes progress, timing, and last message |

#### Status Response `State`

| Value | Name | Description |
| ----- | ---- | ----------- |
| 0 | Pending | The request is created but has a pending status |
| 1 | Starting | The request was created and is in the act of starting |
| 2 | Running | The data movement request is actively running |
| 3 | Completed | The data movement request has completed |
| 4 | Cancelling | The data movement request has started the cancellation process |

#### Status Response `Status`

| Value | Name | Description |
| ----- | ---- | ----------- |
| 0 | Invalid | The request was found to be invalid. See `message` for details |
| 1 | Not Found | The request with the supplied UID was not found |
| 2 | Success | The request completed with success |
| 3 | Failed | The request failed. See `message` for details |

#### Status Response `CommandStatus`

| Field | Type | Description |
| ----- | ---- | ----------- |
| command | string | The full command that is performing data movement |
| progress| int32 | The progress percentage of the data movement (e.g. 50) |
| elapsedTime | string | The elapsed time of the data movement (e.g. 25s) |
| lastMessage | string | The most recent line of output from the data movement command |
| lastMessageTime | string | The time of `lastMessage` |

## List the data movement requests

`namespace` and `workflow` must be supplied in order to retrieve data movement
requests. An empty list of `uids` will be returned if there are no data movement
requests that match both `namespace` and `workflow`.

### List Request

| Field | Type | Description |
| ----- | ---- | ----------- |
| `workflow` | Workflow | The name and namespace of the initiating workflow |

### List Response

| Field | Type | Description |
| ----- | ---- | ----------- |
| `uids` | string[] | List of data movement requests associated with the given workflow and namespace. |

## Cancel a data movement request

Data movement requests can be cancelled after creation of a data movement request. Cancel Responses
will be sent back after the Cancellation has been successfully (or unsuccessfully) initiated. The
user can then poll with a Status Request to verify the end result of the data movement request to
ensure the cancellation was successful.

While the data movement is being cancelled, a Status Response will report a state of `Cancelling`
until the Cancellation is complete. At that time, a state of `Completed` with a status of
`Cancelled` will be returned. See the Status Response section above.

### Cancel Request

| Field | Type | Description |
| ----- | ---- | ----------- |
| `workflow` | Workflow | The name and namespace of the initiating workflow |
| `uid` | string | The unique identifier for the data movement resource to be cancelled |

### Cancel Response

| Field | Type | Description |
| ----- | ---- | ----------- |
| `status` | Status | Cancel status of the data movement request |
| `message` | string | String providing detailed message pertaining to the cancellation of the request |

#### Cancel Response `Status`

| Value | Name | Description |
| ----- | ---- | ----------- |
| 0 | Invalid | The cancel request was found to be invalid |
| 1 | Not Found | The request with the supplied UID was not found |
| 2 | Success | The data movement request was cancelled successfully |
| 3 | Failed | The data movement request cannot be canceled |

## Delete a data movement request

Data movement requests can be deleted once a data movement request has been completed (successful or
otherwise). Delete Responses will be sent back after the Deletion has been successfully (or
unsuccessfully) initiated. The user can then use a Status Request to verify the end result of the
data movement request to ensure the deletion was successful.

### Delete Request

| Field | Type | Description |
| ----- | ---- | ----------- |
| `workflow` | Workflow | The name and namespace of the initiating workflow |
| `uid` | string | The unique identifier for the data movement resource |

### Delete Response

| Field | Type | Description |
| ----- | ---- | ----------- |
| `status` | Status | Deletion status of the data movement request |
| `message` | string | String providing detailed message pertaining to the deletion of the request |

#### Delete Response `Status`

| Value | Name | Description |
| ----- | ---- | ----------- |
| 0 | Invalid | The delete request was found to be invalid |
| 1 | Not Found | The request with the supplied UID was not found |
| 2 | Success | The data movement request was deleted successfully |
| 3 | Active | The data movement request is currently active and cannot be deleted |

# Example Clients

## Go

The Go client, which resides in `./client-go/`, is the most customizable client example. It demonstrates a number of command line options to configure the Create request. It provides a debug hook to _not_ delete the request after completion, so one can test that the requests are cleaned up as part of the workflow.

## C

The C client, which resides in `./_client-c/`, is a simple client that uses CGo as an interface to the server. (The leading underscore "\_" in the directory is so Go ignores directory as part of a larger repository build). As there is no native C GRPC implementation, Go is used to interface with the GRPC server while providing an interface to the C code. A Makefile is provided to build the various components and assemble it into a final executable.

The C client should not be used for anything more than rapid prototyping. No new features will be added to the C client.

## C++

The C++ client, which resides in `./client-cpp/`, is a simple client that uses [gRPC C++](https://grpc.io/docs/languages/cpp/).

## Python

The Python client, which resides in `./client-py/`, is a very simple client. No customization options are provided
