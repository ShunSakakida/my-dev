locals {
  resource_name = "${var.identifier}-for-state"
}


# Stateファイルを保存するS3
resource "aws_s3_bucket" "state_bucket" {
  bucket = local.resource_name
  acl    = "private" # アクセス制御リスト（ACL）

  tags = {
    Name = "${var.identifier}"
  }
}

# Stateファイルのロックを管理するDynamoDB
resource "aws_dynamodb_table" "terraform_locks" {
  name         = local.resource_name
  billing_mode = "PAY_PER_REQUEST"
   attribute {
     name = "LockID"
     type = "S"
   }
   hash_key = "LockID"

  tags = {
    Name = "${var.identifier}"
  }
}