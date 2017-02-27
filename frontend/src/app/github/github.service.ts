import { SortPipe } from './../pipes/sort.pipe';
import { Injectable } from '@angular/core';
import { Http } from '@angular/http';
import { Observable } from 'rxjs/Rx';
import { ApiService } from './../api/api.service';

@Injectable()
export class GithubService {

  api: string = 'https://api.github.com/users/';
  params: string = '/events?per_page=100';
  users: any[] = [];
  usersContribusion: any[] = [];
  client_auth : string = "";

  constructor(private http: Http, private sort: SortPipe,private apiService : ApiService) {
    this.users.push(
      'Gillespie59',
      'GwennaelBuchet',
      'T3kstiil3',
      'RemiEven',
      'looztra',
      'a-cordier',
      'wadendo',
      'NathanDM',
      'Antoinephi',
      'cluster',
      'yyekhlef',
      'gdrouet',
      'Kize',
      'kratisto'
    );

    this.client_auth = 'client_id='+apiService.getApi('github').client_id+'&client_secret='+apiService.getApi('github').client_secret;
    console.log(this.client_auth);
  }

  getTopContributors() {
    return Observable
      .from(this.users)
      .flatMap(user => this.getUserInfo(user))
      .flatMap(user => this.getUserContributions(user))
      .map(contribsAndUser => {
        let nbContributions = 0;
        let contribs = contribsAndUser.contribs;

        contribs
          .filter(contrib => contrib.type === 'PushEvent' && contrib.created_at.indexOf(this.getDate()) !== -1)
          .map(contri => nbContributions += contri.payload.commits.length);

        this.usersContribusion = this.usersContribusion.filter(contrib => {
          return contrib.user !== contribsAndUser.user.name;
        });

        this.usersContribusion.push({
          user: contribsAndUser.user.name,
          contributions: nbContributions,
          avatar: contribs[0].actor.avatar_url
        });

        return this.sort.transform(this.usersContribusion, 'contributions');
      });
  }

  getUserInfo(user: string) {
    return this.http.get(`${this.api}${user}`+"?"+this.client_auth)
      .map(res => {
        let response = res.json();
        return {
          login: user,
          name: response['name']
        }
      });
  }

  getUserContributions(user) {
    return this.http
      .get(this.api + user.login + this.params+"&"+this.client_auth)
      .map(res => {
        return {
          contribs: res.json(),
          user
        };
      });
  }

  private getDate() {
    let today = new Date();
    return `${today.getFullYear()}-${("0" + (today.getMonth() + 1)).slice(-2)}`;
  }

}
