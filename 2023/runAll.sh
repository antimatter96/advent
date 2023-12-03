ADVENT_DEBUG=false
for day in day_*; do
  cd ${day}
  echo ${day}
  cat input.txt | go run main.go
  cd ..
done
