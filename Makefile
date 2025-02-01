.PHONY: gen
gen:
	buf lint
	buf generate
