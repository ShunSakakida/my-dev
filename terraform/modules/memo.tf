# module "file_writer" {
#   # フォルダでモジュールとなるため、フォルダのパスを指定
#   source      = "./modules/local_file_writer"
#   # モジュールに渡す引数
#   file_content = "This is a test content"
#   file_name     = "test_output.txt"
# }

# // データソースを使用してファイルの内容を読み取る → 結果を file_content に代入
# data "local_file" "file_content" {
#   # file_writerというモジュールの戻り値を参照して、読み取るファイル名を指定している
#   filename = module.file_writer.file_path
# }

# // 出力を表示
# output "created_file_path" {
#   value = module.file_writer.file_path
# }

# output "file_content" {
#   value = data.local_file.file_content.content
# }

# resource "local_file" "sample_main1" {
#   filename = "${path.module}/output/sample.main1.txt"
#   content = "Hello World"
# }
