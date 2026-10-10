VERSION = 0.1.0
GO = go

peak: peak.go home.min.html
	CGO_ENABLED=0 ${GO} build -trimpath -ldflags="-s -w -X main.Version=${VERSION}" -o $@ peak.go

%.min.html: %.html
	sed 's/^[ \t]*//' $< | tr -d '\n' > $@

clean:
	rm -f peak home.min.html

.PHONY: clean
