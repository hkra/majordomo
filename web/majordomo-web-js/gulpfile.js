var gulp = require('gulp');
var environments = require('gulp-environments');
var tasks = require('./tools/gulp-tasks/');

// Load application configuration.
var appConfig = require('./config/appconfig.json');

// Predefined environments from gulp-environments. Environent can be specified
// for task invokation in a number of ways, including: gulp <task> --development
// See gulp-environments documentation for additional patterns:
// https://www.npmjs.com/package/gulp-environments 
var development = environments.development;
var production = environments.production;

// Paths and files to remove on clean.
var cleanPaths = [
    appConfig.outputDir
];

gulp.task('clean', () => tasks.clean(cleanPaths));
gulp.task('build', function(){});