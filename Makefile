
ifndef TAG
TAG = release
endif

all: pc

pc:
	cd $(CURDIR) && \
	go build -tags=$(TAG)

mobile:
	cd $(CURDIR) && \
	gomobile build -tags=$(TAG)

clean:
	rm -f balloon balloon.apk

.PHONY: docs
docs:
	make -C $(CURDIR)/docsrc

doccommit:
	git add $(CURDIR)/docs
	git commit -m "update site" $(CURDIR)/docs
