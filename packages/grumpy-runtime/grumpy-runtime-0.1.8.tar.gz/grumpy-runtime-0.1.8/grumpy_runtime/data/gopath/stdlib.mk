all:

build/src/__python__/__builtin__/module.go: build/src/__python__/__builtin__.py
	@mkdir -p $(@D)
	@grumpc -modname=__builtin__ $< > $@

build/src/__python__/__builtin__/module.d: build/src/__python__/__builtin__.py
	@mkdir -p $(@D)
	@pydeps -modname=__builtin__ $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/__builtin__.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/__builtin__.a: build/src/__python__/__builtin__/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/__builtin__ -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/__builtin__.a

-include build/src/__python__/__builtin__/module.d

build/src/__python__/_abcoll/module.go: build/src/__python__/_abcoll.py
	@mkdir -p $(@D)
	@grumpc -modname=_abcoll $< > $@

build/src/__python__/_abcoll/module.d: build/src/__python__/_abcoll.py
	@mkdir -p $(@D)
	@pydeps -modname=_abcoll $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/_abcoll.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/_abcoll.a: build/src/__python__/_abcoll/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/_abcoll -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/_abcoll.a

-include build/src/__python__/_abcoll/module.d

build/src/__python__/_collections/module.go: build/src/__python__/_collections.py
	@mkdir -p $(@D)
	@grumpc -modname=_collections $< > $@

build/src/__python__/_collections/module.d: build/src/__python__/_collections.py
	@mkdir -p $(@D)
	@pydeps -modname=_collections $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/_collections.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/_collections.a: build/src/__python__/_collections/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/_collections -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/_collections.a

-include build/src/__python__/_collections/module.d

build/src/__python__/_csv/module.go: build/src/__python__/_csv.py
	@mkdir -p $(@D)
	@grumpc -modname=_csv $< > $@

build/src/__python__/_csv/module.d: build/src/__python__/_csv.py
	@mkdir -p $(@D)
	@pydeps -modname=_csv $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/_csv.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/_csv.a: build/src/__python__/_csv/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/_csv -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/_csv.a

-include build/src/__python__/_csv/module.d

build/src/__python__/_functools/module.go: build/src/__python__/_functools.py
	@mkdir -p $(@D)
	@grumpc -modname=_functools $< > $@

build/src/__python__/_functools/module.d: build/src/__python__/_functools.py
	@mkdir -p $(@D)
	@pydeps -modname=_functools $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/_functools.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/_functools.a: build/src/__python__/_functools/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/_functools -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/_functools.a

-include build/src/__python__/_functools/module.d

build/src/__python__/_md5/module.go: build/src/__python__/_md5.py
	@mkdir -p $(@D)
	@grumpc -modname=_md5 $< > $@

build/src/__python__/_md5/module.d: build/src/__python__/_md5.py
	@mkdir -p $(@D)
	@pydeps -modname=_md5 $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/_md5.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/_md5.a: build/src/__python__/_md5/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/_md5 -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/_md5.a

-include build/src/__python__/_md5/module.d

build/src/__python__/_random/module.go: build/src/__python__/_random.py
	@mkdir -p $(@D)
	@grumpc -modname=_random $< > $@

build/src/__python__/_random/module.d: build/src/__python__/_random.py
	@mkdir -p $(@D)
	@pydeps -modname=_random $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/_random.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/_random.a: build/src/__python__/_random/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/_random -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/_random.a

-include build/src/__python__/_random/module.d

build/src/__python__/_sha/module.go: build/src/__python__/_sha.py
	@mkdir -p $(@D)
	@grumpc -modname=_sha $< > $@

build/src/__python__/_sha/module.d: build/src/__python__/_sha.py
	@mkdir -p $(@D)
	@pydeps -modname=_sha $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/_sha.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/_sha.a: build/src/__python__/_sha/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/_sha -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/_sha.a

-include build/src/__python__/_sha/module.d

build/src/__python__/_sha256/module.go: build/src/__python__/_sha256.py
	@mkdir -p $(@D)
	@grumpc -modname=_sha256 $< > $@

build/src/__python__/_sha256/module.d: build/src/__python__/_sha256.py
	@mkdir -p $(@D)
	@pydeps -modname=_sha256 $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/_sha256.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/_sha256.a: build/src/__python__/_sha256/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/_sha256 -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/_sha256.a

-include build/src/__python__/_sha256/module.d

build/src/__python__/_sha512/module.go: build/src/__python__/_sha512.py
	@mkdir -p $(@D)
	@grumpc -modname=_sha512 $< > $@

build/src/__python__/_sha512/module.d: build/src/__python__/_sha512.py
	@mkdir -p $(@D)
	@pydeps -modname=_sha512 $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/_sha512.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/_sha512.a: build/src/__python__/_sha512/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/_sha512 -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/_sha512.a

-include build/src/__python__/_sha512/module.d

build/src/__python__/_sre/module.go: build/src/__python__/_sre.py
	@mkdir -p $(@D)
	@grumpc -modname=_sre $< > $@

build/src/__python__/_sre/module.d: build/src/__python__/_sre.py
	@mkdir -p $(@D)
	@pydeps -modname=_sre $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/_sre.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/_sre.a: build/src/__python__/_sre/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/_sre -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/_sre.a

-include build/src/__python__/_sre/module.d

build/src/__python__/_struct/module.go: build/src/__python__/_struct.py
	@mkdir -p $(@D)
	@grumpc -modname=_struct $< > $@

build/src/__python__/_struct/module.d: build/src/__python__/_struct.py
	@mkdir -p $(@D)
	@pydeps -modname=_struct $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/_struct.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/_struct.a: build/src/__python__/_struct/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/_struct -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/_struct.a

-include build/src/__python__/_struct/module.d

build/src/__python__/_syscall/module.go: build/src/__python__/_syscall.py
	@mkdir -p $(@D)
	@grumpc -modname=_syscall $< > $@

build/src/__python__/_syscall/module.d: build/src/__python__/_syscall.py
	@mkdir -p $(@D)
	@pydeps -modname=_syscall $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/_syscall.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/_syscall.a: build/src/__python__/_syscall/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/_syscall -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/_syscall.a

-include build/src/__python__/_syscall/module.d

build/src/__python__/_weakrefset/module.go: build/src/__python__/_weakrefset.py
	@mkdir -p $(@D)
	@grumpc -modname=_weakrefset $< > $@

build/src/__python__/_weakrefset/module.d: build/src/__python__/_weakrefset.py
	@mkdir -p $(@D)
	@pydeps -modname=_weakrefset $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/_weakrefset.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/_weakrefset.a: build/src/__python__/_weakrefset/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/_weakrefset -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/_weakrefset.a

-include build/src/__python__/_weakrefset/module.d

build/src/__python__/abc/module.go: build/src/__python__/abc.py
	@mkdir -p $(@D)
	@grumpc -modname=abc $< > $@

build/src/__python__/abc/module.d: build/src/__python__/abc.py
	@mkdir -p $(@D)
	@pydeps -modname=abc $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/abc.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/abc.a: build/src/__python__/abc/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/abc -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/abc.a

-include build/src/__python__/abc/module.d

build/src/__python__/argparse/module.go: build/src/__python__/argparse.py
	@mkdir -p $(@D)
	@grumpc -modname=argparse $< > $@

build/src/__python__/argparse/module.d: build/src/__python__/argparse.py
	@mkdir -p $(@D)
	@pydeps -modname=argparse $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/argparse.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/argparse.a: build/src/__python__/argparse/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/argparse -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/argparse.a

-include build/src/__python__/argparse/module.d

build/src/__python__/base64/module.go: build/src/__python__/base64.py
	@mkdir -p $(@D)
	@grumpc -modname=base64 $< > $@

