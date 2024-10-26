variable "identifier" {
  description = "識別子"
  type        = string
}
variable "resource_name" {
  description = "リソース名"
  type        = string

#   // バリデーションルールを追加
#   validation {
#     condition     = var.number_input >= 0
#     error_message = "The number_input must be a non-negative number."
#   }
}
