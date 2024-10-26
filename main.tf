terraform {
  # backend：stateファイルをどこに置くか

  # backend "s3" { # S3にstateファイルを置く
  #   bucket         = "kst-shun.sakakida" # ファイルを置くバケット名
  #   key            = "terraform/state/terraform.tfstate"  # バケット内のパス
  #   region         = "ap-northeast-1"  # バケットのリージョン
  #   dynamodb_table = "kst-shun.sakakida"  # DynamoDBテーブル名：ロックを行うために必要
  #   encrypt        = true  # 状態ファイルの暗号化を有効にする
  # }
}


# stateファイルをS3に保存するモジュール
module "state_s3" {
  source         = "./modules/state_s3"  # モジュールのパス
  identifier = var.identifier       # ルートモジュールからプレフィックスを渡す
}