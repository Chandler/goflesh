App = Ember.Application.create
  rootElement: "#app"
  LOG_TRANSITIONS: true  
BaseController = Ember.Controller.extend
  errors: null
  
  clearErrors: ->
    @set 'errors', null
  errorMessages: (->
    @get 'errors'
  ).property 'errors'

  fieldsPopulated: ->
    for k,v of @recordProperties
      return false if v == ''
    true

NewController = BaseController.extend
  createTransition: ->
    @transitionToRoute('discovery');
  create: ->
    @clearErrors()
    @recordProperties = @getProperties(@submitFields)
    if @fieldsPopulated()
      model = @get('model')
      record = model.createRecord(@recordProperties)
      record.transaction.commit()
      record.becameError =  =>
        @set 'errors', 'SERVER ERROR'
      record.didCreate = =>
        @createTransition()
    else
      @set 'errors', 'Empty Fields'
App.Auth = Ember.Auth.create
  modules: ['emberData']

  userModel: 'Em.App.User' # default null
  # requestAdapter: 'jquery' # default 'jquery'
  # responseAdapter: 'json' # default 'json'
  # strategyAdapter: 'token' # default 'token'
  signInEndPoint: 'api/users/authenticate'

  tokenKey: 'api_key'
  tokenIdKey: 'id'

  modules: ['emberData', 'authRedirectable', 'actionRedirectable', 'rememberable']

  # authRedirectable:
  #   route: 'sign-in'

  # actionRedirectable:
  #   signInRoute: 'users'
  #   signInSmart: true
  #   signInBlacklist: ['sign-in']
  #   signOutRoute: 'posts'

  # rememberable:
  #   tokenKey: 'test'
  #   period: 7
  #   autoRecall: true
# Might be useful later http://techblog.fundinggates.com/blog/2012/08/ember-handlebars-helpers-bound-and-unbound/
Handlebars.registerHelper 'avatar', (size, options) ->
  key = "doesn't matter yet"
  new Handlebars.SafeString(Utilities.avatarTag(key, size, options))


Utilities =
  avatarTag: (hash, size, options = {}) ->
    sizes =
      tiny:  25
      small: 50
      large: 100
      profile: 150
    px = sizes[size]
    #random for now
    random = Math.random().toString(16).slice(2)
    hash = random + random + random + random
    "<img class=\"" + options.hash.class +  "\" src=\"http://www.gravatar.com/avatar/" + hash + "?s=" + px + "&d=identicon\"/>"


App.ApplicationController = Ember.Controller.extend
  signOut:  ->
    App.Auth.destroySession()
  currentUser: (->
    App.Auth.get('user')
  ).property()
App.DiscoveryController = Ember.ObjectController.extend
  orgs: (->
    string = @get 'filterString'
    if string == ""
      return []
    else
      @get('model').filter (org) ->
        !!(org.get('name').indexOf(string) != -1)
  ).property 'filterString'
  updateFilter: (arg) ->
    @set 'filterString', arg
App.GameHomeController =  Ember.Controller.extend
    needs: 'game'
    game: null
    gameBinding: 'controllers.game'
    contentBinding: 'game.players'

App.GameSettingsController = BaseController.extend
  needs: 'game'
  game: null
  gameBinding: 'controllers.game'
  edit: ->
    this.clearErrors()
    if @get('game.name') != ''
      record = @get('game.content')
      @get('store').get('defaultTransaction').commit()
      record.on 'becameError', =>
        @set 'errors', 'SERVER ERROR'
      record.on 'didUpdate', =>
        @transitionTo('game.home', record);
    else
      @set 'errors', 'Empty Field'

App.GamesController = Ember.ArrayController.extend
  selectedGame: null
App.GamesNewController = NewController.extend
  submitFields: ['name', 'slug']
  name: '',
  slug: '',

App.LoginController = BaseController.extend
  email: null
  password: null
  login: (arg) ->
    @clearErrors()
    App.Auth.signIn
      data:
        email: @email
        password: @password

    App.Auth.on 'signInSuccess', =>
      @transitionTo('user.home', App.Auth.get('user'))

App.OrganizationHomeController = Ember.Controller.extend
  needs: 'organization'
  organization: null
  organizationBinding: 'controllers.organization'

App.OrganizationSettingsController = BaseController.extend
  needs: 'organization'
  organization: null
  organizationBinding: 'controllers.organization'
  edit: ->
    this.clearErrors()
    if @get('organization.name') != ''
      record = @get('organization.content')
      @get('store').get('defaultTransaction').commit()
      record.on 'becameError', =>
        @set 'errors', 'SERVER ERROR'
      record.on 'didUpdate', =>
        @transitionTo('organization.home', record);
    else
      @set 'errors', 'Empty Field'
App.OrganizationsController = Ember.ArrayController.extend
  selectedOrganization: null
App.OrganizationsNewController = NewController.extend
  submitFields: ['name', 'slug', 'location']
  name: ''
  slug: ''
  location: ''
App.UserHomeController = Ember.Controller.extend
  needs: 'user'
  user: null
  userBinding: 'controllers.user'
