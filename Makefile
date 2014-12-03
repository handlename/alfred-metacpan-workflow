VERSION=${shell cat VERSION}
CMD=cmd/workflow/workflow
WORKFLOW=metacpan-$(VERSION).alfredworkflow

dist/$(WORKFLOW): $(CMD)
	zip -jv $@ $(CMD) icon.png info.plist

$(CMD): *.go
	cd cmd/workflow; go build -v

.PHONY: clean
clean:
	rm $(CMD) $(WORKFLOW)

.PHONY: ship
ship:
	ghr $(VERSION) dist/$(WORKFLOW)
