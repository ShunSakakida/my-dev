terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

provider "aws" {
  region = "ap-northeast-1"  # 東京リージョンを指定
}

# 記載を消す（今回はコメントアウト）ことでapplyしたら実際のリソースも削除される
# resource "aws_s3_bucket" "my_bucket" {
#   bucket = "my-unique-bucket-name-12345-20231010"

#   tags = {
#     Name        = "MyBucket"
#     Environment = "Dev"
#   }
# }

# バケットポリシーを使用してアクセス制御を設定する例 → アタッチできなかった
# resource "aws_s3_bucket_policy" "my_bucket_policy" {
#   bucket = aws_s3_bucket.my_bucket.id

#   policy = jsonencode({
#     Version = "2012-10-17"
#     Statement = [
#       {
#         Effect = "Allow"
#         Principal = "*"
#         Action = [
#           "s3:GetObject"
#         ]
#         Resource = [
#           "${aws_s3_bucket.my_bucket.arn}/*"
#         ]
#       }
#     ]
#   })
# }

# EC2インスタンス
# resource "aws_instance" "example" { # aws_instance = EC2インスタンスを作成するリソース
#   ami           = "ami-00bf64aaeb8d382c9"  # Amazon Linux 2 AMI
#   instance_type = "t2.micro"

#   tags = {
#     Name = "ExampleInstance"
#   }
# }