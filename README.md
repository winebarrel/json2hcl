# json2hcl

A tool to convert JSON to HCL.

### Usage

## Installation

```sh
brew install winebarrel/json2hcl/json2hcl
```

## Usage

```go
Usage: json2hcl [OPTION] [FILE]
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
  Statement = [{
    Action = "service-prefix:action-name"
    Condition = {
      DateGreaterThan = {
        "aws:CurrentTime" = "2020-04-01T00:00:00Z"
      }
      DateLessThan = {
        "aws:CurrentTime" = "2020-06-30T23:59:59Z"
      }
    }
    Effect   = "Allow"
    Resource = "*"
  }]
  Version = "2012-10-17"
}
```
