module.exports = function(grunt) {
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    clean: {
      assets: ["public/js/*.js", 'tests/client/js/*.js'],
    },
    coffee: {
      options: {
        bare: true
      },
      app: {
        flatten: true,
        expand: true,
        cwd: 'client',
        src: ['**/*.coffee'],
        dest: 'public/js/',
        ext: '.js'
      },
      tests: {
        flatten: true,
        expand: true,
        cwd: 'tests/client',
        src: ['*.coffee'],
        dest: 'tests/client/js/',
        ext: '.js'
      }
    },
    ember_handlebars: {
      options: {
        processName: function(filename) {
        /* example template structure
            templates/
              widget.handlebars
              posts/
                posts.handlebars
                index.handlebars
                new.handlebars
                comments/
                  new.handlebars
                  edit.handlebars

            resulting ember handlebar templates.
              TEMPLATES['widget']
              TEMPLATES['posts']
              TEMPLATES['post/index']
              TEMPLATES['post/new']
              TEMPLATES['comments/new']
              TEMPLATES['comments/edit']
          */

          split_path = filename.split('/')
          resource_name = split_path[split_path.length-2]
          template_name = split_path.pop().split('.')[0]

          if(resource_name == template_name || resource_name == 'templates') {
            return template_name
          } else {
            return resource_name + "/" + template_name
          }
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
        files: ['client/**/*.coffee', 'tests/client/**/*.coffee'],
        tasks: 'coffee'
      },
      ember_handlebars: {
        files: ['client/templates/**/*.handlebars'],
        tasks: 'ember_handlebars'
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
  grunt.loadNpmTasks('grunt-contrib-jasmine');
  grunt.loadNpmTasks('grunt-shell');
  grunt.loadNpmTasks('grunt-contrib-connect');
  grunt.loadNpmTasks('grunt-madge');
  
  grunt.loadNpmTasks('grunt-browserify');

  grunt.registerTask('spec', ['coffee:tests', 'jasmine']);
  grunt.registerTask('compile', ['clean:assets','coffee:app', 'coffee:tests','ember_handlebars','stylus']);
  grunt.registerTask('c', ['compile']);
  grunt.registerTask('w', ['compile','watch']);

};