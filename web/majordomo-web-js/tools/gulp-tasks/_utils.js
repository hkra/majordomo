// Run a shell command
// http://stackoverflow.com/questions/14458508/node-js-shell-command-execution
var run_command = function(cmd, args, callBack ) {
    var spawn = require('child_process').spawn;
    var child = spawn(cmd, args);
    var resp = "";

    child.stdout.on('data', function (buffer) { resp += buffer.toString() });
    child.stdout.on('end', function() {
        if (typeof(callBack) !== 'undefined') {
            callBack(resp);
        }
    });
}

module.exports = {
    run: run_command
};