EMDIR=${HOME}/.emacs.d

install: mwnci.el
	../../install.sh -m 0700 -d ${EMDIR}
	cp mwnci.el ${EMDIR}
	touch ${EMDIR}/init.el
	if [ ! "`grep -s load-file.*mwnci.el ${EMDIR}/init.el`" ] ; \
	then \
	    echo '(load-file "~/.emacs.d/mwnci.el")' >> ${EMDIR}/init.el ; \
	fi
clean:
	find . -name \#\* -type f -exec rm {} \;
	find . -name \*~ -type f -exec rm {} \;
