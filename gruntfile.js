module.exports = function(grunt) {
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    clean: {
      assets: ["public/js/*.js", "public/css/*.css"],
    },
    coffee: {
      glob_to_multiple: {
        options: {
          bare: true
        },
        flatten: true,
        expand: true,
        cwd: 'client',
        src: ['**/*.coffee'],
        dest: 'public/js/',
        ext: '.js'
      }
    },
    ember_handlebars: {
      options: {
        amd: true,
        processName: function(filename) {
          return filename.split('/').pop().split('.')[0] // /an/annoying/path/name.handlebars => name
        }
      },
      compile: {
        files: {
          "public/js/templates.js": ["client/templates/**/*.handlebars"]
        }
      }
    },
    stylus: {
      compile: {
        files: {
          'public/css/application.css': ['client/stylesheets/**/*.styl']
        }
      }
    },
    watch: {
      coffee: {
        files: ['client/**/*.coffee'],
        tasks: 'coffee'
      },
      ember_handlebars: {
        files: ['client/templates/**/*.handlebars'],
        tasks: 'handlebars'
      },
      stylus: {
        files: ['client/stylesheets/*.styl'],
        tasks: 'stylus'
      }
    }
  });

  grunt.loadNpmTasks('grunt-contrib-handlebars');
  grunt.loadNpmTasks('grunt-ember-handlebars');
  grunt.loadNpmTasks('grunt-contrib-coffee');
  grunt.loadNpmTasks('grunt-contrib-stylus');
  grunt.loadNpmTasks('grunt-contrib-concat');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-contrib-clean');
  grunt.loadNpmTasks('grunt-shell');
  grunt.loadNpmTasks('grunt-requirejs');
  grunt.loadNpmTasks('grunt-amd-dist');


  grunt.registerTask('compile', ['clean:assets', 'coffee', 'ember_handlebars','stylus']);
  grunt.registerTask('c', ['compile']);
  grunt.registerTask('w', ['compile','watch']);

};