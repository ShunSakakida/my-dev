terraform {
  # backend：stateファイルをどこに置くか：変数をできない
  backend "s3" { # S3にstateファイルを置く
    bucket         = "kst-sakakida-tf-for-state" # ファイルを置くバケット名
    key            = "terraform/state/terraform.tfstate"  # バケット内のパス
    region         = "ap-northeast-1"
    dynamodb_table = "kst-sakakida-tf-for-state"  # DynamoDBテーブル名：ロックを行うために必要
    encrypt        = true  # 状態ファイルの暗号化を有効にする
  }
}


# stateファイルをS3に保存するモジュールを呼び出す
module "state_s3" {
  source        = "./modules/state_s3"  # モジュールのパス
  identifier    = local.identifier
  resource_name = local.resource_name_for_tf_state
}
