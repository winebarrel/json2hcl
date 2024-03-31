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
