VIM=${HOME}/.vim

install:
	../../install.sh -m 0700 -d ${VIM}/ftdetect ${VIM}/syntax
	../../install.sh -t ${VIM}/ftdetect ftdetect/mwnci.vim
	../../install.sh -t ${VIM}/syntax syntax/mwnci.vim
	if [ ! "`grep -s ^au.*mwnci ${VIM}/filetype.vim`" ] ; \
	then \
	    cat filetype.vim >> ${VIM}/filetype.vim ; \
	fi

clean:
	find . -name \#\* -type f -exec rm {} \;
	find . -name \*~ -type f -exec rm {} \;
