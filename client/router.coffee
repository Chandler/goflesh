#http://darthdeus.github.io/blog/2013/02/01/ember-dot-js-router-and-template-naming-convention/

define ["ember"], (Em) ->
  Router = Em.Router.extend
    enableLogging: true
    location: 'history'

  Router.map ->
    @route 'discovery'
    @resource 'organizations', path: "/orgs", ->
      @route 'show', path: ":organization_id", ->
      @route 'new'
    # @route 'signup'
    @resource 'users', ->
      @route 'show', path: ":user_id", ->
      @route 'new', path: "/signup", ->
      @route 'edit', path: ":user_id/edit"
  Router