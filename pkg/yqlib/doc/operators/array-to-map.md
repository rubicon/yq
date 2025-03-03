# Array to Map

Use this operator to convert an array to..a map. The indices are used as map keys, null values in the array are skipped over.

Behind the scenes, this is implemented using reduce:

```
(.[] | select(. != null) ) as $i ireduce({}; .[$i | key] = $i)
```

## Simple example
Given a sample.yml file of:
```yaml
cool:
  - null
  - null
  - hello
```
then
```bash
yq '.cool |= array_to_map' sample.yml
```
will output
```yaml
cool:
  2: hello
```

