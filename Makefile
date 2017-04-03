


README.txt: withenv.1
	groff -C -Tutf8 -man $< | col -b > $@
