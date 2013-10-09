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
    "<img class=\" avatar " + options.hash.class +  "\" src=\"http://i.imgur.com/bITe0Cis.png\"/>"

  avatar2Tag: (key, size, klass) ->
    sizes =
      tiny:  25
      small: 50
      large: 100
      profile: 150
    px = sizes[size]
    "<img class=\" avatar " + klass +  "\" src=\"http://www.gravatar.com/avatar/" + key + "?s=" + px + "&d=identicon\"/>"

  #for when we switch to avatar io
  # avatarIOTag: (key, size, klass) ->
  #   "<img class=\" avatar " + klass +  "\" src=\"http://avatars.io/5219420f20885b315500004c/" + key + "?size=" + size + "\"/>"

# idaho http://i.imgur.com/y9yAWbD.jpg