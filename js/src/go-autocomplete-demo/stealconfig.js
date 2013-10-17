steal.config({
    map: {
        '*': {
            'jquery/jquery.js' : 'jquery'
        }
    },
    paths: {
        jquery: 'js/pkg/jquery/jquery.min.js'
    },
    shim: {
        jquery: {
            exports  : 'jQuery',
            ignore   : true,
            packaged : false
        }
    },
    ext: {
        js       : 'js',
        css      : 'css',
        less     : 'steal/less/less.js',
        mustache : 'can/view/mustache/mustache.js'
    }
});
