
all:
	@(cd sipgen; $(MAKE))
	@(cd siplib; $(MAKE))

install:
	@(cd sipgen; $(MAKE) install)
	@(cd siplib; $(MAKE) install)
	@test -d $(DESTDIR)/usr/lib/python2.7/dist-packages || mkdir -p $(DESTDIR)/usr/lib/python2.7/dist-packages
	cp -f sipconfig.py $(DESTDIR)/usr/lib/python2.7/dist-packages/sipconfig.py
	cp -f /home/pjr/git_neon.2/OGN-Flight-Logger_V3.2-dev/OGN_Flogger/sip-4.19.8/sipdistutils.py $(DESTDIR)/usr/lib/python2.7/dist-packages/sipdistutils.py

clean:
	@(cd sipgen; $(MAKE) clean)
	@(cd siplib; $(MAKE) clean)
