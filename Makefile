.PHONY: gen
gen:
	buf lint
	buf format
	buf generate
