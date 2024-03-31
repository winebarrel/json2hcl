{
  Statement = [{
    Condition = {
      StringNotEquals = {
        "aws:RequestedRegion" = ["eu-central-1", "eu-west-1", "eu-west-2", "eu-west-3"]
      }
    }
    Effect    = "Deny"
    NotAction = ["cloudfront:*", "iam:*", "route53:*", "support:*"]
    Resource  = "*"
    Sid       = "DenyAllOutsideRequestedRegions"
  }]
  Version = "2012-10-17"
}
