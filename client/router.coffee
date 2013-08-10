#http://darthdeus.github.io/blog/2013/02/01/ember-dot-js-router-and-template-naming-convention/

define ["ember"], (Em) ->
  Router = Em.Router.extend
    enableLogging: true
    location: 'history'

  Router.map ->
    @route 'discovery' 
    @resource 'organizations', path: "/orgs", ->
      @resource 'organization', path: ":organization_id", ->
        @route 'settings'
        @route 'home'
      @route 'new'
    @resource 'games', ->
      @resource 'game', path: ":game_id/edit", ->
        @route 'settings'
        @route 'home'
      @route 'new'
    @resource 'users', ->
      @resource 'user', path: ":user_id", ->
        @route 'home', path: ":user_id"
        @route 'settings', path: ":user_id/settings"
    @route 'users.new', path: "/signup"
    @route 'login'
  Router