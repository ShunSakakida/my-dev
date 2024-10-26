# modules/s3/main.tf
resource "aws_s3_bucket" "state_bucket" {
  bucket = "${var.identifier}-for-state"  # プレフィックスを使用
  acl    = "private"                      # アクセス制御リスト（ACL）

  tags = {
    Name        = "${var.identifier}"
  }
}