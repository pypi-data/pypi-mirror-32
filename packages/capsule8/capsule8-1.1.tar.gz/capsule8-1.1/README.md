# Capsule8 api-python

This repo is as a wrapper repo around the [Capsule8 sensor](https://github.com/capsule8/capsule8). Specifically this repo is targeted at the `.proto` definitions for the sensor and generates stubs for python.  

## Install with pip

```
pip install https://github.com/capsule8/api-python/archive/master.zip
```

## Install with python

```
git clone https://github.com/capsule8/api-python.git
cd api-python/
make
```

## Testing

In order to run the functional tests capsule8 python api you need to have a running capsule8 sensor on your machine. Go to [capsule8](https://github.com/capsule8/capsule8) for instructions.

To test capsule8 and all it's packages run

```
sudo chown $USER /var/run/capsule8/sensor.sock
make test
```

## Examples

In order to run the examples you need to have a running capsule8 sensor on your machine. Once you have a running sensor you can run the example

```
sudo chown $USER /var/run/capsule8/sensor.sock
python3 print_process_events.py
```

And you should see output similar to
```
events {
  event {
    id: "aaf19f70cae7b2018f7d55237bd4e0f80ca06d39a0c1876efb481caf42f302b3"
    process_id: "f0dcdb8c5ad84874a7f210326353a5d8fef331f137fc167d25e09ec7c762ad07"
    process_pid: 19959
    sensor_id: "17f25eddd795798d57990b7cb33721ecbe1774ee96a195acfd2f0d84edc67f26"
    sensor_sequence_number: 97
    sensor_monotime_nanos: 24165408812
    process {
      type: PROCESS_EVENT_TYPE_FORK
      fork_child_pid: 9443
    }
    credentials {
      uid: 1001
      gid: 1001
      euid: 1001
      egid: 1001
      suid: 1001
      sgid: 1001
      fsuid: 1001
      fsgid: 1001
    }
    process_tgid: 19959
  }
}

events {
  event {
    id: "53f5e4fbef0ee56acb97074553a0c137c4cb12e626b9b864252efb570562df7f"
    process_id: "046ace5eda2fe335cabe4c9cc11ce626563677181297f38a837ba3a7cad17964"
    process_pid: 9443
    sensor_id: "17f25eddd795798d57990b7cb33721ecbe1774ee96a195acfd2f0d84edc67f26"
    sensor_sequence_number: 98
    sensor_monotime_nanos: 24170786264
    process {
      type: PROCESS_EVENT_TYPE_FORK
      fork_child_pid: 9444
    }
    cpu: 3
    credentials {
      uid: 1001
      gid: 1001
      euid: 1001
      egid: 1001
      suid: 1001
      sgid: 1001
      fsuid: 1001
      fsgid: 1001
    }
    process_tgid: 9443
  }
}
```

## Usage

For usage on how to use generated gRPC code, follow [gRPC basics](https://grpc.io/docs/tutorials/basic/python.html).

## Build new protos

Running `make protos` will create new protos from the capsule8 open source sensor api
