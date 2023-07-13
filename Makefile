commit: format test
	git add -A .
	git commit -m "$(info)"
	git push origin master
cur:
	git tag|sort -r|head -n1
tag:
	git tag $(tag)
	git push origin $(tag)
format:
	@find ./ -type f -name "*.go" -exec gofmt -w {} \;
test:
	go test ./...