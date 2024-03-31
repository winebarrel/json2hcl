{
  Statement = [{
    Action   = "iam:GetAccountPasswordPolicy"
    Effect   = "Allow"
    Resource = "*"
    Sid      = "ViewAccountPasswordRequirements"
    }, {
    Action   = ["iam:GetUser", "iam:ChangePassword"]
    Effect   = "Allow"
    Resource = "arn:aws:iam::*:user/$${aws:username}"
    Sid      = "ChangeOwnPassword"
  }]
  Version = "2012-10-17"
}
