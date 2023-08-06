# Usage

`larson` loads parameters from AWS ssm parameter store, and adds them as shell environment variables.

## Download parameters as a json file

```
$ larson get-parameters /cmacdonaldtest/ > ./example.json
```

```
$ cat ./example.json

{
    "alpha": "the_alpha_value",
    "beta": "the_beta_value",
    "delta": "the_delta_value"
}
```

## Load json parameters as environment variables

```
source larson_json_to_vars ./example.json
```

```
$ env | grep 'alpha\|beta\|delta'
alpha=the_alpha_value
delta=the_delta_value
beta=the_beta_value
```

## Upload parameters to parameter store

```
$ larson put-parameters /cmacdonaldtest/ --input-file=./new-values.json
```

# Installation

git clone, `pip install -e .`

If and when we put it on an accessible pypi, `pip install larson` should work.

# Tests

an excellent idea...