module.exports = function(grunt) {
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    clean: {
      assets: ["public/js/*.js", 'tests/client/js/*.js', 'public/js/bower/*.js'],
    },
    bower: {
      dev: {
        dest: 'public/js/bower',
        options: {
          stripJsAffix: true
        }
      }
    },
    concat: {
      basic: {
        src: [
            'public/js/bower/jquery.js',
            'public/js/bower/jquery-cookie.js',
            'public/js/bower/jquery-ui.js',
            'public/js/bower/d3.js',
            'public/js/bower/underscore.js',
            'public/js/bower/handlebars.js',
            'public/js/bower/moment.js',
            'public/js/lib/ember.js',
            'public/js/lib/ember-data.js',
            'public/js/lib/ember-auth.js',
            'public/js/lib/ember-list-view.js',
            'public/js/lib/avatar.io.js'
          ],
        dest: 'public/js/libraries.js',
      }
    },
    uglify: {
      options: {
        mangle: false
      },
      my_target: {
        files: {
          'public/js/libraries.js': [
            'add/some/files/we/will/need/this/later',
          ]
        }
      }
    },
    coffee: { 
      options: {
        bare: true
      },
      app: {
        options: {
          sourceMap: true
        },
        files: {
          'public/js/main.js': [
            'client/app.coffee',
            'client/controllers/frameworkControllers.coffee',
            'client/**/*.coffee'
          ] 
        }
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
  grunt.loadNpmTasks('grunt-shell');
  grunt.loadNpmTasks('grunt-contrib-connect');
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-bower');


  grunt.registerTask('spec', ['coffee:tests', 'jasmine']);
  grunt.registerTask('compile', ['clean:assets', 'bower', 'concat','coffee:app', 'coffee:tests','ember_handlebars','stylus']);
  grunt.registerTask('c', ['compile']);
  grunt.registerTask('w', ['compile','watch']);

};