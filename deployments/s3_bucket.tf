resource "aws_s3_bucket" "lambda_bucket" {
  bucket = "auth-user-lambda-bucket-go"
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

  key    = "/authuser/${formatdate("YYYYMMDD", timestamp())}/hash${formatdate("hhmmss", timestamp())}-bootstrap.zip"
  source = local.output_path_zip
  acl   = "private"
  etag = data.archive_file.lambda_zip.output_base64sha256
}
