import React, { Component, Fragment } from 'react'
// import { google } from 'googleapis'
import marcel from './marcel'
import { auth } from './client_secret.json'

// const api = (calendarId, tokenID) => `https://www.googleapis.com/calendar/v3/calendars/${calendarId}/events?key=${tokenID}`

export class App extends Component {
  componentDidMount() {
    window.gapi.load('client:auth2', () => {
      window.gapi.auth2.init({
        apiKey: auth.api_key,
        clientId: auth.client_id,
        scope: 'https://www.googleapis.com/auth/calendar.readonly',
      }).then(() => {
        const GoogleAuth = window.gapi.auth2.getAuthInstance();

        // Listen for sign-in state changes.
        GoogleAuth.isSignedIn.listen(this.updateSigninStatus);

        // Handle initial sign-in state. (Determine if user is already signed in.)
        var user = GoogleAuth.currentUser.get();
        this.setSigninStatus();

        // Call handleAuthClick function when user clicks on
        //      "Sign In/Authorize" button.
        // $('#sign-in-or-out-button').click(function() {
        //   this.handleAuthClick(GoogleAuth);
        // });
        // $('#revoke-access-button').click(function() {
        //   this.revokeAccess(GoogleAuth);
        // });
      });
    })
  }

  handleAuthClick = () => {
    const GoogleAuth = window.gapi.auth2.getAuthInstance();
    if (GoogleAuth.isSignedIn.get()) {
      // User is authorized and has clicked 'Sign out' button.
      GoogleAuth.signOut();
    } else {
      // User is not signed in. Start Google auth flow.
      GoogleAuth.signIn();
    }
  }

  revokeAccess = () => {
    const GoogleAuth = window.gapi.auth2.getAuthInstance();
    GoogleAuth.disconnect();
  }

  setSigninStatus = (isSignedIn) => {
    const GoogleAuth = window.gapi.auth2.getAuthInstance();
    var user = GoogleAuth.currentUser.get();
    var SCOPE = 'https://www.googleapis.com/auth/calendar.readonly';
    var isAuthorized = user.hasGrantedScopes(SCOPE);
    if (isAuthorized) {
      console.log('granted')
    } else {
      console.log('failed')
    }
  }

  updateSigninStatus = (isSignedIn) => {
    this.setSigninStatus(isSignedIn);
  }

  render() {
    return (
      <Fragment>
        <button id="sign-in-or-out-button" onClick={() => this.handleAuthClick()}>
          Sign In/Authorize
        </button>
        <button id="revoke-access-button" onClick={() => this.revokeAccess}>Revoke access</button>

        <hr />
      </Fragment>
    )
  }
}

export default marcel(App)