build/src/__python__/base64/module.d: build/src/__python__/base64.py
	@mkdir -p $(@D)
	@pydeps -modname=base64 $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/base64.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/base64.a: build/src/__python__/base64/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/base64 -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/base64.a

-include build/src/__python__/base64/module.d

build/src/__python__/binascii/module.go: build/src/__python__/binascii.py
	@mkdir -p $(@D)
	@grumpc -modname=binascii $< > $@

build/src/__python__/binascii/module.d: build/src/__python__/binascii.py
	@mkdir -p $(@D)
	@pydeps -modname=binascii $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/binascii.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/binascii.a: build/src/__python__/binascii/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/binascii -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/binascii.a

-include build/src/__python__/binascii/module.d

build/src/__python__/bisect/module.go: build/src/__python__/bisect.py
	@mkdir -p $(@D)
	@grumpc -modname=bisect $< > $@

build/src/__python__/bisect/module.d: build/src/__python__/bisect.py
	@mkdir -p $(@D)
	@pydeps -modname=bisect $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/bisect.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/bisect.a: build/src/__python__/bisect/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/bisect -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/bisect.a

-include build/src/__python__/bisect/module.d

build/src/__python__/collections/module.go: build/src/__python__/collections.py
	@mkdir -p $(@D)
	@grumpc -modname=collections $< > $@

build/src/__python__/collections/module.d: build/src/__python__/collections.py
	@mkdir -p $(@D)
	@pydeps -modname=collections $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/collections.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/collections.a: build/src/__python__/collections/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/collections -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/collections.a

-include build/src/__python__/collections/module.d

build/src/__python__/colorsys/module.go: build/src/__python__/colorsys.py
	@mkdir -p $(@D)
	@grumpc -modname=colorsys $< > $@

build/src/__python__/colorsys/module.d: build/src/__python__/colorsys.py
	@mkdir -p $(@D)
	@pydeps -modname=colorsys $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/colorsys.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/colorsys.a: build/src/__python__/colorsys/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/colorsys -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/colorsys.a

-include build/src/__python__/colorsys/module.d

build/src/__python__/contextlib/module.go: build/src/__python__/contextlib.py
	@mkdir -p $(@D)
	@grumpc -modname=contextlib $< > $@

build/src/__python__/contextlib/module.d: build/src/__python__/contextlib.py
	@mkdir -p $(@D)
	@pydeps -modname=contextlib $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/contextlib.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/contextlib.a: build/src/__python__/contextlib/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/contextlib -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/contextlib.a

-include build/src/__python__/contextlib/module.d

build/src/__python__/copy/module.go: build/src/__python__/copy.py
	@mkdir -p $(@D)
	@grumpc -modname=copy $< > $@

build/src/__python__/copy/module.d: build/src/__python__/copy.py
	@mkdir -p $(@D)
	@pydeps -modname=copy $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/copy.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/copy.a: build/src/__python__/copy/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/copy -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/copy.a

-include build/src/__python__/copy/module.d

build/src/__python__/copy_reg/module.go: build/src/__python__/copy_reg.py
	@mkdir -p $(@D)
	@grumpc -modname=copy_reg $< > $@

build/src/__python__/copy_reg/module.d: build/src/__python__/copy_reg.py
	@mkdir -p $(@D)
	@pydeps -modname=copy_reg $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/copy_reg.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/copy_reg.a: build/src/__python__/copy_reg/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/copy_reg -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/copy_reg.a

-include build/src/__python__/copy_reg/module.d

build/src/__python__/cStringIO/module.go: build/src/__python__/cStringIO.py
	@mkdir -p $(@D)
	@grumpc -modname=cStringIO $< > $@

build/src/__python__/cStringIO/module.d: build/src/__python__/cStringIO.py
	@mkdir -p $(@D)
	@pydeps -modname=cStringIO $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/cStringIO.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/cStringIO.a: build/src/__python__/cStringIO/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/cStringIO -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/cStringIO.a

-include build/src/__python__/cStringIO/module.d

build/src/__python__/csv/module.go: build/src/__python__/csv.py
	@mkdir -p $(@D)
	@grumpc -modname=csv $< > $@

build/src/__python__/csv/module.d: build/src/__python__/csv.py
	@mkdir -p $(@D)
	@pydeps -modname=csv $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/csv.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/csv.a: build/src/__python__/csv/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/csv -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/csv.a

-include build/src/__python__/csv/module.d

build/src/__python__/datetime/module.go: build/src/__python__/datetime.py
	@mkdir -p $(@D)
	@grumpc -modname=datetime $< > $@

build/src/__python__/datetime/module.d: build/src/__python__/datetime.py
	@mkdir -p $(@D)
	@pydeps -modname=datetime $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/datetime.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/datetime.a: build/src/__python__/datetime/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/datetime -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/datetime.a

-include build/src/__python__/datetime/module.d

build/src/__python__/difflib/module.go: build/src/__python__/difflib.py
	@mkdir -p $(@D)
	@grumpc -modname=difflib $< > $@

build/src/__python__/difflib/module.d: build/src/__python__/difflib.py
	@mkdir -p $(@D)
	@pydeps -modname=difflib $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/difflib.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/difflib.a: build/src/__python__/difflib/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/difflib -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/difflib.a

-include build/src/__python__/difflib/module.d

build/src/__python__/dircache/module.go: build/src/__python__/dircache.py
	@mkdir -p $(@D)
	@grumpc -modname=dircache $< > $@

build/src/__python__/dircache/module.d: build/src/__python__/dircache.py
	@mkdir -p $(@D)
	@pydeps -modname=dircache $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/dircache.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/dircache.a: build/src/__python__/dircache/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/dircache -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/dircache.a

-include build/src/__python__/dircache/module.d

build/src/__python__/dummy_thread/module.go: build/src/__python__/dummy_thread.py
	@mkdir -p $(@D)
	@grumpc -modname=dummy_thread $< > $@

build/src/__python__/dummy_thread/module.d: build/src/__python__/dummy_thread.py
	@mkdir -p $(@D)
	@pydeps -modname=dummy_thread $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/dummy_thread.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/dummy_thread.a: build/src/__python__/dummy_thread/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/dummy_thread -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/dummy_thread.a

-include build/src/__python__/dummy_thread/module.d

build/src/__python__/errno/module.go: build/src/__python__/errno.py
	@mkdir -p $(@D)
	@grumpc -modname=errno $< > $@

build/src/__python__/errno/module.d: build/src/__python__/errno.py
	@mkdir -p $(@D)
	@pydeps -modname=errno $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/errno.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/errno.a: build/src/__python__/errno/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/errno -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/errno.a

-include build/src/__python__/errno/module.d

build/src/__python__/exceptions/module.go: build/src/__python__/exceptions.py
	@mkdir -p $(@D)
	@grumpc -modname=exceptions $< > $@

build/src/__python__/exceptions/module.d: build/src/__python__/exceptions.py
	@mkdir -p $(@D)
	@pydeps -modname=exceptions $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/exceptions.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/exceptions.a: build/src/__python__/exceptions/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/exceptions -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/exceptions.a

-include build/src/__python__/exceptions/module.d

build/src/__python__/fnmatch/module.go: build/src/__python__/fnmatch.py
	@mkdir -p $(@D)
	@grumpc -modname=fnmatch $< > $@

build/src/__python__/fnmatch/module.d: build/src/__python__/fnmatch.py
	@mkdir -p $(@D)
	@pydeps -modname=fnmatch $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/fnmatch.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/fnmatch.a: build/src/__python__/fnmatch/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/fnmatch -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/fnmatch.a

-include build/src/__python__/fnmatch/module.d

