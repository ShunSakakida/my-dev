// 出力を定義
output "file_path" {
  description = "The path of the created file"
  value       = local_file.file_create.filename # 戻り値
}