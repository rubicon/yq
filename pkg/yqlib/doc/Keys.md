
## Get Keys
Given a sample.yml file of:
```yaml
a:
  b: cat
  c: dog
  d: frog
```
then
```bash
yq eval '.a | keys' sample.yml
```
will output
```yaml
3
```

## Set key style
Given a sample.yml file of:
```yaml
a:
  b: cat
  c: dog
  d: frog
```
then
```bash
yq eval '(.a | keys) style = 'single'' sample.yml
```
will output
```yaml
a:
  b: cat
  c: dog
  d: frog
```