build/src/__python__/fpformat/module.go: build/src/__python__/fpformat.py
	@mkdir -p $(@D)
	@grumpc -modname=fpformat $< > $@

build/src/__python__/fpformat/module.d: build/src/__python__/fpformat.py
	@mkdir -p $(@D)
	@pydeps -modname=fpformat $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/fpformat.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/fpformat.a: build/src/__python__/fpformat/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/fpformat -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/fpformat.a

-include build/src/__python__/fpformat/module.d

build/src/__python__/functools/module.go: build/src/__python__/functools.py
	@mkdir -p $(@D)
	@grumpc -modname=functools $< > $@

build/src/__python__/functools/module.d: build/src/__python__/functools.py
	@mkdir -p $(@D)
	@pydeps -modname=functools $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/functools.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/functools.a: build/src/__python__/functools/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/functools -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/functools.a

-include build/src/__python__/functools/module.d

build/src/__python__/genericpath/module.go: build/src/__python__/genericpath.py
	@mkdir -p $(@D)
	@grumpc -modname=genericpath $< > $@

build/src/__python__/genericpath/module.d: build/src/__python__/genericpath.py
	@mkdir -p $(@D)
	@pydeps -modname=genericpath $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/genericpath.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/genericpath.a: build/src/__python__/genericpath/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/genericpath -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/genericpath.a

-include build/src/__python__/genericpath/module.d

build/src/__python__/getopt/module.go: build/src/__python__/getopt.py
	@mkdir -p $(@D)
	@grumpc -modname=getopt $< > $@

build/src/__python__/getopt/module.d: build/src/__python__/getopt.py
	@mkdir -p $(@D)
	@pydeps -modname=getopt $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/getopt.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/getopt.a: build/src/__python__/getopt/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/getopt -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/getopt.a

-include build/src/__python__/getopt/module.d

build/src/__python__/glob/module.go: build/src/__python__/glob.py
	@mkdir -p $(@D)
	@grumpc -modname=glob $< > $@

build/src/__python__/glob/module.d: build/src/__python__/glob.py
	@mkdir -p $(@D)
	@pydeps -modname=glob $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/glob.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/glob.a: build/src/__python__/glob/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/glob -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/glob.a

-include build/src/__python__/glob/module.d

build/src/__python__/heapq/module.go: build/src/__python__/heapq.py
	@mkdir -p $(@D)
	@grumpc -modname=heapq $< > $@

build/src/__python__/heapq/module.d: build/src/__python__/heapq.py
	@mkdir -p $(@D)
	@pydeps -modname=heapq $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/heapq.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/heapq.a: build/src/__python__/heapq/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/heapq -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/heapq.a

-include build/src/__python__/heapq/module.d

build/src/__python__/itertools/module.go: build/src/__python__/itertools.py
	@mkdir -p $(@D)
	@grumpc -modname=itertools $< > $@

build/src/__python__/itertools/module.d: build/src/__python__/itertools.py
	@mkdir -p $(@D)
	@pydeps -modname=itertools $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/itertools.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/itertools.a: build/src/__python__/itertools/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/itertools -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/itertools.a

-include build/src/__python__/itertools/module.d

build/src/__python__/itertools_test/module.go: build/src/__python__/itertools_test.py
	@mkdir -p $(@D)
	@grumpc -modname=itertools_test $< > $@

build/src/__python__/itertools_test/module.d: build/src/__python__/itertools_test.py
	@mkdir -p $(@D)
	@pydeps -modname=itertools_test $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/itertools_test.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/itertools_test.a: build/src/__python__/itertools_test/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/itertools_test -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/itertools_test.a

-include build/src/__python__/itertools_test/module.d

build/src/__python__/json_scanner/module.go: build/src/__python__/json_scanner.py
	@mkdir -p $(@D)
	@grumpc -modname=json_scanner $< > $@

build/src/__python__/json_scanner/module.d: build/src/__python__/json_scanner.py
	@mkdir -p $(@D)
	@pydeps -modname=json_scanner $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/json_scanner.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/json_scanner.a: build/src/__python__/json_scanner/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/json_scanner -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/json_scanner.a

-include build/src/__python__/json_scanner/module.d

build/src/__python__/keyword/module.go: build/src/__python__/keyword.py
	@mkdir -p $(@D)
	@grumpc -modname=keyword $< > $@

build/src/__python__/keyword/module.d: build/src/__python__/keyword.py
	@mkdir -p $(@D)
	@pydeps -modname=keyword $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/keyword.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/keyword.a: build/src/__python__/keyword/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/keyword -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/keyword.a

-include build/src/__python__/keyword/module.d

build/src/__python__/linecache/module.go: build/src/__python__/linecache.py
	@mkdir -p $(@D)
	@grumpc -modname=linecache $< > $@

build/src/__python__/linecache/module.d: build/src/__python__/linecache.py
	@mkdir -p $(@D)
	@pydeps -modname=linecache $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/linecache.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/linecache.a: build/src/__python__/linecache/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/linecache -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/linecache.a

-include build/src/__python__/linecache/module.d

build/src/__python__/math/module.go: build/src/__python__/math.py
	@mkdir -p $(@D)
	@grumpc -modname=math $< > $@

build/src/__python__/math/module.d: build/src/__python__/math.py
	@mkdir -p $(@D)
	@pydeps -modname=math $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/math.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/math.a: build/src/__python__/math/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/math -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/math.a

-include build/src/__python__/math/module.d

build/src/__python__/math_test/module.go: build/src/__python__/math_test.py
	@mkdir -p $(@D)
	@grumpc -modname=math_test $< > $@

build/src/__python__/math_test/module.d: build/src/__python__/math_test.py
	@mkdir -p $(@D)
	@pydeps -modname=math_test $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/math_test.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/math_test.a: build/src/__python__/math_test/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/math_test -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/math_test.a

-include build/src/__python__/math_test/module.d

build/src/__python__/md5/module.go: build/src/__python__/md5.py
	@mkdir -p $(@D)
	@grumpc -modname=md5 $< > $@

build/src/__python__/md5/module.d: build/src/__python__/md5.py
	@mkdir -p $(@D)
	@pydeps -modname=md5 $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/md5.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/md5.a: build/src/__python__/md5/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/md5 -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/md5.a

-include build/src/__python__/md5/module.d

build/src/__python__/mimetools/module.go: build/src/__python__/mimetools.py
	@mkdir -p $(@D)
	@grumpc -modname=mimetools $< > $@

build/src/__python__/mimetools/module.d: build/src/__python__/mimetools.py
	@mkdir -p $(@D)
	@pydeps -modname=mimetools $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/mimetools.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/mimetools.a: build/src/__python__/mimetools/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/mimetools -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/mimetools.a

-include build/src/__python__/mimetools/module.d

build/src/__python__/mutex/module.go: build/src/__python__/mutex.py
	@mkdir -p $(@D)
	@grumpc -modname=mutex $< > $@

build/src/__python__/mutex/module.d: build/src/__python__/mutex.py
	@mkdir -p $(@D)
	@pydeps -modname=mutex $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/mutex.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/mutex.a: build/src/__python__/mutex/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/mutex -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/mutex.a

-include build/src/__python__/mutex/module.d

build/src/__python__/operator/module.go: build/src/__python__/operator.py
	@mkdir -p $(@D)
	@grumpc -modname=operator $< > $@

build/src/__python__/operator/module.d: build/src/__python__/operator.py
	@mkdir -p $(@D)
	@pydeps -modname=operator $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/operator.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/operator.a: build/src/__python__/operator/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/operator -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/operator.a

-include build/src/__python__/operator/module.d

build/src/__python__/optparse/module.go: build/src/__python__/optparse.py
	@mkdir -p $(@D)
	@grumpc -modname=optparse $< > $@

