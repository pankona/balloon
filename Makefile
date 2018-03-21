
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

.PHONY: docs
docs:
	make -C $(CURDIR)/docsrc
