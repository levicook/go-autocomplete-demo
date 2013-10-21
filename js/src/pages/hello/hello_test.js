
var hello = require('./hello'),
    exports = module.exports;

exports.testSomething = function (test) {
    'use strict';
    test.equal(hello(4, 2), 8, "expected 4");
    test.done();
};
