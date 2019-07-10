git-hash:=$(shell git describe --always --match=NeVeRmAtCh --dirty 2>/dev/null)

$(shell git config --global url."https://9470a8d7c1f3262393ba4afdd2ec1a9d9134724b:x-oauth-basic@github.com/ipweb-group/go-ipws-config".insteadOf "https://github.com/ipweb-group/go-ipws-config")
$(shell git config --global url."https://9470a8d7c1f3262393ba4afdd2ec1a9d9134724b:x-oauth-basic@github.com/ipweb-group/interface-go-ipws-core".insteadOf "https://github.com/ipweb-group/interface-go-ipws-core")
$(shell git config --global url."https://9470a8d7c1f3262393ba4afdd2ec1a9d9134724b:x-oauth-basic@github.com/ipweb-group/ipw".insteadOf "https://github.com/ipweb-group/ipw")
