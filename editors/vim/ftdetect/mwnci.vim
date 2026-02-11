au BufNewFile,BufRead * if getline(1) =~ '#!.*mwnci$' | setlocal filetype=mwnci | endif
