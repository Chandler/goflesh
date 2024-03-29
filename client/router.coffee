#http://darthdeus.github.io/blog/2013/02/01/ember-dot-js-router-and-template-naming-convention/

App.Router = Ember.Router.extend
  enableLogging: true
  location: 'history'

App.Router.map ->
  @route 'discovery'
  @resource 'organizations', path: "/orgs", ->
    @resource 'organization', path: ":organization_id", ->
      #TODO: unhide settings when moderators are implemented
      # @route 'settings'
      @route 'home'
    @route 'new'
  @resource 'games', ->
    @resource 'game', path: ":game_id", ->
      #TODO: unhide settings when moderators are implemented
      # @route 'settings'
      @route 'home'
      @route 'players'
    @route 'new'
  @resource 'users', ->
    @resource 'user', path: ":user_id", ->
      @route 'home'
      @route 'settings'
  @route 'users.new', path: "/signup"
  @route 'login'
