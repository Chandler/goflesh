App.UserView = Ember.View.extend
  templateName: "user"
  upload: ->
    $("#avatar").trigger('click')    
    $("#avatar").trigger('change')
    $("#avatar").trigger('change')

  didInsertElement: ->
    console.log("inserted")
    #yes I'm checking a key into github sue me.
    token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJwcml2YXRlX3Rva2VuIjoiYjdhZGFkNjI2YWI0NTc0ZmMwYTE5M2MyOWQ0YjBiMTkxYjQxYWNlOWQwMzUzNTU5YzhhMzRhMDhkOWRiNGI2ZCJ9.V0JQ5FygLaWiXn1yzN-gdmywkki6V1V8r2y1TqSMPUU"
    client = new AvatarsIO(token) # obtain at http://avatars.io/
    $ ->
      uploader = client.create("#avatar") # selector for input[type="file"] field, here #avatar, for example
      uploader.setAllowedExtensions ["png", "jpg"] # optional, defaults to png, gif, jpg, jpeg
      uploader.on "complete", (url) ->
        alert url # for example, http://avatars.io/ua3aS5a