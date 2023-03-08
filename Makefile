genCode:
	go build .
	./ginkgo_gen -f example/model.go -cf example/tracing_testcase.md -o ./