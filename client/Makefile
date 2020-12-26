OUT_DIR ?= dist
OUT_JS ?= ${OUT_DIR}/bundle.js

TS_MAIN ?= src/main.ts

${OUT_JS}: $(wildcard src/**)
	deno bundle ${TS_MAIN} ${OUT_JS}
