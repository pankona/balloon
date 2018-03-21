
ifndef TAG
TAG = release
endif

all: pc

pc:
	cd $(CURDIR) && \
	go build -tags=$(TAG) -o balloon

mobile:
	cd $(CURDIR) && \
	gomobile build -tags=$(TAG) -o balloon.apk

.PHONY: docs
docs:
	make -C $(CURDIR)/docsrc