build/src/__python__/optparse/module.d: build/src/__python__/optparse.py
	@mkdir -p $(@D)
	@pydeps -modname=optparse $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/optparse.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/optparse.a: build/src/__python__/optparse/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/optparse -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/optparse.a

-include build/src/__python__/optparse/module.d

build/src/__python__/os_test/module.go: build/src/__python__/os_test.py
	@mkdir -p $(@D)
	@grumpc -modname=os_test $< > $@

build/src/__python__/os_test/module.d: build/src/__python__/os_test.py
	@mkdir -p $(@D)
	@pydeps -modname=os_test $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/os_test.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/os_test.a: build/src/__python__/os_test/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/os_test -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/os_test.a

-include build/src/__python__/os_test/module.d

build/src/__python__/pprint/module.go: build/src/__python__/pprint.py
	@mkdir -p $(@D)
	@grumpc -modname=pprint $< > $@

build/src/__python__/pprint/module.d: build/src/__python__/pprint.py
	@mkdir -p $(@D)
	@pydeps -modname=pprint $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/pprint.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/pprint.a: build/src/__python__/pprint/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/pprint -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/pprint.a

-include build/src/__python__/pprint/module.d

build/src/__python__/Queue/module.go: build/src/__python__/Queue.py
	@mkdir -p $(@D)
	@grumpc -modname=Queue $< > $@

build/src/__python__/Queue/module.d: build/src/__python__/Queue.py
	@mkdir -p $(@D)
	@pydeps -modname=Queue $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/Queue.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/Queue.a: build/src/__python__/Queue/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/Queue -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/Queue.a

-include build/src/__python__/Queue/module.d

build/src/__python__/quopri/module.go: build/src/__python__/quopri.py
	@mkdir -p $(@D)
	@grumpc -modname=quopri $< > $@

build/src/__python__/quopri/module.d: build/src/__python__/quopri.py
	@mkdir -p $(@D)
	@pydeps -modname=quopri $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/quopri.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/quopri.a: build/src/__python__/quopri/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/quopri -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/quopri.a

-include build/src/__python__/quopri/module.d

build/src/__python__/random/module.go: build/src/__python__/random.py
	@mkdir -p $(@D)
	@grumpc -modname=random $< > $@

build/src/__python__/random/module.d: build/src/__python__/random.py
	@mkdir -p $(@D)
	@pydeps -modname=random $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/random.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/random.a: build/src/__python__/random/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/random -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/random.a

-include build/src/__python__/random/module.d

build/src/__python__/random_test/module.go: build/src/__python__/random_test.py
	@mkdir -p $(@D)
	@grumpc -modname=random_test $< > $@

build/src/__python__/random_test/module.d: build/src/__python__/random_test.py
	@mkdir -p $(@D)
	@pydeps -modname=random_test $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/random_test.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/random_test.a: build/src/__python__/random_test/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/random_test -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/random_test.a

-include build/src/__python__/random_test/module.d

build/src/__python__/re/module.go: build/src/__python__/re.py
	@mkdir -p $(@D)
	@grumpc -modname=re $< > $@

build/src/__python__/re/module.d: build/src/__python__/re.py
	@mkdir -p $(@D)
	@pydeps -modname=re $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/re.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/re.a: build/src/__python__/re/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/re -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/re.a

-include build/src/__python__/re/module.d

build/src/__python__/re_tests/module.go: build/src/__python__/re_tests.py
	@mkdir -p $(@D)
	@grumpc -modname=re_tests $< > $@

build/src/__python__/re_tests/module.d: build/src/__python__/re_tests.py
	@mkdir -p $(@D)
	@pydeps -modname=re_tests $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/re_tests.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/re_tests.a: build/src/__python__/re_tests/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/re_tests -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/re_tests.a

-include build/src/__python__/re_tests/module.d

build/src/__python__/repr/module.go: build/src/__python__/repr.py
	@mkdir -p $(@D)
	@grumpc -modname=repr $< > $@

build/src/__python__/repr/module.d: build/src/__python__/repr.py
	@mkdir -p $(@D)
	@pydeps -modname=repr $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/repr.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/repr.a: build/src/__python__/repr/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/repr -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/repr.a

-include build/src/__python__/repr/module.d

build/src/__python__/rfc822/module.go: build/src/__python__/rfc822.py
	@mkdir -p $(@D)
	@grumpc -modname=rfc822 $< > $@

build/src/__python__/rfc822/module.d: build/src/__python__/rfc822.py
	@mkdir -p $(@D)
	@pydeps -modname=rfc822 $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/rfc822.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/rfc822.a: build/src/__python__/rfc822/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/rfc822 -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/rfc822.a

-include build/src/__python__/rfc822/module.d

build/src/__python__/sched/module.go: build/src/__python__/sched.py
	@mkdir -p $(@D)
	@grumpc -modname=sched $< > $@

build/src/__python__/sched/module.d: build/src/__python__/sched.py
	@mkdir -p $(@D)
	@pydeps -modname=sched $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/sched.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/sched.a: build/src/__python__/sched/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/sched -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/sched.a

-include build/src/__python__/sched/module.d

build/src/__python__/select_/module.go: build/src/__python__/select_.py
	@mkdir -p $(@D)
	@grumpc -modname=select_ $< > $@

build/src/__python__/select_/module.d: build/src/__python__/select_.py
	@mkdir -p $(@D)
	@pydeps -modname=select_ $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/select_.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/select_.a: build/src/__python__/select_/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/select_ -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/select_.a

-include build/src/__python__/select_/module.d

build/src/__python__/sha/module.go: build/src/__python__/sha.py
	@mkdir -p $(@D)
	@grumpc -modname=sha $< > $@

build/src/__python__/sha/module.d: build/src/__python__/sha.py
	@mkdir -p $(@D)
	@pydeps -modname=sha $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/sha.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/sha.a: build/src/__python__/sha/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/sha -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/sha.a

-include build/src/__python__/sha/module.d

build/src/__python__/sre_compile/module.go: build/src/__python__/sre_compile.py
	@mkdir -p $(@D)
	@grumpc -modname=sre_compile $< > $@

build/src/__python__/sre_compile/module.d: build/src/__python__/sre_compile.py
	@mkdir -p $(@D)
	@pydeps -modname=sre_compile $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/sre_compile.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/sre_compile.a: build/src/__python__/sre_compile/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/sre_compile -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/sre_compile.a

-include build/src/__python__/sre_compile/module.d

build/src/__python__/sre_constants/module.go: build/src/__python__/sre_constants.py
	@mkdir -p $(@D)
	@grumpc -modname=sre_constants $< > $@

build/src/__python__/sre_constants/module.d: build/src/__python__/sre_constants.py
	@mkdir -p $(@D)
	@pydeps -modname=sre_constants $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/sre_constants.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/sre_constants.a: build/src/__python__/sre_constants/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/sre_constants -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/sre_constants.a

-include build/src/__python__/sre_constants/module.d

build/src/__python__/sre_parse/module.go: build/src/__python__/sre_parse.py
	@mkdir -p $(@D)
	@grumpc -modname=sre_parse $< > $@

build/src/__python__/sre_parse/module.d: build/src/__python__/sre_parse.py
	@mkdir -p $(@D)
	@pydeps -modname=sre_parse $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/sre_parse.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/sre_parse.a: build/src/__python__/sre_parse/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/sre_parse -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/sre_parse.a

-include build/src/__python__/sre_parse/module.d

build/src/__python__/stat/module.go: build/src/__python__/stat.py
	@mkdir -p $(@D)
	@grumpc -modname=stat $< > $@

