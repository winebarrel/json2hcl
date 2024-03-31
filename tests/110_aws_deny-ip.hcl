{
  Statement = {
    Action = "*"
    Condition = {
      NotIpAddress = {
        "aws:SourceIp" = ["192.0.2.0/24", "203.0.113.0/24"]
      }
    }
    Effect   = "Deny"
    Resource = "*"
  }
  Version = "2012-10-17"
}
