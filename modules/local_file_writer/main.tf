// local_fileリソースで変数を使用
resource "local_file" "file_create" {
  filename = "${path.module}/${var.file_name}"
  content  = var.file_content
}