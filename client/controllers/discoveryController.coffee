define ["ember"], (Em) ->
  DiscoveryController = Ember.ObjectController.extend
    # orgs: [
    #   Ember.Object.create({
    #     id: 3,
    #     name: 'HZ Stanford',
    #     thumb_url: 'https://fbcdn-profile-a.akamaihd.net/hprofile-ak-ash3/41800_6192688417_1783896943_q.jpg',
    #     member_count: 612,
    #     location: 'Palo Alto, CA'
    #   }),
    #   Ember.Object.create({
    #     id: 2,
    #     name: 'HZ Washington State University',
    #     thumb_url: 'https://sphotos-a.xx.fbcdn.net/hphotos-ash4/c356.0.604.604/s320x320/376136_10151643919862704_963634461_n.jpg',
    #     member_count: 387,
    #     location: 'Pullman, Washington'
    #   }),
    #   Ember.Object.create({
    #     id: 1,
    #     name: 'University of Idaho HZ',
    #     thumb_url: 'https://fbcdn-profile-a.akamaihd.net/hprofile-ak-ash3/c0.50.180.180/22176_305111860995_8060208_a.jpg',
    #     member_count: 199,
    #     location: 'Moscow, Idaho',
    #     show: true 
    #   })
    # ],
    orgs: (->
      this.get('content')
    ).property("orgs")

    organizations: (->
      this.get('orgs')
    ).property("orgs.@each")
    something: ->
      filteredOrgs = this.get('orgs').filter (org) ->
        if (!org.get('show'))
          true
      this.set('orgs', filteredOrgs)