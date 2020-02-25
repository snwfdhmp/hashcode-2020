VERSION = $(shell git describe --tags)

run:
	go run . --input a --cpuprofile ./profiles/a.prof
	go run . --input b --cpuprofile ./profiles/b.prof
	go run . --input c --cpuprofile ./profiles/c.prof
	go run . --input d --cpuprofile ./profiles/d.prof
	go run . --input e --cpuprofile ./profiles/e.prof
	go run . --input f --cpuprofile ./profiles/f.prof
	zip -r output/sources.zip .