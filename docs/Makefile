DOCS_DIR = $(CURDIR)/../docs

all: $(CURDIR)/node_modules
	cd $(CURDIR)      && \
	gitbook build     && \
	rm -r $(DOCS_DIR) && \
	mv $(CURDIR)/_book $(DOCS_DIR)

$(CURDIR)/node_modules:
	cd $(CURDIR) && \
	gitbook install
