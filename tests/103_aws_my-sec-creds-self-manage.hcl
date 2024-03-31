{
  Statement = [{
    Action   = ["iam:GetAccountPasswordPolicy", "iam:ListVirtualMFADevices"]
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
    }, {
    Action   = ["iam:CreateVirtualMFADevice"]
    Effect   = "Allow"
    Resource = "arn:aws:iam::*:mfa/*"
    Sid      = "AllowManageOwnVirtualMFADevice"
    }, {
    Action   = ["iam:DeactivateMFADevice", "iam:EnableMFADevice", "iam:ListMFADevices", "iam:ResyncMFADevice"]
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
    NotAction = ["iam:CreateVirtualMFADevice", "iam:EnableMFADevice", "iam:GetUser", "iam:GetMFADevice", "iam:ListMFADevices", "iam:ListVirtualMFADevices", "iam:ResyncMFADevice", "sts:GetSessionToken"]
    Resource  = "*"
    Sid       = "DenyAllExceptListedIfNoMFA"
  }]
  Version = "2012-10-17"
}
