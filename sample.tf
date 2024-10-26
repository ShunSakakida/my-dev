# terraform.tfstateにリソースの状態を管理している
# リソース作成が実行されている場合は、スキップされる


# resource "resouce_type" "resouce_name" {
#   # 処理内容
# }

resource "local_file" "sample" {
  filename = "${path.module}/output/sample.txt"
  content = "Hello World"
}

resource "local_file" "sample2" {
  filename = "${path.module}/output/sample2.txt"
  content = "Hello World"
}

resource "local_file" "sample3" {
  filename = "${path.module}/output/sample3.txt"
  content = "Hello World"
}

resource "local_file" "sample4" {
  filename = "${path.module}/output/sample4.txt"
  content = "Hello World"
}

resource "local_file" "sample5" {
  filename = "${path.module}/output/sample5.txt"
  content = "Hello World"
}
