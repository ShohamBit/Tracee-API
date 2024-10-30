# Tracee client

## run tracee

__NOTE:__ tracee config file doesn't support: policy, scope, event & capture

``` bash
Tracee \
--config ../path/to/config/config{you connection type}.yaml \
--policy ../path/to/policies 
```

## run client

``` bash
sudo ./dist/traceectl 
```

you have cupule of options:
    1. version
    2. stream event
    3. enable event
    4. disable event
    5. metrics

``` bash
sudo ./dist/traceectl version
sudo ./dist/traceectl streamEvents [policies]
sudo ./dist/traceectl enableEvents [events]
sudo ./dist/traceectl disableEvents [events]
sudo ./dist/traceectl metrics
sudo ./dist/traceectl 
```

## test stream events

the client has a policy directory under `traceectl/policies`

run stream events:

```bash
./dist/traceectl streamEvents
```

you can also format the output of tracee client by adding the flag `output`.
`output` support 3 types of outputs:

1. table (default)
2. json
3. go template

you can run it like this:

```bash
./dist/traceectl streamEvents --output json
```

stream events from policy1:

```bash
./dist/traceectl streamEvents policy1
./dist/traceectl streamEvents policy2
```

stream event from 2 policies

``` bash
./dist/traceectl streamEvents policy1 policy2
```

you can run this commend to check for events,
run this in command in each dir

``` bash
cat hi.txt
```
