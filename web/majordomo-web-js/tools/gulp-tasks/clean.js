var run = require('./_utils').run

// Task: clean
// @param paths string[] directories or files to clean. 
module.exports = function(paths) {
    paths.forEach(function(element) { run("rm", ['-r', element]); }, this);
}