build/src/__python__/stat/module.d: build/src/__python__/stat.py
	@mkdir -p $(@D)
	@pydeps -modname=stat $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/stat.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/stat.a: build/src/__python__/stat/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/stat -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/stat.a

-include build/src/__python__/stat/module.d

build/src/__python__/string/module.go: build/src/__python__/string.py
	@mkdir -p $(@D)
	@grumpc -modname=string $< > $@

build/src/__python__/string/module.d: build/src/__python__/string.py
	@mkdir -p $(@D)
	@pydeps -modname=string $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/string.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/string.a: build/src/__python__/string/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/string -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/string.a

-include build/src/__python__/string/module.d

build/src/__python__/StringIO/module.go: build/src/__python__/StringIO.py
	@mkdir -p $(@D)
	@grumpc -modname=StringIO $< > $@

build/src/__python__/StringIO/module.d: build/src/__python__/StringIO.py
	@mkdir -p $(@D)
	@pydeps -modname=StringIO $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/StringIO.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/StringIO.a: build/src/__python__/StringIO/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/StringIO -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/StringIO.a

-include build/src/__python__/StringIO/module.d

build/src/__python__/sys/module.go: build/src/__python__/sys.py
	@mkdir -p $(@D)
	@grumpc -modname=sys $< > $@

build/src/__python__/sys/module.d: build/src/__python__/sys.py
	@mkdir -p $(@D)
	@pydeps -modname=sys $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/sys.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/sys.a: build/src/__python__/sys/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/sys -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/sys.a

-include build/src/__python__/sys/module.d

build/src/__python__/sys_test/module.go: build/src/__python__/sys_test.py
	@mkdir -p $(@D)
	@grumpc -modname=sys_test $< > $@

build/src/__python__/sys_test/module.d: build/src/__python__/sys_test.py
	@mkdir -p $(@D)
	@pydeps -modname=sys_test $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/sys_test.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/sys_test.a: build/src/__python__/sys_test/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/sys_test -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/sys_test.a

-include build/src/__python__/sys_test/module.d

build/src/__python__/tempfile/module.go: build/src/__python__/tempfile.py
	@mkdir -p $(@D)
	@grumpc -modname=tempfile $< > $@

build/src/__python__/tempfile/module.d: build/src/__python__/tempfile.py
	@mkdir -p $(@D)
	@pydeps -modname=tempfile $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/tempfile.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/tempfile.a: build/src/__python__/tempfile/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/tempfile -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/tempfile.a

-include build/src/__python__/tempfile/module.d

build/src/__python__/tempfile_test/module.go: build/src/__python__/tempfile_test.py
	@mkdir -p $(@D)
	@grumpc -modname=tempfile_test $< > $@

build/src/__python__/tempfile_test/module.d: build/src/__python__/tempfile_test.py
	@mkdir -p $(@D)
	@pydeps -modname=tempfile_test $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/tempfile_test.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/tempfile_test.a: build/src/__python__/tempfile_test/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/tempfile_test -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/tempfile_test.a

-include build/src/__python__/tempfile_test/module.d

build/src/__python__/textwrap/module.go: build/src/__python__/textwrap.py
	@mkdir -p $(@D)
	@grumpc -modname=textwrap $< > $@

build/src/__python__/textwrap/module.d: build/src/__python__/textwrap.py
	@mkdir -p $(@D)
	@pydeps -modname=textwrap $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/textwrap.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/textwrap.a: build/src/__python__/textwrap/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/textwrap -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/textwrap.a

-include build/src/__python__/textwrap/module.d

build/src/__python__/thread/module.go: build/src/__python__/thread.py
	@mkdir -p $(@D)
	@grumpc -modname=thread $< > $@

build/src/__python__/thread/module.d: build/src/__python__/thread.py
	@mkdir -p $(@D)
	@pydeps -modname=thread $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/thread.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/thread.a: build/src/__python__/thread/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/thread -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/thread.a

-include build/src/__python__/thread/module.d

build/src/__python__/threading/module.go: build/src/__python__/threading.py
	@mkdir -p $(@D)
	@grumpc -modname=threading $< > $@

build/src/__python__/threading/module.d: build/src/__python__/threading.py
	@mkdir -p $(@D)
	@pydeps -modname=threading $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/threading.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/threading.a: build/src/__python__/threading/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/threading -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/threading.a

-include build/src/__python__/threading/module.d

build/src/__python__/time/module.go: build/src/__python__/time.py
	@mkdir -p $(@D)
	@grumpc -modname=time $< > $@

build/src/__python__/time/module.d: build/src/__python__/time.py
	@mkdir -p $(@D)
	@pydeps -modname=time $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/time.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/time.a: build/src/__python__/time/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/time -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/time.a

-include build/src/__python__/time/module.d

build/src/__python__/time_test/module.go: build/src/__python__/time_test.py
	@mkdir -p $(@D)
	@grumpc -modname=time_test $< > $@

build/src/__python__/time_test/module.d: build/src/__python__/time_test.py
	@mkdir -p $(@D)
	@pydeps -modname=time_test $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/time_test.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/time_test.a: build/src/__python__/time_test/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/time_test -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/time_test.a

-include build/src/__python__/time_test/module.d

build/src/__python__/traceback/module.go: build/src/__python__/traceback.py
	@mkdir -p $(@D)
	@grumpc -modname=traceback $< > $@

build/src/__python__/traceback/module.d: build/src/__python__/traceback.py
	@mkdir -p $(@D)
	@pydeps -modname=traceback $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/traceback.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/traceback.a: build/src/__python__/traceback/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/traceback -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/traceback.a

-include build/src/__python__/traceback/module.d

build/src/__python__/types/module.go: build/src/__python__/types.py
	@mkdir -p $(@D)
	@grumpc -modname=types $< > $@

build/src/__python__/types/module.d: build/src/__python__/types.py
	@mkdir -p $(@D)
	@pydeps -modname=types $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/types.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/types.a: build/src/__python__/types/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/types -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/types.a

-include build/src/__python__/types/module.d

build/src/__python__/types_test/module.go: build/src/__python__/types_test.py
	@mkdir -p $(@D)
	@grumpc -modname=types_test $< > $@

build/src/__python__/types_test/module.d: build/src/__python__/types_test.py
	@mkdir -p $(@D)
	@pydeps -modname=types_test $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/types_test.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/types_test.a: build/src/__python__/types_test/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/types_test -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/types_test.a

-include build/src/__python__/types_test/module.d

build/src/__python__/unittest_case/module.go: build/src/__python__/unittest_case.py
	@mkdir -p $(@D)
	@grumpc -modname=unittest_case $< > $@

build/src/__python__/unittest_case/module.d: build/src/__python__/unittest_case.py
	@mkdir -p $(@D)
	@pydeps -modname=unittest_case $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/unittest_case.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/unittest_case.a: build/src/__python__/unittest_case/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/unittest_case -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/unittest_case.a

-include build/src/__python__/unittest_case/module.d

build/src/__python__/unittest_loader/module.go: build/src/__python__/unittest_loader.py
	@mkdir -p $(@D)
	@grumpc -modname=unittest_loader $< > $@

build/src/__python__/unittest_loader/module.d: build/src/__python__/unittest_loader.py
	@mkdir -p $(@D)
	@pydeps -modname=unittest_loader $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/unittest_loader.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/unittest_loader.a: build/src/__python__/unittest_loader/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/unittest_loader -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/unittest_loader.a

-include build/src/__python__/unittest_loader/module.d

build/src/__python__/unittest_result/module.go: build/src/__python__/unittest_result.py
	@mkdir -p $(@D)
	@grumpc -modname=unittest_result $< > $@

