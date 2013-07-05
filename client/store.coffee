
define ["ember-data"], (DS) ->
  Store = DS.Store.extend
    adapter: DS.RESTAdapter.create
      namespace: 'api' 
      serializer: DS.JSONSerializer.createWithMixins
        extract: (loader, json, type, record) ->
          # Conforms ember-data to JSONAPI spec
          # by accepting singular resources in an array
          root = this.rootForType(type)
          json[root] = json[root][0]
          @_super(loader, json, type, record)

          