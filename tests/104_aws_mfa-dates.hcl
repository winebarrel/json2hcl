{
  Statement = {
    Action = ["service-prefix-1:*", "service-prefix-2:action-name-a", "service-prefix-2:action-name-b"]
    Condition = {
      Bool = {
        "aws:MultiFactorAuthPresent" = true
      }
      DateGreaterThan = {
        "aws:CurrentTime" = "2017-07-01T00:00:00Z"
      }
      DateLessThan = {
        "aws:CurrentTime" = "2017-12-31T23:59:59Z"
      }
    }
    Effect   = "Allow"
    Resource = "*"
  }
  Version = "2012-10-17"
}
