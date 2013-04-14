module.exports = function(grunt) {
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),

    coffee: {
      compile: {
        files: {
          'flesh/public/js/application.js': ['flesh/client/*.coffee']
        }
      }
    },

    handlebars: {
      options: {
        processName: function(filename) {
          return filename.split('/').pop().split('.')[0] // /an/annoying/path/name.handlebars => name
        }
      },
      compile: {
        files: {
          "flesh/public/js/templates.js": ["flesh/client/templates/*.handlebars"]
        }
      }
    },

    stylus: {
      compile: {
        files: {
          'flesh/public/css/application.css': ['flesh/client/stylesheets/*.styl']
        }
      }
    },

    concat: {
      options: {
        separator: ';'
      },
      dist: {
        src: [ //you must specify these individually because they must be in the correct order
          'flesh/public/js/lib/jquery-1.9.1.min.js',
          'flesh/public/js/lib/handlebars.runtime.js',
          'flesh/public/js/lib/ember.js'
        ],
        dest: 'flesh/public/js/vendor.js'
      }
    },

    watch: {
      coffee: {
        files: ['flesh/client/*.coffee'],
        tasks: 'coffee'
      },
      handlebars: {
        files: ['flesh/client/templates/*.handlebars'],
        tasks: 'handlebars'
      },
      stylus: {
        files: ['flesh/client/stylesheets/*.styl'],
        tasks: 'handlebars'
      }
    },

    shell: {
      run: {
        command: 'revel run flesh'
      }
    }

  });

  grunt.loadNpmTasks('grunt-contrib-handlebars');
  grunt.loadNpmTasks('grunt-contrib-coffee');
  grunt.loadNpmTasks('grunt-contrib-stylus');
  grunt.loadNpmTasks('grunt-contrib-concat');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-shell');

  grunt.registerTask('compile', ['coffee', 'handlebars', 'stylus']);
  grunt.registerTask('server', ['watch','shell']);

};