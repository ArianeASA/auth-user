resource "aws_s3_bucket" "lambda_bucket" {
  bucket = "auth-user-x-lambda-bucket-go"
}

resource "aws_s3_bucket_public_access_block" "auth-public-access-block" {
  bucket = aws_s3_bucket.lambda_bucket.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

resource "aws_s3_bucket_ownership_controls" "lambda_bucket_controls" {
  bucket = aws_s3_bucket.lambda_bucket.id
  rule {
    object_ownership = "BucketOwnerPreferred"
  }
}

resource "aws_s3_bucket_acl" "lambda_bucket_acl" {
  depends_on = [aws_s3_bucket_ownership_controls.lambda_bucket_controls]

  bucket = aws_s3_bucket.lambda_bucket.id
  acl    = "private"
}

locals {
  source_file = "../${path.module}/bootstrap"
  output_path_zip = "../${path.module}/bootstrap.zip"

}

data "archive_file" "lambda_zip" {
  type = "zip"

  source_file  = local.source_file
  output_path = local.output_path_zip
}


resource "aws_s3_object" "lambda_main" {
  bucket = aws_s3_bucket.lambda_bucket.id

  key    = "authuser/${formatdate("YYYYMMDD", timestamp())}/hash${formatdate("hhmmss", timestamp())}-bootstrap.zip"
  source = local.output_path_zip
  acl   = "private"
  etag = data.archive_file.lambda_zip.output_base64sha256
}

resource "aws_iam_policy" "s3_get_object" {
  name        = "s3_get_auth_user_xx_object"
  description = "Allows access to the S3 object"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Action = [
          "s3:GetObject"
        ],
        Resource = [
          "arn:aws:s3:::${aws_s3_bucket.lambda_bucket.id}/${aws_s3_object.lambda_main.key}"
        ]
      },{
        Sid       = "HTTPSOnly"
        Effect    = "Deny"
        Action    = "s3:*"
        Resource = [
          aws_s3_bucket.lambda_bucket.arn,
          "${aws_s3_bucket.lambda_bucket.arn}/*",
        ]
        Condition = {
          Bool = {
            "aws:SecureTransport" = "false"
          }
        }
      },
    ]
  })
}

resource "aws_iam_role_policy_attachment" "lambda_exec_s3_get_object" {
  role       = aws_iam_role.lambda_exec.name
  policy_arn = aws_iam_policy.s3_get_object.arn
}