build/src/__python__/unittest_result/module.d: build/src/__python__/unittest_result.py
	@mkdir -p $(@D)
	@pydeps -modname=unittest_result $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/unittest_result.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/unittest_result.a: build/src/__python__/unittest_result/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/unittest_result -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/unittest_result.a

-include build/src/__python__/unittest_result/module.d

build/src/__python__/unittest_runner/module.go: build/src/__python__/unittest_runner.py
	@mkdir -p $(@D)
	@grumpc -modname=unittest_runner $< > $@

build/src/__python__/unittest_runner/module.d: build/src/__python__/unittest_runner.py
	@mkdir -p $(@D)
	@pydeps -modname=unittest_runner $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/unittest_runner.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/unittest_runner.a: build/src/__python__/unittest_runner/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/unittest_runner -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/unittest_runner.a

-include build/src/__python__/unittest_runner/module.d

build/src/__python__/unittest_signals/module.go: build/src/__python__/unittest_signals.py
	@mkdir -p $(@D)
	@grumpc -modname=unittest_signals $< > $@

build/src/__python__/unittest_signals/module.d: build/src/__python__/unittest_signals.py
	@mkdir -p $(@D)
	@pydeps -modname=unittest_signals $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/unittest_signals.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/unittest_signals.a: build/src/__python__/unittest_signals/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/unittest_signals -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/unittest_signals.a

-include build/src/__python__/unittest_signals/module.d

build/src/__python__/unittest_suite/module.go: build/src/__python__/unittest_suite.py
	@mkdir -p $(@D)
	@grumpc -modname=unittest_suite $< > $@

build/src/__python__/unittest_suite/module.d: build/src/__python__/unittest_suite.py
	@mkdir -p $(@D)
	@pydeps -modname=unittest_suite $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/unittest_suite.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/unittest_suite.a: build/src/__python__/unittest_suite/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/unittest_suite -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/unittest_suite.a

-include build/src/__python__/unittest_suite/module.d

build/src/__python__/unittest_util/module.go: build/src/__python__/unittest_util.py
	@mkdir -p $(@D)
	@grumpc -modname=unittest_util $< > $@

build/src/__python__/unittest_util/module.d: build/src/__python__/unittest_util.py
	@mkdir -p $(@D)
	@pydeps -modname=unittest_util $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/unittest_util.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/unittest_util.a: build/src/__python__/unittest_util/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/unittest_util -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/unittest_util.a

-include build/src/__python__/unittest_util/module.d

build/src/__python__/urlparse/module.go: build/src/__python__/urlparse.py
	@mkdir -p $(@D)
	@grumpc -modname=urlparse $< > $@

build/src/__python__/urlparse/module.d: build/src/__python__/urlparse.py
	@mkdir -p $(@D)
	@pydeps -modname=urlparse $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/urlparse.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/urlparse.a: build/src/__python__/urlparse/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/urlparse -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/urlparse.a

-include build/src/__python__/urlparse/module.d

build/src/__python__/UserDict/module.go: build/src/__python__/UserDict.py
	@mkdir -p $(@D)
	@grumpc -modname=UserDict $< > $@

build/src/__python__/UserDict/module.d: build/src/__python__/UserDict.py
	@mkdir -p $(@D)
	@pydeps -modname=UserDict $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/UserDict.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/UserDict.a: build/src/__python__/UserDict/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/UserDict -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/UserDict.a

-include build/src/__python__/UserDict/module.d

build/src/__python__/UserList/module.go: build/src/__python__/UserList.py
	@mkdir -p $(@D)
	@grumpc -modname=UserList $< > $@

build/src/__python__/UserList/module.d: build/src/__python__/UserList.py
	@mkdir -p $(@D)
	@pydeps -modname=UserList $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/UserList.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/UserList.a: build/src/__python__/UserList/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/UserList -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/UserList.a

-include build/src/__python__/UserList/module.d

build/src/__python__/UserString/module.go: build/src/__python__/UserString.py
	@mkdir -p $(@D)
	@grumpc -modname=UserString $< > $@

build/src/__python__/UserString/module.d: build/src/__python__/UserString.py
	@mkdir -p $(@D)
	@pydeps -modname=UserString $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/UserString.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/UserString.a: build/src/__python__/UserString/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/UserString -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/UserString.a

-include build/src/__python__/UserString/module.d

build/src/__python__/uu/module.go: build/src/__python__/uu.py
	@mkdir -p $(@D)
	@grumpc -modname=uu $< > $@

build/src/__python__/uu/module.d: build/src/__python__/uu.py
	@mkdir -p $(@D)
	@pydeps -modname=uu $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/uu.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/uu.a: build/src/__python__/uu/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/uu -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/uu.a

-include build/src/__python__/uu/module.d

build/src/__python__/warnings/module.go: build/src/__python__/warnings.py
	@mkdir -p $(@D)
	@grumpc -modname=warnings $< > $@

build/src/__python__/warnings/module.d: build/src/__python__/warnings.py
	@mkdir -p $(@D)
	@pydeps -modname=warnings $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/warnings.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/warnings.a: build/src/__python__/warnings/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/warnings -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/warnings.a

-include build/src/__python__/warnings/module.d

build/src/__python__/weakref/module.go: build/src/__python__/weakref.py
	@mkdir -p $(@D)
	@grumpc -modname=weakref $< > $@

build/src/__python__/weakref/module.d: build/src/__python__/weakref.py
	@mkdir -p $(@D)
	@pydeps -modname=weakref $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/weakref.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/weakref.a: build/src/__python__/weakref/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/weakref -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/weakref.a

-include build/src/__python__/weakref/module.d

build/src/__python__/weetest/module.go: build/src/__python__/weetest.py
	@mkdir -p $(@D)
	@grumpc -modname=weetest $< > $@

build/src/__python__/weetest/module.d: build/src/__python__/weetest.py
	@mkdir -p $(@D)
	@pydeps -modname=weetest $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/weetest.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/weetest.a: build/src/__python__/weetest/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/weetest -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/weetest.a

-include build/src/__python__/weetest/module.d

build/src/__python__/weetest_test/module.go: build/src/__python__/weetest_test.py
	@mkdir -p $(@D)
	@grumpc -modname=weetest_test $< > $@

build/src/__python__/weetest_test/module.d: build/src/__python__/weetest_test.py
	@mkdir -p $(@D)
	@pydeps -modname=weetest_test $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/weetest_test.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/weetest_test.a: build/src/__python__/weetest_test/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/weetest_test -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/weetest_test.a

-include build/src/__python__/weetest_test/module.d

build/src/__python__/json/module.go: build/src/__python__/json/__init__.py
	@mkdir -p $(@D)
	@grumpc -modname=json $< > $@

build/src/__python__/json/module.d: build/src/__python__/json/__init__.py
	@mkdir -p $(@D)
	@pydeps -modname=json $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/json.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/json.a: build/src/__python__/json/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/json -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/json.a

-include build/src/__python__/json/module.d

build/src/__python__/json/decoder/module.go: build/src/__python__/json/decoder.py
	@mkdir -p $(@D)
	@grumpc -modname=json.decoder $< > $@

build/src/__python__/json/decoder/module.d: build/src/__python__/json/decoder.py
	@mkdir -p $(@D)
	@pydeps -modname=json.decoder $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/json/decoder.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/json/decoder.a: build/src/__python__/json/decoder/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/json/decoder -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/json/decoder.a

-include build/src/__python__/json/decoder/module.d

build/src/__python__/json/encoder/module.go: build/src/__python__/json/encoder.py
	@mkdir -p $(@D)
	@grumpc -modname=json.encoder $< > $@

build/src/__python__/json/encoder/module.d: build/src/__python__/json/encoder.py
	@mkdir -p $(@D)
	@pydeps -modname=json.encoder $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/json/encoder.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/json/encoder.a: build/src/__python__/json/encoder/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/json/encoder -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/json/encoder.a

