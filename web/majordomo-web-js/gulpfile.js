var gulp = require('gulp');
var environments = require('gulp-environments');
var tasks = require('./tools/gulp-tasks/');

// Predefined environments from gulp-environments. Environent can be specified
// for task invokation in a number of ways, including: gulp <task> --dev/--prod
// See gulp-environments documentation for additional patterns:
// https://www.npmjs.com/package/gulp-environments 
var development = environments.development;
var production = environments.production;

// Paths and files to remove on clean.
var cleanPaths = [];

// External libraries that don't require any processing. They will be injected into the 
// index template--individually for dev and bundled and minified for prod. 
var libScripts = [];
var libStyles = [];

gulp.task('clean', () => tasks.clean(cleanPaths));
gulp.task('build', function(){});