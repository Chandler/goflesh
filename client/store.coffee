#http://www.thomasboyt.com/2013/05/01/why-ember-data-breaks.html
define ["ember-data"], (DS) ->
  DS.RESTAdapter.map 'Em.App.Game',
    organization:
      embedded: 'always'
    players:
      embedded: 'load'
  
  DS.RESTAdapter.map 'Em.App.Organization',
    games:
      embedded: 'always'


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

  Store = DS.Store.extend
    adapter: restAdapter