-include build/src/__python__/json/encoder/module.d

build/src/__python__/os/module.go: build/src/__python__/os/__init__.py
	@mkdir -p $(@D)
	@grumpc -modname=os $< > $@

build/src/__python__/os/module.d: build/src/__python__/os/__init__.py
	@mkdir -p $(@D)
	@pydeps -modname=os $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/os.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/os.a: build/src/__python__/os/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/os -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/os.a

-include build/src/__python__/os/module.d

build/src/__python__/os/path/module.go: build/src/__python__/os/path.py
	@mkdir -p $(@D)
	@grumpc -modname=os.path $< > $@

build/src/__python__/os/path/module.d: build/src/__python__/os/path.py
	@mkdir -p $(@D)
	@pydeps -modname=os.path $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/os/path.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/os/path.a: build/src/__python__/os/path/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/os/path -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/os/path.a

-include build/src/__python__/os/path/module.d

build/src/__python__/os/path_test/module.go: build/src/__python__/os/path_test.py
	@mkdir -p $(@D)
	@grumpc -modname=os.path_test $< > $@

build/src/__python__/os/path_test/module.d: build/src/__python__/os/path_test.py
	@mkdir -p $(@D)
	@pydeps -modname=os.path_test $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/os/path_test.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/os/path_test.a: build/src/__python__/os/path_test/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/os/path_test -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/os/path_test.a

-include build/src/__python__/os/path_test/module.d

build/src/__python__/test/module.go: build/src/__python__/test/__init__.py
	@mkdir -p $(@D)
	@grumpc -modname=test $< > $@

build/src/__python__/test/module.d: build/src/__python__/test/__init__.py
	@mkdir -p $(@D)
	@pydeps -modname=test $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test.a: build/src/__python__/test/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test.a

-include build/src/__python__/test/module.d

build/src/__python__/test/list_tests/module.go: build/src/__python__/test/list_tests.py
	@mkdir -p $(@D)
	@grumpc -modname=test.list_tests $< > $@

build/src/__python__/test/list_tests/module.d: build/src/__python__/test/list_tests.py
	@mkdir -p $(@D)
	@pydeps -modname=test.list_tests $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/list_tests.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/list_tests.a: build/src/__python__/test/list_tests/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/list_tests -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/list_tests.a

-include build/src/__python__/test/list_tests/module.d

build/src/__python__/test/lock_tests/module.go: build/src/__python__/test/lock_tests.py
	@mkdir -p $(@D)
	@grumpc -modname=test.lock_tests $< > $@

build/src/__python__/test/lock_tests/module.d: build/src/__python__/test/lock_tests.py
	@mkdir -p $(@D)
	@pydeps -modname=test.lock_tests $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/lock_tests.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/lock_tests.a: build/src/__python__/test/lock_tests/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/lock_tests -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/lock_tests.a

-include build/src/__python__/test/lock_tests/module.d

build/src/__python__/test/mapping_tests/module.go: build/src/__python__/test/mapping_tests.py
	@mkdir -p $(@D)
	@grumpc -modname=test.mapping_tests $< > $@

build/src/__python__/test/mapping_tests/module.d: build/src/__python__/test/mapping_tests.py
	@mkdir -p $(@D)
	@pydeps -modname=test.mapping_tests $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/mapping_tests.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/mapping_tests.a: build/src/__python__/test/mapping_tests/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/mapping_tests -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/mapping_tests.a

-include build/src/__python__/test/mapping_tests/module.d

build/src/__python__/test/seq_tests/module.go: build/src/__python__/test/seq_tests.py
	@mkdir -p $(@D)
	@grumpc -modname=test.seq_tests $< > $@

build/src/__python__/test/seq_tests/module.d: build/src/__python__/test/seq_tests.py
	@mkdir -p $(@D)
	@pydeps -modname=test.seq_tests $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/seq_tests.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/seq_tests.a: build/src/__python__/test/seq_tests/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/seq_tests -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/seq_tests.a

-include build/src/__python__/test/seq_tests/module.d

build/src/__python__/test/string_tests/module.go: build/src/__python__/test/string_tests.py
	@mkdir -p $(@D)
	@grumpc -modname=test.string_tests $< > $@

build/src/__python__/test/string_tests/module.d: build/src/__python__/test/string_tests.py
	@mkdir -p $(@D)
	@pydeps -modname=test.string_tests $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/string_tests.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/string_tests.a: build/src/__python__/test/string_tests/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/string_tests -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/string_tests.a

-include build/src/__python__/test/string_tests/module.d

build/src/__python__/test/test_argparse/module.go: build/src/__python__/test/test_argparse.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_argparse $< > $@

build/src/__python__/test/test_argparse/module.d: build/src/__python__/test/test_argparse.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_argparse $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_argparse.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_argparse.a: build/src/__python__/test/test_argparse/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_argparse -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_argparse.a

-include build/src/__python__/test/test_argparse/module.d

build/src/__python__/test/test_bisect/module.go: build/src/__python__/test/test_bisect.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_bisect $< > $@

build/src/__python__/test/test_bisect/module.d: build/src/__python__/test/test_bisect.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_bisect $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_bisect.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_bisect.a: build/src/__python__/test/test_bisect/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_bisect -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_bisect.a

-include build/src/__python__/test/test_bisect/module.d

build/src/__python__/test/test_colorsys/module.go: build/src/__python__/test/test_colorsys.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_colorsys $< > $@

build/src/__python__/test/test_colorsys/module.d: build/src/__python__/test/test_colorsys.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_colorsys $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_colorsys.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_colorsys.a: build/src/__python__/test/test_colorsys/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_colorsys -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_colorsys.a

-include build/src/__python__/test/test_colorsys/module.d

build/src/__python__/test/test_datetime/module.go: build/src/__python__/test/test_datetime.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_datetime $< > $@

build/src/__python__/test/test_datetime/module.d: build/src/__python__/test/test_datetime.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_datetime $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_datetime.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_datetime.a: build/src/__python__/test/test_datetime/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_datetime -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_datetime.a

-include build/src/__python__/test/test_datetime/module.d

build/src/__python__/test/test_dict/module.go: build/src/__python__/test/test_dict.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_dict $< > $@

build/src/__python__/test/test_dict/module.d: build/src/__python__/test/test_dict.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_dict $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_dict.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_dict.a: build/src/__python__/test/test_dict/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_dict -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_dict.a

-include build/src/__python__/test/test_dict/module.d

build/src/__python__/test/test_dircache/module.go: build/src/__python__/test/test_dircache.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_dircache $< > $@

build/src/__python__/test/test_dircache/module.d: build/src/__python__/test/test_dircache.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_dircache $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_dircache.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_dircache.a: build/src/__python__/test/test_dircache/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_dircache -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_dircache.a

-include build/src/__python__/test/test_dircache/module.d

build/src/__python__/test/test_dummy_thread/module.go: build/src/__python__/test/test_dummy_thread.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_dummy_thread $< > $@

build/src/__python__/test/test_dummy_thread/module.d: build/src/__python__/test/test_dummy_thread.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_dummy_thread $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_dummy_thread.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_dummy_thread.a: build/src/__python__/test/test_dummy_thread/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_dummy_thread -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_dummy_thread.a

-include build/src/__python__/test/test_dummy_thread/module.d

build/src/__python__/test/test_fpformat/module.go: build/src/__python__/test/test_fpformat.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_fpformat $< > $@

build/src/__python__/test/test_fpformat/module.d: build/src/__python__/test/test_fpformat.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_fpformat $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_fpformat.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_fpformat.a: build/src/__python__/test/test_fpformat/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_fpformat -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_fpformat.a

