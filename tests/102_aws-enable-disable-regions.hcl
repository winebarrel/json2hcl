{
  Statement = [{
    Action = ["account:EnableRegion", "account:DisableRegion"]
    Condition = {
      StringEquals = {
        "account:TargetRegion" = "ap-east-1"
      }
    }
    Effect   = "Allow"
    Resource = "*"
    Sid      = "EnableDisableHongKong"
    }, {
    Action   = ["account:ListRegions"]
    Effect   = "Allow"
    Resource = "*"
    Sid      = "ViewConsole"
  }]
  Version = "2012-10-17"
}
