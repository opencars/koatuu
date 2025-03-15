.PHONY: default all clean
APPS     := server parser
BLDDIR   := bin
IMPORT_BASE := github.com/opencars/koatuu

default: clean all

all: $(APPS)

$(BLDDIR)/%:
	go build -o $@ ./cmd/$*

$(APPS): %: $(BLDDIR)/%

clean:
	@mkdir -p $(BLDDIR)
	@for app in $(APPS) ; do \
		rm -f $(BLDDIR)/$$app ; \
	done