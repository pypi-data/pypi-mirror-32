(window["webpackJsonp"] = window["webpackJsonp"] || []).push([[4],{

/***/ "../../../node_modules/monaco-editor/esm/vs/editor/common/services sync recursive":
/*!**************************************************************************************************************************!*\
  !*** /Users/kylek/code/src/github.com/nteract/nteract-ext/node_modules/monaco-editor/esm/vs/editor/common/services sync ***!
  \**************************************************************************************************************************/
/*! no static exports found */
/***/ (function(module, exports, __webpack_require__) {

var map = {
	"vs/language/css/cssWorker": "../../../node_modules/monaco-editor/esm/vs/language/css/cssWorker.js",
	"vs/language/html/htmlWorker": "../../../node_modules/monaco-editor/esm/vs/language/html/htmlWorker.js",
	"vs/language/json/jsonWorker": "../../../node_modules/monaco-editor/esm/vs/language/json/jsonWorker.js",
	"vs/language/typescript/tsWorker": "../../../node_modules/monaco-editor/esm/vs/language/typescript/tsWorker.js"
};


function webpackContext(req) {
	var id = webpackContextResolve(req);
	return __webpack_require__(id);
}
function webpackContextResolve(req) {
	var id = map[req];
	if(!(id + 1)) { // check for number or string
		var e = new Error("Cannot find module '" + req + "'");
		e.code = 'MODULE_NOT_FOUND';
		throw e;
	}
	return id;
}
webpackContext.keys = function webpackContextKeys() {
	return Object.keys(map);
};
webpackContext.resolve = webpackContextResolve;
module.exports = webpackContext;
webpackContext.id = "../../../node_modules/monaco-editor/esm/vs/editor/common/services sync recursive";

/***/ })

}]);