App.UserSettingsController = BaseController.extend
  needs: 'user'
  user: null
  userBinding: 'controllers.user'
  edit: ->
    this.clearErrors()
    if @get('user.name') != ''
      record = @get('user.content')
      @get('store').get('defaultTransaction').commit()
      record.on 'becameError', =>
        @set 'errors', 'SERVER ERROR'
      record.on 'didUpdate', =>
        @transitionTo('user.home', record);
    else
      @set 'errors', 'Empty Field'

App.UsersController = Ember.ArrayController.extend
  selectedUser: null

App.UsersNewController = NewController.extend
  submitFields: ['first_name', 'last_name', 'email', 'screen_name', 'password']
  first_name: '',
  last_name: '',
  email: '',
  screen_name: '',
  password: ''
App.Game = DS.Model.extend
  name: DS.attr 'string'
  slug: DS.attr 'string'
  organization: DS.belongsTo 'App.Organization'
  players: DS.hasMany 'App.Player'

App.Game.toString = -> 
  "Game"
App.Member = DS.Model.extend
  organization: DS.belongsTo 'App.Organization'

App.Member.toString = -> 
  "Member"
App.Organization = DS.Model.extend
  name: DS.attr 'string'
  slug: DS.attr 'string'
  location: DS.attr 'string'
  games: DS.hasMany 'App.Game',
    inverse: 'organization'

App.Organization.toString = -> 
  "Organization"
App.Player = DS.Model.extend
  game: DS.belongsTo 'App.Game'
  user: DS.belongsTo 'App.User'

App.Player.toString = -> 
  "Player"



App.User = DS.Model.extend
  first_name: DS.attr 'string'
  last_name: DS.attr 'string'
  screen_name: DS.attr 'string'
  email: DS.attr 'string'
  password: DS.attr 'string'
  player: DS.belongsTo 'App.Player'

App.User.toString = -> 
  "User"
#http://darthdeus.github.io/blog/2013/02/01/ember-dot-js-router-and-template-naming-convention/

App.Router = Ember.Router.extend
  enableLogging: true
  location: 'history'

App.Router.map ->
  @route 'discovery' 
  @resource 'organizations', path: "/orgs", ->
    @resource 'organization', path: ":organization_id", ->
      @route 'settings'
      @route 'home'
    @route 'new'
  @resource 'games', ->
    @resource 'game', path: ":game_id", ->
      @route 'settings'
      @route 'home'
    @route 'new'
  @resource 'users', ->
    @resource 'user', path: ":user_id", ->
      @route 'home'
      @route 'settings'
  @route 'users.new', path: "/signup"
  @route 'login'

App.DiscoveryRoute = Ember.Route.extend
  model: ->
    App.Organization.find()
    
  setupController: (controller, model) ->
    @controller.set('filterString', '')

App.GameRoute = Ember.Route.extend
  model: (params) ->
    App.Game.find(params.game_id)

  setupController: (controller, model) ->
    @_super arguments...
    @controllerFor('games').set 'selectedGame', model

App.GamesRoute = Ember.Route.extend
  model: ->
    App.Game

# App.IndexRoute = Ember.Route.extend(redirect: ->
#   @transitionTo 'discovery'
# )
App.OrganizationRoute = Ember.Route.extend
  model: (params) ->
    App.Organization.find(params.organization_id)

  setupController: (controller, model) ->
    @_super arguments...
    @controllerFor('organizations').set 'selectedOrganization', model

App.DiscoveryRoute = Ember.Route.extend
  model: ->
    App.Organization

App.UserRoute = Ember.Route.extend
  model: (params) ->
    App.User.find(params.user_id)

  setupController: (controller, model) ->
    @_super arguments...
    @controllerFor('user').set 'selectedUser', model

App.UsersRoute = Ember.Route.extend
  model: ->
    App.User

#http://www.thomasboyt.com/2013/05/01/why-ember-data-breaks.html
restAdapter = DS.RESTAdapter.create
  namespace: 'api' 
  serializer: DS.JSONSerializer.createWithMixins
    extract: (loader, json, type, record) ->
      # Conforms ember-data to JSONAPI spec
      # by accepting singular resources in an array
      root = this.rootForType(type)
      plural = root + "s"
      json[root] = json[plural][0]
      delete json[plural]
      
      @_super(loader, json, type, record)

App.Store = DS.Store.extend
  adapter: restAdapter
App.ApplicationView = Ember.View.extend
  didInsertElement: ->
   console.log "application view rendered"

App.DiscoveryView = Ember.View.extend
  templateName: "discovery"



App.GameItemView = Ember.View.extend
  templateName: "gameItem"

App.ListItemView = Ember.View.extend
  templateName: "listItem"
  didInsertElement: ->
    this.$().hide()
    this.$().fadeIn(100)

App.LoginView = Ember.View.extend
  templateName: "login"
  login: ->
    console.log "test"

App.ApplicationView = Ember.View.extend
  templateName: 'playerTableRow'

App.ApplicationView = Ember.View.extend
  templateName: 'registerKill'
App.TimeSeriesView = Ember.View.extend
  templateName: 'graph',

  