-include build/src/__python__/test/test_fpformat/module.d

build/src/__python__/test/test_genericpath/module.go: build/src/__python__/test/test_genericpath.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_genericpath $< > $@

build/src/__python__/test/test_genericpath/module.d: build/src/__python__/test/test_genericpath.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_genericpath $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_genericpath.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_genericpath.a: build/src/__python__/test/test_genericpath/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_genericpath -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_genericpath.a

-include build/src/__python__/test/test_genericpath/module.d

build/src/__python__/test/test_list/module.go: build/src/__python__/test/test_list.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_list $< > $@

build/src/__python__/test/test_list/module.d: build/src/__python__/test/test_list.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_list $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_list.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_list.a: build/src/__python__/test/test_list/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_list -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_list.a

-include build/src/__python__/test/test_list/module.d

build/src/__python__/test/test_md5/module.go: build/src/__python__/test/test_md5.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_md5 $< > $@

build/src/__python__/test/test_md5/module.d: build/src/__python__/test/test_md5.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_md5 $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_md5.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_md5.a: build/src/__python__/test/test_md5/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_md5 -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_md5.a

-include build/src/__python__/test/test_md5/module.d

build/src/__python__/test/test_mimetools/module.go: build/src/__python__/test/test_mimetools.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_mimetools $< > $@

build/src/__python__/test/test_mimetools/module.d: build/src/__python__/test/test_mimetools.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_mimetools $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_mimetools.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_mimetools.a: build/src/__python__/test/test_mimetools/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_mimetools -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_mimetools.a

-include build/src/__python__/test/test_mimetools/module.d

build/src/__python__/test/test_mutex/module.go: build/src/__python__/test/test_mutex.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_mutex $< > $@

build/src/__python__/test/test_mutex/module.d: build/src/__python__/test/test_mutex.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_mutex $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_mutex.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_mutex.a: build/src/__python__/test/test_mutex/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_mutex -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_mutex.a

-include build/src/__python__/test/test_mutex/module.d

build/src/__python__/test/test_operator/module.go: build/src/__python__/test/test_operator.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_operator $< > $@

build/src/__python__/test/test_operator/module.d: build/src/__python__/test/test_operator.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_operator $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_operator.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_operator.a: build/src/__python__/test/test_operator/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_operator -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_operator.a

-include build/src/__python__/test/test_operator/module.d

build/src/__python__/test/test_queue/module.go: build/src/__python__/test/test_queue.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_queue $< > $@

build/src/__python__/test/test_queue/module.d: build/src/__python__/test/test_queue.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_queue $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_queue.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_queue.a: build/src/__python__/test/test_queue/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_queue -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_queue.a

-include build/src/__python__/test/test_queue/module.d

build/src/__python__/test/test_quopri/module.go: build/src/__python__/test/test_quopri.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_quopri $< > $@

build/src/__python__/test/test_quopri/module.d: build/src/__python__/test/test_quopri.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_quopri $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_quopri.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_quopri.a: build/src/__python__/test/test_quopri/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_quopri -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_quopri.a

-include build/src/__python__/test/test_quopri/module.d

build/src/__python__/test/test_rfc822/module.go: build/src/__python__/test/test_rfc822.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_rfc822 $< > $@

build/src/__python__/test/test_rfc822/module.d: build/src/__python__/test/test_rfc822.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_rfc822 $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_rfc822.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_rfc822.a: build/src/__python__/test/test_rfc822/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_rfc822 -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_rfc822.a

-include build/src/__python__/test/test_rfc822/module.d

build/src/__python__/test/test_sched/module.go: build/src/__python__/test/test_sched.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_sched $< > $@

build/src/__python__/test/test_sched/module.d: build/src/__python__/test/test_sched.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_sched $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_sched.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_sched.a: build/src/__python__/test/test_sched/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_sched -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_sched.a

-include build/src/__python__/test/test_sched/module.d

build/src/__python__/test/test_select/module.go: build/src/__python__/test/test_select.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_select $< > $@

build/src/__python__/test/test_select/module.d: build/src/__python__/test/test_select.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_select $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_select.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_select.a: build/src/__python__/test/test_select/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_select -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_select.a

-include build/src/__python__/test/test_select/module.d

build/src/__python__/test/test_slice/module.go: build/src/__python__/test/test_slice.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_slice $< > $@

build/src/__python__/test/test_slice/module.d: build/src/__python__/test/test_slice.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_slice $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_slice.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_slice.a: build/src/__python__/test/test_slice/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_slice -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_slice.a

-include build/src/__python__/test/test_slice/module.d

build/src/__python__/test/test_stat/module.go: build/src/__python__/test/test_stat.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_stat $< > $@

build/src/__python__/test/test_stat/module.d: build/src/__python__/test/test_stat.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_stat $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_stat.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_stat.a: build/src/__python__/test/test_stat/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_stat -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_stat.a

-include build/src/__python__/test/test_stat/module.d

build/src/__python__/test/test_string/module.go: build/src/__python__/test/test_string.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_string $< > $@

build/src/__python__/test/test_string/module.d: build/src/__python__/test/test_string.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_string $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_string.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_string.a: build/src/__python__/test/test_string/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_string -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_string.a

-include build/src/__python__/test/test_string/module.d

build/src/__python__/test/test_support/module.go: build/src/__python__/test/test_support.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_support $< > $@

build/src/__python__/test/test_support/module.d: build/src/__python__/test/test_support.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_support $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_support.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_support.a: build/src/__python__/test/test_support/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_support -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_support.a

-include build/src/__python__/test/test_support/module.d

build/src/__python__/test/test_threading/module.go: build/src/__python__/test/test_threading.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_threading $< > $@

build/src/__python__/test/test_threading/module.d: build/src/__python__/test/test_threading.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_threading $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_threading.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_threading.a: build/src/__python__/test/test_threading/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_threading -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_threading.a

-include build/src/__python__/test/test_threading/module.d

build/src/__python__/test/test_tuple/module.go: build/src/__python__/test/test_tuple.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_tuple $< > $@

build/src/__python__/test/test_tuple/module.d: build/src/__python__/test/test_tuple.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_tuple $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_tuple.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_tuple.a: build/src/__python__/test/test_tuple/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_tuple -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_tuple.a

-include build/src/__python__/test/test_tuple/module.d

build/src/__python__/test/test_uu/module.go: build/src/__python__/test/test_uu.py
	@mkdir -p $(@D)
	@grumpc -modname=test.test_uu $< > $@

build/src/__python__/test/test_uu/module.d: build/src/__python__/test/test_uu.py
	@mkdir -p $(@D)
	@pydeps -modname=test.test_uu $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/test/test_uu.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/test/test_uu.a: build/src/__python__/test/test_uu/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/test/test_uu -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/test/test_uu.a

-include build/src/__python__/test/test_uu/module.d

build/src/__python__/unittest/module.go: build/src/__python__/unittest/__init__.py
	@mkdir -p $(@D)
	@grumpc -modname=unittest $< > $@

build/src/__python__/unittest/module.d: build/src/__python__/unittest/__init__.py
	@mkdir -p $(@D)
	@pydeps -modname=unittest $< | awk '{gsub(/\./, "/", $$0); print "build/pkg/darwin_amd64/__python__/unittest.a: build/pkg/darwin_amd64/__python__/" $$0 ".a"}' > $@

build/pkg/darwin_amd64/__python__/unittest.a: build/src/__python__/unittest/module.go
	@mkdir -p $(@D)
	@go tool compile -o $@ -p __python__/unittest -complete -I build/pkg/darwin_amd64 -pack $<

all: build/pkg/darwin_amd64/__python__/unittest.a

-include build/src/__python__/unittest/module.d

