# json2hcl

[![CI](https://github.com/winebarrel/json2hcl/actions/workflows/ci.yml/badge.svg)](https://github.com/winebarrel/json2hcl/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/winebarrel/json2hcl/graph/badge.svg?token=DMI0AG4U48)](https://codecov.io/gh/winebarrel/json2hcl)

A tool to convert JSON to HCL.

### Usage

## Installation

```sh
brew install winebarrel/json2hcl/json2hcl
```

## Usage

```
Usage: json2hcl [OPTION] [FILE]
  -n	do not escape ${...} and %{...} template sequences
  -version
    	print version and exit
```

```sh
$ cat policy.json
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": "service-prefix:action-name",
            "Resource": "*",
            "Condition": {
                "DateGreaterThan": {"aws:CurrentTime": "2020-04-01T00:00:00Z"},
                "DateLessThan": {"aws:CurrentTime": "2020-06-30T23:59:59Z"}
            }
        }
    ]
}

$ json2hcl policy.json # or `cat policy.json | json2hcl`
{
  Version = "2012-10-17"
  Statement = [
    {
      Effect   = "Allow"
      Action   = "service-prefix:action-name"
      Resource = "*"
      Condition = {
        DateGreaterThan = {
          "aws:CurrentTime" = "2020-04-01T00:00:00Z"
        }
        DateLessThan = {
          "aws:CurrentTime" = "2020-06-30T23:59:59Z"
        }
      }
    },
  ]
}
```

## Related tools

- https://github.com/winebarrel/tfjson2hcl
