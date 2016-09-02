var gulp = require('gulp');
var environments = require('gulp-environments');

var bundleTypes = {
    ScriptBundle: "scripts",
    StyleBundle: "styles"
};

/**
 * Not sure...didn't really think this through very well.
 */
function bundle(bundleType, inputPaths) {
    if (typeof(bundleType) === 'undefined') { throw "Bundle type must be specified"; }
    if (typeof(inputPaths) === 'undefined') { throw "Input paths type must be specified"; }
    if (!Array.isArray(inputPaths)) { throw "Input paths should be an array of strings"; }
}