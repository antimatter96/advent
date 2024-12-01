for day in day_*; do
  go test -timeout 5s -run ^TestRun$ github.com/antimatter96/advent/2024/${day}
done
