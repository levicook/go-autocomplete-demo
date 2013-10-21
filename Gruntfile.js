module.exports = function(grunt) {
    grunt.initConfig({
        pkg: grunt.file.readJSON('package.json'),
        browserify: {
            all: {
                files: {
                    'js/pkg/pages/hello/hello.js' : 'js/src/pages/hello/hello.js',
                    'js/pkg/pages/home/home.js'   : 'js/src/pages/home/home.js'
                },
                options: {
                    debug: true
                }
            }
        },
        jslint: {
            main: {
                src: ['js/src/**/*.js'],
                exclude: ['js/src/**/*_test.js'],
                directives: {
                    browser: true,
                    predef: ['console', 'module']
                }
            },
            test: {
                src: ['js/src/**/*_test.js'],
                directives: {
                    browser: true,
                    predef: ['console', 'module']
                }
            }
        },
        watch: {
            files: ['js/src/**/*.js'],
            tasks: [
                'jslint',
                'browserify'
            ]
        }
    });

    grunt.loadNpmTasks('grunt-browserify')
    grunt.loadNpmTasks('grunt-contrib-watch')
    grunt.loadNpmTasks('grunt-jslint');

    grunt.registerTask('default', 'watch');
};
