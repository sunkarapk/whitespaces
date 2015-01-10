VERSION = 1.1.1

GPM = gpm
GO_XC = goxc
GOXC_FILE = .goxc.local.json

all: deps install

install:
	go install -a github.com/pksunkara/whitespaces

deps:
	$(GPM)

goxc:
	$(shell echo '{\n "ArtifactsDest": "build",\n "ConfigVersion": "0.9",' > $(GOXC_FILE))
	$(shell echo ' "PackageVersion": "$(VERSION)",\n "TaskSettings": {' >> $(GOXC_FILE))
	$(shell echo '  "bintray": {\n   "apikey": "",\n   "package": "whitespaces",' >> $(GOXC_FILE))
	$(shell echo '   "repository": "utils",\n   "subject": "pksunkara"' >> $(GOXC_FILE))
	$(shell echo '  }\n }\n}' >> $(GOXC_FILE))
	$(GO_XC)

bintray:
	$(GO_XC) bintray

clean:
	rm -rf build debian .goxc.local.json
