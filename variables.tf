
# # 変数を定義
# variable "file_content" {
#   description = "The content to be written to the file"
#   type = string
#   default = "variable.txt"
# }
# variable "number_input" {
#   description = "A numeric"
#   type = number
#   default = 0

#   // バリデーションルールを追加
#   validation {
#     condition     = var.number_input >= 0
#     error_message = "The number_input must be a non-negative number."
#   }
# }
# resource "local_file" "variable" {
#   filename = "${path.module}/output/variable.txt"
#   content  = "${var.file_content}\nNumber: ${var.number_input}" # 引数で渡された値を使用
# }

# 実行方法
# terraform apply -var="file_content=Your Custom Content"
# ↓
# terraform apply -var="file_content=My Custom Content"
# のように引数を変えた場合でも差分として認識される

# terraform apply -var="file_content=Your Custom Content" -var="number_input=5"