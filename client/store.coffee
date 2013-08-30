#send authorization headers for all logged in users
$.ajaxSetup
  beforeSend: (xhr) ->
    password = App.Auth.get('authToken')
    username = App.Auth.get('userId')
    if(password && username)
      token    = username+":"+password
      xhr.setRequestHeader('Authorization', 'Basic ' + btoa(token))
    


#http://www.thomasboyt.com/2013/05/01/why-ember-data-breaks.html
restAdapter = DS.RESTAdapter.create
  namespace: 'api' 
  serializer: DS.RESTSerializer.createWithMixins
    extract: (loader, json, type, record) ->
      # Conforms ember-data to JSONAPI spec
      # by accepting singular resources in an array
      root = this.rootForType(type) 
      plural = root + "s"
      json[root] = json[plural][0]
      delete json[plural]
      
      @_super(loader, json, type, record)
  
  # Conforms ember-data to JSONAPI spec
  # by posting singular resources in an array
  # TODO mind this when upgrading ember data, it could change.
  createRecord: (store, type, record) ->
    root = this.rootForType(type);
    adapter = this;
    data = {};

    #old version: data[root] = this.serialize(record, { includeId: true });
    data[this.pluralize(root)] = [this.serialize(record, { includeId: true })];

    @ajax(@buildURL(root), "POST",
      data: data
    ).then((json) ->
      adapter.didCreateRecord store, type, record, json
    , (xhr) ->
      adapter.didError store, type, record, xhr
      throw xhr
    ).then null, ->
      DS.rejectionHandler
    

restAdapter.registerTransform 'avatar', 
  serialize: (value) ->
    return {avatar: {hash: value}}
  
  deserialize: (value) ->
    return value["hash"]
  
App.Store = DS.Store.extend
  adapter: restAdapter


DS.RESTAdapter.configure('App.PlayerEvent',
    sideLoadAs: 'events'
)

App.Store.registerAdapter 'App.PlayerEvent', DS.RESTAdapter.extend
  mappings:
    events: App.PlayerEvent
  namespace: 'api'
  customUrl: 'events/players'
  buildURL: (root, suffix, record) ->
    url = [@url]
    url.push @namespace  unless Ember.isNone(@namespace)
    url.push @customUrl
    return url.join("/");



