{
  Statement = [{
    Action   = "iam:ListVirtualMFADevices"
    Effect   = "Allow"
    Resource = "*"
    Sid      = "AllowViewAccountInfo"
    }, {
    Action   = ["iam:CreateVirtualMFADevice"]
    Effect   = "Allow"
    Resource = "arn:aws:iam::*:mfa/*"
    Sid      = "AllowManageOwnVirtualMFADevice"
    }, {
    Action   = ["iam:DeactivateMFADevice", "iam:EnableMFADevice", "iam:GetUser", "iam:GetMFADevice", "iam:ListMFADevices", "iam:ResyncMFADevice"]
    Effect   = "Allow"
    Resource = "arn:aws:iam::*:user/$${aws:username}"
    Sid      = "AllowManageOwnUserMFA"
    }, {
    Condition = {
      BoolIfExists = {
        "aws:MultiFactorAuthPresent" = "false"
      }
    }
    Effect    = "Deny"
    NotAction = ["iam:CreateVirtualMFADevice", "iam:EnableMFADevice", "iam:GetUser", "iam:ListMFADevices", "iam:ListVirtualMFADevices", "iam:ResyncMFADevice", "sts:GetSessionToken"]
    Resource  = "*"
    Sid       = "DenyAllExceptListedIfNoMFA"
  }]
  Version = "2012-10-17"
}
