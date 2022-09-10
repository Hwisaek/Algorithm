i=1
while read line; do
  if [ ${i} -eq 10 ]; then
    echo "$line"
  fi
  ((i += 1))
done <file.txt