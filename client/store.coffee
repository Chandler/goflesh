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

#TODO: make this generic
# restAdapter.registerTransform 'avatarHash', 
#   serialize: (value) ->
#     return {}
  
#   deserialize: (value) ->
#     return Ember.create({ hash: value["hash"] })
  
App.Store = DS.Store.extend
  adapter: restAdapter