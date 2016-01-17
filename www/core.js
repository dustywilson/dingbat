var s = gotalk.connection().on('open', onConnect).on('close', onDisconnect);

function onSignIn(googleUser) {
  var id_token = googleUser.getAuthResponse().id_token;

  var profile = googleUser.getBasicProfile();
  console.log('ID: ' + profile.getId()); // Do not send to your backend! Use an ID token instead.
  console.log('Name: ' + profile.getName());
  console.log('Image URL: ' + profile.getImageUrl());
  console.log('Email: ' + profile.getEmail());

  s.request("client.init-google", id_token, function (err, result) {
    if (err) {
      console.error('client.init-google failed:', err);
      return signOut(); // or something else?
    }
    $("body").addClass("loggedin");
    console.log(result);
  })
}

gotalk.handleNotification('connection.info', function(c) {
  console.log('connection.info.ID:', c.ID);
});

function signOut() {
  var auth2 = gapi.auth2.getAuthInstance();
  auth2.signOut().then(function () {
    console.log('User signed out.');
    $("body").removeClass("loggedin");
    s.request("client.signout-google", {}, function(err){
      location.reload();
    });
  });
}

function onConnect() {
  //   s.request("person.create", {
  //     firstName: "Fred",
  //     lastName:  "Flintstone"
  //   }, function (err, result) {
  //     if (err) return console.error('create failed:', err);
  //     console.log('person.create result:', result);
  //
  //     s.request("person.get", {ID:result.ID}, function (err, result) {
  //       if (err) return console.error('echo failed:', err);
  //       console.log('person.get result:', result);
  //     });
  //   });
  // });
}

function onDisconnect(err) {
  if (err != null && err.isGotalkProtocolError) return console.error(err);
}
