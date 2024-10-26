

# 実行方法
# terraform apply -var="file_content=Your Custom Content"
# ↓
# terraform apply -var="file_content=My Custom Content"
# のように引数を変えた場合でも差分として認識される

# terraform apply -var="file_content=Your Custom Content" -var="number_input=5"