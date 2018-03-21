
all: pc

pc:
	cd $(CURDIR) && go build -o balloon

mobile:
	cd $(CURDIR) && gomobile build -o balloon.apk

.PHONY: docs
docs:
	make -C $(CURDIR)/docsrc
