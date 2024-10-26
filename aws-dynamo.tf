resource "aws_dynamodb_table" "example" {
  name         = "example-table"
  billing_mode = "PAY_PER_REQUEST"

  hash_key  = "ID"         // パーティションキー
  range_key = "Timestamp"  // ソートキー

  attribute {
    name = "ID"
    type = "S"
  }

  attribute {
    name = "Timestamp"
    type = "N"  // ソートキーのデータ型
  }

  tags = {
    Name = "ExampleTable"
  }
}