#send authorization headers in all Ajax calls when user is logged in.
$.ajaxSetup
  beforeSend: (xhr) ->
    password = App.Auth.get('authToken')
    username = App.Auth.get('userId')
    if(password && username)
      token    = username+":"+password
      xhr.setRequestHeader('Authorization', 'Basic ' + btoa(token))
    

#http://www.thomasboyt.com/2013/05/01/why-ember-data-breaks.html

#Custom extensions of the ember rest adapter. Review these changes when upgrading ember-data
#works with ember-data commit: ef11bff (2013-08-26 20:54:06 -0700)
FleshRestAdapter = DS.RESTAdapter.extend
  namespace: 'api' 
  serializer: DS.RESTSerializer.createWithMixins
    extract: (loader, json, type, record) ->
      # Conforms ember-data to JSONAPI spec
      # by *accepting* singular resources in an array
      # with a plural key
      root = this.rootForType(type) 
      plural = this.pluralize(root)
      json[root] = json[plural][0]
      delete json[plural]
      
      @_super(loader, json, type, record)
  
  # Conforms ember-data to JSONAPI spec
  # by *sending* singular resources in an array
  # with a pural key
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

  didError: (store, type, record, xhr) ->
    series = xhr.status.toString()[0]
    if (series == "4" or series == "5")
      errors = xhr.responseText
      store.recordWasInvalid record, errors
    else
      @_super.apply this, arguments_

#custom attribute type for the rest adapter
FleshRestAdapter.registerTransform 'avatar', 
  serialize: (value) ->
    {avatar: {hash: value}}
  
  deserialize: (value) ->
    value["hash"]

FleshRestAdapter.registerTransform 'isodate', 
  serialize: (value) ->
    null
  
  deserialize: (value) ->
    moment(value).format('MMM Do ha')
  
FleshRestAdapter.map 'App.Event',
  tag: { embedded: 'always' }

App.Store = DS.Store.extend
  adapter: FleshRestAdapter

