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
    Action   = ["iam:DeleteSigningCertificate", "iam:ListSigningCertificates", "iam:UpdateSigningCertificate", "iam:UploadSigningCertificate"]
    Effect   = "Allow"
    Resource = "arn:aws:iam::*:user/$${aws:username}"
    Sid      = "AllowManageOwnSigningCertificates"
    }, {
    Action   = ["iam:DeleteSSHPublicKey", "iam:GetSSHPublicKey", "iam:ListSSHPublicKeys", "iam:UpdateSSHPublicKey", "iam:UploadSSHPublicKey"]
    Effect   = "Allow"
    Resource = "arn:aws:iam::*:user/$${aws:username}"
    Sid      = "AllowManageOwnSSHPublicKeys"
    }, {
    Action   = ["iam:CreateServiceSpecificCredential", "iam:DeleteServiceSpecificCredential", "iam:ListServiceSpecificCredentials", "iam:ResetServiceSpecificCredential", "iam:UpdateServiceSpecificCredential"]
    Effect   = "Allow"
    Resource = "arn:aws:iam::*:user/$${aws:username}"
    Sid      = "AllowManageOwnGitCredentials"
  }]
  Version = "2012-10-17"
}
