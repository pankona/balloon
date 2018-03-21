DOCS_DIR = $(CURDIR)/../docs

all:
	gitbook build &&        \
	rm -r $(DOCS_DIR) && \
	mv $(CURDIR)/_book $(DOCS_DIR)
