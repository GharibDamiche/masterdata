#!/usr/bin/env bash

cd /sources/hyperion/contract

readarray -d '' list < <(find ./ -name "*.pb.go" -print0)

for f in "${list[@]}"; do
  out=$(LANG=en_US.UTF-8 git diff --shortstat $f 2>&1)
  if [[ $out != *"1 file changed, 1 insertion(+), 1 deletion(-)"* && $out != *"1 file changed, 2 insertions(+), 2 deletions(-)"* ]]; then
    continue
  fi

  out=$(LANG=en_US.UTF-8 git diff --word-diff $f 2>&1)

  if [[ $out != *"protoc-gen-go"* ]]; then
    continue
  fi

  out=$(LANG=en_US.UTF-8 git checkout $f 2>&1)

  #basename=$(basename "$f")
  #filename="${basename%.*}"
  #dir=$(dirname "$f")
  #mkdir -p "./docs/api/${dir}"
  #protoc --proto_path=. --doc_out="./docs/api/$dir" --doc_opt=markdown,"${filename}.md" "${f}"
  #sed -i "s/Protocol Documentation/${filename}/g" "./docs/api/${dir}/${filename}.md"
done
