{
  Statement = [{
    Action   = ["iam:GetAccountPasswordPolicy", "iam:GetAccountSummary"]
    Effect   = "Allow"
    Resource = "*"
    Sid      = "AllowViewAccountInfo"
    }, {
    Action   = ["iam:ChangePassword", "iam:GetUser"]
    Effect   = "Allow"
    Resource = "arn:aws:iam::*:user/$${aws:username}"
    Sid      = "AllowManageOwnPasswords"
    }, {
    Action   = ["iam:CreateAccessKey", "iam:DeleteAccessKey", "iam:ListAccessKeys", "iam:UpdateAccessKey", "iam:GetAccessKeyLastUsed"]
    Effect   = "Allow"
    Resource = "arn:aws:iam::*:user/$${aws:username}"
    Sid      = "AllowManageOwnAccessKeys"
    }, {
    Action   = ["iam:DeleteSSHPublicKey", "iam:GetSSHPublicKey", "iam:ListSSHPublicKeys", "iam:UpdateSSHPublicKey", "iam:UploadSSHPublicKey"]
    Effect   = "Allow"
    Resource = "arn:aws:iam::*:user/$${aws:username}"
    Sid      = "AllowManageOwnSSHPublicKeys"
  }]
  Version = "2012-10-17"
}
