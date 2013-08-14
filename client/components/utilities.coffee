define [], ->
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

