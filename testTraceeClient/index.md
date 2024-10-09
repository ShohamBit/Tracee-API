# Tracee client

## run tracee

``` bash
sudo ./dist/tracee \
--config ../TraceeClient/config.yaml \
--policy ../TraceeClient/policies \
```

## run client

``` bash
sudo ./dist/TraceeClient 
```

you have cupule of options:
    1. version
    2. stream event
    3. enable event
    4. disable event
    5. metrics

``` bash
sudo ./dist/TraceeClient version
sudo ./dist/TraceeClient streamEvents [policies]
sudo ./dist/TraceeClient enableEvents [events]
sudo ./dist/TraceeClient disableEvents [events]
sudo ./dist/TraceeClient metrics
sudo ./dist/TraceeClient 
```

## test stream events

the client has a policy directory under `TraceeClient/policies`

run stream events:

```bash
./dist/TraceeClient streamEvents
```

stream events from policy1:

```bash
./dist/TraceeClient streamEvents policy1
./dist/TraceeClient streamEvents policy2
```

stream evenet from 2 policies

``` bash
./dist/TraceeClient streamEvents policy1 policy2
```

you can run this commend to check for events,
run this in command in each dir

``` bash
cat hi.txt
```
