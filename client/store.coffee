define ["ember-data"], (DS) ->
  JSONAPIAdapter = DS.RESTAdapter.extend()

  Store = DS.Store.extend
    adapter: JSONAPIAdapter.create({ namespace: 'api' })
