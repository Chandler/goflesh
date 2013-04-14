module.exports = function(grunt) {
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),

    coffee: {
      compile: {
        files: {
          'public/js/application.js': ['client/**/*.coffee']
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
          "public/js/templates.js": ["client/templates/*.handlebars"]
        }
      }
    },

    stylus: {
      compile: {
        files: {
          'public/css/application.css': ['client/stylesheets/*.styl']
        }
      }
    },

    concat: {
      options: {
        separator: ';'
      },
      dist: {
        src: [ //you must specify these individually because they must be in the correct order
          'public/js/lib/jquery-1.9.1.min.js',
          'public/js/lib/handlebars.runtime.js',
          'public/js/lib/ember.js'
        ],
        dest: 'public/js/vendor.js'
      }
    },

    watch: {
      coffee: {
        files: ['client/*.coffee'],
        tasks: 'coffee'
      },
      handlebars: {
        files: ['client/templates/*.handlebars'],
        tasks: 'handlebars'
      },
      stylus: {
        files: ['client/stylesheets/*.styl'],
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
  grunt.registerTask('w', ['compile','watch']);

};